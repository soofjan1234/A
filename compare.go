package sortFile

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func CompareFileNames(a, b string) bool {
	// 将文件名转换为小写进行比较
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	// 逐个字符比较
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		// 获取字符的ASCII值
		asciiA := int(a[i])
		asciiB := int(b[j])

		// 判断是否为数字
		isDigitA := isDigit(a[i])
		isDigitB := isDigit(b[j])

		if isDigitA && isDigitB {
			// 如果都是数字，则将连续的数字字符转换为整数进行比较
			numA := extractNumber(a, i)
			numB := extractNumber(b, j)

			if numA != numB {
				// 数字不相等，按数值大小进行比较
				return numA < numB
			}
		} else if isDigitA {
			// 数字优先于字母和特殊字符
			return false
		} else if isDigitB {
			// 数字优先于字母和特殊字符
			return true
		} else {
			// 两个字符都是字母或特殊字符，直接按照ASCII值比较
			if asciiA != asciiB {
				return asciiA < asciiB
			}
			// 如果相等，则比较下一个字符
			i++
			j++
		}
	}

	// 如果一个文件名是另一个文件名的前缀，则较短的文件名较小
	return len(a) < len(b)
}

func CompareFileUpdatedTime(a, b time.Time) bool {
	return a.After(b)
}

func CompareFileSize(a, b int64) bool {
	return a < b
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func extractNumber(str string, i int) int {
	start := i
	for i < len(str) && isDigit(str[i]) {
		i++
	}
	num, _ := strconv.Atoi(str[start:i])
	return num
}

func GetFileExtension(str string) string {
	fileExtension := filepath.Ext(str)
	fileExtension = strings.TrimPrefix(fileExtension, ".")
	return fileExtension
}
