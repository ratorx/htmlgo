package elements

import (
  "github.com/ratorx/htmlgo"
)
[[- range .ElementFuncs ]]

// [[.FuncName]] represents the HTML element '[[.TagName]]'.
// For more information visit https://www.w3schools.com/tags/tag_[[.TagName]].asp.
func [[.FuncName]](attrs []htmlgo.Attribute, children ...HTML) HTML {
	return &htmlgo.Tree{Tag: "[[.TagName]]", Attributes: attrs, Children: children}
}

// [[.FuncName]]_ is a convenience wrapper for [[.FuncName]] without the attrs argument.
func [[.FuncName]]_(children ...HTML) HTML {
  return [[.FuncName]](nil, children...)
}
[[- end ]]

// Void Elements
[[- range .VoidElementFuncs ]]

// [[.FuncName]] represents the HTML void element '[[.TagName]]'.
// For more information visit https://www.w3schools.com/tags/tag_[[.TagName]].asp.
func [[.FuncName]](attrs []htmlgo.Attribute) HTML {
	return &htmlgo.Tree{Tag: "[[.TagName]]", Attributes: attrs, SelfClosing: true}
}

// [[.FuncName]]_ is a convenience wrapper for [[.FuncName]] without the attrs argument.
func [[.FuncName]]_() HTML {
  return [[.FuncName]](nil)
}
[[- end ]]
