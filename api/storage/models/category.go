package models

import "time"

type Category struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	*Timestamp
	*SoftDelete
}

func CreateCategory() *Category {
	category := &Category{
		Name:       "",
		Timestamp:  &Timestamp{},
		SoftDelete: &SoftDelete{},
	}

	category.SetTimestamt()
	return category
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
	softDeleteTs := time.Now().Format(time.RFC3339)
	c.DeletedAt = &softDeleteTs
}

func (c *Category) Undo() {
	softDeleteTs := ""
	c.DeletedAt = &softDeleteTs
}

func (m *Category) TableCreate() *TableCreation {
	return &TableCreation{
		Sql: `
			create table categories
			(
				ID          serial
					constraint categories_pk
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

func GetCategories() string {
	return `SELECT * FROM categories ORDER BY "CreatedAt"`
}
