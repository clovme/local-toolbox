package article

type ApiMdImagesVO struct {
	Url  string `json:"url"`
	Name string `json:"alt"`
	Type string `json:"type"`
	Size int64  `json:"size"`
	Hash string `json:"hash"`
}
