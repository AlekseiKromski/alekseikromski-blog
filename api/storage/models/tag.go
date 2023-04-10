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

func (t *Tag) SetTimestamp() {
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

func (m *Tag) TableCreate() string {
	return `
		create table tags
		(
			ID          serial
				constraint tags_pk
					primary key,
			"name"  varchar(60)      not null,
			"CreatedAt" timestamp not null,
			"UpdatedAt" timestamp not null,
			"DeletedAt" timestamp
		);
	`
}
