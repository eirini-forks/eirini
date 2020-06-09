package v1

import (
	"encoding/json"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LRP describes an Long Running Process
type LRP struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec LRPSpec `json:"spec"`
}

type LRPSpec struct {
	GUID                    string                     `json:"GUID"`
	Version                 string                     `json:"version"`
	ProcessGUID             string                     `json:"processGUID"`
	ProcessType             string                     `json:"processType"`
	AppGUID                 string                     `json:"appGUID"`
	AppName                 string                     `json:"appName"`
	SpaceGUID               string                     `json:"spaceGUID"`
	SpaceName               string                     `json:"spaceName"`
	OrganizationGUID        string                     `json:"organizationGUID"`
	OrganizationName        string                     `json:"organizationName"`
	PlacementTags           []string                   `json:"placementTags"`
	Ports                   []int32                    `json:"ports"`
	Routes                  map[string]json.RawMessage `json:"routes"`
	Environment             map[string]string          `json:"environment"`
	EgressRules             []json.RawMessage          `json:"egressRules"`
	NumInstances            int                        `json:"instances"`
	LastUpdated             string                     `json:"lastUpdated"`
	HealthCheckType         string                     `json:"healthCheckType"`
	HealthCheckHTTPEndpoint string                     `json:"healthCheckHttpEndpoint"`
	HealthCheckTimeoutMs    uint                       `json:"healthCheckTimeoutMs"`
	StartTimeoutMs          uint                       `json:"startTimeoutMs"`
	MemoryMB                int64                      `json:"memoryMB"`
	DiskMB                  int64                      `json:"diskMB"`
	CPUWeight               uint8                      `json:"cpuWeight"`
	VolumeMounts            []VolumeMount              `json:"volumeMounts"`
	Lifecycle               Lifecycle                  `json:"lifecycle"`
	DropletHash             string                     `json:"dropletHash"`
	DropletGUID             string                     `json:"dropletGUID"`
	StartCommand            string                     `json:"startCommand"`
	UserDefinedAnnotations  map[string]string          `json:"userDefinedAnnotations"`
}

type Lifecycle struct {
	DockerLifecycle    *DockerLifecycle    `json:"docker"`
	BuildpackLifecycle *BuildpackLifecycle `json:"buildpack"`
}

type DockerLifecycle struct {
	Image            string   `json:"image"`
	Command          []string `json:"command"`
	RegistryUsername string   `json:"registryUsername"`
	RegistryPassword string   `json:"registryPassword"`
}

type BuildpackLifecycle struct {
	DropletHash  string `json:"dropletHash"`
	DropletGUID  string `json:"dropletGuid"`
	StartCommand string `json:"startCommand"`
}

type VolumeMount struct {
	VolumeID string `json:"volumeID"`
	MountDir string `json:"mountDir"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LRPList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []LRP `json:"items"`
}