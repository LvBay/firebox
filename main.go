package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "firebox",
		Width:     1110,
		Height:    768,
		Assets:    assets,
		Frameless: false,
		//     background-color: rgb(73,64,106);
		BackgroundColour: options.NewRGB(73, 64, 106),
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			Appearance:           "",
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About:                nil,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
