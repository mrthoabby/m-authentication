package interfaces

type IRepository[T any] interface {
	ValideBasicAuth(user, password string) bool
	GetClaims(data T, tableName string) (T, error)
}
