package main

import (
	. "github.com/0daryo/ghostscad/primitive"
	"github.com/0daryo/ghostscad/sys"
	. "github.com/go-gl/mathgl/mgl64"
)

func newIntersection() Primitive {
	return NewIntersection(
		NewCylinder(4, 1),
		NewRotation(
			Vec3{90, 0, 0},
			NewCylinder(4, 0.9),
		))
}

func newUnion() Primitive {
	return NewTranslation(
		Vec3{-5, 0, 0},
		NewUnion(
			NewCylinder(4, 1),
			NewRotation(
				Vec3{90, 0, 0},
				NewCylinder(4, 0.9),
			)))
}

func newDifference() Primitive {
	return NewTranslation(
		Vec3{5, 0, 0},
		NewRender(
			10,
			NewDifference(
				NewCylinder(4, 1),
				NewRotation(
					Vec3{90, 0, 0},
					NewCylinder(4, 0.9),
				))))
}

func main() {
	sys.Initialize()
	sys.SetFn(360)
	sys.RenderOne(NewList(
		newIntersection(),
		newUnion(),
		newDifference(),
	))
}
