package interfaces

type DatabaseConnection interface {
	Connect()
	Disconnect()
	Query()
}
