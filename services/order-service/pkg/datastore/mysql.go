package datastore

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ErwinSalas/go-eda/services/order-service/pkg/types"
	_ "github.com/go-sql-driver/mysql"
)

type IDatastore interface {
	InsertOrder(ctx context.Context, order *types.Order) (int64, error)
	GetLastOrder(ctx context.Context) (*types.Order, error)
	Migrate(ctx context.Context) error
	GetOrders(ctx context.Context) ([]*types.Order, error)
	UpdateOrderStatus(ctx context.Context, id int, status string) error
}

type Order struct {
	database *sql.DB
}

func NewDataStore(db *sql.DB) (*Order, error) {
	return &Order{database: db}, nil
}

const (
	insertOrderQuery  = "INSERT INTO orders (customer_details, items, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	getLastOrderQuery = "SELECT id, customer_details, items, status, created_at, updated_at FROM orders ORDER BY id DESC LIMIT 1"
	migrateQuery      = `
        CREATE TABLE IF NOT EXISTS orders (
            id INT AUTO_INCREMENT PRIMARY KEY,
            customer_details VARCHAR(255),
            items JSON,
            status VARCHAR(50),
            created_at TIMESTAMP,
            updated_at TIMESTAMP
        )
    `
	getOrdersQuery         = "SELECT id, customer_details, items, status, created_at, updated_at FROM orders"
	updateOrderStatusQuery = "UPDATE orders SET status = ? WHERE id = ?"
)

func (d *Order) InsertOrder(ctx context.Context, order *types.Order) (int64, error) {
	items, err := json.Marshal(order.Items)
	if err != nil {
		return 0, err
	}

	result, err := d.database.ExecContext(ctx, insertOrderQuery, order.CustomerDetails, items, order.Status, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Order) GetLastOrder(ctx context.Context) (*types.Order, error) {
	order := &types.Order{}
	var items json.RawMessage
	var createdAt, updatedAt sql.NullTime

	err := d.database.QueryRowContext(ctx, getLastOrderQuery).Scan(&order.ID, &order.CustomerDetails, &items, &order.Status, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(items, &order.Items); err != nil {
		return nil, err
	}

	if createdAt.Valid {
		order.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		order.UpdatedAt = updatedAt.Time
	}

	return order, nil
}

func (d *Order) Migrate(ctx context.Context) error {
	_, err := d.database.ExecContext(ctx, migrateQuery)
	return err
}

func (d *Order) GetOrders(ctx context.Context) ([]*types.Order, error) {
	rows, err := d.database.QueryContext(ctx, getOrdersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*types.Order

	for rows.Next() {
		order := &types.Order{}
		var items json.RawMessage
		var createdAt, updatedAt sql.NullTime

		err = rows.Scan(&order.ID, &order.CustomerDetails, &items, &order.Status, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(items, &order.Items); err != nil {
			return nil, err
		}

		if createdAt.Valid {
			order.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			order.UpdatedAt = updatedAt.Time
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (d *Order) UpdateOrderStatus(ctx context.Context, id int, status string) error {
	_, err := d.database.ExecContext(ctx, updateOrderStatusQuery, status, id)
	return err
}
