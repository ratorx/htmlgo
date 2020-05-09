//go:generate go run github.com/ratorx/htmlgo/htmlgogen
package htmlgo

import (
	"strings"
	"text/template"

	"github.com/ratorx/htmlgo/attributes"
)

// Builder is a wrapper around strings.Builder that contains extra configuration for building HTML.
type Builder struct {
	strings.Builder
	// TabSize specified the number of spaces to use for indentation.
	TabSize int
}

// Indent is a convenience function to indent a new line upto indentLevel
func Indent(indentLevel int, builder *Builder) {
	builder.WriteString(strings.Repeat(" ", indentLevel*builder.TabSize))
}

// HTML interface represents all types that can be converted to an HTML string.
type HTML interface {
	Build(indentLevel int, builder *Builder)
}

// HTMLTree is the main way to build HTML documents. It represents an HTML element with pointers to its children. If SelfClosing is true, it represents a HTML element without children. The Children field is ignored when SelfClosing is true.
type HTMLTree struct {
	Tag         string
	Attributes  []attributes.Attribute
	Children    []HTML
	SelfClosing bool
}

func (h *HTMLTree) Build(indentLevel int, builder *Builder) {
	// Tag open
	Indent(indentLevel, builder)
	builder.WriteRune('<')
	builder.WriteString(h.Tag)
	for _, attr := range h.Attributes {
		builder.WriteRune(' ')
		attr.Build(&builder.Builder)
	}
	builder.WriteRune('>')

	if !h.SelfClosing {
		if len(h.Children) != 0 {
			builder.WriteRune('\n')
		}

		// Children
		for _, child := range h.Children {
			child.Build(indentLevel+1, builder)
			builder.WriteRune('\n')
		}

		// Tag close
		Indent(indentLevel, builder)
		builder.WriteRune('<')
		builder.WriteRune('/')
		builder.WriteString(h.Tag)
		builder.WriteRune('>')
	}
}

// HTMLString provides a way to insert HTML from outside this library.
type HTMLString string

func (h HTMLString) Build(_ int, builder *Builder) {
	builder.WriteString(string(h))
}

// HTMLText represents unsanitised HTML content. It is escaped when building the document.
type HTMLText string

func (h HTMLText) Build(_ int, builder *Builder) {
	template.HTMLEscape(builder, []byte(h))
}

// HTMLFunc is an adapter type to use any Build function without creating a new type.
type HTMLFunc func(indentLevel int, builder *Builder)

func (h HTMLFunc) Build(indentLevel int, builder *Builder) {
	h(indentLevel, builder)
}

// Html initialises a new HTML document with the doctype set.
func Html(attrs []attributes.Attribute, children ...HTML) HTML {
	return HTMLFunc(func(indentLevel int, builder *Builder) {
		(&HTMLTree{Tag: "!DOCTYPE HTML", SelfClosing: true}).Build(indentLevel, builder)
		(&HTMLTree{Tag: "html", Attributes: attrs, Children: children}).Build(indentLevel, builder)
	})
}

// Html_ is a convenience wrapper for Html without the attrs argument.
func Html_(children ...HTML) HTML {
	return Html(nil, children...)
}
