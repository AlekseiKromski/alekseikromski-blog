package models

import "time"

type Category struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	*Timestamp
	*SoftDelete
}

func (c *Category) Validate() bool {
	if len(c.Name) != 0 {
		return true
	}
	return false
}

func (c *Category) SetTimestamt() {
	c.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(c.CreatedAt) != 0 {
		c.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (c *Category) Soft() {
	c.DeletedAt = time.Now().Format(time.RFC3339)
}

func (c *Category) Undo() {
	c.DeletedAt = ""
}
