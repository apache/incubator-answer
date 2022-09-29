package notification

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
)

// NotificationReadRepo notificationRead repository
type NotificationReadRepo interface {
	AddNotificationRead(ctx context.Context, notificationRead *entity.NotificationRead) (err error)
	RemoveNotificationRead(ctx context.Context, id int) (err error)
	UpdateNotificationRead(ctx context.Context, notificationRead *entity.NotificationRead) (err error)
	GetNotificationRead(ctx context.Context, id int) (notificationRead *entity.NotificationRead, exist bool, err error)
	GetNotificationReadList(ctx context.Context, notificationRead *entity.NotificationRead) (notificationReads []*entity.NotificationRead, err error)
	GetNotificationReadPage(ctx context.Context, page, pageSize int, notificationRead *entity.NotificationRead) (notificationReads []*entity.NotificationRead, total int64, err error)
}

// NotificationReadService user service
type NotificationReadService struct {
	notificationReadRepo NotificationReadRepo
}

func NewNotificationReadService(notificationReadRepo NotificationReadRepo) *NotificationReadService {
	return &NotificationReadService{
		notificationReadRepo: notificationReadRepo,
	}
}

// AddNotificationRead add notification read record
func (ns *NotificationReadService) AddNotificationRead(ctx context.Context, req *schema.AddNotificationReadReq) (err error) {
	notificationRead := &entity.NotificationRead{}
	_ = copier.Copy(notificationRead, req)
	return ns.notificationReadRepo.AddNotificationRead(ctx, notificationRead)
}

// RemoveNotificationRead delete notification read record
func (ns *NotificationReadService) RemoveNotificationRead(ctx context.Context, id int) (err error) {
	return ns.notificationReadRepo.RemoveNotificationRead(ctx, id)
}

// UpdateNotificationRead update notification read record
func (ns *NotificationReadService) UpdateNotificationRead(ctx context.Context, req *schema.UpdateNotificationReadReq) (err error) {
	notificationRead := &entity.NotificationRead{}
	_ = copier.Copy(notificationRead, req)
	return ns.notificationReadRepo.UpdateNotificationRead(ctx, notificationRead)
}

// GetNotificationRead get notification read record one
func (ns *NotificationReadService) GetNotificationRead(ctx context.Context, id int) (resp *schema.GetNotificationReadResp, err error) {
	notificationRead, exist, err := ns.notificationReadRepo.GetNotificationRead(ctx, id)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.UnknownError)
	}

	resp = &schema.GetNotificationReadResp{}
	_ = copier.Copy(resp, notificationRead)
	return resp, nil
}

// GetNotificationReadList get notification read record list all
func (ns *NotificationReadService) GetNotificationReadList(ctx context.Context, req *schema.GetNotificationReadListReq) (resp *[]schema.GetNotificationReadResp, err error) {
	notificationRead := &entity.NotificationRead{}
	_ = copier.Copy(notificationRead, req)

	notificationReads, err := ns.notificationReadRepo.GetNotificationReadList(ctx, notificationRead)
	if err != nil {
		return
	}

	resp = &[]schema.GetNotificationReadResp{}
	_ = copier.Copy(resp, notificationReads)
	return
}

// GetNotificationReadWithPage get notification read record list page
func (ns *NotificationReadService) GetNotificationReadWithPage(ctx context.Context, req *schema.GetNotificationReadWithPageReq) (pageModel *pager.PageModel, err error) {
	notificationRead := &entity.NotificationRead{}
	_ = copier.Copy(notificationRead, req)

	page := req.Page
	pageSize := req.PageSize

	notificationReads, total, err := ns.notificationReadRepo.GetNotificationReadPage(ctx, page, pageSize, notificationRead)
	if err != nil {
		return
	}

	resp := &[]schema.GetNotificationReadResp{}
	_ = copier.Copy(resp, notificationReads)

	return pager.NewPageModel(total, resp), nil
}
