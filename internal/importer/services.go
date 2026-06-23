package importer

import (
	"context"
	"monitoring-edc/internal/monitoring"
)

type Service interface {
	ImportFile(
		ctx context.Context,
		path string,
	) error
}

type service struct {
	monitoringService monitoring.Service
}

func (s *service) ImportFile(
	ctx context.Context,
	path string,
) error {

	data, err := ParseMonitoring(path)
	if err != nil {
		return err
	}

	return s.monitoringService.Import(
		ctx,
		data,
	)
}

func NewService(
	monitoringService monitoring.Service,
) Service {
	return &service{
		monitoringService: monitoringService,
	}
}
