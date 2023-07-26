package alg

import (
	"fmt"
)

type B struct { // Rational
	b *I  // numerator integer optional
	a N32 // denominator natural >= 1
}

func NewB0() *B {
	return &B{
		a:1,
	}
}

func NewB(num Z, den Z) *B {
	if den == 0 {
		return nil // infinite
	}
	if num == 0 {
		return NewB0() // zero
	}
	s := false
	if num < 0 {
		num = -num
		s = !s
	}
	if den < 0 {
		den = -den
		s = !s
	}
	g := gcd(num, den)
	num /= g
	den /= g
	if num > MaxN {
		return nil // numerator overflow
	} else if den > MaxN {
		return nil // denominator overflow
	}
	var b *I
	if s { // fast negative
		b = newIminus(N32(num))
	} else { // fast positive
		b = newIplus(N32(num))
	}
	return &B{
		b: b,
		a: N32(den),
	}
}

func NewBplus(num, den N32) *B {
	return newB(false, num, den)
}

func NewBminus(num, den N32) *B {
	return newB(true, num, den)	
}

func newB(s bool, num, den N32) *B {
	if den == 0 {
		return nil // infinity
	}
	if num == 0 {
		return NewB0() // zero
	}
	n, d := num, den
	g := n.gcd(d)
	return &B{
		b: &I{ s:s, n: n / g },
		a: d / g,
	}
}

func (x *B) clone() *B {
	if x == nil || x.a == 0 {
		return nil // infinite
	} else if x.b == nil || x.b.n == 0 {
		return NewB0()
	} else {
		return newB(x.b.s, x.b.n, x.a)
	}
}

func (x *B) AddB(y *B) *B {
	if x == nil || y == nil || x.a == 0 || y.a == 0 {
		return nil // infinite
	} else if x.b == nil || x.b.n == 0 {
		return y.clone() // y
	} else if y.b == nil || y.b.n == 0 {
		return x.clone() // x
	}
	num := x.b.mul(y.a) + y.b.mul(x.a)
	den := Z(x.a) * Z(y.a)
	return NewB(num, den)
}

func (x *B) Inv() *B {
	if x == nil || x.a == 0 || x.b == nil || x.b.n == 0 {
		return nil
	} 
	return &B{
		b: &I{ s:x.b.s, n: x.a },
		a: x.b.n,
	}
}

func (x *B) MulB(y *B) *B {
	if x == nil || y == nil {
		return nil
	} else if x.b == nil || x.b.n == 0 {
		return NewB0()
	} else if y.b == nil || y.b.n == 0 {
		return NewB0()
	}
	num := Z(x.b.n) * Z(y.b.n)
	if x.b.s != y.b.s {
		num = -num
	}
	den := Z(x.a) * Z(y.a)
	return NewB(num, den)
}

func (x *B) String() string {
	if x == nil || x.a == 0 {
		return "" // infinity
	} else if x.b == nil || x.b.n == 0 {
		return "0"
	} else if x.a == 1 {
		return x.b.string()
	} else {
		return fmt.Sprintf("%s/%d", x.b.string(), x.a)
	}
}
