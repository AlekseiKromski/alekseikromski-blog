package models

import "time"

type Tag struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	*Timestamp
	*SoftDelete
}

func CreateTag() *Tag {
	tag := &Tag{
		Name:       "",
		Timestamp:  &Timestamp{},
		SoftDelete: &SoftDelete{},
	}

	tag.SetTimestamp()
	return tag
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
	softDeleteTs := time.Now().Format(time.RFC3339)
	t.DeletedAt = &softDeleteTs
}

func (t *Tag) Undo() {
	softDeleteTs := time.Now().Format(time.RFC3339)
	t.DeletedAt = &softDeleteTs
}

func (m *Tag) TableCreate() *TableCreation {
	return &TableCreation{
		Sql: `
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
		`,
		Dependencies: []string{},
	}
}

func GetTags() string {
	return `SELECT * FROM tags ORDER BY "CreatedAt"`
}
