package main

import (
	"reflect"
)

<<<<<<< Updated upstream
//Reverse reverses a slice.
func Reverse(slice interface{}) {
	n := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
=======
// Parallelsum does parallel vector sum,
// in each loop,buffered input capacity will be cut half
// for the next loop to Sum goroutine to consume.
// it terminates until the input remains one.
func ParallelSum(slcs ...[]int) []int {
	input := make(chan []int, len(slcs))
	output := make(chan []int)
	var result []int
	go func(input chan []int) {
		for _, slc := range slcs {
			input <- slc
		}
		close(input)
	}(input)

	for {
		var wg sync.WaitGroup
		wg.Add(cap(input) / 2)
		for i := 0; i < cap(input)/2; i++ {
			out := Sum(input)
			go func() {
				defer wg.Done()
				for o := range out {
					output <- o
				}
			}()
		}
		go func(output chan []int) {
			wg.Wait()
			close(output)
		}(output)

		input = make(chan []int, cap(input)/2)
		if cap(input) < 2 {
			result = <-output
			break
		}
		for o := range output {
			input <- o
		}
		output = make(chan []int)
		close(input)
>>>>>>> Stashed changes
	}
}

<<<<<<< Updated upstream
func main() {
=======
//TODO:complete the Sum for the parallel sum function.
func Sum(sum chan []int) (output chan []int) {
	out := make(chan []int, cap(sum))
	go func() {
		total := <-sum
		for t := range sum {
			for i := range total {
				total[i] += t[i]
			}
		}
		out <- total
	}()
	return out
}

func main() {
	fmt.Println(`Please edit main.go,and complete the 'Sum' function for the parallel sum to pass the test.
		Concurrency is the most important feature of Go,and the principle is
		'Do not communicate by sharing memory; instead, share memory by communicating.'
		In this exercise you need to catch many features of channels.This is a tour for you to figure out!
		Because here the focus is pipleline model (link:http://blog.golang.org/pipelines).
		It's different from the custom parallel vector sum in which sum numbers at every index of the vectors in a goroutine.
		In this exercies,vector is just a abstract,you can change it to a struct or any thing else that can be sumed up.
	`)
>>>>>>> Stashed changes
}
