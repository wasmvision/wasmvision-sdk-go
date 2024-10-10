// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package http represents the imported interface "wasmvision:platform/http".
package http

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// HTTPError represents the enum "wasmvision:platform/http#http-error".
//
// HTTP errors returned by the runtime.
//
//	enum http-error {
//		success,
//		destination-not-allowed,
//		invalid-url,
//		request-error,
//		runtime-error,
//		too-many-requests
//	}
type HTTPError uint8

const (
	HTTPErrorSuccess HTTPError = iota
	HTTPErrorDestinationNotAllowed
	HTTPErrorInvalidURL
	HTTPErrorRequestError
	HTTPErrorRuntimeError
	HTTPErrorTooManyRequests
)

var stringsHTTPError = [6]string{
	"success",
	"destination-not-allowed",
	"invalid-url",
	"request-error",
	"runtime-error",
	"too-many-requests",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e HTTPError) String() string {
	return stringsHTTPError[e]
}

// Get represents the imported function "get".
//
// Get the content at the specified URL.
// Returns either the content or an error.
//
//	get: func(url: string) -> result<list<u8>, http-error>
//
//go:nosplit
func Get(url string) (result cm.Result[cm.List[uint8], cm.List[uint8], HTTPError]) {
	url0, url1 := cm.LowerString(url)
	wasmimport_Get((*uint8)(url0), (uint32)(url1), &result)
	return
}

//go:wasmimport wasmvision:platform/http get
//go:noescape
func wasmimport_Get(url0 *uint8, url1 uint32, result *cm.Result[cm.List[uint8], cm.List[uint8], HTTPError])

// Post represents the imported function "post".
//
// Post the content to the specified URL.
// Returns either the response content or an error.
//
//	post: func(url: string, content-type: string, body: list<u8>) -> result<list<u8>,
//	http-error>
//
//go:nosplit
func Post(url string, contentType string, body cm.List[uint8]) (result cm.Result[cm.List[uint8], cm.List[uint8], HTTPError]) {
	url0, url1 := cm.LowerString(url)
	contentType0, contentType1 := cm.LowerString(contentType)
	body0, body1 := cm.LowerList(body)
	wasmimport_Post((*uint8)(url0), (uint32)(url1), (*uint8)(contentType0), (uint32)(contentType1), (*uint8)(body0), (uint32)(body1), &result)
	return
}

//go:wasmimport wasmvision:platform/http post
//go:noescape
func wasmimport_Post(url0 *uint8, url1 uint32, contentType0 *uint8, contentType1 uint32, body0 *uint8, body1 uint32, result *cm.Result[cm.List[uint8], cm.List[uint8], HTTPError])
