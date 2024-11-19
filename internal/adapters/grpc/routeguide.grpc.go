package adapters

import (
	"context"
	"log"
	"test-grpc/grpc/internal/model"
	"test-grpc/grpc/internal/pb/routeguide"
	"test-grpc/grpc/internal/port"
)

type RouteGuideGrpcAdapter struct {
	routeguide.UnimplementedRouteGuideServer
	Service port.RouteGuideServicePort
}

func NewRouteGuideGrpcAdapter(service port.RouteGuideServicePort) *RouteGuideGrpcAdapter {
	return &RouteGuideGrpcAdapter{
		Service: service,
	}
}

func (ra *RouteGuideGrpcAdapter) GetFeature(_ context.Context, req *routeguide.Point) (*routeguide.Feature, error) {
	point := model.Point{
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
	}
	feature, err := ra.Service.GetFeature(point)
	if err != nil {
		log.Fatal(err)
	}

	return feature.ToProto(), nil
}

func (ra *RouteGuideGrpcAdapter) ListFeatures(rect *routeguide.Rectangle, stream routeguide.RouteGuide_ListFeaturesServer) error {
	if err := ra.Service.ListFeatures(rect, stream); err != nil {
		return err
	}
	return nil
}
