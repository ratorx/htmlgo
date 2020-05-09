package attributes

import (
	"fmt"

	"github.com/ratorx/htmlgo"
)

// Dataset creates HTML data-* attributes.
func Dataset(key, value string) htmlgo.Attribute {
	return htmlgo.Attribute{Name: fmt.Sprintf("data-%s", key), Values: []string{value}}
}
