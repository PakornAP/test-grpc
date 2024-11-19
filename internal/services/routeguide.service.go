package services

import (
	"fmt"
	"io"
	"log"
	"math"
	"sync"
	"test-grpc/grpc/internal/model"
	"test-grpc/grpc/internal/pb/routeguide"
	"test-grpc/grpc/internal/port"
)

type RouteGuideService struct {
	savedFeatures []model.Feature
	mu            sync.Mutex
	routeNotes    map[string][]*routeguide.RouteNote
}

func NewRouteGuideService(features []model.Feature) port.RouteGuideServicePort {
	return &RouteGuideService{
		savedFeatures: features,
		routeNotes:    map[string][]*routeguide.RouteNote{},
	}
}

func inRange(point *routeguide.Point, rect *routeguide.Rectangle) bool {
	top := math.Max(float64(rect.Hi.Longitude), float64(rect.Lo.Longitude))
	bottom := math.Min(float64(rect.Hi.Longitude), float64(rect.Lo.Longitude))
	right := math.Max(float64(rect.Hi.Latitude), float64(rect.Lo.Latitude))
	left := math.Min(float64(rect.Hi.Latitude), float64(rect.Lo.Latitude))
	x, y := float64(point.Latitude), float64(point.Longitude)
	return x <= right && x >= left && y <= top && y >= bottom

}

func serialize(point *routeguide.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}

func (rs *RouteGuideService) GetFeature(point model.Point) (model.Feature, error) {
	for _, ft := range rs.savedFeatures {
		if ft.Location == point {
			return ft, nil
		}
	}
	return model.Feature{Name: "not found", Location: point}, nil
}

func (rs *RouteGuideService) ListFeatures(rect *routeguide.Rectangle, stream routeguide.RouteGuide_ListFeaturesServer) error {
	var res []*routeguide.Feature
	for _, f := range rs.savedFeatures {
		feature := f.ToProto()
		if inRange(feature.Location, rect) {
			res = append(res, feature)
			if err := stream.Send(feature); err != nil {
				log.Println("error !! : ", err)
				return err
			}
		}
	}
	return nil
}

// func (rs *RouteGuideService) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
// 	...
// }
// ...

func (rs *RouteGuideService) RouteChat(stream routeguide.RouteGuide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println(rs.routeNotes)
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)
		rs.mu.Lock()
		rs.routeNotes[key] = append(rs.routeNotes[key], in)
		rn := make([]*routeguide.RouteNote, len(rs.routeNotes[key]))
		copy(rn, rs.routeNotes[key])
		rs.mu.Unlock()
		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}

}
