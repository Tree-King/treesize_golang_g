package scanFolder

import (
	"os"
	"path/filepath"

	"9.suarha.com/root/tree_go.git/filesystemprotocol"
)

// 遍历目录下一切文件与文件夹，返回一个按层级结构存储的filesystemprotocol.GFileSystem对象
func G获取目录结构(根路径 string) (*filesystemprotocol.GFileSystem, error) {
	根 := &filesystemprotocol.GFileSystem{}

	文件夹, err := os.Open(根路径)
	if err != nil {
		return 根, err
	}
	defer 文件夹.Close()

	根.G名称 = filepath.Base(根路径)
	根.G是目录 = true
	根.G大小 = 0

	文件信息, err := 文件夹.Readdir(0)
	if err != nil {
		return 根, err
	}

	for _, 文件 := range 文件信息 {
		子文件系统 := &filesystemprotocol.GFileSystem{
			G名称:    文件.Name(),
			G大小:    文件.Size(),
			G是目录:   文件.IsDir(),
			G子文件系统: nil,
		}

		if 文件.IsDir() {
			子文件系统, err = G获取目录结构(filepath.Join(根路径, 文件.Name()))
			if err != nil {
				return 根, err
			}
		}

		根.G子文件系统 = append(根.G子文件系统, 子文件系统)
		根.G大小 += 子文件系统.G大小
	}

	return 根, nil
}
