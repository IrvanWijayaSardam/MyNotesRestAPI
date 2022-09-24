package post

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/IrvanWijayaSardam/GOData/models"
	pRepo "github.com/IrvanWijayaSardam/GOData/repository"
)

func NewSQLUserRepo(Conn *sql.DB) pRepo.UserRepo {
	return &user{
		db: Conn,
	}
}

type user struct {
	db *sql.DB
}

func (r *user) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.UserID,
			&data.Username,
			&data.Email,
			&data.Password,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil

}

func (r *user) Fetch(ctx context.Context, num int64) ([]*models.User, error) {
	query := "SELECT * FROM user limit ?"

	return r.fetch(ctx, query, num)
}

func (r *user) GetByID(ctx context.Context, UserID int64) (*models.User, error) {
	query := "SELECT * FROM user where UserID=?"

	rows, err := r.fetch(ctx, query, UserID)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}

	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}
func (r *user) Create(ctx context.Context, p *models.User) (int64, error) {
	query := "INSERT INTO user SET Username=?, Email=?, Password =?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Username, p.Email, p.Password)
	defer stmt.Close()

	fmt.Println(p.Username, p.Email, p.Password)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
func (r *user) Update(ctx context.Context, p *models.User) (*models.User, error) {
	query := "UPDATE user SET Username=?, Email=?, Password=? WHERE UserID=?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(
		ctx,
		p.Username,
		p.Email,
		p.Password,
		p.UserID,
	)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return p, nil

}
func (r *user) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM user WHERE UserID=?"

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
