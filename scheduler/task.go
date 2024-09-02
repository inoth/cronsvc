package scheduler

import "context"

type TaskEvent interface {
	Execute(ctx context.Context, collector <-chan Accumulator)
}

type Accumulator interface {
}
