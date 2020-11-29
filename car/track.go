package car

type track struct {
	cars		[]*car
	height		int
	width		int
	places 		[]int
	lap			int
	laps		int
}

func newTrack(c []*cars, h, w, laps int) *track {

	p := make([]int, 8)
	for i := range p {
		p[i] = c[i].getPlace
	}

	t := &track {
		cars: c,
		height: h,
		width: w,
		places: p,
		lap: 0,
		laps: laps,
	}
}