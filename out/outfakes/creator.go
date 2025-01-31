// This file was generated by counterfeiter
package outfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet/v2"
)

type Creator struct {
	CreateStub        func() (go_pivnet.Release, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct{}
	createReturns     struct {
		result1 go_pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Creator) Create() (go_pivnet.Release, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.recordInvocation("Create", []interface{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub()
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *Creator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Creator) CreateReturns(result1 go_pivnet.Release, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *Creator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.invocations
}

func (fake *Creator) recordInvocation(key string, args []interface{}) {
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
