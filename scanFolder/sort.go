package scanFolder

import (
	"path/filepath"
	"sort"

	"9.suarha.com/root/tree_go.git/filesystemprotocol"
)

// GBySize 结构体表示按大小排序的结构体
type GBySize []*filesystemprotocol.GFileSystem

func (a GBySize) Len() int           { return len(a) }
func (a GBySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a GBySize) Less(i, j int) bool { return a[i].G大小 > a[j].G大小 }

// 按文件夹大小排序输出，返回一个按照文件夹大小排序的列表，列表中包含每个文件夹的名称、大小和绝对路径
func G按文件夹排序输出(根 *filesystemprotocol.GFileSystem) []*filesystemprotocol.GFileSystem {
	var 结果 []*filesystemprotocol.GFileSystem
	var 前缀路径 string

	for _, 子文件系统 := range 根.G子文件系统 {
		if 子文件系统.G是目录 {
			结果 = append(结果, G按文件夹排序输出(子文件系统)...)
		} else {
			子文件系统.G名称 = filepath.Join(前缀路径, 子文件系统.G名称)
			结果 = append(结果, 子文件系统)
		}
	}

	sort.Sort(GBySize(结果))
	for i := range 结果 {
		结果[i].G名称 = filepath.Join(前缀路径, 结果[i].G名称)
	}
	return 结果
}

// 按文件大小排序输出，返回一个按照文件大小排序的列表，列表中包含每个文件的名称、大小和绝对路径
func G按文件大小排序输出(根 *filesystemprotocol.GFileSystem) []*filesystemprotocol.GFileSystem {
	var 结果 []*filesystemprotocol.GFileSystem
	var 前缀路径 string

	for _, 子文件系统 := range 根.G子文件系统 {
		if 子文件系统.G是目录 {
			结果 = append(结果, G按文件大小排序输出(子文件系统)...)
		} else {
			子文件系统.G名称 = filepath.Join(前缀路径, 子文件系统.G名称)
			结果 = append(结果, 子文件系统)
		}
	}

	sort.Sort(GBySize(结果))
	for i := range 结果 {
		结果[i].G名称 = filepath.Join(前缀路径, 结果[i].G名称)
	}
	return 结果
}
