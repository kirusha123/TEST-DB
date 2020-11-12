package DB

//DBconf . . .
type DBconf struct {
	host   string
	port   int
	user   string
	pass   string
	dbname string
}

//NewDBconfig config for DB
func NewDBconfig() *DBconf {
	return &DBconf{
		host:   "localhost",
		port:   5432,
		user:   "postgres",
		pass:   "admin",
		dbname: "PSQLtest",
	}

}

type database struct {
	cfg *DBconf
}

//NewDB sad...
func NewDB(config *DBconf) *database {
	return &database{
		cfg: config,
	}
}
