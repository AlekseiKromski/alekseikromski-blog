package models

import (
	"fmt"
	"time"
)

type Post struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	CategoryID  int        `json:"category_id,omitempty"`
	Category    *Category  `json:"category,omitempty"`
	Comments    []*Comment `json:"comments"`
	Description string     `json:"description"`
	*Timestamp
	*SoftDelete
}

func CreatePostWithData(title, desc string, categoryID int) *Post {
	post := &Post{
		Title:       title,
		Description: desc,
		CategoryID:  categoryID,
		Comments:    []*Comment{},
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
		Comments:    []*Comment{},
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
	return fmt.Sprintf(`INSERT INTO posts ("title", "description", "category_id", "CreatedAt", "UpdatedAt", "DeletedAt") VALUES ('%s','%s', %d,'%s','%s', NULL)`, m.Title, m.Description, m.CategoryID, m.CreatedAt, m.UpdatedAt)
}

func UpdatePost(post *Post) string {
	return fmt.Sprintf(`UPDATE posts SET title = '%s', category_id = %d, description = '%s', "UpdatedAt" = '%s', "DeletedAt" = '%s' WHERE posts.id = %d`, post.Title, post.CategoryID, post.Description, post.UpdatedAt, *post.DeletedAt, post.ID)
}

func GetLastPosts(limit int, offset int, categoryID int) string {
	if categoryID == 0 {
		return fmt.Sprintf(`SELECT * FROM posts INNER JOIN categories c on c.id = posts.category_id ORDER BY posts."CreatedAt" DESC LIMIT %d OFFSET %d`, limit, offset)
	}
	return fmt.Sprintf(`SELECT * FROM posts INNER JOIN categories c on c.id = posts.category_id WHERE category_id = '%d' ORDER BY posts."CreatedAt" DESC LIMIT %d OFFSET %d`, categoryID, limit, offset)
}

func GetPost(postID int) string {
	return fmt.Sprintf(`SELECT * FROM posts INNER JOIN categories c on c.id = posts.category_id WHERE posts."id" = %d`, postID)
}
