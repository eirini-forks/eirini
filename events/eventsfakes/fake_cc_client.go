// Code generated by counterfeiter. DO NOT EDIT.
package eventsfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/events"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
)

type FakeCcClient struct {
	AppCrashedStub        func(string, cc_messages.AppCrashedRequest, lager.Logger) error
	appCrashedMutex       sync.RWMutex
	appCrashedArgsForCall []struct {
		arg1 string
		arg2 cc_messages.AppCrashedRequest
		arg3 lager.Logger
	}
	appCrashedReturns struct {
		result1 error
	}
	appCrashedReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCcClient) AppCrashed(arg1 string, arg2 cc_messages.AppCrashedRequest, arg3 lager.Logger) error {
	fake.appCrashedMutex.Lock()
	ret, specificReturn := fake.appCrashedReturnsOnCall[len(fake.appCrashedArgsForCall)]
	fake.appCrashedArgsForCall = append(fake.appCrashedArgsForCall, struct {
		arg1 string
		arg2 cc_messages.AppCrashedRequest
		arg3 lager.Logger
	}{arg1, arg2, arg3})
	stub := fake.AppCrashedStub
	fakeReturns := fake.appCrashedReturns
	fake.recordInvocation("AppCrashed", []interface{}{arg1, arg2, arg3})
	fake.appCrashedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCcClient) AppCrashedCallCount() int {
	fake.appCrashedMutex.RLock()
	defer fake.appCrashedMutex.RUnlock()
	return len(fake.appCrashedArgsForCall)
}

func (fake *FakeCcClient) AppCrashedCalls(stub func(string, cc_messages.AppCrashedRequest, lager.Logger) error) {
	fake.appCrashedMutex.Lock()
	defer fake.appCrashedMutex.Unlock()
	fake.AppCrashedStub = stub
}

func (fake *FakeCcClient) AppCrashedArgsForCall(i int) (string, cc_messages.AppCrashedRequest, lager.Logger) {
	fake.appCrashedMutex.RLock()
	defer fake.appCrashedMutex.RUnlock()
	argsForCall := fake.appCrashedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCcClient) AppCrashedReturns(result1 error) {
	fake.appCrashedMutex.Lock()
	defer fake.appCrashedMutex.Unlock()
	fake.AppCrashedStub = nil
	fake.appCrashedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCcClient) AppCrashedReturnsOnCall(i int, result1 error) {
	fake.appCrashedMutex.Lock()
	defer fake.appCrashedMutex.Unlock()
	fake.AppCrashedStub = nil
	if fake.appCrashedReturnsOnCall == nil {
		fake.appCrashedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.appCrashedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCcClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appCrashedMutex.RLock()
	defer fake.appCrashedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCcClient) recordInvocation(key string, args []interface{}) {
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

var _ events.CcClient = new(FakeCcClient)
