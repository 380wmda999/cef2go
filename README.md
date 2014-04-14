CEF2go - HTML 5 based GUI toolkit for the Go language
=====================================================

CEF2go is an open source project founded by [Czarek Tomczak]
(http://www.linkedin.com/in/czarektomczak) in 2014
to provide Go bindings for the [Chromium Embedded Framework]
(https://code.google.com/p/chromiumembedded/) (CEF).
CEF2go can act as a GUI toolkit, allowing you to create an HTML 5
based GUI in your application. Or you can provide browser
capabilities to your application.

Supported platforms: Windows, Linux, Mac OSX.

Currently the CEF2go example creates just a simple window with
the Chromium browser embedded. You can set a few options for
your application like the cache directory. More advanced bindings
are in plans, and that includes javascript bindings and callbacks, so
that you can have bidirectional communication between Go and
Javascript in a native way. Though, it is already possible to
communicate with Go from Javascript, see the next section for
an example.

CEF2go is licensed under the BSD 3-clause license, see the LICENSE
file.

This Fork
---------
This fork just restructures the source so its easier to 'go get' and get started.
 * It removes copies of 3rd party libraries.
 * This fork has also only been tested on Windows.
 * Wraps some callbacks from CEF.
 * Wraps some structures for easier access from Go.
 * Changes from using log.Logger interface to a more verbose one.
 * Implements a naive version of the reference counting.



