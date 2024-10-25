// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package logging represents the imported interface "wasmvision:platform/logging".
package logging

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Log represents the imported function "log".
//
// DEPRECATED: Log a message on the host.
// Instead use error, warn, info, or debug.
//
//	log: func(msg: string)
//
//go:nosplit
func Log(msg string) {
	msg0, msg1 := cm.LowerString(msg)
	wasmimport_Log((*uint8)(msg0), (uint32)(msg1))
	return
}

// Error represents the imported function "error".
//
// Log an error on the host.
//
//	error: func(msg: string)
//
//go:nosplit
func Error(msg string) {
	msg0, msg1 := cm.LowerString(msg)
	wasmimport_Error((*uint8)(msg0), (uint32)(msg1))
	return
}

// Warn represents the imported function "warn".
//
// Log a warning on the host.
//
//	warn: func(msg: string)
//
//go:nosplit
func Warn(msg string) {
	msg0, msg1 := cm.LowerString(msg)
	wasmimport_Warn((*uint8)(msg0), (uint32)(msg1))
	return
}

// Info represents the imported function "info".
//
// Log some non-critical information on the host.
//
//	info: func(msg: string)
//
//go:nosplit
func Info(msg string) {
	msg0, msg1 := cm.LowerString(msg)
	wasmimport_Info((*uint8)(msg0), (uint32)(msg1))
	return
}

// Debug represents the imported function "debug".
//
// Log some debugging info on the host.
//
//	debug: func(msg: string)
//
//go:nosplit
func Debug(msg string) {
	msg0, msg1 := cm.LowerString(msg)
	wasmimport_Debug((*uint8)(msg0), (uint32)(msg1))
	return
}

// Println represents the imported function "println".
//
// Print a message on the host. Intended to bypassing the normal logging system.
//
//	println: func(msg: string)
//
//go:nosplit
func Println(msg string) {
	msg0, msg1 := cm.LowerString(msg)
	wasmimport_Println((*uint8)(msg0), (uint32)(msg1))
	return
}
