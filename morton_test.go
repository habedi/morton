package morton

import (
	"fmt"
	"testing"
)

func TestUInt32ToBitArray(t *testing.T) {
	var a uint32 = 0b10000001 // 129
	// Expected bit array for 129 of length 32 (uint32)
	var expected = []uint8{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 1}

	var result = uInt32ToBitArray(a)
	// Check if the two bit arrays are equal
	for i := 0; i < 32; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Result: %v\n", result)
}

func TestUInt64ToBitArray(t *testing.T) {
	var a uint64 = 0b10000001 // 129
	// Expected bit array for 129 of length 64 (uint64)
	var expected = []uint8{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 1}

	var result = uInt64ToBitArray(a)
	// Check if the two bit arrays are equal
	for i := 0; i < 64; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Result: %v\n", result)
}

func TestInterleave2D(t *testing.T) {
	var a uint32 = 0b10000001 // 129
	var b uint32 = 0b01000000 // 64
	// Expected bit array for 129 and 64 of length 32 (uint32)
	var expected = []uint8{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0}

	var result = interleave2D(uInt32ToBitArray(a), uInt32ToBitArray(b))
	// Check if the two bit arrays are equal
	for i := 0; i < 32; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected value at index %d to be %d, got %d", i, expected[i], result[i])
		}
	}
	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Result: %v\n", result)
}

func TestInterleave3D(t *testing.T) {

	var a uint32 = 0b10000001 // 129
	var b uint32 = 0b01000000 // 64
	var c uint32 = 0b00100000 // 32

	// Expected bit array for 129, 64, and 32 of length 21 (uint32)
	var expected = []uint8{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 1, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0}

	var result = interleave3D(uInt32ToBitArray(a), uInt32ToBitArray(b), uInt32ToBitArray(c))
	// Check if the two bit arrays are equal
	for i := 0; i < 21; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected value at index %d to be %d, got %d", i, expected[i], result[i])
		}
	}

	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Result: %v\n", result)
}

func TestFromBitArray(t *testing.T) {
	var a = 0b10000001 // 129
	var expected = uint64(a)

	var intermediateResult = uInt32ToBitArray(uint32(a))

	fmt.Printf("Intermediate result: %v\n", intermediateResult)

	var result = fromBitArray(intermediateResult)

	fmt.Printf("Result: %d\n", uInt32ToBitArray(uint32(result)))

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("Result: %d\n", result)
}

func TestEncode2D(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 6
	var result = Encode2D(uint32(a), uint32(b))
	if result != uint64(expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("Result: %d\n", result)
}

func TestDecode2D(t *testing.T) {
	var encoded = 6
	var expectedX = 1
	var expectedY = 2
	var x, y = Decode2D(uint64(encoded))
	if x != uint32(expectedX) || y != uint32(expectedY) {
		t.Errorf("Expected (%d, %d), got (%d, %d)", expectedX, expectedY, x, y)
	}
	fmt.Printf("Expected: (%d, %d)\n", expectedX, expectedY)
	fmt.Printf("Result: (%d, %d)\n", x, y)
}

func TestEncode3D(t *testing.T) {
	var a uint32 = 1
	var b uint32 = 2
	var c uint32 = 3
	var expected uint64 = 29
	var result = Encode3D(a, b, c)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("Result: %d\n", result)
}

func TestDecode3D(t *testing.T) {
	var encoded uint64 = 31
	var expectedX uint32 = 1
	var expectedY uint32 = 3
	var expectedZ uint32 = 3
	var x, y, z = Decode3D(encoded)
	if x != expectedX || y != expectedY || z != expectedZ {
		t.Errorf("Expected (%d, %d, %d), got (%d, %d, %d)", expectedX, expectedY, expectedZ, x, y, z)
	}
	fmt.Printf("Expected: (%d, %d, %d)\n", expectedX, expectedY, expectedZ)
	fmt.Printf("Result: (%d, %d, %d)\n", x, y, z)
}

// Examples from the README

func TestMyMain(t *testing.T) {

	// Encode and decode (10, 13)
	demo2D(10, 13)

	// Encode and decode (10, 13, 5)
	demo3D(10, 13, 5)
}

func demo2D(x uint32, y uint32) {
	code := Encode2D(x, y)
	fmt.Printf("Morton code of (%d, %d) is %d\n", x, y, code)

	x, y = Decode2D(code)
	fmt.Printf("Decoded Morton code %d to (%d, %d)\n", code, x, y)
}

func demo3D(x uint32, y uint32, z uint32) {
	code := Encode3D(x, y, z)
	fmt.Printf("Morton code of (%d, %d, %d) is %d\n", x, y, z, code)

	x, y, z = Decode3D(code)
	fmt.Printf("Decoded Morton code %d to (%d, %d, %d)\n", code, x, y, z)
}
