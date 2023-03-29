package database

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

type transactor struct {
	db *gorm.DB
}

func (t *transactor) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	txCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx := t.db.WithContext(txCtx).Begin()

	// run callback
	if err := tFunc(injectTx(ctx, tx)); err != nil {
		// if error, rollback
		if err := tx.Rollback().Error; err != nil {
			log.Printf("cannot rollback transaction: %v", err)
		}

		return err
	}

	// if no error, commit
	if err := tx.Commit().Error; err != nil {
		log.Printf("cannot commit transaction: %v", err)
	}

	return nil
}

// tx returns context with or without transaction extracted from context
func (t *transactor) tx(ctx context.Context) *gorm.DB {
	if tx := extractTx(ctx); tx != nil {
		return tx.WithContext(ctx)
	}

	return t.db.WithContext(ctx)
}
