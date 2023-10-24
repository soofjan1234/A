package test

import (
	. "github.com/soofjan1234/sortFile"
	"testing"
	"time"
)

func TestUpdatedTime(t *testing.T) {
	// 创建测试用的文件切片
	files := []File{
		{Name: "file1", ModifiedAt: time.Now().Add(-time.Hour)},
		{Name: "file2", ModifiedAt: time.Now()},
		{Name: "file3", ModifiedAt: time.Now().Add(-2 * time.Hour)},
	}

	// 调用 SortFileByUpdatedTime 函数进行排序
	SortByUpdatedTime(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files :")
	for _, file := range files {
		t.Log(file.ModifiedAt)
	}

	// 调用 SortFileByUpdatedTimeDesc 函数进行排序
	SortByUpdatedTimeDesc(files)

	// 检查排序后的结果是否正确
	t.Log("After sorting files desc:")
	for _, file := range files {
		t.Log(file.ModifiedAt)
	}
}
