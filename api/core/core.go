package core

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type PaginatedResponse[T any] struct {
	Count    int
	Next     string
	Previous string
	Results  []T
}

const SwapiUrl string = "https://swapi.dev/api"

var RedisClient *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "redis_cache:6379",
	Password: "",
	DB:       0,
})

func FetchRelationships[T any, U any](
	urls []string,
	urlScheme string,
	redisClient *redis.Client,
	context context.Context,
	create func(source U) T,
) ([]T, error) {
	var err error = nil
	relationships := make([]T, len(urls))

	for i := range urls {
		url := urls[i]

		if !strings.HasPrefix(url, urlScheme) {
			err = errors.New("invalid url scheme provided : " + url)
			break
		}

		rawContent, _ := GetOrFetch(url, context, *redisClient, 0)

		var content U
		parsingError := json.Unmarshal([]byte(rawContent), &content)

		if parsingError != nil {
			err = parsingError
			break
		}

		relationships[i] = create(content)
	}

	return relationships, err
}

func GetOrFetch(
	url string,
	context context.Context,
	client redis.Client,
	expiration time.Duration,
) (string, error) {
	content, _ := client.Get(context, url).Result()

	if content != "" {
		return content, nil
	}

	apiResponse, err := http.Get(url)
	if err != nil || apiResponse.StatusCode != http.StatusOK {
		return "", err
	}

	defer apiResponse.Body.Close()
	bodyContent, parsingError := io.ReadAll(apiResponse.Body)

	if parsingError != nil {
		return "", parsingError
	}

	content = string(bodyContent[:])
	client.Set(context, url, content, expiration) // 0 means no expiration

	return content, nil
}