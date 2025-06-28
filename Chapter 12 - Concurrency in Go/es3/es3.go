package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func generateMap() map[int]float64 {
	mappa := map[int]float64{}

	for i := 0; i < 100_000; i++ {
		mappa[i] = math.Sqrt(float64(i))
	}

	return mappa
}

func main() {
	for i := 0; i < 100_000; i += 1000 {
		mappa := generateMap()
		fmt.Println(mappa[i])
	}

	var generateMapCached func() map[int]float64 = sync.OnceValue(generateMap)

	fmt.Println()
	time.Sleep(time.Second * 2)

	for i := 0; i < 100_000; i += 1000 {
		mappa := generateMapCached()
		fmt.Println(mappa[i])
	}
}
