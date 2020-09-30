package model

type Quote struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Path    string `json:"path,omitempty"`
}
