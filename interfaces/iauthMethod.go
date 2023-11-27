package interfaces

import "reflect"

// IAuthMethod is an interface that represents an authentication method.
type IAuthMethod interface {
	GetType() reflect.Type
}
