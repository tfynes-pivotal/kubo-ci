// Code generated by counterfeiter. DO NOT EDIT.
package directorfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/director"
	semver "github.com/cppforlife/go-semi-semantic/version"
)

type FakeRelease struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() semver.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 semver.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 semver.Version
	}
	VersionMarkStub        func(mark string) string
	versionMarkMutex       sync.RWMutex
	versionMarkArgsForCall []struct {
		mark string
	}
	versionMarkReturns struct {
		result1 string
	}
	versionMarkReturnsOnCall map[int]struct {
		result1 string
	}
	CommitHashWithMarkStub        func(mark string) string
	commitHashWithMarkMutex       sync.RWMutex
	commitHashWithMarkArgsForCall []struct {
		mark string
	}
	commitHashWithMarkReturns struct {
		result1 string
	}
	commitHashWithMarkReturnsOnCall map[int]struct {
		result1 string
	}
	JobsStub        func() ([]director.Job, error)
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct{}
	jobsReturns     struct {
		result1 []director.Job
		result2 error
	}
	jobsReturnsOnCall map[int]struct {
		result1 []director.Job
		result2 error
	}
	PackagesStub        func() ([]director.Package, error)
	packagesMutex       sync.RWMutex
	packagesArgsForCall []struct{}
	packagesReturns     struct {
		result1 []director.Package
		result2 error
	}
	packagesReturnsOnCall map[int]struct {
		result1 []director.Package
		result2 error
	}
	DeleteStub        func(force bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		force bool
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRelease) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeRelease) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeRelease) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) Version() semver.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeRelease) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeRelease) VersionReturns(result1 semver.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 semver.Version
	}{result1}
}

func (fake *FakeRelease) VersionReturnsOnCall(i int, result1 semver.Version) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 semver.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 semver.Version
	}{result1}
}

func (fake *FakeRelease) VersionMark(mark string) string {
	fake.versionMarkMutex.Lock()
	ret, specificReturn := fake.versionMarkReturnsOnCall[len(fake.versionMarkArgsForCall)]
	fake.versionMarkArgsForCall = append(fake.versionMarkArgsForCall, struct {
		mark string
	}{mark})
	fake.recordInvocation("VersionMark", []interface{}{mark})
	fake.versionMarkMutex.Unlock()
	if fake.VersionMarkStub != nil {
		return fake.VersionMarkStub(mark)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionMarkReturns.result1
}

func (fake *FakeRelease) VersionMarkCallCount() int {
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	return len(fake.versionMarkArgsForCall)
}

func (fake *FakeRelease) VersionMarkArgsForCall(i int) string {
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	return fake.versionMarkArgsForCall[i].mark
}

func (fake *FakeRelease) VersionMarkReturns(result1 string) {
	fake.VersionMarkStub = nil
	fake.versionMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) VersionMarkReturnsOnCall(i int, result1 string) {
	fake.VersionMarkStub = nil
	if fake.versionMarkReturnsOnCall == nil {
		fake.versionMarkReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionMarkReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMark(mark string) string {
	fake.commitHashWithMarkMutex.Lock()
	ret, specificReturn := fake.commitHashWithMarkReturnsOnCall[len(fake.commitHashWithMarkArgsForCall)]
	fake.commitHashWithMarkArgsForCall = append(fake.commitHashWithMarkArgsForCall, struct {
		mark string
	}{mark})
	fake.recordInvocation("CommitHashWithMark", []interface{}{mark})
	fake.commitHashWithMarkMutex.Unlock()
	if fake.CommitHashWithMarkStub != nil {
		return fake.CommitHashWithMarkStub(mark)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.commitHashWithMarkReturns.result1
}

func (fake *FakeRelease) CommitHashWithMarkCallCount() int {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return len(fake.commitHashWithMarkArgsForCall)
}

func (fake *FakeRelease) CommitHashWithMarkArgsForCall(i int) string {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return fake.commitHashWithMarkArgsForCall[i].mark
}

func (fake *FakeRelease) CommitHashWithMarkReturns(result1 string) {
	fake.CommitHashWithMarkStub = nil
	fake.commitHashWithMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMarkReturnsOnCall(i int, result1 string) {
	fake.CommitHashWithMarkStub = nil
	if fake.commitHashWithMarkReturnsOnCall == nil {
		fake.commitHashWithMarkReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.commitHashWithMarkReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) Jobs() ([]director.Job, error) {
	fake.jobsMutex.Lock()
	ret, specificReturn := fake.jobsReturnsOnCall[len(fake.jobsArgsForCall)]
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct{}{})
	fake.recordInvocation("Jobs", []interface{}{})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.jobsReturns.result1, fake.jobsReturns.result2
}

func (fake *FakeRelease) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *FakeRelease) JobsReturns(result1 []director.Job, result2 error) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 []director.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeRelease) JobsReturnsOnCall(i int, result1 []director.Job, result2 error) {
	fake.JobsStub = nil
	if fake.jobsReturnsOnCall == nil {
		fake.jobsReturnsOnCall = make(map[int]struct {
			result1 []director.Job
			result2 error
		})
	}
	fake.jobsReturnsOnCall[i] = struct {
		result1 []director.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeRelease) Packages() ([]director.Package, error) {
	fake.packagesMutex.Lock()
	ret, specificReturn := fake.packagesReturnsOnCall[len(fake.packagesArgsForCall)]
	fake.packagesArgsForCall = append(fake.packagesArgsForCall, struct{}{})
	fake.recordInvocation("Packages", []interface{}{})
	fake.packagesMutex.Unlock()
	if fake.PackagesStub != nil {
		return fake.PackagesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.packagesReturns.result1, fake.packagesReturns.result2
}

func (fake *FakeRelease) PackagesCallCount() int {
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	return len(fake.packagesArgsForCall)
}

func (fake *FakeRelease) PackagesReturns(result1 []director.Package, result2 error) {
	fake.PackagesStub = nil
	fake.packagesReturns = struct {
		result1 []director.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeRelease) PackagesReturnsOnCall(i int, result1 []director.Package, result2 error) {
	fake.PackagesStub = nil
	if fake.packagesReturnsOnCall == nil {
		fake.packagesReturnsOnCall = make(map[int]struct {
			result1 []director.Package
			result2 error
		})
	}
	fake.packagesReturnsOnCall[i] = struct {
		result1 []director.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeRelease) Delete(force bool) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		force bool
	}{force})
	fake.recordInvocation("Delete", []interface{}{force})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(force)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeRelease) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRelease) DeleteArgsForCall(i int) bool {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].force
}

func (fake *FakeRelease) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeRelease) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRelease) recordInvocation(key string, args []interface{}) {
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

var _ director.Release = new(FakeRelease)