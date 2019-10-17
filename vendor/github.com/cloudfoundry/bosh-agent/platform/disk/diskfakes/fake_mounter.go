// Code generated by counterfeiter. DO NOT EDIT.
package diskfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/platform/disk"
)

type FakeMounter struct {
	IsMountPointStub        func(string) (string, bool, error)
	isMountPointMutex       sync.RWMutex
	isMountPointArgsForCall []struct {
		arg1 string
	}
	isMountPointReturns struct {
		result1 string
		result2 bool
		result3 error
	}
	isMountPointReturnsOnCall map[int]struct {
		result1 string
		result2 bool
		result3 error
	}
	IsMountedStub        func(string) (bool, error)
	isMountedMutex       sync.RWMutex
	isMountedArgsForCall []struct {
		arg1 string
	}
	isMountedReturns struct {
		result1 bool
		result2 error
	}
	isMountedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	MountStub        func(string, string, ...string) error
	mountMutex       sync.RWMutex
	mountArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []string
	}
	mountReturns struct {
		result1 error
	}
	mountReturnsOnCall map[int]struct {
		result1 error
	}
	MountFilesystemStub        func(string, string, string, ...string) error
	mountFilesystemMutex       sync.RWMutex
	mountFilesystemArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 []string
	}
	mountFilesystemReturns struct {
		result1 error
	}
	mountFilesystemReturnsOnCall map[int]struct {
		result1 error
	}
	MountTmpfsStub        func(string, string) error
	mountTmpfsMutex       sync.RWMutex
	mountTmpfsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	mountTmpfsReturns struct {
		result1 error
	}
	mountTmpfsReturnsOnCall map[int]struct {
		result1 error
	}
	RemountStub        func(string, string, ...string) error
	remountMutex       sync.RWMutex
	remountArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []string
	}
	remountReturns struct {
		result1 error
	}
	remountReturnsOnCall map[int]struct {
		result1 error
	}
	RemountAsReadonlyStub        func(string) error
	remountAsReadonlyMutex       sync.RWMutex
	remountAsReadonlyArgsForCall []struct {
		arg1 string
	}
	remountAsReadonlyReturns struct {
		result1 error
	}
	remountAsReadonlyReturnsOnCall map[int]struct {
		result1 error
	}
	RemountInPlaceStub        func(string, ...string) error
	remountInPlaceMutex       sync.RWMutex
	remountInPlaceArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	remountInPlaceReturns struct {
		result1 error
	}
	remountInPlaceReturnsOnCall map[int]struct {
		result1 error
	}
	SwapOnStub        func(string) error
	swapOnMutex       sync.RWMutex
	swapOnArgsForCall []struct {
		arg1 string
	}
	swapOnReturns struct {
		result1 error
	}
	swapOnReturnsOnCall map[int]struct {
		result1 error
	}
	UnmountStub        func(string) (bool, error)
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		arg1 string
	}
	unmountReturns struct {
		result1 bool
		result2 error
	}
	unmountReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMounter) IsMountPoint(arg1 string) (string, bool, error) {
	fake.isMountPointMutex.Lock()
	ret, specificReturn := fake.isMountPointReturnsOnCall[len(fake.isMountPointArgsForCall)]
	fake.isMountPointArgsForCall = append(fake.isMountPointArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsMountPoint", []interface{}{arg1})
	fake.isMountPointMutex.Unlock()
	if fake.IsMountPointStub != nil {
		return fake.IsMountPointStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.isMountPointReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeMounter) IsMountPointCallCount() int {
	fake.isMountPointMutex.RLock()
	defer fake.isMountPointMutex.RUnlock()
	return len(fake.isMountPointArgsForCall)
}

func (fake *FakeMounter) IsMountPointCalls(stub func(string) (string, bool, error)) {
	fake.isMountPointMutex.Lock()
	defer fake.isMountPointMutex.Unlock()
	fake.IsMountPointStub = stub
}

func (fake *FakeMounter) IsMountPointArgsForCall(i int) string {
	fake.isMountPointMutex.RLock()
	defer fake.isMountPointMutex.RUnlock()
	argsForCall := fake.isMountPointArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMounter) IsMountPointReturns(result1 string, result2 bool, result3 error) {
	fake.isMountPointMutex.Lock()
	defer fake.isMountPointMutex.Unlock()
	fake.IsMountPointStub = nil
	fake.isMountPointReturns = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeMounter) IsMountPointReturnsOnCall(i int, result1 string, result2 bool, result3 error) {
	fake.isMountPointMutex.Lock()
	defer fake.isMountPointMutex.Unlock()
	fake.IsMountPointStub = nil
	if fake.isMountPointReturnsOnCall == nil {
		fake.isMountPointReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
			result3 error
		})
	}
	fake.isMountPointReturnsOnCall[i] = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeMounter) IsMounted(arg1 string) (bool, error) {
	fake.isMountedMutex.Lock()
	ret, specificReturn := fake.isMountedReturnsOnCall[len(fake.isMountedArgsForCall)]
	fake.isMountedArgsForCall = append(fake.isMountedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsMounted", []interface{}{arg1})
	fake.isMountedMutex.Unlock()
	if fake.IsMountedStub != nil {
		return fake.IsMountedStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.isMountedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMounter) IsMountedCallCount() int {
	fake.isMountedMutex.RLock()
	defer fake.isMountedMutex.RUnlock()
	return len(fake.isMountedArgsForCall)
}

func (fake *FakeMounter) IsMountedCalls(stub func(string) (bool, error)) {
	fake.isMountedMutex.Lock()
	defer fake.isMountedMutex.Unlock()
	fake.IsMountedStub = stub
}

func (fake *FakeMounter) IsMountedArgsForCall(i int) string {
	fake.isMountedMutex.RLock()
	defer fake.isMountedMutex.RUnlock()
	argsForCall := fake.isMountedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMounter) IsMountedReturns(result1 bool, result2 error) {
	fake.isMountedMutex.Lock()
	defer fake.isMountedMutex.Unlock()
	fake.IsMountedStub = nil
	fake.isMountedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMounter) IsMountedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isMountedMutex.Lock()
	defer fake.isMountedMutex.Unlock()
	fake.IsMountedStub = nil
	if fake.isMountedReturnsOnCall == nil {
		fake.isMountedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isMountedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMounter) Mount(arg1 string, arg2 string, arg3 ...string) error {
	fake.mountMutex.Lock()
	ret, specificReturn := fake.mountReturnsOnCall[len(fake.mountArgsForCall)]
	fake.mountArgsForCall = append(fake.mountArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Mount", []interface{}{arg1, arg2, arg3})
	fake.mountMutex.Unlock()
	if fake.MountStub != nil {
		return fake.MountStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mountReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) MountCallCount() int {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return len(fake.mountArgsForCall)
}

func (fake *FakeMounter) MountCalls(stub func(string, string, ...string) error) {
	fake.mountMutex.Lock()
	defer fake.mountMutex.Unlock()
	fake.MountStub = stub
}

func (fake *FakeMounter) MountArgsForCall(i int) (string, string, []string) {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	argsForCall := fake.mountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeMounter) MountReturns(result1 error) {
	fake.mountMutex.Lock()
	defer fake.mountMutex.Unlock()
	fake.MountStub = nil
	fake.mountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) MountReturnsOnCall(i int, result1 error) {
	fake.mountMutex.Lock()
	defer fake.mountMutex.Unlock()
	fake.MountStub = nil
	if fake.mountReturnsOnCall == nil {
		fake.mountReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) MountFilesystem(arg1 string, arg2 string, arg3 string, arg4 ...string) error {
	fake.mountFilesystemMutex.Lock()
	ret, specificReturn := fake.mountFilesystemReturnsOnCall[len(fake.mountFilesystemArgsForCall)]
	fake.mountFilesystemArgsForCall = append(fake.mountFilesystemArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 []string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("MountFilesystem", []interface{}{arg1, arg2, arg3, arg4})
	fake.mountFilesystemMutex.Unlock()
	if fake.MountFilesystemStub != nil {
		return fake.MountFilesystemStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mountFilesystemReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) MountFilesystemCallCount() int {
	fake.mountFilesystemMutex.RLock()
	defer fake.mountFilesystemMutex.RUnlock()
	return len(fake.mountFilesystemArgsForCall)
}

func (fake *FakeMounter) MountFilesystemCalls(stub func(string, string, string, ...string) error) {
	fake.mountFilesystemMutex.Lock()
	defer fake.mountFilesystemMutex.Unlock()
	fake.MountFilesystemStub = stub
}

func (fake *FakeMounter) MountFilesystemArgsForCall(i int) (string, string, string, []string) {
	fake.mountFilesystemMutex.RLock()
	defer fake.mountFilesystemMutex.RUnlock()
	argsForCall := fake.mountFilesystemArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeMounter) MountFilesystemReturns(result1 error) {
	fake.mountFilesystemMutex.Lock()
	defer fake.mountFilesystemMutex.Unlock()
	fake.MountFilesystemStub = nil
	fake.mountFilesystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) MountFilesystemReturnsOnCall(i int, result1 error) {
	fake.mountFilesystemMutex.Lock()
	defer fake.mountFilesystemMutex.Unlock()
	fake.MountFilesystemStub = nil
	if fake.mountFilesystemReturnsOnCall == nil {
		fake.mountFilesystemReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountFilesystemReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) MountTmpfs(arg1 string, arg2 string) error {
	fake.mountTmpfsMutex.Lock()
	ret, specificReturn := fake.mountTmpfsReturnsOnCall[len(fake.mountTmpfsArgsForCall)]
	fake.mountTmpfsArgsForCall = append(fake.mountTmpfsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("MountTmpfs", []interface{}{arg1, arg2})
	fake.mountTmpfsMutex.Unlock()
	if fake.MountTmpfsStub != nil {
		return fake.MountTmpfsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mountTmpfsReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) MountTmpfsCallCount() int {
	fake.mountTmpfsMutex.RLock()
	defer fake.mountTmpfsMutex.RUnlock()
	return len(fake.mountTmpfsArgsForCall)
}

func (fake *FakeMounter) MountTmpfsCalls(stub func(string, string) error) {
	fake.mountTmpfsMutex.Lock()
	defer fake.mountTmpfsMutex.Unlock()
	fake.MountTmpfsStub = stub
}

func (fake *FakeMounter) MountTmpfsArgsForCall(i int) (string, string) {
	fake.mountTmpfsMutex.RLock()
	defer fake.mountTmpfsMutex.RUnlock()
	argsForCall := fake.mountTmpfsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMounter) MountTmpfsReturns(result1 error) {
	fake.mountTmpfsMutex.Lock()
	defer fake.mountTmpfsMutex.Unlock()
	fake.MountTmpfsStub = nil
	fake.mountTmpfsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) MountTmpfsReturnsOnCall(i int, result1 error) {
	fake.mountTmpfsMutex.Lock()
	defer fake.mountTmpfsMutex.Unlock()
	fake.MountTmpfsStub = nil
	if fake.mountTmpfsReturnsOnCall == nil {
		fake.mountTmpfsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountTmpfsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) Remount(arg1 string, arg2 string, arg3 ...string) error {
	fake.remountMutex.Lock()
	ret, specificReturn := fake.remountReturnsOnCall[len(fake.remountArgsForCall)]
	fake.remountArgsForCall = append(fake.remountArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Remount", []interface{}{arg1, arg2, arg3})
	fake.remountMutex.Unlock()
	if fake.RemountStub != nil {
		return fake.RemountStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.remountReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) RemountCallCount() int {
	fake.remountMutex.RLock()
	defer fake.remountMutex.RUnlock()
	return len(fake.remountArgsForCall)
}

func (fake *FakeMounter) RemountCalls(stub func(string, string, ...string) error) {
	fake.remountMutex.Lock()
	defer fake.remountMutex.Unlock()
	fake.RemountStub = stub
}

func (fake *FakeMounter) RemountArgsForCall(i int) (string, string, []string) {
	fake.remountMutex.RLock()
	defer fake.remountMutex.RUnlock()
	argsForCall := fake.remountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeMounter) RemountReturns(result1 error) {
	fake.remountMutex.Lock()
	defer fake.remountMutex.Unlock()
	fake.RemountStub = nil
	fake.remountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) RemountReturnsOnCall(i int, result1 error) {
	fake.remountMutex.Lock()
	defer fake.remountMutex.Unlock()
	fake.RemountStub = nil
	if fake.remountReturnsOnCall == nil {
		fake.remountReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.remountReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) RemountAsReadonly(arg1 string) error {
	fake.remountAsReadonlyMutex.Lock()
	ret, specificReturn := fake.remountAsReadonlyReturnsOnCall[len(fake.remountAsReadonlyArgsForCall)]
	fake.remountAsReadonlyArgsForCall = append(fake.remountAsReadonlyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemountAsReadonly", []interface{}{arg1})
	fake.remountAsReadonlyMutex.Unlock()
	if fake.RemountAsReadonlyStub != nil {
		return fake.RemountAsReadonlyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.remountAsReadonlyReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) RemountAsReadonlyCallCount() int {
	fake.remountAsReadonlyMutex.RLock()
	defer fake.remountAsReadonlyMutex.RUnlock()
	return len(fake.remountAsReadonlyArgsForCall)
}

func (fake *FakeMounter) RemountAsReadonlyCalls(stub func(string) error) {
	fake.remountAsReadonlyMutex.Lock()
	defer fake.remountAsReadonlyMutex.Unlock()
	fake.RemountAsReadonlyStub = stub
}

func (fake *FakeMounter) RemountAsReadonlyArgsForCall(i int) string {
	fake.remountAsReadonlyMutex.RLock()
	defer fake.remountAsReadonlyMutex.RUnlock()
	argsForCall := fake.remountAsReadonlyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMounter) RemountAsReadonlyReturns(result1 error) {
	fake.remountAsReadonlyMutex.Lock()
	defer fake.remountAsReadonlyMutex.Unlock()
	fake.RemountAsReadonlyStub = nil
	fake.remountAsReadonlyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) RemountAsReadonlyReturnsOnCall(i int, result1 error) {
	fake.remountAsReadonlyMutex.Lock()
	defer fake.remountAsReadonlyMutex.Unlock()
	fake.RemountAsReadonlyStub = nil
	if fake.remountAsReadonlyReturnsOnCall == nil {
		fake.remountAsReadonlyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.remountAsReadonlyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) RemountInPlace(arg1 string, arg2 ...string) error {
	fake.remountInPlaceMutex.Lock()
	ret, specificReturn := fake.remountInPlaceReturnsOnCall[len(fake.remountInPlaceArgsForCall)]
	fake.remountInPlaceArgsForCall = append(fake.remountInPlaceArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("RemountInPlace", []interface{}{arg1, arg2})
	fake.remountInPlaceMutex.Unlock()
	if fake.RemountInPlaceStub != nil {
		return fake.RemountInPlaceStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.remountInPlaceReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) RemountInPlaceCallCount() int {
	fake.remountInPlaceMutex.RLock()
	defer fake.remountInPlaceMutex.RUnlock()
	return len(fake.remountInPlaceArgsForCall)
}

func (fake *FakeMounter) RemountInPlaceCalls(stub func(string, ...string) error) {
	fake.remountInPlaceMutex.Lock()
	defer fake.remountInPlaceMutex.Unlock()
	fake.RemountInPlaceStub = stub
}

func (fake *FakeMounter) RemountInPlaceArgsForCall(i int) (string, []string) {
	fake.remountInPlaceMutex.RLock()
	defer fake.remountInPlaceMutex.RUnlock()
	argsForCall := fake.remountInPlaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMounter) RemountInPlaceReturns(result1 error) {
	fake.remountInPlaceMutex.Lock()
	defer fake.remountInPlaceMutex.Unlock()
	fake.RemountInPlaceStub = nil
	fake.remountInPlaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) RemountInPlaceReturnsOnCall(i int, result1 error) {
	fake.remountInPlaceMutex.Lock()
	defer fake.remountInPlaceMutex.Unlock()
	fake.RemountInPlaceStub = nil
	if fake.remountInPlaceReturnsOnCall == nil {
		fake.remountInPlaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.remountInPlaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) SwapOn(arg1 string) error {
	fake.swapOnMutex.Lock()
	ret, specificReturn := fake.swapOnReturnsOnCall[len(fake.swapOnArgsForCall)]
	fake.swapOnArgsForCall = append(fake.swapOnArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SwapOn", []interface{}{arg1})
	fake.swapOnMutex.Unlock()
	if fake.SwapOnStub != nil {
		return fake.SwapOnStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.swapOnReturns
	return fakeReturns.result1
}

func (fake *FakeMounter) SwapOnCallCount() int {
	fake.swapOnMutex.RLock()
	defer fake.swapOnMutex.RUnlock()
	return len(fake.swapOnArgsForCall)
}

func (fake *FakeMounter) SwapOnCalls(stub func(string) error) {
	fake.swapOnMutex.Lock()
	defer fake.swapOnMutex.Unlock()
	fake.SwapOnStub = stub
}

func (fake *FakeMounter) SwapOnArgsForCall(i int) string {
	fake.swapOnMutex.RLock()
	defer fake.swapOnMutex.RUnlock()
	argsForCall := fake.swapOnArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMounter) SwapOnReturns(result1 error) {
	fake.swapOnMutex.Lock()
	defer fake.swapOnMutex.Unlock()
	fake.SwapOnStub = nil
	fake.swapOnReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) SwapOnReturnsOnCall(i int, result1 error) {
	fake.swapOnMutex.Lock()
	defer fake.swapOnMutex.Unlock()
	fake.SwapOnStub = nil
	if fake.swapOnReturnsOnCall == nil {
		fake.swapOnReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.swapOnReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) Unmount(arg1 string) (bool, error) {
	fake.unmountMutex.Lock()
	ret, specificReturn := fake.unmountReturnsOnCall[len(fake.unmountArgsForCall)]
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Unmount", []interface{}{arg1})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unmountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMounter) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeMounter) UnmountCalls(stub func(string) (bool, error)) {
	fake.unmountMutex.Lock()
	defer fake.unmountMutex.Unlock()
	fake.UnmountStub = stub
}

func (fake *FakeMounter) UnmountArgsForCall(i int) string {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	argsForCall := fake.unmountArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMounter) UnmountReturns(result1 bool, result2 error) {
	fake.unmountMutex.Lock()
	defer fake.unmountMutex.Unlock()
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMounter) UnmountReturnsOnCall(i int, result1 bool, result2 error) {
	fake.unmountMutex.Lock()
	defer fake.unmountMutex.Unlock()
	fake.UnmountStub = nil
	if fake.unmountReturnsOnCall == nil {
		fake.unmountReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.unmountReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMounter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isMountPointMutex.RLock()
	defer fake.isMountPointMutex.RUnlock()
	fake.isMountedMutex.RLock()
	defer fake.isMountedMutex.RUnlock()
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	fake.mountFilesystemMutex.RLock()
	defer fake.mountFilesystemMutex.RUnlock()
	fake.mountTmpfsMutex.RLock()
	defer fake.mountTmpfsMutex.RUnlock()
	fake.remountMutex.RLock()
	defer fake.remountMutex.RUnlock()
	fake.remountAsReadonlyMutex.RLock()
	defer fake.remountAsReadonlyMutex.RUnlock()
	fake.remountInPlaceMutex.RLock()
	defer fake.remountInPlaceMutex.RUnlock()
	fake.swapOnMutex.RLock()
	defer fake.swapOnMutex.RUnlock()
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMounter) recordInvocation(key string, args []interface{}) {
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

var _ disk.Mounter = new(FakeMounter)
