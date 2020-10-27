// Code generated by counterfeiter. DO NOT EDIT.
package routefakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s/informers/route"
	v1 "k8s.io/api/apps/v1"
)

type FakeStatefulSetUpdateEventHandler struct {
	HandleStub        func(*v1.StatefulSet, *v1.StatefulSet)
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
		arg1 *v1.StatefulSet
		arg2 *v1.StatefulSet
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatefulSetUpdateEventHandler) Handle(arg1 *v1.StatefulSet, arg2 *v1.StatefulSet) {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
		arg1 *v1.StatefulSet
		arg2 *v1.StatefulSet
	}{arg1, arg2})
	stub := fake.HandleStub
	fake.recordInvocation("Handle", []interface{}{arg1, arg2})
	fake.handleMutex.Unlock()
	if stub != nil {
		fake.HandleStub(arg1, arg2)
	}
}

func (fake *FakeStatefulSetUpdateEventHandler) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeStatefulSetUpdateEventHandler) HandleCalls(stub func(*v1.StatefulSet, *v1.StatefulSet)) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeStatefulSetUpdateEventHandler) HandleArgsForCall(i int) (*v1.StatefulSet, *v1.StatefulSet) {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	argsForCall := fake.handleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatefulSetUpdateEventHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatefulSetUpdateEventHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ route.StatefulSetUpdateEventHandler = new(FakeStatefulSetUpdateEventHandler)
