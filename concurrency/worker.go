package concurrency

func Worker(id int, jobs <-chan int, results chan <-chan int) {
	// Do process the jobs

	// Send the results in the results channel
}