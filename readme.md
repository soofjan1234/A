# sortFile
sortFile is a method that sorts files by filename or file extension.

## Usage
```txt
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
	//SortFilesByName(fileNames)

	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}
}
```
