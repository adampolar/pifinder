package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	//fmt.Printf("%v\n", IterPi())
	//fmt.Printf("%v\n", SmartParallelPi())
	//fmt.Printf("%v\n", ParallelPi())
	//fmt.Printf("%v\n", SmarterParallelPi())
	fmt.Printf("%v\n", ParallelPi10())
}

func IterPi() float64 {
	const numberOfIter = 1000000

	trueCount := 0
	for i := 0; i < numberOfIter; i++ {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
			trueCount++
		}
	}
	return float64(trueCount) / (float64(numberOfIter) / 4.0)
}

func ParallelPi() float64 {
	const numberOfGoRoutines = 1000000

	resultChan := make(chan bool, numberOfGoRoutines)

	wg := sync.WaitGroup{}
	wg.Add(numberOfGoRoutines)

	for i := 0; i < numberOfGoRoutines; i++ {
		go func() {
			defer wg.Done()
			x := rand.Float64()*2 - 1
			y := rand.Float64()*2 - 1
			resultChan <- ((math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1)
		}()
	}
	wg.Wait()
	close(resultChan)
	trueCount := 0
	for v := range resultChan {
		if v {
			trueCount++
		}
	}
	return float64(trueCount) / (float64(numberOfGoRoutines) / 4.0)
}

func SmartParallelPi() float64 {
	numCores := runtime.NumCPU()
	numberOfGoRoutines := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores)

	wg := sync.WaitGroup{}
	wg.Add(numCores)

	for i := 0; i < numCores; i++ {
		go func() {
			defer wg.Done()
			count := 0
			for j := 0; j < (numberOfGoRoutines / numCores); j++ {
				x := rand.Float64()*2 - 1
				y := rand.Float64()*2 - 1
				if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	trueCount := 0
	for v := range resultChan {
		trueCount = trueCount + v
	}
	return float64(trueCount) / (float64(numberOfGoRoutines) / 4.0)
}

func SmarterParallelPi() float64 {
	numCores := runtime.NumCPU()
	numberOfGoRoutines := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores)

	for i := 0; i < numCores; i++ {
		go func() {
			count := 0
			for j := 0; j < (numberOfGoRoutines / numCores); j++ {
				x := rand.Float64()*2 - 1
				y := rand.Float64()*2 - 1
				if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores {
				return float64(trueCount) / (float64(numberOfGoRoutines) / 4.0)
			}
		}
	}
}

func SmartestParallelPi() float64 {
	numCores := runtime.NumCPU()
	numberOfIter := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores)
	randSource := make([]float64, numberOfIter*2)

	for k := 0; k < numberOfIter*2; k++ {
		randSource[k] = rand.Float64()
	}

	for i := 0; i < numCores; i++ {
		i := i
		randSourceSection := i * (numberOfIter / numCores)
		go func() {
			count := 0
			for j := 0; j < (numberOfIter / numCores); j++ {
				x := randSource[randSourceSection+2*j]*2 - 1
				y := randSource[randSourceSection+2*j+1]*2 - 1
				if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores {
				return float64(trueCount) / (float64(numberOfIter) / 4.0)
			}
		}
	}
}

func ParallelPi6() float64 {
	numCores := runtime.NumCPU()
	numberOfIter := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores)
	randSources := make([]*rand.Rand, numCores)

	for i := 0; i < numCores; i++ {
		randSources[i] = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))

		i := i
		go func() {
			count := 0
			for j := 0; j < (numberOfIter / numCores); j++ {
				x := randSources[i].Float64()*2 - 1
				y := randSources[i].Float64()*2 - 1
				if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores {
				return float64(trueCount) / (float64(numberOfIter) / 4.0)
			}
		}
	}
}

func ParallelPi7() float64 {
	numCores := runtime.NumCPU()
	numberOfIter := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores)
	randSources := make([]*rand.Rand, numCores)

	for i := 0; i < numCores; i++ {
		randSources[i] = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))

		i := i
		go func() {
			count := 0
			for j := 0; j < (numberOfIter / numCores); j++ {
				x := randSources[i].Float64()*2 - 1
				y := randSources[i].Float64()*2 - 1
				if (math.Sqrt(x*x + y*y)) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores {
				return float64(trueCount) / (float64(numberOfIter) / 4.0)
			}
		}
	}
}

func ParallelPi8() float32 {
	numCores := runtime.NumCPU()
	numberOfIter := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores)
	randSources := make([]*rand.Rand, numCores)

	for i := 0; i < numCores; i++ {
		randSources[i] = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))

		i := i
		go func() {
			count := 0
			for j := 0; j < (numberOfIter / numCores); j++ {
				x := randSources[i].Float32()*2 - 1
				y := randSources[i].Float32()*2 - 1
				if math.Sqrt(float64(x*x+y*y)) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores {
				return float32(trueCount) / (float32(numberOfIter) / 4.0)
			}
		}
	}
}

func ParallelPi9() float64 {
	numCores := int64(runtime.NumCPU())
	numberOfIter := int64((1000000 / numCores) * numCores)

	resultChan := make(chan int, numCores)
	var i int64
	for i = 0; i < numCores; i++ {
		i := i
		go func() {
			count := 0
			randSource := rand.New(rand.NewSource(time.Now().UnixNano() + i))
			var j int64
			for j = 0; j < (numberOfIter / numCores); j++ {
				x := randSource.Float64()*2 - 1
				y := randSource.Float64()*2 - 1
				if (math.Sqrt(x*x + y*y)) < 1 {
					count++
				}
			}
			resultChan <- count
		}()
	}
	trueCount := 0
	var doneCount int64
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores {
				return float64(trueCount) / (float64(numberOfIter) / 4.0)
			}
		}
	}
}

func ParallelPi10() float64 {
	numCores := int64(runtime.NumCPU())
	numberOfIter := int64((10000000000 / numCores) * numCores)

	results := make([]int, numCores)

	wg := sync.WaitGroup{}
	wg.Add(int(numCores))
	var i int64
	for i = 0; i < numCores; i++ {
		i := i
		go func() {
			count := 0
			randSource := rand.New(rand.NewSource(time.Now().UnixNano() + i))
			var j int64
			for j = 0; j < (numberOfIter / numCores); j++ {
				x := randSource.Float64()*2 - 1
				y := randSource.Float64()*2 - 1
				if (math.Sqrt(x*x + y*y)) < 1 {
					count++
				}
			}
			results[i] = count
			wg.Done()
		}()
	}
	wg.Wait()
	trueCount := 0
	for _, v := range results {
		trueCount += v
	}
	return float64(trueCount) / (float64(numberOfIter) / 4.0)
}

type PairOfFloats struct {
	x float64
	y float64
}

func EvenSmarterParallelPi() float64 {
	numCores := runtime.NumCPU()
	numberOfIter := (1000000 / numCores) * numCores

	resultChan := make(chan int, numCores-1)

	randChan := make(chan *PairOfFloats, numberOfIter) //healthy buffer
	go func(randChan chan<- *PairOfFloats) {
		for k := 0; k < numberOfIter; k++ {
			randChan <- &PairOfFloats{x: rand.Float64(), y: rand.Float64()}
		}
		close(randChan)
	}(randChan)

	for i := 0; i < numCores-1; i++ {
		go func(incomingRandStream <-chan *PairOfFloats) {
			count := 0
			func() {
				for {
					select {
					case r, anyLeft := <-incomingRandStream:

						if !anyLeft {
							return
						}
						x := (*r).x*2 - 1
						y := (*r).y*2 - 1
						if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
							count++
						}
					}
				}
			}()
			resultChan <- count
		}(randChan)
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores-1 {
				return float64(trueCount) / (float64(numberOfIter) / 4.0)
			}
		}
	}
}

func ParallelPi5() float64 {
	numCores := runtime.NumCPU()
	numberOfIter := (1000000 / (numCores - 1)) * (numCores - 1)

	resultChan := make(chan int, numCores-1)

	coresForIter := numCores - 1
	chansForIter := make([]*chan *PairOfFloats, coresForIter)

	for i := 0; i < coresForIter; i++ {
		a := make(chan *PairOfFloats)
		chansForIter[i] = &a
	}

	for _, v := range chansForIter {
		go func(randChan chan<- *PairOfFloats) {
			for k := 0; k < numberOfIter/(numCores-1); k++ {
				randChan <- &PairOfFloats{x: rand.Float64(), y: rand.Float64()}
			}
			close(randChan)
		}(*v)
		go func(incomingRandStream <-chan *PairOfFloats, resultChan chan<- int) {
			count := 0
			func() {
				for {
					select {
					case r, anyLeft := <-incomingRandStream:

						if !anyLeft {
							return
						}
						x := (*r).x*2 - 1
						y := (*r).y*2 - 1
						if (math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))) < 1 {
							count++
						}
					}
				}
			}()
			resultChan <- count
		}(*v, resultChan)
	}
	trueCount := 0
	doneCount := 0
	for {
		select {
		case v := <-resultChan:
			doneCount++
			trueCount = trueCount + v
			if doneCount >= numCores-1 {
				return float64(trueCount) / (float64(numberOfIter) / 4.0)
			}
		}
	}
}
