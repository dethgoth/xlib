package xlib

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lX11
#include <stdlib.h>
#include <X11/Xlib.h>
#include "xlib.h"
extern int xErrorHandlerCallback(Display* display, XErrorEvent* event);
*/
import "C"
import (
	"unsafe"
)

const (
	NoEventMask              int64 = 0
	KeyPressMask             int64 = 1
	KeyReleaseMask           int64 = 2
	ButtonPressMask          int64 = 4
	ButtonReleaseMask        int64 = 8
	EnterWindowMask          int64 = 16
	LeaveWindowMask          int64 = 32
	PointerMotionMask        int64 = 64
	PointerMotionHintMask    int64 = 128
	Button1MotionMask        int64 = 256
	Button2MotionMask        int64 = 512
	Button3MotionMask        int64 = 1024
	Button4MotionMask        int64 = 2048
	Button5MotionMask        int64 = 4096
	ButtonMotionMask         int64 = 8192
	KeymapStateMask          int64 = 16384
	ExposureMask             int64 = 32768
	VisibilityChangeMask     int64 = 65536
	StructureNotifyMask      int64 = 131072
	ResizeRedirectMask       int64 = 262144
	SubstructureNotifyMask   int64 = 524288
	SubstructureRedirectMask int64 = 1048576
	FocusChangeMask          int64 = 2097152
	PropertyChangeMask       int64 = 4194304
	ColormapChangeMask       int64 = 8388608
	OwnerGrabButtonMask      int64 = 16777216
	LineSolid                int   = 0
	LineOnOffDash            int   = 1
	LineDoubleDash           int   = 2
    CapNotLast int = 0
    CapButt int = 1
    CapRound int = 2
    CapProjecting int = 3
    JoinMitter int = 0
    JoinRound int = 1
    JoinBevel int = 2
    FillSolid int = 0
    FillTiled int = 1
    FillStippled int = 2
    FillOpaqueStippled int = 3
    EvenOddRule int = 0
    WindingRule int = 1
)

type Display C.Display
type Screen C.Screen
type XErrorEvent struct {
	Type         int
	Display      *Display
	ResourceId   uint64
	Serial       uint64
	Error_code   uint8
	Request_code uint8
	Minor_code   uint8
}
type Cursor C.Cursor
type XEvent C.XEvent
type GC C.GC

var (
	XErrorCallback func(*Display, *XErrorEvent)
)

func strConcat(a []interface{}) string {
	str := ""
	for _, strPart := range a {
		switch s := strPart.(type) {
		case string:
			str += s
		}
	}
	return str
}

/*
* function for calling user's go callback function
 */
//export xErrorHandlerCallback
func xErrorHandlerCallback(displayC *C.Display, eventC *C.XErrorEvent) C.int {
	display := (*Display)(unsafe.Pointer(displayC))
	event := (*XErrorEvent)(unsafe.Pointer(eventC))
	XErrorCallback(display, event)
	return 0
}

func XSetErrorHandler(callback func(*Display, *XErrorEvent)) func(*Display, *XErrorEvent) {
	XErrorCallback = callback
	C.XSetErrorHandler((C.XErrorHandler)(unsafe.Pointer(C.ErrorHandlerWrapper)))
	return XErrorCallback
}

func XSelectInput(display *Display, window uint32, event_mask int64) int {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	event_maskC := (C.long)(event_mask)
	return (int)(C.XSelectInput(displayC, windowC, event_maskC))
}

func XSync(display *Display, discard bool) int {
	displayC := (*C.Display)(display)
	var discardC C.Bool
	if discard {
		discardC = (C.Bool)(0)
	} else {
		discardC = (C.Bool)(1)
	}
	return (int)(C.XSync(displayC, discardC))
}
func XOpenDisplay(displayNameParts ...interface{}) *Display {
	if len(displayNameParts) == 0 {
		display := C.XOpenDisplay(nil)
		return (*Display)(display)
	} else {
		displayNameComplete := strConcat(displayNameParts)
		if len(displayNameComplete) > 0 {
			displayNameCompleteC := C.CString(displayNameComplete)
			display := C.XOpenDisplay(displayNameCompleteC)
			C.free(unsafe.Pointer(displayNameCompleteC))
			return (*Display)(display)

		} else {
			display := C.XOpenDisplay(nil)
			return (*Display)(display)
		}
	}
}
func XDefaultRootWindow(display *Display) uint32 {
	displayC := (*C.Display)(display)
	return (uint32)(C.XDefaultRootWindow(displayC))
}

func XQueryTree(display *Display, window uint32, root_return *uint32, parent_return *uint32, children_return **uint32, nchildren_return *uint32) {

}

func WidthOfScreen(screen Screen) int {
	return (int)(screen.width)
}

func HeightOfScreen(screen Screen) int {
	return (int)(screen.height)
}

func RootWindow(display *Display, screen int) uint32 {
	screenC := (C.int)(screen)
	displayC := (C.Display)(display)
	return (uint32)(C.XRootWindow(displayC, screenC))
}

func DefaultScreen(display *Display) int {
	displayC := (*C.Display)(display)
	return int(C.XDefaultScreen)
}

func DefaultDepth(display *Display, screen int) int {
	displayC := (*C.Display)(display)
	screenC := (C.int)(screen)
	return (int)(C.XDefaultDepth(displayC, screenC))
}

func XCreateGC(display *Display, drawable uint32, valuemask uint32, values *XGCValues) GC {
	displayC := (C.Display)(display)
	drawableC := (C.ulong)(drawable)
	valuemaskC := (C.ulong)(valuemask)
	valuesC := (*C.XGCValues)(values)
	return (GC)(C.XCreateGC(displayC, drawableC, valuemaskC, valuesC))
}

func XSetLineAttributes(display *Display, gc GC, line_width uint, line_style int, cap_style int, join_style int) int {
	displayC := (*C.Display)(display)
	gcC := (C.GC)(gc)
	line_widthC := (c.uint)(line_width)
	line_styleC := (c.int)(line_style)
	lcap_styleC := (c.int)(cap_style)
	join_styleC := (c.int)(join_style)
	return (int)(C.XSetLineAttributes(displayC, gcC, line_widthC, line_styleC, cap_styleC, join_styleC))
}
