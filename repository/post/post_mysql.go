package post

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/IrvanWijayaSardam/GOData/models"
	pRepo "github.com/IrvanWijayaSardam/GOData/repository"
)

func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &Post{
		db: Conn,
	}
}

type Post struct {
	db *sql.DB
}

func (r *Post) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Post, 0)
	for rows.Next() {
		data := new(models.Post)

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Description,
			&data.Date,
			&data.UserID,
			&data.Image,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil

}

func (r *Post) Fetch(ctx context.Context, num int64) ([]*models.Post, error) {
	query := "SELECT * FROM notes limit ?"

	return r.fetch(ctx, query, num)
}

func (r *Post) GetByID(ctx context.Context, UserID int64) (*models.Post, error) {
	query := "SELECT * FROM notes where UserID=?"

	rows, err := r.fetch(ctx, query, UserID)
	if err != nil {
		return nil, err
	}

	payload := &models.Post{}

	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}
func (r *Post) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "INSERT INTO notes SET ID=?, Title=?, Description=?, Date=?,UserID =?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.ID, p.Title, p.Description, p.Date, p.UserID)
	defer stmt.Close()

	fmt.Println(p.ID, p.Title, p.Description, p.UserID)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
func (r *Post) Update(ctx context.Context, p *models.Post) (*models.Post, error) {
	query := "UPDATE notes SET title=?, description=?,Date =?, userid=? WHERE id=?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Description,
		p.Date,
		p.UserID,
		p.ID,
	)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return p, nil

}
func (r *Post) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM notes WHERE id=?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
