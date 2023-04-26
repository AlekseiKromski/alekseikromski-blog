package models

import (
	"fmt"
	"time"
)

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CategoryID  int       `json:"category_id,omitempty"`
	Category    *Category `json:"category,omitempty"`
	Description string    `json:"description"`
	*Timestamp
	*SoftDelete
}

func CreatePostWithData(title, desc string) *Post {
	post := &Post{
		Title:       title,
		Description: desc,
		Timestamp:   &Timestamp{},
		SoftDelete:  &SoftDelete{},
	}

	post.SetTimestamp()
	return post
}

func CreatePost() *Post {
	post := &Post{
		Title:       "",
		Description: "",
		Timestamp:   &Timestamp{},
		SoftDelete:  &SoftDelete{},
	}

	post.SetTimestamp()
	return post
}

func (p *Post) Validate() bool {
	if len(p.Title) != 0 && len(p.Description) != 0 {
		return true
	}

	return false
}

func (p *Post) SetTimestamp() {
	p.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(p.CreatedAt) == 0 {
		p.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (p *Post) Soft() {
	softDeleteTs := time.Now().Format(time.RFC3339)
	p.DeletedAt = &softDeleteTs
}

func (p *Post) Undo() {
	softDeleteTs := time.Now().Format(time.RFC3339)
	p.DeletedAt = &softDeleteTs
}

func (m *Post) TableCreate() *TableCreation {
	return &TableCreation{
		Sql: `
			create table posts
			(
				ID          serial
					constraint posts_pk
						primary key,
				"title"  varchar(60)      not null,
				"description"  text      not null,
				"category_id" serial not null,
				"CreatedAt" timestamp not null,
				"UpdatedAt" timestamp not null,
				"DeletedAt" timestamp,
				CONSTRAINT FK_CATEGORY
					FOREIGN KEY(category_id)
						REFERENCES categories(id)
			);
		`,
		Dependencies: []string{"Category"},
	}
}

func (m *Post) CreateRecord() string {
	return fmt.Sprintf(`INSERT INTO posts ("title", "description", "CreatedAt", "UpdatedAt", "DeletedAt") VALUES ('%s','%s','%s','%s', NULL)`, m.Title, m.Description, m.CreatedAt, m.UpdatedAt)
}

func UpdatePost(post *Post) string {
	return fmt.Sprintf(`UPDATE posts SET title = '%s', category_id = %d, description = '%s', "UpdatedAt" = '%s' WHERE posts.id = %d`, post.Title, post.CategoryID, post.Description, time.Now().Format(time.RFC3339), post.ID)
}

func GetLastPosts(limit int, offset int, categoryID int) (string, bool) {
	if categoryID == 0 {
		return fmt.Sprintf(`SELECT * FROM posts ORDER BY "CreatedAt" DESC LIMIT %d OFFSET %d`, limit, offset), false
	}
	return fmt.Sprintf(`SELECT * FROM posts INNER JOIN categories c on c.id = posts.category_id WHERE category_id = '%d' ORDER BY posts."CreatedAt" DESC LIMIT %d OFFSET %d`, categoryID, limit, offset), true
}

func GetPost(postID int) string {
	return fmt.Sprintf(`SELECT * FROM posts INNER JOIN categories c on c.id = posts.category_id WHERE posts."id" = %d`, postID)
}
