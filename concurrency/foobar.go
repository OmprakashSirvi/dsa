package concurrency

type FooBar struct {
	n       int
	FooChan chan struct{}
	BarChan chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb:= &FooBar{
		n:       n,
		FooChan: make(chan struct{}, 1),
		BarChan: make(chan struct{}, 1),
	}

	// Start the execution
	fb.FooChan <- struct{}{}

	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.FooChan
		printFoo()
		fb.BarChan <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.BarChan
		printBar()
		fb.FooChan <- struct{}{}
	}
}
