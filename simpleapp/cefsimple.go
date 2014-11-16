package main


import(
    "github.com/fromkeith/cef2go"
    "github.com/fromkeith/w32"
    "unsafe"
    "log"
)



func main() {
    log.Println("Start")
    hInstance := w32.GetModuleHandle("")
    log.Printf("hInstances: %x\n", hInstance)

    var app myAppHandler
    app.handler = cef2go.NewAppHandlerT(&app)
    var browserPHandler myBrowserProcessHandler
    browserPHandler.handler = cef2go.NewBrowserProcessHandlerT(&browserPHandler)
    app.browserProcessHandler = &browserPHandler

    ret := cef2go.ExecuteProcess(unsafe.Pointer(hInstance), &app)
    log.Printf("ExecuteResult: %d\n", ret)
    if ret >= 0 {
        return
    }

    settings := cef2go.Settings{}
    log.Printf("Initialize: %#v\n", settings)
    cef2go.Initialize(settings, &app)

    log.Println("Run the message loop")
    cef2go.RunMessageLoop()

    log.Println("Shutdown")
    cef2go.Shutdown()

    log.Println("Cleanup")

    cef2go.DumpRefs()
}