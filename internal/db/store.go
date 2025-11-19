package db

import (
	"context"
	"database/sql"
	"fmt"

)

// 定义商店结构体， 可以单独运行数据库查询的所有功能，以及在事务组合运行查询这些功能
type Store struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db),
	}
}
// 通用事务执行器
func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error{
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil{
		return err
	}
	q := New(tx)  // 这里用 tx 构造 Queries，所有操作在事务内
	err = fn(q)
	if err != nil{
		if rbErr := tx.Rollback(); rbErr != nil{
			return fmt.Errorf("Err is %v, RollbackErr is %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// 该结构体包含两个账户转账的所有参数
type TransferTxParams struct{
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTXResult struct{
	Transfer			Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount   Account `json:"to_account"`
	FromEntry     Entry `json:"from_Entry"`
	ToEntry     Entry `json:"to_Entry"`
}

// 采用闭包传回结果
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTXResult, error){
	var result TransferTXResult
	// 具体的事务
	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID:arg.FromAccountID,
			ToAccountID:arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil{
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID:arg.FromAccountID,
			Amount:-arg.Amount,
		})
		if err != nil{
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID:arg.ToAccountID,
			Amount:arg.Amount,
		})
		if err != nil{
			return err
		}

		// Update accounts balanc
		if arg.FromAccountID < arg.ToAccountID{
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		}else{
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
)(accoun1 Account, account2 Account, err error){
			accoun1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID: accountID1,
			Amount: amount1,
		})
		if err != nil{
			return
		}

		account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID: accountID2,
			Amount: amount2,
		})

		return
		
}