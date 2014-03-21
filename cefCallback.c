// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go


#include "cefBase.h"
#include "include/capi/cef_callback_capi.h"


void cef_callback_t_cont(struct _cef_callback_t* self) {
    self->cont(self);
}


void cef_callback_t_cancel(struct _cef_callback_t* self) {
    self->cancel(self);
}