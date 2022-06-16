package schedule

import(
	"fmt"
    "os"
    "os/signal"

	"hackernew-scrap/infrastructure"
	"hackernew-scrap/domain/scrap"

	"github.com/robfig/cron/v3"
	"github.com/kelseyhightower/envconfig"
)

type CronExpression struct {
	Minutes string `envconfig:"CRON_MINUTES" default:"*"`
	Hours       string `envconfig:"CRON_HOURS" default:"*"`
	DayOfMonth       string `envconfig:"CRON_DAY_OF_MONTH" default:"*"`
	Month       string `envconfig:"CRON_MONTH" default:"*"`
	DayOfWeek   string `envconfig:"CRON_DAY_OF_WEEK" default:"*"`
}

type newsSchedule struct {
	usecase scrap.NewsUsecase
}

func NewNewsSchedule(us scrap.NewsUsecase) scrap.NewsSchedule{
	return &newsSchedule{usecase: us}
}


func (schedule *newsSchedule) CronJob() {
	cronExpression := CronExpression{}

	err := envconfig.Process("", &cronExpression)
	if err != nil {
		infrastructure.ErrLog.Fatal("Fail to get cron expression ", err)
	}
	cronTime := fmt.Sprintf("%s %s %s %s %s", cronExpression.Minutes, cronExpression.Hours, cronExpression.DayOfMonth, cronExpression.Month, cronExpression.DayOfWeek)
	c := cron.New()
	c.AddFunc(cronTime, schedule.usecase.FetchData)
	go c.Start()
	sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    <-sig
}
