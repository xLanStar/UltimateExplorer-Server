package fileManager

type Folder struct {
	Id         uint32   `json:"id"`
	FolderName string   `json:"name"`
	FolderPath string   `json:"-"`
	Files      []File   `json:"-"`
	SubFolders []Folder `json:"subfolders,omitempty"`
}
