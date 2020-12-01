package main

import (

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	
	"image"
	"os"
	_"image/png"

	"fmt"
	"time"
	"math/rand"
)


type Car struct {
	id int
	lap int
	position int
	//speed chan int
	sprite *pixel.Sprite
	mat pixel.Matrix
}

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

	// Text in the screen
	// just say "score"
	scoreAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	scoreTxt := text.New(pixel.V(500, 750), scoreAtlas)
	fmt.Fprintln(scoreTxt, "Score: ")


	infoAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	infoTxt := text.New(pixel.V(650, 820), infoAtlas)


	// winnerAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	// winnerTxt := text.New(pixel.V(50, 500), winnerAtlas)

	// Loads images
	carPic, err := loadPicture("car.png")
	if err != nil {
		panic(err)
	}

	circuit, err := loadPicture("pista.png")
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Black)


	// edit position and scale of images
	// circuit
	circuitSprite := pixel.NewSprite(circuit, circuit.Bounds())
	mat := pixel.IM
	mat = mat.Moved(pixel.V(500,500))
	mat = mat.ScaledXY(pixel.V(300,540),pixel.V(2.5, 5))
	circuitSprite.Draw(win, mat)

	initSpaceCar:= 570.0
	spaceCar:= 500.0 / totalCars
	
	// car init
	for i:= 0; i< totalCars; i++{
		cars = append(cars, Car{i,0,0,pixel.NewSprite(carPic, carPic.Bounds()), pixel.IM.Moved(pixel.V(70, initSpaceCar))})
		initSpaceCar -= float64(spaceCar)
	}

// -------------------------------------------------------------//

	for !win.Closed() {
		scoreTxt.Draw(win, pixel.IM.Scaled(scoreTxt.Orig, 4))

		infoTxt.Clear()

		// winners
		if winners == 3{

		} else{
			//infoTxt.Clear()
			

			for i:= 0; i< totalCars; i++{
				// update positions

				//print car info
				//fmt.Fprintf(infoTxt, "Car: %d  Lap: %d  Position: %d  Speed: %d mp/h \n", cars[i].id+1,cars[i].lap,cars[i].position,speed)
				fmt.Fprintf(infoTxt, "Car: %d   Lap: %d   Position: %d  \n", cars[i].id+1,cars[i].lap,cars[i].position)

			}



			for i:= 0; i< totalCars; i++{
				cars[i].sprite.Draw(win, cars[i].mat)
			}

		}

		infoTxt.Draw(win, pixel.IM.Scaled(scoreTxt.Orig, 1.5))

		// last thing to do
		win.Update()
	}
}

// global vars
var winners int
var totalCars int
var cars []Car

// var tempTotalSpeed []int
// var tempXPos []int
// var laps int
// var positionArray []int
// var top []Car

func main() {

	winners = 0
	tempTotal := 10
	totalCars = tempTotal
	


	pixelgl.Run(run)


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