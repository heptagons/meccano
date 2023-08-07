package alg

import (
	"testing"
)

func TestQ32(t *testing.T) {

	factory := NewN32s() // with primes for reductions

	f := func(p ...Z) []Z { return p }

	qs := NewQ32s(factory)

	m := map[string]*Q32 {

	}

	for _, s := range []struct { nums []Z ; den N; exp string; } {
		{ f(),  1, "err" },
		{ f(0), 0, "err" },
		{ f(0), 1, "0"   },
		{ f(4), 2, "2"   },
		{ f(1), 2, "1/2" },

		{ f( 1,2),   2, "err"  },
		{ f( 0,1,1), 1, "1"    }, // (0+1√1)/1 = (1)/1 = 1
		{ f( 1,1,1), 1, "2"    }, // (1+1√1)/1 = (1+1)/1 = 2
		{ f( 0,1,2), 1, "√2"   }, // (0+1√2)/1 = (√2)/1 = √2
		{ f( 1,1,2), 1, "1+√2" }, // 1+1√2
		{ f( 1,1,3), 1, "1+√3" }, // 1+1√3
		{ f( 1,1,4), 1, "3"    }, // 1+1√4 = 1+2 = 3
		{ f( 1,1,5), 1, "1+√5" }, // 1+1√5
		
		{ f( 1,1,5), 2, "(1+√5)/2" }, // (1+1√5)/2
		{ f( 1,1,6), 2, "(1+√6)/2" }, // (1+1√5)/2
		{ f( 2,1,7), 2, "(2+√7)/2" }, // (1+1√5)/2
		{ f( 2,1,8), 2, "1+√2"     }, // (2+2√2)/2 = 1+√2
		{ f( 2,1,9), 2, "5/2"      }, // (2+1√9)/2 = (2+3)/2 = 5/2
		{ f(-2,1,9), 2, "1/2"      }, // (-2+1√9)/2 = (-2+3)/2 = 1/2

		{ f( 0,2,1), 2,  "1"       }, // (2√1)/2 = 1
		{ f( 0,2,2), 2,  "√2"      }, // (2√2)/2 = √2
		{ f( 0,4,5), 6,  "2√5/3"   }, // (0+4√5)/6 = 4√5/6 = 2√5/3
	} {
		if q, err := qs.newQ32(s.den, s.nums...); err != nil {
			if s.exp != "err" {
				t.Fatalf("reduceQ unexpected overflow for %d %v", s.den, s.nums)
			}
		} else if got := q.String(); got != s.exp {
			t.Fatalf("qs.new got %s exp %s", got, s.exp)
		} else {
			m[got] = q // feed map to be used in multiplications
		}
	}

	// mul
	for _, s := range []struct { a, b, exp string } {
		{ "0",   "√2",  "0"   },
		{ "1",   "1",   "1"   },
		{ "1",   "2",   "2"   },
		{ "3",   "3",   "9"   },
		{ "5/2", "1/2", "5/4" },
		{ "1",   "√2",  "√2"  },
		{ "√2",  "√2",  "2"   },
		
		{ "1+√5", "1+√5",     "6+2√5"      },
		{ "1+√5", "(1+√5)/2", "3+√5"       },
		{ "1+√5", "2√5/3",    "(10+2√5)/3" },
	} {
		a := m[s.a]
		b := m[s.b]
		if r, err := qs.MulQ(a, b); err != nil {
			t.Fatalf("qs.MulQ error for %s %s %v", s.a, s.b, err)
		} else if got := r.String(); got != s.exp {
			t.Fatalf("qs.MulQ got %s exp %s", got, s.exp)
		}
	}

	// mul
	for _, s := range []struct { a, b, exp string } {
		{ "0",   "√2",  "√2"   },
		{ "1",   "1",   "2"    },
		{ "1",   "2",   "3"    },
		{ "3",   "3",   "6"    },
		{ "5/2", "1/2", "3"    },
		{ "1",   "√2",  "1+√2" },
		{ "√2",  "√2",  "2√2"  },
		
		{ "1+√5", "1+√5",     "2+2√5"     },
		{ "1+√5", "(1+√5)/2", "(3+3√5)/2" },
		{ "1+√5", "2√5/3",    "(3+5√5)/3" },
	} {
		a := m[s.a]
		b := m[s.b]
		if r, err := qs.AddQ(a, b); err != nil {
			t.Fatalf("qs.MulQ error for %s %s %v", s.a, s.b, err)
		} else if got := r.String(); got != s.exp {
			t.Fatalf("qs.MulQ got %s exp %s", got, s.exp)
		}
	}

}