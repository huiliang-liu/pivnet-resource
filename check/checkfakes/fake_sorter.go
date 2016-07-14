// This file was generated by counterfeiter
package checkfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/go-pivnet"
	"github.com/pivotal-cf-experimental/pivnet-resource/check"
)

type FakeSorter struct {
	SortBySemverStub        func([]pivnet.Release) ([]pivnet.Release, error)
	sortBySemverMutex       sync.RWMutex
	sortBySemverArgsForCall []struct {
		arg1 []pivnet.Release
	}
	sortBySemverReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSorter) SortBySemver(arg1 []pivnet.Release) ([]pivnet.Release, error) {
	var arg1Copy []pivnet.Release
	if arg1 != nil {
		arg1Copy = make([]pivnet.Release, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.sortBySemverMutex.Lock()
	fake.sortBySemverArgsForCall = append(fake.sortBySemverArgsForCall, struct {
		arg1 []pivnet.Release
	}{arg1Copy})
	fake.recordInvocation("SortBySemver", []interface{}{arg1Copy})
	fake.sortBySemverMutex.Unlock()
	if fake.SortBySemverStub != nil {
		return fake.SortBySemverStub(arg1)
	} else {
		return fake.sortBySemverReturns.result1, fake.sortBySemverReturns.result2
	}
}

func (fake *FakeSorter) SortBySemverCallCount() int {
	fake.sortBySemverMutex.RLock()
	defer fake.sortBySemverMutex.RUnlock()
	return len(fake.sortBySemverArgsForCall)
}

func (fake *FakeSorter) SortBySemverArgsForCall(i int) []pivnet.Release {
	fake.sortBySemverMutex.RLock()
	defer fake.sortBySemverMutex.RUnlock()
	return fake.sortBySemverArgsForCall[i].arg1
}

func (fake *FakeSorter) SortBySemverReturns(result1 []pivnet.Release, result2 error) {
	fake.SortBySemverStub = nil
	fake.sortBySemverReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeSorter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sortBySemverMutex.RLock()
	defer fake.sortBySemverMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSorter) recordInvocation(key string, args []interface{}) {
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

var _ check.Sorter = new(FakeSorter)
