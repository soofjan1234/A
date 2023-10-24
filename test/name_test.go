package test

import (
	. "github.com/soofjan1234/sortFile"
	"testing"
)

func TestName(t *testing.T) {
	// 创建测试用的文件切片
	files := []File{
		{Name: "a"},
		{Name: "file20"},
		{Name: "file3"},
		{Name: "3"},
		{Name: "9"},
		{Name: "_hhh"},
		{Name: "_a"},
		{Name: "%file3"},
	}

	// 调用 SortFileByName 函数进行排序
	SortByName(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files :")
	for _, file := range files {
		t.Log(file.Name)
	}

	// 调用 SortFileByNameDesc 函数进行排序
	SortByNameDesc(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files desc:")
	for _, file := range files {
		t.Log(file.Name)
	}
}
