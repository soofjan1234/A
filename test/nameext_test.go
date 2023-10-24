package test

import (
	. "github.com/soofjan1234/sortFile"
	"testing"
)

func TestNameExtension(t *testing.T) {
	// 创建测试用的文件切片
	files := []File{
		{Name: "a.txt"},
		{Name: "file20.txt"},
		{Name: "file3.doc"},
		{Name: "3.docx"},
		{Name: "9.mp3"},
		{Name: "_hhh"},
		{Name: "_a.jpg"},
	}

	// 调用 SortFileByNameExtension 函数进行排序
	SortByNameExtension(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files :")
	for _, file := range files {
		t.Log(file.Name)
	}

	// 调用 SortFileByNameExtensionDesc 函数进行排序
	SortByNameExtensionDesc(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files desc:")
	for _, file := range files {
		t.Log(file.Name)
	}
}
