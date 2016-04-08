package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/kkirsche/hn/api"
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	// if err := g.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, saveMain); err != nil {
	// 	return err
	// }
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("main", 0, 0, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// User section
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter Hacker News username you would like to retrieve: ")
		// username, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Panicln(err.Error())
		// }
		//
		// user := hnapi.GetUser(username)
		// fmt.Println(user.Karma)

		// Top Stories section
		topStores := hnapi.RetrieveTopStoriesItemNumbers()
		for i, tsNumber := range *topStores {
			if i > 25 {
				break
			}
			item := hnapi.GetItem(tsNumber)

			fmt.Fprintf(v, "%s\n\n", item.Title)
		}

		v.Editable = false
		v.Wrap = false
		if err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layout)
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}
	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
