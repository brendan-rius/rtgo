package main

type Primitive interface {
	Intersect(Ray) bool
}