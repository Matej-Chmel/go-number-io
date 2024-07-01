package internal

import r "reflect"

type DescendInfo struct {
	Dimensions  uint
	ElementType r.Type
}

func Descend[T any]() DescendInfo {
	aType := r.TypeOf((*T)(nil)).Elem()
	return DynamicDescend(aType)
}

func DynamicDescend(aType r.Type) DescendInfo {
	var dims uint = 0
	kind := aType.Kind()

	for kind == r.Slice {
		aType = aType.Elem()
		kind = aType.Kind()
		dims++
	}

	return DescendInfo{
		Dimensions:  dims,
		ElementType: aType,
	}
}

func GetType[T any]() r.Type {
	return r.TypeOf((*T)(nil)).Elem()
}

func GetTypeKind[T any]() r.Kind {
	return GetType[T]().Kind()
}
