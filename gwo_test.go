package ggwo

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	gwo := New(50, 3)
	gwo.Run()
	for i, wolf := range gwo.HistoryBests() {
		fmt.Printf("%diter %v, fitness = %v\n", i, wolf.Values(), wolf.Fitness())
	}
}
