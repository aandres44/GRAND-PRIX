package car

// car directions,letter stands for cardinal point
const (
	N direction = 1 + iota
	S
	W
	E
	NW
	NE
	SW
	SE
)

type direction int

type car struct {
	position []coord
	direction direction
	place int
	number int
	color color
	speed int
}

func newCar(pos coord, d direction, p place, n number, c color) *car {
	return &car {
		position: pos
		direction: d
		place: p
		number: n
		color: c
		speed: 0
	}
}

func changeDir(d direction) {
	opposites := map[direction]direction{
		N: S,
		S: N,
		W: E,
		E
		NW
		NE
		SW
		SE
	}

	if o := opposites[d]; o != 0 && o != s.direction {
		s.direction = d
	}
}

func (c *car) getPos() coord {
	return c.position
}

func (c *car) move() error() {

	h := getPos()
	coord := coord{x: h.x, y: h.y}

	switch c.direction {
	case N:
		coord.y++
	case S:
		coord.y--
	case W:
		coord.x--
	case E:
		coord.x++
	case NW:
		coord.y++
		coord.x--
	case NE:
		coord.y++
		coord.x++
	case SW:
		coord.y--
		coord.x--
	case SE:
		coord.y--
		coord.x++
		
	}

}