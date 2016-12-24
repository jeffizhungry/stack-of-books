package testing

/******************************************************************************
 * Testing as a public interface
 *****************************************************************************/
// I usually put all my testing related functions and structs in testing.go

type MockStudent struct {
}

func GetFakeHomework() string {
	return "canned data"
}

func GetFakeLateHomework() string {
	return "canned data"
}

/******************************************************************************
 * Testing helpers
 *****************************************************************************/
// Testing helper functions and structs that many not be useful to users of
// this package can just be unexported.

func getFakeHomework() string {
	return "canned data"
}
