package geom

import "math"

type Triangle [3]Vec3I

func (t *Triangle) Area() float64 {
	// Derive the vectors of two sides of the triangle
	a, b, c := t[0], t[1], t[2]
	v1 := b.Subtract(a)
	v2 := c.Subtract(a)
	// The answer is half the magnitude of the cross product
	cp := v1.CrossProd(&v2)
	return 0.5 * cp.Magnitude()
}

func (t *Triangle) IncludesVertex(v1 Vec3I) bool {
	for _, v2 := range t {
		if v1 == v2 {
			return true
		}
	}
	return false
}

// Compute the normal of the plane of the triangle using the right hand rule
func (t *Triangle) Normal() Vec3 {
	// Derive the vectors of two sides of the triangle
	a, b, c := t[0], t[1], t[2]
	v1 := Vec3{b.GetX() - a.GetX(), b.GetY() - a.GetY(), b.GetZ() - a.GetZ()}
	v2 := Vec3{c.GetX() - a.GetX(), c.GetY() - a.GetY(), c.GetZ() - a.GetZ()}

	// Calculate the cross product
	cp := v1.CrossProd(&v2)

	// Normalize the cross product to arrive at the normal vector
	l := math.Sqrt(cp.GetX()*cp.GetX() + cp.GetY()*cp.GetY() + cp.GetZ()*cp.GetZ())
	return cp.DivideScalar(l)
}
