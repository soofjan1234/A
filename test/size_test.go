package test

import (
	. "github.com/soofjan1234/sortFile"
	"testing"
)

func TestSize(t *testing.T) {
	// 创建测试用的文件切片
	files := []File{
		{Name: "file1", Size: 5000},
		{Name: "file2", Size: 500},
		{Name: "file3", Size: 50000},
	}

	// 调用 SortFileBySize 函数进行排序
	SortBySize(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files :")
	for _, file := range files {
		t.Log(file.Size)
	}

	// 调用 SortFileBySizeDesc 函数进行排序
	SortBySizeDesc(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files desc:")
	for _, file := range files {
		t.Log(file.Size)
	}
}
