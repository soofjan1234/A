package main

import (
	"fmt"
	"github.com/soofjan1234/sortFile"
	"sort"
)

func SortFilesByName(fileNames []string) {
	sort.Slice(fileNames, func(i, j int) bool {
		return sortFile.CompareFileNames(fileNames[i], fileNames[j])
	})
}

func SortFilesByNameExtension(fileNames []string) {
	sort.Slice(fileNames, func(i, j int) bool {
		ext1 := sortFile.GetFileExtension(fileNames[i])
		ext2 := sortFile.GetFileExtension(fileNames[j])
		return sortFile.CompareFileNames(ext1, ext2)
	})
}

func main() {
	fileNames := []string{
		"file2.txt",
		"a.doc",
		"100.tar.gz",
		"20.jpg",
		"File3.png",
		"file4.docx",
		"file30.mp3",
		"_file.mp4",
		"#file2.zip",
		"***file_1.exe",
		"",
	}

	SortFilesByNameExtension(fileNames)

	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}
}
