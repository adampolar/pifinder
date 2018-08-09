package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkParallelPi(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi()
	}
}

func BenchmarkSmartParallelPi(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		SmartParallelPi()
	}
}

func BenchmarkSmarterParallelPi(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		SmarterParallelPi()
	}
}

func BenchmarkSmartestParallelPi(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		SmartestParallelPi()
	}
}

func BenchmarkEvenSmarterParallelPi(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		EvenSmarterParallelPi()
	}
}

func BenchmarkParallelPi5(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi5()
	}
}

func BenchmarkParallelPi6(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi6()
	}
}

func BenchmarkParallelPi7(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi7()
	}
}

func BenchmarkParallelPi8(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi8()
	}
}

func BenchmarkParallelPi9(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi9()
	}
}

func BenchmarkParallelPi10(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ParallelPi10()
	}
}

func BenchmarkIterPi(b *testing.B) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		IterPi()
	}
}

func BenchmarkMakingGoRoutineAndWritingToChannel(b *testing.B) {

	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		chan1 := make(chan int, 4)
		for j := 0; j < 4; j++ {
			go func() {
				chan1 <- j
			}()
		}
		i := 0
		func() {
			for {
				select {
				case k := <-chan1:
					k = k
					i++
					if i == 4 {
						return
					}
				}
			}
		}()
	}
}

func BenchmarkMaths(b *testing.B) {

	rand.Seed(time.Now().Unix())

	trueCount := 0
	for i := 0; i < b.N; i++ {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
			trueCount++
		}
	}
}
