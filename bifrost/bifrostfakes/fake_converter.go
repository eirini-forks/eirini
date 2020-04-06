// Code generated by counterfeiter. DO NOT EDIT.
package bifrostfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/bifrost"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/eirini/opi"
)

type FakeConverter struct {
	ConvertLRPStub        func(cf.DesireLRPRequest) (opi.LRP, error)
	convertLRPMutex       sync.RWMutex
	convertLRPArgsForCall []struct {
		arg1 cf.DesireLRPRequest
	}
	convertLRPReturns struct {
		result1 opi.LRP
		result2 error
	}
	convertLRPReturnsOnCall map[int]struct {
		result1 opi.LRP
		result2 error
	}
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

func (fake *FakeConverter) ConvertLRP(arg1 cf.DesireLRPRequest) (opi.LRP, error) {
	fake.convertLRPMutex.Lock()
	ret, specificReturn := fake.convertLRPReturnsOnCall[len(fake.convertLRPArgsForCall)]
	fake.convertLRPArgsForCall = append(fake.convertLRPArgsForCall, struct {
		arg1 cf.DesireLRPRequest
	}{arg1})
	fake.recordInvocation("ConvertLRP", []interface{}{arg1})
	fake.convertLRPMutex.Unlock()
	if fake.ConvertLRPStub != nil {
		return fake.ConvertLRPStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.convertLRPReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConverter) ConvertLRPCallCount() int {
	fake.convertLRPMutex.RLock()
	defer fake.convertLRPMutex.RUnlock()
	return len(fake.convertLRPArgsForCall)
}

func (fake *FakeConverter) ConvertLRPCalls(stub func(cf.DesireLRPRequest) (opi.LRP, error)) {
	fake.convertLRPMutex.Lock()
	defer fake.convertLRPMutex.Unlock()
	fake.ConvertLRPStub = stub
}

func (fake *FakeConverter) ConvertLRPArgsForCall(i int) cf.DesireLRPRequest {
	fake.convertLRPMutex.RLock()
	defer fake.convertLRPMutex.RUnlock()
	argsForCall := fake.convertLRPArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConverter) ConvertLRPReturns(result1 opi.LRP, result2 error) {
	fake.convertLRPMutex.Lock()
	defer fake.convertLRPMutex.Unlock()
	fake.ConvertLRPStub = nil
	fake.convertLRPReturns = struct {
		result1 opi.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeConverter) ConvertLRPReturnsOnCall(i int, result1 opi.LRP, result2 error) {
	fake.convertLRPMutex.Lock()
	defer fake.convertLRPMutex.Unlock()
	fake.ConvertLRPStub = nil
	if fake.convertLRPReturnsOnCall == nil {
		fake.convertLRPReturnsOnCall = make(map[int]struct {
			result1 opi.LRP
			result2 error
		})
	}
	fake.convertLRPReturnsOnCall[i] = struct {
		result1 opi.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeConverter) ConvertTask(arg1 string, arg2 cf.TaskRequest) (opi.Task, error) {
	fake.convertTaskMutex.Lock()
	ret, specificReturn := fake.convertTaskReturnsOnCall[len(fake.convertTaskArgsForCall)]
	fake.convertTaskArgsForCall = append(fake.convertTaskArgsForCall, struct {
		arg1 string
		arg2 cf.TaskRequest
	}{arg1, arg2})
	fake.recordInvocation("ConvertTask", []interface{}{arg1, arg2})
	fake.convertTaskMutex.Unlock()
	if fake.ConvertTaskStub != nil {
		return fake.ConvertTaskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.convertTaskReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConverter) ConvertTaskCallCount() int {
	fake.convertTaskMutex.RLock()
	defer fake.convertTaskMutex.RUnlock()
	return len(fake.convertTaskArgsForCall)
}

func (fake *FakeConverter) ConvertTaskCalls(stub func(string, cf.TaskRequest) (opi.Task, error)) {
	fake.convertTaskMutex.Lock()
	defer fake.convertTaskMutex.Unlock()
	fake.ConvertTaskStub = stub
}

func (fake *FakeConverter) ConvertTaskArgsForCall(i int) (string, cf.TaskRequest) {
	fake.convertTaskMutex.RLock()
	defer fake.convertTaskMutex.RUnlock()
	argsForCall := fake.convertTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConverter) ConvertTaskReturns(result1 opi.Task, result2 error) {
	fake.convertTaskMutex.Lock()
	defer fake.convertTaskMutex.Unlock()
	fake.ConvertTaskStub = nil
	fake.convertTaskReturns = struct {
		result1 opi.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeConverter) ConvertTaskReturnsOnCall(i int, result1 opi.Task, result2 error) {
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

func (fake *FakeConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convertLRPMutex.RLock()
	defer fake.convertLRPMutex.RUnlock()
	fake.convertTaskMutex.RLock()
	defer fake.convertTaskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConverter) recordInvocation(key string, args []interface{}) {
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

var _ bifrost.Converter = new(FakeConverter)
