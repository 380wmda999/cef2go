package cef2go




/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"
extern void initialize_request_handler(struct _cef_request_handler_t * requestHandler);
extern char * cef_request_get_url(struct _cef_request_t * self);
*/
import "C"
import (
    "unsafe"
    "fmt"
)

var _RequestHandler *C.struct__cef_request_handler_t // requires reference counting


//export go_OnBeforeBrowse
func go_OnBeforeBrowse(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    frame *C.struct__cef_frame_t,
    request *C.struct__cef_request_t,
    is_redirect int) int {

    cString := C.cef_request_get_url(request)
    fmt.Printf("go_OnBeforeBrowse: %p\n", cString)
    str := C.GoString(cString)
    defer C.free(unsafe.Pointer(cString))
    fmt.Printf("go_OnBeforeBrowse.url: %s\n", str)

    return 0
}

//export go_OnBeforeResourceLoad
func go_OnBeforeResourceLoad(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    frame *C.struct__cef_frame_t,
    request *C.struct__cef_request_t) int {
    return 0
}

//export go_GetResourceHandler
func go_GetResourceHandler(
        self *C.struct__cef_request_handler_t,
        browser *C.struct__cef_browser_t,
        frame *C.struct__cef_frame_t,
        request *C.struct__cef_request_t) *C.struct__cef_resource_handler_t {
    return nil
}

//export go_OnResourceRedirect
func go_OnResourceRedirect(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    frame *C.struct__cef_frame_t,
    old_url * C.char,
    new_url * C.char) {
}

//export go_GetAuthCredentials
func go_GetAuthCredentials(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    frame *C.struct__cef_frame_t,
    isProxy int,
    host * C.char,
    port int,
    realm * C.char,
    scheme * C.char,
    callback *C.struct__cef_auth_callback_t) int {
    return 1
}

//export go_OnQuotaRequest
func go_OnQuotaRequest(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    origin_url *C.char,
    new_size int64,
    callback *C.struct__cef_quota_callback_t) int {
    return 1
}


//export go_OnProtocolExecution
func go_OnProtocolExecution(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    url *C.char,
    allow_os_execution unsafe.Pointer) {

}


//export go_OnCertificateError
func go_OnCertificateError(
    self *C.struct__cef_request_handler_t,
    cert_error int, //C.enum_cef_errorcode_t,
    request_url *C.char,
    callback *C.struct__cef_allow_certificate_error_callback_t) int {
    return 0
}


//export go_OnBeforePluginLoad
func go_OnBeforePluginLoad(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    url *C.char,
    policy_url *C.char,
    info *C.struct__cef_web_plugin_info_t) int {
    return 1
}

//export go_OnPluginCrashed
func go_OnPluginCrashed(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    plugin_path *C.char) {

}

//export go_OnRenderProcessTerminated
func go_OnRenderProcessTerminated(
    self *C.struct__cef_request_handler_t,
    browser *C.struct__cef_browser_t,
    status int, //C.enum_cef_termination_status_t
    ) {

}


func InitializeRequestHandler() *C.struct__cef_request_handler_t {
    var handler *C.struct__cef_request_handler_t
    handler = (*C.struct__cef_request_handler_t)(
            C.calloc(1, C.sizeof_struct__cef_request_handler_t))
    C.initialize_request_handler(handler)
    return handler
}