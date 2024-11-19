package port

import (
	"test-grpc/grpc/internal/model"
	"test-grpc/grpc/internal/pb/routeguide"
)

type RouteGuideServicePort interface {
	GetFeature(point model.Point) (model.Feature, error)
	ListFeatures(rectangle *routeguide.Rectangle, stream routeguide.RouteGuide_ListFeaturesServer) error
	// RecordRoute()
	// RouteChat()
}
