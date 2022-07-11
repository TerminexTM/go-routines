package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder //the data channel holds PizzaOrder types
	quit chan chan error //the quit channel holds a channel of errors
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

//any type of Producer has this function
func (p *Producer) Close() error {
	ch := make(chan error) //means of sending back an error if something goes wrong
	p.quit <- ch           //hands the error channel to the channel of quit in the Producer struct :: STILL DON'T UNDERSTAND THIS!
	return <-ch            //return error if any exists
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++                      //receive number from our count and increment by one to begin the order
	if pizzaNumber <= NumberOfPizzas { //as long as we still have pizzas to be made it will enter the if statement
		delay := rand.Intn(5) + 1                        //setup delay variable
		fmt.Printf("Received order #%d!\n", pizzaNumber) //print order received

		// lets actually make the pizza now
		rnd := rand.Intn(12) + 1 //generate a random number between 1 and 12, if number is less than 5 we fail to make pizza
		msg := ""                //default empty msg string
		success := false         //default set pizza to fail

		if rnd < 5 {
			pizzasFailed++ //if we failed we add to the failed pizza counter
		} else {
			pizzasMade++ //success adds to success counter
		}
		total++ //total always goes up

		fmt.Printf("Making pizza #%d. It will take %d seconds.... \n", pizzaNumber, delay) //trach the pizza you are making
		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second) //how long you wait

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d! ***", pizzaNumber) //failed this way
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The Cook quit while making #%d! ***", pizzaNumber) //failed that way
		} else {
			success = true                                              //success sets to true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber) //gives pizza ready and order number
		}
		//build the object for the pizzaOrder
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		//return it
		return &p
	}

	//this happens because we should be done
	//we only need to return the pizzaNumber because everything else should be done. Value should be 11 if NumberOfPizza set to 10
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzaeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification :: received from quit channel

	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		// try to make a pizza

		// decision structure, did we make the pizza, did something go wrong, did we quit?
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select { //statement only for channels but similar to switch statement in appearance
			//we tried to make a pizza (we sent something to the data channel).
			case pizzaMaker.data <- *currentPizza: //populate the pizzaMaker.data with any data we got from the pointer to currentPizza

			case quitChan := <-pizzaMaker.quit:
				//close channels
				close(pizzaMaker.data)
				close(quitChan)
				return //return nothing, go routine done and will exit
			}
		}
	}
}

func main() {
	// seed the random number generator :: need to use seudo random numbers
	rand.Seed(time.Now().UnixNano()) //uses current time in nanoseconds, good for always being random!

	// print out a message
	color.Cyan("the Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{ //create one producer with pointer to Producer struct
		data: make(chan PizzaOrder), //in order to create a new channel you call the make method.
		quit: make(chan chan error), //we call make for a channel and assign the type, in this case make chan of chan errors
	}

	// run the producer in the background
	go pizzaeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel! ***", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("-----------------------------------")
	color.Cyan("The Pizzaria is closed for the day!")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a great day!!")
	}
}
