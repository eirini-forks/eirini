// Code generated by counterfeiter. DO NOT EDIT.
package bifrostfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/bifrost"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/eirini/opi"
)

type FakeTaskConverter struct {
	ConvertTaskStub        func(string, cf.TaskRequest) (opi.Task, error)
	convertTaskMutex       sync.RWMutex
	convertTaskArgsForCall []struct {
		arg1 string
		arg2 cf.TaskRequest
	}
	convertTaskReturns struct {
		result1 opi.Task
		result2 error
	}
	convertTaskReturnsOnCall map[int]struct {
		result1 opi.Task
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskConverter) ConvertTask(arg1 string, arg2 cf.TaskRequest) (opi.Task, error) {
	fake.convertTaskMutex.Lock()
	ret, specificReturn := fake.convertTaskReturnsOnCall[len(fake.convertTaskArgsForCall)]
	fake.convertTaskArgsForCall = append(fake.convertTaskArgsForCall, struct {
		arg1 string
		arg2 cf.TaskRequest
	}{arg1, arg2})
	stub := fake.ConvertTaskStub
	fakeReturns := fake.convertTaskReturns
	fake.recordInvocation("ConvertTask", []interface{}{arg1, arg2})
	fake.convertTaskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskConverter) ConvertTaskCallCount() int {
	fake.convertTaskMutex.RLock()
	defer fake.convertTaskMutex.RUnlock()
	return len(fake.convertTaskArgsForCall)
}

func (fake *FakeTaskConverter) ConvertTaskCalls(stub func(string, cf.TaskRequest) (opi.Task, error)) {
	fake.convertTaskMutex.Lock()
	defer fake.convertTaskMutex.Unlock()
	fake.ConvertTaskStub = stub
}

func (fake *FakeTaskConverter) ConvertTaskArgsForCall(i int) (string, cf.TaskRequest) {
	fake.convertTaskMutex.RLock()
	defer fake.convertTaskMutex.RUnlock()
	argsForCall := fake.convertTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskConverter) ConvertTaskReturns(result1 opi.Task, result2 error) {
	fake.convertTaskMutex.Lock()
	defer fake.convertTaskMutex.Unlock()
	fake.ConvertTaskStub = nil
	fake.convertTaskReturns = struct {
		result1 opi.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskConverter) ConvertTaskReturnsOnCall(i int, result1 opi.Task, result2 error) {
	fake.convertTaskMutex.Lock()
	defer fake.convertTaskMutex.Unlock()
	fake.ConvertTaskStub = nil
	if fake.convertTaskReturnsOnCall == nil {
		fake.convertTaskReturnsOnCall = make(map[int]struct {
			result1 opi.Task
			result2 error
		})
	}
	fake.convertTaskReturnsOnCall[i] = struct {
		result1 opi.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convertTaskMutex.RLock()
	defer fake.convertTaskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskConverter) recordInvocation(key string, args []interface{}) {
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

var _ bifrost.TaskConverter = new(FakeTaskConverter)
