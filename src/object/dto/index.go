package dto

type PageBasicReq struct {
	PageSize int    `json:"pageSize,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	Type     int    `json:"type,omitempty"`
	Id       int    `json:"id,omitempty"`
}
