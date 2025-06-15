package main

import (
	"Xs0s/internal/game"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {

	app := app.New()
	window := app.NewWindow("Hello World")
	defer window.Close()
	window.Resize(fyne.NewSize(600, 600))
	game.StartNewGame()

	content := container.NewWithoutLayout(&game.Field)

	window.SetContent(content)
	window.ShowAndRun()

}
