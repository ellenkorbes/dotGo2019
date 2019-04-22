package main

import (
	s "github.com/deadsy/sdfx/sdf"
	"math"
)

func shape(vector, mesh string) {
	output2d, output3d := generate_2D_3D()
	s.RenderSVG(output2d, 50, vector, "fill:none;stroke:#02f2b4;stroke-width:3px;")
	s.RenderSTL(output3d, 300, mesh)
}

// #######                   ###
// #        #    #  #    #   ###
// #        #    #  ##   #   ###
// #####    #    #  # #  #    #
// #        #    #  #  # #
// #        #    #  #   ##   ###
// #         ####   #    #   ###

func generate_2D_3D() (s.SDF2, s.SDF3) {
	var output2d s.SDF2
	var output3d s.SDF3

	// 1. Start here
	circle := s.Circle2D(20)
	coin := s.Extrude3D(circle, 10)
	output2d, output3d = circle, coin

	// 2. Then uncomment this
	// square := s.Box2D(s.V2{10, 80}, 0)
	// cube := s.Extrude3D(square, 30)
	// cube = s.Transform3D(cube, s.Translate3d(s.V3{90, 0, 95}))
	// output2d = s.Union2D(circle, square)
	// output3d = s.Union3D(coin, cube)

	// 3. Then this
	// coin = s.Transform3D(coin, s.RotateY(s.DtoR(90)))
	// coin = s.Transform3D(coin, s.Translate3d(s.V3{30, 70, 30}))
	// coin2 := s.Extrude3D(circle, 10)
	// coin2 = s.Transform3D(coin2, s.RotateY(s.DtoR(90)))
	// coin2 = s.Transform3D(coin2, s.Translate3d(s.V3{30, -70, 30}))
	// output3d = s.Union3D(cube, coin, coin2)

	// 4. Now let's try some different stuff
	// b := s.NewBezier()
	// b.Add(0, 0).HandleFwd(s.DtoR(0), 150)
	// b.Add(50, 150).Mid()
	// b.Add(0, 300).HandleRev(s.DtoR(0), 150)
	// b.Close()
	// b2d := s.Polygon2D(b.Polygon().Vertices())
	// output2d = b2d

	// 5. And in 3D
	// b3d := s.Revolve3D(b2d)
	// output3d = b3d

	// 6. And voilá, we have a gopher!
	// cube = s.Transform3D(s.Extrude3D(s.Box2D(s.V2{20, 20}, 0), 30), s.Translate3d(s.V3{90, 0, 95}))
	// q := bezierBlobs(25, 25, 50, 37, 65, 32, 30)
	// w := bezierBlobs(25, 25, 50, 37, 65, -32, 30)
	// e := bezierBlobs(13, 26, 26, 30, 90, 0, 70)
	// r := bezierBlobs(10, 10, 20, 10, 86, -40, 37)
	// t := bezierBlobs(10, 10, 20, 10, 86, 40, 43)
	// y := bezierBlobs(5, 5, 13, 13, 105, 0, 70)
	// gopher := s.Union3D(cube, coin, coin2, b3d, q, w, e, r, t, y)
	// output3d = gopher

	// 7. Now let's make a dilator
	// output2d, output3d = dilator()

	// 8. And one more thing!
	// gopher = s.ScaleUniform3D(gopher, 0.1925)
	// gopher = s.Transform3D(gopher, s.Translate3d(s.V3{0, 0, -235}))
	// output3d = s.Union3D(gopher, output3d)

	return output2d, output3d
}

// #######                   ###
// #        #    #  #    #   ###
// #        #    #  ##   #   ###
// #####    #    #  # #  #    #
// #        #    #  #  # #
// #        #    #  #   ##   ###
// #         ####   #    #   ###

func bezierBlobs(mX, mY, fY, h, tX, tY, tZ float64) s.SDF3 {
	b := s.NewBezier()
	b.Add(0, 0).HandleFwd(s.DtoR(0), h)
	b.Add(mX, mY).Mid()
	b.Add(0, fY).HandleRev(s.DtoR(0), h)
	b.Close()
	output2d := s.Polygon2D(b.Polygon().Vertices())
	output3d := s.Revolve3D(output2d)
	output3d = s.Transform3D(output3d, s.Translate3d(s.V3{tX, tY, tZ}))
	return output3d
}

func dilator() (s.SDF2, s.SDF3) {
	d := s.NewBezier()
	length := 200.0
	radius := 34.0 / 2
	steps := 50.0
	step := radius / steps
	for x := 1.0; x < radius; x += step {

		pt1 := (x * x) * 0.3
		pt2 := (x / math.Abs(radius-x)) * 7
		y := (pt1 + pt2) * 0.2

		if y > length {
			y = length
		}
		d.Add(x, -y)
	}
	d.Add(0, -length)
	d.Close()
	output2d := s.Polygon2D(d.Polygon().Vertices())
	output3d := s.Revolve3D(output2d)
	return output2d, output3d
}
