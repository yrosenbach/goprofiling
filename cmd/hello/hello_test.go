package main

import "testing"

var input = "hi hello HellO helloo HELO helloo hi hello HellO helloo HELO helloo hi hello HellO helloo HELO helloo hi hello HellO helloo HELO helloo hi hello HellO helloo HELO helloo hi hello HellO helloo HELO helloo"

var output string

func BenchmarkHelloToWorldAlg1(b *testing.B) {
	var result string

	for n := 0; n < b.N; n++ {
		result = HelloToWorld(input)
	}

	output = result
}
