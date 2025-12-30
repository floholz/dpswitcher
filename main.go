package main

import (
	_ "embed"
	"time"

	"github.com/floholz/dpswitch/cmd"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

//go:embed assets/icon.png
var iconData []byte

func main() {
	a := app.NewWithID("dpswitch")
	if desk, ok := a.(desktop.App); ok {
		icon := fyne.NewStaticResource("icon", iconData)
		desk.SetSystemTrayIcon(icon)
		cmd.InitMenu(desk)
		startMenuPolling(desk)
	}
	a.Run()
}

func startMenuPolling(app desktop.App) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			// This will update the tray menu every 5 seconds
			cmd.InitMenu(app)
		}
	}()
}
