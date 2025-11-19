package db

import (
	"context"
	"testing"
	"time"

	"github.com/chen/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry{
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}


func TestCreateEntry(t *testing.T){
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}


func TestGetEntry(t *testing.T){
	account := createRandomAccount(t)
	arg := createRandomEntry(t, account)

	entry, err := testQueries.GetEntry(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.ID, arg.ID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.WithinDuration(t, entry.CreatedAt, arg.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T){
	account := createRandomAccount(t)
	for i:=0; i<10; i++{
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries{
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
	
}


func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := createRandomEntry(t, account)

	arg1 := UpdateentriesParams{
		ID: arg.ID,
		Amount: arg.Amount,
	}


	entry, err := testQueries.Updateentries(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.ID, arg1.ID)
	require.Equal(t, entry.Amount, arg1.Amount)
}