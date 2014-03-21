// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package cef2go




/*
#include "cefBase.h"
#include "include/capi/cef_response_capi.h"

extern int cef_response_t_is_read_only(struct _cef_response_t* self);
extern int cef_response_t_get_status(struct _cef_response_t* self);
extern void cef_response_t_set_status(struct _cef_response_t* self, int status);
extern cef_string_utf8_t * cef_response_t_get_status_text(struct _cef_response_t* self);
extern void cef_response_t_set_status_text(struct _cef_response_t* self, char* statusText);
extern cef_string_utf8_t * cef_response_t_get_mime_type( struct _cef_response_t* self);
extern void cef_response_t_set_mime_type(struct _cef_response_t* self, char* mimeType);
extern cef_string_utf8_t * cef_response_t_get_header(struct _cef_response_t* self, char* name);
extern void cef_response_t_get_header_map(struct _cef_response_t* self, cef_string_multimap_t headerMap);
extern void cef_response_t_set_header_map(struct _cef_response_t* self, cef_string_multimap_t headerMap);
*/
import "C"

import "unsafe"

type CefResponseT struct {
    CStruct         *C.struct__cef_response_t
}


func (r CefResponseT) IsReadOnly() bool {
    return C.cef_response_t_is_read_only(r.CStruct) == C.int(1)
}
func (r CefResponseT) GetStatus() int {
    return int(C.cef_response_t_get_status(r.CStruct))
}
func (r CefResponseT) SetStatus(status int) {
    C.cef_response_t_set_status(r.CStruct, C.int(status))
}
func (r CefResponseT) GetStatusText() string {
    cefString := C.cef_response_t_get_status_text(r.CStruct)
    defer C.cef_string_userfree_utf8_free(cefString)
    return C.GoString(cefString.str)
}
func (r CefResponseT) SetStatusText(text string) {
    textCs := C.CString(text)
    defer C.free(unsafe.Pointer(textCs))
    C.cef_response_t_set_status_text(r.CStruct, textCs)
}
func (r CefResponseT) GetMimeType() string {
    cefString :=  C.cef_response_t_get_mime_type(r.CStruct)
    defer C.cef_string_userfree_utf8_free(cefString)
    return C.GoString(cefString.str)
}
func (r CefResponseT) SetMimeType(mimeType string) {
    mimeTypeCs := C.CString(mimeType)
    defer C.free(unsafe.Pointer(mimeTypeCs))
    C.cef_response_t_set_mime_type(r.CStruct, mimeTypeCs)
}
func (r CefResponseT) GetHeader(name string) string {
    nameCs := C.CString(name)
    defer C.free(unsafe.Pointer(nameCs))
    valCef := C.cef_response_t_get_header(r.CStruct, nameCs)
    defer C.cef_string_userfree_utf8_free(valCef)
    return C.GoString(valCef.str)
}
func (r CefResponseT) GetHeaderMap() map[string][]string {
    headerMapPtr := C.cef_string_multimap_alloc()
    defer C.cef_string_multimap_free(headerMapPtr)
    C.cef_response_t_get_header_map(r.CStruct, headerMapPtr)
    return extractCefMultiMap(headerMapPtr)
}
func (r CefResponseT) SetHeaderMap(map[string][]string) {
    headerMapPtr := C.cef_string_multimap_alloc()
    defer C.cef_string_multimap_free(headerMapPtr)
    C.cef_response_t_set_header_map(r.CStruct, headerMapPtr)
}