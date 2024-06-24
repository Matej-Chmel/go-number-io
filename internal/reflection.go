package internal

import "reflect"

func GetTypeKind[T any]() reflect.Kind {
	return reflect.TypeOf((*T)(nil)).Elem().Kind()
}
