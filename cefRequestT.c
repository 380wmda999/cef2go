
#include <stdlib.h>
#include "cefBase.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"



cef_string_utf8_t* cef_request_t_get_url(struct _cef_request_t * self) {
    cef_string_t * str = self->get_url(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}


int cef_request_t_is_read_only(struct _cef_request_t * self) {
    return self->is_read_only(self);
}

void cef_request_t_set_url(struct _cef_request_t * self, char * url) {
    cef_string_t output;
    output.str = NULL;
    output.length = 0;
    output.dtor = NULL;

    cef_string_from_utf8(url, strlen(url), &output);

    self->set_url(self, &output);
}

cef_string_utf8_t* cef_request_t_get_method(struct _cef_request_t * self) {
    cef_string_t * str = self->get_method(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

void cef_request_t_set_method(struct _cef_request_t * self, char * meth){
    cef_string_t output;
    output.str = NULL;
    output.length = 0;
    output.dtor = NULL;

    cef_string_from_utf8(meth, strlen(meth), &output);

    self->set_url(self, &output);
}

struct _cef_post_data_t * cef_request_t_get_post_data(struct _cef_request_t * self) {
    return self->get_post_data(self);
}

cef_string_multimap_t cef_request_t_get_header_map(struct _cef_request_t * self) {
    cef_string_multimap_t mapPointer = cef_string_multimap_alloc();
    self->get_header_map(self, mapPointer);
    return mapPointer;
}

int cef_request_t_get_flags(struct _cef_request_t * self) {
    return self->get_flags(self);
}

void cef_request_t_set_flags(struct _cef_request_t * self, int f) {
    self->set_flags(self, f);
}

cef_string_utf8_t* cef_request_t_get_first_party_for_cookies(struct _cef_request_t * self) {
    cef_string_t * str = self->get_first_party_for_cookies(self);
    cef_string_utf8_t* out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

int cef_request_t_get_resource_type(struct _cef_request_t * self) {
    return self->get_resource_type(self);
}

int cef_request_t_get_transition_type(struct _cef_request_t * self) {
    return self->get_transition_type(self);
}