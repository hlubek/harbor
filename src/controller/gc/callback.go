package gc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor/src/lib/log"
	"github.com/goharbor/harbor/src/lib/orm"
	"github.com/goharbor/harbor/src/pkg/scheduler"
	"github.com/goharbor/harbor/src/pkg/task"
)

func init() {
	err := scheduler.RegisterCallbackFunc(SchedulerCallback, gcCallback)
	if err != nil {
		log.Fatalf("failed to registry GC call back, %v", err)
	}
}

func gcCallback(ctx context.Context, p string) error {
	param := &Policy{}
	if err := json.Unmarshal([]byte(p), param); err != nil {
		return fmt.Errorf("failed to unmarshal the param: %v", err)
	}
	_, err := Ctl.Start(orm.Context(), *param, task.ExecutionTriggerSchedule)
	return err
}
