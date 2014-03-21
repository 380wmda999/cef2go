// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package cef2go

/*
#include <stdlib.h>
#include "cefBase.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
extern void initialize_life_span_handler(struct _cef_life_span_handler_t* lifeHandler);
*/
import "C"

import (
    "unsafe"
)


type LifeSpanHandler interface {
    OnAfterCreated(browser CefBrowserT)
    RunModal(browser CefBrowserT) int
    DoClose(browser CefBrowserT) int
    BeforeClose(browser CefBrowserT)
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler LifeSpanHandler


//export go_OnBeforePopup
func go_OnBeforePopup(
        self * C.struct__cef_life_span_handler_t,
        browser * C.struct__cef_browser_t,
        frame * C.struct__cef_frame_t,
        target_url *C.cef_string_t,
        target_frame_name *C.cef_string_t,
        popupFeatures * C.struct__cef_popup_features_t,
        windowInfo * C.struct__cef_window_info_t,
        client ** C.struct__cef_client_t,
        settings * C.struct__cef_browser_settings_t,
        no_javascript_access * C.int) int {

    C.releaseVoid(unsafe.Pointer(browser))
    C.releaseVoid(unsafe.Pointer(frame))
    C.releaseVoid(unsafe.Pointer(popupFeatures))
    C.releaseVoid(unsafe.Pointer(windowInfo))
    //C.releaseVoid(unsafe.Pointer(client))
    C.releaseVoid(unsafe.Pointer(settings))
    return 1
}

//export go_OnAfterCreated
func go_OnAfterCreated(
        self *C.struct__cef_life_span_handler_t,
        browser *C.struct__cef_browser_t) {
    if globalLifespanHandler != nil {
        globalLifespanHandler.OnAfterCreated(CefBrowserT{browser})
    } else {
        C.releaseVoid(unsafe.Pointer(browser))
    }
}

//export go_RunModal
func go_RunModal(
        self *C.struct__cef_life_span_handler_t,
        browser *C.struct__cef_browser_t) int {
    if globalLifespanHandler != nil {
        return globalLifespanHandler.RunModal(CefBrowserT{browser})
    } else {
        C.releaseVoid(unsafe.Pointer(browser))
    }
    return 0
}

//export go_DoClose
func go_DoClose(
        self *C.struct__cef_life_span_handler_t,
        browser *C.struct__cef_browser_t) int {
    if globalLifespanHandler != nil {
        globalLifespanHandler.DoClose(CefBrowserT{browser})
    } else {
        C.releaseVoid(unsafe.Pointer(browser))
    }
    return 0
}

//export go_BeforeClose
func go_BeforeClose(
        self *C.struct__cef_life_span_handler_t,
        browser *C.struct__cef_browser_t) {
    if globalLifespanHandler != nil {
        globalLifespanHandler.BeforeClose(CefBrowserT{browser})
    } else {
        C.releaseVoid(unsafe.Pointer(browser))
    }
}



func InitializeLifeSpanHandler() *C.struct__cef_life_span_handler_t {
    var handler *C.struct__cef_life_span_handler_t
    handler = (*C.struct__cef_life_span_handler_t)(
            C.calloc(1, C.sizeof_struct__cef_life_span_handler_t))
    C.initialize_life_span_handler(handler)
    go_AddRef(unsafe.Pointer(handler))
    Logger.Println("_LifespanHandler: ", unsafe.Pointer(handler))
    return handler
}

func SetLifespanHandler(handler LifeSpanHandler) {
    globalLifespanHandler = handler
}