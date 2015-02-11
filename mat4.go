package geom

type SymMat4 [10]float64

func (m1 *SymMat4) Add(m2 *SymMat4) {
	m1[0] += m2[0]
	m1[1] += m2[1]
	m1[2] += m2[2]
	m1[3] += m2[3]
	m1[4] += m2[4]
	m1[5] += m2[5]
	m1[6] += m2[6]
	m1[7] += m2[7]
	m1[8] += m2[8]
	m1[9] += m2[9]
}

func (m1 *SymMat4) Subtract(m2 *SymMat4) {
	m1[0] -= m2[0]
	m1[1] -= m2[1]
	m1[2] -= m2[2]
	m1[3] -= m2[3]
	m1[4] -= m2[4]
	m1[5] -= m2[5]
	m1[6] -= m2[6]
	m1[7] -= m2[7]
	m1[8] -= m2[8]
	m1[9] -= m2[9]
}

func (m *SymMat4) IsEmpty() bool {
	return m[0] == 0 &&
		m[1] == 0 &&
		m[2] == 0 &&
		m[3] == 0 &&
		m[4] == 0 &&
		m[5] == 0 &&
		m[6] == 0 &&
		m[7] == 0 &&
		m[8] == 0 &&
		m[9] == 0
}

func (m *SymMat4) Inverse() (result_ref *SymMat4, can_invert bool) {
	result := SymMat4{}
	result_ref = &result
	can_invert = true
	det := m.Determinant()
	if det == 0 {
		can_invert = false
		return
	}
	result[0] = (m[5]*m[8]*m[6] - m[6]*m[7]*m[6] + m[6]*m[5]*m[8] - m[4]*m[8]*m[8] - m[5]*m[5]*m[9] + m[4]*m[7]*m[9]) / det
	result[1] = (m[3]*m[7]*m[6] - m[2]*m[8]*m[6] - m[3]*m[5]*m[8] + m[1]*m[8]*m[8] + m[2]*m[5]*m[9] - m[1]*m[7]*m[9]) / det
	result[2] = (m[2]*m[6]*m[6] - m[3]*m[5]*m[6] + m[3]*m[4]*m[8] - m[1]*m[6]*m[8] - m[2]*m[4]*m[9] + m[1]*m[5]*m[9]) / det
	result[3] = (m[3]*m[5]*m[5] - m[2]*m[6]*m[5] - m[3]*m[4]*m[7] + m[1]*m[6]*m[7] + m[2]*m[4]*m[8] - m[1]*m[5]*m[8]) / det
	result[4] = (m[2]*m[8]*m[3] - m[3]*m[7]*m[3] + m[3]*m[2]*m[8] - m[0]*m[8]*m[8] - m[2]*m[2]*m[9] + m[0]*m[7]*m[9]) / det
	result[5] = (m[3]*m[5]*m[3] - m[2]*m[6]*m[3] - m[3]*m[1]*m[8] + m[0]*m[6]*m[8] + m[2]*m[1]*m[9] - m[0]*m[5]*m[9]) / det
	result[6] = (m[2]*m[6]*m[2] - m[3]*m[5]*m[2] + m[3]*m[1]*m[7] - m[0]*m[6]*m[7] - m[2]*m[1]*m[8] + m[0]*m[5]*m[8]) / det
	result[7] = (m[1]*m[6]*m[3] - m[3]*m[4]*m[3] + m[3]*m[1]*m[6] - m[0]*m[6]*m[6] - m[1]*m[1]*m[9] + m[0]*m[4]*m[9]) / det
	result[8] = (m[3]*m[4]*m[2] - m[1]*m[6]*m[2] - m[3]*m[1]*m[5] + m[0]*m[6]*m[5] + m[1]*m[1]*m[8] - m[0]*m[4]*m[8]) / det
	result[9] = (m[1]*m[5]*m[2] - m[2]*m[4]*m[2] + m[2]*m[1]*m[5] - m[0]*m[5]*m[5] - m[1]*m[1]*m[7] + m[0]*m[4]*m[7]) / det
	return
}

func (m *SymMat4) Determinant() float64 {
	return m[3]*m[5]*m[5]*m[3] - m[2]*m[6]*m[5]*m[3] - m[3]*m[4]*m[7]*m[3] + m[1]*m[6]*m[7]*m[3] +
		m[2]*m[4]*m[8]*m[3] - m[1]*m[5]*m[8]*m[3] - m[3]*m[5]*m[2]*m[6] + m[2]*m[6]*m[2]*m[6] +
		m[3]*m[1]*m[7]*m[6] - m[0]*m[6]*m[7]*m[6] - m[2]*m[1]*m[8]*m[6] + m[0]*m[5]*m[8]*m[6] +
		m[3]*m[4]*m[2]*m[8] - m[1]*m[6]*m[2]*m[8] - m[3]*m[1]*m[5]*m[8] + m[0]*m[6]*m[5]*m[8] +
		m[1]*m[1]*m[8]*m[8] - m[0]*m[4]*m[8]*m[8] - m[2]*m[4]*m[2]*m[9] + m[1]*m[5]*m[2]*m[9] +
		m[2]*m[1]*m[5]*m[9] - m[0]*m[5]*m[5]*m[9] - m[1]*m[1]*m[7]*m[9] + m[0]*m[4]*m[7]*m[9]
}

func (m *SymMat4) VertexError(v Vec3) float64 {
	// v(transpose) * m * v
	x := v.X
	y := v.Y
	z := v.Z
	return x*x*m[0] + 2*x*y*m[1] + 2*x*z*m[2] + 2*x*m[3] +
		y*y*m[4] + 2*y*z*m[5] + 2*y*m[6] +
		z*z*m[7] + 2*z*m[8] +
		m[9]
}

func (m *SymMat4) Clone() *SymMat4 {
	new_m := SymMat4{}
	copy(new_m[:], m[:])
	return &new_m
}
