package database

import "database/sql"

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}
func (c *Category) Create(name, description string) (*Category, error) {
	category := &Category{
		db:          c.db,
		Name:        name,
		Description: description,
	}

	err := c.db.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id, name, description",
		category.Name,
		category.Description,
	).Scan(&category.ID, &category.Name, &category.Description)

	return category, err
}

func (c *Category) GetByID(id string) error {
	return c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description)
}

func (c *Category) Update() error {
	_, err := c.db.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", c.Name, c.Description, c.ID)
	return err
}

func (c *Category) Delete() error {
	_, err := c.db.Exec("DELETE FROM categories WHERE id = $1", c.ID)
	return err
}
