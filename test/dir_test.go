package test

import (
	. "github.com/soofjan1234/sortFile"
	"testing"
)

func TestDirName(t *testing.T) {
	// 创建测试用的文件切片
	dirs := []Dir{
		{Name: "a"},
		{Name: "file20"},
		{Name: "file3"},
		{Name: "3"},
		{Name: "9"},
		{Name: "_hhh"},
		{Name: "_a"},
		{Name: "%file3"},
	}

	// 调用 SortDirByName 函数进行排序
	SortByDirName(dirs)

	// 检查排序后的结果是否正确
	t.Log("After sorting dirs :")
	for _, file := range dirs {
		t.Log(file.Name)
	}

	// 调用 SortDirByNameDesc 函数进行排序
	SortByDirNameDesc(dirs)

	// 检查排序后的结果是否正确
	t.Log("After sorting dirs desc:")
	for _, file := range dirs {
		t.Log(file.Name)
	}
}
