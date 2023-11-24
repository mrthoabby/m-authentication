package interfaces

import "reflect"

// AuthMethod is an interface that represents an authentication method.
type AuthMethod interface {
	GetType() reflect.Type
}
