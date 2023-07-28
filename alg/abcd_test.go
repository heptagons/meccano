package alg

import (
	"fmt"
	"testing"
)

func TestZ(t *testing.T) {
	for _, u := range []struct{ a, b, c Z; exp string } {
		{ a:0,    b:0,    c:0,    exp:"0,0,0" },
		{ a:1,    b:0,    c:0,    exp:"1,0,0" },
		{ a:1,    b:1,    c:0,    exp:"1,1,0" },
		{ a:1,    b:1,    c:1,    exp:"1,1,1" },
		{ a:19,   b:23,   c:29,   exp:"19,23,29" },
		{ a:1001, b:1001, c:1001, exp:"1,1,1" },
		{ a:3003, b:1001, c:7007, exp:"3,1,7" },
		{ a:1000, b:1001, c:1002, exp:"1000,1001,1002" },
		{ a:-30,  b:+45,  c:-60,  exp:"-2,3,-4" },
		{ a:2,    b:0,    c:1,    exp:"2,0,1" },
	} {
		(&u.a).Reduce3(&u.b, &u.c)
		if got := fmt.Sprintf("%d,%d,%d", u.a, u.b, u.c); got != u.exp {
			t.Fatalf("Z.Reduce3 got %s exp %s", got, u.exp)
		}
	}
}

func TestB(t *testing.T) {

	none     := NewB( 0,0)
	zero     := NewB( 0,1)
	minus1   := NewB(-1,1)
	half     := NewB( 2,4)
	ten      := NewB( 10,1)
	one_17th := NewB( 2*3*5*7*11*13, 2*3*5*7*11*13*17)
	for exp, b := range map[string]*B {
		"∞":     none,
		"+0":    zero,
		"-1":    minus1,
		"+1/2":  half,
		"-7/3":  NewB(-2*5*7*11, 2*3*5*11),
		"+1/17": one_17th,
		"+10":   ten,
	} {
		if got := b.String(); got != exp {
			t.Fatalf("B got %s exp %s", got, exp)
		}
	}
	// AddB
	for exp, b := range map[string]*B {
		"∞":      zero.AddB(none),
		"+0":     zero.AddB(zero),
		"-1":     zero.AddB(minus1),
		"-2":     minus1.AddB(minus1),
		"+1":     half.AddB(half),
		"+2/17":  one_17th.AddB(one_17th),
		"+19/34": half.AddB(one_17th),
		"+21/34": half.AddB(one_17th).AddB(one_17th),
		"+19/17": half.AddB(one_17th).AddB(one_17th).AddB(half),
		"+21/2":  ten.AddB(half),
	} {
		if got := b.String(); got != exp {
			t.Fatalf("B-add got %s exp %s", got, exp)
		}
	}
	// InvB
	for _, b := range []struct { e string; b *B } {
		{ e:"∞",     b: none.Inv() },
		{ e:"∞",     b: zero.Inv() },
		{ e:"-1",    b: minus1.Inv() },
		{ e:"-1",    b: minus1.Inv().Inv().Inv() },
		{ e:"+2",    b: half.Inv() },
		{ e:"-3/7",  b: NewB(-2*5*7*11, 2*3*5*11).Inv() },
		{ e:"+17",   b: one_17th.Inv() },
		{ e:"+1/10", b: ten.Inv() },
		{ e:"+10",   b: ten.Inv().Inv() },
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B-inv got %s exp %s", got, b.e)
		}
	}

	// MulB
	for _, b := range []struct { e string; b *B } {
		{ e:"∞",        b: zero.MulB(none) },
		{ e:"+0",       b: zero.MulB(zero) },
		{ e:"+0",       b: zero.MulB(minus1) },
		{ e:"+1",       b: minus1.MulB(minus1) },
		{ e:"+1/4",     b: half.MulB(half) },
		{ e:"+1/289",   b: one_17th.MulB(one_17th) },
		{ e:"+1/34",    b: half.MulB(one_17th) },
		{ e:"+1/578",   b: half.MulB(one_17th).MulB(one_17th) },
		{ e:"+1/1156",  b: half.MulB(one_17th).MulB(one_17th).MulB(half) },
		{ e:"+5",       b: ten.MulB(half) },
		{ e:"-12/5",    b: NewB(-144,130).MulB(NewB(26,12)) },
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B-mul got %s exp %s", got, b.e)
		}
	}
}

func TestD(t *testing.T) {
	a1 := Z(1)
	a2 := Z(2)
	n32s := NewN32s()
	infinite := NewD(n32s,   0, 0,0,  0)
	zero     := NewD(n32s,   0, 0,0, a1)
	minus1   := NewD(n32s,  -1, 0,0, a1)
	half     := NewD(n32s,   1, 0,0, a2)
	ten      := NewD(n32s,  10, 0,0, a1)
	one_17th := NewD(n32s,  2*3*5*7*11*13, 0,0, 2*3*5*7*11*13*17)

	for p, r := range []struct { d *D; exp string } {
		{ d: infinite, exp:"∞" },
		{ d: zero,     exp:"+0" },
		{ d: minus1,   exp:"-1" },
		{ d: half,     exp:"+1/2" },
		{ d: ten,      exp:"+10" },
		{ d: one_17th, exp:"+1/17" },

		{ d: NewD(n32s, 0, 1,2, a1),   exp: "+1√2"    },
		{ d: NewD(n32s, 0, 1,2, a2),   exp:"+1√2/2" },
	} {
		if p >= 7 {
			fmt.Printf("b=%v out=%v in=%v a=%v\n", r.d.ab.b, r.d.cd.out, r.d.cd.in, r.d.ab.a)
		}

		if got := r.d.String(); got != r.exp {
			t.Fatalf("D-new got %s exp %s", got, r.exp)	
		}
	}
}