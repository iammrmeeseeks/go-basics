package main

import (
	"fmt"
)

func fibbonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibbonacci(c, quit)
}

/*
 *
Imagine you have a toy factory that makes Fibonacci toys, and you have two workers.
One worker makes the toys, and the other worker puts them on display.

Worker 1 (Maker): Makes Fibonacci toys and puts them on a conveyor belt.
Worker 2 (Displayer): Takes toys from the conveyor belt and puts them on display.


*** Step 1: Setting Up the Factory
Factory Setup:
c is the conveyor belt where toys (numbers) are placed.
quit is a signal that tells Worker 1 to stop making toys.

Factory Setup:
     +-----------------+
     | Conveyor Belt (c)|
     +-----------------+
     |    Signal (quit) |
     +-----------------+


*** Step 2: Worker 2 (Displayer) Starts

Worker 2 Starts:
Worker 2 starts a loop to take 10 toys from the conveyor belt and put them on display.
After displaying 10 toys, Worker 2 sends a stop signal on quit.

Worker 2 (Displayer):
     +---------------------+
     | Display 10 Toys     |
     | from Conveyor (c)   |
     +---------------------+
     |    Send Stop Signal |
     |    on (quit)        |
     +---------------------+


*** Step 3: Worker 1 (Maker) Starts
Worker 1 Starts:
Worker 1 starts making Fibonacci toys and puts them on the conveyor belt.
Worker 1 stops making toys when they receive the stop signal from quit.

Worker 1 (Maker):
     +---------------------+
     | Make Fibonacci Toys |
     | Put on Conveyor (c) |
     +---------------------+
     |    Stop on Signal   |
     |    from (quit)      |
     +---------------------+


Initialization:

	c := make(chan int)				// Create the conveyor belt (c).
	quit := make(chan int)			// Create the stop signal (quit).

Worker 2 (Displayer):

	go func() {						// Worker 2 starts and runs in parallel.
    for i := 0; i < 10; i++ {
        fmt.Println(<-c)			// Displays 10 toys (numbers) taken from the conveyor belt (c).
    	}
    	quit <- 0					// Sends stop signal (quit <- 0).
	}()


Worker 1 (Maker):

	fibonacci(c, quit)		   		// Worker 1 starts making Fibonacci toys.
									// Keeps making and putting toys on the conveyor belt (c <- x).
									// Stops when it receives the stop signal (<-quit).


1) Conveyor Belt and Signals:

 +--------+   +-----+   +-----+   +-----+
 |  c (1) |-> |  c  |-> |  c  |-> |  c  |
 +--------+   +-----+   +-----+   +-----+
 (Fibonacci numbers: 0, 1, 1, 2, 3, 5, ...)

2) Stop Signal:

 +---------+
 | quit (0)|
 +---------+

3) Workers Interaction:

 Worker 1 (Maker)                     Worker 2 (Displayer)
 +-------------------+                +------------------+
 |                   |                |                  |
 | c <- Fibonacci    |                |  Display c       |
 |   0, 1, 1, ...    |                |  0, 1, 1, ...    |
 |                   |                |                  |
 | <-quit            |                |  quit <- 0       |
 +-------------------+                +------------------+

Worker 1 (Maker): Creates Fibonacci numbers and puts them on the conveyor belt.
Worker 2 (Displayer): Takes 10 numbers from the conveyor belt and puts them on display.
	Then tells Worker 1 to stop.

*/
