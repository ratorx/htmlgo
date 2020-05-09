package attributes

[[ range .AttributeFuncs ]]

// [[.FuncName]] represents the HTML attribute '[[.AttrName]]'.
func [[.FuncName]](values ...string) Attribute {
	return Attribute{Name: "[[.AttrName]]", Values: values}
}
[[ end ]]
