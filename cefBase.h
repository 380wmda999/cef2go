#ifndef CEF_BASE_H
#define CEF_BASE_H

#include <stdio.h>
#include "include/capi/cef_base_capi.h"


// Set to 1 to check if add_ref() and release()
// are called and to track the total number of calls.
// add_ref will be printed as "+", release as "-".
#define DEBUG_REFERENCE_COUNTING 0

// Print only the first execution of the callback,
// ignore the subsequent.
#define DEBUG_CALLBACK(x) { static int first_call = 1; if (first_call == 1) { first_call = 0; printf(x); } }


int CEF_CALLBACK add_ref(cef_base_t* self);
int CEF_CALLBACK release(cef_base_t* self);
int CEF_CALLBACK get_refct(cef_base_t* self);

#endif
