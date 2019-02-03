package main

import (
	"syscall/js"
	"time"
	"unicode"

	"github.com/dennwc/dom"
)

// Format Duration
func fd(d time.Duration) string {
	var ns string
	var period bool
	for _, r := range d.String() {
		if (period && unicode.IsDigit(r)) || r == '.' {
			period = true
			continue
		}
		ns += string(r)
	}
	return ns
}

func main() {
	var working stopwatch
	var distracted stopwatch

	js.Global().Set("start", js.NewCallback(func(v []js.Value) {
		dom.GetDocument().GetElementById("startbutton")
		working = newStopwatch(true)
		distracted = stopwatch{}
		js.Global().Call("disableButtons", true)
	}))
	js.Global().Set("flip", js.NewCallback(func(v []js.Value) {
		working.Flip()
		distracted.Flip()
	}))
	js.Global().Set("stop", js.NewCallback(func(v []js.Value) {
		working.Stop()
		distracted.Stop()
		js.Global().Call("disableButtons", false)
	}))

	for {
		dom.GetDocument().GetElementById("working").SetInnerHTML(fd(working.Elapsed()))
		dom.GetDocument().GetElementById("distracted").SetInnerHTML(fd(distracted.Elapsed()))
		dom.ConsoleLog("\rWorking: %s (%s)\t\tDistracted %s (%s)\t\t", fd(working.Elapsed()), fd(working.lastElapsed), fd(distracted.Elapsed()), fd(distracted.lastElapsed))
		time.Sleep(time.Millisecond * 100)
	}
}
