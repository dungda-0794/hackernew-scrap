package usecase

import (
	"fmt"
	"hackernew-scrap/core/errors"
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/models"
	"sort"
	"strconv"
	"strings"

	"hackernew-scrap/infrastructure"

	"github.com/gocolly/colly"
)

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

func (n *newsUsecase) FetchData() {
	const url = "https://news.ycombinator.com"
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
		infrastructure.InfoLog.Println("Visiting", r.URL)
	})

	c.Visit(url)

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Point > articles[j].Point
	})

	for i := 0; i < len(articles); i++ {
		result, err := n.repository.CheckExists(articles[i].IDExternal)
		if err != nil {
			infrastructure.ErrLog.Fatal("Fail to check exists", err)
		}

		if result == true {
			continue
		}

		n.repository.Create(articles[i])
		break
	}

	return
}
