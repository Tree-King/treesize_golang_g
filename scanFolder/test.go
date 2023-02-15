package scanFolder

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

const (
	文件数量   = 10
	目录层数   = 5
	最大文件大小 = 1024 * 1024 * 3  // 3MB
	最大目录大小 = 1024 * 1024 * 50 // 50MB
)

func Test获取目录结构(t *testing.T) {
	// 创建一个临时目录，用于测试
	临时目录, err := ioutil.TempDir("", "scanFolder-test-")
	if err != nil {
		t.Fatalf("无法创建临时目录：%v", err)
	}
	defer os.RemoveAll(临时目录)

	// 在临时目录中创建一些文件和目录
	创建测试文件和目录(临时目录, "", 0)

	// 测试获取目录结构
	根, err := G获取目录结构(临时目录)
	if err != nil {
		t.Fatalf("G获取目录结构 函数出错：%v", err)
	}

	// 测试按文件夹大小排序输出
	结果 := G按文件夹排序输出(根)
	if len(结果) == 0 {
		t.Fatalf("G按文件夹大小排序输出 函数返回空列表")
	}
	for i := 1; i < len(结果); i++ {
		if 结果[i-1].G大小 < 结果[i].G大小 {
			t.Errorf("G按文件夹大小排序输出 函数结果排序错误")
			break
		}
	}

	// 测试按文件大小排序输出
	结果 = G按文件大小排序输出(根)
	if len(结果) == 0 {
		t.Fatalf("G按文件大小排序输出 函数返回空列表")
	}
	for i := 1; i < len(结果); i++ {
		if 结果[i-1].G大小 < 结果[i].G大小 {
			t.Errorf("G按文件大小排序输出 函数结果排序错误")
			break
		}
	}
}

// 创建测试文件和目录
func 创建测试文件和目录(目录, 前缀路径 string, 层数 int) {
	// 创建一些文件和目录
	文件和目录, _ := ioutil.ReadDir(目录)
	文件数量 := len(文件和目录)

	for i := 文件数量; i < 文件数量+rand.Intn(文件数量); i++ {
		if 层数 >= 目录层数 {
			return
		}

		名称 := fmt.Sprintf("测试文件-%d", i)
		if rand.Intn(2) == 1 { // 模拟目录
			err := os.Mkdir(filepath.Join(目录, 名称), 0700)
			if err != nil {
				return
			}

			// 递归创建目录
			if 层数 < 目录层数-1 {
				创建测试文件和目录(filepath.Join(目录, 名称), filepath.Join(前缀路径, 名称), 层数+1)
			}
		} else { // 模拟文件
			大小 := int64((math.Sin(float64(i)/float64(文件数量)*2*math.Pi)+1)/2*float64(最大文件大小)) + rand.Int63n(int64(最大文件大小/10))

			文件, err := os.Create(filepath.Join(目录, 名称))
			if err != nil {
				return
			}
			defer 文件.Close()

			数据 := make([]byte, 大小)
			_, err = rand.Read(数据)
			if err != nil {
				return
			}

			_, err = 文件.Write(数据)
			if err != nil {
				return
			}
		}
	}

	// 更新目录大小
	文件和目录, _ = ioutil.ReadDir(目录)
	大小 := int64(0)
	for _, 文件 := range 文件和目录 {
		if 文件.IsDir() {
			创建测试文件和目录(filepath.Join(目录, 文件.Name()), filepath.Join(前缀路径, 文件.Name()), 层数+1)
		} else {
			大小 += 文件.Size()
		}
	}
	if 大小 < 最大目录大小 {
		残留大小 := 最大目录大小 - 大小
		垃圾数据 := make([]byte, 残留大小)
		_, _ = rand.Read(垃圾数据)

		垃圾文件, _ := os.Create(filepath.Join(目录, fmt.Sprintf("测试垃圾文件-%d", 文件数量)))
		垃圾文件.Write(垃圾数据)
		垃圾文件.Close()
	}
}
