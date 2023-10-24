# sortFile
sortFile is a method that sorts files by filename or file extension.
And Also can sort directories

## Methods
```txt
func chooseAndSortDirMethod(sort string, ascend bool, dirs []Dir) {}
func chooseAndSortFileMethod(sort string, ascend bool, files []File) {}
```

## Notes
Here are the explanation of the arguments
```txt
// sort所应填写的参数
const (
	UpdateTime = "UpdateTime"
	FileName = "FileName"
	FileCategory = "FileCategory"
	FileSize = "FileSize"
)

// 应该转成的文件模型
type File struct {
	gorm.Model
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	ModifiedAt time.Time `json:"modified_at"`
	Size       int64     `json:"size"`
	Type       string    `json:"type"`
}

// 应该转成的目录模型
type Dir struct {
	gorm.Model
	Name       string    `json:"name"`
	Path       string    `json:"path"`
}
```

Directories can only be sorted by name.

