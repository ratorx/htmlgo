package elements

import (
	"github.com/ratorx/htmlgo"
)

type HTML = htmlgo.HTML
type HTMLString = htmlgo.String
type HTMLUnescaped = htmlgo.Unescaped
type HTMLFunc = htmlgo.Func

func Attrs(attrs ...htmlgo.Attribute) []htmlgo.Attribute {
	return attrs
}

// Html initialises a new HTML document with the doctype set.
func Html(attrs []htmlgo.Attribute, children ...HTML) HTML {
	return HTMLFunc(func(indentLevel int, renderer *htmlgo.Renderer) {
		(&htmlgo.Tree{Tag: "!DOCTYPE HTML", SelfClosing: true}).Serialize(indentLevel, renderer)
		(&htmlgo.Tree{Tag: "html", Attributes: attrs, Children: children}).Serialize(indentLevel, renderer)
	})
}

// Html_ is a convenience wrapper for Html without the attrs argument.
func Html_(children ...HTML) HTML {
	return Html(nil, children...)
}
