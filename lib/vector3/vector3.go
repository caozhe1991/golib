package vector3

import (
	"fmt"
	"math"
)

// 表示1除以10的8次方
const EPSILON = 1e-8
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// 重载string方法
func(p *Vector3) String() string {
	return fmt.Sprintf("Vector3(%f, %f, %f)", p.X, p.Y, p.Z)
}

// 向量 + 向量
func(p *Vector3) Add(a *Vector3) *Vector3 {
	return &Vector3{
		X: p.X + a.X,
		Y: p.Y + a.Y,
		Z: p.Z + a.Z,
	}
}

// 向量 - 向量
func (p *Vector3) Subtract(a *Vector3) *Vector3 {
	return &Vector3{
		X: p.X - a.X,
		Y: p.Y - a.Y,
		Z: p.Z - a.Z,
	}
}

// 向量 * 标量
func(p *Vector3) Multiply(v float64) *Vector3 {
	return &Vector3{
		X: p.X * v,
		Y: p.Y * v,
		Z: p.Z * v,
	}
}

// 求模，计算向量的长度
func(p *Vector3) Norm() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

// 向量归一化，是指向量的长度为1，而不是x、y、z都为1
func(p *Vector3)Normalize() *Vector3 {
	norm := p.Norm()
	// 因为被除数不能为零，所以先先判断。
	// float 类型由于精度原因，不能用 == 判断是否为零，所以通常认为小于某个足够小的值时既为0
	if norm < EPSILON {
		return Zero()
	}
	return p.Multiply(1 / norm)
}



// 零向量
func Zero() *Vector3 {
	return &Vector3{X:0, Y:0, Z:0}
}

func New(x, y, z float64) *Vector3 {
	return &Vector3{X: x, Y:y, Z:z}
}

