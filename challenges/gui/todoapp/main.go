package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func updateTimer(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	fmt.Println("Todo App Started...")

	a := app.New()
	w := a.NewWindow("Todo App")

	// Clock Widget
	clockWidget := createClock()

	// Grid Layout For The Button, Clock Widget
	grid := container.New(layout.NewFormLayout(), clockWidget, widget.NewButton("Click Me!", func() {
		fmt.Println("Button has been clicked!")

		w2 := a.NewWindow("Sub Window")
		w2.SetContent(widget.NewLabel("You clicked the button!"))
		w2.Resize(fyne.NewSize(200, 200))
		w2.Show()
	}))
	updateTimer(clockWidget)

	// Setting The Content Of The Main Window
	w.SetContent(grid)

	// Resizing The Window
	w.Resize(fyne.NewSize(400, 400))

	go func() {
		for range time.Tick(time.Second) {
			updateTimer(clockWidget)
		}
	}()

	w.ShowAndRun()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Tidying up...\nExited.")
}

func createClock() *widget.Label {
	clock := widget.NewLabel("")
	return clock
}
