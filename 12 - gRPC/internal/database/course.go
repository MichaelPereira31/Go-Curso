package database

import "database/sql"

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}
 func (c *Course) Create() error {
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", c.ID, c.Name, c.Description, c.CategoryID)
	return err
}

func (c *Course) GetByID(id string) error {
	return c.db.QueryRow("SELECT id, name, description, category_id FROM courses WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description, &c.CategoryID)
}

func (c *Course) Update() error {
	_, err := c.db.Exec("UPDATE courses SET name = $1, description = $2, category_id = $3 WHERE id = $4", c.Name, c.Description, c.CategoryID, c.ID)
	return err
}

func (c *Course) Delete() error {
	_, err := c.db.Exec("DELETE FROM courses WHERE id = $1", c.ID)
	return err
}