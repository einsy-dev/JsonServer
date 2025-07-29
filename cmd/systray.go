package main

import (
	_ "embed"
	"fmt"
	"time"

	"fyne.io/systray"
)

func SysTray() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}

	systray.Run(onReady, onExit)
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

//go:embed icon.ico
var icon []byte

func onReady() {
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("Local Server")
	systray.SetTooltip("Local Server")

	clear := systray.AddMenuItem("Clear", "Clear Saved Data")
	systray.AddSeparator()

	restart := systray.AddMenuItem("Restart", "Restart Server")
	systray.AddSeparator()

	addQuitItem()

	go func() {

		for {
			select {
			case <-clear.ClickedCh:
				clear.SetTitle("Cleared")

			case <-restart.ClickedCh:
				systray.ResetMenu()
				addQuitItem()

			}
		}
	}()
}
