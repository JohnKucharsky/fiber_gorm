package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type OrderStore struct {
	db *pgxpool.Pool
}

func NewOrderStore(db *pgxpool.Pool) *OrderStore {
	return &OrderStore{db: db}
}

func (us *OrderStore) Create(u domain.OrderRequest) (
	int,
	error,
) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx, `
        INSERT INTO orders (user_id, product_id, updated_at)
        VALUES (@user_id, @product_id, @updated_at)
        RETURNING id`,
		pgx.NamedArgs{
			"user_id":    u.UserID,
			"product_id": u.ProductID,
			"updated_at": time.Now().Local(),
		},
	)
	if err != nil {
		return 0, err
	}

	type returnedRow struct {
		ID int `db:"id"`
	}
	row, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[returnedRow],
	)
	if err != nil {
		return 0, err
	}

	return row.ID, nil
}

func query() string {
	return `SELECT
			JSON_BUILD_OBJECT(
				'id', orders.id,
				'updated_at', orders.updated_at,
				'user', JSON_BUILD_OBJECT(
					'id', users.id,
				    'created_at', users.created_at,
				    'updated_at', users.updated_at,
				    'first_name', users.first_name,
				    'last_name', users.last_name
				),
				'product', JSON_BUILD_OBJECT(
					'id', products.id,
				    'created_at', products.created_at,
				    'updated_at', products.updated_at,
				    'name', products.name,
				    'serial_number', products.serial_number
				)
			)
		FROM orders
		    left join users on orders.user_id = users.id
		    left join products on orders.product_id = products.id`
}

func (us *OrderStore) GetMany() ([]*domain.Order, error) {
	ctx := context.Background()

	rows, err := us.db.Query(
		ctx, query(),
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pginit.JSONRowToAddrOfStruct[domain.Order],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *OrderStore) GetOne(id int) (*domain.Order, error) {
	ctx := context.Background()

	if id == 0 {
		return nil, errors.New("id is 0")
	}
	query := query()
	query += fmt.Sprintf(" where orders.id = %d", id)

	rows, err := us.db.Query(
		ctx, query,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pginit.JSONRowToAddrOfStruct[domain.Order],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *OrderStore) Update(u domain.OrderRequest, id int) error {
	ctx := context.Background()

	_, err := us.db.Exec(
		ctx,
		`UPDATE orders SET 
                updated_at = @updated_at,
                user_id = @user_id,
    			product_id = @product_id 
             WHERE id = @id returning id`,
		pgx.NamedArgs{
			"id":         id,
			"updated_at": time.Now().Local(),
			"user_id":    u.UserID,
			"product_id": u.ProductID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (us *OrderStore) Delete(id int) error {
	ctx := context.Background()

	_, err := us.db.Query(
		ctx,
		`delete from orders where id = @id returning id`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
