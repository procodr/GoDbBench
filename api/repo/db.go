package repo

import (
	"api/core"
	"context"
	"database/sql"
	"time"
)

type DbTest struct {
	db *sql.DB
}

func NewDbTest(db *sql.DB) *DbTest {
	return &DbTest{db: db}
}

func (p *DbTest) Create(d *core.Data) error {
	query := `
	        INSERT INTO data (c1, c2, c3)
	        VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query, d.C1, d.C2, d.C3)
	return err
}
