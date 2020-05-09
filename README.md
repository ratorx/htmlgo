# ratorx/htmlgo

A library for writing type-safe HTML in Go. Forked from [github.com/julvo/htmlgo](https://github.com/julvo/htmlgo).

Differences:

* The threat model is different. All text sanitisation is explicitly user-specified using `HTMLUString`.
* `HTML` is now an interface. This allows easier creation of stateful components (rather than building closures).
* The actual HTML is generated lazily.

Potential future directions:
* Prevent invalid direct parent-child element combinations. This should be relatively easy to do. A dummy interface can be generated for each parent class and the HTML elements that implement the interface can be restricted. The tedious bit is populating the list of constraints.
* Prevent invalid attribute-element combinations.
* Investigate transitively invalid trees (e.g. when a certain parent cannot contain have a child node of specific types anywhere below it).
* Investigate rejecting input at compile time when the above points combine.
