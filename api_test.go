package gonumberio_test

import (
	"fmt"
	"math"
	"os"
	r "reflect"
	"testing"

	nio "github.com/Matej-Chmel/go-number-io"
	ite "github.com/Matej-Chmel/go-number-io/internal"
)

const (
	epsilon float64 = 0.000001
)

var (
	bool1D = []bool{false, true, false, false, true}

	bool2D = [][]bool{
		{false, true},
		{true},
		{false, false, false, false, false, false},
		{true, true},
	}

	bool3D = [][][]bool{
		{
			{true},
			{false, false, false, false, false, false},
			{true, true},
		},
		{
			{false, true, false, true},
			{false, true, false, true},
			{false, true, false, true},
		},
		{
			{false, true},
			{false, true, false, true, true},
			{false, false, true, true, true},
			{false, true, true},
			{true},
		},
	}

	float1D = []float32{1.12, 3.45, -3.90, .23, -.23, .201, 0, 0, 0}

	float2D = [][]float32{
		{3.4, 5.6, -201.2},
		{1.1, .231, -23.},
		{90.001, 9000.1, -01000.1},
	}

	float3D = [][][]float32{
		{
			{1, 1, 1},
			{1.1, 1.1, 1.1},
			{.001, .001, .001},
		},
		{
			{-.2, -.1, -.2, -.2},
			{.001, .901, .89},
			{240001.432, 3456.989, .908757},
			{.232, .451, .23, .123, .12},
		},
		{
			{0.0, 1.},
			{1, 2.},
		},
	}

	int1D = []int32{1, -2, 3, -4, 5, 0}

	int2D = [][]int32{
		{1, 2, -3},
		{4, -5, 6},
		{-7, -8, -9},
		{0, 0},
	}

	int3D = [][][]int32{
		{
			{-1, 2, 3},
			{0, 0, 0},
		},
		{
			{-1, -2, -3},
			{-4, -5, -6},
			{7, 8, 9},
		},
		{
			{0},
		},
		{
			{100, 10000},
			{-20132, -2121},
			{-3000, 10300, 12001, 14001},
			{9091, 8091, 17003},
			{90123},
		},
	}

	uint1D = []uint32{1, 2, 3, 4, 5}

	uint2D = [][]uint32{
		{9, 0, 90},
		{89, 78},
		{102, 0, 0, 0, 3022, 19283},
		{18293},
	}

	uint3D = [][][]uint32{
		{
			{0, 1, 2, 3, 4},
			{5, 6, 7, 7},
			{7, 9, 0, 8},
			{8, 7, 7},
			{1},
		},
		{
			{1, 2, 3},
			{4, 5},
			{6},
		},
		{
			{8, 9},
			{0, 0, 1},
			{2, 3, 4, 5},
			{10, 11, 12, 13},
			{14},
		},
	}
)

func check1D[T SliceItem](err error, a, b []T, cmp func(T, T) bool, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare1D(a, b, cmp) {
		throw(a, b, t)
	}
}

func check2D[T SliceItem](err error, a, b [][]T, cmp func(T, T) bool, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare2D(a, b, cmp) {
		throw(a, b, t)
	}
}

func check3D[T SliceItem](err error, a, b [][][]T, cmp func(T, T) bool, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare3D(a, b, cmp) {
		throw(a, b, t)
	}
}

func compare1D[T SliceItem](a1, b1 []T, cmp func(T, T) bool) bool {
	if len(a1) != len(b1) {
		return false
	}

	for t1 := 0; t1 < len(a1); t1++ {
		if !cmp(a1[t1], b1[t1]) {
			return false
		}
	}

	return true
}

func compare2D[T SliceItem](a2, b2 [][]T, cmp func(T, T) bool) bool {
	if len(a2) != len(b2) {
		return false
	}

	for t2 := 0; t2 < len(a2); t2++ {
		if val := compare1D(a2[t2], b2[t2], cmp); !val {
			return false
		}
	}

	return true
}

func compare3D[T SliceItem](a3, b3 [][][]T, cmp func(T, T) bool) bool {
	if len(a3) != len(b3) {
		return false
	}

	for t3 := 0; t3 < len(a3); t3++ {
		if val := compare2D(a3[t3], b3[t3], cmp); !val {
			return false
		}
	}

	return true
}

func epsilonCompare[T ite.Number](a, b T) bool {
	diff := float64(a - b)
	return math.Abs(diff) <= epsilon
}

func equals[T SliceItem](a, b T) bool {
	return a == b
}

func getEpsilonCmp[T any]() interface{} {
	switch kind := ite.GetTypeKind[T](); kind {
	case r.Float32:
		return epsilonCompare[float32]
	case r.Float64:
		return epsilonCompare[float64]
	}

	return nil
}

func getExpected[T SliceItem]() ([]T, [][]T, [][][]T) {
	var i1D interface{}
	var i2D interface{}
	var i3D interface{}

	switch kind := ite.GetTypeKind[T](); kind {
	case r.Bool:
		i1D, i2D, i3D = bool1D, bool2D, bool3D
	case r.Float32:
		i1D, i2D, i3D = float1D, float2D, float3D
	case r.Int32:
		i1D, i2D, i3D = int1D, int2D, int3D
	case r.Uint32:
		i1D, i2D, i3D = uint1D, uint2D, uint3D
	}

	r1D, _ := i1D.([]T)
	r2D, _ := i2D.([][]T)
	r3D, _ := i3D.([][][]T)
	return r1D, r2D, r3D
}

func getTypeName[T SliceItem]() string {
	switch kind := ite.GetTypeKind[T](); kind {
	case r.Bool:
		return "bool"
	case r.Float32:
		return "float"
	case r.Float64:
		return "float64"
	case r.Int8:
		return "int8"
	case r.Int16:
		return "int16"
	case r.Int32:
		return "int"
	case r.Int64:
		return "int64"
	case r.Uint8:
		return "uint8"
	case r.Uint16:
		return "uint16"
	case r.Uint32:
		return "uint"
	case r.Uint64:
		return "uint64"
	}

	return "unknown"
}

type file struct {
	*os.File
	error
}

func openFile(dim uint, typeName string) file {
	filePath := fmt.Sprintf("data/%s/%dD.txt", typeName, dim)
	data, err := os.Open(filePath)
	return file{File: data, error: err}
}

func openFiles[T SliceItem]() (res [3]file) {
	typeName := getTypeName[T]()
	res[0] = openFile(1, typeName)
	res[1] = openFile(2, typeName)
	res[2] = openFile(3, typeName)
	return
}

type SliceItem interface {
	bool | ite.Number
}

func throw[T any](a, b T, t *testing.T) {
	t.Errorf("\n\n%v\n\n!=\n\n%v", a, b)
}

func runTest[T SliceItem](t *testing.T) {
	e1D, e2D, e3D := getExpected[T]()
	files := openFiles[T]()

	defer func() {
		for _, v := range files {
			v.Close()
		}
	}()

	var cmp func(T, T) bool = equals[T]
	eps, ok := getEpsilonCmp[T]().(func(T, T) bool)

	if eps != nil && ok {
		cmp = eps
	}

	a1D, err := nio.Read1D[T](files[0])
	check1D(err, a1D, e1D, cmp, t)

	a2D, err := nio.Read2D[T](files[1])
	check2D(err, a2D, e2D, cmp, t)

	a3D, err := nio.Read3D[T](files[2])
	check3D(err, a3D, e3D, cmp, t)
}

func TestBool(t *testing.T) {
	runTest[bool](t)
}

func TestFloat32(t *testing.T) {
	runTest[float32](t)
}

func TestInt32(t *testing.T) {
	runTest[int32](t)
}

func TestUint32(t *testing.T) {
	runTest[uint32](t)
}
