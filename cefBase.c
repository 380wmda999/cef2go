// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi
// Website: https://github.com/fromkeith/cefcapi

#include "cefBase.h"
#include <string.h>

// ----------------------------------------------------------------------------
// cef_base_t
// ----------------------------------------------------------------------------

///
// Structure defining the reference count implementation functions. All
// framework structures must include the cef_base_t structure first.
///

///
// Increment the reference count.
///
int CEF_CALLBACK add_ref(cef_base_t* self) {
    DEBUG_CALLBACK("cef_base_t.add_ref\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("+");
    return 1;
}

///
// Decrement the reference count.  Delete this object when no references
// remain.
///
int CEF_CALLBACK release(cef_base_t* self) {
    DEBUG_CALLBACK("cef_base_t.release\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("-");
    return 1;
}

///
// Returns the current number of references.
///
int CEF_CALLBACK get_refct(cef_base_t* self) {
    DEBUG_CALLBACK("cef_base_t.get_refct\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("=");
    return 1;
}

void initialize_cef_base(cef_base_t* base) {
    printf("initialize_cef_base\n");
    // Check if "size" member was set.
    size_t size = base->size;
    // Let's print the size in case sizeof was used
    // on a pointer instead of a structure. In such
    // case the number will be very high.
    printf("cef_base_t.size = %lu\n", (unsigned long)size);
    if (size <= 0) {
        printf("FATAL: initialize_cef_base failed, size member not set\n");
        _exit(1);
    }
    base->add_ref = add_ref;
    base->release = release;
    base->get_refct = get_refct;
}

//
// other base/shared items
//


// returns a utf8 encoded string that you need to delete
cef_string_utf8_t * cefStringToUtf8(cef_string_t * source) {
    cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
    cef_string_to_utf8(source->str, source->length, output);
    return output;
}
