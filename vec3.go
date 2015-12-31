package geom

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3) Clone() *Vec3 {
	return &Vec3{v.X, v.Y, v.Z}
}

func (v *Vec3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v1 *Vec3) Add(v2 Vec3) *Vec3 {
	return &Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 *Vec3) Subtract(v2 *Vec3) *Vec3 {
	return &Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 *Vec3) Multiply(n float64) *Vec3 {
	return &Vec3{v1.X * n, v1.Y * n, v1.Z * n}
}

func (v1 *Vec3) Devide(n float64) *Vec3 {
	return &Vec3{v1.X / n, v1.Y / n, v1.Z / n}
}

func (v1 *Vec3) Mean(vs ...*Vec3) *Vec3 {
	acc := Vec3{0, 0, 0}
	for _, v := range vs {
		acc.X += v.X
		acc.Y += v.Y
		acc.Z += v.Z
	}
	return acc.Devide(float64(len(vs)))
}

func (v1 *Vec3) CrossProd(v2 *Vec3) *Vec3 {
	return &Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 *Vec3) DotProd(v2 *Vec3) (cp float64) {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Compute the normal of the face which includes by the three given points
func TriNormal(a, b, c *Vec3) *Vec3 {
	// Derive the vectors of two sides of the triangle
	v1 := Vec3{b.X - a.X, b.Y - a.Y, b.Z - a.Z}
	v2 := Vec3{c.X - a.X, c.Y - a.Y, c.Z - a.Z}

	// Calculate the cross product
	cpx := v1.Y*v2.Z - v1.Z*v2.Y
	cpy := v1.Z*v2.X - v1.X*v2.Z
	cpz := v1.X*v2.Y - v1.Y*v2.X

	// Normalize the cross product to arrive at the normal vector
	l := math.Sqrt(cpx*cpx + cpy*cpy + cpz*cpz)
	return &Vec3{cpx / l,
		cpy / l,
		cpz / l,
	}
}

func (v1 *Vec3) Angle(v2 *Vec3) float64 {
	dot_prod := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
	v1_len := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
	v2_len := math.Sqrt(v2.X*v2.X + v2.Y*v2.Y + v2.Z*v2.Z)
	return math.Abs(math.Acos(dot_prod / (v1_len * v2_len)))
}

func (v1 *Vec3) LessThan(v2 *Vec3) bool {
	return (v1.X < v2.X ||
		v1.X == v2.X && (v1.Y < v2.Y ||
			v1.Y == v2.Y && v1.Z < v2.Z))
}
