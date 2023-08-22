package notice_queue

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/log"
)

type ExternalNotificationQueueService interface {
	Send(ctx context.Context, msg *schema.ExternalNotificationMsg)
	RegisterHandler(handler func(ctx context.Context, msg *schema.ExternalNotificationMsg) error)
}

type externalNotificationQueueService struct {
	Queue   chan *schema.ExternalNotificationMsg
	Handler func(ctx context.Context, msg *schema.ExternalNotificationMsg) error
}

func (ns *externalNotificationQueueService) Send(ctx context.Context, msg *schema.ExternalNotificationMsg) {
	ns.Queue <- msg
}

func (ns *externalNotificationQueueService) RegisterHandler(
	handler func(ctx context.Context, msg *schema.ExternalNotificationMsg) error) {
	ns.Handler = handler
}

func (ns *externalNotificationQueueService) working() {
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

// NewNewQuestionNotificationQueueService create a new notification queue service
func NewNewQuestionNotificationQueueService() ExternalNotificationQueueService {
	ns := &externalNotificationQueueService{}
	ns.Queue = make(chan *schema.ExternalNotificationMsg, 128)
	ns.working()
	return ns
}
