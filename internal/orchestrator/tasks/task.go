package tasks

import (
	"context"
	"time"

	v1 "github.com/garethgeorge/backrest/gen/go/v1"
	"github.com/garethgeorge/backrest/internal/hook"
	"github.com/garethgeorge/backrest/internal/oplog"
	"github.com/garethgeorge/backrest/internal/orchestrator/repo"
	"google.golang.org/protobuf/proto"
)

var NeverScheduledTask = ScheduledTask{}

const (
	PlanForUnassociatedOperations = "_unassociated_"

	TaskPriorityStats          = -1
	TaskPriorityDefault        = 0
	TaskPriorityInteractive    = 1 << 1
	TaskPriorityIndexSnapshots = 1 << 2
	TaskPriorityForget         = 1 << 3
	TaskPriorityPrune          = 1 << 4
)

// TaskRunner is an interface for running tasks. It is used by tasks to create operations and write logs.
type TaskRunner interface {
	// CreateOperation creates the operation in storage and sets the operation ID in the task.
	CreateOperation(*v1.Operation) error
	// UpdateOperation updates the operation in storage. It must be called after CreateOperation.
	UpdateOperation(*v1.Operation) error
	// ExecuteHooks
	ExecuteHooks(events []v1.Hook_Condition, vars hook.HookVars) error
	// OpLog returns the oplog for the operations.
	OpLog() *oplog.OpLog
	// GetRepo returns the repo with the given ID.
	GetRepo(repoID string) (*v1.Repo, error)
	// GetPlan returns the plan with the given ID.
	GetPlan(planID string) (*v1.Plan, error)
	// GetRepoOrchestrator returns the orchestrator for the repo with the given ID.
	GetRepoOrchestrator(repoID string) (*repo.RepoOrchestrator, error)
	// ScheduleTask schedules a task to run at a specific time.
	ScheduleTask(task Task, priority int) error
}

// ScheduledTask is a task that is scheduled to run at a specific time.
type ScheduledTask struct {
	Task        Task          // the task to run
	RunAt       time.Time     // the time at which the task should be run.
	Op          *v1.Operation // operation associated with this execution of the task.
	configModno int32         // used by to detect if the config has changed since this task was scheduled.
	callbacks   []func(error) // callbacks to run after the task is run.
}

func (s ScheduledTask) Eq(other ScheduledTask) bool {
	return s.Task == other.Task && s.RunAt.Equal(other.RunAt)
}

func (s ScheduledTask) Less(other ScheduledTask) bool {
	if s.RunAt.Equal(other.RunAt) {
		return s.Task.Name() < other.Task.Name()
	}
	return s.RunAt.Before(other.RunAt)
}

func (s ScheduledTask) cancel(oplog *oplog.OpLog) error {
	if s.Op == nil {
		return nil
	}

	opCopy := proto.Clone(s.Op).(*v1.Operation)
	opCopy.Status = v1.OperationStatus_STATUS_SYSTEM_CANCELLED
	opCopy.DisplayMessage = "operation cancelled"
	opCopy.UnixTimeEndMs = curTimeMillis()
	if err := oplog.Update(opCopy); err != nil {
		return err
	}

	return nil
}

// Task is a task that can be scheduled to run at a specific time.
type Task interface {
	Name() string                                                       // human readable name for this task.
	Next(now time.Time, runner TaskRunner) ScheduledTask                // returns the next scheduled task.
	Run(ctx context.Context, st ScheduledTask, runner TaskRunner) error // run the task.
	PlanID() string                                                     // the ID of the plan this task is associated with.
	RepoID() string                                                     // the ID of the repo this task is associated with.
}

type BaseTask struct {
	TaskName   string
	TaskPlanID string
	TaskRepoID string
}

func (b BaseTask) Name() string {
	return b.TaskName
}

func (b BaseTask) PlanID() string {
	return b.TaskPlanID
}

func (b BaseTask) RepoID() string {
	return b.TaskRepoID
}

type OneoffTask struct {
	BaseTask
	RunAt       time.Time
	FlowID      int64 // the ID of the flow this task is associated with.
	DidSchedule bool
	ProtoOp     *v1.Operation // the prototype operation for this class of task.
}

func (o *OneoffTask) Next(now time.Time, runner TaskRunner) ScheduledTask {
	if o.DidSchedule {
		return NeverScheduledTask
	}
	o.DidSchedule = true

	var op *v1.Operation
	if o.ProtoOp != nil {
		op = proto.Clone(o.ProtoOp).(*v1.Operation)
		op.PlanId = o.PlanID()
		op.RepoId = o.RepoID()
		op.FlowId = o.FlowID
		op.UnixTimeStartMs = timeToUnixMillis(o.RunAt) // TODO: this should be updated before Run is called.
		op.Status = v1.OperationStatus_STATUS_PENDING
	}

	return ScheduledTask{
		RunAt: o.RunAt,
		Op:    op,
	}
}

type GenericOneoffTask struct {
	BaseTask
	OneoffTask
	Do func(ctx context.Context, st ScheduledTask, runner TaskRunner) error
}

func (g *GenericOneoffTask) Run(ctx context.Context, st ScheduledTask, runner TaskRunner) error {
	return g.Do(ctx, st, runner)
}

func timeToUnixMillis(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond()/1000000)
}

func curTimeMillis() int64 {
	return timeToUnixMillis(time.Now())
}
