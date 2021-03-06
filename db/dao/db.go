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

	// key value
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

	// uptime watch request
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

	// uptime result
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

	// uptime conclusion
	if q.addUptimeConclusion, err = db.PrepareContext(ctx, addUptimeConclusion); err != nil {
		return nil, fmt.Errorf("error preparing query addUptimeConclusion: %w", err)
	}
	if q.deleteUptimeConclusionByUWRID, err = db.PrepareContext(ctx, deleteUptimeConclusionByUWRID); err != nil {
		return nil, fmt.Errorf("error preparing query deleteUptimeConclusionByID: %w", err)
	}
	if q.getUptimeConclusionByUWRID, err = db.PrepareContext(ctx, getUptimeConclusionByUWRID); err != nil {
		return nil, fmt.Errorf("error preparing query getUptimeConclusionByID: %w", err)
	}
	if q.getAllUptimeConclusion, err = db.PrepareContext(ctx, getAllUptimeConclusion); err != nil {
		return nil, fmt.Errorf("error preparing query getAllUptimeConclusion: %w", err)
	}

	//uptime ssl info
	if q.addUptimeSSLInfo, err = db.PrepareContext(ctx, addUptimeSSLInfo); err != nil {
		return nil, fmt.Errorf("error preparing query addUptimeSSLInfo: %w", err)
	}
	if q.deleteUptimeSSLInfoByUWRID, err = db.PrepareContext(ctx, deleteUptimeSSLInfoByUWRID); err != nil {
		return nil, fmt.Errorf("error preparing query deleteUptimeSSLInfoByUWRID: %w", err)
	}
	if q.updateUptimeSSLInfoByUWRID, err = db.PrepareContext(ctx, updateUptimeSSLInfoByUWRID); err != nil {
		return nil, fmt.Errorf("error preparing query updateUptimeSSLInfoByUWRID: %w", err)
	}
	if q.getUptimeSSLInfoByUWRID, err = db.PrepareContext(ctx, getUptimeSSLInfoByUWRID); err != nil {
		return nil, fmt.Errorf("error preparing query getUptimeSSLInfoByUWRID: %w", err)
	}
	if q.getAllUptimeSSLInfo, err = db.PrepareContext(ctx, getAllUptimeSSLInfo); err != nil {
		return nil, fmt.Errorf("error preparing query getAllUptimeSSLInfo: %w", err)
	}

	// service
	if q.addService, err = db.PrepareContext(ctx, addService); err != nil {
		return nil, fmt.Errorf("error preparing query addService: %w", err)
	}
	if q.deleteServiceByID, err = db.PrepareContext(ctx, deleteServiceByID); err != nil {
		return nil, fmt.Errorf("error preparing query deleteServiceByID: %w", err)
	}
	if q.getServiceByID, err = db.PrepareContext(ctx, getServiceByID); err != nil {
		return nil, fmt.Errorf("error preparing query getServiceByID: %w", err)
	}
	if q.getAllService, err = db.PrepareContext(ctx, getAllService); err != nil {
		return nil, fmt.Errorf("error preparing query getAllService: %w", err)
	}

	return &q, nil
}

func (q *Queries) Close() error {
	var err error

	// key value
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

	// uptime watch request
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

	// uptime result
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

	// uptime conclusion
	if q.addUptimeConclusion != nil {
		if cerr := q.addUptimeConclusion.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUptimeConclusion: %w", cerr)
		}
	}
	if q.deleteUptimeConclusionByUWRID != nil {
		if cerr := q.deleteUptimeConclusionByUWRID.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUptimeConclusionByID: %w", cerr)
		}
	}
	if q.getUptimeConclusionByUWRID != nil {
		if cerr := q.getUptimeConclusionByUWRID.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUptimeConclusionByID: %w", cerr)
		}
	}
	if q.getAllUptimeConclusion != nil {
		if cerr := q.getAllUptimeConclusion.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUptimeConclusion: %w", cerr)
		}
	}

	// uptime ssl info
	if q.addUptimeSSLInfo != nil {
		if cerr := q.addUptimeSSLInfo.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUptimeSSLInfo: %w", cerr)
		}
	}
	if q.deleteUptimeSSLInfoByUWRID != nil {
		if cerr := q.deleteUptimeSSLInfoByUWRID.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUptimeSSLInfoByUWRID: %w", cerr)
		}
	}
	if q.updateUptimeSSLInfoByUWRID != nil {
		if cerr := q.updateUptimeSSLInfoByUWRID.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUptimeSSLInfoByUWRID: %w", cerr)
		}
	}
	if q.getUptimeSSLInfoByUWRID != nil {
		if cerr := q.getUptimeSSLInfoByUWRID.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUptimeSSLInfoByUWRID: %w", cerr)
		}
	}
	if q.getAllUptimeSSLInfo != nil {
		if cerr := q.getAllUptimeSSLInfo.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUptimeSSLInfo: %w", cerr)
		}
	}

	// service
	if q.addService != nil {
		if cerr := q.addService.Close(); cerr != nil {
			err = fmt.Errorf("error closing addService: %w", cerr)
		}
	}
	if q.deleteServiceByID != nil {
		if cerr := q.deleteServiceByID.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteServiceByID: %w", cerr)
		}
	}
	if q.getServiceByID != nil {
		if cerr := q.getServiceByID.Close(); cerr != nil {
			err = fmt.Errorf("error closing getServiceByID: %w", cerr)
		}
	}
	if q.getAllService != nil {
		if cerr := q.getAllService.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllService: %w", cerr)
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
	db DB

	// keyvalue
	addKeyValue    *sql.Stmt
	updateKeyValue *sql.Stmt
	getKeyValue    *sql.Stmt
	deleteKeyValue *sql.Stmt

	// uptime watch request
	addUptimeWatchRequest        *sql.Stmt
	getUptimeWatchRequestByID    *sql.Stmt
	getAllUptimeWatchRequest     *sql.Stmt
	deleteUptimeWatchRequestById *sql.Stmt

	// uptime watch result
	addUptimeResult           *sql.Stmt
	getUptimeResultCount      *sql.Stmt
	getUptimeResults          *sql.Stmt
	deleteUptimeResults       *sql.Stmt
	getUptimeResultStatsForID *sql.Stmt

	// uptime conclusion
	addUptimeConclusion           *sql.Stmt
	deleteUptimeConclusionByUWRID *sql.Stmt
	getUptimeConclusionByUWRID    *sql.Stmt
	getAllUptimeConclusion        *sql.Stmt

	//uptime SSL info
	addUptimeSSLInfo           *sql.Stmt
	deleteUptimeSSLInfoByUWRID *sql.Stmt
	updateUptimeSSLInfoByUWRID *sql.Stmt
	getUptimeSSLInfoByUWRID    *sql.Stmt
	getAllUptimeSSLInfo        *sql.Stmt

	// service
	addService        *sql.Stmt
	deleteServiceByID *sql.Stmt
	getServiceByID    *sql.Stmt
	getAllService     *sql.Stmt
}

type Store interface {
	Close() error

	// keyvalue
	AddKeyValue(ctx context.Context, arg KeyValue) (KeyValue, error)
	UpdateKeyValue(ctx context.Context, arg KeyValue) (KeyValue, error)
	GetKeyValue(ctx context.Context, key string) (KeyValue, error)
	DeleteKeyValue(ctx context.Context, key string) error

	// uptime watch request
	AddUptimeWatchRequest(ctx context.Context, arg AddUptimeWatchRequestParams) (UptimeWatchRequest, error)
	GetUptimeWatchRequestByID(ctx context.Context, id int) (UptimeWatchRequest, error)
	GetAllUptimeWatchRequest(ctx context.Context) ([]UptimeWatchRequest, error)
	DeleteUptimeWatchRequestById(ctx context.Context, id int) error
	UpdateUptimeWatchRequestById(ctx context.Context, updateData map[string]interface{}, id int) (UptimeWatchRequest, error)

	// uptime result
	AddUptimeResult(ctx context.Context, arg AddUptimeResultParams) (UptimeResult, error)
	GetUptimeResultCount(ctx context.Context, UWRID int) (int, error)
	GetUptimeResults(ctx context.Context, arg GetUptimeResultsParams) ([]UptimeResult, error)
	DeleteUptimeResults(ctx context.Context, UWRID int) error
	GetUptimeResultStatsForID(ctx context.Context, UWRID int) (UptimeResultStats, error)

	// uptime conclusion
	AddUptimeConclusion(ctx context.Context, arg UptimeConclusion) (UptimeConclusion, error)
	DeleteUptimeConclusionByUWRID(ctx context.Context, UWRID int) error
	GetUptimeConclusionByUWRID(ctx context.Context, UWRID int) (UptimeConclusion, error)
	GetAllUptimeConclusion(ctx context.Context, arg getAllUptimeConclusionParams) ([]UptimeConclusion, error)

	// ssl info
	AddUptimeSSLInfo(ctx context.Context, arg UptimeSSLInfo) (UptimeSSLInfo, error)
	DeleteUptimeSSLInfoByUWRID(ctx context.Context, UWRID int) error
	UpdateUptimeSSLInfoByUWRID(ctx context.Context, arg UptimeSSLInfo) (UptimeSSLInfo, error)
	GetUptimeSSLInfoByUWRID(ctx context.Context, UWRID int) (UptimeSSLInfo, error)
	GetAllUptimeSSLInfo(ctx context.Context, arg getAllUptimeSSLInfoParams) ([]UptimeSSLInfo, error)

	// service
	AddService(ctx context.Context, arg Service) (Service, error)
	DeleteServiceByID(ctx context.Context, ID string) error
	GetServiceByID(ctx context.Context, ID string) (Service, error)
	GetAllService(ctx context.Context) ([]Service, error)
}
