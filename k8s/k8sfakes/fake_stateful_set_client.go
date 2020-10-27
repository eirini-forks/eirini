// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s"
	"code.cloudfoundry.org/eirini/opi"
	v1 "k8s.io/api/apps/v1"
)

type FakeStatefulSetClient struct {
	CreateStub        func(string, *v1.StatefulSet) (*v1.StatefulSet, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 *v1.StatefulSet
	}
	createReturns struct {
		result1 *v1.StatefulSet
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1.StatefulSet
		result2 error
	}
	DeleteStub        func(string, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetByLRPIdentifierStub        func(opi.LRPIdentifier) ([]v1.StatefulSet, error)
	getByLRPIdentifierMutex       sync.RWMutex
	getByLRPIdentifierArgsForCall []struct {
		arg1 opi.LRPIdentifier
	}
	getByLRPIdentifierReturns struct {
		result1 []v1.StatefulSet
		result2 error
	}
	getByLRPIdentifierReturnsOnCall map[int]struct {
		result1 []v1.StatefulSet
		result2 error
	}
	GetBySourceTypeStub        func(string) ([]v1.StatefulSet, error)
	getBySourceTypeMutex       sync.RWMutex
	getBySourceTypeArgsForCall []struct {
		arg1 string
	}
	getBySourceTypeReturns struct {
		result1 []v1.StatefulSet
		result2 error
	}
	getBySourceTypeReturnsOnCall map[int]struct {
		result1 []v1.StatefulSet
		result2 error
	}
	UpdateStub        func(string, *v1.StatefulSet) (*v1.StatefulSet, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 *v1.StatefulSet
	}
	updateReturns struct {
		result1 *v1.StatefulSet
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1.StatefulSet
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatefulSetClient) Create(arg1 string, arg2 *v1.StatefulSet) (*v1.StatefulSet, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 *v1.StatefulSet
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

func (fake *FakeStatefulSetClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeStatefulSetClient) CreateCalls(stub func(string, *v1.StatefulSet) (*v1.StatefulSet, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeStatefulSetClient) CreateArgsForCall(i int) (string, *v1.StatefulSet) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatefulSetClient) CreateReturns(result1 *v1.StatefulSet, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) CreateReturnsOnCall(i int, result1 *v1.StatefulSet, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1.StatefulSet
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) Delete(arg1 string, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStatefulSetClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeStatefulSetClient) DeleteCalls(stub func(string, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeStatefulSetClient) DeleteArgsForCall(i int) (string, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatefulSetClient) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatefulSetClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatefulSetClient) GetByLRPIdentifier(arg1 opi.LRPIdentifier) ([]v1.StatefulSet, error) {
	fake.getByLRPIdentifierMutex.Lock()
	ret, specificReturn := fake.getByLRPIdentifierReturnsOnCall[len(fake.getByLRPIdentifierArgsForCall)]
	fake.getByLRPIdentifierArgsForCall = append(fake.getByLRPIdentifierArgsForCall, struct {
		arg1 opi.LRPIdentifier
	}{arg1})
	stub := fake.GetByLRPIdentifierStub
	fakeReturns := fake.getByLRPIdentifierReturns
	fake.recordInvocation("GetByLRPIdentifier", []interface{}{arg1})
	fake.getByLRPIdentifierMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatefulSetClient) GetByLRPIdentifierCallCount() int {
	fake.getByLRPIdentifierMutex.RLock()
	defer fake.getByLRPIdentifierMutex.RUnlock()
	return len(fake.getByLRPIdentifierArgsForCall)
}

func (fake *FakeStatefulSetClient) GetByLRPIdentifierCalls(stub func(opi.LRPIdentifier) ([]v1.StatefulSet, error)) {
	fake.getByLRPIdentifierMutex.Lock()
	defer fake.getByLRPIdentifierMutex.Unlock()
	fake.GetByLRPIdentifierStub = stub
}

func (fake *FakeStatefulSetClient) GetByLRPIdentifierArgsForCall(i int) opi.LRPIdentifier {
	fake.getByLRPIdentifierMutex.RLock()
	defer fake.getByLRPIdentifierMutex.RUnlock()
	argsForCall := fake.getByLRPIdentifierArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStatefulSetClient) GetByLRPIdentifierReturns(result1 []v1.StatefulSet, result2 error) {
	fake.getByLRPIdentifierMutex.Lock()
	defer fake.getByLRPIdentifierMutex.Unlock()
	fake.GetByLRPIdentifierStub = nil
	fake.getByLRPIdentifierReturns = struct {
		result1 []v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) GetByLRPIdentifierReturnsOnCall(i int, result1 []v1.StatefulSet, result2 error) {
	fake.getByLRPIdentifierMutex.Lock()
	defer fake.getByLRPIdentifierMutex.Unlock()
	fake.GetByLRPIdentifierStub = nil
	if fake.getByLRPIdentifierReturnsOnCall == nil {
		fake.getByLRPIdentifierReturnsOnCall = make(map[int]struct {
			result1 []v1.StatefulSet
			result2 error
		})
	}
	fake.getByLRPIdentifierReturnsOnCall[i] = struct {
		result1 []v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) GetBySourceType(arg1 string) ([]v1.StatefulSet, error) {
	fake.getBySourceTypeMutex.Lock()
	ret, specificReturn := fake.getBySourceTypeReturnsOnCall[len(fake.getBySourceTypeArgsForCall)]
	fake.getBySourceTypeArgsForCall = append(fake.getBySourceTypeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetBySourceTypeStub
	fakeReturns := fake.getBySourceTypeReturns
	fake.recordInvocation("GetBySourceType", []interface{}{arg1})
	fake.getBySourceTypeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatefulSetClient) GetBySourceTypeCallCount() int {
	fake.getBySourceTypeMutex.RLock()
	defer fake.getBySourceTypeMutex.RUnlock()
	return len(fake.getBySourceTypeArgsForCall)
}

func (fake *FakeStatefulSetClient) GetBySourceTypeCalls(stub func(string) ([]v1.StatefulSet, error)) {
	fake.getBySourceTypeMutex.Lock()
	defer fake.getBySourceTypeMutex.Unlock()
	fake.GetBySourceTypeStub = stub
}

func (fake *FakeStatefulSetClient) GetBySourceTypeArgsForCall(i int) string {
	fake.getBySourceTypeMutex.RLock()
	defer fake.getBySourceTypeMutex.RUnlock()
	argsForCall := fake.getBySourceTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStatefulSetClient) GetBySourceTypeReturns(result1 []v1.StatefulSet, result2 error) {
	fake.getBySourceTypeMutex.Lock()
	defer fake.getBySourceTypeMutex.Unlock()
	fake.GetBySourceTypeStub = nil
	fake.getBySourceTypeReturns = struct {
		result1 []v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) GetBySourceTypeReturnsOnCall(i int, result1 []v1.StatefulSet, result2 error) {
	fake.getBySourceTypeMutex.Lock()
	defer fake.getBySourceTypeMutex.Unlock()
	fake.GetBySourceTypeStub = nil
	if fake.getBySourceTypeReturnsOnCall == nil {
		fake.getBySourceTypeReturnsOnCall = make(map[int]struct {
			result1 []v1.StatefulSet
			result2 error
		})
	}
	fake.getBySourceTypeReturnsOnCall[i] = struct {
		result1 []v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) Update(arg1 string, arg2 *v1.StatefulSet) (*v1.StatefulSet, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 *v1.StatefulSet
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

func (fake *FakeStatefulSetClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeStatefulSetClient) UpdateCalls(stub func(string, *v1.StatefulSet) (*v1.StatefulSet, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeStatefulSetClient) UpdateArgsForCall(i int) (string, *v1.StatefulSet) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatefulSetClient) UpdateReturns(result1 *v1.StatefulSet, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) UpdateReturnsOnCall(i int, result1 *v1.StatefulSet, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1.StatefulSet
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getByLRPIdentifierMutex.RLock()
	defer fake.getByLRPIdentifierMutex.RUnlock()
	fake.getBySourceTypeMutex.RLock()
	defer fake.getBySourceTypeMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatefulSetClient) recordInvocation(key string, args []interface{}) {
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

var _ k8s.StatefulSetClient = new(FakeStatefulSetClient)
