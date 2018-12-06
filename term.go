package term

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	Input chan [][]rune
}

func (s Screen) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()
	for {
		frameRate := time.Duration(100 * time.Millisecond)
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
			termbox.SetCell(x, y, rune(cell), termbox.ColorBlue, termbox.ColorWhite)
		}
	}
	termbox.Flush()
}
