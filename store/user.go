package store

import (
	"context"
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) Create(u domain.UserRequest) (*domain.User, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx, `
        INSERT INTO users (first_name, last_name, created_at, updated_at)
        VALUES (@first_name, @last_name, @created_at, @updated_at)
        RETURNING id, first_name,last_name, created_at, updated_at`,
		pgx.NamedArgs{
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"created_at": time.Now().Local(),
			"updated_at": time.Now().Local(),
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.User],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (us *UserStore) GetMany() ([]*domain.User, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx, `select * from users`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.User],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *UserStore) GetOne(id int) (*domain.User, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx,
		`select * from users where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.User],
	)
	if err != nil {

		return nil, err
	}

	return res, nil
}

func (us *UserStore) Update(u domain.UserRequest, id int) (
	*domain.User,
	error,
) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx,
		`UPDATE users SET 
                updated_at = @updated_at,
                first_name = @first_name,
    			last_name = @last_name 
    	WHERE id = @id 
    	returning id,created_at,updated_at,first_name,last_name`,
		pgx.NamedArgs{
			"id":         id,
			"updated_at": time.Now().Local(),
			"first_name": u.FirstName,
			"last_name":  u.LastName,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.User],
	)
	if err != nil {

		return nil, err
	}

	return res, nil
}

func (us *UserStore) Delete(id int) (*domain.User, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx,
		`delete from users where id = @id 
        returning id, created_at, updated_at, first_name, last_name`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.User],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
