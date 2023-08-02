package activity_queue

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/log"
)

type ActivityQueueService interface {
	Send(ctx context.Context, msg *schema.ActivityMsg)
	RegisterHandler(handler func(ctx context.Context, msg *schema.ActivityMsg) error)
}

type activityQueueService struct {
	Queue   chan *schema.ActivityMsg
	Handler func(ctx context.Context, msg *schema.ActivityMsg) error
}

func (ns *activityQueueService) Send(ctx context.Context, msg *schema.ActivityMsg) {
	ns.Queue <- msg
}

func (ns *activityQueueService) RegisterHandler(
	handler func(ctx context.Context, msg *schema.ActivityMsg) error) {
	ns.Handler = handler
}

func (ns *activityQueueService) working() {
	go func() {
		for msg := range ns.Queue {
			log.Debugf("received activity %+v", msg)
			if ns.Handler == nil {
				log.Warnf("no handler for activity")
				continue
			}
			if err := ns.Handler(context.Background(), msg); err != nil {
				log.Error(err)
			}
		}
	}()
}

// NewActivityQueueService create a new activity queue service
func NewActivityQueueService() ActivityQueueService {
	ns := &activityQueueService{}
	ns.Queue = make(chan *schema.ActivityMsg, 128)
	ns.working()
	return ns
}
