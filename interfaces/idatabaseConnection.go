package interfaces

//Adapter pattern for database connection
type IDatabaseConnection[T, R any] interface {
	Connect(connectionString string) (bool, error)
	Disconnect()
	Query(query T) R
}
