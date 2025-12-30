package cmd

import (
	"dpswitch/cmd/display-tools"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func InitMenu(app desktop.App) *fyne.Menu {
	var menuItems []*fyne.MenuItem
	if isKDESession() && hasKScreenDoctor() {
		menuItems = SetupMenu(app.(desktop.App), &display_tools.KScreenDoctor{})
	} else {
		menuItems = []*fyne.MenuItem{
			fyne.NewMenuItem("Your current desktop environment is not supported", nil),
		}
	}

	menu := fyne.NewMenu("dpswitch", menuItems...)
	app.SetSystemTrayMenu(menu)
	return menu
}

func SetupMenu(app desktop.App, tool display_tools.DPConfigTool) []*fyne.MenuItem {
	displayList, err := tool.ListDisplays()
	if err != nil {
		fyne.LogError("Couldn't get screens", err)
		return nil
	}

	var items []*fyne.MenuItem
	for _, display := range displayList {
		if display.Connected {
			item := fyne.NewMenuItem(display.ID, func() {
				err := tool.ToggleDisplay(display.ID)
				if err != nil {
					log.Println(err)
					return
				}
				InitMenu(app)
			})
			item.Checked = display.Active
			if display.Primary {
				item.Disabled = true
			}
			items = append(items, item)
		}
	}
	return items
}

func isKDESession() bool {
	// KDE sets this
	if os.Getenv("XDG_CURRENT_DESKTOP") == "KDE" {
		return true
	}

	// Plasma also sets this
	if os.Getenv("DESKTOP_SESSION") == "plasma" {
		return true
	}

	return false
}

func hasKScreenDoctor() bool {
	_, err := exec.LookPath("kscreen-doctor")
	return err == nil
}
