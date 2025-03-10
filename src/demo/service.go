package demo

import (
	"context"

	"gorm.io/gorm"
)

type DemoService struct {
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) where(ctx context.Context, data *Demo) *gorm.DB {
	db := *&gorm.DB{}

	return &db
}
