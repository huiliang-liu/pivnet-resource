// This file was generated by counterfeiter
package checkfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/go-pivnet"
	"github.com/pivotal-cf-experimental/pivnet-resource/check"
)

type FakeFilter struct {
	ReleasesByReleaseTypeStub        func(releases []pivnet.Release, releaseType string) ([]pivnet.Release, error)
	releasesByReleaseTypeMutex       sync.RWMutex
	releasesByReleaseTypeArgsForCall []struct {
		releases    []pivnet.Release
		releaseType string
	}
	releasesByReleaseTypeReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	ReleasesByVersionStub        func(releases []pivnet.Release, version string) ([]pivnet.Release, error)
	releasesByVersionMutex       sync.RWMutex
	releasesByVersionArgsForCall []struct {
		releases []pivnet.Release
		version  string
	}
	releasesByVersionReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilter) ReleasesByReleaseType(releases []pivnet.Release, releaseType string) ([]pivnet.Release, error) {
	var releasesCopy []pivnet.Release
	if releases != nil {
		releasesCopy = make([]pivnet.Release, len(releases))
		copy(releasesCopy, releases)
	}
	fake.releasesByReleaseTypeMutex.Lock()
	fake.releasesByReleaseTypeArgsForCall = append(fake.releasesByReleaseTypeArgsForCall, struct {
		releases    []pivnet.Release
		releaseType string
	}{releasesCopy, releaseType})
	fake.recordInvocation("ReleasesByReleaseType", []interface{}{releasesCopy, releaseType})
	fake.releasesByReleaseTypeMutex.Unlock()
	if fake.ReleasesByReleaseTypeStub != nil {
		return fake.ReleasesByReleaseTypeStub(releases, releaseType)
	} else {
		return fake.releasesByReleaseTypeReturns.result1, fake.releasesByReleaseTypeReturns.result2
	}
}

func (fake *FakeFilter) ReleasesByReleaseTypeCallCount() int {
	fake.releasesByReleaseTypeMutex.RLock()
	defer fake.releasesByReleaseTypeMutex.RUnlock()
	return len(fake.releasesByReleaseTypeArgsForCall)
}

func (fake *FakeFilter) ReleasesByReleaseTypeArgsForCall(i int) ([]pivnet.Release, string) {
	fake.releasesByReleaseTypeMutex.RLock()
	defer fake.releasesByReleaseTypeMutex.RUnlock()
	return fake.releasesByReleaseTypeArgsForCall[i].releases, fake.releasesByReleaseTypeArgsForCall[i].releaseType
}

func (fake *FakeFilter) ReleasesByReleaseTypeReturns(result1 []pivnet.Release, result2 error) {
	fake.ReleasesByReleaseTypeStub = nil
	fake.releasesByReleaseTypeReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) ReleasesByVersion(releases []pivnet.Release, version string) ([]pivnet.Release, error) {
	var releasesCopy []pivnet.Release
	if releases != nil {
		releasesCopy = make([]pivnet.Release, len(releases))
		copy(releasesCopy, releases)
	}
	fake.releasesByVersionMutex.Lock()
	fake.releasesByVersionArgsForCall = append(fake.releasesByVersionArgsForCall, struct {
		releases []pivnet.Release
		version  string
	}{releasesCopy, version})
	fake.recordInvocation("ReleasesByVersion", []interface{}{releasesCopy, version})
	fake.releasesByVersionMutex.Unlock()
	if fake.ReleasesByVersionStub != nil {
		return fake.ReleasesByVersionStub(releases, version)
	} else {
		return fake.releasesByVersionReturns.result1, fake.releasesByVersionReturns.result2
	}
}

func (fake *FakeFilter) ReleasesByVersionCallCount() int {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return len(fake.releasesByVersionArgsForCall)
}

func (fake *FakeFilter) ReleasesByVersionArgsForCall(i int) ([]pivnet.Release, string) {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return fake.releasesByVersionArgsForCall[i].releases, fake.releasesByVersionArgsForCall[i].version
}

func (fake *FakeFilter) ReleasesByVersionReturns(result1 []pivnet.Release, result2 error) {
	fake.ReleasesByVersionStub = nil
	fake.releasesByVersionReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releasesByReleaseTypeMutex.RLock()
	defer fake.releasesByReleaseTypeMutex.RUnlock()
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
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

var _ check.Filter = new(FakeFilter)
