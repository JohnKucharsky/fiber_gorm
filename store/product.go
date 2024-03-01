package store

import (
	"context"
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type ProductStore struct {
	db *pgxpool.Pool
}

func NewProductStore(db *pgxpool.Pool) *ProductStore {
	return &ProductStore{db: db}
}

func (us *ProductStore) Create(u domain.ProductRequest) (
	*domain.Product,
	error,
) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx, `
        INSERT INTO products (name, serial_number, created_at, updated_at)
        VALUES (@name, @serial_number, @created_at, @updated_at)
        RETURNING id, name, serial_number, created_at, updated_at`,
		pgx.NamedArgs{
			"name":          u.Name,
			"serial_number": u.SerialNumber,
			"created_at":    time.Now().Local(),
			"updated_at":    time.Now().Local(),
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.Product],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (us *ProductStore) GetMany() ([]*domain.Product, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx, `select * from products`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.Product],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *ProductStore) GetOne(id int) (*domain.Product, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx,
		`select * from products where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Product],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *ProductStore) Update(u domain.ProductRequest, id int) (
	*domain.Product,
	error,
) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx,
		`UPDATE products SET 
                updated_at = @updated_at,
                name = @name,
    			serial_number = @serial_number 
             WHERE id = @id 
        returning id,created_at,updated_at,name,serial_number`,
		pgx.NamedArgs{
			"id":            id,
			"updated_at":    time.Now().Local(),
			"name":          u.Name,
			"serial_number": u.SerialNumber,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Product],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *ProductStore) Delete(id int) (*domain.Product, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx,
		`delete from products where id = @id 
        returning id, created_at, updated_at, name, serial_number`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Product],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
