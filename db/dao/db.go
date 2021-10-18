package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func NewDB(db DB) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DB) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addKeyValue, err = db.PrepareContext(ctx, addKeyValue); err != nil {
		return nil, fmt.Errorf("error preparing query addKeyValue: %w", err)
	}
	if q.updateKeyValue, err = db.PrepareContext(ctx, updateKeyValue); err != nil {
		return nil, fmt.Errorf("error preparing query updateKeyValue: %w", err)
	}
	if q.getKeyValue, err = db.PrepareContext(ctx, getKeyValue); err != nil {
		return nil, fmt.Errorf("error preparing query getKeyValue: %w", err)
	}
	if q.deleteKeyValue, err = db.PrepareContext(ctx, deleteKeyValue); err != nil {
		return nil, fmt.Errorf("error preparing query deleteKeyValue: %w", err)
	}
	if q.addUptimeWatchRequest, err = db.PrepareContext(ctx, addUptimeWatchRequest); err != nil {
		return nil, fmt.Errorf("error preparing query addUptimeWatchRequest: %w", err)
	}
	if q.getUptimeWatchRequestByID, err = db.PrepareContext(ctx, getUptimeWatchRequestByID); err != nil {
		return nil, fmt.Errorf("error preparing query getUptimeWatchRequestByID: %w", err)
	}
	if q.getAllUptimeWatchRequest, err = db.PrepareContext(ctx, getAllUptimeWatchRequest); err != nil {
		return nil, fmt.Errorf("error preparing query getAllUptimeWatchRequest: %w", err)
	}
	if q.deleteUptimeWatchRequestById, err = db.PrepareContext(ctx, deleteUptimeWatchRequestById); err != nil {
		return nil, fmt.Errorf("error preparing query getAllUptimeWatchRequest: %w", err)
	}
	if q.addUptimeResult, err = db.PrepareContext(ctx, addUptimeResult); err != nil {
		return nil, fmt.Errorf("error preparing query addUptimeResult: %w", err)
	}
	if q.getUptimeResults, err = db.PrepareContext(ctx, getUptimeResults); err != nil {
		return nil, fmt.Errorf("error preparing query getUptimeResults: %w", err)
	}
	if q.getUptimeResultCount, err = db.PrepareContext(ctx, getUptimeResultCount); err != nil {
		return nil, fmt.Errorf("error preparing query getUptimeResultCount: %w", err)
	}
	if q.deleteUptimeResults, err = db.PrepareContext(ctx, deleteUptimeResults); err != nil {
		return nil, fmt.Errorf("error preparing query deleteUptimeResults: %w", err)
	}
	if q.getUptimeResultStatsForID, err = db.PrepareContext(ctx, getUptimeResultStatsForID); err != nil {
		return nil, fmt.Errorf("error preparing query getUptimeResultStatsForID: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addKeyValue != nil {
		if cerr := q.addKeyValue.Close(); cerr != nil {
			err = fmt.Errorf("error closing addKeyValue: %w", cerr)
		}
	}
	if q.updateKeyValue != nil {
		if cerr := q.updateKeyValue.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateKeyValue: %w", cerr)
		}
	}
	if q.getKeyValue != nil {
		if cerr := q.getKeyValue.Close(); cerr != nil {
			err = fmt.Errorf("error closing getKeyValue: %w", cerr)
		}
	}
	if q.deleteKeyValue != nil {
		if cerr := q.deleteKeyValue.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteKeyValue: %w", cerr)
		}
	}
	if q.addUptimeWatchRequest != nil {
		if cerr := q.addUptimeWatchRequest.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUptimeWatchRequest: %w", cerr)
		}
	}
	if q.getUptimeWatchRequestByID != nil {
		if cerr := q.getUptimeWatchRequestByID.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUptimeWatchRequestByID: %w", cerr)
		}
	}
	if q.getAllUptimeWatchRequest != nil {
		if cerr := q.getAllUptimeWatchRequest.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUptimeWatchRequest: %w", cerr)
		}
	}
	if q.deleteUptimeWatchRequestById != nil {
		if cerr := q.deleteUptimeWatchRequestById.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUptimeWatchRequestById: %w", cerr)
		}
	}
	if q.addUptimeResult != nil {
		if cerr := q.addUptimeResult.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUptimeResult: %w", cerr)
		}
	}
	if q.getUptimeResults != nil {
		if cerr := q.getUptimeResults.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUptimeResults: %w", cerr)
		}
	}
	if q.getUptimeResultCount != nil {
		if cerr := q.getUptimeResultCount.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUptimeResultCount: %w", cerr)
		}
	}
	if q.deleteUptimeResults != nil {
		if cerr := q.deleteUptimeResults.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUptimeResults: %w", cerr)
		}
	}
	if q.getUptimeResultStatsForID != nil {
		if cerr := q.getUptimeResultStatsForID.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUptimeResultStatsForID: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DB
	addKeyValue                  *sql.Stmt
	updateKeyValue               *sql.Stmt
	getKeyValue                  *sql.Stmt
	deleteKeyValue               *sql.Stmt
	addUptimeWatchRequest        *sql.Stmt
	getUptimeWatchRequestByID    *sql.Stmt
	getAllUptimeWatchRequest     *sql.Stmt
	deleteUptimeWatchRequestById *sql.Stmt
	addUptimeResult              *sql.Stmt
	getUptimeResultCount         *sql.Stmt
	getUptimeResults             *sql.Stmt
	deleteUptimeResults          *sql.Stmt
	getUptimeResultStatsForID    *sql.Stmt
}
