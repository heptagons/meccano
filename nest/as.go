package nest

import (
	"fmt"
)

// A32s is a factory to produce algebraic numbers A32
// uses Z32s factory which produce 32 bit integers
type A32s struct {
	*Z32s
}

// NewA32s creates A32s factory
func NewA32s() *A32s {
	return &A32s{
		Z32s: NewZ32s(),
	}
}

// 1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32
//                                                                                                                                    ,-
//                                                                   ,-------------------------------------------------------        /
//                                  ,------------------------       /                               ,------------------------       /
//                 ,---------      /               ,---------      /               ,---------      /               ,---------      /
//        ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /
// b + c √ d + e √ f + g √ h + i √ j + k √ l + m √ n + o √ p + q √ r + s √ t + u √ v + w √ x + y √ z + A √ B + C √ D + E √ F + G √ ...
// =-----------=---------------=-------------------------------=---------------------------------------------------------------=------
//                                                             a
// aNew produces a A32 number trying to simplify and reduce.
func (qs *A32s) aNew(d N, n ...Z) (q *A32, err error) {
	switch len(n) {
	case 1:
		// b/a
		return qs.aNew1(d, n[0])
	case 3:
		// (b+c√d)/a
		return qs.aNew3(d, n[0], n[1], n[2])
	case 5:
		// (b+c√d+e√f)/a
		return qs.aNew5(d, n[0], n[1], n[2], n[3], n[4])
	case 7:
		// (b+c√d+e√f+g√h)/a
		return qs.aNew7(d, n[0], n[1], n[2], n[3], n[4], n[5], n[6])
	default:
		return nil, ErrInvalid
	}
}

// aNew1 produces a most a number of size 1, that is number like b/a.
func (qs *A32s) aNew1(a N, b Z) (*A32, error) {
	if A, B, err := qs.zFrac(a, b); err != nil {
		return nil, err
	} else {
		return newA32(A, B), nil
	}
}

// aNew3 produces at most a number of size 3, that is a number like (b + c√d)/a.
// If c=0 or d=0 or d=1 produces a number of size 1.
func (qs *A32s) aNew3(a N, b, c, d Z) (*A32, error) {
	if c == 0 || d == 0 {
		return qs.aNew1(a, b)                                   // b/a
	} else if d == 1 {
		return qs.aNew1(a, b + c)                               // (b+c)/a
	} else if C, D, err := qs.zSqrt(c, d); err != nil {         // (b + C√D)/a
		return nil, err
	} else if D == 1 {
		return qs.aNew1(a, b + Z(C))                            // (b+c)/a
	} else if A, BC, err := qs.zFracN(a, b, Z(C)); err != nil { // (B + C√D)/A
		return nil, err
	} else {
		B, C := BC[0], BC[1]
		return newA32(A, B, C, D), nil
	}
}

// aNew5 produces at most a number of size 5, that is a number like (b + c√d + e√f)/a.
// If e=0 or f=0 or d=f, produces a number of size 3.
func (qs *A32s) aNew5(a N, b, c, d, e, f Z) (*A32, error) {
	if e == 0 || f == 0 {                              // (b + c√d)/a
		return qs.aNew3(a, b, c, d)
	} else if f == 1 {                                 // (b+e + c√d)/a
		return qs.aNew3(a, b+e, c, d)
	} else if d == f {                                 // (b + (c+e)√d)/a
		return qs.aNew3(a, b, c+e, d)
	} else if C, D, err := qs.zSqrt(c, d); err != nil { // (b + C√D + e√f)/a
		return nil, err
	} else if D == +1 {                                // (b+C + e√f)/a
		return qs.aNew3(a, b+Z(C), e, f) 
	} else if E, F, err := qs.zSqrt(e, f); err != nil { // (b + C√D + E√F)/a
		return nil, err
	} else if F == 1 {                                 // (b+E, C√D)/a
		return qs.aNew3(a, b+Z(E), Z(C), Z(D))
	} else if D == F {                                 // (b, (C+E)√D)/a)
		return qs.aNew3(a, b, Z(C) + Z(E), Z(D))
	} else if A, BCE, err := qs.zFracN(a, b, Z(C), Z(E)); err != nil { // (B + C√D + E√F)/A
		return nil, err
	} else {
		B, C, E := BCE[0], BCE[1], BCE[2]
		if D < F {
			return newA32(A, B, C, D, E, F), nil
		} else {
			return newA32(A, B, E, F, C, D), nil
		}
	}
}

// aNew7 produces at most a number of size 7, that is a number like (b+c√d+e√(f+g√h))/a.
// If g=0 or h=+1, produces a number of size 5. If e=0 produces a number of size 3.
func (qs *A32s) aNew7(a N, b, c, d, e, f, g, h Z) (*A32, error) {
	if g == 0 {                                         // (b+c√d+e√f)/a
		return qs.aNew5(a, b, c, d, e, f)            
	} else if h == +1 {                                 // (b+c√d+e√f+g)/a
		return qs.aNew5(a, b, c, d, e, f+g)          
	} else if e == 0 {                                  // (b+c√d)/a
		return qs.aNew3(a, b, c, d)               
	} else if G, H, err := qs.zSqrt(g, h); err != nil { // (b+c√d+e√(f+G√H))/a
		return nil, err
	} else if H == + 1 {                                // (b+c√d+e√f+G)/a
		return qs.aNew5(a, b, c, d, e, f+Z(G))

	} else if E, FG, err := qs.zSqrtN(e, f, Z(G)); err != nil { // (b+c√d+E√(F+G√H))/a
		return nil, err
	} else if C, D, err := qs.zSqrt(c, d); err != nil {         // (b+C√D+E√(F+G√H))/a
		return nil, err
	} else if A, BCE, err := qs.zFracN(a, b, Z(C), Z(E)); err != nil { // (B + C√E + D√(F+G√H))/A
		return nil, err
	} else {
		B, C, E := BCE[0], BCE[1], BCE[2]
		F, G := FG[0], FG[1]
		return newA32(A, B, C, D, E, F, G, H), nil
	}
}

// aAddN adds the given numbers and returns a new result number.
// Call multiple times this function aAdd and follow its limitations.
// Return an error if summand sizes cannot be added.
func (qs *A32s) aAddN(q ...*A32) (s *A32, err error) {
	n := len(q)
	if n == 1 {
		return q[0], nil
	}
	max := q[0]
	for i := 1; i < n; i++ {
		min := q[i]
		if max == nil || min == nil {
			return nil, nil
		}
		if s, err = qs.aAdd(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

// aMulN multiplies the given numbers and return a new result number.
// Call multiple times this function aMul and follows its limitations.
// Return an error if multiplands sizes cannot be multiplied.
func (qs *A32s) aMulN(q ...*A32) (s *A32, err error) {
	n := len(q)
	if n == 1 {
		return q[0], nil
	}
	max := q[0]
	for i := 1; i < n; i++ {
		min := q[i]
		if max == nil || min == nil {
			return nil, nil
		}
		if s, err = qs.aMul(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

// aAdd return the addition of the two given numbers.
// Limitations:
//	- Both numbers sizes=1 are added ok.
//	- One number size=3 and the other size=1 are added ok.
//	- Both numbers sizes=3 are added ok.
//	- For other sizes combination "Can't add pair" error is returned.
func (qs *A32s) aAdd(q, r *A32) (s *A32, err error) {
	if q == nil || r == nil {
		return nil, nil
	}
	max, min := q, r
	if len(q.num) < len(r.num) {
		max, min = r, q
	}
	A, B := max.ab()
	a, b := min.ab()
	// Use lcm always to prevent overflows
	w := (A / Ngcd(A, a)) * a // lcm denominator
	U := Z(w / A) // upper factor for max numerator terms
	u := Z(w / a) // upper factor for min numerator terms
	switch len(max.num) {
	case 1:
		//  B     b     BU + bu    x
		// --- + --- = -------- = ---
		//  A     a        w       w
		return qs.aNew(w,
			B*U + b*u) // x

	case 3:
		C, D := max.cd()
		switch len(min.num) {
		case 1:
			// B + C√D    b    BU + bu + CU√D    x + y√z
			// ------- + --- = -------------- = --------
			//    A       a          w             w
			return qs.aNew(w,
				B*U + b*u, // x
				C*U, D)    // y√z

		case 3:
			rc, rd := Z(min.num[1]), Z(min.num[2])
			if D == rd { // simpler case
				// B + C√D   b + rc√D  B*U + b*u + (C*U + rc*u)√D
				// ---------- + --------- = ----------------------------------
				//     A          a                     w
				return qs.aNew(w, B*U + b*u, C*U + rc*u, D)
			}
			if B == b && B == 0 { // simpler case both b's=0
				// C√D   rc√rd   C*U√D + rc*u√rd
				// ----- + ----- = -------------------
				//   A     a              w
				return qs.aNew(w, 0, C*U, D, rc*u, rd)
			}
		}
	}
	return nil, fmt.Errorf("Can't add pair %s and %s", q, r)
}

// aMul return the multiplication of the two given numbers.
// Limitations:
//	- Both numbers sizes=1 are multiplied ok.
//	- One number size=3 and the other size=1 are multiplied ok.
//	- Both numbers sizes=3 are multiplied ok.
//	- For other sizes combination "Can't mul pair" error is returned.
func (qs *A32s) aMul(q, r *A32) (s *A32, err error) {
	if q == nil || r == nil {
		return nil, nil
	}
	max, min := q, r
	if len(q.num) < len(r.num) {
		max, min = r, q
	}
	switch len(max.num) {
	case 1:
		A, B := max.ab()
		a, b := min.ab()
		//  B * b = Bb 
		return qs.aNew(A*a,
			B*b) // x

	case 3:
		A, B, C, D := max.abcd()
		switch len(min.num) {
		case 1:
			a, b := min.ab()
			// = (B + C√D) * b
			// = Bb + Cb√D
			// =  x +  y√z
			return qs.aNew(A*a,
				B*b,    // x
				C*b, D) // y√z

		case 3:
			a, b, c, d := min.abcd()
			if D < d {
				A, B, C, D = min.abcd()
				a, b, c, d = max.abcd()
			}
			if B == 0 {
				if b == 0 {
					// C√D * c√d = Cc√Dd = y√z
					return qs.aNew(A*a,
						0,        // x
						C*c, D*d) // y/z
				}
				// = C√D * (b + c√d)
				// = Cc√Dd + bC√D
				// =  y√z  +  e√f
				return qs.aNew(A*a,
					0,        // x
					C*c, D*d, // y√z
					b*C, D)   // e√f
			}
			if b == 0 {
				// = (B + C√D) * c√d
				// = Bc√d + Cc√Dd
				// =  y√z +  e√f
				return qs.aNew(A*a,
					0,          // x
					B*c, d,     // y√z
					C*c,   D*d) // e√f
			}
			if D == d {
				// = (B + C√D) * (b * c√D)
				// = Bb + CcD + Bc√D + bC√D
				// = Bb + CcD + (Bc+bC)√D
				// =    x     +    y   √z
				return qs.aNew(A*a,
					B*b + C*c*D,  // x
					B*c + b*C, D) // y√z
			}//  ;fmt.Println("5")
			// = (B + C√D) * (b * c√d) 
			// = Bb + Bc√d + bC√D + Cc√Dd
			// = Bb + Bc√d + (C√D)*(b+c√d)
			// = Bb + Bc√d + (C√D)*√((b+c√d)(b+c√d))
			// = Bb + Bc√d + (C√D)*√(b²+2bc√d+c²d)
			// = Bb + Bc√d + (C√D)*√(b²+c²d + 2bc√d)
			// = Bb + Bc√d + C√(bD²+c²Dd + 2bcD√d)
			// =  x +  y√z + e√(    f    +   g √h)
			return qs.aNew(A*a,
				B*b,             // x
				B*c, d,          // y√z
				C,               // e
				b*D*D + c*c*D*d, // f
				2*b*c*D, d)      // g√h
		}
	}
	return nil, fmt.Errorf("Can't mul pair %s and %s", q, r)
}


// aSqrt returns the square root of the given number.
// Limitations:
//	Square root of numbers of size 1 are returned ok.
//	Square root of numbers of size 3 are returned ok.
//	For other sizes a "Can't square root" error is returned
func (qs *A32s) aSqrt(q *A32) (s *A32, err error) {
	if q == nil {
		return nil, nil
	}
	switch len(q.num) {
	case 1:
		//   / b     √b     √ab    √ab    B + C√D
		//  / --- = ----- = ---- = ---- = -------
		// √   a     √a     √aa     a       A
		a, b := q.ab()
		return qs.aNew(a, // A
			0,            // B
			1, Z(a)*b)    // C√D

	case 3:
		a, b, c, d := q.abcd()
		if b == 0 {
			//   / c√d    √(c√d)   √a√(c√d)   √(ac√d)  B + C√D + E√(F+G√H)
			//  / ----- = ------ = -------- = ------ = ------------------
			// √    a       √a       √a√a        a             A
			return qs.aNew(a, // A
				0,         // B = 0
				0, 1,      // C√D = 0
				1,         // E
				0,         // F
				Z(a)*c, d) // G√H
		}
		// First from √(b + c√d) look if b² - c²d = x²
		// In other words, look a x such that 1√(b²-c²d) = x√1
		// Case example: √(6+2√5) = 1+√5
		if x, r, _ := qs.zSqrt(1, b*b - c*c*d); r == 1 {
			a := N(2)
			// √(b + c√d) = (√(2b+2x) + √(2b-2x))/2
			// √(b - c√d) = (√(2b+2x) - √(2b-2x))/2
			o1, i1, _ := qs.zSqrt(1, 2*(b + Z(x))) // o1√i1 = √(b+x)
			o2, i2, _ := qs.zSqrt(1, 2*(b - Z(x))) // o2√i2 = √(b-x)
			if i1 == +1 && i2 != +1 { // √(b+x) is integer, √(b-x) is not.
				return qs.aNew(a, // C
					Z(o1),        // B
					Z(o2), Z(i2)) // C√D
			}
			if i1 != +1 && i2 == +1 { // √(b+x) is not integer, √(b-x) is.
				return qs.aNew(a, // C
					Z(o2),        // B
					Z(o1), Z(i1)) // C√D
			}
			if i1 >= i2 {
				return qs.aNew(a, // C
					0,            // B
					Z(o1), Z(i1), // C√D
					Z(o2), Z(i2)) // E√F
			}
			return qs.aNew(a, // C
				0,            // B
				Z(o2), Z(i2), // C√D
				Z(o1), Z(i1)) // E√F
		}
		//   / b + c√d    √(b + c√d)   √a√(b + c√d)   √(ab + ac√d)   √(F + G√H))
		//  / --------- = ---------- = ------------ = ------------ = -----------
		// √      a           √a           √a√a             a             A
		return qs.aNew(a, // A
			0,      // B = 0
			0, 1,   // C√D = 0
			1,      // E
			Z(a)*b, // F
			Z(a)*c, // G
			d,      // H
		)
	}
	return nil, fmt.Errorf("Can't square root of %s", q.String())
}

// aPow2 returns the given number raised to the power 2.
// Limitations:
//	Numbers of sizes 1, 3, 5 and 7 are raised ok.
//	For the rest of sizes a "Can't pow2" error is returned.
func (qs *A32s) aPow2(q *A32) (s *A32, err error) {
	if q == nil {
		return nil, nil
	}
	switch len(q.num) {
	case 1:
		// = b*b = b²
		// = B
		a, b := q.ab()
		return qs.aNew(a*a, // A
			b*b)            // B
	
	case 3:
		a, b, c, d := q.abcd()
		if b == 0 {
			// = (c√d)(c√d) = c²d
			// = B
			return qs.aNew(a*a, // A
				c*c*d)          // B
		}
		// x² = (b + c√d)(b + c√d) 
		//    = b² + 2bc√d + c²d
		//    = b² + c²d + 2bc√d
		//    = B + C√D
		return qs.aNew(a*a, // A
			b*b + c*c*d,    // B 
			2*b*c, d)       // C√D

	case 5: 
		a, b, c, d, e, f := q.abcdef()
		if b == 0 {
			// = (c√d + e√f)(c√d + e√f)
			// = c²d + e²f + 2ce√(df)
			// = B + C√D
			return qs.aNew(a*a, // A
				c*c*d + e*e*f,  // B
				2*c*e, d*f)     // C√D
		}
		// = (b + c√d + e√f)(b + c√d + e√f)
		// = (x + e√f)(x + e√f)
		// = x² + 2xe√f + e²f
		// = x² + e²f + 2√(b + c√d)*e√f
		// = x² + e²f + 2e√(bf + cf√d)
		// = b² + c²d + e²f + 2bc√d + 2e√(bf + cf√d)
		return qs.aNew(a*a,      // A
			b*b + c*c*d + e*e*f, // B 
			2*b*c, d,            // C√D
			2*e,                 // E
			b*f,                 // F
			c*f, d)              // G√H

	case 7:
		a, b, c, d, e, f, g, h := q.abcdefgh()
		if b == 0 {
			if c == 0 {
				// = e√(f+g√h) * e√(f+g√h)
				// = e²(f + g√h)
				// = e²f + e²g√h
				// = B + C√D
				return qs.aNew(a*a, // A
					e*e*f,          // B
					e*e*g, h)       // C√D
			} else {
				// = (c√d + e√(f+g√h)) * (c√d + e√(f+g√h))
				// = c²d + 2ce√(df + dg√h) + e²(f + g√h)
				// = c²d + e²f + e²g√h + 2ce√(df + dg√h)
				// = B + C√D + E√(F + G√H)
				return qs.aNew(a*a, // A
					c*c*d + e*e*f,  // B
					e*e*g, h,       // C√D
					2*c*e,          // E
					d*f,            // F
					d*g, h)         // G√H
			}
		} else {
			if c == 0 { // b != 0 and c == 0 is like b = 0 and c=b d=1
				// = (b + e√(f+g√h)) * (b + e√(f+g√h))
				// = b² + 2be√(f+g√h) + e²(f+g√h)
				// = b² + e²f + e²g√h + 2be√(f+g√h)
				// = B + C√D + E√(F + G√H)
				return qs.aNew(a*a, // A
					b*b + e*e*f,    // B
					e*e*g, h,       // C√D
					2*b*e,          // E
					f,              // F
					g, h)           // G√H
			}
		}
	}
	return nil, fmt.Errorf("Can't pow2 of %s", q.String())
}









