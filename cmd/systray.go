package main

import (
	_ "embed"
	"fmt"
	"time"

	"fyne.io/systray"
)

//go:embed icon.ico
var icon []byte

func SysTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Local Server")
	systray.SetTooltip("Local Server")
	systray.SetTemplateIcon(icon, icon)

	clear := systray.AddMenuItem("Clear", "Clear Saved Data")
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()

	go func() {
		for {
			select {
			case <-clear.ClickedCh:
				ClearServer()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	now := time.Now()
	fmt.Println("Exit at", now.String())
}
