package math

type Math struct {
	a int
	b int
}

func NewMath(a, b int) Math {
	return Math{a, b}
}

func (m *Math) Add() int {
	return m.a + m.b
}
