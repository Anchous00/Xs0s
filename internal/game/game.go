package game

import (
	"Xs0s/utils/customButton"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var Field fyne.Container = *container.NewWithoutLayout()

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

	Field.Add(line1)
	Field.Add(line2)
	Field.Add(line3)
	Field.Add(line4)

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
	line1.StrokeWidth = 10
	line1.Position1 = fyne.NewPos(x+7.06, y+7.06)
	line1.Position2 = fyne.NewPos(x+192.94, y+192.94)

	line2 := canvas.NewLine(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	line2.StrokeWidth = 10
	line2.Position1 = fyne.NewPos(x+192.94, y+7.06)
	line2.Position2 = fyne.NewPos(x+7.06, y+192.94)

	Field.Add(line1)
	Field.Add(line2)
	Field.Refresh()

}

func AddButton(x, y float32) {
	button := customButton.NewCustomTappableRectangle(color.Black, func() {
		if g.field[int(x/200)][int(y/200)] == "" && !CheckDraw() && !CheckWin() {
			MakeMove(x, y)
		}
	})
	button.Resize(fyne.NewSize(195, 195))
	button.Move(fyne.NewPos(x+2.5, y+2.5))

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
	i := x / 200
	j := y / 200

	switch g.current {
	case "X":
		DrawX(x, y)
		g.current = "0"
		g.field[int(i)][int(j)] = "X"
	case "0":
		DrawCircle(x, y)
		g.current = "X"
		g.field[int(i)][int(j)] = "O"
	}
	if CheckWin() {
		OfferNewGame(g.field[int(i)][int(j)])
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
	button.Resize(fyne.NewSize(400, 300))
	button.Move(fyne.NewPos(100, 150))
	Field.Add(button)
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
	MakeGrid()
	AddButttons()
	g = NewGame()
}
