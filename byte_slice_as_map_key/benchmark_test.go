package byte_slice_indexed_maps

import (
	"testing"
	"unsafe"
)

const (
	mapSize   = int(^uint(0) >> 1)
	arraySize = 10
)

func initialize() (map[[arraySize]byte]any, map[string]any, [arraySize]byte, any) {
	// allocate a large map to avoid reallocation.
	mBytes := make(map[[arraySize]byte]any, mapSize)
	mString := make(map[string]any, mapSize)

	key := [arraySize]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var value any
	value = nil

	return mBytes, mString, key, value
}

func BenchmarkArrayKeyed(b *testing.B) {
	m, _, key, value := initialize()

	b.ResetTimer()
	bArray(m, key, value, b)
}

func BenchmarkOptimizedStringKeyed(b *testing.B) {
	_, m, key, value := initialize()

	b.ResetTimer()
	bOptimizedString(m, key[:], value, b)
}

func BenchmarkUnsafeStringKeyed(b *testing.B) {
	_, m, key, value := initialize()

	b.ResetTimer()
	bUnsafeString(m, key[:], value, b)
}

func bArray(m map[[arraySize]byte]any, key [arraySize]byte, value any, b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertArrayKey(m, key, value)
		_, _ = retrieveArrayKey(m, key)
		deleteArrayKey(m, key)
	}
}

func bUnsafeString(m map[string]any, key []byte, value any, b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertUnsafeStringKey(m, key, value)
		_, _ = retrieveUnsafeStringKey(m, key)
		deleteUnsafeStringKey(m, key)
	}
}

func bOptimizedString(m map[string]any, key []byte, value any, b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertOptimizedStringKey(m, key, value)
		_, _ = retrieveOptimizedStringKey(m, key)
		deleteOptimizedStringKey(m, key)
	}
}

// optimized string

func retrieveOptimizedStringKey(m map[string]any, key []byte) (any, bool) {
	v, ok := m[string(key)]
	return v, ok
}

func insertOptimizedStringKey(m map[string]any, key []byte, val any) {
	m[string(key)] = val
}

func deleteOptimizedStringKey(m map[string]any, key []byte) {
	delete(m, string(key))
}

// unsafe string

func unsafeByteSliceToString(key []byte) string {
	return *(*string)(unsafe.Pointer(&key))
}

func retrieveUnsafeStringKey(m map[string]any, key []byte) (any, bool) {
	v, ok := m[unsafeByteSliceToString(key)]
	return v, ok
}

func insertUnsafeStringKey(m map[string]any, key []byte, val any) {
	m[unsafeByteSliceToString(key)] = val
}

func deleteUnsafeStringKey(m map[string]any, key []byte) {
	delete(m, unsafeByteSliceToString(key))
}

// byte array

func retrieveArrayKey(m map[[arraySize]byte]any, key [arraySize]byte) (any, bool) {
	v, ok := m[key]
	return v, ok
}

func insertArrayKey(m map[[arraySize]byte]any, key [arraySize]byte, val any) {
	m[key] = val
}

func deleteArrayKey(m map[[arraySize]byte]any, key [arraySize]byte) {
	delete(m, key)
}
