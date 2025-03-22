package service

import (
	"context"

	"video-balancer/api/proto"
	"video-balancer/internal/balancer"
)

type Service struct {
	proto.UnimplementedVideoBalancerServer
	balancer *balancer.Balancer
}

func NewService(balancer *balancer.Balancer) *Service {
	return &Service{
		balancer: balancer,
	}
}

func (s *Service) GetVideoURL(ctx context.Context, req *proto.VideoRequest) (*proto.VideoResponse, error) {
	redirectURL, err := s.balancer.GetRedirectURL(req.GetVideo())
	if err != nil {
		return nil, err
	}

	return &proto.VideoResponse{
		Url: redirectURL,
	}, nil
}
