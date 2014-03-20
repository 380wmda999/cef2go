// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/fromkeith/cef2go

package cef2go


/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
var _MainArgs *C.struct__cef_main_args_t


const (
    LOGSEVERITY_DEFAULT = C.LOGSEVERITY_DEFAULT
    LOGSEVERITY_VERBOSE = C.LOGSEVERITY_VERBOSE
    LOGSEVERITY_INFO = C.LOGSEVERITY_INFO
    LOGSEVERITY_WARNING = C.LOGSEVERITY_WARNING
    LOGSEVERITY_ERROR = C.LOGSEVERITY_ERROR
    LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
    LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)


func _InitializeGlobalCStructuresBase() {
     _MainArgs = (*C.struct__cef_main_args_t)(
            C.calloc(1, C.sizeof_struct__cef_main_args_t))
}
