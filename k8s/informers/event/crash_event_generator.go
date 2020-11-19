package event

import (
	"strconv"

	"code.cloudfoundry.org/eirini/events"
	"code.cloudfoundry.org/eirini/k8s"
	"code.cloudfoundry.org/eirini/util"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
	v1 "k8s.io/api/core/v1"
)

type DefaultCrashEventGenerator struct {
	eventsClient k8s.EventsClient
}

func NewDefaultCrashEventGenerator(eventsClient k8s.EventsClient) DefaultCrashEventGenerator {
	return DefaultCrashEventGenerator{
		eventsClient: eventsClient,
	}
}

func (g DefaultCrashEventGenerator) Generate(pod *v1.Pod, logger lager.Logger) (events.CrashEvent, bool) {
	logger = logger.Session("generate-crash-event",
		lager.Data{
			"pod-name": pod.Name,
			"guid":     pod.Annotations[k8s.AnnotationProcessGUID],
			"version":  pod.Annotations[k8s.AnnotationVersion],
		})

	statuses := pod.Status.ContainerStatuses
	if len(statuses) == 0 {
		logger.Debug("skipping-empty-container-statuseses")

		return events.CrashEvent{}, false
	}

	if pod.Labels[k8s.LabelSourceType] != k8s.AppSourceType {
		logger.Debug("skipping-non-eirini-pod")

		return events.CrashEvent{}, false
	}

	appStatus := getOPIContainerStatus(pod.Status.ContainerStatuses)
	if appStatus == nil {
		logger.Debug("skipping-eirini-pod-has-no-opi-container-statuses")

		return events.CrashEvent{}, false
	}

	lastTerminatedEventSent := pod.Annotations[k8s.AnnotationLastReportedAppCrash]
	if appStatus.State.Terminated != nil {
		if lastTerminatedEventSent == strconv.FormatInt(appStatus.State.Terminated.FinishedAt.Unix(), 10) {
			return events.CrashEvent{}, false
		}

		return g.generateReportForTerminatedPod(pod, appStatus, logger)
	}

	if appStatus.LastTerminationState.Terminated != nil {
		if lastTerminatedEventSent == strconv.FormatInt(appStatus.LastTerminationState.Terminated.FinishedAt.Unix(), 10) {
			return events.CrashEvent{}, false
		}

		exitStatus := int(appStatus.LastTerminationState.Terminated.ExitCode)
		exitDescription := appStatus.LastTerminationState.Terminated.Reason
		crashTimestamp := appStatus.LastTerminationState.Terminated.FinishedAt.Unix()

		return generateReport(pod, appStatus.LastTerminationState.Terminated.Reason, exitStatus, exitDescription, crashTimestamp, calculateCrashCount(appStatus)), true
	}

	logger.Debug("skipping-pod-healthy")

	return events.CrashEvent{}, false
}

func (g DefaultCrashEventGenerator) generateReportForTerminatedPod(pod *v1.Pod, status *v1.ContainerStatus, logger lager.Logger) (events.CrashEvent, bool) {
	podEvents, err := g.eventsClient.GetByPod(*pod)
	if err != nil {
		logger.Error("skipping-failed-to-get-k8s-events", err)

		return events.CrashEvent{}, false
	}

	if k8s.IsStopped(podEvents) {
		logger.Debug("skipping-pod-stopped")

		return events.CrashEvent{}, false
	}

	terminated := status.State.Terminated

	return generateReport(pod, terminated.Reason, int(terminated.ExitCode), terminated.Reason, terminated.FinishedAt.Unix(), calculateCrashCount(status)), true
}

func generateReport(
	pod *v1.Pod,
	reason string,
	exitStatus int,
	exitDescription string,
	crashTimestamp int64,
	restartCount int,
) events.CrashEvent {
	index, _ := util.ParseAppIndex(pod.Name)

	return events.CrashEvent{
		ProcessGUID: pod.Annotations[k8s.AnnotationProcessGUID],
		AppCrashedRequest: cc_messages.AppCrashedRequest{
			Reason:          reason,
			Instance:        pod.Name,
			Index:           index,
			ExitStatus:      exitStatus,
			ExitDescription: exitDescription,
			CrashTimestamp:  crashTimestamp,
			CrashCount:      restartCount,
		},
	}
}

func getOPIContainerStatus(statuses []v1.ContainerStatus) *v1.ContainerStatus {
	for _, status := range statuses {
		if status.Name == k8s.OPIContainerName {
			return &status
		}
	}

	return nil
}

func calculateCrashCount(containerState *v1.ContainerStatus) int {
	// warning: apparently the RestartCount is limited to 5 by K8s Garbage
	// Collection. However, we have observed it at 6 at least!

	// If container is running, the restart count will be the crash count.  If
	// container is terminated or waiting, we need to add 1, as it has not yet
	// been restarted

	if containerState.State.Running != nil {
		return int(containerState.RestartCount)
	}

	return int(containerState.RestartCount + 1)
}
