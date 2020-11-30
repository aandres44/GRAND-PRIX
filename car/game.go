package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	
	"image"
	"os"
	_"image/png"

	// "fmt"
	"time"
	"math/rand"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

//Function to run the window
func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Grand Prix!",
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Loads images
	score, scoreErr := loadPicture("score.png")
	if scoreErr != nil {
		panic(scoreErr)
	}

	circuit, err := loadPicture("pista.png")
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Black)


	// edit position and scale of images
	scoreSprite := pixel.NewSprite(score, score.Bounds())
	scoreSprite.Draw(win, pixel.IM.Moved(pixel.V(800, 800)))

	circuitSprite := pixel.NewSprite(circuit, circuit.Bounds())
	mat := pixel.IM
	mat = mat.Moved(pixel.V(500,500))
	mat = mat.ScaledXY(pixel.V(300,540),pixel.V(2.5, 5))
	circuitSprite.Draw(win, mat)

	for !win.Closed() {
		win.Update()
	}
}


func main() {
	pixelgl.Run(run)

	// rand.Seed(time.Now().UnixNano())

	// // 50 works like the number of step in each lap
	// // here 
	// laps:= 4
	// jobs := make(chan int,laps*50)
	// results := make(chan int,laps*50)

	// go worker(jobs,results)
	// go worker(jobs,results)

	// for i := 0; i <  laps*50; i++{
	// 	jobs <- i
	// }

	// close(jobs)

	// for i := 0; i <  laps*50/2; i++{
	// 	fmt.Println( "Car1 speed: ", <- results)
	// 	fmt.Println( "Car2 speed: ", <- results)
	// 	fmt.Println( " ")

	// 	fmt.Println( <- results)
	// 	fmt.Println( <- results)
	// 	fmt.Println( " ")

	// }

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