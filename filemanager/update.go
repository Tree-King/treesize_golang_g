package filemanager

import (
	"sort"

	"9.suarha.com/root/tree_go.git/filesystemprotocol"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (fm *FileManager) updateFolderHierarchyTab(fs *filesystemprotocol.GFileSystem, tab *container.TabItem) {
	// Create a tree view of the folder hierarchy
	treeView := widget.NewTreeView(widget.NewTreeModel(
		&folderHierarchyTree{root: fs},
	))

	// Set the tree view as the content of the tab
	tab.Content = container.New(layout.NewVBoxLayout(), treeView)
}

func (fm *FileManager) updateSortBySizeTab(fs *filesystemprotocol.GFileSystem, tab *container.TabItem, sortOption int) {
	// Create a list view of the sorted files or folders based on the sort option
	var data []interface{}
	if sortOption == 0 {
		// Sort by file size
		for _, f := range fs.Files {
			data = append(data, f)
		}
	} else {
		// Sort by folder size
		for _, d := range fs.Directories {
			data = append(data, d)
		}
	}
	sort.Slice(data, func(i, j int) bool {
		var size1, size2 int64
		if sortOption == 0 {
			size1 = *data[i].(*filesystemprotocol.GFile).Size
			size2 = *data[j].(*filesystemprotocol.GFile).Size
		} else {
			size1 = *data[i].(*filesystemprotocol.GDirectory).Size
			size2 = *data[j].(*filesystemprotocol.GDirectory).Size
		}
		return size1 > size2
	})
	listView := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			if sortOption == 0 {
				file := data[i].(*filesystemprotocol.GFile)
				obj.(*widget.Label).SetText(file.Name + " - " + humanize.Bytes(uint64(*file.Size)))
			} else {
				dir := data[i].(*filesystemprotocol.GDirectory)
				obj.(*widget.Label).SetText(dir.Name + " - " + humanize.Bytes(uint64(*dir.Size)))
			}
		},
	)
	listView.Resize(fyne.NewSize(0, 500))

	// Set the list view as the content of the tab
	tab.Content = container.New(layout.NewVBoxLayout(), widget.NewSelect([]string{"File", "Folder"}, func(selected string) {
		if selected == "File" {
			fm.updateSortBySizeTab(fs, tab, 0)
		} else {
			fm.updateSortBySizeTab(fs, tab, 1)
		}
	}), listView)
}
