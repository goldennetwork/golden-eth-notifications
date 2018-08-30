package ethNotification

import "errors"

var (
	ErrTransactionNotFound = errors.New("Transaction not found")
	ErrBlockNotFound       = errors.New("Block not found")
)
