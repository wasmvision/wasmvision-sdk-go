// Code generated by wit-bindgen-go. DO NOT EDIT.

package http

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasmvision:platform".

//go:wasmimport wasmvision:platform/http get
//go:noescape
func wasmimport_Get(url0 *uint8, url1 uint32, result *cm.Result[cm.List[uint8], cm.List[uint8], HTTPError])

//go:wasmimport wasmvision:platform/http post
//go:noescape
func wasmimport_Post(url0 *uint8, url1 uint32, contentType0 *uint8, contentType1 uint32, body0 *uint8, body1 uint32, result *cm.Result[cm.List[uint8], cm.List[uint8], HTTPError])

//go:wasmimport wasmvision:platform/http post-image
//go:noescape
func wasmimport_PostImage(url0 *uint8, url1 uint32, contentType0 *uint8, contentType1 uint32, requestTemplate0 *uint8, requestTemplate1 uint32, responseItem0 *uint8, responseItem1 uint32, mat0 uint32, result *cm.Result[cm.List[uint8], cm.List[uint8], HTTPError])
