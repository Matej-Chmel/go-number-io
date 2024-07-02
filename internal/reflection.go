package internal

import r "reflect"

// Result of counting number of dimensions of a generic type
type DescendInfo struct {
	Dimensions  uint
	ElementType r.Type
	Supported   bool
}

// Counts number of dimensions of a generic type
func Descend[T any]() DescendInfo {
	aType := r.TypeOf((*T)(nil)).Elem()
	return DynamicDescend(aType)
}

// Counts number of dimensions of a reflection Type
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
		Supported:   kind != r.Array,
	}
}

// Returns reflection Type corresponding to the specified generic type
func GetType[T any]() r.Type {
	return r.TypeOf((*T)(nil)).Elem()
}

// Returns reflection Kind corresponding to the specified generic type
func GetTypeKind[T any]() r.Kind {
	return GetType[T]().Kind()
}
