## Starting Out
- [Effective Go] (https://golang.org/doc/effective_go.html)

    * Great article detailing the nuances of the language.

- [Slices vs arrays] (https://blog.golang.org/slices)

    * TLDR; Arrays are values, slices are pointers (C idea of an array)

- [Concurrency vs Parallelism: VIDEO] (https://www.youtube.com/watch?v=cN_DpYBzKso)

    * TLDR; Concurrency inter-leaves tasks, Parallel is simultaneous tasks
    * Concurrent -> Parallel easy
    * But you should still watch anyways, because it is a great talk

- [Go Concurrency Patterns: VIDEO] (https://www.youtube.com/watch?v=f6kdp27TYZs)

    * Common Go patterns using go routines and channels



## Best Practices / Quick Tips

- [7 common mistakes in Go: VIDEO] (https://www.youtube.com/watch?v=29LLRKIL_TI)

    * By Docker COO, very good lessons

- [Using Go at The New York Times: VIDEO] (https://www.youtube.com/watch?v=bAQ9ShmXYLY)
    
    * good example of a full technology stack with Go at its core
    * this gives examples of building rpms, unittest coverage, gathering metrics, CICD

- [Gotchas and common mistakes in Go] (http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html)

- [Ten Useful Techniques in Go] (http://arslan.io/ten-useful-techniques-in-go)

- [Writing Testable Apps in Go] (http://relistan.com/writing-testable-apps-in-go/)

- [Stupid Gopher Tricks] (http://talks.golang.org/2015/tricks.slide#1)



## Articles
- [Using context structs rather than relying on globals] (http://elithrar.github.io/article/custom-handlers-avoiding-globals/)

- [Http handler error handling revisisted] (http://elithrar.github.io/article/http-handler-error-handling-revisited/)

- [Organizing Database Access] (http://www.alexedwards.net/blog/organising-database-access)

  * global vs. dependency injection vs others

- [Custom Request Scope Contexts] (https://joeshaw.org/net-context-and-http-handler/#custom-context-handler-types)

  * expand on golang.org/x/net/context

- [Quick Overview of Go database/sql library] (http://go-database-sql.org/index.html)

  * good explaination of nuances of the database/sql package

- [Structuring Applications in Go] (https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091#.867ujkhod)

- [Implementing RSA Encryption in Golang] (http://blog.giorgis.io/golang-rsa-encryption)

- [Why Go gets exceptions right] (http://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right)



## Development Environment
- [Kubernetes Development Workflow] (https://github.com/kubernetes/kubernetes/blob/master/docs/devel/development.md)

    * VERY GOOD description of dev workflow

- [How to write Go code] (https://golang.org/doc/code.html)

    * explain how to setup GOPATH

- [vim-go] (https://github.com/fatih/vim-go)

    * advanced IDE like vim plugin



## Advanced Concepts
- [Go Scheduler] (https://morsmachine.dk/go-scheduler)

    * a simplified version of the PDF below

- [Analysis of the Go runtime scheduler] (http://www.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf)
    
    * a more indepth look at the scheduler internals

- [Go blog: Garbage collection] (https://blog.golang.org/go15gc)

- [Go 1.5+ Gargage Collection In-depth] (https://docs.google.com/document/d/1wmjrocXIWTr1JxU-3EQBI6BK6KgtiFArkG47XK73xIQ/edit)

- [Go Execution Modes] (https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/edit)

    * go conditional compilation

- [Go blog: Request-scoped Contexts (net/context package)] (https://blog.golang.org/context)

- [Go-kit: Library for building large-scale distributed systems: VIDEO] (https://www.youtube.com/watch?v=1AjaZi4QuGo)



## Books
- [The Go Programming Language] (http://www.gopl.io/)

    * Basics, frankly effective go is probably more useful.

- [Build Web Application with Golang] (https://github.com/astaxie/build-web-application-with-golang/blob/master/en/preface.md)

- [MEAP: Go in Practice] (https://www.manning.com/books/go-in-practice)

    * goes beyond just language syntax, VERY GOOD
