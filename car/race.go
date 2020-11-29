package car

type Race struct {
	track  *track
	places  []int
	isOver  bool
	lap   	int
	laps	int
	//fN 		int
}

func (r *Race) end() {
	r.isOver = true
}

func (r *Race) start() {

}
