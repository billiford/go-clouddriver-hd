// Code generated by counterfeiter. DO NOT EDIT.
package kubernetesfakes

import (
	"sync"

	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FakeClient struct {
	ApplyStub        func(*unstructured.Unstructured) (kubernetes.Metadata, error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 *unstructured.Unstructured
	}
	applyReturns struct {
		result1 kubernetes.Metadata
		result2 error
	}
	applyReturnsOnCall map[int]struct {
		result1 kubernetes.Metadata
		result2 error
	}
	GetStub        func(string, string, string) (*unstructured.Unstructured, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getReturns struct {
		result1 *unstructured.Unstructured
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *unstructured.Unstructured
		result2 error
	}
	ListStub        func(schema.GroupVersionResource, metav1.ListOptions) (*unstructured.UnstructuredList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 schema.GroupVersionResource
		arg2 metav1.ListOptions
	}
	listReturns struct {
		result1 *unstructured.UnstructuredList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *unstructured.UnstructuredList
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Apply(arg1 *unstructured.Unstructured) (kubernetes.Metadata, error) {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 *unstructured.Unstructured
	}{arg1})
	fake.recordInvocation("Apply", []interface{}{arg1})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.applyReturns.result1, fake.applyReturns.result2
}

func (fake *FakeClient) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeClient) ApplyArgsForCall(i int) *unstructured.Unstructured {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].arg1
}

func (fake *FakeClient) ApplyReturns(result1 kubernetes.Metadata, result2 error) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 kubernetes.Metadata
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ApplyReturnsOnCall(i int, result1 kubernetes.Metadata, result2 error) {
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Metadata
			result2 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 kubernetes.Metadata
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Get(arg1 string, arg2 string, arg3 string) (*unstructured.Unstructured, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeClient) GetArgsForCall(i int) (string, string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3
}

func (fake *FakeClient) GetReturns(result1 *unstructured.Unstructured, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *unstructured.Unstructured
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetReturnsOnCall(i int, result1 *unstructured.Unstructured, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *unstructured.Unstructured
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *unstructured.Unstructured
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) List(arg1 schema.GroupVersionResource, arg2 metav1.ListOptions) (*unstructured.UnstructuredList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 schema.GroupVersionResource
		arg2 metav1.ListOptions
	}{arg1, arg2})
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakeClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeClient) ListArgsForCall(i int) (schema.GroupVersionResource, metav1.ListOptions) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].arg1, fake.listArgsForCall[i].arg2
}

func (fake *FakeClient) ListReturns(result1 *unstructured.UnstructuredList, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *unstructured.UnstructuredList
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListReturnsOnCall(i int, result1 *unstructured.UnstructuredList, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *unstructured.UnstructuredList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *unstructured.UnstructuredList
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ kubernetes.Client = new(FakeClient)