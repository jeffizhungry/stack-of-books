// Package singleton describe the singleton design pattern in Go.
//
// Singletons are essentially a global variable wrapped in a class, and this
// pattern should be avoided when possible. Moreover, they make testing difficult
// if this singleton encapsulates access to external services. However there are
// some cases where it does make the code cleaner.
package singleton

// Exported package level singleton.
//
// I only recommend this if this singleton is going to be a READ-ONLY object,
// i.e. storing config vars. Otherwise this needs to be locked whenever it is
// set. If this is the case, it might be better to use a unexported singleton.
var One Singleton

// Unexported package level singleton.
//
// You can have a private singleton and give it package level functions that
// just access the singleton struct's methods. This allow users to call the singleton
// in a much cleaner manner.
//
//	Other Languages: singleton.GetInstance().DoSomething
//
//	Here:            singleton.DoSomething()
//
// This pattern in used in https://github.com/sirupsen/logrus/blob/master/exported.go
var hidden Singleton

func DoSomething() string {
	return hidden.DoSomething()
}

// You can also provide functions to set the global function.
func SetSingeton(s Singleton) {
	hidden = s
}

// If this object doesn't technically need to be singleton, and you are only providing
// a singleton to make it easier for users to use your package. You should give
// users the option to just use the class rather than the singleton. This will enable
// them inject this dependency in their testing.
type Singleton struct{}

func (s Singleton) DoSomething() string {
	return ""
}
