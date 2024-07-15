package bloomfilter

import "testing"

func TestBloomNew(t *testing.T) {
	filter, err := New(20, 0.05)
	if err != nil {
		t.Error("Could not create the bloom filter")
	}
	if filter.size != 124 || filter.hash_count != 4 {
		t.Errorf("Filter Size: %v (expected 124), Hash Count: %v (expected 4)", filter.size, filter.hash_count)
	}
}

func TestAddToBloom(t *testing.T) {
	itemCount := 100
	fpProb := 0.01
	bloom, err := New(itemCount, fpProb)
	if err != nil {
		t.Fatalf("Failed to create Bloom filter: %v", err)
	}

	item := "hello"
	bloom.Add(item)

	// Since Bloom filters have a small probability of false positives,
	// this test might not be 100% accurate for all cases.
	// Here we assume that the probability is low enough to rely on.
	if !bloom.Contains(item) {
		t.Errorf("Expected Bloom filter to contain item %s", item)
	}
}

func TestCheckInBloom(t *testing.T) {
	itemCount := 100
	fpProb := 0.01
	bloom, err := New(itemCount, fpProb)
	if err != nil {
		t.Fatalf("Failed to create Bloom filter: %v", err)
	}

	item := "world"
	bloom.Add(item)

	if !bloom.Contains(item) {
		t.Errorf("Expected Bloom filter to contain item %s", item)
	}

	nonExistentItem := "not_in_filter"
	if bloom.Contains(nonExistentItem) {
		t.Errorf("Expected Bloom filter not to contain item %s", nonExistentItem)
	}
}
