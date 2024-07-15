# BloomFilter

[![Build Status](https://github.com/beastop/bloomfilter/actions/workflows/go.yml/badge.svg)](https://github.com/beastop/bloomfilter/actions)
[![GoDoc](https://pkg.go.dev/badge/github.com/beastop/bloomfilter.svg)](https://pkg.go.dev/github.com/beastop/bloomfilter)

BloomFilter is a simple and efficient implementation of a Bloom filter in Go.

## Installation

To install the package, run:
```sh
go get github.com/beastop/bloomfilter
```

## Usage

### Creating a Bloom Filter

```go
package main

import (
	"fmt"
	"github.com/beastop/bloomfilter"
)

func main() {
	// Create a new Bloom filter
	bloom, err := bloomfilter.New(100, 0.01)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Add items to the Bloom filter
	bloom.Add("foo")
	bloom.Add("bar")

	// Check if items are in the Bloom filter
	fmt.Println("foo in bloom filter:", bloom.Contains("foo"))
	fmt.Println("baz in bloom filter:", bloom.Contains("baz"))
}
```

## Documentation

For detailed API documentation, please visit the [GoDoc](https://pkg.go.dev/github.com/beastop/bloomfilter).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
