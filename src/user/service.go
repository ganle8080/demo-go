package user

import (
	"context"

	"gorm.io/gorm"
)

type Service struct {
	ctx context.Context
	db  *gorm.DB
}

func NewService(ctx context.Context) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) PageList() {

}
