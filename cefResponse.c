// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

#include "_cgo_export.h"

#include "cefBase.h"
#include "include/capi/cef_response_capi.h"





int cef_response_t_is_read_only(struct _cef_response_t* self) {
    return self->is_read_only(self);
}


int cef_response_t_get_status(struct _cef_response_t* self) {
    return self->get_status(self);
}


void cef_response_t_set_status(struct _cef_response_t* self, int status) {
    self->set_status(self, status);
}

cef_string_utf8_t * cef_response_t_get_status_text(struct _cef_response_t* self) {

    cef_string_t * str = self->get_status_text(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

void cef_response_t_set_status_text(struct _cef_response_t* self, char* statusText) {
    cef_string_t * statusTextCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(statusText, strlen(statusText), statusTextCef);

    self->set_status_text(self, statusTextCef);

    cef_string_userfree_utf16_free(statusTextCef);
}

cef_string_utf8_t * cef_response_t_get_mime_type(struct _cef_response_t* self) {

    cef_string_t * str = self->get_mime_type(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

void cef_response_t_set_mime_type(struct _cef_response_t* self, char* mimeType) {
    cef_string_t * mimeTypeCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(mimeType, strlen(mimeType), mimeTypeCef);

    self->set_mime_type(self, mimeTypeCef);

    cef_string_userfree_utf16_free(mimeTypeCef);
}

cef_string_utf8_t * cef_response_t_get_header(struct _cef_response_t* self, char* name) {

    cef_string_t * nameCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(name, strlen(name), nameCef);

    cef_string_t * str = self->get_header(self, nameCef);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    cef_string_userfree_utf16_free(nameCef);
    return out;
}

void cef_response_t_get_header_map(struct _cef_response_t* self, cef_string_multimap_t headerMap) {
    self->get_header_map(self, headerMap);
}

void cef_response_t_set_header_map(struct _cef_response_t* self, cef_string_multimap_t headerMap) {
    self->set_header_map(self, headerMap);
}