// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go
package cef2go

/*
#include "cefBase.h"
#include "include/capi/cef_scheme_capi.h"

extern void intialize_cef_scheme_handler_factory(struct _cef_scheme_handler_factory_t * factory);
extern int cef_scheme_handler_register(char * schemeName, char * domainName, struct _cef_scheme_handler_factory_t* factory);
extern int whitelist_cef_add_cross_origin_whitelist_entry(char* sourceOrigin, char* targetProtocol, char* targetDomain, int allow_target_subdomains);
*/
import "C"

import (
    "unsafe"
    //"fmt"
)

type SchemeHandlerFactory interface {
    CreateSchemeHandler(browser CefBrowserT, frame CefFrameT, schemeName string, request CefRequestT) CefResourceHandlerT
}

type CefSchemeHandlerFactory struct {
    CStruct         *C.struct__cef_scheme_handler_factory_t // memory manage?
    Factory         SchemeHandlerFactory
}

var (
    schemeHandlerMap = make(map[unsafe.Pointer]CefSchemeHandlerFactory)
)


//export go_CreateSchemeHandler
func go_CreateSchemeHandler(
        self *C.struct__cef_scheme_handler_factory_t,
        browser *C.struct__cef_browser_t,
        frame *C.struct__cef_frame_t,
        scheme_name *C.cef_string_utf8_t,
        request *C.struct__cef_request_t) *C.struct__cef_resource_handler_t {

    schemeName := C.GoString(scheme_name.str)
    defer C.cef_string_userfree_utf8_free(scheme_name)

    if handler, ok := schemeHandlerMap[unsafe.Pointer(self)]; ok {
        return handler.Factory.CreateSchemeHandler(
            CefBrowserT{browser},
            CefFrameT{frame},
            schemeName,
            CefRequestT{request},
        ).CStruct
    }
    return nil
}

func deleteCustomScheme(it unsafe.Pointer) {
    // delete it from our map
    delete(schemeHandlerMap, it)
}

func RegisterCustomScheme(schemeName, domainName string, schemeHandler SchemeHandlerFactory) (int, CefSchemeHandlerFactory) {
    var handler CefSchemeHandlerFactory
    handler.Factory = schemeHandler

    handler.CStruct = (*C.struct__cef_scheme_handler_factory_t)(
            C.calloc(1, C.sizeof_struct__cef_scheme_handler_factory_t))
    C.intialize_cef_scheme_handler_factory(handler.CStruct)
    RegisterDestructor(unsafe.Pointer(handler.CStruct), deleteCustomScheme)

    schemeNameCs := C.CString(schemeName)
    defer C.free(unsafe.Pointer(schemeNameCs))

    domainNameCs := C.CString(domainName)
    defer C.free(unsafe.Pointer(domainNameCs))

    retCode := C.cef_scheme_handler_register(schemeNameCs, domainNameCs, handler.CStruct)

    schemeHandlerMap[unsafe.Pointer(handler.CStruct)] = handler

    return int(retCode), handler
}

func AddCrossOriginWhitelistEntry(sourceOrigin, targetProtocol, targetDomain string, allowSubdomains bool) bool {

    sourceOriginCs := C.CString(sourceOrigin)
    defer C.free(unsafe.Pointer(sourceOriginCs))
    targetProtocolCs := C.CString(targetProtocol)
    defer C.free(unsafe.Pointer(targetProtocolCs))
    targetDomainCs := C.CString(targetDomain)
    defer C.free(unsafe.Pointer(targetDomainCs))

    allowDomainsInt := 1
    if !allowSubdomains {
        allowDomainsInt = 0
    }

    return C.whitelist_cef_add_cross_origin_whitelist_entry(sourceOriginCs, targetProtocolCs, targetDomainCs, C.int(allowDomainsInt)) == C.int(1)

}