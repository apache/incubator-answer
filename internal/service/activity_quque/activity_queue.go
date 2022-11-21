package activity_quque

import (
	"github.com/answerdev/answer/internal/schema"
)

var (
	ActivityQueue = make(chan *schema.ActivityMsg, 128)
)

// AddActivity add new activity
func AddActivity(msg *schema.ActivityMsg) {
	ActivityQueue <- msg
}
