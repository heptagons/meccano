package angles

import (
	"fmt"
)

type Rat struct {
	Neg bool // negative when true
	Num Nat  // natural numerator
	Den Nat  // natural denominator
}

// NewRat creates a new quotient when n and d are numerator and denominator.
// Greatest common divisor set N and D at minimum.
func NewRat(n, d int) (q *Rat) {
	// return nil as NaN or +/- infinity
	if d == 0 {
		return
	}
	q = &Rat{}
	// set negative sign and convert to uint
	num := Nat(0)
	den := Nat(0)
	if n < 0 {
		if d > 0 {
			q.Neg = true
		}
		num = Nat(-n)
	} else {
		num = Nat(n)
	}
	if d < 0 {
		if n > 0 {
			q.Neg = true
		}
		den = Nat(-d)
	} else {
		den = Nat(d)
	}
	// return zero
	if num == 0 {
		q.Den = den
		return
	}
	// greatest common divisor (GCD) via Euclidean algorithm
	a, b := num, den
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	// set reduced num/dem
	q.Num = num / a
	q.Den = den / a
	return
}

// NewRatCosZ returns the cosines law:
// cosC = (a² + b² - c²) / 2ab
func NewRatCosC(a, b, c int) *Rat {
	return NewRat(a*a + b*b - c*c, 2*a*b)
}

// NewRatSin2Z return the sines law:
// sinC = math.Sqrt(4a²b² - (a²+b²-c²)²) / 2ab
// TODO: Return algebraic to the squared rational version!!!
func NewRatSin2C(a, b, c int) *Rat {
	m := 4*a*a*b*b
	n := (a*a + b*b - c*c)
	return NewRat(m - n*n, 4*a*a*b*b)
}

// Times returns a new Rat with the addition of this quotient q with given quotient p.
func (q *Rat) Add(p *Rat) *Rat {
	if q == nil || p == nil {
		return nil
	}
	n1 := int(q.Num * p.Den)
	if q.Neg {
		n1 *= -1
	}
	n2 := int(p.Num * q.Den)
	if p.Neg {
		n2 *= -1
	}
	return NewRat(n1 + n2, int(q.Den * p.Den))
}

// Mul returns new Rat with the multiplication of this quotient q with given quotient p.
func (q *Rat) Mul(p *Rat) *Rat {
	if q == nil || p == nil {
		return nil
	}
	if q.Num == 0 || p.Num == 0 {
		return &Rat{ Den:1 }
	}
	n1 := int(q.Num)
	if q.Neg {
		n1 *= -1
	}
	n2 := int(p.Num)
	if p.Neg {
		n2 *= -1
	}
	return NewRat(n1 * n2, int(q.Den * p.Den))
}

// Negate returns new Rat with the sign changed.
func (r *Rat) Negate() *Rat {
	if r == nil {
		return nil
	}
	return &Rat{
		Neg: !r.Neg,
		Num: r.Num,
		Den: r.Den,
	}
}

// Invert returns new Rat with the numerator and denominator reversed
func (q *Rat) Invert() *Rat {
	if q == nil {
		return nil
	}
	if q.Num == 0 {
		return nil
	}
	return &Rat{
		Neg: q.Neg,
		Num: q.Den,
		Den: q.Num,
	}
}

func (q *Rat) String() string {
	if q == nil {
		return ""
	}
	if q.Num == 0 {
		return "0"
	}
	neg := ""
	if q.Neg {
		neg = "-"
	}
	if q.Den == 1 {
		return fmt.Sprintf("%s%d", neg, q.Num)	
	}
	return fmt.Sprintf("%s%d/%d", neg, q.Num, q.Den)
}
