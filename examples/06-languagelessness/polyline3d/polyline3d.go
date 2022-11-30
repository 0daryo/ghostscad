package main

import (
	"github.com/0daryo/ghostscad/lib/shapes"
	"github.com/0daryo/ghostscad/sys"

	. "github.com/0daryo/ghostscad/lib/degmath"
	. "github.com/0daryo/ghostscad/primitive"
	. "github.com/go-gl/mathgl/mgl64"
)

func main() {
	sys.Initialize()

	r := 50.0
	points := []Vec3{}
	for a := 0.0; a <= 180; a++ {
		points = append(points, Vec3{
			r * Cos(-90.0+a) * Cos(a),
			r * Cos(-90.0+a) * Sin(a),
			r * Sin(-90.0+a),
		})
	}

	list := NewList()

	for i := 0; i < 8; i++ {
		list.Add(
			NewRotation(Vec3{0, 0, float64(i * 45)}, shapes.NewPolyline3d(points, 2).Build()),
		)
	}
	sys.RenderOne(list)
}
