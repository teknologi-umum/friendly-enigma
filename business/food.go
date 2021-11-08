package business

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/georgysavva/scany/sqlscan"
)

// GetAllFoods returns all the foods in the refrigerator.
func GetAllFoods(conn *sql.Conn, ctx context.Context) ([]Food, error) {
	rows, err := conn.QueryContext(
		ctx,
		"SELECT * FROM refrigerator",
	)
	if err != nil {
		return []Food{}, err
	}
	defer rows.Close()

	// This is how you would scan a result from the SQLite rows
	// to array of structs with scanny.
	var foods []Food
	err = sqlscan.ScanAll(&foods, rows)
	if err != nil {
		return []Food{}, err
	}

	return foods, nil
}

// AddNewFood adds a new food to the refrigerator.
// This function is broken, but why?
// I'm sure you can fix it!  :D
func AddNewFood(conn *sql.Conn, ctx context.Context, food Food) (Food, error) {
	// Let's create a SQL transaction instance as one of the way
	// to handle race condition.
	//
	// Learn about SQL race condition:
	// Timestamped: https://youtu.be/Wb0DM9I8RDo?t=8758
	// If you prefer reading: https://stackoverflow.com/questions/34510/what-is-a-race-condition
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return Food{}, err
	}

	// Check if the food already exists
	exists, err := tx.QueryContext(ctx, "SELECT id FROM refrigerator WHERE name = ?", food.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		tx.Rollback()
		return Food{}, err
	}
	defer exists.Close()

	var existingFoodName string
	if exists.Next() {
		err = exists.Scan(&existingFoodName)
		if err != nil {
			tx.Rollback()
			return Food{}, err
		}
	}

	if existingFoodName != "" {
		return Food{}, errors.New("Food already exists")
	}

	// If the food does not exist, we can insert it.
	_, err = tx.ExecContext(ctx, "INSERT INTO refrigerator (name, quantity, updated_at) VALUES ($1, $2, $2)", food.Name, food.Quantity, time.Now())
	if err != nil {
		tx.Rollback()
		return Food{}, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return Food{}, err
	}

	return food, nil
}
