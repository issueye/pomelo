package tray

import (
	"context"
	_ "embed"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icon/icon.ico
var icon []byte

type Tray struct {
	// app *backend.App
}

func onReady(ctx context.Context) func() {
	return func() {
		systray.SetIcon(icon)
		systray.SetTitle("一个小工具集")

		mMain := systray.AddMenuItem("Pomelo", "一个小工具")
		mQuit := systray.AddMenuItem("退出", "退出程序")

		go func() {
			for {
				select {
				case <-mQuit.ClickedCh:
					{
						systray.Quit()
						runtime.Quit(ctx)
					}
				case <-mMain.ClickedCh:
					{
						runtime.Show(ctx)
					}
				}
			}
		}()
	}
}

func onExit(_ context.Context) func() {
	return func() {}
}

func RunTray(ctx context.Context) {
	systray.Run(onReady(ctx), onExit(ctx))
}
