package interfaces

type Repository[T any] interface {
	ValideBasicAuth(user, password string) bool
	GetClaims(data T, tableName string) (T, error)
}
