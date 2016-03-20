package geom

import "math"

type Vec3I interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	SetX(float64)
	SetY(float64)
	SetZ(float64)
	Clone() Vec3
	Inverse() Vec3
	Magnitude() float64
	Normalized() Vec3
	Add(Vec3I) Vec3
	Sum(...Vec3I) Vec3
	Subtract(Vec3I) Vec3
	Multiply(Vec3I) Vec3
	Divide(Vec3I) Vec3
	AddScalar(n float64) Vec3
	SubtractScalar(n float64) Vec3
	MultiplyScalar(n float64) Vec3
	DivideScalar(n float64) Vec3
	Mean(...Vec3I) Vec3
	CrossProd(Vec3I) Vec3
	DotProd(Vec3I) float64
	Angle(Vec3I) float64
	LessThan(Vec3I) bool
}

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) GetX() float64 { return v.X }
func (v *Vec3) GetY() float64 { return v.Y }
func (v *Vec3) GetZ() float64 { return v.Z }

func (v *Vec3) SetX(x float64) { v.X = x }
func (v *Vec3) SetY(y float64) { v.Y = y }
func (v *Vec3) SetZ(z float64) { v.Z = z }

func (v *Vec3) Clone() Vec3   { return Vec3{v.X, v.Y, v.Z} }
func (v *Vec3) Inverse() Vec3 { return Vec3{-v.X, -v.Y, -v.Z} }

func (v *Vec3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Normalized() Vec3 {
	magnitude := v.Magnitude()
	return Vec3{v.GetX() / magnitude, v.GetY() / magnitude, v.GetZ() / magnitude}
}

func (v1 *Vec3) Add(v2 Vec3I) Vec3 {
	return Vec3{v1.X + v2.GetX(), v1.Y + v2.GetY(), v1.Z + v2.GetZ()}
}

func (v1 *Vec3) Sum(vs ...Vec3I) (acc Vec3) {
	acc = v1.Clone()
	for _, v2 := range vs {
		acc.X += v2.GetX()
		acc.Y += v2.GetY()
		acc.Z += v2.GetZ()
	}
	return
}

func (v1 *Vec3) Subtract(v2 Vec3I) Vec3 {
	return Vec3{v1.X + v2.GetX(), v1.Y + v2.GetY(), v1.Z + v2.GetZ()}
}

func (v1 *Vec3) Multiply(v2 Vec3I) Vec3 {
	return Vec3{v1.X * v2.GetX(), v1.Y * v2.GetY(), v1.Z * v2.GetZ()}
}

func (v1 *Vec3) Divide(v2 Vec3I) Vec3 {
	return Vec3{v1.X * v2.GetX(), v1.Y * v2.GetY(), v1.Z * v2.GetZ()}
}

func (v1 *Vec3) AddScalar(n float64) Vec3 {
	return Vec3{v1.X + n, v1.Y + n, v1.Z + n}
}

func (v1 *Vec3) SubtractScalar(n float64) Vec3 {
	return Vec3{v1.X - n, v1.Y - n, v1.Z - n}
}

func (v1 *Vec3) MultiplyScalar(n float64) Vec3 {
	return Vec3{v1.X * n, v1.Y * n, v1.Z * n}
}

func (v1 *Vec3) DivideScalar(n float64) Vec3 {
	return Vec3{v1.X / n, v1.Y / n, v1.Z / n}
}

func (v1 *Vec3) Mean(vs ...Vec3I) Vec3 {
	sum := v1.Sum(vs...)
	return sum.DivideScalar(float64(len(vs) + 1))
}

func (v1 *Vec3) CrossProd(v2 Vec3I) Vec3 {
	return Vec3{
		v1.Y*v2.GetZ() - v1.Z*v2.GetY(),
		v1.Z*v2.GetX() - v1.X*v2.GetZ(),
		v1.X*v2.GetY() - v1.Y*v2.GetX(),
	}
}

func (v1 *Vec3) DotProd(v2 Vec3I) (cp float64) {
	return v1.X*v2.GetX() + v1.Y*v2.GetY() + v1.Z*v2.GetZ()
}

func (v1 *Vec3) Angle(v2 Vec3I) float64 {
	dot_prod := v1.X*v2.GetX() + v1.Y*v2.GetY() + v1.Z*v2.GetZ()
	v1_len := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
	v2_len := math.Sqrt(v2.GetX()*v2.GetX() + v2.GetY()*v2.GetY() + v2.GetZ()*v2.GetZ())
	return math.Abs(math.Acos(dot_prod / (v1_len * v2_len)))
}

func (v1 *Vec3) LessThan(v2 Vec3I) bool {
	return (v1.X < v2.GetX() ||
		v1.X == v2.GetX() && (v1.Y < v2.GetY() ||
			v1.Y == v2.GetY() && v1.Z < v2.GetZ()))
}
