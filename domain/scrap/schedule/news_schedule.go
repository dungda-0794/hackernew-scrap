package schedule

import (
	"fmt"
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/infrastructure"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/robfig/cron/v3"
)

type cronExpression struct {
	Minutes    string `envconfig:"CRON_MINUTES" default:"0"`
	Hours      string `envconfig:"CRON_HOURS" default:"8"`
	DayOfMonth string `envconfig:"CRON_DAY_OF_MONTH" default:"*"`
	Month      string `envconfig:"CRON_MONTH" default:"*"`
	DayOfWeek  string `envconfig:"CRON_DAY_OF_WEEK" default:"*"`
}

type newsSchedule struct {
	usecase scrap.NewsUsecase
}

// NewNewsSchedule for create news schedule
func NewNewsSchedule(us scrap.NewsUsecase) scrap.NewsSchedule {
	return &newsSchedule{usecase: us}
}

func (schedule *newsSchedule) CronJob() {
	cronExpression := cronExpression{}
	err := envconfig.Process("CRON", &cronExpression)
	if err != nil {
		infrastructure.ErrLog.Fatal("Fail to get cron expression ", err)
	}
	cronTime := fmt.Sprintf("%s %s %s %s %s",
		cronExpression.Minutes,
		cronExpression.Hours,
		cronExpression.DayOfMonth,
		cronExpression.Month,
		cronExpression.DayOfWeek)
	c := cron.New()
	_, err = c.AddFunc(cronTime, schedule.usecase.FetchData)
	if err != nil {
		infrastructure.ErrLog.Fatal("Fail to add func ", err)
	}
	go c.Start()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
