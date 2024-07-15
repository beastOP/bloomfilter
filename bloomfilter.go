// Package bloomfilter is a lightweight implimentation of bloomfilter algorithm in pure golang.
package bloomfilter

import (
	"fmt"
	"hash/fnv"
	"math"
)

// Bloom represents a Bloom filter data structure.
type Bloom struct {
	fp_prob    float64 // Desired false positive probability.
	size       uint64  // Size of the bit array.
	hash_count int     // Number of hash functions.
	bit_array  []bool  // Bit array representing the Bloom filter.
}

// New creates a new Bloom filter with the specified item count and false positive probability.
// Returns a pointer to the Bloom filter and an error if any.
func New(item_count int, fp_prob float64) (*Bloom, error) {
	if item_count <= 0 {
		return nil, fmt.Errorf("item_count must be greater than 0, got %d", item_count)
	}
	if fp_prob <= 0 || fp_prob >= 1 {
		return nil, fmt.Errorf("fp_prob must be between 0 and 1, got %v", fp_prob)
	}

	// Calculate the size of bit array(m) to used using following formula
	// m = -(n * lg(p)) / (lg(2)^2)
	// n is int
	//      number of items expected to be stored in filter
	// p is float
	//      False Positive probability in decimal
	size := -(float64(item_count) * math.Log(fp_prob)) / math.Pow(math.Log(2), 2)

	// Calculate the hash function(k) to be used using following formula
	// k = (m/n) * lg(2)
	// m is int
	//      size of bit array
	// n is int
	//      number of items expected to be stored in filter
	hash_count := int(size / float64(item_count) * math.Log(2))
	return &Bloom{
		fp_prob:    fp_prob,
		size:       uint64(size),
		hash_count: hash_count,
		bit_array:  make([]bool, uint64(size)),
	}, nil
}

// hash generates a hash value for the given item and hash index/
// Returns the computed hash value as uint64.
func (b *Bloom) hash(item string, i int) uint64 {
	hash1 := fnv.New64()
	hash2 := fnv.New64a()
	hash1.Write([]byte(item))
	hash2.Write([]byte(item))
	return (hash1.Sum64() + uint64(i)*hash2.Sum64()) % b.size
}

// Add inserts an item into the Bloom filter.
func (b *Bloom) Add(item string) {
	for i := 0; i < b.hash_count; i++ {
		pos := b.hash(item, i)
		b.bit_array[pos] = true
	}
}

// Contains checks if an item is in the Bloom filter.
// Returns true if the item is probably in the filter, false otherwise.
func (b *Bloom) Contains(item string) bool {
	for i := 0; i < b.hash_count; i++ {
		pos := b.hash(item, i)
		if !b.bit_array[pos] {
			return false
		}
	}
	return true
}
