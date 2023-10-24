package sortFile

import (
	"gorm.io/gorm"
	"time"
)

const (
	UpdateTime   = "UpdateTime"
	FileName     = "FileName"
	FileCategory = "FileCategory"
	FileSize     = "FileSize"
)

type File struct {
	gorm.Model
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	ModifiedAt time.Time `json:"modified_at"`
	Size       int64     `json:"size"`
	Type       string    `json:"type"`
}

type Dir struct {
	gorm.Model
	Name string `json:"name"`
	Path string `json:"path"`
}
