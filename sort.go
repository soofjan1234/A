package sortFile

import "sort"

func SortByUpdatedTime(files []File) {
	sort.Slice(files, func(i, j int) bool {
		return CompareFileUpdatedTime(files[i].ModifiedAt, files[j].ModifiedAt)
	})
}

func SortBySize(files []File) {
	sort.Slice(files, func(i, j int) bool {
		return CompareFileSize(files[i].Size, files[j].Size)
	})
}

func SortByName(files []File) {
	sort.Slice(files, func(i, j int) bool {
		return CompareFileNames(files[i].Name, files[j].Name)
	})
}

func SortByDirName(dirs []Dir) {
	sort.Slice(dirs, func(i, j int) bool {
		return CompareFileNames(dirs[i].Name, dirs[j].Name)
	})
}

func SortByNameExtension(files []File) {
	sort.Slice(files, func(i, j int) bool {
		ext1 := GetFileExtension(files[i].Name)
		ext2 := GetFileExtension(files[j].Name)
		return CompareFileNames(ext1, ext2)
	})
}

func SortByUpdatedTimeDesc(files []File) {
	SortByUpdatedTime(files)
	reverseFile(files)
}

func SortBySizeDesc(files []File) {
	SortBySize(files)
	reverseFile(files)
}

func SortByNameDesc(files []File) {
	SortByName(files)
	reverseFile(files)
}

func SortByDirNameDesc(dirs []Dir) {
	SortByDirName(dirs)
	reverseDir(dirs)
}

func SortByNameExtensionDesc(files []File) {
	SortByNameExtension(files)
	reverseFile(files)
}

func reverseFile(arr []File) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
}

func reverseDir(dirs []Dir) {
	length := len(dirs)
	for i := 0; i < length/2; i++ {
		dirs[i], dirs[length-i-1] = dirs[length-i-1], dirs[i]
	}
}
