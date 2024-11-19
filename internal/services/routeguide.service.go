package services

import (
	"log"
	"math"
	"test-grpc/grpc/internal/model"
	"test-grpc/grpc/internal/pb/routeguide"
	"test-grpc/grpc/internal/port"
)

type RouteGuideService struct {
	savedFeatures []model.Feature
}

func NewRouteGuideService(features []model.Feature) port.RouteGuideServicePort {
	return &RouteGuideService{
		savedFeatures: features,
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
	log.Println(res)
	return nil
}

// func (rs *RouteGuideService) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
// 	...
// }
// ...

// func (rs *RouteGuideService) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
// 	...
// }
