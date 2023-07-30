package alg

type H struct {
	gh *R32
	ef *R32
	cd *R32
	ab *B
}

type Hs struct {
	*R32s
}

func NewHs(rs *R32s) *Hs {
	return &Hs{
		R32s: rs,
	}
}

func (s *Hs) NewH(b, c, d, e, f, g, h Z, a N) *H {
	gh := s.Radical(g, h, nil)
	if gh == nil {
		return nil // overflow
	}
	ef := s.Radical(e, f, nil) // TODO
	if ef == nil {
		return nil
	}
	ab := NewBnotReduce(b, a)
	return &H {
		gh: gh,
		ab: ab,
	}


	/*
	(&a).Reduce4(&b, &c, &e)
	if ab == nil {
		return nil // infinite
	}
	if c == 0 || d == 0 { // degenerated D is B
		return &D{
			ab: ab,
			cd: NewR32zero(),
		}
	}
	if cd := s.NewR32(c, d); cd == nil {
		return nil // overflow
	} else {
		// after the d simplification, c was increased
		// specially when b is 0, we need to try reduce a and c
		ab.Reduce3(cd.out)
		return &D{
			ab: ab,
			cd: cd,
		}
	}*/
}

func (x *H) String() string {
	str := &Str{}
	x.gh.Str(str)
	return str.String()
}
