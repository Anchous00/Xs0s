package game

import (
	server "Xs0s/internal/connection"
	"Xs0s/internal/user"
	"Xs0s/utils/customButton"
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var Field = *container.NewWithoutLayout()
var Menu = *container.NewWithoutLayout()
var Logger = *container.NewWithoutLayout()
var App = app.New()
var window = App.NewWindow("Конченная игра")

var Player user.User

var g = NewGame()

func NewGame() *server.Game {
	game := &server.Game{Current: 'X'}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Field[i][j] = ' '
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
	line6.Position1 = fyne.NewPos(0, 0)
	line6.Position2 = fyne.NewPos(603, 0)

	line7 := canvas.NewLine(color.White)
	line7.StrokeWidth = 5
	line7.Position1 = fyne.NewPos(600, 0)
	line7.Position2 = fyne.NewPos(600, 603)

	line8 := canvas.NewLine(color.White)
	line8.StrokeWidth = 5
	line8.Position1 = fyne.NewPos(0, 600)
	line8.Position2 = fyne.NewPos(603, 600)

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

func DrawField() {
	for i := float32(0); i < 3; i++ {
		for j := float32(0); j < 3; j++ {
			if g.Field[int(i)][int(j)] == 'X' {
				DrawX(i*200, j*200)
			}
			if g.Field[int(i)][int(j)] == '0' {
				DrawCircle(i*200, j*200)
			}
		}
	}
}

func AddButton(x, y float32) {
	button := customButton.NewCustomTappableRectangle(color.Black, func() {
		if g.Field[int(x/200)][int(y/200)] == ' ' && !CheckDraw() && !CheckWin() {
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
	if (Player.Char) != g.Current {
		return
	}
	switch g.Current {
	case 'X':
		if g.Field[i][j] == ' ' && !CheckDraw() && !CheckWin() {
			g.Current = '0'
			g.Field[i][j] = 'X'
			signal := make([]byte, 3)
			signal = append(signal, byte('X'))
			signal = append(signal, byte(i))
			signal = append(signal, byte(j))
			server.SendMove(signal)
			DrawField()
		}
	case '0':
		if g.Field[i][j] == ' ' && !CheckDraw() && !CheckWin() {
			g.Current = 'X'
			g.Field[i][j] = '0'
			fmt.Println(g.Field)
			DrawField()
			var signal []byte
			signal = append(signal, byte('0'))
			signal = append(signal, byte(i))
			signal = append(signal, byte(j))
			server.SendMove(signal)
		}

	}

	if CheckWin() {
		OfferNewGame(g.Field[i][j])
	}
	if CheckDraw() {
		OfferNewGame(' ')
	}

	DrawField()
}

func OfferNewGame(winner byte) {

	TextNewGame := container.NewCenter(
		widget.NewLabel("winner is "+string(winner)),
		widget.NewLabel("\n\n\nNew game"),
	)

	btn := container.New(
		layout.NewStackLayout(),
		widget.NewButton("", func() {
			Field.RemoveAll()
			StartNewGame()
		}),
		canvas.NewRectangle(color.RGBA{R: 0, G: 127, B: 0, A: 255}),
		TextNewGame,
	)

	TextExit := container.NewCenter(
		widget.NewLabel("Exit game"),
	)
	btnExit := container.New(
		layout.NewStackLayout(),
		widget.NewButton("", func() { ShowMenu() }),
		canvas.NewRectangle(color.RGBA{R: 127, G: 0, B: 0, A: 255}),
		TextExit,
	)

	btn.Resize(fyne.NewSize(200, 100))
	btn.Move(fyne.NewPos(0, 0))

	btnExit.Resize(fyne.NewSize(200, 100))
	btnExit.Move(fyne.NewPos(0, 100))

	buttons := container.NewWithoutLayout(btn, btnExit)
	buttons.Move(fyne.NewPos(200, 200))

	Field.Add(buttons)
	MakeGrid()
}

func CheckWin() bool {
	for row := 0; row < 3; row++ {
		if g.Field[row][0] != ' ' && g.Field[row][0] == g.Field[row][1] && g.Field[row][1] == g.Field[row][2] {
			return true
		}
	}

	for col := 0; col < 3; col++ {
		if g.Field[0][col] != ' ' && g.Field[0][col] == g.Field[1][col] && g.Field[1][col] == g.Field[2][col] {
			return true
		}
	}

	if g.Field[0][0] != ' ' && g.Field[0][0] == g.Field[1][1] && g.Field[1][1] == g.Field[2][2] {
		return true
	}
	if g.Field[0][2] != ' ' && g.Field[0][2] == g.Field[1][1] && g.Field[1][1] == g.Field[2][0] {
		return true
	}

	return false
}

func CheckDraw() bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g.Field[row][col] == ' ' || CheckWin() {
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

	window.Resize(fyne.NewSize(608, 608))
	StartNewGame()

	content := container.NewWithoutLayout(&Field)
	window.SetContent(content)
}

func ShowMenu() {

	window.Resize(fyne.NewSize(608, 608))

	OLEG := canvas.NewImageFromFile("internal/user/OLEG.jpg")
	content := container.NewStack(OLEG)
	content.Resize(fyne.NewSize(600, 600))
	content.Move(fyne.NewPos(0, 0))
	if Player.Username == "Oleg" {
		Menu.Add(content)
	}

	CreateButton := widget.NewButton("Create Game", func() {
		server.StartServer(g.Field)
		time.Sleep(10 * time.Millisecond)
		Player.Char = 'X'
		StartGame()
	})
	CreateButton.Resize(fyne.NewSize(400, 200))
	CreateButton.Move(fyne.NewPos(100, 100))

	ExitButton := widget.NewButton("Exit", func() { App.Quit() })
	ExitButton.Resize(fyne.NewSize(200, 50))
	ExitButton.Move(fyne.NewPos(200, 350))

	LogInButton := widget.NewButton("Log In", func() { LogIn() })
	LogInButton.Resize(fyne.NewSize(100, 100))
	LogInButton.Move(fyne.NewPos(0, 0))

	FindButton := widget.NewButton("Find game", func() {
		Player.Char = '0'
		StartGame()
		time.Sleep(100 * time.Millisecond)
		server.StartClient(g.Field)
		g.Field, g.Current = server.WaitMove(g.Field)
		DrawField()
	})
	FindButton.Resize(fyne.NewSize(200, 50))
	FindButton.Move(fyne.NewPos(200, 300))

	var labelUsername = widget.NewLabel(Player.Username)
	if Player.Username == "" {
		labelUsername = widget.NewLabel("Guest")
	}
	LabelUsername := container.NewCenter(labelUsername)

	Icon := widget.NewIcon(theme.AccountIcon())

	UserInfo := container.NewHBox(
		Icon,
		LabelUsername,
	)
	User := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 127, A: 255}),
		UserInfo,
	)
	User.Resize(fyne.NewSize(200, 100))
	User.Move(fyne.NewPos(400, 0))

	Menu.Add(User)
	Menu.Add(CreateButton)
	Menu.Add(FindButton)
	Menu.Add(ExitButton)
	Menu.Add(LogInButton)

	window.SetContent(&Menu)
}

func RunApp() {
	ShowMenu()
	window.SetFixedSize(true)
	ic, _ := fyne.LoadResourceFromPath("internal/user/OLEG.jpg")
	window.SetIcon(ic)

	window.ShowAndRun()
}

func LogIn() {
	var err error
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter your username")
	input.Resize(fyne.NewSize(350, 40))
	input.Move(fyne.NewPos(100, 200))

	ExitButton := widget.NewButton("Exit", func() { ShowMenu() })
	ExitButton.Resize(fyne.NewSize(200, 50))
	ExitButton.Move(fyne.NewPos(200, 300))

	SaveButton := widget.NewButton("Enter", func() {
		User := user.User{
			Username: input.Text,
		}
		Player.Username = User.Username
		if err = user.WriteUsers(User); err != nil {
			log.Println("Error while writing users file:", err)
		}
		ShowMenu()
	})
	SaveButton.Resize(fyne.NewSize(200, 50))
	SaveButton.Move(fyne.NewPos(200, 250))

	Logger.Add(SaveButton)
	Logger.Add(ExitButton)
	Logger.Add(input)
	window.SetContent(&Logger)

}
