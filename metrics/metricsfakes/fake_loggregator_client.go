// Code generated by counterfeiter. DO NOT EDIT.
package metricsfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/metrics"
	loggregator "code.cloudfoundry.org/go-loggregator"
)

type FakeLoggregatorClient struct {
	EmitGaugeStub        func(...loggregator.EmitGaugeOption)
	emitGaugeMutex       sync.RWMutex
	emitGaugeArgsForCall []struct {
		arg1 []loggregator.EmitGaugeOption
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoggregatorClient) EmitGauge(arg1 ...loggregator.EmitGaugeOption) {
	fake.emitGaugeMutex.Lock()
	fake.emitGaugeArgsForCall = append(fake.emitGaugeArgsForCall, struct {
		arg1 []loggregator.EmitGaugeOption
	}{arg1})
	stub := fake.EmitGaugeStub
	fake.recordInvocation("EmitGauge", []interface{}{arg1})
	fake.emitGaugeMutex.Unlock()
	if stub != nil {
		fake.EmitGaugeStub(arg1...)
	}
}

func (fake *FakeLoggregatorClient) EmitGaugeCallCount() int {
	fake.emitGaugeMutex.RLock()
	defer fake.emitGaugeMutex.RUnlock()
	return len(fake.emitGaugeArgsForCall)
}

func (fake *FakeLoggregatorClient) EmitGaugeCalls(stub func(...loggregator.EmitGaugeOption)) {
	fake.emitGaugeMutex.Lock()
	defer fake.emitGaugeMutex.Unlock()
	fake.EmitGaugeStub = stub
}

func (fake *FakeLoggregatorClient) EmitGaugeArgsForCall(i int) []loggregator.EmitGaugeOption {
	fake.emitGaugeMutex.RLock()
	defer fake.emitGaugeMutex.RUnlock()
	argsForCall := fake.emitGaugeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLoggregatorClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitGaugeMutex.RLock()
	defer fake.emitGaugeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLoggregatorClient) recordInvocation(key string, args []interface{}) {
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

var _ metrics.LoggregatorClient = new(FakeLoggregatorClient)
