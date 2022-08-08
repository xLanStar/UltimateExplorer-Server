package fileManager

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	rootFolders   []Folder           = make([]Folder, 0, 10) // 實際 Folders 儲存地
	mapFolder     map[uint32]*Folder = make(map[uint32]*Folder)
	mapFile       map[uint32]*File   = make(map[uint32]*File, 1000)
	folderCounter uint32             = 0
	fileCounter   uint32             = 0

	folderPaths []string
)

// Private 函數
func loadFolder(folderPath string) Folder {

	fmt.Println("[FileManager] 讀取資料夾 路徑:", folderPath)

	// 建立 Folder
	node := Folder{
		Id:         folderCounter,
		FolderName: filepath.Base(folderPath),
		FolderPath: folderPath,
		SubFolders: make([]Folder, 0, 5),
		Files:      make([]File, 0, 20),
	}
	mapFolder[node.Id] = &node
	folderCounter++

	// 讀取 Folder
	f, _ := os.Open(folderPath)
	fileInfos, _ := f.Readdir(-1)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// 讀取 SubFolder
			node.SubFolders = append(node.SubFolders, loadFolder(filepath.Join(folderPath, fileInfo.Name())))
		} else {
			// 讀取 File
			file := File{
				Id:       fileCounter,
				FileName: fileInfo.Name(),
				FilePath: filepath.Join(folderPath, fileInfo.Name()),
				Type:     GetType(fileInfo.Name()),
			}
			node.Files = append(node.Files, file)
			mapFile[file.Id] = &file
			fileCounter++
		}
	}

	return node
}

// public 函數
func Init(_folderPaths []string, dbpath string) {
	folderPaths = _folderPaths[:]

	for _, folderPath := range folderPaths {
		rootFolders = append(rootFolders, loadFolder(folderPath))
	}
}

func GetFile(fileId uint32) *File {
	return mapFile[fileId]
}

func GetRootFolders() []Folder {
	return rootFolders
}

func GetFolderFiles(folderId uint32) []File {
	return mapFolder[folderId].Files
}
