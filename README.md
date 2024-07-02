# Go array reader
Generic library for reading 1D, 2D and 3D slices from a text file.

## Example
Suppose that a user wants to read the following text file as a 3D slice of integers.

```none
-1 2 3
0 0 0

-1 -2 -3
-4    -5 -6
7   8 9

0

100 10000
-20132 -2121
-3000           10300 12001 14001
9091        8091 17003
90123
```

There are 2 approaches.

- User specifies the number of dimensions by specifying the function &mdash; `Read1D`, `Read2D` or `Read3D`
- User specifies the whole type as a generic parameter to the function `Read` and lets the library use reflection to determine the number of dimensions

```go
package main

import (
	"fmt"
	"os"

	nio "github.com/Matej-Chmel/go-number-io"
)

func main() {
	file, err := os.Open("data/int/3D.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	// data, err := nio.Read3D[int32](file) // Concrete approach
	data, err := nio.Read[[][][]int32](file) // Dynamic approach

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(data)
}
```

### Output
```none
[[[-1 2 3] [0 0 0]] [[-1 -2 -3] [-4 -5 -6] [7 8 9]] [[0]] [[100 10000] [-20132 -2121] [-3000 10300 12001 14001] [9091 8091 17003] [90123]]]
```

## Custom conversion
The library provides default conversions for `bool`, `byte`, `float`, `int` and `uint` types and their bit specific versions. For other types, user has to specify a conversion function to a reading function with `Custom` suffix.

Conversion function has the following requirements:

- Input is `*nio.ByteReader`
- Output is `(T, uint, error)` where `T` is element of slice
- `uint` in output represents flags, there 2 important ones:
	- `HasNewlineFlag` &mdash; Newline was found
	- `HasValueFlag` &mdash; An element was found

### Input
Suppose that a user wants to read a 2D slice of a custom type that represents a pair of integers. The format is the following.

```none
{1, 1} {2, 2} {3, 3}
{4, 4} {5, 5} {6, 6}
{7, 7} {8, 8} {9, 9}
```

### Code
User specifies a custom conversion function `convertIntPair` and uses methods of `ByteReader` &mdash; `LookAheadFor` and `SkipByte` to process the curly braces.

```go
package main

import (
	"fmt"
	"os"

	nio "github.com/Matej-Chmel/go-number-io"
)

type intPair struct {
	a int
	b int
}

func convertIntPair(r *nio.ByteReader) (intPair, uint, error) {
	if err := r.SkipByte('{'); err != nil {
		return intPair{}, 0, err
	}

	firstInt, err := getIntHelper(r)

	if err != nil {
		return intPair{}, 0, err
	}

	if err := r.SkipByte(','); err != nil {
		return intPair{}, 0, err
	}

	secondInt, err := getIntHelper(r)

	if err != nil {
		return intPair{}, 0, err
	}

	if err := r.SkipByte('}'); err != nil {
		return intPair{}, 0, err
	}

	flags := nio.HasValueFlag
	isNewLine, err := r.LookAheadFor('{')

	if isNewLine {
		flags |= nio.HasNewlineFlag
	}

	return intPair{a: firstInt, b: secondInt}, flags, err
}

func getIntHelper(r *nio.ByteReader) (int, error) {
	data, _, err := nio.ConvertSigned[int](r)

	if err != nil {
		return 0, err
	}

	return data, nil
}

func main() {
	file, err := os.Open("data/intPair/2D.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	data, err := nio.Read2DCustom(file, nio.DefaultChunkSize, convertIntPair)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(data)
}
```

### Output
```none
[[{1 1} {2 2} {3 3}] [{4 4} {5 5} {6 6}] [{7 7} {8 8} {9 9}]]
```
