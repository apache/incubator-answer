package notice_queue

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/log"
)

type NotificationQueueService interface {
	Send(ctx context.Context, msg *schema.NotificationMsg)
	RegisterHandler(handler func(ctx context.Context, msg *schema.NotificationMsg) error)
}

type notificationQueueService struct {
	Queue   chan *schema.NotificationMsg
	Handler func(ctx context.Context, msg *schema.NotificationMsg) error
}

func (ns *notificationQueueService) Send(ctx context.Context, msg *schema.NotificationMsg) {
	ns.Queue <- msg
}

func (ns *notificationQueueService) RegisterHandler(
	handler func(ctx context.Context, msg *schema.NotificationMsg) error) {
	ns.Handler = handler
}

func (ns *notificationQueueService) working() {
	go func() {
		for msg := range ns.Queue {
			log.Debugf("received notification %+v", msg)
			if ns.Handler == nil {
				log.Warnf("no handler for notification")
				continue
			}
			if err := ns.Handler(context.Background(), msg); err != nil {
				log.Error(err)
			}
		}
	}()
}

// NewNotificationQueueService create a new notification queue service
func NewNotificationQueueService() NotificationQueueService {
	ns := &notificationQueueService{}
	ns.Queue = make(chan *schema.NotificationMsg, 128)
	ns.working()
	return ns
}

func AddNotification2(msg *schema.NotificationMsg) {
}
