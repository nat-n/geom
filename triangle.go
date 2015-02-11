package geom

type Triangle [3]*Vec3

func (t *Triangle) Area() float64 {
	// Derive the vectors of two sides of the triangle
	v1 := t[1].Subtract(t[0])
	v2 := t[2].Subtract(t[0])
	// Calculate the cross product
	cp := v1.CrossProd(v2)
	// The answer is half the magnitude of the cross product
	return 0.5 * cp.Magnitude()
}

func (t *Triangle) IncludesVertex(v1 *Vec3) bool {
	for _, v2 := range t {
		if v1 == v2 {
			return true
		}
	}
	return false
}
