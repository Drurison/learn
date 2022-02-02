package main

import (
	"fmt"
	"math/rand"
	"time"
)

type CustomGenerator struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (g *CustomGenerator) Generate() bool {
	if g.remaining == 0 {
		g.cache, g.remaining = g.src.Int63(), 63
	}

	result := g.cache&0x01 == 1
	g.cache >>= 1
	g.remaining--

	return result
}

func main() {
	r := &CustomGenerator{src: rand.NewSource(time.Now().UnixNano())}
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Print(r.Generate(), " ")
	}
}
