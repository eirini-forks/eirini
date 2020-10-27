// Code generated by counterfeiter. DO NOT EDIT.
package reconcilerfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s/reconciler"
	v1 "k8s.io/api/core/v1"
	v1a "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FakeEventsClient struct {
	CreateStub        func(string, *v1.Event) (*v1.Event, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 *v1.Event
	}
	createReturns struct {
		result1 *v1.Event
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1.Event
		result2 error
	}
	GetByInstanceAndReasonStub        func(string, v1a.OwnerReference, int, string) (*v1.Event, error)
	getByInstanceAndReasonMutex       sync.RWMutex
	getByInstanceAndReasonArgsForCall []struct {
		arg1 string
		arg2 v1a.OwnerReference
		arg3 int
		arg4 string
	}
	getByInstanceAndReasonReturns struct {
		result1 *v1.Event
		result2 error
	}
	getByInstanceAndReasonReturnsOnCall map[int]struct {
		result1 *v1.Event
		result2 error
	}
	UpdateStub        func(string, *v1.Event) (*v1.Event, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 *v1.Event
	}
	updateReturns struct {
		result1 *v1.Event
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1.Event
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventsClient) Create(arg1 string, arg2 *v1.Event) (*v1.Event, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 *v1.Event
	}{arg1, arg2})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventsClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeEventsClient) CreateCalls(stub func(string, *v1.Event) (*v1.Event, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeEventsClient) CreateArgsForCall(i int) (string, *v1.Event) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventsClient) CreateReturns(result1 *v1.Event, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) CreateReturnsOnCall(i int, result1 *v1.Event, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1.Event
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) GetByInstanceAndReason(arg1 string, arg2 v1a.OwnerReference, arg3 int, arg4 string) (*v1.Event, error) {
	fake.getByInstanceAndReasonMutex.Lock()
	ret, specificReturn := fake.getByInstanceAndReasonReturnsOnCall[len(fake.getByInstanceAndReasonArgsForCall)]
	fake.getByInstanceAndReasonArgsForCall = append(fake.getByInstanceAndReasonArgsForCall, struct {
		arg1 string
		arg2 v1a.OwnerReference
		arg3 int
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetByInstanceAndReasonStub
	fakeReturns := fake.getByInstanceAndReasonReturns
	fake.recordInvocation("GetByInstanceAndReason", []interface{}{arg1, arg2, arg3, arg4})
	fake.getByInstanceAndReasonMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventsClient) GetByInstanceAndReasonCallCount() int {
	fake.getByInstanceAndReasonMutex.RLock()
	defer fake.getByInstanceAndReasonMutex.RUnlock()
	return len(fake.getByInstanceAndReasonArgsForCall)
}

func (fake *FakeEventsClient) GetByInstanceAndReasonCalls(stub func(string, v1a.OwnerReference, int, string) (*v1.Event, error)) {
	fake.getByInstanceAndReasonMutex.Lock()
	defer fake.getByInstanceAndReasonMutex.Unlock()
	fake.GetByInstanceAndReasonStub = stub
}

func (fake *FakeEventsClient) GetByInstanceAndReasonArgsForCall(i int) (string, v1a.OwnerReference, int, string) {
	fake.getByInstanceAndReasonMutex.RLock()
	defer fake.getByInstanceAndReasonMutex.RUnlock()
	argsForCall := fake.getByInstanceAndReasonArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEventsClient) GetByInstanceAndReasonReturns(result1 *v1.Event, result2 error) {
	fake.getByInstanceAndReasonMutex.Lock()
	defer fake.getByInstanceAndReasonMutex.Unlock()
	fake.GetByInstanceAndReasonStub = nil
	fake.getByInstanceAndReasonReturns = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) GetByInstanceAndReasonReturnsOnCall(i int, result1 *v1.Event, result2 error) {
	fake.getByInstanceAndReasonMutex.Lock()
	defer fake.getByInstanceAndReasonMutex.Unlock()
	fake.GetByInstanceAndReasonStub = nil
	if fake.getByInstanceAndReasonReturnsOnCall == nil {
		fake.getByInstanceAndReasonReturnsOnCall = make(map[int]struct {
			result1 *v1.Event
			result2 error
		})
	}
	fake.getByInstanceAndReasonReturnsOnCall[i] = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) Update(arg1 string, arg2 *v1.Event) (*v1.Event, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 *v1.Event
	}{arg1, arg2})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventsClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeEventsClient) UpdateCalls(stub func(string, *v1.Event) (*v1.Event, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeEventsClient) UpdateArgsForCall(i int) (string, *v1.Event) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventsClient) UpdateReturns(result1 *v1.Event, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) UpdateReturnsOnCall(i int, result1 *v1.Event, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1.Event
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getByInstanceAndReasonMutex.RLock()
	defer fake.getByInstanceAndReasonMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventsClient) recordInvocation(key string, args []interface{}) {
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

var _ reconciler.EventsClient = new(FakeEventsClient)
