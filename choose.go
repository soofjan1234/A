package sortFile

func chooseAndSortDirMethod(sort string, ascend bool, dirs []Dir) {
	if sort == FileName {
		if ascend {
			SortByDirName(dirs)
		} else {
			SortByDirNameDesc(dirs)
		}
	}
}

func chooseAndSortFileMethod(sort string, ascend bool, files []File) {
	if sort == UpdateTime {
		if ascend {
			SortByUpdatedTime(files)
		} else {
			SortByUpdatedTimeDesc(files)
		}
	} else if sort == FileSize {
		if ascend {
			SortBySize(files)
		} else {
			SortBySizeDesc(files)
		}
	} else if sort == FileName {
		if ascend {
			SortByName(files)
		} else {
			SortByNameDesc(files)
		}
	} else if sort == FileCategory {
		if ascend {
			SortByNameExtension(files)
		} else {
			SortByNameExtensionDesc(files)
		}
	}
}
