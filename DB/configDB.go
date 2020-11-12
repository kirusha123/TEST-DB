package db

//DBconf . . .
type DBconf struct {
	host   string
	port   int
	user   string
	pass   string
	dbname string
}

//New config for DB
func NewDBconfig() *DBconf {
	return &DBconf{
		host:   "localhost",
		port:   5432,
		user:   "postgres",
		pass:   "admin",
		dbname: "PSQLtest",
	}

}
