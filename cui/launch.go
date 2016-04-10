package hncui

import (
	"github.com/jroimartin/gocui"
)

// Launch is used to begin the command line user interface portion of the
// application.
func Launch() error {
	g := gocui.NewGui()
	err := g.Init()
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetLayout(layout)
	err = keybindings(g)
	if err != nil {
		return err
	}
	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
