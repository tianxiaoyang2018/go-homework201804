package service

import "github.com/p1cn/tantan-backend-common/health"

type IService interface {
	Start() error
	Stop() error
	GetHealthChecks() []health.HealthCheck
}
