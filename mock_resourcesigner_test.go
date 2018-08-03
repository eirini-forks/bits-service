// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/cloudfoundry-incubator/bits-service (interfaces: ResourceSigner)

package bitsgo_test

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	time "time"
)

type MockResourceSigner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockResourceSigner() *MockResourceSigner {
	return &MockResourceSigner{fail: pegomock.GlobalFailHandler}
}

func (mock *MockResourceSigner) Sign(resource string, method string, expirationTime time.Time) string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMockResourceSigner().")
	}
	params := []pegomock.Param{resource, method, expirationTime}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Sign", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockResourceSigner) VerifyWasCalledOnce() *VerifierResourceSigner {
	return &VerifierResourceSigner{mock, pegomock.Times(1), nil}
}

func (mock *MockResourceSigner) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierResourceSigner {
	return &VerifierResourceSigner{mock, invocationCountMatcher, nil}
}

func (mock *MockResourceSigner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierResourceSigner {
	return &VerifierResourceSigner{mock, invocationCountMatcher, inOrderContext}
}

type VerifierResourceSigner struct {
	mock                   *MockResourceSigner
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierResourceSigner) Sign(resource string, method string, expirationTime time.Time) *ResourceSigner_Sign_OngoingVerification {
	params := []pegomock.Param{resource, method, expirationTime}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Sign", params)
	return &ResourceSigner_Sign_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ResourceSigner_Sign_OngoingVerification struct {
	mock              *MockResourceSigner
	methodInvocations []pegomock.MethodInvocation
}

func (c *ResourceSigner_Sign_OngoingVerification) GetCapturedArguments() (string, string, time.Time) {
	resource, method, expirationTime := c.GetAllCapturedArguments()
	return resource[len(resource)-1], method[len(method)-1], expirationTime[len(expirationTime)-1]
}

func (c *ResourceSigner_Sign_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []time.Time) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]time.Time, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(time.Time)
		}
	}
	return
}
