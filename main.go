package main

import (
	"9.suarha.com/root/tree_go.git/filesystemprotocol"
	"9.suarha.com/root/tree_go.git/scanFolder"
)

func main() {
	// Create a new file manager and set the scan function to use the ScanFolder function from the scanFolder package
	fm := filemanager.New()
	fm.SetScanFunc(func(path string) *filesystemprotocol.GFileSystem {
		fs, err := scanFolder.G获取目录结构(path)
		if err != nil {
			return nil
		}

		return fs
	})

	// Run the file manager
	fm.Run()
}
