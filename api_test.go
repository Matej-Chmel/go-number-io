package gonumberio_test

import (
	"fmt"
	"os"
	"testing"

	nio "github.com/Matej-Chmel/go-number-io"
)

var (
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

func check1D[T comparable](err error, a, b []T, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare1D(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

func check2D[T comparable](err error, a, b [][]T, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare2D(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

func check3D[T comparable](err error, a, b [][][]T, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if !compare3D(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

func compare1D[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func compare2D[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		aa, bb := a[i], b[i]

		if len(aa) != len(bb) {
			return false
		}

		for j := 0; j < len(aa); j++ {
			if aa[j] != bb[j] {
				return false
			}
		}
	}

	return true
}

func compare3D[T comparable](a, b [][][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		aa, bb := a[i], b[i]

		if len(aa) != len(bb) {
			return false
		}

		for j := 0; j < len(aa); j++ {
			aaa, bbb := aa[i], bb[i]

			if len(aaa) != len(bbb) {
				return false
			}

			for k := 0; k < len(aa); k++ {
				if aaa[k] != bbb[k] {
					return false
				}
			}
		}
	}

	return true
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

func Test1DFloat32(t *testing.T) {
	startTest(1, "float", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertFloat[float32])
		check1D(err, actual, float1D, t)
	})
}

func Test2DFloat32(t *testing.T) {
	startTest(2, "float", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertFloat[float32])
		check2D(err, actual, float2D, t)
	})
}

func Test3DFloat32(t *testing.T) {
	startTest(3, "float", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertFloat[float32])
		check3D(err, actual, float3D, t)
	})
}

func Test1DInt32(t *testing.T) {
	startTest(1, "int", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertSigned[int32])
		check1D(err, actual, int1D, t)
	})
}

func Test2DInt32(t *testing.T) {
	startTest(2, "int", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertSigned[int32])
		check2D(err, actual, int2D, t)
	})
}

func Test3DInt32(t *testing.T) {
	startTest(3, "int", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertSigned[int32])
		check3D(err, actual, int3D, t)
	})
}

func Test1DUint32(t *testing.T) {
	startTest(1, "uint", t, func(file *os.File) {
		actual, err := nio.Read1D(file, nio.DefaultChunkSize, nio.ConvertUnsigned[uint32])
		check1D(err, actual, uint1D, t)
	})
}

func Test2DUint32(t *testing.T) {
	startTest(2, "uint", t, func(file *os.File) {
		actual, err := nio.Read2D(file, nio.DefaultChunkSize, nio.ConvertUnsigned[uint32])
		check2D(err, actual, uint2D, t)
	})
}

func Test3DUint32(t *testing.T) {
	startTest(3, "uint", t, func(file *os.File) {
		actual, err := nio.Read3D(file, nio.DefaultChunkSize, nio.ConvertUnsigned[uint32])
		check3D(err, actual, uint3D, t)
	})
}
