// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/fromkeith/cef2go

package cef2go



/*
#include "include/capi/cef_app_capi.h"
extern void initialize_app_handler(cef_app_t* app);
*/
import "C"

import (
    "unsafe"
)

var _AppHandler *C.cef_app_t


func _InitializeGlobalCStructuresApp() {

    _AppHandler = (*C.cef_app_t)(
           C.calloc(1, C.sizeof_cef_app_t))
    C.initialize_app_handler(_AppHandler)
    Logger.Infof("_AppHandler: %x", unsafe.Pointer(_AppHandler))
}


//export goDebugLog
func goDebugLog(toLog *C.char) {
    ll := C.GoString(toLog)
    Logger.Infof(ll)
}