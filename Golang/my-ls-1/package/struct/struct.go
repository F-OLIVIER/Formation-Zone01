package Struct

import (
	"time"
)

type Option struct {
	OptionPetitA bool
	OptionPetitL bool
	OptionGrandR bool
	OptionPetitR bool
	OptionPetitT bool
	Argument     []string
}

type FileItem struct {
	OriginalName string
	Name         string
	Linkname	 string
	Link         int
	Permission   string
	User         string
	Group        string
	Size         int
	Lastmod      time.Time
	Minor        uint32
	Major        uint32
	Ftype        string
	FolderPath   string
	FolderPathT  string
	Total        int
}
