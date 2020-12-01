# GRAND-PRIX

## Team

Andres Aguirre Alvarez & Luis Garcia Miramontes

A01228159 & A01540063

## Professor 

Obed Nehemías Muñóz

Technical Requirements
----------------------
- The race's lap can be static or automatically generated.
- Racers number can be configured on start.
- Define a random max speed for each car.
- You will emulate the increasing speed behaviour that any car has.
- If a car detect another car on his route and it's slower, it must slow down its speed or try to rebase it.
- Each racer behaviour will be implemented as a separated thread.
- Cars threads must use the same map or city layout data structure resource.
- Define how many laps before starting the race.
- Display each car's speed, position, racing time and lap.
- At the end, display the top 3 winners.


## Description

Final project for Advanced Programming Course

In this project we implemented a video game that uses goroutines to be able to place a car in the screen. Also we implemented a crashing
where a car can return to the race if is crashing with another car.

28/11/2020

link to video
image of diagram

### Info

### Packages
We used the pexternal ackages image and faiface/pixel.

The first 3 were used to edit the images and text on them. Then the last 2 were used to import font types and colors
 
	"github.com/faiface/pixel/text"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
 
  "golang.org/x/image/font/basicfont"
	"golang.org/x/image/colornames"
  
### Structures
Car
	id int - An identifier for the cars
	currentLap int - Current lap of the cars
	speed chan int - Channel for the goroutine to get the speed
	boundPos chan int - Channel for the gourutine to get the position in the Y bounds of the car
	sprite *pixel.Sprite - The sprite(image) of the car
	mat pixel.Matrix  - Matrix of the sprite
	seed rand.Source - 
	botBound int
	topBound int
	crashing bool
	finished bool
	position int
	finalPosition int
	timeElapsed time.Duration
### Functions
