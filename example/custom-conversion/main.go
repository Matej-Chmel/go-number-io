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
