// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package cef2go


/*
#include "cefBase.h"
#include "include/capi/cef_callback_capi.h"

extern void cef_callback_t_cont(struct _cef_callback_t* self);
extern void cef_callback_t_cancel(struct _cef_callback_t* self);
*/
import "C"

import (
    "unsafe"
)

type CefCallbackT struct {
    CStruct         *C.struct__cef_callback_t
}



func (b CefCallbackT) Release() {
    C.releaseVoid(unsafe.Pointer(b.CStruct))
}

func (b CefCallbackT) AddRef() {
    C.add_refVoid(unsafe.Pointer(b.CStruct))
}


// continue processing
func (c CefCallbackT) Cont() {
    C.cef_callback_t_cont(c.CStruct)
}

// cancel processing
func (c CefCallbackT) Cancel() {
    C.cef_callback_t_cancel(c.CStruct)
}