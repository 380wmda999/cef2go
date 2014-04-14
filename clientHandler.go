// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go
// Website: https://github.com/CzarekTomczak/cef2go

package cef2go


/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
extern void initialize_client_handler(struct _cef_client_t* client);
*/
import "C"

import (
//    "fmt"
    "unsafe"
)


var _ClientHandler *C.struct__cef_client_t // requires reference counting


//export go_GetContextMenuHandler
func go_GetContextMenuHandler(self *C.struct__cef_client_t) *C.struct__cef_context_menu_handler_t {
    return nil
}

//export go_GetDialogHandler
func go_GetDialogHandler(self *C.struct__cef_client_t) *C.struct__cef_dialog_handler_t {
    return nil
}

//export go_GetDisplayHandler
func go_GetDisplayHandler(self *C.struct__cef_client_t) *C.struct__cef_display_handler_t {
    //Logger.Infof("Add ref to displayHandler\n")
    go_AddRef(unsafe.Pointer(_DisplayHandler))
    return _DisplayHandler
    //return nil
}

//export go_GetDownloadHandler
func go_GetDownloadHandler(self *C.struct__cef_client_t) *C.struct__cef_download_handler_t {
    return nil
}

//export go_GetDragHandler
func go_GetDragHandler(self *C.struct__cef_client_t) *C.struct__cef_drag_handler_t {
    return nil
}

//export go_GetFocusHandler
func go_GetFocusHandler(self *C.struct__cef_client_t) *C.struct__cef_focus_handler_t {
    return nil
}

//export go_GetGeoLocationHandler
func go_GetGeoLocationHandler(self *C.struct__cef_client_t) *C.struct__cef_geolocation_handler_t {
    return nil
}

//export go_GetJsDialogHandler
func go_GetJsDialogHandler(self *C.struct__cef_client_t) *C.struct__cef_jsdialog_handler_t {
    return nil
}

//export go_GetKeyboardHandler
func go_GetKeyboardHandler(self *C.struct__cef_client_t) *C.struct__cef_keyboard_handler_t {
    return nil
}

//export go_GetLifespanHandler
func go_GetLifespanHandler(self *C.struct__cef_client_t) *C.struct__cef_life_span_handler_t {
    go_AddRef(unsafe.Pointer(_LifeSpanHandler))
    return _LifeSpanHandler
}

//export go_GetLoadHandler
func go_GetLoadHandler(self *C.struct__cef_client_t) *C.struct__cef_load_handler_t {
    return nil
}

//export go_GetRenderHandler
func go_GetRenderHandler(self *C.struct__cef_client_t) *C.struct__cef_render_handler_t {
    return nil
}

//export go_GetRequestHandler
func go_GetRequestHandler(self *C.struct__cef_client_t) *C.struct__cef_request_handler_t {
    go_AddRef(unsafe.Pointer(_RequestHandler))
    return _RequestHandler
}



func InitializeHandler() *C.struct__cef_client_t {
    var handler *C.struct__cef_client_t
    handler = (*C.struct__cef_client_t)(
            C.calloc(1, C.sizeof_struct__cef_client_t))
    C.initialize_client_handler(handler)
    go_AddRef(unsafe.Pointer(handler))
    Logger.Infof("_ClientHandler: %x", unsafe.Pointer(handler))
    return handler
}

