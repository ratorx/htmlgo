package htmlgo

import (
  a "github.com/ratorx/htmlgo/attributes"
)
[[ range .ElementFuncs ]]

// [[.FuncName]] represents the HTML element '[[.TagName]]'.
// For more information visit https://www.w3schools.com/tags/tag_[[.TagName]].asp.
func [[.FuncName]](attrs []a.Attribute, children ...HTML) HTML {
	return &HTMLTree{Tag: "[[.TagName]]", Attributes: attrs, Children: children}
}

// [[.FuncName]]_ is a convenience wrapper for [[.FuncName]] without the attrs argument.
func [[.FuncName]]_(children ...HTML) HTML {
  return [[.FuncName]](nil, children...)
}
[[ end ]]

// Void Elements
[[ range .VoidElementFuncs ]]

// [[.FuncName]] represents the HTML void element '[[.TagName]]'.
// For more information visit https://www.w3schools.com/tags/tag_[[.TagName]].asp.
func [[.FuncName]](attrs []a.Attribute) HTML {
	return &HTMLTree{Tag: "[[.TagName]]", Attributes: attrs, SelfClosing: true}
}

// [[.FuncName]]_ is a convenience wrapper for [[.FuncName]] without the attrs argument.
func [[.FuncName]]_() HTML {
  return [[.FuncName]](nil)
}
[[ end ]]
