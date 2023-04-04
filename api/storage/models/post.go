package models

import "time"

type Post struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	*Timestamp
	*SoftDelete
}

func (p *Post) Validate() bool {
	if len(p.Title) != 0 && len(p.Description) != 0 && p.ID != 0 {
		return true
	}

	return false
}

func (p *Post) SetTimestamp() {
	p.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(p.CreatedAt) != 0 {
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
		create table if not exists posts
		(
			id          serial
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
