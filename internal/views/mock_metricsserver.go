// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/internal/resource (interfaces: MetricsServer)

package views

import (
	k8s "github.com/derailed/k9s/internal/k8s"
	pegomock "github.com/petergtz/pegomock"
	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"reflect"
	"time"
)

type MockMetricsServer struct {
	fail func(message string, callerSkip ...int)
}

func NewMockMetricsServer() *MockMetricsServer {
	return &MockMetricsServer{fail: pegomock.GlobalFailHandler}
}

func (mock *MockMetricsServer) ClusterLoad(_param0 []v1.Node, _param1 []v1beta1.NodeMetrics) k8s.ClusterMetrics {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ClusterLoad", params, []reflect.Type{reflect.TypeOf((*k8s.ClusterMetrics)(nil)).Elem()})
	var ret0 k8s.ClusterMetrics
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(k8s.ClusterMetrics)
		}
	}
	return ret0
}

func (mock *MockMetricsServer) FetchNodesMetrics() ([]v1beta1.NodeMetrics, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchNodesMetrics", params, []reflect.Type{reflect.TypeOf((*[]v1beta1.NodeMetrics)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []v1beta1.NodeMetrics
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]v1beta1.NodeMetrics)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsServer) FetchPodsMetrics(_param0 string) ([]v1beta1.PodMetrics, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchPodsMetrics", params, []reflect.Type{reflect.TypeOf((*[]v1beta1.PodMetrics)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []v1beta1.PodMetrics
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]v1beta1.PodMetrics)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsServer) HasMetrics() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasMetrics", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockMetricsServer) NodesMetrics(_param0 []v1.Node, _param1 []v1beta1.NodeMetrics, _param2 k8s.NodesMetrics) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	pegomock.GetGenericMockFrom(mock).Invoke("NodesMetrics", params, []reflect.Type{})
}

func (mock *MockMetricsServer) PodsMetrics(_param0 []v1beta1.PodMetrics, _param1 k8s.PodsMetrics) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("PodsMetrics", params, []reflect.Type{})
}

func (mock *MockMetricsServer) VerifyWasCalledOnce() *VerifierMetricsServer {
	return &VerifierMetricsServer{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockMetricsServer) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMetricsServer {
	return &VerifierMetricsServer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockMetricsServer) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMetricsServer {
	return &VerifierMetricsServer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockMetricsServer) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMetricsServer {
	return &VerifierMetricsServer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMetricsServer struct {
	mock                   *MockMetricsServer
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMetricsServer) ClusterLoad(_param0 []v1.Node, _param1 []v1beta1.NodeMetrics) *MetricsServer_ClusterLoad_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ClusterLoad", params, verifier.timeout)
	return &MetricsServer_ClusterLoad_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsServer_ClusterLoad_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsServer_ClusterLoad_OngoingVerification) GetCapturedArguments() ([]v1.Node, []v1beta1.NodeMetrics) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MetricsServer_ClusterLoad_OngoingVerification) GetAllCapturedArguments() (_param0 [][]v1.Node, _param1 [][]v1beta1.NodeMetrics) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]v1.Node, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]v1.Node)
		}
		_param1 = make([][]v1beta1.NodeMetrics, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]v1beta1.NodeMetrics)
		}
	}
	return
}

func (verifier *VerifierMetricsServer) FetchNodesMetrics() *MetricsServer_FetchNodesMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchNodesMetrics", params, verifier.timeout)
	return &MetricsServer_FetchNodesMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsServer_FetchNodesMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsServer_FetchNodesMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MetricsServer_FetchNodesMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMetricsServer) FetchPodsMetrics(_param0 string) *MetricsServer_FetchPodsMetrics_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchPodsMetrics", params, verifier.timeout)
	return &MetricsServer_FetchPodsMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsServer_FetchPodsMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsServer_FetchPodsMetrics_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MetricsServer_FetchPodsMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMetricsServer) HasMetrics() *MetricsServer_HasMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasMetrics", params, verifier.timeout)
	return &MetricsServer_HasMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsServer_HasMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsServer_HasMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MetricsServer_HasMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMetricsServer) NodesMetrics(_param0 []v1.Node, _param1 []v1beta1.NodeMetrics, _param2 k8s.NodesMetrics) *MetricsServer_NodesMetrics_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodesMetrics", params, verifier.timeout)
	return &MetricsServer_NodesMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsServer_NodesMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsServer_NodesMetrics_OngoingVerification) GetCapturedArguments() ([]v1.Node, []v1beta1.NodeMetrics, k8s.NodesMetrics) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MetricsServer_NodesMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 [][]v1.Node, _param1 [][]v1beta1.NodeMetrics, _param2 []k8s.NodesMetrics) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]v1.Node, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]v1.Node)
		}
		_param1 = make([][]v1beta1.NodeMetrics, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]v1beta1.NodeMetrics)
		}
		_param2 = make([]k8s.NodesMetrics, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(k8s.NodesMetrics)
		}
	}
	return
}

func (verifier *VerifierMetricsServer) PodsMetrics(_param0 []v1beta1.PodMetrics, _param1 k8s.PodsMetrics) *MetricsServer_PodsMetrics_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PodsMetrics", params, verifier.timeout)
	return &MetricsServer_PodsMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsServer_PodsMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsServer_PodsMetrics_OngoingVerification) GetCapturedArguments() ([]v1beta1.PodMetrics, k8s.PodsMetrics) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MetricsServer_PodsMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 [][]v1beta1.PodMetrics, _param1 []k8s.PodsMetrics) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]v1beta1.PodMetrics, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]v1beta1.PodMetrics)
		}
		_param1 = make([]k8s.PodsMetrics, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(k8s.PodsMetrics)
		}
	}
	return
}
