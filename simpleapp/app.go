package main


import (
    "github.com/fromkeith/cef2go"
    "log"
    "github.com/fromkeith/w32"
    "syscall"
    "unsafe"
)

type myAppHandler struct {
    handler         cef2go.AppHandlerT
    browserProcessHandler  cef2go.BrowserProcessHandler
}



func (app *myAppHandler) OnBeforeCommandLineProcessing(processType string, commandLine cef2go.CommandLineT) {

}
func (app *myAppHandler) OnRegisterCustomSchemes() {

}
func (app *myAppHandler) GetResourceBundleHandler() {

}
func (app *myAppHandler) GetBrowserProcessHandler() cef2go.BrowserProcessHandler {
    cef2go.AddRef(unsafe.Pointer(app.browserProcessHandler.GetBrowserProcessHandlerT().CStruct))
    return app.browserProcessHandler
}
func (app *myAppHandler) GetRenderProcessHandler() {

}
func (app *myAppHandler) GetAppHandlerT() cef2go.AppHandlerT {
    return app.handler
}

type myBrowserProcessHandler struct {
    handler             cef2go.BrowserProcessHandlerT
}

func (bph * myBrowserProcessHandler) OnContextInitialized() {
    log.Println("OnContextInitialized")
    settings := cef2go.BrowserSettings{}

    var rootClientHandler myClientHandler
    rootClientHandler.clientHandlerT = cef2go.NewClientHandlerT(&rootClientHandler)
    rootClientHandler.lifeSpan = new(myLifeSpanHandler)
    rootClientHandler.lifeSpan.lifeSpan = cef2go.NewLifeSpanHandlerT(rootClientHandler.lifeSpan)
    rootClientHandler.requestHandler = new(myRequestHandler)
    rootClientHandler.requestHandler.handler = cef2go.NewRequestHandlerT(rootClientHandler.requestHandler)
    rootClientHandler.displayHandler = new(myDisplayHandler)
    rootClientHandler.displayHandler.handler = cef2go.NewDisplayHandlerT(rootClientHandler.displayHandler)


    root := createRootWindow()
    cef2go.CreateBrowser(unsafe.Pointer(root), &rootClientHandler, settings, "http://www.google.com")
}

func (bph * myBrowserProcessHandler) OnBeforeChildProcessLaunch(commandLine cef2go.CommandLineT) {
    defer commandLine.Release()
}


func (bph * myBrowserProcessHandler) OnRenderProcessThreadCreated(extraInfo cef2go.CefListValueT) {
    defer extraInfo.Release()
    log.Printf("OnRenderProcessThreadCreated: %X", unsafe.Pointer(extraInfo.CStruct))
}


func (bph * myBrowserProcessHandler) GetBrowserProcessHandlerT() cef2go.BrowserProcessHandlerT {
    return bph.handler
}






/// heere be windows window creating code.

func createRootWindow() w32.HWND {
    windowClassName, _ := syscall.UTF16PtrFromString("myWindowClass")
    moduleHandle := w32.GetModuleHandle("")

    var wc w32.WNDCLASSEX
    wc.Size = uint32(unsafe.Sizeof(wc))
    wc.WndProc = syscall.NewCallback(WndProc)
    wc.Instance = moduleHandle
    wc.MenuName = nil
    wc.ClassName = windowClassName
    wc.Style = w32.CS_HREDRAW | w32.CS_VREDRAW
    if regClass := w32.RegisterClassEx(&wc); regClass == 0 {
        panic("Failed to register window class")
    }


    windowName, _ := syscall.UTF16PtrFromString("SimpleApp")

    // CreateWindowEx
    wh := w32.CreateWindowEx(
        0,
        windowClassName,
        windowName,
        w32.WS_OVERLAPPEDWINDOW,
        w32.CW_USEDEFAULT,
        w32.CW_USEDEFAULT,
        w32.CW_USEDEFAULT,
        w32.CW_USEDEFAULT,
        0,
        0,
        moduleHandle,
        nil,
    )
    if wh == 0 {
        panic("Parent.CreateWindowEX returned 0")
    }

    // ShowWindow
    w32.ShowWindow(wh, w32.SW_MAXIMIZE)

    // UpdateWindow
    if success := w32.UpdateWindow(wh); !success {
        panic("Parent.UpdateWindow returned 0")
    }
    return wh
}


func WndProc(hwndRaw syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
    hwnd := w32.HWND(hwndRaw)
    /*switch msg {
        case w32.WM_DESTROY:
            cef2go.QuitMessageLoop()
            break
    }*/
    rc = w32.DefWindowProc(hwnd, msg, wparam, lparam)
    return
}