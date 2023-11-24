package interfaces

//Adapter pattern for database connection
type DatabaseConnection[T, R any] interface {
	Connect(connectionString string) (bool, error)
	Disconnect()
	Query(query T) R
}
