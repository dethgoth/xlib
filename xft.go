package xlib

//#cgo pkg-config: xft
//#include <stdlib.h>
//#include <X11/Xft/Xft.h>
import "C"

type XRenderColor C.XRenderColor

type XftColor struct {
    pixel uint32
    color XRenderColor
}
