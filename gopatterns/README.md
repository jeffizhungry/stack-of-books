#### Go Patterns

The purpose of this guide is to provide ideas, not to define explicit structures. 
This is why I did not define names for each pattern. Developers should pick and 
choose from this guide as necessary to fit their use case.

Go patterns are [not the same as OOP patterns](https://groups.google.com/forum/#!msg/golang-nuts/3fOIZ1VLn1o/GeE1z5qUA6YJ). These are patterns specific
to the Go language, but contains ideas such as [SOLID design principles] (https://dave.cheney.net/2016/08/20/solid-go-design).

I believe patterns allow for cleaner code, more modular and testable code when 
used appropriately.

The benefits of using patterns:
    - Familiar patterns across different codebases
    - Better modularity and testabilty
        - there are clear places where code boundaries should be defined
    - Better code quality
        - reminds developers to add validators, backoffs, retries, timeouts etc...
        - conventions to make code easier to reason about

The issues with using patterns:
    - People try forcing pattern onto a simple problem
    - People stop thinking of new patterns that might be better suited for the prob.
    - Using patterns to overly abstract your code unecessarily
        - abstractions are nice for users of a package
        - too many abstractions make it difficult to debug a package
