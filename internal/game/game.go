package game

import (
	"Xs0s/utils/customButton"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var Field = *container.NewWithoutLayout()
var Menu = *container.NewWithoutLayout()
var App = app.New()
var window = App.NewWindow("Hello World")

type Game struct {
	current string
	field   [3][3]string
}

var g = NewGame()

func NewGame() *Game {
	game := &Game{current: "X"}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.field[i][j] = ""
		}
	}
	return game
}

func MakeGrid() {
	line1 := canvas.NewLine(color.White)
	line1.StrokeWidth = 5
	line1.Position1 = fyne.NewPos(200, 0)
	line1.Position2 = fyne.NewPos(200, 600)

	line2 := canvas.NewLine(color.White)
	line2.StrokeWidth = 5
	line2.Position1 = fyne.NewPos(400, 0)
	line2.Position2 = fyne.NewPos(400, 600)

	line3 := canvas.NewLine(color.White)
	line3.StrokeWidth = 5
	line3.Position1 = fyne.NewPos(0, 200)
	line3.Position2 = fyne.NewPos(600, 200)

	line4 := canvas.NewLine(color.White)
	line4.StrokeWidth = 5
	line4.Position1 = fyne.NewPos(0, 400)
	line4.Position2 = fyne.NewPos(600, 400)

	line5 := canvas.NewLine(color.White)
	line5.StrokeWidth = 5
	line5.Position1 = fyne.NewPos(0, -3)
	line5.Position2 = fyne.NewPos(0, 603)

	line6 := canvas.NewLine(color.White)
	line6.StrokeWidth = 5
	line6.Position1 = fyne.NewPos(-3, 0)
	line6.Position2 = fyne.NewPos(603, 0)

	line7 := canvas.NewLine(color.White)
	line7.StrokeWidth = 5
	line7.Position1 = fyne.NewPos(595, 0)
	line7.Position2 = fyne.NewPos(595, 603)

	line8 := canvas.NewLine(color.White)
	line8.StrokeWidth = 5
	line8.Position1 = fyne.NewPos(0, 595)
	line8.Position2 = fyne.NewPos(603, 595)

	Field.Add(line1)
	Field.Add(line2)
	Field.Add(line3)
	Field.Add(line4)
	Field.Add(line5)
	Field.Add(line6)
	Field.Add(line7)
	Field.Add(line8)

}

func DrawCircle(x, y float32) {
	circle := canvas.NewCircle(color.NRGBA{R: 0, G: 255, B: 255, A: 255})
	circle.Resize(fyne.NewSize(200, 200))
	circle.Move(fyne.NewPos(x, y))
	circle.StrokeWidth = 5

	circle2 := canvas.NewCircle(color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	circle2.Resize(fyne.NewSize(195, 195))
	circle2.Move(fyne.NewPos(x+2.5, y+2.5))
	circle2.StrokeWidth = 5

	Field.Add(circle)
	Field.Add(circle2)

	Field.Refresh()
}

func DrawX(x, y float32) {
	line1 := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line1.StrokeWidth = 5
	line1.Position1 = fyne.NewPos(x, y)
	line1.Position2 = fyne.NewPos(x+200, y+200)

	line2 := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line2.StrokeWidth = 5
	line2.Position1 = fyne.NewPos(x+200, y)
	line2.Position2 = fyne.NewPos(x, y+200)

	Field.Add(line1)
	Field.Add(line2)
	MakeGrid()
	Field.Refresh()

}

func AddButton(x, y float32) {
	button := customButton.NewCustomTappableRectangle(color.Black, func() {
		if g.field[int(x/200)][int(y/200)] == "" && !CheckDraw() && !CheckWin() {
			MakeMove(x, y)
		}
	})
	button.Resize(fyne.NewSize(200, 200))
	button.Move(fyne.NewPos(x, y))

	Field.Add(button)
}

func AddButttons() {
	for i := float32(0); i < 3; i++ {
		for j := float32(0); j < 3; j++ {
			AddButton(i*200, j*200)
		}
	}
}

func MakeMove(x, y float32) {
	i := int(x / 200)
	j := int(y / 200)

	switch g.current {
	case "X":
		if g.field[i][j] == "" {
			DrawX(x, y)
			g.current = "0"
			g.field[i][j] = "X"
		}
	case "0":
		if g.field[i][j] == "" {
			DrawCircle(x, y)
			g.current = "X"
			g.field[i][j] = "O"
		}
	}
	if CheckWin() {
		OfferNewGame(g.field[i][j])
	}
	if CheckDraw() {
		OfferNewGame("nobody")
	}
}

func OfferNewGame(winner string) {
	button := widget.NewButton("winner is "+winner+"\n\n\nNew game", func() {
		Field.RemoveAll()
		StartNewGame()
	})
	exitButton := widget.NewButton("exit game", func() { ShowMenu() })
	exitButton.Resize(fyne.NewSize(200, 100))
	exitButton.Move(fyne.NewPos(0, 100))

	button.Resize(fyne.NewSize(200, 100))
	button.Move(fyne.NewPos(0, 0))
	buttons := container.NewWithoutLayout(button, exitButton)
	buttons.Move(fyne.NewPos(200, 200))
	Field.Add(buttons)
}

func CheckWin() bool {
	for row := 0; row < 3; row++ {
		if g.field[row][0] != "" && g.field[row][0] == g.field[row][1] && g.field[row][1] == g.field[row][2] {
			return true
		}
	}

	for col := 0; col < 3; col++ {
		if g.field[0][col] != "" && g.field[0][col] == g.field[1][col] && g.field[1][col] == g.field[2][col] {
			return true
		}
	}

	if g.field[0][0] != "" && g.field[0][0] == g.field[1][1] && g.field[1][1] == g.field[2][2] {
		return true
	}
	if g.field[0][2] != "" && g.field[0][2] == g.field[1][1] && g.field[1][1] == g.field[2][0] {
		return true
	}

	return false
}

func CheckDraw() bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g.field[row][col] == "" {
				return false
			}
		}
	}
	return true
}

func StartNewGame() {
	AddButttons()
	MakeGrid()
	g = NewGame()
}

func StartGame() {

	window.Resize(fyne.NewSize(600, 600))
	StartNewGame()

	content := container.NewWithoutLayout(&Field)
	window.SetContent(content)
}

func ShowMenu() {

	window.Resize(fyne.NewSize(600, 600))

	NewGameButton := widget.NewButton("New Game", func() {
		StartGame()
	})

	NewGameButton.Resize(fyne.NewSize(400, 200))
	NewGameButton.Move(fyne.NewPos(100, 100))

	ExitButton := widget.NewButton("Exit", func() { App.Quit() })
	ExitButton.Resize(fyne.NewSize(200, 100))
	ExitButton.Move(fyne.NewPos(200, 300))
	content := container.NewWithoutLayout(NewGameButton, ExitButton)
	window.SetContent(content)
}

func RunApp() {
	ShowMenu()
	window.SetFixedSize(true)
	window.ShowAndRun()
}
