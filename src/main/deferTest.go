package main

type foo struct {
	s string
	i int
}

func (f *foo) close() {
	println("foo close:%s,%d", f.s, f.i)
}

func deferTest() {
	f1 := &foo{"a", 1}
	defer f1.close()

	*f1 = foo{"b", 2}
}
