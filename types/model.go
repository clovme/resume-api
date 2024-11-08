package types

type ModelItem struct {
	Name  string
	Model interface{}
}

type FileInfo struct {
	Size int64 `json:"size"` // 文件大小
}
