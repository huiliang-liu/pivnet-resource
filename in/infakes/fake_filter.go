// This file was generated by counterfeiter
package infakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet/v2"
)

type FakeFilter struct {
	ProductFileKeysByGlobsStub        func(productFiles []go_pivnet.ProductFile, globs []string) ([]go_pivnet.ProductFile, error)
	productFileKeysByGlobsMutex       sync.RWMutex
	productFileKeysByGlobsArgsForCall []struct {
		productFiles []go_pivnet.ProductFile
		globs        []string
	}
	productFileKeysByGlobsReturns struct {
		result1 []go_pivnet.ProductFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilter) ProductFileKeysByGlobs(productFiles []go_pivnet.ProductFile, globs []string) ([]go_pivnet.ProductFile, error) {
	var productFilesCopy []go_pivnet.ProductFile
	if productFiles != nil {
		productFilesCopy = make([]go_pivnet.ProductFile, len(productFiles))
		copy(productFilesCopy, productFiles)
	}
	var globsCopy []string
	if globs != nil {
		globsCopy = make([]string, len(globs))
		copy(globsCopy, globs)
	}
	fake.productFileKeysByGlobsMutex.Lock()
	fake.productFileKeysByGlobsArgsForCall = append(fake.productFileKeysByGlobsArgsForCall, struct {
		productFiles []go_pivnet.ProductFile
		globs        []string
	}{productFilesCopy, globsCopy})
	fake.recordInvocation("ProductFileKeysByGlobs", []interface{}{productFilesCopy, globsCopy})
	fake.productFileKeysByGlobsMutex.Unlock()
	if fake.ProductFileKeysByGlobsStub != nil {
		return fake.ProductFileKeysByGlobsStub(productFiles, globs)
	} else {
		return fake.productFileKeysByGlobsReturns.result1, fake.productFileKeysByGlobsReturns.result2
	}
}

func (fake *FakeFilter) ProductFileKeysByGlobsCallCount() int {
	fake.productFileKeysByGlobsMutex.RLock()
	defer fake.productFileKeysByGlobsMutex.RUnlock()
	return len(fake.productFileKeysByGlobsArgsForCall)
}

func (fake *FakeFilter) ProductFileKeysByGlobsArgsForCall(i int) ([]go_pivnet.ProductFile, []string) {
	fake.productFileKeysByGlobsMutex.RLock()
	defer fake.productFileKeysByGlobsMutex.RUnlock()
	return fake.productFileKeysByGlobsArgsForCall[i].productFiles, fake.productFileKeysByGlobsArgsForCall[i].globs
}

func (fake *FakeFilter) ProductFileKeysByGlobsReturns(result1 []go_pivnet.ProductFile, result2 error) {
	fake.ProductFileKeysByGlobsStub = nil
	fake.productFileKeysByGlobsReturns = struct {
		result1 []go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.productFileKeysByGlobsMutex.RLock()
	defer fake.productFileKeysByGlobsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFilter) recordInvocation(key string, args []interface{}) {
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
