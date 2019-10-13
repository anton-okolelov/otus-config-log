package service

import (
	"github.com/anton-okolelov/otus-config-log/config"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	config config.Config
}

func NewService(logger *zap.Logger, config config.Config) *Service {
	return &Service{logger: logger, config: config}
}

func (s *Service) Start() {
	s.logger.Sugar().Infof("Start %v, %v", s.config.Host, s.config.Port)
}
