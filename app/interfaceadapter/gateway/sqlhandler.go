package gateway

type SqlHandler interface {
	Query(string, ...interface{}) (Row, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
