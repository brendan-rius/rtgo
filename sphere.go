package main

import (
	"math"
	"github.com/ungerik/go3d/vec3"
)

type Sphere struct {
	size float64
}

func (this Sphere) Intersect(ray Ray) bool {
	a := vec3.Dot(ray.Direction, ray.Direction)
	b := 2 * vec3.Dot(ray.Origin, ray.Direction)
	c := vec3.Dot(ray.Origin, ray.Origin) - float32(this.size * this.size)
	det := b * b - 4 * a * c
	if det < 0 {
		return false
	}
	sqrtDet := float32(math.Sqrt(float64(det)))
	inv2a := float32(1. / (2 * a))
	t1 := (-b + sqrtDet) * inv2a
	t2 := (-b - sqrtDet) * inv2a
	return t1 >= 0 || t2 >= 0
}