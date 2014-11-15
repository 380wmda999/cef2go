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

var (
    clientHandlerMap        = make(map[unsafe.Pointer]ClientHandler)
)

type ClientHandlerT struct {
    CStruct             *C.struct__cef_client_t
}

type ClientHandler interface {
    GetContextMenuHandler() ContextMenuHandlerT
    GetDialogHandler() DialogHandlerT
    GetDisplayHandler() DisplayHandler
    GetDownloadHandler() DownloadHandler
    GetDragHandler() DragHandlerT
    GetFocusHandler() FocusHandlerT
    GetGeoLocationHandler() GeolocationHandlerT
    GetJsDialogHandler() JsdialogHandlerT
    GetKeyboardHandler() KeyboardHandlerT
    GetLifeSpanHandler() LifeSpanHandler
    GetLoadHandler() LoadHandlerT
    GetRenderHandler() RenderHandlerT
    GetRequestHandler() RequestHandler

    GetClientHandlerT() ClientHandlerT
}


// these structs haven't been defined yet
type ContextMenuHandlerT struct {
    CStruct                 *C.struct__cef_context_menu_handler_t
}
type DialogHandlerT struct {
    CStruct                 *C.struct__cef_dialog_handler_t
}
type DragHandlerT struct {
    CStruct                 *C.struct__cef_drag_handler_t
}
type FocusHandlerT struct {
    CStruct                 *C.struct__cef_focus_handler_t
}
type GeolocationHandlerT struct {
    CStruct                 *C.struct__cef_geolocation_handler_t
}
type JsdialogHandlerT struct {
    CStruct                 *C.struct__cef_jsdialog_handler_t
}
type KeyboardHandlerT struct {
    CStruct                 *C.struct__cef_keyboard_handler_t
}
type LoadHandlerT struct {
    CStruct                 *C.struct__cef_load_handler_t
}
type RenderHandlerT struct {
    CStruct                 *C.struct__cef_render_handler_t
}


//export go_GetContextMenuHandler
func go_GetContextMenuHandler(self *C.struct__cef_client_t) *C.struct__cef_context_menu_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetContextMenuHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetDialogHandler
func go_GetDialogHandler(self *C.struct__cef_client_t) *C.struct__cef_dialog_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetDialogHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetDisplayHandler
func go_GetDisplayHandler(self *C.struct__cef_client_t) *C.struct__cef_display_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetDisplayHandler()
        if res != nil {
            res.GetDisplayHandlerT().AddRef()
            return res.GetDisplayHandlerT().CStruct
        }
    }
    return nil
}

//export go_GetDownloadHandler
func go_GetDownloadHandler(self *C.struct__cef_client_t) *C.struct__cef_download_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetDownloadHandler()
        if res != nil {
            res.GetDownloadHandlerT().AddRef()
            return res.GetDownloadHandlerT().CStruct
        }
    }
    return nil
}

//export go_GetDragHandler
func go_GetDragHandler(self *C.struct__cef_client_t) *C.struct__cef_drag_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetDragHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetFocusHandler
func go_GetFocusHandler(self *C.struct__cef_client_t) *C.struct__cef_focus_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetFocusHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetGeoLocationHandler
func go_GetGeoLocationHandler(self *C.struct__cef_client_t) *C.struct__cef_geolocation_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetGeoLocationHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetJsDialogHandler
func go_GetJsDialogHandler(self *C.struct__cef_client_t) *C.struct__cef_jsdialog_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetJsDialogHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetKeyboardHandler
func go_GetKeyboardHandler(self *C.struct__cef_client_t) *C.struct__cef_keyboard_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetKeyboardHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetLifespanHandler
func go_GetLifespanHandler(self *C.struct__cef_client_t) *C.struct__cef_life_span_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetLifeSpanHandler()
        if res != nil {
            res.GetLifeSpanHandlerT().AddRef()
            return res.GetLifeSpanHandlerT().CStruct
        }
    }
    return nil
}

//export go_GetLoadHandler
func go_GetLoadHandler(self *C.struct__cef_client_t) *C.struct__cef_load_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetLoadHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetRenderHandler
func go_GetRenderHandler(self *C.struct__cef_client_t) *C.struct__cef_render_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetRenderHandler()
        return res.CStruct
    }
    return nil
}

//export go_GetRequestHandler
func go_GetRequestHandler(self *C.struct__cef_client_t) *C.struct__cef_request_handler_t {
    if handler, ok := clientHandlerMap[unsafe.Pointer(self)]; ok {
        res := handler.GetRequestHandler()
        if res != nil {
            res.GetRequestHandlerT().AddRef()
            return res.GetRequestHandlerT().CStruct
        }
    }
    return nil
}



func NewClientHandlerT(handler ClientHandler) ClientHandlerT {
    var c ClientHandlerT
    c.CStruct = (*C.struct__cef_client_t)(
            C.calloc(1, C.sizeof_struct__cef_client_t))
    C.initialize_client_handler(c.CStruct)
    go_AddRef(unsafe.Pointer(c.CStruct))
    Logger.Infof("_ClientHandler: %x", unsafe.Pointer(c.CStruct))
    clientHandlerMap[unsafe.Pointer(c.CStruct)] = handler
    return c
}

