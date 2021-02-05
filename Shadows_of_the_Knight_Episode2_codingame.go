package main

import (
	"fmt"
)

var (
	doone  bool = false
	dotwo  bool = false
	icmove bool
)

const (
	UP   = 1
	DOWN = 2
)

type line struct {
	l, h int
	pos  int
}

func (lin *line) within(p int) bool {
	if p >= lin.l && p <= lin.h {
		return true
	}
	return false
}

func (lin *line) transformline(v int) int {
	if lin.pos == UP {
		return lin.l + v
	}
	if lin.pos == DOWN {
		return lin.h - v
	}
	return -1
}

func divline(lin line) (line, line) {
	d := 0
	p := 0
	if islinex(lin.l, lin.h) {
		return lin, lin
	}
	if centerline(lin.l, lin.h) {
		d++
	} else {
		p++
	}
	l1 := line{lin.l, ((lin.l + lin.h) / 2) - d, UP}

	l2 := line{((lin.l + lin.h) / 2) + d + p, lin.h, DOWN}

	return l1, l2
}

func centerline(x, X int) bool {
	return (X+x)%2 == 0
}

func isliney(x, W int) bool {
	return (W - x) == 0
}

func islinex(y, H int) bool {
	return (H - y) == 0
}

type statem struct {
	lock bool
	lm   bool
	exc  bool
}

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)

	lx := line{0, W - 1, UP}
	vx := X0
	ly := line{0, H - 1, UP}
	vy := Y0

	var fx, fy statem

	var vpx, vpy int
	fx.lock = true
	fy.lock = false
	if islinex(lx.h, lx.l) {
		fx.lm = true
		fy.lock = true
		fx.lock = false
	}

	if isliney(ly.h, ly.l) {
		fy.lm = true
		fx.lock = true
		fy.lock = false

	}

	lineonex, linetwox := divline(lx)
	lineoney, linetwoy := divline(ly)

	//var runx int = 0
	//var ignore bool
	vpy = vy
	vpx = vx

	var bombDir int
	var bombDi string
	for i := 0; i < N; i++ {
		//var bombDir string
		if fx.lock || fy.lock {
			fmt.Scan(&bombDi)
			switch bombDi {
			case "UNkNOWN":
				bombDir = 0
			case "WARMER":
				bombDir = 1
			case "COLDER":
				bombDir = 2
			case "SAME":
				bombDir = 3
			}
			//fmt.Println("nolock")
			//	bombDir = 0
		}
		updatest(&fx, &fy)
		if !fy.lock && !fx.lm {
			vpx = vx
			linxmode(&lineonex, &linetwox, bombDir, &vx, &fx)
			if fx.lock {
				fmt.Printf("%v %v\n", vx, vpy)
			}
			if fx.lm {
				vpx = vx

			}

			if fx.lock == false {
				bombDir = 1
			}
		}
		updatest(&fx, &fy)

		if !fx.lock && !fy.lm {
			vpy = vy
			linymode(&lineoney, &linetwoy, bombDir, &vy, &fy)
			if fy.lock {
				fmt.Printf("%v %v\n", vpx, vy)
			}
			if fy.lm {
				vpy = vy
			}
			if fy.lock == false {
				fmt.Printf("%v %v\n", vx, vy)
				fx.lock = true
				vpy = vy
			}
		}
	}

}

func updatest(fx, fy *statem) {
	if fx.lm == true {
		fy.exc = true
		fy.lock = true
	}
	if fy.lm == true {
		fx.exc = true
		fx.lock = true
	}

}
func linxmode(l1, l2 *line, bombDir int, vx *int, st *statem) {
	if !doone {
		bombDir = 0
		doone = true
	}

	if st.exc {
		*vx, st.lock, st.lm = Linproc(l1, l2, *vx, bombDir, false)
		st.lock = true

	} else {
		*vx, st.lock, st.lm = Linproc(l1, l2, *vx, bombDir, false)

	}

}

func linymode(l1, l2 *line, bombDir int, vy *int, st *statem) {
	if !dotwo {
		bombDir = 0
		dotwo = true
	}

	if st.exc {
		*vy, st.lock, st.lm = Linproc(l1, l2, *vy, bombDir, false)
		st.lock = true

	} else {
		*vy, st.lock, st.lm = Linproc(l1, l2, *vy, bombDir, false)

	}

}

func whichline(l1, l2 *line, v int) int {
	switch {
	case l1.within(v):

		return l1.pos
	case l2.within(v):

		return l2.pos

	default:
		return 0
	}
}
func divl(l1, l2 *line, status int) {
	if status == UP {
		*l1, *l2 = divline(*l1)
	} else {
		*l1, *l2 = divline(*l2)
	}
}
func transformlinenew(l1, l2 *line, v, status int) (int, int) {

	switch status {
	case UP:
		p := v - l1.l
		v = l2.transformline(p)
		status = l2.pos
	case DOWN:
		p := l2.h - v
		v = l1.transformline(p)
		status = l1.pos
	default:
		v = l1.l
		status = l1.pos

	}
	return v, status
}

func Linproc(l1, l2 *line, v, bombd int, exc bool) (int, bool, bool) {
	var lock, lm bool
	var status int

	switch bombd {
	case 0:

		status = whichline(l1, l2, v)
		//case "UNKNOWN":
		if status == 0 {
			icmove = true
		}
		v, status = transformlinenew(l1, l2, v, status)
		lock = true
		lm = false

	case 1:
		icmove = false

		status = whichline(l1, l2, v)

		divl(l1, l2, status)

		if islinex(l1.l, l2.l) {
			lock = false
			lm = true
			return v, lock, lm

		}

		status = whichline(l1, l2, v)
		if status == 0 {
			icmove = true
		}
		v, status = transformlinenew(l1, l2, v, status)

		lock = true
		lm = false

	case 2:

		status = whichline(l1, l2, v)

		v, status = transformlinenew(l1, l2, v, status)

		lock = false
		lm = false
		if icmove {
			lock = true

			icmove = false
		}
	case 3:

		//	case "SAME":

		v = l1.h + 1

		lock = false
		lm = true
	}
	return v, lock, lm

}
