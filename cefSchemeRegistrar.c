// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

#include "_cgo_export.h"

#include "cefBase.h"
#include "include/capi/cef_scheme_capi.h"
#include "include/capi/cef_origin_whitelist_capi.h"




struct _cef_resource_handler_t* cef_scheme_handler_factory_t_create(
      struct _cef_scheme_handler_factory_t* self,
      struct _cef_browser_t* browser, struct _cef_frame_t* frame,
      const cef_string_t* scheme_name, struct _cef_request_t* request) {
    cef_string_utf8_t * out = cefStringToUtf8(scheme_name);
    return go_CreateSchemeHandler(self, browser, frame, out, request);
}

void intialize_cef_scheme_handler_factory(struct _cef_scheme_handler_factory_t * factory) {
    goDebugLog("initializeSchemeHandler\n");
    factory->base.size = sizeof(cef_scheme_handler_factory_t);
    initialize_cef_base((cef_base_t*) factory);
    go_AddRef((cef_base_t*) factory);
    factory->create = cef_scheme_handler_factory_t_create;
}


int cef_scheme_handler_register(char * schemeName, char * domainName, struct _cef_scheme_handler_factory_t* factory) {
    cef_string_t * schemeNameCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(schemeName, strlen(schemeName), schemeNameCef);
    cef_string_t * domainNameCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(domainName, strlen(domainName), domainNameCef);

    int ret =
        cef_register_scheme_handler_factory(schemeNameCef, domainNameCef, factory);

    cef_string_userfree_utf16_free(domainNameCef);
    cef_string_userfree_utf16_free(schemeNameCef);
    return ret;
}


int whitelist_cef_add_cross_origin_whitelist_entry(
    char* sourceOrigin,
    char* targetProtocol,
    char* targetDomain,
    int allow_target_subdomains) {

    cef_string_t * sourceOriginCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(sourceOrigin, strlen(sourceOrigin), sourceOriginCef);
    cef_string_t * targetProtocolCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(targetProtocol, strlen(targetProtocol), targetProtocolCef);
    cef_string_t * targetDomainCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(targetDomain, strlen(targetDomain), targetDomainCef);


    int ret = cef_add_cross_origin_whitelist_entry(sourceOriginCef, targetProtocolCef, targetDomainCef, allow_target_subdomains);

    cef_string_userfree_utf16_free(targetDomainCef);
    cef_string_userfree_utf16_free(targetProtocolCef);
    cef_string_userfree_utf16_free(sourceOriginCef);

    return ret;
}