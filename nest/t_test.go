package nest

import (
	"fmt"
	"strings"
	"testing"
)

func TestTCosSin(t *testing.T) {

	factory := NewT32s()

	frac := func(t *TRat) string {
		a, _ := factory.aNew1(N(t.den), t.num)
		return a.String()
	}
	for p, r := range []struct { a, b, c N32; cosines, sines string } {
		{ 1, 1, 1, "1/2 1/2 1/2",       "√3/2 √3/2 √3/2"          }, // √3
		{ 2, 2, 1, "1/4 1/4 7/8",       "√15/4 √15/4 √15/8"       }, // √15
		{ 3, 2, 2, "-1/8 3/4 3/4",      "3√7/8 √7/4 √7/4"         }, // √7
		{ 3, 3, 1, "1/6 1/6 17/18",     "√35/6 √35/6 √35/18"      }, // √35
		{ 3, 3, 2, "1/3 1/3 7/9",       "2√2/3 2√2/3 4√2/9"       }, // √2
		{ 4, 3, 2, "-1/4 11/16 7/8",    "√15/4 3√15/16 √15/8"     }, // √15
		{ 4, 3, 3, "1/9 2/3 2/3",       "4√5/9 √5/3 √5/3"         }, // √5
		{ 4, 4, 1, "1/8 1/8 31/32",     "3√7/8 3√7/8 3√7/32"      }, // √7
		{ 4, 4, 3, "3/8 3/8 23/32",     "√55/8 √55/8 3√55/32"     }, // √55
		{ 5, 3, 3, "-7/18 5/6 5/6",     "5√11/18 √11/6 √11/6"     }, // √11
		{ 5, 4, 2, "-5/16 13/20 37/40", "√231/16 √231/20 √231/40" }, // √231
		{ 5, 4, 3, "0 3/5 4/5",         "1 4/5 3/5"               }, // √1
		{ 5, 4, 4, "7/32 5/8 5/8",      "5√39/32 √39/8 √39/8"     }, // √39
		{ 5, 5, 1, "1/10 1/10 49/50",   "3√11/10 3√11/10 3√11/50" }, // √11
		{ 5, 5, 2, "1/5 1/5 23/25",     "2√6/5 2√6/5 4√6/25"      }, // √6
		{ 5, 5, 3, "3/10 3/10 41/50",   "√91/10 √91/10 3√91/50"   }, // √91
		{ 5, 5, 4, "2/5 2/5 17/25",     "√21/5 √21/5 4√21/25"     }, // √21

		{ 7, 6, 5, "1/5 19/35 5/7",     "2√6/5 12√6/35 2√6/7"     }, // √6

	} {
		tr := newT(r.a, r.b, r.c)
		if tr == nil {
			t.Fatalf("T: nil for %d %d %d", r.a, r.b, r.c)
		}
		cosA, cosB, cosC := factory.tRatCosinesAll(tr) // always three angles A,B,C including repetitions
		sinA, sinB, sinC := factory.tRatSinesAll(tr)

		if cosines := fmt.Sprintf("%s %s %s", frac(cosA), frac(cosB), frac(cosC)); cosines != r.cosines {
			t.Fatalf("T.cosines got %s exp %s", cosines, r.cosines)
		} else if sines := fmt.Sprintf("%s %s %s", sinA.String(), sinB.String(), sinC.String()); sines != r.sines {
			t.Fatalf("T.sines got %s exp %s", sines, r.sines)
		} else {
			fmt.Printf("%d & (%s) & $%s$ & $%s$ & $%s$ & $%s$ & $%s$ & $%s$\\\\\n", (p+1), tr.String(),
				cosA.Tex(), cosB.Tex(), cosC.Tex(),
				sinA.Tex(), sinB.Tex(), sinC.Tex())
		}
	}
}

func TestTCosXY_(t *testing.T) {
	factory := NewT32s()
	ts := &Ts{}
	ts.AddTris(3)
	sums := make(map[string]bool, 0)
	listN := 0
	for i, m := range ts.tris {
		xCosines := factory.tRatCosines(m)
		for j:=0; j <= i; j++ {
			n := ts.tris[j]
			yCosines := factory.tRatCosines(n)
			
			for _, xRat := range xCosines.rats {
				for _, yRat := range yCosines.rats {
					ts := fmt.Sprintf("(%s)[%c] & (%s)[%c]", m, xRat.angle, n, yRat.angle)
					if xy, err := factory.tRatCosXY(xRat, yRat); err == nil {
						//if len(xy.num) > 2 && xy.num[2] == 5 {
							key := xy.Key()
							if _, ok := sums[key]; !ok {
								sums[key] = true
								listN++
								fmt.Printf("%d & %s & $%s$\\\\ %%%s\n", listN, ts, xy.Tex(), xy.String())
							}
						//}
					}
				}
			}
		}
	}
}

func TestTCos2XY_(t *testing.T) {
	factory := NewT32s()
	ts := &Ts{}
	ts.AddTris(3)
	listN := 0
	for i:=0; i < len(ts.tris); i++ {
		m := ts.tris[i]
		xCosines := factory.tRatCosines(m)
		for j:=0; j <= i; j++ {
			n := ts.tris[j]
			yCosines := factory.tRatCosines(n)
			for _, xRat := range xCosines.rats {
				for _, yRat := range yCosines.rats {
					ts := fmt.Sprintf("(%s)[%c] & (%s)[%c]", m, xRat.angle, n, yRat.angle)

					if cos1, err := factory.tRatCos2XY(xRat, yRat); err != nil {

					} else if cos2, err := factory.tRatCos2XY(yRat, xRat); err != nil {

					} else {
						listN++
						fmt.Printf("%d & %s & $%s$ & $%s$\\\\\n", listN, ts, cos1.Tex(), cos2.Tex())
					}
				}
			}
		}
	}
}

// prints Latex string rows with columns: 
// triangle-angle-1 | triangle-angle-2 | triangle-angle-3 | cos(X,Y,Z)
func TestTCosXYZ(t *testing.T) {
	factory := NewT32s()
	
	ts := &Ts{}
	ts.AddTris(3)
	listN := 0
	for i:=0; i < len(ts.tris); i++ {
		x := ts.tris[i]
		xCosines := factory.tRatCosines(x)
		for j:=0; j <= i; j++ {
			y := ts.tris[j]
			yCosines := factory.tRatCosines(y)
			for k := 0; k <= j; k++ {
				z := ts.tris[k]
				zCosines := factory.tRatCosines(z)

				for _, xRat := range xCosines.rats {
					for _, yRat := range yCosines.rats {
						if i == j && xRat.angle == yRat.angle { continue }
						for _, zRat := range zCosines.rats {
							if i == k && xRat.angle == zRat.angle { continue }
							if j == k && yRat.angle == zRat.angle { continue }
							ts := fmt.Sprintf("(%s)[%c] & (%s)[%c] & (%s)[%c]", x, xRat.angle, y, yRat.angle, z, zRat.angle)
							an, ad, as := xRat.num, xRat.den, xRat.S()
							dn, dd, ds := yRat.num, yRat.den, yRat.S()
							gn, gd, gs := zRat.num, zRat.den, zRat.S()
							s := NewS32s(factory.Z32s)
							s.sAddSqrt(an*dn*gn,     1)
							s.sAddSqrt(     -gn, as*ds)
							s.sAddSqrt(     -dn, as*gs)
							s.sAddSqrt(     -an, ds*gs)
							s.sDivide (N(ad)*N(dd)*N(gd))
							listN++
							// to be used in /github.com/heptagons/meccano/nest/doc/triangles.tex table
							fmt.Printf("%d & %s & $%s$\\\\ %%%s\n", listN, ts, s.Tex(), s.String())
						}
					}
				}
			}
		}
	}
}















func TestTdiag1(t *testing.T) {

	factory := NewT32s()

	for _, tr := range []*T { 
		newT(3,3,3), // equilateral
		newT(4,3,3), // isoceles
		newT(7,6,5), // scalene
	} {
		fmt.Printf("[%s] diagonals:\n", tr.String())
		for _, ang := range []Tang { TangC, TangB, TangA } {
			diagsAng, den := factory.tDiagsAng(tr, ang)
			fmt.Printf("  ang=%c den=%d:\n", ang, den)
			for d, diags := range diagsAng {
				var surds strings.Builder
				for pos, diag := range diags {
					if pos > 0 {
						surds.WriteString(", ")
					}
					if a, err := factory.aNew3(den, 0, 1, Z(diag)); err != nil {
						continue
					} else {
						surds.WriteString(a.String())
					}
				}
				fmt.Printf("    %d: %v -> %s\n", d, diags, surds.String())
			}
		}
	}
}


func TestTdiag2(t *testing.T) {
	factory := NewT32s()
	/*   *
		/|\
     5 / | \ 6
      /  |  \
     /   |   \
    x1   |8   x2  x1x2 = 3 + 2√5
     \   |   /
      \  |  /
     5 \ | / 6
        \|/
         *
	*/
    for _, pair := range []struct { t1, t2 *T; max, min N32 } {
    	{ newT(8,5,5), newT(8,6,6), 6, 5 },
    	{ newT(7,6,5), newT(6,5,4), 6, 5 },
    	{ newT(5,5,5), newT(5,5,5), 5, 3 },
    	{ newT(5,5,5), newT(5,5,4), 5, 4 },
    	{ newT(5,5,5), newT(5,4,4), 5, 4 },
    } {
    	fmt.Printf("Pair [%s] [%s]\n", pair.t1.String(), pair.t2.String())
		cosX, _ := factory.tCosAplusB(pair.t1, TangB, pair.t2, TangB)
		fmt.Printf("  cosBB = %s\n", cosX.String())
		for p1 := N32(1) ; p1 <= pair.max; p1++ {
			for p2 := N32(1) ; p2 <= pair.min; p2++ {
				if p1 >= p2 {
					x, _ := factory.tLawOfCos(p1, p2, cosX)
					fmt.Printf("    [%d,%d,%s]\n", p1, p2, x)
				}
			}
		}
    }
}

func TestT765diags(t *testing.T) {
	factory := NewT32s()
	a, b, c := N(7), N(6), N(5)
	aa, bb, cc := Z(a)*Z(a), Z(b)*Z(b), Z(c)*Z(c)
	
	m := make(map[string]*A32)

	bc := N(b)*N(c)
	bc2 := Z(bc)*Z(bc)
	b2_c2_a2 := bb + cc - aa
	for _, g := range []struct { y, z Z } {
		{ 1,1 },{ 2,2 },{ 3,3 },{ 4,4 },{ 5,5 },
		{ 2,1 },{ 3,2 },{ 4,3 },{ 5,4 },
		{ 3,1 },{ 4,2 },{ 5,3 },{ 6,4 },
		{ 4,1 },{ 5,2 },{ 6,3 },
		{ 5,1 },{ 6,2 },
		{ 6,1 },
	} {
		in := bc2*(g.y*g.y + g.z*g.z) - Z(bc)*g.y*g.z*b2_c2_a2
		if diag, err := factory.aNew3(bc, 0, 1, in); err == nil {
			m[fmt.Sprintf("b_%d,c_%d", g.y, g.z)] = diag
		}
	}

	ac := N(a)*N(c)
	ac2 := Z(ac)*Z(ac)
	a2_c2_b2 := aa + cc - bb
	for _, g := range []struct { x, z Z } {
		{ 1,1 },{ 2,2 },{ 3,3 },{ 4,4 },{ 5,5 },
		{ 2,1 },{ 3,2 },{ 4,3 },{ 5,4 },{ 6,5 },
		{ 3,1 },{ 4,2 },{ 5,3 },{ 6,4 },
		{ 4,1 },{ 5,2 },{ 6,3 },{ 7,4 },
		{ 5,1 },{ 6,2 },{ 7,3 },
		{ 6,1 },{ 7,2 },
		{ 7,1 },
	} {
		in := ac2*(g.x*g.x + g.z*g.z) - Z(ac)*g.x*g.z*a2_c2_b2
		if diag, err := factory.aNew3(ac, 0, 1, in); err == nil {
			m[fmt.Sprintf("a_%d,c_%d", g.x, g.z)] = diag
		}
	}

	ab := N(a)*N(b)
	ab2 := Z(ab)*Z(ab)
	a2_b2_c2 := aa + bb - cc
	for _, g := range []struct { x, y Z } {
		{ 1,1 },{ 2,2 },{ 3,3 },{ 4,4 },{ 5,5 },{ 6,6 },
		{ 2,1 },{ 3,2 },{ 4,3 },{ 5,4 },{ 6,5 },
		{ 3,1 },{ 4,2 },{ 5,3 },{ 6,4 },{ 7,5 },
		{ 4,1 },{ 5,2 },{ 6,3 },{ 7,4 },
		{ 5,1 },{ 6,2 },{ 7,3 },
		{ 6,1 },{ 7,2 },
		{ 7,1 },
	} {
		in := ab2*(g.x*g.x + g.y*g.y) - Z(ab)*g.x*g.y*a2_b2_c2
		if diag, err := factory.aNew3(ab, 0, 1, in); err == nil {
			m[fmt.Sprintf("a_%d,b_%d", g.x, g.y)] = diag	
		}
	}
	print := func(X, Y N, f string) {
		for x:=N(1); x <= X; x++ {
			first := false
			for y:=N(1); y <= Y; y++ {
				if !first { first=true } else { fmt.Print("& " ); }
				if diag, ok := m[fmt.Sprintf(f, y, x)]; ok {
					fmt.Printf("%s ", diag.Tex())
				}
			}
			fmt.Println("\\\\")
		}
	}
	fmt.Println("A[b,c]"); print(c, b, "b_%d,c_%d")
	fmt.Println("B[a,c]"); print(c, a, "a_%d,c_%d")
	fmt.Println("C[a,b]"); print(b, a, "a_%d,b_%d")
}

func TestTalphas(t *testing.T) {
	factory := NewT32s()
	sqrt := N32(24)
	slur, _ := factory.aNew3(1, 0, 1, Z(sqrt))
	for _, tri := range newTalphas(sqrt) {
		if cosA, cosB, cosC, err := factory.tAlphaCosines(tri); err != nil {
			t.Fatalf("error %v", err)
		} else {
			t.Logf("(%s,%d,%d) %s %s %s",
				slur.String(), tri.b, tri.c, cosA, cosB, cosC)
		}
	}
}

func TestTbetas(t *testing.T) {
	factory := NewT32s()
	sqrt := N32(24)
	slur, _ := factory.aNew3(1, 0, 1, Z(sqrt))
	for _, tri := range newTbetas(sqrt,9) {
		if cosA, cosB, cosC, err := factory.tBetaCosines(tri); err != nil {
			t.Fatalf("error %v", err)
		} else {
			t.Logf("(%d,%s,%d) %s %s %s",
				tri.a, slur.String(), tri.c, cosA, cosB, cosC)
		}
	}
}

func TestTgammas(t *testing.T) {
	factory := NewT32s()
	sqrt := N32(24)
	slur, _ := factory.aNew3(1, 0, 1, Z(sqrt))
	for _, tri := range newTgammas(sqrt,9) {
		if cosA, cosB, cosC, err := factory.tGammaCosines(tri); err != nil {
			t.Fatalf("error %v", err)
		} else {
			t.Logf("(%d,%d,%s) %s %s %s",
				tri.a, tri.b, slur.String(), cosA, cosB, cosC)
		}
	}
}

