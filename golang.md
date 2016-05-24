## Starting Out
- [Golang Sandbox] (https://tour.golang.org/welcome/1)
   
   * Great first dive into Golang syntax.

- [Effective Go] (https://golang.org/doc/effective_go.html)

    * Great article detailing the nuances of the language.

- [Simplicity is Complicated: VIDEO] (https://www.youtube.com/watch?v=rFejpH_tAHM)

    * Great intro on various aspects of the language from a high level perspective.

- [Concurrency vs Parallelism: VIDEO] (https://www.youtube.com/watch?v=cN_DpYBzKso)

    * Concurrency is managing multiple tasks (on one core, as a pipeline, etc.)
    * Parallelism doing multiple tasks simultaneously
    * But you should still watch anyways, because it is a great talk

- [Channel Axioms] (http://dave.cheney.net/2014/03/19/channel-axioms)

    * You need to know this!

- [Go Proverbs] (https://go-proverbs.github.io/)

    * This is actually really important if you understand what each of these mean.

- [Visualizing Concurrency in Go] (https://divan.github.io/posts/go_concurrency_visualize/)

    * Really good visual of concurrency patterns in Go, and how go routines
    interact with each other through message passing


## Best Practices / Quick Tips
- [7 common mistakes in Go: VIDEO] (https://www.youtube.com/watch?v=29LLRKIL_TI)

    * By Docker COO, very good lessons

- [Using Go at The New York Times: VIDEO] (https://www.youtube.com/watch?v=bAQ9ShmXYLY)
    
    * good example of a full technology stack with Go at its core
    * this gives examples of building rpms, unittest coverage, gathering metrics, CICD

- [50 Shades of Go: Gotchas and common mistakes in Go] (http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html)

- [Ten Useful Techniques in Go] (http://arslan.io/ten-useful-techniques-in-go)

- [Writing Testable Apps in Go] (http://relistan.com/writing-testable-apps-in-go/)

- [Stupid Gopher Tricks] (http://talks.golang.org/2015/tricks.slide#1)

- [Whats in a name] (https://talks.golang.org/2014/names.slide#1)

   * Good presentation on good Golangic naming conventions

- [CodeReviewComments] (https://github.com/golang/go/wiki/CodeReviewComments)

   * Stylistic Guide that highlights various Golang philosophies used in the Go standard library

- [Go best practices, six years in] (https://peter.bourgon.org/go-best-practices-2016/) 

   * Some interesting insights, and useful links


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

- [Golang concepts from an OOP point of view] (https://github.com/luciotato/golang-notes/blob/master/OOP.md)


## Advanced Concepts
- [Go Scheduler] (https://morsmachine.dk/go-scheduler)

    * a simplified version of the PDF below

- [Analysis of the Go runtime scheduler] (http://www.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf)
    
    * a more indepth look at the scheduler internals

- [Go blog: Garbage collection] (https://blog.golang.org/go15gc)

- [Go 1.5+ Gargage Collection In-depth] (https://docs.google.com/document/d/1wmjrocXIWTr1JxU-3EQBI6BK6KgtiFArkG47XK73xIQ/edit)

- [Go GC: Solving the Latency Problem] (https://www.youtube.com/watch?v=aiv1JOfMjm0)
    
    * advances on traditional GCs with write barriers

- [Go Execution Modes] (https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/edit)

    * go conditional compilation --> necessary for platform specifc code

- [Summary of Go Generics Discussions] (https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/edit#)

- [Go blog: Request-scoped Contexts (net/context package)] (https://blog.golang.org/context)

- [Go-kit: Library for building large-scale distributed systems: VIDEO] (https://www.youtube.com/watch?v=1AjaZi4QuGo)

- [Make a RESTful JSON API in Go] (http://thenewstack.io/make-a-restful-json-api-go/)

- [Five things that make Go fast] (http://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)


## Pitfalls
- [Go channels are bad and you should feel bad] (http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/)

    * interesting perspective


## Testing
* **go vet** is for correctness, **golint** is for coding style

- [Interfaces and composition for effective unittesting in Golang] (http://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/)

- [Writing table driven tests in Go] (http://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)


## Development Environment
- [Kubernetes Development Workflow] (https://github.com/kubernetes/kubernetes/blob/master/docs/devel/development.md)

    * VERY GOOD description of dev workflow

- [vim-go] (https://github.com/fatih/vim-go)

    * advanced IDE like vim plugin


## Books
- [MEAP: Go in Practice] (https://www.manning.com/books/go-in-practice)

    * goes beyond just language syntax, VERY GOOD

- [Build Web Application with Golang] (https://github.com/astaxie/build-web-application-with-golang/blob/master/en/preface.md)
