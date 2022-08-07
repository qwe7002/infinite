package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {
	sp := inf.NewSpinner(
		spinner.WithShape(components.Dot),
		//spinner.WithDisableOutputResult(),
	).Display()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 100)
			sp.Refreshf("hello world %d", i)
		}

		sp.Finish("finish")

		sp.Refresh("is finish?")

	}()

	time.Sleep(time.Millisecond * 100 * 15)
}