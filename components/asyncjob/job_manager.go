package asyncjob

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

type group struct {
	isConcurrent bool
	jobs         []Job
	wg           *sync.WaitGroup
	logger       *logrus.Entry
}

func NewGroup(isConcurrent bool, logger *logrus.Logger, jobs ...Job) *group {
	g := &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
		logger:       logger.WithField("service_name", "cron job"),
	}

	return g
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs)) // buffer

	for i, _ := range g.jobs {
		if g.isConcurrent {
			// use goroutine to run concurrent jobs
			go func(aj Job) {
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])

			continue
		}

		job := g.jobs[i]
		errChan <- g.runJob(ctx, job)
		g.wg.Done()
	}

	var err error

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}

	g.wg.Wait()
	return err
}

func (g *group) runJob(ctx context.Context, j Job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			g.logger.Error(err)
			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}

	return nil
}
