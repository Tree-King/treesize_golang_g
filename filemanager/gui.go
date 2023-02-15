package guifiles

import (
	"fmt"

	"9.suarha.com/root/tree_go.git/filesystemprotocol"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

// GUI is a struct that contains the GUI elements and data for the file manager
type GUI struct {
	window       fyne.Window
	treeView     *widget.Tree
	currentPath  string
	scanFunction ScanFunction
}

// NewGUI creates a new GUI for the file manager
func NewGUI() *GUI {
	return &GUI{}
}

// SetScanFunction sets the function used to scan folders
func (gui *GUI) SetScanFunction(scanFunction ScanFunction) {
	gui.scanFunction = scanFunction
}

// Run starts and runs the GUI
func (gui *GUI) Run() {
	a := app.New()
	gui.window = a.NewWindow("TreeSize")

	// Create the folder selection elements
	folderEntry := widget.NewEntry()
	folderEntry.Disable()
	folderButton := widget.NewButtonWithIcon("Select Folder", theme.FolderOpenIcon(), func() {
		folderDialog := dialog.Directory()
		folderDialog.Title = "Select Folder"
		folderDialog.Show()
		if folderDialog.Selected() {
			folderEntry.SetText(folderDialog.Directory())
			gui.currentPath = folderDialog.Directory()
		}
	})

	// Create the scan button
	scanButton := widget.NewButtonWithIcon("Scan", theme.SearchIcon(), func() {
		if gui.currentPath == "" {
			dialog.ShowInformation("Error", "Please select a folder to scan", gui.window)
			return
		}
		results := gui.scanFunction(gui.currentPath)
		gui.updateTreeView(results)
	})

	// Create the tree view and add it to a scrollable container
	gui.treeView = widget.NewTree(gui.createTreeNode, nil)
	treeContainer := container.NewScroll(gui.treeView)

	// Add the folder selection elements, scan button, and tree view to the window
	gui.window.SetContent(container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Folder:"),
			folderEntry,
			folderButton,
			scanButton,
		),
		treeContainer,
	))

	gui.window.Resize(fyne.NewSize(800, 600))
	gui.window.ShowAndRun()
}

// createTreeNode creates a widget.TreeNode for a GFileSystem object
func (gui *GUI) createTreeNode(fs *filesystemprotocol.GFileSystem) *widget.TreeNode {
	if fs == nil {
		return nil
	}

	node := &widget.TreeNode{
		Value:    fs.G名称,
		IsBranch: fs.G是目录,
	}

	if fs.G是目录 {
		for _, child := range fs.G子文件系统 {
			node.Children = append(node.Children, gui.createTreeNode(child))
		}
	} else {
		sizeString := fmt.Sprintf("%.2f MB", float64(fs.G大小)/1024/1024)
		node.Value = fmt.Sprintf("%s (%s)", fs.G名称, sizeString)
	}

	return node
}
