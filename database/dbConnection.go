package database

type DBConnection interface {
	Connect() DBConnection
	Insert(...interface{})
	Disconnect()
}
