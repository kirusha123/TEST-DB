package db

type DB struct {
	cfg *DBconf
}

func NewDB(config *DBconf) *DB {
	return &DB{
		cfg: config,
	}
}
