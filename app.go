package main

import (
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMouseClick() int {
	return 1
}

func (a *App) GreetAsyncViaEvent() {
	count := 0
	go func() {
		ch := robotgo.EventStart()
		for event := range ch {
			if event.Kind == hook.MouseUp {
				count++
				runtime.EventsEmit(a.ctx, "rcv:greet", count)
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (a *App) OpenFileDialog() error {
	p, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{Title: "选取lol路径"})
	fmt.Println("OpenFileDialog", p, err)
	return err
}
