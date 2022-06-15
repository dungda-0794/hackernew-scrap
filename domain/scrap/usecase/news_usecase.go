package usecase

import (
	"encoding/json"
	"fmt"
	"hackernew-scrap/core/errors"
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/models"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kelseyhightower/envconfig"
	"github.com/robfig/cron/v3"
)

type CronExpression struct {
	Minutes    string `envconfig:"CRON_MINUTES" default:"*"`
	Hours      string `envconfig:"CRON_HOURS" default:"*"`
	DayOfMonth string `envconfig:"CRON_DAY_OF_MONTH" default:"*"`
	Month      string `envconfig:"CRON_MONTH" default:"*"`
	DayOfWeek  string `envconfig:"CRON_DAY_OF_WEEK" default:"*"`
}

type newsUsecase struct {
	repository scrap.Repository
}

// NewNewsUsecase for create news usecase.
func NewNewsUsecase(repository scrap.Repository) scrap.NewsUsecase {
	return &newsUsecase{repository: repository}
}

func (n *newsUsecase) CreateNews(news models.News) (*models.News, error) {
	result, err := n.repository.Create(news)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

func NewCronJob() {
	cronExpression := CronExpression{}

	err := envconfig.Process("", &cronExpression)
	if err != nil {
		return
	}
	cronTime := fmt.Sprintf("%s %s %s %s %s", cronExpression.Minutes, cronExpression.Hours, cronExpression.DayOfMonth, cronExpression.Month, cronExpression.DayOfWeek)
	c := cron.New()
	c.AddFunc(cronTime, FetchData)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func FetchData() {
	articles := []models.News{}
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("table.itemlist tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr.athing", func(i int, el *colly.HTMLElement) {
			temp := models.News{}
			temp.Title = el.ChildText(".titlelink")
			temp.Link = el.ChildAttr(".titlelink", "href")
			temp.IDExternal = el.Attr("id")
			infoSelector := fmt.Sprintf("tr:nth-child(%d)", (i+1)*3-1)
			info := strings.Split(e.ChildText(infoSelector), " ")
			temp.Author = info[3]
			points, err := strconv.Atoi(info[0])
			if err != nil {
				temp.Point = 0
			} else {
				temp.Point = points
			}
			articles = append(articles, temp)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://news.ycombinator.com")

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Point > articles[j].Point
	})

	// Convert results to JSON data if the scraping job has finished
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		panic(err)
	}

	// Dump json to the standard output (can be redirected to a file)
	fmt.Println(string(jsonData))
}
