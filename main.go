package main

import (
	"context"
	"embed"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	JARUPER_SETTINGS_FILE_PATH string
	TMP_SETTINGS_FILE_PATH     string
	REPOSITORY_LIST_FILE_PATH  string
	TMP_DIR_PATH               string
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

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}
	JARUPER_SETTINGS_FILE_PATH = filepath.Join(currentDir, "config", "libsynchro-settings.xml")
	TMP_SETTINGS_FILE_PATH = filepath.Join(currentDir, "config", "tmp", "settings.xml")
	REPOSITORY_LIST_FILE_PATH = filepath.Join(currentDir, "config", "repository-list.json")
	TMP_DIR_PATH = filepath.Join(currentDir, "config", "tmp")
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "libsynchro",
		Width:  1400,
		Height: 1100,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
