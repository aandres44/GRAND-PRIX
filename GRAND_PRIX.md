# Multithreaded Grand Prix

This is an implementation of a a multithreaded grand prix simulator where every car will be an independent entity. 
It will be simulating an N laps race.

## Technical Requirements

- The race's lap is introduced by the user as well as the number of cars.
- Random max speed for each car.
- Emulate the increasing speed behaviour that any car has.
- If a car detect another car on his route it will slow down its speed simulating a crash.
- Each racer behaviour is a separated thread.
- Cars threads use the same map or city layout data structure resource.
- Displays each car's speed, position, racing time and lap.
- At the end, top 3 winners is displayed, if there is a tie for 3rd more cars will be shown.

## Build

### On __Linux__ or __MacOs__

```bash
# To compile the go proyect
make build

# To remove the executable created after compilation.
make clean
```

### On __Windows__

```bash
# To get dependencies
go get -v

# To build the proyect
go build
```

## Run

### On __Linux__

```bash
make run
# or
make
```
### On __Windows__

```bash
go run game.go
```
