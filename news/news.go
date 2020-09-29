package news

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strings"
	"time"
)

const apiBaseUrl = "https://newsapi.org/v2/top-headlines"
const apiPageSizeParam = "100"

// Built automagically using json-to-go: https://mholt.github.io/json-to-go/
type Resp struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

var nr = &Resp{}

func FetchPeriodically() error {
	c, err := getConfig()
	if err != nil {
		return fmt.Errorf("news.FetchPeriodically()::getConfig() threw %v\n", err)
	}
	ticker := time.NewTicker(time.Duration(c.RefreshFrequency) * time.Minute)

	// Ensure that we always start with a non-empty Resp
	nr, err = fetchNewsFromAPI(c)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Print("fetching!")
				nr, _ = fetchNewsFromAPI(c)
			}
		}
	}()
	return nil
}

func GetNews() *Resp {
	return nr
}

func fetchNewsFromAPI(c *config) (*Resp, error) {
	nr := &Resp{}
	client := resty.New()

	_, err := client.R().
		SetQueryParamsFromValues(constructQueryParams(c)).
		SetHeader("Accept", "application/json").
		SetAuthToken(c.ApiKey).
		SetResult(nr).
		Get(apiBaseUrl)

	return nr, err
}

func constructQueryParams(c *config) url.Values {
	queryParams := make(url.Values)

	if c.Sources != nil {
		queryParams.Set("sources", strings.Join(c.Sources[:], ","))
	}
	if c.Country != "" {
		queryParams.Set("country", c.Country)
	}
	if c.Categories != nil {
		for _, category := range c.Categories {
			queryParams.Add("category", category)
		}
	}
	queryParams.Set("pageSize", apiPageSizeParam)
	return queryParams
}
