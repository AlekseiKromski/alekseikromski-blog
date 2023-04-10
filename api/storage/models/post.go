package models

import (
	"fmt"
	"time"
)

type Post struct {
	ID          int    `json:"omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
	p.DeletedAt = time.Now().Format(time.RFC3339)
}

func (p *Post) Undo() {
	p.DeletedAt = ""
}

func (m *Post) TableCreate() string {
	return `
		create table posts
		(
			ID          serial
				constraint posts_pk
					primary key,
			"title"  varchar(60)      not null,
			"description"  text      not null,
			"CreatedAt" timestamp not null,
			"UpdatedAt" timestamp not null,
			"DeletedAt" timestamp
		);
	`
}

func (m *Post) CreateRecord() string {
	return fmt.Sprintf(`INSERT INTO posts ("title", "description", "CreatedAt", "UpdatedAt", "DeletedAt") VALUES ('%s','%s','%s','%s', NULL)`, m.Title, m.Description, m.CreatedAt, m.UpdatedAt)
}

func GetLastPosts(limit int) string {
	return fmt.Sprintf(`SELECT * FROM posts ORDER BY "CreatedAt" DESC LIMIT %d`, limit)
}
