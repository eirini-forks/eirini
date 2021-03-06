// Code generated by counterfeiter. DO NOT EDIT.
package routefakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/route"
)

type FakeEmitter struct {
	EmitStub        func(route.Message)
	emitMutex       sync.RWMutex
	emitArgsForCall []struct {
		arg1 route.Message
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) Emit(arg1 route.Message) {
	fake.emitMutex.Lock()
	fake.emitArgsForCall = append(fake.emitArgsForCall, struct {
		arg1 route.Message
	}{arg1})
	stub := fake.EmitStub
	fake.recordInvocation("Emit", []interface{}{arg1})
	fake.emitMutex.Unlock()
	if stub != nil {
		fake.EmitStub(arg1)
	}
}

func (fake *FakeEmitter) EmitCallCount() int {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return len(fake.emitArgsForCall)
}

func (fake *FakeEmitter) EmitCalls(stub func(route.Message)) {
	fake.emitMutex.Lock()
	defer fake.emitMutex.Unlock()
	fake.EmitStub = stub
}

func (fake *FakeEmitter) EmitArgsForCall(i int) route.Message {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	argsForCall := fake.emitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
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

var _ route.Emitter = new(FakeEmitter)
