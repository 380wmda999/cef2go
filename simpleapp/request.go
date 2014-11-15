package main


import (
    "github.com/fromkeith/cef2go"
    "log"
)

type myRequestHandler struct {
    handler         cef2go.RequestHandlerT
}



func (reqH *myRequestHandler) OnBeforeBrowse(browser cef2go.CefBrowserT, frame cef2go.CefFrameT, request cef2go.CefRequestT, isRedirect int) int {
    defer browser.Release()
    defer frame.Release()
    defer request.Release()
    log.Println("Before browse: ", request.GetUrl())
    return 0
}
func (reqH *myRequestHandler) OnBeforeResourceLoad(browser cef2go.CefBrowserT, frame cef2go.CefFrameT, request cef2go.CefRequestT) int {
    defer browser.Release()
    defer frame.Release()
    defer request.Release()
    return 0
}
func (reqH *myRequestHandler) OnCertificateError(errorCode cef2go.CefErrorCode, requestUrl string, errorCallback cef2go.CefCertErrorCallbackT) int {
    errorCallback.Release()
    return 0
}

func (reqH *myRequestHandler) GetRequestHandlerT() cef2go.RequestHandlerT {
    return reqH.handler
}