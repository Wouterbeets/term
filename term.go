package term

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	Input     chan [][]rune
	UserInput chan rune
}

func NewScreen() *Screen {
	return &Screen{
		Input:     make(chan [][]rune),
		UserInput: make(chan rune),
	}
}

func (s Screen) Run(frameRate time.Duration) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()
	go func() {
		for {
			e := termbox.PollEvent()
			fmt.Println(e)
			s.UserInput <- e.Ch
		}
	}()
	for {
		start := time.Now()
		select {
		case inp := <-s.Input:
			if len(inp) == 0 {
				return
			}
			display(inp)
		case <-time.After(frameRate):
		}
		end := time.Now()
		elapsed := end.Sub(start)
		time.Sleep(frameRate - elapsed)
	}
}

func display(inp [][]rune) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlue)
	for y, row := range inp {
		for x, cell := range row {
			termbox.SetCell(x*2, y, rune(cell), termbox.ColorBlue, termbox.ColorWhite)
			termbox.SetCell(x*2-1, y, rune(cell), termbox.ColorBlue, termbox.ColorWhite)
		}
	}
	termbox.Flush()
}
