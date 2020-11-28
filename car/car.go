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

func (c *car) changeDir(d direction) { // <------- hay que revisar si es necesario ver los opuestos
	opposites := map[direction]direction{
		N: S,
		S: N,
		W: E,
		E: W,
		NW: NE,
		NE: NW,
		SW: SE,
		SE: SW, // <--- va una coma ahi??
	}

	if o := opposites[d]; o != 0 && o != c.direction {
		c.direction = d // changes direction
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