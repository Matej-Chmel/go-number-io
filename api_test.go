package gonumberio_test

import (
	"fmt"
	"math"
	"os"
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

func check1D[T ite.SliceItem](err error, a, b []T, cmp func(T, T) bool, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare1D(a, b, cmp) {
		throw(a, b, t)
	}
}

func check2D[T ite.SliceItem](err error, a, b [][]T, cmp func(T, T) bool, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare2D(a, b, cmp) {
		throw(a, b, t)
	}
}

func check3D[T ite.SliceItem](err error, a, b [][][]T, cmp func(T, T) bool, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare3D(a, b, cmp) {
		throw(a, b, t)
	}
}

func compare1D[T ite.SliceItem](a1, b1 []T, cmp func(T, T) bool) bool {
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

func compare2D[T ite.SliceItem](a2, b2 [][]T, cmp func(T, T) bool) bool {
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

func compare3D[T ite.SliceItem](a3, b3 [][][]T, cmp func(T, T) bool) bool {
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

func equals[T ite.SliceItem](a, b T) bool {
	return a == b
}

func throw[T any](a, b T, t *testing.T) {
	t.Errorf("\n\n%v\n\n!=\n\n%v", a, b)
}

func startTest(dim uint, typeName string, t *testing.T, f func(*os.File)) {
	filePath := fmt.Sprintf("data/%s/%dD.txt", typeName, dim)
	file, err := os.Open(filePath)

	if err != nil {
		t.Error(err)
		return
	}

	defer file.Close()

	f(file)
}

func Test1DBool(t *testing.T) {
	startTest(1, "bool", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertBool)
		check1D(err, actual, bool1D, equals, t)
	})
}

func Test2DBool(t *testing.T) {
	startTest(2, "bool", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertBool)
		check2D(err, actual, bool2D, equals, t)
	})
}

func Test3DBool(t *testing.T) {
	startTest(3, "bool", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertBool)
		check3D(err, actual, bool3D, equals, t)
	})
}

func Test1DFloat32(t *testing.T) {
	startTest(1, "float", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertFloat[float32])
		check1D(err, actual, float1D, epsilonCompare, t)
	})
}

func Test2DFloat32(t *testing.T) {
	startTest(2, "float", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertFloat[float32])
		check2D(err, actual, float2D, epsilonCompare, t)
	})
}

func Test3DFloat32(t *testing.T) {
	startTest(3, "float", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertFloat[float32])
		check3D(err, actual, float3D, epsilonCompare, t)
	})
}

func Test1DInt32(t *testing.T) {
	startTest(1, "int", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertSigned[int32])
		check1D(err, actual, int1D, equals, t)
	})
}

func Test2DInt32(t *testing.T) {
	startTest(2, "int", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertSigned[int32])
		check2D(err, actual, int2D, equals, t)
	})
}

func Test3DInt32(t *testing.T) {
	startTest(3, "int", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertSigned[int32])
		check3D(err, actual, int3D, equals, t)
	})
}

func Test1DUint32(t *testing.T) {
	startTest(1, "uint", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertUnsigned[uint32])
		check1D(err, actual, uint1D, equals, t)
	})
}

func Test2DUint32(t *testing.T) {
	startTest(2, "uint", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertUnsigned[uint32])
		check2D(err, actual, uint2D, equals, t)
	})
}

func Test3DUint32(t *testing.T) {
	startTest(3, "uint", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertUnsigned[uint32])
		check3D(err, actual, uint3D, equals, t)
	})
}
