// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"03testifyexample/service"
	"sync"
)

type FakeStudent struct {
	GetAgeStub        func() int
	getAgeMutex       sync.RWMutex
	getAgeArgsForCall []struct {
	}
	getAgeReturns struct {
		result1 int
	}
	getAgeReturnsOnCall map[int]struct {
		result1 int
	}
	GetFullNameStub        func() string
	getFullNameMutex       sync.RWMutex
	getFullNameArgsForCall []struct {
	}
	getFullNameReturns struct {
		result1 string
	}
	getFullNameReturnsOnCall map[int]struct {
		result1 string
	}
	IsFullTimeStub        func() bool
	isFullTimeMutex       sync.RWMutex
	isFullTimeArgsForCall []struct {
	}
	isFullTimeReturns struct {
		result1 bool
	}
	isFullTimeReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStudent) GetAge() int {
	fake.getAgeMutex.Lock()
	ret, specificReturn := fake.getAgeReturnsOnCall[len(fake.getAgeArgsForCall)]
	fake.getAgeArgsForCall = append(fake.getAgeArgsForCall, struct {
	}{})
	stub := fake.GetAgeStub
	fakeReturns := fake.getAgeReturns
	fake.recordInvocation("GetAge", []interface{}{})
	fake.getAgeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStudent) GetAgeCallCount() int {
	fake.getAgeMutex.RLock()
	defer fake.getAgeMutex.RUnlock()
	return len(fake.getAgeArgsForCall)
}

func (fake *FakeStudent) GetAgeCalls(stub func() int) {
	fake.getAgeMutex.Lock()
	defer fake.getAgeMutex.Unlock()
	fake.GetAgeStub = stub
}

func (fake *FakeStudent) GetAgeReturns(result1 int) {
	fake.getAgeMutex.Lock()
	defer fake.getAgeMutex.Unlock()
	fake.GetAgeStub = nil
	fake.getAgeReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeStudent) GetAgeReturnsOnCall(i int, result1 int) {
	fake.getAgeMutex.Lock()
	defer fake.getAgeMutex.Unlock()
	fake.GetAgeStub = nil
	if fake.getAgeReturnsOnCall == nil {
		fake.getAgeReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getAgeReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeStudent) GetFullName() string {
	fake.getFullNameMutex.Lock()
	ret, specificReturn := fake.getFullNameReturnsOnCall[len(fake.getFullNameArgsForCall)]
	fake.getFullNameArgsForCall = append(fake.getFullNameArgsForCall, struct {
	}{})
	stub := fake.GetFullNameStub
	fakeReturns := fake.getFullNameReturns
	fake.recordInvocation("GetFullName", []interface{}{})
	fake.getFullNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStudent) GetFullNameCallCount() int {
	fake.getFullNameMutex.RLock()
	defer fake.getFullNameMutex.RUnlock()
	return len(fake.getFullNameArgsForCall)
}

func (fake *FakeStudent) GetFullNameCalls(stub func() string) {
	fake.getFullNameMutex.Lock()
	defer fake.getFullNameMutex.Unlock()
	fake.GetFullNameStub = stub
}

func (fake *FakeStudent) GetFullNameReturns(result1 string) {
	fake.getFullNameMutex.Lock()
	defer fake.getFullNameMutex.Unlock()
	fake.GetFullNameStub = nil
	fake.getFullNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStudent) GetFullNameReturnsOnCall(i int, result1 string) {
	fake.getFullNameMutex.Lock()
	defer fake.getFullNameMutex.Unlock()
	fake.GetFullNameStub = nil
	if fake.getFullNameReturnsOnCall == nil {
		fake.getFullNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getFullNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStudent) IsFullTime() bool {
	fake.isFullTimeMutex.Lock()
	ret, specificReturn := fake.isFullTimeReturnsOnCall[len(fake.isFullTimeArgsForCall)]
	fake.isFullTimeArgsForCall = append(fake.isFullTimeArgsForCall, struct {
	}{})
	stub := fake.IsFullTimeStub
	fakeReturns := fake.isFullTimeReturns
	fake.recordInvocation("IsFullTime", []interface{}{})
	fake.isFullTimeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStudent) IsFullTimeCallCount() int {
	fake.isFullTimeMutex.RLock()
	defer fake.isFullTimeMutex.RUnlock()
	return len(fake.isFullTimeArgsForCall)
}

func (fake *FakeStudent) IsFullTimeCalls(stub func() bool) {
	fake.isFullTimeMutex.Lock()
	defer fake.isFullTimeMutex.Unlock()
	fake.IsFullTimeStub = stub
}

func (fake *FakeStudent) IsFullTimeReturns(result1 bool) {
	fake.isFullTimeMutex.Lock()
	defer fake.isFullTimeMutex.Unlock()
	fake.IsFullTimeStub = nil
	fake.isFullTimeReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStudent) IsFullTimeReturnsOnCall(i int, result1 bool) {
	fake.isFullTimeMutex.Lock()
	defer fake.isFullTimeMutex.Unlock()
	fake.IsFullTimeStub = nil
	if fake.isFullTimeReturnsOnCall == nil {
		fake.isFullTimeReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isFullTimeReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStudent) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAgeMutex.RLock()
	defer fake.getAgeMutex.RUnlock()
	fake.getFullNameMutex.RLock()
	defer fake.getFullNameMutex.RUnlock()
	fake.isFullTimeMutex.RLock()
	defer fake.isFullTimeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStudent) recordInvocation(key string, args []interface{}) {
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

var _ service.Student = new(FakeStudent)
