package htmlgo

import "strings"

// Attribute represents an HTML attribute.
type Attribute struct {
	Name   string
	Values []string
}

// Serialize serializes the Attribute to a strings.Builder
func (a Attribute) Serialize(builder *strings.Builder) {
	builder.WriteString(a.Name)
	builder.WriteRune('=')
	builder.WriteRune('"')
	builder.WriteString(strings.Join(a.Values, " "))
	builder.WriteRune('"')
}

// Tree is the main way to build HTML documents. It represents an HTML element with a slice of its children. If SelfClosing is true, it represents a HTML element with no children. The Children field is ignored when SelfClosing is true.
type Tree struct {
	Tag         string
	Attributes  []Attribute
	Children    []HTML
	SelfClosing bool
}

func (h *Tree) Serialize(indentLevel int, renderer *Renderer) {
	// Tag open
	Indent(indentLevel, renderer)
	renderer.WriteRune('<')
	renderer.WriteString(h.Tag)
	for _, attr := range h.Attributes {
		renderer.WriteRune(' ')
		attr.Serialize(&renderer.Builder)
	}
	renderer.WriteRune('>')

	if !h.SelfClosing {
		if len(h.Children) != 0 {
			renderer.WriteRune('\n')
		}

		// Children
		for _, child := range h.Children {
			child.Serialize(indentLevel+1, renderer)
			renderer.WriteRune('\n')
		}

		// Tag close
		Indent(indentLevel, renderer)
		renderer.WriteRune('<')
		renderer.WriteRune('/')
		renderer.WriteString(h.Tag)
		renderer.WriteRune('>')
	}
}
