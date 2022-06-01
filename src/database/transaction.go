package database

import (
	"context"

	"github.com/jinzhu/gorm"
)

type Transaction interface {
    DoInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}


var txKey = struct{}{}

type tx struct {
    db *gorm.DB
}

func NewTransaction(db *gorm.DB) Transaction {
    return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
    tx := t.db.Begin()
    ctx = context.WithValue(ctx, &txKey, tx)
    v, err := f(ctx)
    if err != nil {
        tx.Rollback()
        return nil, err
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return nil, err
    }
    return v, nil
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
    tx, ok := ctx.Value(&txKey).(*gorm.DB)
    return tx, ok
}
