package octa

import (
	"fmt"
	"math"

	"github.com/heptagons/meccano"
)

// Find bars forming octagon internal angle 135°.
//  
//  D . _ 
//  | \   ` - _ C
//  |   \       /
//  |     \   /
//  A-------B

//  Angles: ABD=45°, DBC=90°, ABC=135°

//  i = A-D = A-D, B-D = i*sqrt(2)
//  we iterate bc to find DC integer
//  
func Method_1() {
	for ab := 1; ab <= 30; ab++ {
		h_2 := 2*ab*ab
		for bc := 1; bc <= 4*ab; bc++ {
			h_3 :=  float64(bc*bc + h_2)
			cd := math.Sqrt(h_3)
			_cd := int(cd)
			if cd == float64(_cd) {
				if meccano.Gcd(_cd, meccano.Gcd(ab, bc)) == 1 {
					fmt.Printf("✔ ab=%2d bc=%2d cd=%2d\n", ab, bc, _cd)
				}
			}
		}
	}
	// Expected
	// ✔ ab= 2 bc= 1 cd= 3
	// ✔ ab= 4 bc= 7 cd= 9
	// ✔ ab= 6 bc= 7 cd=11
	// ✔ ab= 6 bc=17 cd=19
	// ✔ ab=10 bc=23 cd=27
	// ✔ ab=12 bc= 1 cd=17
	// ✔ ab=14 bc=47 cd=51
	// ✔ ab=20 bc=17 cd=33
	// ✔ ab=24 bc=23 cd=41
	// ✔ ab=28 bc=41 cd=57
	// ✔ ab=30 bc= 7 cd=43
	// ✔ ab=30 bc=41 cd=59
}

// Octagons bars second method
// 1) i iterate cd=1,2,3,...
// 2) j iterate bc=1,2,3 < cd
// 3) calculate bd*bd = cd**cd - bc*bc
// 4) accept when bd*db = 2 * square
func Angles135(max int) {
	for x := 1; x < max; x++ {
		for y := 1; y < x; y++ {
			if zz := x*x - y*y; zz % 2 == 0 {
				f := math.Sqrt(float64(zz / 2))
				if z := int(f); f == float64(z) {
					if meccano.Gcd(z, meccano.Gcd(x, y)) == 1 {
						a := int(math.Max(float64(y), f))
						fmt.Printf("a=%3d x=%3d y=%3d z=%3d\n", a, x, y, z)
					}
				}
			}
		}
	}
}

/*
func Diagonals(max int) {

	sols := &meccano.Sols{}

	for a := 1; a < max; a++ {
		for b := 1; b < a; b++ {
			cc := a*a - b*b
			if cc % 2 == 0 {
				f := math.Sqrt(float64(cc/2))
				if c := int(math.Sqrt(f)); math.Pow(float64(c), 2) == f {
					sols.Add(a, b, c)
				}
			}
		}
	}
}

}
	for cd := ; a < max; a++ {
		for b := 1; b <= a; b++ {
			f := float64(2*a*a + b*b)
			if c := int(math.Sqrt(f)); math.Pow(float64(c), 2) == f {
				sols.Add(a, b, c)
			}
		}
	}
}*/