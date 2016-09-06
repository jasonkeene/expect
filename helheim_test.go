// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package expect_test

type mockT struct {
	ErrorfCalled chan bool
	ErrorfInput  struct {
		Format chan string
		Args   chan []interface{}
	}
	FatalCalled chan bool
	FatalInput  struct {
		Arg0 chan []interface{}
	}
	FailNowCalled chan bool
}

func newMockT() *mockT {
	m := &mockT{}
	m.ErrorfCalled = make(chan bool, 100)
	m.ErrorfInput.Format = make(chan string, 100)
	m.ErrorfInput.Args = make(chan []interface{}, 100)
	m.FatalCalled = make(chan bool, 100)
	m.FatalInput.Arg0 = make(chan []interface{}, 100)
	m.FailNowCalled = make(chan bool, 100)
	return m
}
func (m *mockT) Errorf(format string, args ...interface{}) {
	m.ErrorfCalled <- true
	m.ErrorfInput.Format <- format
	m.ErrorfInput.Args <- args
}
func (m *mockT) Fatal(arg0 ...interface{}) {
	m.FatalCalled <- true
	m.FatalInput.Arg0 <- arg0
}
func (m *mockT) FailNow() {
	m.FailNowCalled <- true
}
