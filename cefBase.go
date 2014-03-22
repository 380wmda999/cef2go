// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/fromkeith/cef2go

package cef2go


/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
import (
    "unsafe"
    "sync"
    //"runtime/debug"
)
var _MainArgs *C.struct__cef_main_args_t


const (
    LOGSEVERITY_DEFAULT = C.LOGSEVERITY_DEFAULT
    LOGSEVERITY_VERBOSE = C.LOGSEVERITY_VERBOSE
    LOGSEVERITY_INFO = C.LOGSEVERITY_INFO
    LOGSEVERITY_WARNING = C.LOGSEVERITY_WARNING
    LOGSEVERITY_ERROR = C.LOGSEVERITY_ERROR
    LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
    LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)

// a niave memory management.
// allows us to keep track of allocated resources, and their counts
// furthermore, the deconstructor lets us free any go resources once
// their C versions have been released
// General info about the ref count in CEF:
//      http://code.google.com/p/chromiumembedded/wiki/UsingTheCAPI
type MemoryManagedBridge struct {
    Count           int
    Deconstructor   func(it unsafe.Pointer)
}

var (
    memoryBridge = make(map[unsafe.Pointer]MemoryManagedBridge)
    refCountLock sync.Mutex
)

//export go_AddRef
func go_AddRef(it unsafe.Pointer) int {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        //Logger.Println("Known Ref_Add: ", it)
        m.Count++
        memoryBridge[it] = m
        return m.Count
    }
    Logger.Println("Unknown Ref_Add: ", it)
    return 1
}

//export go_Release
func go_Release(it unsafe.Pointer) int {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        m.Count--
        if m.Count == 0 {
            if m.Deconstructor != nil {
                m.Deconstructor(it)
            }
            Logger.Println("Known Ref_Free: ", it)
            C.free(it)
            delete(memoryBridge, it)
        } else {
            //Logger.Println("Known Ref_Release: ", it)
            memoryBridge[it] = m
        }
        return m.Count
    }
    Logger.Println("Unknown Ref_Release: ", it)
    return 1
}
//export go_GetRefCount
func go_GetRefCount(it unsafe.Pointer) int {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        return m.Count
    }
    Logger.Println("Unknown Ref_Count: ", it)
    return 1
}

//export go_CreateRef
func go_CreateRef(it unsafe.Pointer) {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if _, ok := memoryBridge[it]; !ok {
        Logger.Println("Ref_Create: ", it)
        var m MemoryManagedBridge
        m.Deconstructor = nil
        memoryBridge[it] = m
        return
    }
    Logger.Println("Ref Already exists Ref_Create: ", it)
}

func RegisterDestructor(it unsafe.Pointer, decon func(it unsafe.Pointer)) bool {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        m.Deconstructor = decon
        memoryBridge[it] = m
        return true
    }
    return false
}

func _InitializeGlobalCStructuresBase() {
     _MainArgs = (*C.struct__cef_main_args_t)(
            C.calloc(1, C.sizeof_struct__cef_main_args_t))
     Logger.Println("_MainArgs: ", unsafe.Pointer(_MainArgs))
}
