package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type sqlRepo struct {
	db *sql.DB
}

func NewRepo(c *Config) (*sqlRepo, error) {
	driver, dsn, _ := strings.Cut(c.DataSourceName, "://")
	fmt.Printf("driver: %s, dsn: %s\n", driver, dsn)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	repo := &sqlRepo{
		db: db,
	}
	if err := repo.createTables(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (repo *sqlRepo) createTables() error {
	_, err := repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS results (
			district		text 		NOT NULL PRIMARY KEY,
			democrats 		integer		NOT NULL,
			republicans 	integer		NOT NULL,
			invalid 		integer		NOT NULL
		);
	`)
	return err
}

func (repo *sqlRepo) List(ctx context.Context) (res []*result, err error) {
	res = make([]*result, 0)
	row, err := repo.db.QueryContext(ctx, `SELECT * FROM results ORDER BY district`)
	if err != nil {
		return res, err
	}

	defer row.Close()
	for row.Next() {
		r := &result{}
		if err := row.Scan(&r.District, &r.Democrats, &r.Republicans, &r.Invalid); err != nil {
			return nil, err
		}
		res = append(res, r)
	}

	return res, nil
}

func (repo *sqlRepo) Submit(ctx context.Context, district string, r *result) error {
	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO results(district, democrats, republicans, invalid) VALUES (?, ?, ?, ?)",
		district, r.Democrats, r.Republicans, r.Invalid)
	return err
}
