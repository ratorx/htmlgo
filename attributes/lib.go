package attributes

import (
	"fmt"
	"strings"
)

// Attribute represents an HTML attribute.
type Attribute struct {
	Name   string
	Values []string
}

// Build serializes the Attribute to a strings.Builder
func (a Attribute) Build(builder *strings.Builder) {
	builder.WriteString(a.Name)
	builder.WriteRune('=')
	builder.WriteRune('"')
	builder.WriteString(strings.Join(a.Values, " "))
	builder.WriteRune('"')
}

// Dataset creates HTML data-* attributes.
func Dataset(key, value string) Attribute {
	return Attribute{Name: fmt.Sprintf("data-%s", key), Values: []string{value}}
}
