//go:generate go run github.com/ratorx/htmlgo/htmlgogen
package htmlgo

import (
	"html/template"
	"strings"
)

// Renderer is a wrapper around strings.Builder that contains extra configuration for rendering HTML.
type Renderer struct {
	strings.Builder
	// TabSize specified the number of spaces to use for indentation.
	TabSize int
}

func (h *Renderer) Render(html HTML) string {
	h.Builder.Reset()
	html.Serialize(0, h)
	return h.Builder.String()
}

// Indent is a convenience function to indent a new line upto indentLevel
func Indent(indentLevel int, renderer *Renderer) {
	renderer.WriteString(strings.Repeat(" ", indentLevel*renderer.TabSize))
}

// HTML interface represents all types that can be converted to an HTML string.
type HTML interface {
	Serialize(indentLevel int, renderer *Renderer)
}

// String provides a way to insert HTML from outside this library.
type String string

func (h String) Serialize(_ int, renderer *Renderer) {
	renderer.WriteString(string(h))
}

// Unescaped represents unescaped HTML content. It is escaped when building the document.
type Unescaped string

func (h Unescaped) Serialize(_ int, renderer *Renderer) {
	template.HTMLEscape(renderer, []byte(h))
}

// Func is an adapter type to implement HTML without creating a new type.
type Func func(indentLevel int, renderer *Renderer)

func (h Func) Serialize(indentLevel int, renderer *Renderer) {
	h(indentLevel, renderer)
}
