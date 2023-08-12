package alg

import (
	"testing"
)

func TestQ32_(t *testing.T) {

	q := newQ32(0, 0)
	// bcd
	for _, r := range []struct { x, y, z Z; exp string } {
		{ 0, 0, 0, "0"       },
		{ 0, 0, 1, "0"       },
		{ 0, 1, 1, "1"       },
		{ 0, 1, 4, "√4"      },
		{ 0, 2, 4, "2√4"     },
		{ 1, 2, 4, "1+2√4"   },
		{ 2, 2,-4, "2+2i√4"  },
		{ 3,-2,-4, "3-2i√4"  },
		{-4,-2,-4, "-4-2i√4" },
		{-5,-2,-1, "-5-2i"   },
		{-5,-1,-1, "-5-i"    },
		{-6,-2, 0, "-6"      },
		{-7, 0, 0, "-7"      },
		{-8, 2, 1, "-8+2"    },
	} {
		s := &Str{}
		q.bcd(s, r.x, r.y, r.z)
		if got := s.String(); got != r.exp {
			t.Fatalf("bcd got %s exp %s", got, r.exp)
		}
	}

	// Test without simplifications done with a factory
	for _, s := range []struct { q *Q32; exp string } {
		{ newQ32(0),             "NaN"     },
		{ newQ32(1),             "Invalid" },
		{ newQ32(1,2,3),         "Invalid" },
		{ newQ32(1,2,3,4,5),     "Invalid" },
		{ newQ32(1,2,3,4,5,6,7), "Invalid" },

		{ newQ32(0,1), "NaN" },
		{ newQ32(7,0), "0"   },
		{ newQ32(7,7), "7/7" },
		{ newQ32(7,1), "1/7" },
		{ newQ32(1,7), "7"   },

		{ newQ32(1,2,3,4), "2+3√4"     },
		{ newQ32(2,3,4,5), "(3+4√5)/2" },
		{ newQ32(2,0,1,1), "1/2"       },
		{ newQ32(1,0,1,1), "1"         },
		{ newQ32(1,0,0,0), "0"         },
		{ newQ32(1,2,3,0), "2"         },

		// len(num)=5
		{ newQ32(1,2,3,4,5,6), "2+3√4+5√6"     },

		{ newQ32(2,3,4,5,6,7), "(3+4√5+6√7)/2" },

		{ newQ32(2,0,4,5,6,7), "(4√5+6√7)/2"   }, // b=0
		{ newQ32(2,3,4,0,6,7), "(3+6√7)/2"     }, // d=0
		{ newQ32(2,3,0,5,6,7), "(3+6√7)/2"     }, // c=0

		{ newQ32(2,3,4,5,0,7), "(3+4√5)/2"     }, // e=0
		{ newQ32(2,3,4,5,6,0), "(3+4√5)/2"     }, // f=0

		{ newQ32(2,0,4,0,6,7), "6√7/2"         }, // b=d=0
		{ newQ32(2,0,4,5,0,7), "4√5/2"         }, // b=e=0
		{ newQ32(2,3,0,5,0,7), "3/2"           }, // c=e=0
		{ newQ32(2,0,0,5,0,7), "0"             }, // b=c=e=0

		{ newQ32(2,3,4,+1,6,7), "(3+4+6√7)/2"   }, // d=+1
		{ newQ32(2,3,4,-1,6,7), "(3+4i+6√7)/2"  }, // d=-1
		{ newQ32(2,0,4,5,6,+1), "(4√5+6)/2"     }, // f=+1
		{ newQ32(2,0,4,5,6,-1), "(4√5+6i)/2"    }, // f=-1

		{ newQ32(2, 2, 2, 2, 2, 2), "(2+2√2+2√2)/2" },
		{ newQ32(1,-1,-1,-1,-1,-1), "-1-i-i" },
		{ newQ32(1,-1,-1,+1,-1,+1), "-1-1-1" },

		// len(num)=7
		{ newQ32(2,3,4,5,6,7,8,9), "(3+4√5+6√(7+8√9))/2" },
		
		{ newQ32(2,3,4,5,+1,7,8,9), "(3+4√5+√(7+8√9))/2" }, // e=+1
		{ newQ32(2,3,4,5,-1,7,8,9), "(3+4√5-√(7+8√9))/2" }, // e=-1

		{ newQ32(2,3,4,5,6,7,8,0), "(3+4√5+6√(7))/2" }, // h=0
		{ newQ32(2,3,4,5,6,7,0,9), "(3+4√5+6√(7))/2" }, // g=0

		{ newQ32(2,3,4,5,6,7,8,1), "(3+4√5+6√(7+8))/2"  }, // h=1
		{ newQ32(2,3,4,5,6,7,1,9), "(3+4√5+6√(7+√9))/2" }, // g=1
		{ newQ32(2,3,4,5,6,0,8,9), "(3+4√5+6√(8√9))/2"  }, // f=0

		{ newQ32(2,3,4,5,6,0,0,9), "(3+4√5)/2"  }, // f=g=0
		{ newQ32(2,3,4,5,6,0,8,0), "(3+4√5)/2"  }, // f=h=0

	} {
		if got := s.q.String(); got != s.exp {
			t.Fatalf("newQ32 got %s exp %s", got, s.exp)
		}
	}
}

func TestQ32s(t *testing.T) {

	f := func(p ...Z) []Z { return p }

	qs := NewQ32s()

	m := make(map[string]*Q32, 0)

	// qNew
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

		{ f(9,8,7,6,5,4,3), 2, "(9+8√7+6√(5+4√3))/2" },
		{ f(0,8,7,6,5,4,3), 2, "4√7+3√(5+4√3)" },
		{ f(0,1,7,6,5,4,3), 2, "(√7+6√(5+4√3))/2" },
		{ f(0,1,1,6,5,4,3), 2, "(1+6√(5+4√3))/2" },
		{ f(0,0,1,6,5,4,3), 2, "3√(5+4√3)" },
		{ f(0,0,1,1,5,4,3), 2, "√(5+4√3)/2" },
		{ f(0,0,1,1,0,4,3), 2, "√(√3)" },
		{ f(0,0,1,1,0,1,3), 2, "√(√3)/2" },
		{ f(0,0,1,1,0,1,1), 2, "1/2" },
		{ f(0,0,1,1,0,0,1), 2, "0" },

	} {
		if q, err := qs.qNew(s.den, s.nums...); err != nil {
			if s.exp != "err" {
				t.Fatalf("reduceQ unexpected overflow for %d %v", s.den, s.nums)
			}
		} else if got := q.String(); got != s.exp {
			t.Fatalf("qs.new got %s exp %s", got, s.exp)
		} else {
			m[got] = q // feed map to be used in multiplications
		}
	}

	// qAdd
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
		if a, ok := m[s.a]; !ok {
			t.Fatalf("qs.AddQ error map has no %s", s.a)
		} else if b, ok := m[s.b]; !ok {
			t.Fatalf("qs.AddQ error map has no %s", s.b)
		} else if r, err := qs.qAdd(a, b); err != nil {
			t.Fatalf("qs.AddQ error for %s %s %v", s.a, s.b, err)
		} else if got := r.String(); got != s.exp {
			t.Fatalf("qs.AddQ got %s exp %s", got, s.exp)
		}
	}

	// qMul
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
		if r, err := qs.qMul(a, b); err != nil {
			t.Fatalf("qs.MulQ error for %s %s %v", s.a, s.b, err)
		} else if got := r.String(); got != s.exp {
			t.Fatalf("qs.MulQ got %s exp %s", got, s.exp)
		}
	}

	// qSqrt
	for _, s := range []struct { q, exp string } {
		{ "0",   "0"     },
		{ "1",   "1"     },
		{ "3",   "√3"    },
		{ "5/2", "√10/2" },
		{ "7/5", "√35/5" },

		//{ "1+√5", "?"},
	} {
		q := m[s.q]
		if r, err := qs.qSqrt(q); err != nil {
			t.Fatalf("qs.sqrt error for %s %v", s.q, err)
		} else if got := r.String(); got != s.exp {
			t.Fatalf("qs.sqrt got %s exp %s", got, s.exp)
		}
	}
}




	/*// roie
	/*
	// Reduce general
	for _, s := range []struct{ exp string; is []int64 } {
		{ "", f(1) }, // 1 is invalid
		{ "", f(1,1,1) }, // 3 is invalid
		{ "", f(1,1,1,1,1) }, // 5 is invalid
		{ "", f(1,1,1,1,1,1) }, // 6 is invalid
		{ "", f(1,1,1,1,1,1,1) }, // 7 is invalid


		{ "+1√(1+1+1√(2+0))",     f(1,1,1,1,1,1,1,1) }, // TODO
		{ "+2√(2+2√2+2√(2+2√2))", f(2,2,2,2,2,2,2,2) },
		{ "+3√(3+3√3+3√(3+3√3))", f(3,3,3,3,3,3,3,3) },
		{ "+8√(1+2+2√(3+0))",     f(4,4,4,4,4,4,4,4) }, // TODO

		// +1√(1+1√1+1√(1+1√1))
		// +1√(1+1√1+1√(1+1))
		// +1√(1+1√1+1√2)
		// +1√(1+1+1√2)
		// +1√(2+1√2)

		// +4√(4+4√4+4√(4+4√4))
		// +4√(4+4√4+4√(4+8))
		// +4√(4+4√4+4√12)
		// +4√(4+4√4+8√3)
		// +4√(4+8+8√3)
		// +8√(1+2+2√3)
		// +8√(3+2√3)

	} {
		if r, overflow := factory.Reduce(s.is...); overflow {
			if s.exp != "∞" {
				t.Fatalf("Reduce unexpected overflow for %+v", s.is)
			}
		} else if got := newAZ32(r...).String(); got != s.exp {
			t.Fatalf("Reduce get %s exp %s", got, s.exp)
		}
	}
	*/
	//t.Logf("22222222 %v", newAZ32(2,2,2,2,2,2,2,2))
