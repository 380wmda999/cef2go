// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package cef2go


/*
#include <stdlib.h>
#include "cefBase.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"

extern cef_string_utf8_t * cef_request_t_get_url(struct _cef_request_t * self);
extern int cef_request_t_is_read_only(struct _cef_request_t * self);
extern void cef_request_t_set_url(struct _cef_request_t * self, char * url);
extern cef_string_utf8_t * cef_request_t_get_method(struct _cef_request_t * self);
extern void cef_request_t_set_method(struct _cef_request_t * self, char * meth);
extern struct _cef_post_data_t * cef_request_t_get_post_data(struct _cef_request_t * self);
extern cef_string_multimap_t cef_request_t_get_header_map(struct _cef_request_t * self);
extern int cef_request_t_get_flags(struct _cef_request_t * self);
extern void cef_request_t_set_flags(struct _cef_request_t * self, int f);
extern cef_string_utf8_t * cef_request_t_get_first_party_for_cookies(struct _cef_request_t * self);
extern int cef_request_t_get_resource_type(struct _cef_request_t * self);
extern int cef_request_t_get_transition_type(struct _cef_request_t * self);
*/
import "C"

import (
    "unsafe"
    //"fmt"
)

type CefRequestT struct {
    CStruct *C.struct__cef_request_t
}

type CefPostDataT struct {
    CStruct *C.struct__cef_post_data_t
}


func (r CefRequestT) GetUrl() string {
    stringStruct := C.cef_request_t_get_url(r.CStruct)
    defer C.cef_string_userfree_utf8_free(stringStruct)
    str := C.GoString(stringStruct.str)
    return str
}

func (r CefRequestT) IsReadOnly() bool {
    result := C.cef_request_t_is_read_only(r.CStruct)
    return result == 1
}

func (r CefRequestT) SetUrl(newUrl string) {
    var newUrlCString *C.char = C.CString(newUrl)
    defer C.free(unsafe.Pointer(newUrlCString))
    C.cef_request_t_set_url(r.CStruct, newUrlCString)
}

func (r CefRequestT) GetMethod() string {
    stringStruct := C.cef_request_t_get_method(r.CStruct)
    defer C.cef_string_userfree_utf8_free(stringStruct)
    str := C.GoString(stringStruct.str)
    return str
}

func (r CefRequestT) SetMethod(method string) {
    var methodCString *C.char = C.CString(method)
    defer C.free(unsafe.Pointer(methodCString))
    C.cef_request_t_set_url(r.CStruct, methodCString)
}

func (r CefRequestT) GetPostData() CefPostDataT {
    postDataCStruct := C.cef_request_t_get_post_data(r.CStruct)
    var toRet CefPostDataT
    toRet.CStruct = postDataCStruct
    return toRet
}

func (r CefRequestT) SetPostData(data CefPostDataT) {
    // do nothing for now
}

func (r CefRequestT) GetHeaderMap() map[string][]string {
    cefMapPointer := C.cef_request_t_get_header_map(r.CStruct)
    if cefMapPointer == nil {
        return nil
    }
    numKeys := C.cef_string_multimap_size(cefMapPointer)
    goMap := make(map[string][]string)
    for i := 0; i < int(numKeys); i++ {
        var key *C.cef_string_t = C.cef_string_userfree_utf16_alloc()
        C.cef_string_multimap_key(cefMapPointer, C.int(i), key)
        charKeyUtf8 := C.cefStringToUtf8(key)
        goKey := C.GoString(charKeyUtf8.str)
        if _, ok := goMap[goKey]; ok {
            continue
        }
        numValsForKey := C.cef_string_multimap_find_count(cefMapPointer, key)

        if numValsForKey >= 0 {
            goVals := make([]string, numValsForKey)
            for k := 0; k < int(numValsForKey); k++ {
                var val *C.cef_string_t = C.cef_string_userfree_utf16_alloc()
                C.cef_string_multimap_enumerate(cefMapPointer,
                    key, C.int(k), val)
                charValUtf8 := C.cefStringToUtf8(val)
                goVals[k] = C.GoString(charValUtf8.str)
                C.cef_string_userfree_utf8_free(charValUtf8)
                C.cef_string_userfree_utf16_free(val)
            }
            goMap[goKey] = goVals
        }
        C.cef_string_userfree_utf8_free(charKeyUtf8)
        C.cef_string_userfree_utf16_free(key)
    }
    C.cef_string_multimap_free(cefMapPointer)
    return goMap
}

func (r CefRequestT) SetHeaderMap(headerMap map[string]string) {
    // do nothing for now
}

func (r CefRequestT) Set(url, method string, postData CefPostDataT, headerMap map[string]string) {
    // do nothing for now
}

func (r CefRequestT) GetFlags() int {
    return int(C.cef_request_t_get_flags(r.CStruct))
}

func (r CefRequestT) SetFlags(flags int) {
    C.cef_request_t_set_flags(r.CStruct, C.int(flags))
}

func (r CefRequestT) GetFirstPartyForCookies() string {
    stringStruct := C.cef_request_t_get_first_party_for_cookies(r.CStruct)
    defer C.cef_string_userfree_utf8_free(stringStruct)
    str := C.GoString(stringStruct.str)
    return str
}

func (r CefRequestT) SetFirstPartyForCookies(url string) {
    // do nothing cause i'm lazy right now
}

func (r CefRequestT) GetResourceType() int {
    return int(C.cef_request_t_get_resource_type(r.CStruct))
}

func (r CefRequestT) GetTransitionType() int {
    return int(C.cef_request_t_get_transition_type(r.CStruct))
}