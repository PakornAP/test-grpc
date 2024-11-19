package model

import "test-grpc/grpc/internal/pb/routeguide"

type Point struct {
	Latitude  int32
	Longitude int32
}

func (p *Point) ToProto() *routeguide.Point {
	return &routeguide.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}

type Rectangle struct {
	Lo Point
	Hi Point
}

func (r *Rectangle) ToProto() *routeguide.Rectangle {
	return &routeguide.Rectangle{
		Lo: r.Lo.ToProto(),
		Hi: r.Hi.ToProto(),
	}
}

type Feature struct {
	Name     string
	Location Point
}

func (f *Feature) ToProto() *routeguide.Feature {
	return &routeguide.Feature{
		Name: f.Name,
		Location: &routeguide.Point{
			Latitude:  f.Location.Latitude,
			Longitude: f.Location.Longitude,
		},
	}
}
