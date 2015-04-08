// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/fromkeith/cef2go

package cef2go


/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "cefBase.h"

extern int releaseVoid(void * self);
extern void appendToMultiMap(cef_string_multimap_t into, char *key, char *val);
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
    Name            string
}

var (
    memoryBridge = make(map[unsafe.Pointer]MemoryManagedBridge)
    refCountLock sync.Mutex
)




//export go_AddRef
func go_AddRef(it unsafe.Pointer) {
    if it == nil {
        return
    }
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        //Logger.Println("Known Ref_Add: ", it)
        m.Count++
        memoryBridge[it] = m
        return
    }
    Logger.Warnf("Unknown Ref_Add: %x\n", it)
}

//export go_Release
func go_Release(it unsafe.Pointer) int {
    if it == nil {
        return 0
    }
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        m.Count--
        if m.Count == 0 {
            if m.Deconstructor != nil {
                m.Deconstructor(it)
            }
            C.free(it)
            delete(memoryBridge, it)
            return 1
        } else {
            //Logger.Println("Known Ref_Release: ", it)
            memoryBridge[it] = m
        }
        return 0
    }
    Logger.Warnf("Unknown Ref_Release: %x\n", it)
    return 0
}
//export go_HasOneReferenceCount
func go_HasOneReferenceCount(it unsafe.Pointer) int {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if m, ok := memoryBridge[it]; ok {
        if m.Count == 1 {
            return 1
        }
        return 0
    }
    Logger.Warnf("Unknown Ref_Count: %x\n", it)
    return 0
}

func Release(w unsafe.Pointer) int {
    return int(C.releaseVoid(w))
}

func AddRef(it unsafe.Pointer) {
    C.add_refVoid(it)
}



//export go_CreateRef
func go_CreateRef(it unsafe.Pointer, name *C.char) {
    goname := ""
    if name != nil {
        goname = C.GoString(name)
    }
    CreateRef(it, goname)
}

func CreateRef(it unsafe.Pointer, name string) {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    if _, ok := memoryBridge[it]; !ok {
        var m MemoryManagedBridge
        m.Deconstructor = nil
        m.Name = name
        memoryBridge[it] = m
        return
    }
    Logger.Warnf("Ref Already exists Ref_Create: %x\n", it)
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
     CreateRef(unsafe.Pointer(_MainArgs), "MainArgs")
     Logger.Infof("_MainArgs: %x", unsafe.Pointer(_MainArgs))
}


func DumpRefs() {
    refCountLock.Lock()
    defer refCountLock.Unlock()

    for k, v := range memoryBridge {
        Logger.Infof("%X : %#v", k, v)
    }
}


func copyIntoMultiMap(into C.cef_string_multimap_t, from map[string][]string) {
    for k, vals := range from {
        copyKeyIntoMultiMap(into, k, vals)
    }
}
func copyKeyIntoMultiMap(into C.cef_string_multimap_t, key string, vals []string) {
    keyCStr := C.CString(key)
    defer C.free(unsafe.Pointer(keyCStr))
    for _, v := range vals {
        copyKeyValIntoMultiMap(into, keyCStr, v)
    }
}
func copyKeyValIntoMultiMap(into C.cef_string_multimap_t, keyCStr *C.char, v string) {
    valCStr := C.CString(v)
    defer C.free(unsafe.Pointer(valCStr))
    C.appendToMultiMap(into, keyCStr, valCStr)
}