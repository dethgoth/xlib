package xlib

//#cgo pkg-config: pango pangoxft
//#include <pango/pango.h>
//#include <pango/pangoxft.h>
import "C"

const (
	PANGO_SCALE_XX_SMALL float64 = 0.5787037037037
	PANGO_SCALE_X_SMALL  float64 = 0.6944444444444
	PANGO_SCALE_SMALL    float64 = 0.8333333333333
	PANGO_SCALE_MEDIUM   float64 = 1.0
	PANGO_SCALE_LARGE    float64 = 1.2
	PANGO_SCALE_X_LARGE  float64 = 1.44
	PANGO_SCALE_XX_LARGE float64 = 1.728
	PANGO_SCALE          int     = 1024
)

type PangoLayout C.PangoLayout
type PangoFontMap C.PangoFontMap
type PangoContext C.PangoContext
type PangoFontDescription C.PangoFontDescription
type PangoFontMetrics C.PangoFontMetrics
type PangoLanguage C.PangoLanguage

func Pango_xft_get_font_map(display *Display, screen int) *PangoFontMap {
	displayC := (*C.Display)(display)
	screenC := (*C.int)(screen)
	return (*PangoFontMap)(C.Pango_xft_get_font_map(displayC, screenC))
}

func pango_font_map_create_context(fontmap *PangoFontMap) *PangoContext {
	fontmapC := (*C.PangoFontMap)(fontmap)
	return (*PangoContext)(C.pango_font_map_create_context(fontmapC))
}

func pango_font_description_from_string(fontname string) *PangoFontDescription {
	fontnameC := C.CString(fontname)
	result := (*PangoFontDescription)(C.pango_font_description_from_string(fontnameC))
	C.Free(fontnameC)
	return result
}

func Pango_layout_new(context *PangoContext) *PangoLayout {
	contextC := (*C.PangoContext)(context)
	return (*PangoLayout)(C.Pango_layout_new(contextC))
}

func Pango_layout_set_font_description(layout *PangoLayout, description *PangoFontDescription) {
	layoutC := (*C.PangoLayout)(layout)
	descriptionC := (*C.PangoFontDescription)(description)
	C.Pango_layout_set_font_description(layoutC, descriptionC)
}

func Pango_context_get_metrics(context *PangoContext, description *PangoFontDescription, language *PangoLanguage) *PangoFontMetrics {
	contextC := (*C.PangoContext)(context)
	descriptionC := (*C.PangoFontDescription)(description)
	languageC := (*C.PangoLanguage)
	return (*PangoFontMetrics)(C.pango_context_get_metrics(contextC, descriptionC, languageC))
}

func Pango_font_metrics_unref(metrics) {

}
