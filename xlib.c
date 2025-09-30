#include "xlib.h"
#include <X11/Xlib.h>
int ErrorHandlerWrapper(Display* display, XErrorEvent* event) {
    int xErrorHandlerCallback(Display*, XErrorEvent*);
    return xErrorHandlerCallback(display, event);
}
