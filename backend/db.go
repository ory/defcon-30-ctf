package main

import (
	"context"
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type (
	sqliteRepo struct {
		db *sql.DB
	}
)

func NewRepo(c *Config) (*sqliteRepo, error) {
	db, err := sql.Open("sqlite3", c.DataSourceName)
	if err != nil {
		return nil, err
	}
	repo := &sqliteRepo{
		db: db,
	}
	if err := repo.createTables(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (repo *sqliteRepo) createTables() error {
	_, err := repo.db.Exec(`
		CREATE TABLE results (
			district_id		integer 	NOT NULL PRIMARY KEY,
			votes 			text		NOT NULL
		);
	`)
	return err
}

func (repo *sqliteRepo) List(ctx context.Context) (res []*result, err error) {
	res = make([]*result, 0)
	row, err := repo.db.QueryContext(ctx, `SELECT * FROM results ORDER BY district_id`)
	if err != nil {
		return res, err
	}

	defer row.Close()
	for row.Next() {
		var (
			r        = &result{}
			rawVotes string
		)
		if err := row.Scan(&r.DistrictID, &rawVotes); err != nil {
			return res, err
		}
		if err := json.Unmarshal([]byte(rawVotes), &r.Votes); err != nil {
			return res, err
		}
		res = append(res, r)
	}

	return res, nil
}

func (repo *sqliteRepo) Submit(ctx context.Context, r *result) error {
	rawVotes, err := json.Marshal(r.Votes)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(ctx,
		"INSERT INTO results(district_id, votes) VALUES (?, JSON(?))",
		r.DistrictID, rawVotes,
	)
	return err
}
