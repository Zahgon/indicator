// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Ring represents a ring structure that can be instantiated
// using the NewRing function.
//
// Example:
//
//	ring := helper.NewRing[int](2)
//
//	fmt.Println(ring.Insert(1)) // 0
//	fmt.Println(ring.Insert(2)) // 0
//	fmt.Println(ring.Insert(3)) // 1
//	fmt.Println(ring.Insert(4)) // 2
type Ring[T any] struct {
	buffer []T
	begin  int
	end    int
	empty  bool
}

// NewRing creates a new ring instance with the given size.
func NewRing[T any](size int) *Ring[T] { _ = "STUB: not implemented"; return nil }

// Put inserts the specified value into the ring and returns the
// value that was previously stored at that index.
func (r *Ring[T]) Put(t T) T { _ = "STUB: not implemented"; return *new(T) }

// Get retrieves the available value from the ring buffer. If empty,
// it returns the default value (T) and false.
func (r *Ring[T]) Get() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// At returns the value at the given index.
func (r *Ring[T]) At(index int) T { _ = "STUB: not implemented"; return *new(T) }

// IsEmpty checks if the current ring buffer is empty.
func (r *Ring[T]) IsEmpty() bool {
	_ = "STUB: not implemented"

	// IsFull checks if the current ring buffer is full.
	return false
}

func (r *Ring[T]) IsFull() bool { _ = "STUB: not implemented"; return false }

// nextIndex returns the next index in a ring buffer, wrapping
// around if it reaches the capacity.
func (r *Ring[T]) nextIndex(i int) int { _ = "STUB: not implemented"; return 0 }
