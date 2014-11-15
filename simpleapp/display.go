package main


import (
    "github.com/fromkeith/cef2go"
    "log"
)



type myDisplayHandler struct {
    handler         cef2go.DisplayHandlerT
}


func (d *myDisplayHandler) OnAddressChange(browser cef2go.CefBrowserT, frame cef2go.CefFrameT, url string) {
    defer browser.Release()
    defer frame.Release()
    log.Println("OnAddressChange" + url)
}
func (d *myDisplayHandler) OnTitleChange(browser cef2go.CefBrowserT, title string) {
    defer browser.Release()
    log.Println("OnTitleChange" + title)
}
func (d *myDisplayHandler) OnToolTip(browser cef2go.CefBrowserT, text string) bool {
    defer browser.Release()
    return false
}
func (d *myDisplayHandler) OnStatusMessage(browser cef2go.CefBrowserT, value string) {
    defer browser.Release()
}
func (d *myDisplayHandler) OnConsoleMessage(browser cef2go.CefBrowserT, message, source string, line int) bool {
    defer browser.Release()
    log.Println("[Console] ", message)
    return true
}

func (d *myDisplayHandler) GetDisplayHandlerT() cef2go.DisplayHandlerT {
    return d.handler
}