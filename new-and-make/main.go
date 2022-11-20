package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
	foo    int
	bar    string
}

func main() {
	// p := new(SyncedBuffer)
	p := NewSynced(100, "foobar")
	fmt.Println("foo:", p.foo)
	fmt.Println("bar:", p.bar)
	fmt.Printf("%#v\n", p)

	// make
	foobar := make(map[string]string)
	foobar["foo"] = "bar"
	foobar["bar"] = "foo"
	fmt.Println(foobar)

	Foo(&foobar)
}

func NewSynced(foo int, bar string) *SyncedBuffer {
	return &SyncedBuffer{
		foo: foo,
		bar: bar,
	}
}

func Foo(v *map[string]string) {
	fmt.Printf("%#v\n", v)
}
