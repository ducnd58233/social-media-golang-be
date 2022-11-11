package subscriber

import (
	"context"
	component "social-media-be/components"
	"social-media-be/components/asyncjob"
	"social-media-be/pubsub"
)

type consumerJob struct {
	Title string
	Hdl   func(ctx context.Context, message *pubsub.Message) error // handler
}

type consumerEngine struct {
	appCtx component.AppContext
}

func NewEngine(appCtx component.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

// Put all subscriber here
func (engine *consumerEngine) Start() error {
	return nil
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	logger := engine.appCtx.GetLogger("start sub topic")

	for _, item := range consumerJobs {
		logger.Info("Setup consumer for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			logger.Info("running job for", job.Title, ". Value: ", message.Data())
			return job.Hdl(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, logger, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				logger.Error(err)
			}
		}
	}()

	return nil
}
