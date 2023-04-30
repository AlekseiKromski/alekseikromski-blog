package models

import (
	"fmt"
	"time"
)

type Comment struct {
	ID     int    `json:"ID"`
	Name   string `json:"name"`
	Text   string `json:"text"`
	PostID int    `json:"post_id"`
	*Timestamp
	*SoftDelete
}

func CreateComment() *Comment {
	category := &Comment{
		Name:       "",
		Text:       "",
		Timestamp:  &Timestamp{},
		SoftDelete: &SoftDelete{},
	}

	category.SetTimestamp()
	return category
}

func CreateCommentWithData(name, text string, postID int) *Comment {
	comment := &Comment{
		Name:       name,
		Text:       text,
		PostID:     postID,
		Timestamp:  &Timestamp{},
		SoftDelete: &SoftDelete{},
	}

	comment.SetTimestamp()
	return comment
}

func (c *Comment) Validate() bool {
	if len(c.Name) != 0 {
		return true
	}
	return false
}

func (c *Comment) SetTimestamp() {
	c.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(c.CreatedAt) == 0 {
		c.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (c *Comment) Soft() {
	softDeleteTs := time.Now().Format(time.RFC3339)
	c.DeletedAt = &softDeleteTs
}

func (c *Comment) Undo() {
	softDeleteTs := ""
	c.DeletedAt = &softDeleteTs
}

func (m *Comment) TableCreate() *TableCreation {
	return &TableCreation{
		Sql: `
			create table comments
			(
				ID          serial
					constraint comments_pk
						primary key,
				"name"  varchar(60)      not null,
				"text"  text     not null,
				"post_id"  serial     not null,
				"CreatedAt" timestamp not null,
				"UpdatedAt" timestamp not null,
				"DeletedAt" timestamp,
				CONSTRAINT FK_POST
					FOREIGN KEY(post_ID)
						REFERENCES posts(id)
			);
		`,
		Dependencies: []string{},
	}
}

func (c *Comment) CreateRecord() string {
	return fmt.Sprintf(`INSERT INTO comments ("name", "text", "post_id", "CreatedAt", "UpdatedAt", "DeletedAt") VALUES ('%s','%s','%d','%s','%s', NULL)`, c.Name, c.Text, c.PostID, c.CreatedAt, c.UpdatedAt)
}

func GetComments(post_ID int) string {
	return fmt.Sprintf(`SELECT * FROM comments WHERE post_id = %d AND comments."DeletedAt" IS NULL ORDER BY comments."CreatedAt" DESC`, post_ID)
}
