package models

import (
	"fmt"
	"time"
)

type Category struct {
	ID   int    `json:"id"`
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

func CreateCategoryWithData(name string) *Category {
	category := &Category{
		Name:       name,
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
	if len(c.CreatedAt) == 0 {
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

func (c *Category) CreateRecord() string {
	return fmt.Sprintf(`INSERT INTO categories ("name", "CreatedAt", "UpdatedAt", "DeletedAt") VALUES ('%s','%s','%s', NULL)`, c.Name, c.CreatedAt, c.UpdatedAt)
}

func GetCategories() string {
	return `SELECT * FROM categories WHERE categories."DeletedAt" IS NULL ORDER BY "CreatedAt" DESC`
}
func GetCategory(id int) string {
	return fmt.Sprintf(`SELECT * FROM categories WHERE id = %d AND categories."DeletedAt" IS NULL`, id)
}

func UpdateCategory(category *Category) string {
	if category.DeletedAt == nil {
		return fmt.Sprintf(
			`UPDATE categories SET name = '%s', "UpdatedAt" = '%s' WHERE categories.id = %d`,
			category.Name, category.UpdatedAt, category.ID,
		)
	}
	return fmt.Sprintf(
		`UPDATE categories SET name = '%s', "UpdatedAt" = '%s', "DeletedAt" = '%s' WHERE categories.id = %d`,
		category.Name, category.UpdatedAt, *category.DeletedAt, category.ID,
	)
}
