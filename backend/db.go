package main

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"strings"
)

type sqlRepo struct {
	db *sqlx.DB
}

func NewRepo(c *Config) (*sqlRepo, error) {
	driver, dsn, _ := strings.Cut(c.DataSourceName, "://")
	db, err := sqlx.Open(driver, dsn)
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

		CREATE TABLE IF NOT EXISTS flags (
		    "when"			integer 	NOT NULL,
		    "user"			text 		NOT NULL,
		    flag 			text 		NOT NULL,
		    email 			text 		NOT NULL PRIMARY KEY
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
		repo.db.Rebind("INSERT INTO results(district, democrats, republicans, invalid) VALUES (?, ?, ?, ?)"),
		district, r.Democrats, r.Republicans, r.Invalid)
	return err
}

func (repo *sqlRepo) FlagSubmission(ctx context.Context, when, user, flag, email string) error {
	_, err := repo.db.ExecContext(ctx,
		repo.db.Rebind("INSERT INTO flags(\"when\", \"user\", flag, email) VALUES (?, ?, ?, ?)"),
		when, user, flag, email)
	return err
}
