package main

import (
	// "github.com/faiface/pixel"
	// "github.com/faiface/pixel/pixelgl"
	// "golang.org/x/image/colornames"

	"fmt"
	"time"
	"math/rand"
)

// Function to run the window

// func run() {
// 	cfg := pixelgl.WindowConfig{
// 		Title:  "Grand Prix!",
// 		Bounds: pixel.R(0, 0, 1024, 768),
// 		VSync:  true,
// 	}
// 	win, err := pixelgl.NewWindow(cfg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	win.Clear(colornames.Blue)

// 	for !win.Closed() {
// 		win.Update()
// 	}
// }

func main() {
	//pixelgl.Run(run)
	rand.Seed(time.Now().UnixNano())

	// 50 works like the number of step in each lap
	// here 
	laps:= 4
	jobs := make(chan int,laps*50)
	results := make(chan int,laps*50)

	go worker(jobs,results)
	go worker(jobs,results)

	for i := 0; i <  laps*50; i++{
		jobs <- i
	}

	close(jobs)

	for i := 0; i <  laps*50/2; i++{
		// fmt.Println( "Car1 speed: ", <- results)
		// fmt.Println( "Car2 speed: ", <- results)
		// fmt.Println( " ")

		fmt.Println( <- results)
		fmt.Println( <- results)
		fmt.Println( " ")

	}

}


// 
func worker(jobs <- chan int, results chan<- int){
	for n:= range jobs{
		results <- randomSpeed(n)
		time.Sleep(time.Millisecond *500)
	}
}

func randomSpeed(n int) int{
	return 260 + rand.Intn(10)
}