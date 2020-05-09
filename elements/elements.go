package elements

import (
  "github.com/ratorx/htmlgo"
)

// A represents the HTML element 'a'.
// For more information visit https://www.w3schools.com/tags/tag_a.asp.
func A(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "a", Attributes: attrs, Children: children}
}

// A_ is a convenience wrapper for A without the attrs argument.
func A_(children ...HTML) HTML {
  return A(nil, children...)
}

// Abbr represents the HTML element 'abbr'.
// For more information visit https://www.w3schools.com/tags/tag_abbr.asp.
func Abbr(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "abbr", Attributes: attrs, Children: children}
}

// Abbr_ is a convenience wrapper for Abbr without the attrs argument.
func Abbr_(children ...HTML) HTML {
  return Abbr(nil, children...)
}

// Acronym represents the HTML element 'acronym'.
// For more information visit https://www.w3schools.com/tags/tag_acronym.asp.
func Acronym(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "acronym", Attributes: attrs, Children: children}
}

// Acronym_ is a convenience wrapper for Acronym without the attrs argument.
func Acronym_(children ...HTML) HTML {
  return Acronym(nil, children...)
}

// Address represents the HTML element 'address'.
// For more information visit https://www.w3schools.com/tags/tag_address.asp.
func Address(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "address", Attributes: attrs, Children: children}
}

// Address_ is a convenience wrapper for Address without the attrs argument.
func Address_(children ...HTML) HTML {
  return Address(nil, children...)
}

// Applet represents the HTML element 'applet'.
// For more information visit https://www.w3schools.com/tags/tag_applet.asp.
func Applet(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "applet", Attributes: attrs, Children: children}
}

// Applet_ is a convenience wrapper for Applet without the attrs argument.
func Applet_(children ...HTML) HTML {
  return Applet(nil, children...)
}

// Article represents the HTML element 'article'.
// For more information visit https://www.w3schools.com/tags/tag_article.asp.
func Article(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "article", Attributes: attrs, Children: children}
}

// Article_ is a convenience wrapper for Article without the attrs argument.
func Article_(children ...HTML) HTML {
  return Article(nil, children...)
}

// Aside represents the HTML element 'aside'.
// For more information visit https://www.w3schools.com/tags/tag_aside.asp.
func Aside(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "aside", Attributes: attrs, Children: children}
}

// Aside_ is a convenience wrapper for Aside without the attrs argument.
func Aside_(children ...HTML) HTML {
  return Aside(nil, children...)
}

// Audio represents the HTML element 'audio'.
// For more information visit https://www.w3schools.com/tags/tag_audio.asp.
func Audio(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "audio", Attributes: attrs, Children: children}
}

// Audio_ is a convenience wrapper for Audio without the attrs argument.
func Audio_(children ...HTML) HTML {
  return Audio(nil, children...)
}

// B represents the HTML element 'b'.
// For more information visit https://www.w3schools.com/tags/tag_b.asp.
func B(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "b", Attributes: attrs, Children: children}
}

// B_ is a convenience wrapper for B without the attrs argument.
func B_(children ...HTML) HTML {
  return B(nil, children...)
}

// Basefont represents the HTML element 'basefont'.
// For more information visit https://www.w3schools.com/tags/tag_basefont.asp.
func Basefont(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "basefont", Attributes: attrs, Children: children}
}

// Basefont_ is a convenience wrapper for Basefont without the attrs argument.
func Basefont_(children ...HTML) HTML {
  return Basefont(nil, children...)
}

// Bdi represents the HTML element 'bdi'.
// For more information visit https://www.w3schools.com/tags/tag_bdi.asp.
func Bdi(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "bdi", Attributes: attrs, Children: children}
}

// Bdi_ is a convenience wrapper for Bdi without the attrs argument.
func Bdi_(children ...HTML) HTML {
  return Bdi(nil, children...)
}

// Bdo represents the HTML element 'bdo'.
// For more information visit https://www.w3schools.com/tags/tag_bdo.asp.
func Bdo(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "bdo", Attributes: attrs, Children: children}
}

// Bdo_ is a convenience wrapper for Bdo without the attrs argument.
func Bdo_(children ...HTML) HTML {
  return Bdo(nil, children...)
}

// Bgsound represents the HTML element 'bgsound'.
// For more information visit https://www.w3schools.com/tags/tag_bgsound.asp.
func Bgsound(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "bgsound", Attributes: attrs, Children: children}
}

// Bgsound_ is a convenience wrapper for Bgsound without the attrs argument.
func Bgsound_(children ...HTML) HTML {
  return Bgsound(nil, children...)
}

// Big represents the HTML element 'big'.
// For more information visit https://www.w3schools.com/tags/tag_big.asp.
func Big(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "big", Attributes: attrs, Children: children}
}

// Big_ is a convenience wrapper for Big without the attrs argument.
func Big_(children ...HTML) HTML {
  return Big(nil, children...)
}

// Blink represents the HTML element 'blink'.
// For more information visit https://www.w3schools.com/tags/tag_blink.asp.
func Blink(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "blink", Attributes: attrs, Children: children}
}

// Blink_ is a convenience wrapper for Blink without the attrs argument.
func Blink_(children ...HTML) HTML {
  return Blink(nil, children...)
}

// Blockquote represents the HTML element 'blockquote'.
// For more information visit https://www.w3schools.com/tags/tag_blockquote.asp.
func Blockquote(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "blockquote", Attributes: attrs, Children: children}
}

// Blockquote_ is a convenience wrapper for Blockquote without the attrs argument.
func Blockquote_(children ...HTML) HTML {
  return Blockquote(nil, children...)
}

// Body represents the HTML element 'body'.
// For more information visit https://www.w3schools.com/tags/tag_body.asp.
func Body(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "body", Attributes: attrs, Children: children}
}

// Body_ is a convenience wrapper for Body without the attrs argument.
func Body_(children ...HTML) HTML {
  return Body(nil, children...)
}

// Button represents the HTML element 'button'.
// For more information visit https://www.w3schools.com/tags/tag_button.asp.
func Button(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "button", Attributes: attrs, Children: children}
}

// Button_ is a convenience wrapper for Button without the attrs argument.
func Button_(children ...HTML) HTML {
  return Button(nil, children...)
}

// Canvas represents the HTML element 'canvas'.
// For more information visit https://www.w3schools.com/tags/tag_canvas.asp.
func Canvas(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "canvas", Attributes: attrs, Children: children}
}

// Canvas_ is a convenience wrapper for Canvas without the attrs argument.
func Canvas_(children ...HTML) HTML {
  return Canvas(nil, children...)
}

// Caption represents the HTML element 'caption'.
// For more information visit https://www.w3schools.com/tags/tag_caption.asp.
func Caption(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "caption", Attributes: attrs, Children: children}
}

// Caption_ is a convenience wrapper for Caption without the attrs argument.
func Caption_(children ...HTML) HTML {
  return Caption(nil, children...)
}

// Center represents the HTML element 'center'.
// For more information visit https://www.w3schools.com/tags/tag_center.asp.
func Center(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "center", Attributes: attrs, Children: children}
}

// Center_ is a convenience wrapper for Center without the attrs argument.
func Center_(children ...HTML) HTML {
  return Center(nil, children...)
}

// Cite represents the HTML element 'cite'.
// For more information visit https://www.w3schools.com/tags/tag_cite.asp.
func Cite(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "cite", Attributes: attrs, Children: children}
}

// Cite_ is a convenience wrapper for Cite without the attrs argument.
func Cite_(children ...HTML) HTML {
  return Cite(nil, children...)
}

// Code represents the HTML element 'code'.
// For more information visit https://www.w3schools.com/tags/tag_code.asp.
func Code(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "code", Attributes: attrs, Children: children}
}

// Code_ is a convenience wrapper for Code without the attrs argument.
func Code_(children ...HTML) HTML {
  return Code(nil, children...)
}

// Colgroup represents the HTML element 'colgroup'.
// For more information visit https://www.w3schools.com/tags/tag_colgroup.asp.
func Colgroup(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "colgroup", Attributes: attrs, Children: children}
}

// Colgroup_ is a convenience wrapper for Colgroup without the attrs argument.
func Colgroup_(children ...HTML) HTML {
  return Colgroup(nil, children...)
}

// Datalist represents the HTML element 'datalist'.
// For more information visit https://www.w3schools.com/tags/tag_datalist.asp.
func Datalist(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "datalist", Attributes: attrs, Children: children}
}

// Datalist_ is a convenience wrapper for Datalist without the attrs argument.
func Datalist_(children ...HTML) HTML {
  return Datalist(nil, children...)
}

// Dd represents the HTML element 'dd'.
// For more information visit https://www.w3schools.com/tags/tag_dd.asp.
func Dd(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "dd", Attributes: attrs, Children: children}
}

// Dd_ is a convenience wrapper for Dd without the attrs argument.
func Dd_(children ...HTML) HTML {
  return Dd(nil, children...)
}

// Del represents the HTML element 'del'.
// For more information visit https://www.w3schools.com/tags/tag_del.asp.
func Del(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "del", Attributes: attrs, Children: children}
}

// Del_ is a convenience wrapper for Del without the attrs argument.
func Del_(children ...HTML) HTML {
  return Del(nil, children...)
}

// Details represents the HTML element 'details'.
// For more information visit https://www.w3schools.com/tags/tag_details.asp.
func Details(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "details", Attributes: attrs, Children: children}
}

// Details_ is a convenience wrapper for Details without the attrs argument.
func Details_(children ...HTML) HTML {
  return Details(nil, children...)
}

// Dfn represents the HTML element 'dfn'.
// For more information visit https://www.w3schools.com/tags/tag_dfn.asp.
func Dfn(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "dfn", Attributes: attrs, Children: children}
}

// Dfn_ is a convenience wrapper for Dfn without the attrs argument.
func Dfn_(children ...HTML) HTML {
  return Dfn(nil, children...)
}

// Dir represents the HTML element 'dir'.
// For more information visit https://www.w3schools.com/tags/tag_dir.asp.
func Dir(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "dir", Attributes: attrs, Children: children}
}

// Dir_ is a convenience wrapper for Dir without the attrs argument.
func Dir_(children ...HTML) HTML {
  return Dir(nil, children...)
}

// Div represents the HTML element 'div'.
// For more information visit https://www.w3schools.com/tags/tag_div.asp.
func Div(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "div", Attributes: attrs, Children: children}
}

// Div_ is a convenience wrapper for Div without the attrs argument.
func Div_(children ...HTML) HTML {
  return Div(nil, children...)
}

// Dl represents the HTML element 'dl'.
// For more information visit https://www.w3schools.com/tags/tag_dl.asp.
func Dl(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "dl", Attributes: attrs, Children: children}
}

// Dl_ is a convenience wrapper for Dl without the attrs argument.
func Dl_(children ...HTML) HTML {
  return Dl(nil, children...)
}

// Dt represents the HTML element 'dt'.
// For more information visit https://www.w3schools.com/tags/tag_dt.asp.
func Dt(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "dt", Attributes: attrs, Children: children}
}

// Dt_ is a convenience wrapper for Dt without the attrs argument.
func Dt_(children ...HTML) HTML {
  return Dt(nil, children...)
}

// Em represents the HTML element 'em'.
// For more information visit https://www.w3schools.com/tags/tag_em.asp.
func Em(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "em", Attributes: attrs, Children: children}
}

// Em_ is a convenience wrapper for Em without the attrs argument.
func Em_(children ...HTML) HTML {
  return Em(nil, children...)
}

// Fieldset represents the HTML element 'fieldset'.
// For more information visit https://www.w3schools.com/tags/tag_fieldset.asp.
func Fieldset(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "fieldset", Attributes: attrs, Children: children}
}

// Fieldset_ is a convenience wrapper for Fieldset without the attrs argument.
func Fieldset_(children ...HTML) HTML {
  return Fieldset(nil, children...)
}

// Figcaption represents the HTML element 'figcaption'.
// For more information visit https://www.w3schools.com/tags/tag_figcaption.asp.
func Figcaption(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "figcaption", Attributes: attrs, Children: children}
}

// Figcaption_ is a convenience wrapper for Figcaption without the attrs argument.
func Figcaption_(children ...HTML) HTML {
  return Figcaption(nil, children...)
}

// Figure represents the HTML element 'figure'.
// For more information visit https://www.w3schools.com/tags/tag_figure.asp.
func Figure(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "figure", Attributes: attrs, Children: children}
}

// Figure_ is a convenience wrapper for Figure without the attrs argument.
func Figure_(children ...HTML) HTML {
  return Figure(nil, children...)
}

// Font represents the HTML element 'font'.
// For more information visit https://www.w3schools.com/tags/tag_font.asp.
func Font(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "font", Attributes: attrs, Children: children}
}

// Font_ is a convenience wrapper for Font without the attrs argument.
func Font_(children ...HTML) HTML {
  return Font(nil, children...)
}

// Footer represents the HTML element 'footer'.
// For more information visit https://www.w3schools.com/tags/tag_footer.asp.
func Footer(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "footer", Attributes: attrs, Children: children}
}

// Footer_ is a convenience wrapper for Footer without the attrs argument.
func Footer_(children ...HTML) HTML {
  return Footer(nil, children...)
}

// Form represents the HTML element 'form'.
// For more information visit https://www.w3schools.com/tags/tag_form.asp.
func Form(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "form", Attributes: attrs, Children: children}
}

// Form_ is a convenience wrapper for Form without the attrs argument.
func Form_(children ...HTML) HTML {
  return Form(nil, children...)
}

// Frame represents the HTML element 'frame'.
// For more information visit https://www.w3schools.com/tags/tag_frame.asp.
func Frame(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "frame", Attributes: attrs, Children: children}
}

// Frame_ is a convenience wrapper for Frame without the attrs argument.
func Frame_(children ...HTML) HTML {
  return Frame(nil, children...)
}

// Frameset represents the HTML element 'frameset'.
// For more information visit https://www.w3schools.com/tags/tag_frameset.asp.
func Frameset(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "frameset", Attributes: attrs, Children: children}
}

// Frameset_ is a convenience wrapper for Frameset without the attrs argument.
func Frameset_(children ...HTML) HTML {
  return Frameset(nil, children...)
}

// H1 represents the HTML element 'h1'.
// For more information visit https://www.w3schools.com/tags/tag_h1.asp.
func H1(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "h1", Attributes: attrs, Children: children}
}

// H1_ is a convenience wrapper for H1 without the attrs argument.
func H1_(children ...HTML) HTML {
  return H1(nil, children...)
}

// H2 represents the HTML element 'h2'.
// For more information visit https://www.w3schools.com/tags/tag_h2.asp.
func H2(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "h2", Attributes: attrs, Children: children}
}

// H2_ is a convenience wrapper for H2 without the attrs argument.
func H2_(children ...HTML) HTML {
  return H2(nil, children...)
}

// H3 represents the HTML element 'h3'.
// For more information visit https://www.w3schools.com/tags/tag_h3.asp.
func H3(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "h3", Attributes: attrs, Children: children}
}

// H3_ is a convenience wrapper for H3 without the attrs argument.
func H3_(children ...HTML) HTML {
  return H3(nil, children...)
}

// H4 represents the HTML element 'h4'.
// For more information visit https://www.w3schools.com/tags/tag_h4.asp.
func H4(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "h4", Attributes: attrs, Children: children}
}

// H4_ is a convenience wrapper for H4 without the attrs argument.
func H4_(children ...HTML) HTML {
  return H4(nil, children...)
}

// H5 represents the HTML element 'h5'.
// For more information visit https://www.w3schools.com/tags/tag_h5.asp.
func H5(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "h5", Attributes: attrs, Children: children}
}

// H5_ is a convenience wrapper for H5 without the attrs argument.
func H5_(children ...HTML) HTML {
  return H5(nil, children...)
}

// H6 represents the HTML element 'h6'.
// For more information visit https://www.w3schools.com/tags/tag_h6.asp.
func H6(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "h6", Attributes: attrs, Children: children}
}

// H6_ is a convenience wrapper for H6 without the attrs argument.
func H6_(children ...HTML) HTML {
  return H6(nil, children...)
}

// Head represents the HTML element 'head'.
// For more information visit https://www.w3schools.com/tags/tag_head.asp.
func Head(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "head", Attributes: attrs, Children: children}
}

// Head_ is a convenience wrapper for Head without the attrs argument.
func Head_(children ...HTML) HTML {
  return Head(nil, children...)
}

// Header represents the HTML element 'header'.
// For more information visit https://www.w3schools.com/tags/tag_header.asp.
func Header(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "header", Attributes: attrs, Children: children}
}

// Header_ is a convenience wrapper for Header without the attrs argument.
func Header_(children ...HTML) HTML {
  return Header(nil, children...)
}

// Hgroup represents the HTML element 'hgroup'.
// For more information visit https://www.w3schools.com/tags/tag_hgroup.asp.
func Hgroup(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "hgroup", Attributes: attrs, Children: children}
}

// Hgroup_ is a convenience wrapper for Hgroup without the attrs argument.
func Hgroup_(children ...HTML) HTML {
  return Hgroup(nil, children...)
}

// I represents the HTML element 'i'.
// For more information visit https://www.w3schools.com/tags/tag_i.asp.
func I(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "i", Attributes: attrs, Children: children}
}

// I_ is a convenience wrapper for I without the attrs argument.
func I_(children ...HTML) HTML {
  return I(nil, children...)
}

// Iframe represents the HTML element 'iframe'.
// For more information visit https://www.w3schools.com/tags/tag_iframe.asp.
func Iframe(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "iframe", Attributes: attrs, Children: children}
}

// Iframe_ is a convenience wrapper for Iframe without the attrs argument.
func Iframe_(children ...HTML) HTML {
  return Iframe(nil, children...)
}

// Ins represents the HTML element 'ins'.
// For more information visit https://www.w3schools.com/tags/tag_ins.asp.
func Ins(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "ins", Attributes: attrs, Children: children}
}

// Ins_ is a convenience wrapper for Ins without the attrs argument.
func Ins_(children ...HTML) HTML {
  return Ins(nil, children...)
}

// Isindex represents the HTML element 'isindex'.
// For more information visit https://www.w3schools.com/tags/tag_isindex.asp.
func Isindex(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "isindex", Attributes: attrs, Children: children}
}

// Isindex_ is a convenience wrapper for Isindex without the attrs argument.
func Isindex_(children ...HTML) HTML {
  return Isindex(nil, children...)
}

// Kbd represents the HTML element 'kbd'.
// For more information visit https://www.w3schools.com/tags/tag_kbd.asp.
func Kbd(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "kbd", Attributes: attrs, Children: children}
}

// Kbd_ is a convenience wrapper for Kbd without the attrs argument.
func Kbd_(children ...HTML) HTML {
  return Kbd(nil, children...)
}

// Keygen represents the HTML element 'keygen'.
// For more information visit https://www.w3schools.com/tags/tag_keygen.asp.
func Keygen(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "keygen", Attributes: attrs, Children: children}
}

// Keygen_ is a convenience wrapper for Keygen without the attrs argument.
func Keygen_(children ...HTML) HTML {
  return Keygen(nil, children...)
}

// Label represents the HTML element 'label'.
// For more information visit https://www.w3schools.com/tags/tag_label.asp.
func Label(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "label", Attributes: attrs, Children: children}
}

// Label_ is a convenience wrapper for Label without the attrs argument.
func Label_(children ...HTML) HTML {
  return Label(nil, children...)
}

// Legend represents the HTML element 'legend'.
// For more information visit https://www.w3schools.com/tags/tag_legend.asp.
func Legend(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "legend", Attributes: attrs, Children: children}
}

// Legend_ is a convenience wrapper for Legend without the attrs argument.
func Legend_(children ...HTML) HTML {
  return Legend(nil, children...)
}

// Li represents the HTML element 'li'.
// For more information visit https://www.w3schools.com/tags/tag_li.asp.
func Li(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "li", Attributes: attrs, Children: children}
}

// Li_ is a convenience wrapper for Li without the attrs argument.
func Li_(children ...HTML) HTML {
  return Li(nil, children...)
}

// Listing represents the HTML element 'listing'.
// For more information visit https://www.w3schools.com/tags/tag_listing.asp.
func Listing(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "listing", Attributes: attrs, Children: children}
}

// Listing_ is a convenience wrapper for Listing without the attrs argument.
func Listing_(children ...HTML) HTML {
  return Listing(nil, children...)
}

// Main represents the HTML element 'main'.
// For more information visit https://www.w3schools.com/tags/tag_main.asp.
func Main(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "main", Attributes: attrs, Children: children}
}

// Main_ is a convenience wrapper for Main without the attrs argument.
func Main_(children ...HTML) HTML {
  return Main(nil, children...)
}

// Map represents the HTML element 'map'.
// For more information visit https://www.w3schools.com/tags/tag_map.asp.
func Map(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "map", Attributes: attrs, Children: children}
}

// Map_ is a convenience wrapper for Map without the attrs argument.
func Map_(children ...HTML) HTML {
  return Map(nil, children...)
}

// Mark represents the HTML element 'mark'.
// For more information visit https://www.w3schools.com/tags/tag_mark.asp.
func Mark(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "mark", Attributes: attrs, Children: children}
}

// Mark_ is a convenience wrapper for Mark without the attrs argument.
func Mark_(children ...HTML) HTML {
  return Mark(nil, children...)
}

// Marquee represents the HTML element 'marquee'.
// For more information visit https://www.w3schools.com/tags/tag_marquee.asp.
func Marquee(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "marquee", Attributes: attrs, Children: children}
}

// Marquee_ is a convenience wrapper for Marquee without the attrs argument.
func Marquee_(children ...HTML) HTML {
  return Marquee(nil, children...)
}

// Menu represents the HTML element 'menu'.
// For more information visit https://www.w3schools.com/tags/tag_menu.asp.
func Menu(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "menu", Attributes: attrs, Children: children}
}

// Menu_ is a convenience wrapper for Menu without the attrs argument.
func Menu_(children ...HTML) HTML {
  return Menu(nil, children...)
}

// Meter represents the HTML element 'meter'.
// For more information visit https://www.w3schools.com/tags/tag_meter.asp.
func Meter(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "meter", Attributes: attrs, Children: children}
}

// Meter_ is a convenience wrapper for Meter without the attrs argument.
func Meter_(children ...HTML) HTML {
  return Meter(nil, children...)
}

// Nav represents the HTML element 'nav'.
// For more information visit https://www.w3schools.com/tags/tag_nav.asp.
func Nav(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "nav", Attributes: attrs, Children: children}
}

// Nav_ is a convenience wrapper for Nav without the attrs argument.
func Nav_(children ...HTML) HTML {
  return Nav(nil, children...)
}

// Nobr represents the HTML element 'nobr'.
// For more information visit https://www.w3schools.com/tags/tag_nobr.asp.
func Nobr(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "nobr", Attributes: attrs, Children: children}
}

// Nobr_ is a convenience wrapper for Nobr without the attrs argument.
func Nobr_(children ...HTML) HTML {
  return Nobr(nil, children...)
}

// Noframes represents the HTML element 'noframes'.
// For more information visit https://www.w3schools.com/tags/tag_noframes.asp.
func Noframes(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "noframes", Attributes: attrs, Children: children}
}

// Noframes_ is a convenience wrapper for Noframes without the attrs argument.
func Noframes_(children ...HTML) HTML {
  return Noframes(nil, children...)
}

// Noscript represents the HTML element 'noscript'.
// For more information visit https://www.w3schools.com/tags/tag_noscript.asp.
func Noscript(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "noscript", Attributes: attrs, Children: children}
}

// Noscript_ is a convenience wrapper for Noscript without the attrs argument.
func Noscript_(children ...HTML) HTML {
  return Noscript(nil, children...)
}

// Object represents the HTML element 'object'.
// For more information visit https://www.w3schools.com/tags/tag_object.asp.
func Object(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "object", Attributes: attrs, Children: children}
}

// Object_ is a convenience wrapper for Object without the attrs argument.
func Object_(children ...HTML) HTML {
  return Object(nil, children...)
}

// Ol represents the HTML element 'ol'.
// For more information visit https://www.w3schools.com/tags/tag_ol.asp.
func Ol(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "ol", Attributes: attrs, Children: children}
}

// Ol_ is a convenience wrapper for Ol without the attrs argument.
func Ol_(children ...HTML) HTML {
  return Ol(nil, children...)
}

// Optgroup represents the HTML element 'optgroup'.
// For more information visit https://www.w3schools.com/tags/tag_optgroup.asp.
func Optgroup(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "optgroup", Attributes: attrs, Children: children}
}

// Optgroup_ is a convenience wrapper for Optgroup without the attrs argument.
func Optgroup_(children ...HTML) HTML {
  return Optgroup(nil, children...)
}

// Option represents the HTML element 'option'.
// For more information visit https://www.w3schools.com/tags/tag_option.asp.
func Option(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "option", Attributes: attrs, Children: children}
}

// Option_ is a convenience wrapper for Option without the attrs argument.
func Option_(children ...HTML) HTML {
  return Option(nil, children...)
}

// Output represents the HTML element 'output'.
// For more information visit https://www.w3schools.com/tags/tag_output.asp.
func Output(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "output", Attributes: attrs, Children: children}
}

// Output_ is a convenience wrapper for Output without the attrs argument.
func Output_(children ...HTML) HTML {
  return Output(nil, children...)
}

// P represents the HTML element 'p'.
// For more information visit https://www.w3schools.com/tags/tag_p.asp.
func P(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "p", Attributes: attrs, Children: children}
}

// P_ is a convenience wrapper for P without the attrs argument.
func P_(children ...HTML) HTML {
  return P(nil, children...)
}

// Plaintext represents the HTML element 'plaintext'.
// For more information visit https://www.w3schools.com/tags/tag_plaintext.asp.
func Plaintext(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "plaintext", Attributes: attrs, Children: children}
}

// Plaintext_ is a convenience wrapper for Plaintext without the attrs argument.
func Plaintext_(children ...HTML) HTML {
  return Plaintext(nil, children...)
}

// Pre represents the HTML element 'pre'.
// For more information visit https://www.w3schools.com/tags/tag_pre.asp.
func Pre(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "pre", Attributes: attrs, Children: children}
}

// Pre_ is a convenience wrapper for Pre without the attrs argument.
func Pre_(children ...HTML) HTML {
  return Pre(nil, children...)
}

// Progress represents the HTML element 'progress'.
// For more information visit https://www.w3schools.com/tags/tag_progress.asp.
func Progress(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "progress", Attributes: attrs, Children: children}
}

// Progress_ is a convenience wrapper for Progress without the attrs argument.
func Progress_(children ...HTML) HTML {
  return Progress(nil, children...)
}

// Q represents the HTML element 'q'.
// For more information visit https://www.w3schools.com/tags/tag_q.asp.
func Q(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "q", Attributes: attrs, Children: children}
}

// Q_ is a convenience wrapper for Q without the attrs argument.
func Q_(children ...HTML) HTML {
  return Q(nil, children...)
}

// Rp represents the HTML element 'rp'.
// For more information visit https://www.w3schools.com/tags/tag_rp.asp.
func Rp(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "rp", Attributes: attrs, Children: children}
}

// Rp_ is a convenience wrapper for Rp without the attrs argument.
func Rp_(children ...HTML) HTML {
  return Rp(nil, children...)
}

// Rt represents the HTML element 'rt'.
// For more information visit https://www.w3schools.com/tags/tag_rt.asp.
func Rt(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "rt", Attributes: attrs, Children: children}
}

// Rt_ is a convenience wrapper for Rt without the attrs argument.
func Rt_(children ...HTML) HTML {
  return Rt(nil, children...)
}

// Ruby represents the HTML element 'ruby'.
// For more information visit https://www.w3schools.com/tags/tag_ruby.asp.
func Ruby(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "ruby", Attributes: attrs, Children: children}
}

// Ruby_ is a convenience wrapper for Ruby without the attrs argument.
func Ruby_(children ...HTML) HTML {
  return Ruby(nil, children...)
}

// S represents the HTML element 's'.
// For more information visit https://www.w3schools.com/tags/tag_s.asp.
func S(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "s", Attributes: attrs, Children: children}
}

// S_ is a convenience wrapper for S without the attrs argument.
func S_(children ...HTML) HTML {
  return S(nil, children...)
}

// Samp represents the HTML element 'samp'.
// For more information visit https://www.w3schools.com/tags/tag_samp.asp.
func Samp(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "samp", Attributes: attrs, Children: children}
}

// Samp_ is a convenience wrapper for Samp without the attrs argument.
func Samp_(children ...HTML) HTML {
  return Samp(nil, children...)
}

// Script represents the HTML element 'script'.
// For more information visit https://www.w3schools.com/tags/tag_script.asp.
func Script(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "script", Attributes: attrs, Children: children}
}

// Script_ is a convenience wrapper for Script without the attrs argument.
func Script_(children ...HTML) HTML {
  return Script(nil, children...)
}

// Section represents the HTML element 'section'.
// For more information visit https://www.w3schools.com/tags/tag_section.asp.
func Section(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "section", Attributes: attrs, Children: children}
}

// Section_ is a convenience wrapper for Section without the attrs argument.
func Section_(children ...HTML) HTML {
  return Section(nil, children...)
}

// Select represents the HTML element 'select'.
// For more information visit https://www.w3schools.com/tags/tag_select.asp.
func Select(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "select", Attributes: attrs, Children: children}
}

// Select_ is a convenience wrapper for Select without the attrs argument.
func Select_(children ...HTML) HTML {
  return Select(nil, children...)
}

// Small represents the HTML element 'small'.
// For more information visit https://www.w3schools.com/tags/tag_small.asp.
func Small(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "small", Attributes: attrs, Children: children}
}

// Small_ is a convenience wrapper for Small without the attrs argument.
func Small_(children ...HTML) HTML {
  return Small(nil, children...)
}

// Spacer represents the HTML element 'spacer'.
// For more information visit https://www.w3schools.com/tags/tag_spacer.asp.
func Spacer(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "spacer", Attributes: attrs, Children: children}
}

// Spacer_ is a convenience wrapper for Spacer without the attrs argument.
func Spacer_(children ...HTML) HTML {
  return Spacer(nil, children...)
}

// Span represents the HTML element 'span'.
// For more information visit https://www.w3schools.com/tags/tag_span.asp.
func Span(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "span", Attributes: attrs, Children: children}
}

// Span_ is a convenience wrapper for Span without the attrs argument.
func Span_(children ...HTML) HTML {
  return Span(nil, children...)
}

// Strike represents the HTML element 'strike'.
// For more information visit https://www.w3schools.com/tags/tag_strike.asp.
func Strike(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "strike", Attributes: attrs, Children: children}
}

// Strike_ is a convenience wrapper for Strike without the attrs argument.
func Strike_(children ...HTML) HTML {
  return Strike(nil, children...)
}

// Strong represents the HTML element 'strong'.
// For more information visit https://www.w3schools.com/tags/tag_strong.asp.
func Strong(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "strong", Attributes: attrs, Children: children}
}

// Strong_ is a convenience wrapper for Strong without the attrs argument.
func Strong_(children ...HTML) HTML {
  return Strong(nil, children...)
}

// Style represents the HTML element 'style'.
// For more information visit https://www.w3schools.com/tags/tag_style.asp.
func Style(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "style", Attributes: attrs, Children: children}
}

// Style_ is a convenience wrapper for Style without the attrs argument.
func Style_(children ...HTML) HTML {
  return Style(nil, children...)
}

// Sub represents the HTML element 'sub'.
// For more information visit https://www.w3schools.com/tags/tag_sub.asp.
func Sub(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "sub", Attributes: attrs, Children: children}
}

// Sub_ is a convenience wrapper for Sub without the attrs argument.
func Sub_(children ...HTML) HTML {
  return Sub(nil, children...)
}

// Summary represents the HTML element 'summary'.
// For more information visit https://www.w3schools.com/tags/tag_summary.asp.
func Summary(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "summary", Attributes: attrs, Children: children}
}

// Summary_ is a convenience wrapper for Summary without the attrs argument.
func Summary_(children ...HTML) HTML {
  return Summary(nil, children...)
}

// Sup represents the HTML element 'sup'.
// For more information visit https://www.w3schools.com/tags/tag_sup.asp.
func Sup(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "sup", Attributes: attrs, Children: children}
}

// Sup_ is a convenience wrapper for Sup without the attrs argument.
func Sup_(children ...HTML) HTML {
  return Sup(nil, children...)
}

// Table represents the HTML element 'table'.
// For more information visit https://www.w3schools.com/tags/tag_table.asp.
func Table(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "table", Attributes: attrs, Children: children}
}

// Table_ is a convenience wrapper for Table without the attrs argument.
func Table_(children ...HTML) HTML {
  return Table(nil, children...)
}

// Tbody represents the HTML element 'tbody'.
// For more information visit https://www.w3schools.com/tags/tag_tbody.asp.
func Tbody(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "tbody", Attributes: attrs, Children: children}
}

// Tbody_ is a convenience wrapper for Tbody without the attrs argument.
func Tbody_(children ...HTML) HTML {
  return Tbody(nil, children...)
}

// Td represents the HTML element 'td'.
// For more information visit https://www.w3schools.com/tags/tag_td.asp.
func Td(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "td", Attributes: attrs, Children: children}
}

// Td_ is a convenience wrapper for Td without the attrs argument.
func Td_(children ...HTML) HTML {
  return Td(nil, children...)
}

// Textarea represents the HTML element 'textarea'.
// For more information visit https://www.w3schools.com/tags/tag_textarea.asp.
func Textarea(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "textarea", Attributes: attrs, Children: children}
}

// Textarea_ is a convenience wrapper for Textarea without the attrs argument.
func Textarea_(children ...HTML) HTML {
  return Textarea(nil, children...)
}

// Tfoot represents the HTML element 'tfoot'.
// For more information visit https://www.w3schools.com/tags/tag_tfoot.asp.
func Tfoot(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "tfoot", Attributes: attrs, Children: children}
}

// Tfoot_ is a convenience wrapper for Tfoot without the attrs argument.
func Tfoot_(children ...HTML) HTML {
  return Tfoot(nil, children...)
}

// Th represents the HTML element 'th'.
// For more information visit https://www.w3schools.com/tags/tag_th.asp.
func Th(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "th", Attributes: attrs, Children: children}
}

// Th_ is a convenience wrapper for Th without the attrs argument.
func Th_(children ...HTML) HTML {
  return Th(nil, children...)
}

// Thead represents the HTML element 'thead'.
// For more information visit https://www.w3schools.com/tags/tag_thead.asp.
func Thead(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "thead", Attributes: attrs, Children: children}
}

// Thead_ is a convenience wrapper for Thead without the attrs argument.
func Thead_(children ...HTML) HTML {
  return Thead(nil, children...)
}

// Time represents the HTML element 'time'.
// For more information visit https://www.w3schools.com/tags/tag_time.asp.
func Time(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "time", Attributes: attrs, Children: children}
}

// Time_ is a convenience wrapper for Time without the attrs argument.
func Time_(children ...HTML) HTML {
  return Time(nil, children...)
}

// Title represents the HTML element 'title'.
// For more information visit https://www.w3schools.com/tags/tag_title.asp.
func Title(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "title", Attributes: attrs, Children: children}
}

// Title_ is a convenience wrapper for Title without the attrs argument.
func Title_(children ...HTML) HTML {
  return Title(nil, children...)
}

// Tr represents the HTML element 'tr'.
// For more information visit https://www.w3schools.com/tags/tag_tr.asp.
func Tr(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "tr", Attributes: attrs, Children: children}
}

// Tr_ is a convenience wrapper for Tr without the attrs argument.
func Tr_(children ...HTML) HTML {
  return Tr(nil, children...)
}

// Tt represents the HTML element 'tt'.
// For more information visit https://www.w3schools.com/tags/tag_tt.asp.
func Tt(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "tt", Attributes: attrs, Children: children}
}

// Tt_ is a convenience wrapper for Tt without the attrs argument.
func Tt_(children ...HTML) HTML {
  return Tt(nil, children...)
}

// U represents the HTML element 'u'.
// For more information visit https://www.w3schools.com/tags/tag_u.asp.
func U(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "u", Attributes: attrs, Children: children}
}

// U_ is a convenience wrapper for U without the attrs argument.
func U_(children ...HTML) HTML {
  return U(nil, children...)
}

// Ul represents the HTML element 'ul'.
// For more information visit https://www.w3schools.com/tags/tag_ul.asp.
func Ul(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "ul", Attributes: attrs, Children: children}
}

// Ul_ is a convenience wrapper for Ul without the attrs argument.
func Ul_(children ...HTML) HTML {
  return Ul(nil, children...)
}

// Var represents the HTML element 'var'.
// For more information visit https://www.w3schools.com/tags/tag_var.asp.
func Var(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "var", Attributes: attrs, Children: children}
}

// Var_ is a convenience wrapper for Var without the attrs argument.
func Var_(children ...HTML) HTML {
  return Var(nil, children...)
}

// Video represents the HTML element 'video'.
// For more information visit https://www.w3schools.com/tags/tag_video.asp.
func Video(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "video", Attributes: attrs, Children: children}
}

// Video_ is a convenience wrapper for Video without the attrs argument.
func Video_(children ...HTML) HTML {
  return Video(nil, children...)
}

// Void Elements

// Area represents the HTML void element 'area'.
// For more information visit https://www.w3schools.com/tags/tag_area.asp.
func Area(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "area", Attributes: attrs, SelfClosing: true}
}

// Area_ is a convenience wrapper for Area without the attrs argument.
func Area_() HTML {
  return Area(nil)
}

// Base represents the HTML void element 'base'.
// For more information visit https://www.w3schools.com/tags/tag_base.asp.
func Base(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "base", Attributes: attrs, SelfClosing: true}
}

// Base_ is a convenience wrapper for Base without the attrs argument.
func Base_() HTML {
  return Base(nil)
}

// Br represents the HTML void element 'br'.
// For more information visit https://www.w3schools.com/tags/tag_br.asp.
func Br(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "br", Attributes: attrs, SelfClosing: true}
}

// Br_ is a convenience wrapper for Br without the attrs argument.
func Br_() HTML {
  return Br(nil)
}

// Col represents the HTML void element 'col'.
// For more information visit https://www.w3schools.com/tags/tag_col.asp.
func Col(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "col", Attributes: attrs, SelfClosing: true}
}

// Col_ is a convenience wrapper for Col without the attrs argument.
func Col_() HTML {
  return Col(nil)
}

// Embed represents the HTML void element 'embed'.
// For more information visit https://www.w3schools.com/tags/tag_embed.asp.
func Embed(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "embed", Attributes: attrs, SelfClosing: true}
}

// Embed_ is a convenience wrapper for Embed without the attrs argument.
func Embed_() HTML {
  return Embed(nil)
}

// Hr represents the HTML void element 'hr'.
// For more information visit https://www.w3schools.com/tags/tag_hr.asp.
func Hr(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "hr", Attributes: attrs, SelfClosing: true}
}

// Hr_ is a convenience wrapper for Hr without the attrs argument.
func Hr_() HTML {
  return Hr(nil)
}

// Img represents the HTML void element 'img'.
// For more information visit https://www.w3schools.com/tags/tag_img.asp.
func Img(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "img", Attributes: attrs, SelfClosing: true}
}

// Img_ is a convenience wrapper for Img without the attrs argument.
func Img_() HTML {
  return Img(nil)
}

// Input represents the HTML void element 'input'.
// For more information visit https://www.w3schools.com/tags/tag_input.asp.
func Input(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "input", Attributes: attrs, SelfClosing: true}
}

// Input_ is a convenience wrapper for Input without the attrs argument.
func Input_() HTML {
  return Input(nil)
}

// Link represents the HTML void element 'link'.
// For more information visit https://www.w3schools.com/tags/tag_link.asp.
func Link(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "link", Attributes: attrs, SelfClosing: true}
}

// Link_ is a convenience wrapper for Link without the attrs argument.
func Link_() HTML {
  return Link(nil)
}

// Meta represents the HTML void element 'meta'.
// For more information visit https://www.w3schools.com/tags/tag_meta.asp.
func Meta(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "meta", Attributes: attrs, SelfClosing: true}
}

// Meta_ is a convenience wrapper for Meta without the attrs argument.
func Meta_() HTML {
  return Meta(nil)
}

// Param represents the HTML void element 'param'.
// For more information visit https://www.w3schools.com/tags/tag_param.asp.
func Param(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "param", Attributes: attrs, SelfClosing: true}
}

// Param_ is a convenience wrapper for Param without the attrs argument.
func Param_() HTML {
  return Param(nil)
}

// Source represents the HTML void element 'source'.
// For more information visit https://www.w3schools.com/tags/tag_source.asp.
func Source(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "source", Attributes: attrs, SelfClosing: true}
}

// Source_ is a convenience wrapper for Source without the attrs argument.
func Source_() HTML {
  return Source(nil)
}

// Track represents the HTML void element 'track'.
// For more information visit https://www.w3schools.com/tags/tag_track.asp.
func Track(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "track", Attributes: attrs, SelfClosing: true}
}

// Track_ is a convenience wrapper for Track without the attrs argument.
func Track_() HTML {
  return Track(nil)
}

// Wbr represents the HTML void element 'wbr'.
// For more information visit https://www.w3schools.com/tags/tag_wbr.asp.
func Wbr(attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "wbr", Attributes: attrs, SelfClosing: true}
}

// Wbr_ is a convenience wrapper for Wbr without the attrs argument.
func Wbr_() HTML {
  return Wbr(nil)
}
