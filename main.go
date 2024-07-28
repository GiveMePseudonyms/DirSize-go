package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Directory struct {
	Name string
	Size int64
}

func GetDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

func ListDirectories(root string, depth int) ([]Directory, error) {
	var dirs []Directory
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			if len(filepath.SplitList(relPath)) <= depth {
				size, err := GetDirSize(path)
				if err != nil {
					return err
				}
				dirs = append(dirs, Directory{Name: path, Size: size})
			}
		}
		return nil
	})
	return dirs, err
}

func main() {
	var root string
	var depth int

	fmt.Print("Enter the directory to search: ")
	fmt.Scan(&root)
	fmt.Print("Enter the search depth: ")
	fmt.Scan(&depth)

	dirs, err := ListDirectories(root, depth)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Size < dirs[j].Size
	})

	fmt.Println("Directories sorted by size:")
	for _, dir := range dirs {
		fmt.Printf("%s: %d bytes\n", dir.Name, dir.Size)
	}
}
