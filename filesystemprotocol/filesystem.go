package filesystemprotocol

// GFileSystem 结构体表示文件系统中的文件和文件夹
type GFileSystem struct {
	G名称    string
	G大小    int64
	G是目录   bool
	G子文件系统 []*GFileSystem
}
