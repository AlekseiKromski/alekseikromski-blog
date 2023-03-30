package models

import "time"

type Tag struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	*Timestamp
	*SoftDelete
}

func (t *Tag) Validate() bool {
	if len(t.Name) != 0 {
		return true
	}
	return false
}

func (t *Tag) SetTimestamt() {
	t.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(t.CreatedAt) != 0 {
		t.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (t *Tag) Soft() {
	t.DeletedAt = time.Now().Format(time.RFC3339)
}

func (t *Tag) Undo() {
	t.DeletedAt = ""
}
