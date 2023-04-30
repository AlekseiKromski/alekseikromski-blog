package models

import (
	"fmt"
	"time"
)

type Tag struct {
	ID     int    `json:"ID"`
	Name   string `json:"name"`
	PostID int    `json:"post_id"`
	*Timestamp
	*SoftDelete
}

func CreateTag() *Tag {
	tag := &Tag{
		Name:       "",
		PostID:     0,
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
				"post_id"  serial     not null,
				"CreatedAt" timestamp not null,
				"UpdatedAt" timestamp not null,
				"DeletedAt" timestamp,
				CONSTRAINT FK_TAG
					FOREIGN KEY(post_id)
						REFERENCES posts(id)
			);
		`,
		Dependencies: []string{},
	}
}

func GetTags(postID *int) string {
	if postID == nil {
		return `SELECT * FROM tags WHERE tags."DeletedAt" is NULL ORDER BY "CreatedAt" DESC`
	}
	return fmt.Sprintf(`SELECT * FROM tags WHERE tags.post_id = %d AND tags."DeletedAt" is NULL ORDER BY "CreatedAt" DESC`, *postID)
}

func GetTagByID(id int) string {
	return fmt.Sprintf(`SELECT * FROM tags WHERE tags.id = %d AND tags."DeletedAt" is NULL`, id)
}

func UpdateTag(tag *Tag) string {
	if tag.DeletedAt == nil {
		return fmt.Sprintf(
			`UPDATE tags SET name = '%s', post_id = %d, "UpdatedAt" = '%s'WHERE tags.id = %d`,
			tag.Name, tag.PostID, tag.UpdatedAt, tag.ID,
		)
	}
	return fmt.Sprintf(
		`UPDATE tags SET name = '%s', post_id = %d, "UpdatedAt" = '%s', "DeletedAt" = '%s' WHERE tags.id = %d`,
		tag.Name, tag.PostID, tag.UpdatedAt, *tag.DeletedAt, tag.ID,
	)
}
