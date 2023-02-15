package filemanager

import (
	"9.suarha.com/root/tree_go.git/filesystemprotocol"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ScanFunc func(path string) *filesystemprotocol.GFileSystem

type FileManager struct {
	window   fyne.Window
	scanFunc ScanFunc
}

func New() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) SetScanFunc(scanFunc ScanFunc) {
	fm.scanFunc = scanFunc
}

func (fm *FileManager) Run() {
	a := app.New()
	fm.window = a.NewWindow("File Manager")
	folderEntry := widget.NewEntry()
	folderEntry.Disable()
	folderButton := widget.NewButton("Select Folder", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				folderEntry.SetText(uri.Path())
			}
		}, fm.window)
	})
	tabContent := container.NewAppTabs(
		container.NewTabItem("Folder Hierarchy", container.New(layout.NewVBoxLayout(), widget.NewLabel("Folder Hierarchy"))),
		container.NewTabItem("Sort by Size", container.New(layout.NewVBoxLayout(), widget.NewSelect([]string{"File", "Folder"}, nil))),
	)
	scanButton := widget.NewButton("Scan", nil)
	fm.window.SetContent(container.New(layout.NewVBoxLayout(),
		container.New(layout.NewHBoxLayout(),
			widget.NewLabel("Folder:"),
			folderEntry,
			folderButton,
			scanButton,
		),
		tabContent,
	))

	scanButton.OnTapped = func() {
		path := folderEntry.Text
		results := fm.scanFunc(path)
		_ = results
		// TODO: Add logic to update the folder hierarchy and sort by size tabs with the scanned data
	}

	fm.window.ShowAndRun()
}
