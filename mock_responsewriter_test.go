// Automatically generated by pegomock. DO NOT EDIT!
// Source: net/http (interfaces: ResponseWriter)

package main_test

import (
	http "net/http"
	"reflect"

	pegomock "github.com/petergtz/pegomock"
)

type MockResponseWriter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockResponseWriter() *MockResponseWriter {
	return &MockResponseWriter{fail: pegomock.GlobalFailHandler}
}

func (mock *MockResponseWriter) Header() http.Header {
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Header", params, []reflect.Type{reflect.TypeOf((*http.Header)(nil)).Elem()})
	var ret0 http.Header
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(http.Header)
		}
	}
	return ret0
}

func (mock *MockResponseWriter) Write(_param0 []byte) (int, error) {
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Write", params, []reflect.Type{reflect.TypeOf((*int)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 int
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(int)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockResponseWriter) WriteHeader(_param0 int) {
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("WriteHeader", params, []reflect.Type{})
}

func (mock *MockResponseWriter) VerifyWasCalledOnce() *VerifierResponseWriter {
	return &VerifierResponseWriter{mock, pegomock.Times(1), nil}
}

func (mock *MockResponseWriter) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierResponseWriter {
	return &VerifierResponseWriter{mock, invocationCountMatcher, nil}
}

func (mock *MockResponseWriter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierResponseWriter {
	return &VerifierResponseWriter{mock, invocationCountMatcher, inOrderContext}
}

type VerifierResponseWriter struct {
	mock                   *MockResponseWriter
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierResponseWriter) Header() *ResponseWriter_Header_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Header", params)
	return &ResponseWriter_Header_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ResponseWriter_Header_OngoingVerification struct {
	mock              *MockResponseWriter
	methodInvocations []pegomock.MethodInvocation
}

func (c *ResponseWriter_Header_OngoingVerification) GetCapturedArguments() {
}

func (c *ResponseWriter_Header_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierResponseWriter) Write(_param0 []byte) *ResponseWriter_Write_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Write", params)
	return &ResponseWriter_Write_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ResponseWriter_Write_OngoingVerification struct {
	mock              *MockResponseWriter
	methodInvocations []pegomock.MethodInvocation
}

func (c *ResponseWriter_Write_OngoingVerification) GetCapturedArguments() []byte {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *ResponseWriter_Write_OngoingVerification) GetAllCapturedArguments() (_param0 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]byte, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]byte)
		}
	}
	return
}

func (verifier *VerifierResponseWriter) WriteHeader(_param0 int) *ResponseWriter_WriteHeader_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "WriteHeader", params)
	return &ResponseWriter_WriteHeader_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ResponseWriter_WriteHeader_OngoingVerification struct {
	mock              *MockResponseWriter
	methodInvocations []pegomock.MethodInvocation
}

func (c *ResponseWriter_WriteHeader_OngoingVerification) GetCapturedArguments() int {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *ResponseWriter_WriteHeader_OngoingVerification) GetAllCapturedArguments() (_param0 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
	}
	return
}