// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

package cef2go


/*
#include "cefBase.h"
#include "include/capi/cef_resource_handler_capi.h"

extern void intialize_cef_resource_handler(struct _cef_resource_handler_t* handler);
*/
import "C"

import "unsafe"
import "fmt"

type CefResourceHandlerT struct {
    CStruct         *C.struct__cef_resource_handler_t
    resourceHandler     ResourceHandler
}

type CefCookieT struct {
    CStruct         *C.struct__cef_cookie_t
}

type ResourceHandler interface {
    ///
    // Begin processing the request. To handle the request return true (1) and
    // call cef_callback_t::cont() once the response header information is
    // available (cef_callback_t::cont() can also be called from inside this
    // function if header information is available immediately). To cancel the
    // request return false (0).
    ///
    ProcessRequest(request CefRequestT, callback CefCallbackT) int
    ///
    // Retrieve response header information. If the response length is not known
    // set |response_length| to -1 and read_response() will be called until it
    // returns false (0). If the response length is known set |response_length| to
    // a positive value and read_response() will be called until it returns false
    // (0) or the specified number of bytes have been read. Use the |response|
    // object to set the mime type, http status code and other optional header
    // values. To redirect the request to a new URL set |redirectUrl| to the new
    // URL.
    ///
    GetResponseHeaders(response CefResponseT) (responseLength int64)
    ///
    // Read response data. If data is available immediately copy up to
    // |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
    // bytes copied, and return true (1). To read the data at a later time set
    // |bytes_read| to 0, return true (1) and call cef_callback_t::cont() when the
    // data is available. To indicate response completion return false (0).
    ///
    ReadResponse(bytesToRead int, callback CefCallbackT) (dataOut []byte, bytesRead int, result int)
    ///
    // Return true (1) if the specified cookie can be sent with the request or
    // false (0) otherwise. If false (0) is returned for any cookie then no
    // cookies will be sent with the request.
    ///
    GetCookie(constCookie CefCookieT) int
    ///
    // Return true (1) if the specified cookie returned with the response can be
    // set or false (0) otherwise.
    ///
    SetCookie(constCookie CefCookieT) int
    ///
    // Request processing has been canceled.
    ///
    Cancel()
}

var (
    handlerMap = make(map[unsafe.Pointer]CefResourceHandlerT)
)


//export go_ProcessRequest
func go_ProcessRequest(
    self *C.struct__cef_resource_handler_t,
    request *C.struct__cef_request_t,
    callback *C.struct__cef_callback_t) int {

    if handler, ok := handlerMap[unsafe.Pointer(self)]; ok {
        return handler.resourceHandler.ProcessRequest(
            CefRequestT{request},
            CefCallbackT{callback},
        )
    }

    return 0
}

//export go_GetResponseHeaders
func go_GetResponseHeaders(
  self *C.struct__cef_resource_handler_t,
  response *C.struct__cef_response_t,
  response_length *C.int64) {

    if handler, ok := handlerMap[unsafe.Pointer(self)]; ok {
        *response_length = C.int64(handler.resourceHandler.GetResponseHeaders(
            CefResponseT{response},
        ))
        return
    }

    *response_length = 0
}

//export go_ReadResponse
func go_ReadResponse(
    self *C.struct__cef_resource_handler_t,
    data_out unsafe.Pointer,
    bytes_to_read C.int,
    bytes_read *C.int,
    callback *C.struct__cef_callback_t) int {


    if handler, ok := handlerMap[unsafe.Pointer(self)]; ok {
        dataOut, bytesRead, result := handler.resourceHandler.ReadResponse(
            int(bytes_to_read),
            CefCallbackT{callback},
        )
        if len(dataOut) != bytesRead {
            panic(fmt.Sprintf("The response given to ReadResponse is invalid. bytes given %d is does not equal bytesRead %d", len(dataOut), bytesRead))
        }
        dataOutCString := C.CString(string(dataOut))
        defer C.free(unsafe.Pointer(dataOutCString))
        C.memcpy(data_out, unsafe.Pointer(dataOutCString), C.size_t(bytesRead))
        *bytes_read = C.int(bytesRead)
        return result
    }

    return 0

}

//export go_GetCookie
func go_GetCookie(
    self *C.struct__cef_resource_handler_t,
    constCookie *C.struct__cef_cookie_t) int {
    if handler, ok := handlerMap[unsafe.Pointer(self)]; ok {
        return handler.resourceHandler.GetCookie(CefCookieT{constCookie})
    }
    return 0
}

//export go_SetCookie
func go_SetCookie(
    self *C.struct__cef_resource_handler_t,
    constCookie *C.struct__cef_cookie_t) int {

    if handler, ok := handlerMap[unsafe.Pointer(self)]; ok {
        return handler.resourceHandler.SetCookie(CefCookieT{constCookie})
    }

    return 0
}

//export go_Cancel
func go_Cancel(self *C.struct__cef_resource_handler_t) {

}

func deleteResourceHandler(it unsafe.Pointer) {
    // delete it from our map
    delete(handlerMap, it)
}


func CreateResourceHandler(resHandler ResourceHandler) CefResourceHandlerT {
    var handler CefResourceHandlerT
    handler.resourceHandler = resHandler
    handler.CStruct = (*C.struct__cef_resource_handler_t)(
            C.calloc(1, C.sizeof_struct__cef_resource_handler_t))
    C.intialize_cef_resource_handler(handler.CStruct)
    unsafeIt := unsafe.Pointer(handler.CStruct)
    go_AddRef(unsafeIt)
    handlerMap[unsafeIt] = handler
    // register the destructor so we can properly remove it from our internal map
    RegisterDestructor(unsafeIt, deleteResourceHandler)

    return handler
}