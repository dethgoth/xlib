#ifndef XLIB_WRAPPER_H
#define XLIB_WRAPPER_H
#include <X11/Xlib.h>
extern int xErrorHandlerCallback(Display* display, XErrorEvent* event);
int ErrorHandlerWrapper(Display* display, XErrorEvent* event);

#endif
