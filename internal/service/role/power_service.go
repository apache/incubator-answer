package role

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
)

// PowerRepo power repository
type PowerRepo interface {
	GetPowerList(ctx context.Context, power *entity.Power) (powers []*entity.Power, err error)
}

//// PowerService user service
//type PowerService struct {
//	powerRepo           PowerRepo
//	rolePowerRelService *RolePowerRelService
//}
//
//// NewPowerService new power service
//func NewPowerService(powerRepo PowerRepo, rolePowerRelService *RolePowerRelService) *PowerService {
//	return &PowerService{
//		powerRepo: powerRepo,
//	}
//}
//
//// GetRolePowerList get role power list
//func (ps *PowerService) GetRolePowerList(ctx context.Context, roleID string) (powers []string, err error) {
//	power := &entity.Power{}
//	powerList, err := ps.powerRepo.GetPowerList(ctx, power)
//	if err != nil {
//		return
//	}
//}
//
//// GetPowerList get  list all
//func (ps *PowerService) GetPowerList(ctx context.Context, req *schema.GetPowerListReq) (resp *[]schema.GetPowerResp, err error) {
//	power := &entity.Power{}
//	powers, err := ps.powerRepo.GetPowerList(ctx, power)
//	if err != nil {
//		return
//	}
//
//	resp = &[]schema.GetPowerResp{}
//	_ = copier.Copy(resp, powers)
//	return
//}
