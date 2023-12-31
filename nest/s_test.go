package nest

import (
	"testing"
)

func TestS32s(t *testing.T) {
	factory := NewZ32s()
	s1 := NewS32s(factory)
	s1.sAdd(2)  // √2
	s1.sAdd(12) // √2 + √12 = √2 + 2√3
	s1.sAdd(8)  // √2 + 2√3 + √8 = √2 + 2√3 + 2√2 = 3√2 + 2√3
	if got, exp := s1.String(), "3√2+2√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s2 := NewS32s(factory)
	s2.sAdd(2) // √2
	s2.sAdd(3) // √2+√3
	if s3, err := s2.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "5+2√6"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	} else if s4, err := s3.sNewSqrt(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s4.String(), "√2+√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s2.sAdd(5) // √2+√3+√5 
	if s3, err := s2.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "10+2√6+2√10+2√15"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}
	s2.sAdd(7) // √2+√3+√5+√7
	if s3, err := s2.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "17+2√6+2√10+2√14+2√15+2√21+2√35"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	// test only basic squareables: base -> pow2 -> sqrt -> base
	for _, r := range []struct { base, pow2 string; surds []Z } {
		{ "1-√5",       "6-2√5", []Z{ 1, -5      } },
		{ "4+7√11", "555+56√11", []Z{ 16, 7*7*11 } },
		{ "4-7√11", "555-56√11", []Z{ 16, -7*7*11 } }, // -4+7√11
	} {
		s := NewS32s(factory)
		if err := s.sAddZ(r.surds); err != nil {
			t.Fatalf("error %v", err)
		} else if got, exp := s.String(), r.base; got != exp {
			t.Fatalf("got %s exp %s", got, exp)
		} else if s1, err := s.sNewPow2(); err != nil {
			t.Fatalf("err %v", err)		
		} else if got, exp := s1.String(), r.pow2; got != exp {
			t.Fatalf("got %s exp %s", got, exp)
		} else if s2, err := s1.sNewSqrt(); err != nil {
			t.Fatalf("err %v", err)
		} else if got, exp := s2.String(), r.base; got != exp {
			t.Fatalf("got %s exp %s", got, exp)
		}
	}

	// test triple sums pow2
	for _, r := range []struct { a, b, c N; pow2 string } {
		{   2,   3,    5, "10+2√6+2√10+2√15"      }, 
		{   6,  10,   15, "31+10√6+6√10+4√15"     }, 
		{ 4*3, 9*5, 16*7, "169+12√15+16√21+24√35" }, // 2√3 + 3√5 + 4√7
	} {
		s := NewS32s(factory)
		s.sAdd(r.a)
		s.sAdd(r.b)
		s.sAdd(r.c)
		if s1, err := s.sNewPow2(); err != nil {
			t.Fatalf("err %v", err)		
		} else if got, exp := s1.String(), r.pow2; got != exp {
			t.Fatalf("got %s exp %s", got, exp)
		}
	}

	// test sFloorCeil simple
	for _, r := range []struct { surds []Z; floor, ceil Z32 } {
		{ []Z{   2,   3 },        2,  4 }, // √2+√3        = 3.146...  floor=2 ceil=4
		{ []Z{   3,   8 },        3,  5 }, // √3+√8        = 4.560...  floor=3 ceil=5
		{ []Z{   5,   7 },        4,  6 }, // √5+√7        = 4.882...  floor=4 ceil=6

		{ []Z{   3,   5,    7 },  5,  8 }, // √3+√5+√7     = 6.614...  floor=5  ceil=8
		{ []Z{ 4*3, 9*5, 16*7 }, 19, 22 }, // 2√3+3√5+4√7  = 20.755... floor=19 ceil=22

		{ []Z{ 3, 5, 7, 11 },     8, 12 }, // √3+√5+√7+√11 = 9.930...  floor=8 ceil=12
	} {
		s := NewS32s(factory)
		if err := s.sAddZ(r.surds); err != nil {
			t.Fatalf("error %v", err)
		} else if floor, ceil, err := s.sFloorCeil(); err != nil {
			t.Fatalf("error %v", err)
		} else if floor != r.floor || ceil != r.ceil {
			t.Fatalf("floor got %d exp %d, ceil got %d exp %d", floor, r.floor, ceil, r.ceil)
		} else {
			t.Logf("%s = %2.3f... floor=%d ceil=%d", s.String(), s.sFloat(), floor, ceil)
		}
	}

	// test sFloorCeil pow2 -> sqrt
	for _, r := range []struct { surds []Z; floor, ceil Z32 } {
		{ []Z{   2,   3 },        3,  4 }, // √2+√3        = 3.146...  floor=3,  ceil=4
		{ []Z{   3,   8 },        4,  5 }, // √3+√8        = 4.560...  floor=4,  ceil=5
		{ []Z{   5,   7 },        4,  5 }, // √5+√7        = 4.882...  floor=4,  ceil=5

		{ []Z{   3,   5,    7 },  6,  7 }, // √3+√5+√7     = 6.614...  floor=6,  ceil=7
		{ []Z{ 4*3, 9*5, 16*7 }, 20, 21 }, // 2√3+3√5+4√7  = 20.755... floor=20, ceil=21

		{ []Z{ 3, 5, 7, 11 },     9, 11 }, // √3+√5+√7+√11 = 9.930...  floor=9 ceil=11

	} {
		s := NewS32s(factory)
		if err := s.sAddZ(r.surds); err != nil {
			t.Fatalf("error %v", err)
		} else if floor, ceil, err := s.sFloorCeil2(); err != nil {
			t.Fatalf("error %v", err)
		} else if floor != r.floor || ceil != r.ceil {
			t.Fatalf("floor got %d exp %d, ceil got %d exp %d", floor, r.floor, ceil, r.ceil)
		} else {
			t.Logf("%s = %2.3f... floor=%d ceil=%d", s.String(), s.sFloat(), floor, ceil)
		}
	}
}

/*
 √2 + √3  <  √4 + √9 
          <  2+3 = 5
 √3 + √8  <  √4 + √9
          <  2+3 = 5
 √5 + √7  < √9 + √9
          < 3+3 = 6

 √3 + √5 + √7 < √4 + √9 + √9
              < 2 + 3 + 3
              < 8
 √3 + √5 + √7 > √1 + √4 + √4
              > 1 + 2 + 2
              > 5

Integers:
2√3 + 3√5 + 4√7 = 20.755...
                > min√1
                > 2√1 + 3√4 + 4√4 
                > 2*1 + 3*2 + 4*2
                > 2 + 6 + 8
                > 16
                
                < max√1
                < 2√4 + 3√9 + 4√9
                < 2*2 + 3*3 + 4*3
                < 4 + 9 + 12
                < 25

2√3 + 3√5 + 4√7 > 16
                < 25

2√3 + 3√5 + 4√7
===================== 
169+12√15+16√21+24√35 > min√1
                      > 169 + 12√9 + 16√16 + 25√36
                      > 169 + 12*3 + 16*4 + 25*6
                      > 419
                      > 20.469...

                      < 169 + 12√16 + 16√25 + 24√36
                      < 169 + 12*4 + 16*5 + 24*6
                      < 441
                      < 21

2√3 + 3√5 + 4√7 = 20.755...
                > 2√3 + 3√3 + 4√3
                > (2+3+4)√3 = 9√3
                > 15.588
                < 2√3 + 3√12 + 4√12
                < 2√3 + 6√3 + 8√3
                < 16√3
                < 27.712...

                > 9√3
                < 16√3

2√3 + 3√5 + 4√7 cmp














*/