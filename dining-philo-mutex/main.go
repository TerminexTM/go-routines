package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	The Dining philosophers problem is well known in computer science circles.
	Five philosophers, numbered from 0 through 4, live in a house where
	the table is laid for theml each philosopher has their own place at
	the table. Their only difficulty - besides those of philosophy - is
	that the dish served is a very difficult kind of spaghetti which has
	to be eaten with two forks. There are two forks next to each plate, so
	that presents no difficulty. As a consequence, however, this means
	that no two neighbours may be eating simultaneously.
*/

/*
	1) Create a variable for Philosophers dining
	2) create for loop that starts a routine representing each diner
	3) we know we need waitgroups so define wg
	4) we know we need to add = to diners
	5) we know we need to wait until all diners are looped
*/

//constants
const hunger = 3

// variables - philosophers
var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 0 * time.Second
var eatTime = 0 * time.Second
var thinkTime = 0 * time.Second
var finishedEating []string
var orderFinishedMutex sync.Mutex

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Println(philosopher, "is seated")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry")
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right\n", philosopher)

		// Philosopher has both forks and is eating?
		fmt.Println(philosopher, "has both forks, and is eating")
		time.Sleep(eatTime)

		//give philosopher some time to think
		fmt.Println(philosopher, "is thinking.")
		time.Sleep(thinkTime)

		// unlock the mutexes (forks)
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.\n", philosopher)
		time.Sleep(sleepTime)
	}

	//print done message
	fmt.Println(philosopher, "is satisfied")
	time.Sleep(sleepTime)

	//have philosopher leave
	fmt.Println(philosopher, "has left the table")

	//track the order that the philosophers finish eating!
	//look out for hidden race conditions!
	orderFinishedMutex.Lock()
	finishedEating = append(finishedEating, philosopher)
	orderFinishedMutex.Unlock()

}

func main() {
	// print introduction
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("-------------------------------")

	wg.Add(len(philosophers))

	//first left fork to be used
	forkLeft := &sync.Mutex{}
	// spawn one goroutine for each philosopher - because each is trying to eat
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		forkRight := &sync.Mutex{}

		//call routine
		go diningProblem(philosophers[i], forkLeft, forkRight)

		//as you travel down the right fork will be left fork for each philosopher
		forkLeft = forkRight
	}
	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("The Philosophers finished in this order")
	fmt.Println("--------------------------------------")
	for i, philo := range finishedEating {
		fmt.Println(i+1, "...", philo)
	}
}
