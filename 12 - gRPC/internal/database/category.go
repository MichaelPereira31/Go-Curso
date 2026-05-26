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
 func (c *Category) Create() error {
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", c.ID, c.Name, c.Description)
	return err
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