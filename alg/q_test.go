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

		// (b + c√d)/a		
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

		{ f(0,147,11),2500, "147√11/2500" },
		{ f(0, 6,33),  99,  "2√33/33"     },

		// (b + c√d e√f)/a

		{ f(1,1,1,1,1), 0, "err"  },

		{ f(0,1,2,3,4), 5, "(6+√2)/5"  }, // (1√2+3√4)/5 = (1√2+6)/5 = (6+√2)/5
		{ f(1,0,2,3,4), 5, "7/5"       }, // (1+0√2+3√4)/5 = (1+3√4)/5 = (1+6)/5 = 7/5
		{ f(1,2,0,3,4), 5, "7/5"       }, // (1+2√0+3√4)/5 = (1+3√4)/5 = (1+6)/5 = 7/5
		{ f(1,2,3,0,4), 5, "(1+2√3)/5" }, // (1+2√3+0√4)/5 = (1+2√3)/5
		{ f(1,2,3,4,0), 5, "(1+2√3)/5" }, // (1+2√3+0√4)/5 = (1+2√3)/5

		{ f(1,2,3,4,5), 6, "(1+2√3+4√5)/6" }, // (1+2√3+4√5)/6 
		{ f(1,2,3,0,5), 6, "(1+2√3)/6"     }, // (1+2√3+0√5)/6 = (1+2√3)/6  f==0
		{ f(1,2,3,4,0), 6, "(1+2√3)/6"     }, // (1+2√3+4√0)/6 = (1+2√3)/6  f==0
		{ f(1,2,3,4,1), 6, "(5+2√3)/6"     }, // (1+2√3+4)/6 = (5+2√3)/6    f==1
		{ f(1,2,3,4,3), 6, "(1+6√3)/6"     }, // (1+2√3+4√3)/6 = (1+6√3)/6  d==f==3
		{ f(1,2,1,4,3), 6, "(3+4√3)/6"     }, // (1+2√1+4√3)/6 = (3+4√3)/6  c=1

		{ f(1,1,1,1,1), 1, "3"     }, // (1+1√1+1√1)/1 = (1+1+1)/1 = 3
		{ f(2,2,2,2,2), 2, "1+2√2" }, // (2+2√2+2√2)/2 = (2+4√2)/2 = 1+2√2
		{ f(3,3,3,3,3), 3, "1+2√3" }, // (3+3√3+3√3)/3 = (3+6√3)/3 = 1+2√2
		{ f(4,4,4,4,4), 4, "5"     }, // (4+4√4+4√4)/4 = (4+8+8)/4 = 5
		{ f(5,5,5,5,5), 5, "1+2√5" },
		
		{ f(2,3,4,5,6), 7, "(8+5√6)/7"      }, // (2+3√4+5√6)/7 = (2+6+5√6)/7 = (8+5√6)/7
		{ f(3,4,5,6,7), 8, "(3+4√5+6√7)/8"  }, // (3+4√5+6√7)/8
		{ f(4,5,6,7,8), 9, "(4+14√2+5√6)/9" }, // (4+5√6+7√8)/9 = (4+5√6+14√2)/9

		{ f(6,5,4,3,2), 1, "16+3√2"         }, // (6+5√4+3√2)/1 = (6+10+3√2)/1 = 16+3√2
		{ f(7,6,5,4,3), 2, "(7+4√3+6√5)/2"  }, // (7+6√5+4√3)/2
		{ f(8,7,6,5,4), 3, "(18+7√6)/3"     }, // (8+7√6+5√4)/3 = (8+7√6+10)/3 = (18+7√6)/3
		{ f(9,8,7,6,5), 4, "(9+6√5+8√7)/4"  }, // (9+8√7+6√5)/4

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

	// add
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

		{ "147√11/2500", "147√11/2500", "147√11/1250" }, // big

		// both bs=0
		{ "√2",    "2√5/3",       "(3√2+2√5)/3"           }, // (3√2)/3 + (2√5)/3 = (3√2+2√5)/3
		{ "√2",    "√2",          "2√2"                   },
		{ "√2",    "147√11/2500", "(2500√2+147√11)/2500"  },
		{ "2√5/3", "147√11/2500", "(5000√5+441√11)/7500"  },
		{ "2√5/3", "2√33/33",     "(22√5+2√33)/33"        },
		{ "√2",    "2√33/33",     "(33√2+2√33)/33"        },

	} {
		a := m[s.a]
		b := m[s.b]
		if r, err := qs.AddQ(a, b); err != nil {
			t.Fatalf("qs.AddQ error for %s %s %v", s.a, s.b, err)
		} else if got := r.String(); got != s.exp {
			t.Fatalf("qs.AddQ got %s exp %s", got, s.exp)
		}
	}

}