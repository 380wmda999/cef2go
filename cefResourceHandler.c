// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go


#include "_cgo_export.h"


#include "cefBase.h"
#include "include/capi/cef_resource_handler_capi.h"



int cef_resource_handler_t_process_request(struct _cef_resource_handler_t* self,
      struct _cef_request_t* request, struct _cef_callback_t* callback) {
    return go_ProcessRequest(self, request, callback);
}

void cef_resource_handler_t_get_response_headers(
  struct _cef_resource_handler_t* self, struct _cef_response_t* response,
  int64* response_length, cef_string_t* redirectUrl) {
    // not supporting redirect right now
    go_GetResponseHeaders(self, response, response_length);
}

int cef_resource_handler_t_read_response(struct _cef_resource_handler_t* self,
  void* data_out, int bytes_to_read, int* bytes_read,
  struct _cef_callback_t* callback) {
    return go_ReadResponse(self, data_out, bytes_to_read, bytes_read, callback);
}

int cef_resource_handler_t_can_get_cookie(struct _cef_resource_handler_t* self,
  const struct _cef_cookie_t* cookie) {
    return go_GetCookie(self, (struct _cef_cookie_t*) cookie);
}

int cef_resource_handler_t_can_set_cookie(struct _cef_resource_handler_t* self,
  const struct _cef_cookie_t* cookie) {
    return go_SetCookie(self, (struct _cef_cookie_t*) cookie);
}

void cef_resource_handler_t_cancel(struct _cef_resource_handler_t* self) {
    go_Cancel(self);
}

void intialize_cef_resource_handler(struct _cef_resource_handler_t* handler) {
    DEBUG_CALLBACK("initializeResourceHandler\n");
    handler->base.size = sizeof(cef_resource_handler_t);
    initialize_cef_base((cef_base_t*) handler);

    handler->process_request = cef_resource_handler_t_process_request;
    handler->get_response_headers = cef_resource_handler_t_get_response_headers;
    handler->read_response = cef_resource_handler_t_read_response;
    handler->can_get_cookie = cef_resource_handler_t_can_get_cookie;
    handler->can_set_cookie = cef_resource_handler_t_can_set_cookie;
    handler->cancel = cef_resource_handler_t_cancel;
}