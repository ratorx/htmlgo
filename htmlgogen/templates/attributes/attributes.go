package attributes

import (
	"github.com/ratorx/htmlgo"
)
[[- range .AttributeFuncs ]]

// [[.FuncName]] represents the HTML attribute '[[.AttrName]]'.
func [[.FuncName]](values ...string) htmlgo.Attribute {
	return htmlgo.Attribute{Name: "[[.AttrName]]", Values: values}
}
[[- end ]]
