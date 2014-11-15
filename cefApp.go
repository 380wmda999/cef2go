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

import "unsafe"


var (
    knownAppHandlers       = make(map[unsafe.Pointer]AppHandler)
)

// create the underlying C structure for an App Handler.
func NewAppHandlerT(handler AppHandler) AppHandlerT {

    var a AppHandlerT
    a.CStruct = (*C.cef_app_t)(
           C.calloc(1, C.sizeof_cef_app_t))
    C.initialize_app_handler(a.CStruct)
    Logger.Infof("_AppHandler: %x", unsafe.Pointer(a.CStruct))

    knownAppHandlers[unsafe.Pointer(a.CStruct)] = handler
    return a
}


//export goDebugLog
func goDebugLog(toLog *C.char) {
    ll := C.GoString(toLog)
    Logger.Infof(ll)
}

//export go_GetBrowserProcessHandler
func go_GetBrowserProcessHandler(self *C.cef_app_t) *C.struct__cef_browser_process_handler_t {
    if handler, ok := knownAppHandlers[unsafe.Pointer(self)]; ok {
        bph := handler.GetBrowserProcessHandler()
        if bph != nil {
            return bph.GetBrowserProcessHandlerT().CStruct
        }
    }
    return nil
}


type AppHandler interface {
    OnBeforeCommandLineProcessing(processType string, commandLine CommandLineT)
    OnRegisterCustomSchemes()
    GetResourceBundleHandler()
    GetBrowserProcessHandler() BrowserProcessHandler
    GetRenderProcessHandler()
    // called to get the underlying c struct.
    GetAppHandlerT() AppHandlerT
}
type AppHandlerT struct {
    CStruct         *C.cef_app_t
}