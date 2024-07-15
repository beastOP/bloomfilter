package bloomfilter

import (
	"fmt"
	"hash/fnv"
	"math"
)

type Bloom struct {
	fp_prob    float64
	size       uint64
	hash_count int
	bit_array  []bool
}

func New(item_count int, fp_prob float64) (*Bloom, error) {
	if item_count <= 0 {
		return nil, fmt.Errorf("item_count must be greater than 0, got %d", item_count)
	}
	if fp_prob <= 0 || fp_prob >= 1 {
		return nil, fmt.Errorf("fp_prob must be between 0 and 1, got %v", fp_prob)
	}

	// Calculate the size of bit array(m) to used using following formula
	// m = -(n * lg(p)) / (lg(2)^2)
	// n : int
	//     number of items expected to be stored in filter
	// p : float
	//     False Positive probability in decimal
	size := -(float64(item_count) * math.Log(fp_prob)) / math.Pow(math.Log(2), 2)

	// Calculate the hash function(k) to be used using following formula
	// k = (m/n) * lg(2)
	// m : int
	//     size of bit array
	// n : int
	//     number of items expected to be stored in filter
	hash_count := int(size / float64(item_count) * math.Log(2))
	return &Bloom{
		fp_prob:    fp_prob,
		size:       uint64(size),
		hash_count: hash_count,
		bit_array:  make([]bool, uint64(size)),
	}, nil
}

func (b *Bloom) hash(item string, i int) uint64 {
	hash1 := fnv.New64()
	hash2 := fnv.New64a()
	hash1.Write([]byte(item))
	hash2.Write([]byte(item))
	return (hash1.Sum64() + uint64(i)*hash2.Sum64()) % b.size
}

func (b *Bloom) Add(item string) {
	for i := 0; i < b.hash_count; i++ {
		pos := b.hash(item, i)
		b.bit_array[pos] = true
	}
}

func (b *Bloom) Contains(item string) bool {
	for i := 0; i < b.hash_count; i++ {
		pos := b.hash(item, i)
		if !b.bit_array[pos] {
			return false
		}
	}
	return true
}
