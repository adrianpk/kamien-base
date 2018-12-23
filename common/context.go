package common

type txContextKey string

const (
	// TxContextKey - Database transaction will ve stored in request context under this key.
	TxContextKey txContextKey = "contextTX"
)
