package learning_golang

import (
	"sync"
)

var wg sync.WaitGroup

func init() {
	wg.Add(5000)
}

func a() {
	defer wg.Done()
	// TODO
}

/*
	Q2.
		Client facing service - 10K req/s

	Client -> Service
		-> Respond to request
		-> Async job (fetch profileID from request to do this job)
		Response doesn''t depend on async job

	Problem statement:
		Mimic it in golang
		Don’t overload the resources ~ limit the resource usage ->
		Respond the call as soon as you get the API call and fire the job
		X : No external queues or DBs

	Possible Solution:
		1. Using SQS queues and then popping and processing. (DL queues) -> only 10 workers will work at a time and resources are not choked -> resource usage can be predicted
		2. Solution 1 can be implemented using Redis

*/

func someAPI() {

	/*
		-> 5K req/s is supported and go routines will not cause any problems
		-> We have to control the num of parallel go routines running at a single time


		go a()

		return jsonResponse

	*/
	wg.Add(50)
}

func fireGoRoutine() {
	/*
		1. Check if we can fire
		2. If we cannot we wait
		3. Fire if we can
	*/

	wg.Wait() // waits for all go routines to end
	go a()    // if we can fire, add a new go routine

	// fire a new go routine

}
