package scheduler

import (
	"context"

	"github.com/robfig/cron/v3"
)

func NewJobScheduler(ctx context.Context) {
	c := cron.New()
	priceJob := NewPriceJob(ctx)
	c.AddFunc("0 */10 10-16 * * *", priceJob.Run)
	c.Start()
}
