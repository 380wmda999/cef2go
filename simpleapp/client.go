package main


import (
    "github.com/fromkeith/cef2go"
    "log"
)

type myClientHandler struct {
    clientHandlerT      cef2go.ClientHandlerT
    lifeSpan            *myLifeSpanHandler
    requestHandler      *myRequestHandler
    displayHandler      *myDisplayHandler
}




func (ch *myClientHandler) GetContextMenuHandler() cef2go.ContextMenuHandlerT {
    return cef2go.ContextMenuHandlerT{nil}
}
func (ch *myClientHandler) GetDialogHandler() cef2go.DialogHandlerT {
    return cef2go.DialogHandlerT{nil}
}
func (ch *myClientHandler) GetDisplayHandler() cef2go.DisplayHandler {
    return ch.displayHandler
}
func (ch *myClientHandler) GetDownloadHandler() cef2go.DownloadHandler {
    return nil
}
func (ch *myClientHandler) GetDragHandler() cef2go.DragHandlerT {
    return cef2go.DragHandlerT{nil}
}
func (ch *myClientHandler) GetFocusHandler() cef2go.FocusHandlerT {
    return cef2go.FocusHandlerT{nil}
}
func (ch *myClientHandler) GetGeoLocationHandler() cef2go.GeolocationHandlerT {
    return cef2go.GeolocationHandlerT{nil}
}
func (ch *myClientHandler) GetJsDialogHandler() cef2go.JsdialogHandlerT {
    return cef2go.JsdialogHandlerT{nil}
}
func (ch *myClientHandler) GetKeyboardHandler() cef2go.KeyboardHandlerT {
    return cef2go.KeyboardHandlerT{nil}
}
func (ch *myClientHandler) GetLifeSpanHandler() cef2go.LifeSpanHandler {
    return ch.lifeSpan
}
func (ch *myClientHandler) GetLoadHandler() cef2go.LoadHandlerT {
    return cef2go.LoadHandlerT{nil}
}
func (ch *myClientHandler) GetRenderHandler() cef2go.RenderHandlerT {
    return cef2go.RenderHandlerT{nil}
}
func (ch *myClientHandler) GetRequestHandler() cef2go.RequestHandler {
    return ch.requestHandler
}

func (ch *myClientHandler) GetClientHandlerT() cef2go.ClientHandlerT {
    return ch.clientHandlerT
}


type myLifeSpanHandler struct {
    lifeSpan            cef2go.LifeSpanHandlerT
}


func (l *myLifeSpanHandler) OnAfterCreated(browser cef2go.CefBrowserT) {
    defer browser.Release()
    log.Println("lifespan::OnAfterCreated")
}
func (l *myLifeSpanHandler) RunModal(browser cef2go.CefBrowserT) int {
    defer browser.Release()
    log.Println("lifespan::RunModal")
    return 0
}
func (l *myLifeSpanHandler) DoClose(browser cef2go.CefBrowserT) int {
    defer browser.Release()
    log.Println("lifespan::DoClose")
    return 0
}
func (l *myLifeSpanHandler) BeforeClose(browser cef2go.CefBrowserT) {
    defer browser.Release()
    log.Println("lifespan::BeforeClose")
    cef2go.QuitMessageLoop()
}
func (l *myLifeSpanHandler) GetLifeSpanHandlerT() cef2go.LifeSpanHandlerT {
    return l.lifeSpan
}