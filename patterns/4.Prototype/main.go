package main

import "fmt"

func main() {
	file1 := &File{"File1"}
	file2 := &File{"File2"}
	file3 := &File{"File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{file2, file3},
		name:     "Folder2",
	}

	fmt.Println("Print hierarhy for Folder1")
	folder1.print("  ")
	fmt.Println("Print hierarhy for clone Folder2")
	folder2.clone().print("  ")
}
