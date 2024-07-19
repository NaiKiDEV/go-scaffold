package config

type File struct {
	Name     string
	Template string
}

type Directory struct {
	Name           string
	Files          []File
	SubDirectories []Directory
}

type Config struct {
	DirectoryConfig []Directory
}
