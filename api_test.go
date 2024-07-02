package gonumberio_test

import (
	"fmt"
	"io"
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
	boolD1 = []bool{false, true, false, false, true}

	boolD2 = [][]bool{
		{false, true},
		{true},
		{false, false, false, false, false, false},
		{true, true},
	}

	boolD3 = [][][]bool{
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

	floatD1 = []float32{1.12, 3.45, -3.9, .23, -.23, .201, 0., 0., 0.}

	floatD2 = [][]float32{
		{3.4, 5.6, -201.2},
		{1.1, .231, -23.},
		{90.001, 9000.1, -01000.1},
	}

	floatD3 = [][][]float32{
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

	int32D1 = []int32{1, -2, 3, -4, 5, 0}

	int32D2 = [][]int32{
		{1, 2, -3},
		{4, -5, 6},
		{-7, -8, -9},
		{0, 0},
	}

	int32D3 = [][][]int32{
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

	intD1 = []int{1, -2, 3, -4, 5, 0}

	intD2 = [][]int{
		{1, 2, -3},
		{4, -5, 6},
		{-7, -8, -9},
		{0, 0},
	}

	intD3 = [][][]int{
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

	uint32D1 = []uint32{1, 2, 3, 4, 5}

	uint32D2 = [][]uint32{
		{9, 0, 90},
		{89, 78},
		{102, 0, 0, 0, 3022, 19283},
		{18293},
	}

	uint32D3 = [][][]uint32{
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

	uintD1 = []uint{1, 2, 3, 4, 5}

	uintD2 = [][]uint{
		{9, 0, 90},
		{89, 78},
		{102, 0, 0, 0, 3022, 19283},
		{18293},
	}

	uintD3 = [][][]uint{
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

func getEpsilonCmp[T SliceItem]() any {
	switch kind := ite.GetTypeKind[T](); kind {
	case r.Float32:
		return epsilonCompare[float32]
	case r.Float64:
		return epsilonCompare[float64]
	}

	return nil
}

func getExpected0D[T SliceItem]() T {
	var a any

	switch kind := ite.GetTypeKind[T](); kind {
	case r.Int, r.Int32:
		a = 127
	}

	res, _ := a.(T)
	return res
}

func getExpected1D[T SliceItem]() []T {
	var a any

	switch kind := ite.GetTypeKind[T](); kind {
	case r.Bool:
		a = boolD1
	case r.Float32:
		a = floatD1
	case r.Int:
		a = intD1
	case r.Int32:
		a = int32D1
	case r.Uint:
		a = uintD1
	case r.Uint32:
		a = uint32D1
	}

	res, _ := a.([]T)
	return res
}

func getExpected2D[T SliceItem]() [][]T {
	var a any

	switch kind := ite.GetTypeKind[T](); kind {
	case r.Bool:
		a = boolD2
	case r.Float32:
		a = floatD2
	case r.Int:
		a = intD2
	case r.Int32:
		a = int32D2
	case r.Uint:
		a = uintD2
	case r.Uint32:
		a = uint32D2
	}

	res, _ := a.([][]T)
	return res
}

func getExpected3D[T SliceItem]() [][][]T {
	var a any

	switch kind := ite.GetTypeKind[T](); kind {
	case r.Bool:
		a = boolD3
	case r.Float32:
		a = floatD3
	case r.Int:
		a = intD3
	case r.Int32:
		a = int32D3
	case r.Uint:
		a = uintD3
	case r.Uint32:
		a = uint32D3
	}

	res, _ := a.([][][]T)
	return res
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
	case r.Int, r.Int32:
		return "int"
	case r.Int64:
		return "int64"
	case r.Uint8:
		return "uint8"
	case r.Uint16:
		return "uint16"
	case r.Uint, r.Uint32:
		return "uint"
	case r.Uint64:
		return "uint64"
	}

	return "unknown"
}

func openFile[T SliceItem](dim uint) (*os.File, error) {
	filePath := fmt.Sprintf("data/%s/%dD.txt", getTypeName[T](), dim)
	return os.Open(filePath)
}

type SliceItem interface {
	bool | ite.Number
}

type tester[T SliceItem] struct {
	cmp       func(T, T) bool
	dim       uint
	file      *os.File
	isDynamic bool
	t         *testing.T
}

func newTester[T SliceItem](dim uint, isDynamic bool, t *testing.T) (tester[T], error) {
	if dim > 3 {
		var res tester[T]
		return res, fmt.Errorf("%dD slices not supported", dim)
	}

	var cmp func(T, T) bool = equals[T]
	eps, ok := getEpsilonCmp[T]().(func(T, T) bool)

	if eps != nil && ok {
		cmp = eps
	}

	file, err := openFile[T](dim)

	if err != nil {
		var res tester[T]
		return res, err
	}

	return tester[T]{
		cmp:       cmp,
		dim:       dim,
		file:      file,
		isDynamic: isDynamic,
		t:         t,
	}, nil
}

func (t *tester[T]) check0D(getActual func(io.Reader) (T, error)) {
	if actual, err := getActual(t.file); err != nil {
		t.throwError(err)
	} else if expected := getExpected0D[T](); !t.cmp(actual, expected) {
		t.throwNotEqual(actual, expected)
	}
}

func (t *tester[T]) check1D(getActual func(io.Reader) ([]T, error)) {
	if actual, err := getActual(t.file); err != nil {
		t.throwError(err)
	} else if expected := getExpected1D[T](); !compare1D(actual, expected, t.cmp) {
		t.throwNotEqual(actual, expected)
	}
}

func (t *tester[T]) check2D(getActual func(io.Reader) ([][]T, error)) {
	if actual, err := getActual(t.file); err != nil {
		t.throwError(err)
	} else if expected := getExpected2D[T](); !compare2D(actual, expected, t.cmp) {
		t.throwNotEqual(actual, expected)
	}
}

func (t *tester[T]) check3D(getActual func(io.Reader) ([][][]T, error)) {
	if actual, err := getActual(t.file); err != nil {
		t.throwError(err)
	} else if expected := getExpected3D[T](); !compare3D(actual, expected, t.cmp) {
		t.throwNotEqual(actual, expected)
	}
}

func (t *tester[T]) logDynamic() {
	t.t.Logf("Is dynamic: %t", t.isDynamic)
}

func (t *tester[T]) runTest() {
	if t.isDynamic {
		if t.dim == 0 {
			t.check0D(nio.Read[T])
		} else if t.dim == 1 {
			t.check1D(nio.Read[[]T])
		} else if t.dim == 2 {
			t.check2D(nio.Read[[][]T])
		} else if t.dim == 3 {
			t.check3D(nio.Read[[][][]T])
		}
	} else {
		if t.dim == 0 {
			t.check0D(nio.Read0D[T])
		} else if t.dim == 1 {
			t.check1D(nio.Read1D[T])
		} else if t.dim == 2 {
			t.check2D(nio.Read2D[T])
		} else if t.dim == 3 {
			t.check3D(nio.Read3D[T])
		}
	}

	t.file.Close()
}

func (t *tester[T]) throwError(err error) {
	t.logDynamic()
	t.t.Error(err)
}

func (t *tester[T]) throwNotEqual(a, b any) {
	t.logDynamic()
	t.t.Errorf("\n\n%v\n\n!=\n\n%v", a, b)
}

func runAllTests[T SliceItem](t *testing.T) {
	runTest[T](1, false, t)
	runTest[T](1, true, t)
	runTest[T](2, false, t)
	runTest[T](2, true, t)
	runTest[T](3, false, t)
	runTest[T](3, true, t)
}

func runTest[T SliceItem](dim uint, isDynamic bool, t *testing.T) {
	tr, err := newTester[T](dim, isDynamic, t)

	if err != nil {
		t.Error(err)
		return
	}

	tr.runTest()
}

func TestBool(t *testing.T) {
	runAllTests[bool](t)
}

func TestFloat32(t *testing.T) {
	runAllTests[float32](t)
}

func TestInt(t *testing.T) {
	runAllTests[int](t)
}

func TestInt0D(t *testing.T) {
	runTest[int](0, false, t)
	runTest[int](0, true, t)
}

func TestInt32(t *testing.T) {
	runAllTests[int32](t)
}

func TestUint(t *testing.T) {
	runAllTests[uint](t)
}

func TestUint32(t *testing.T) {
	runAllTests[uint32](t)
}
