package service

import (
	"context"
	"encoding/json"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"net"
	"net/http"
	"time"
)

var (
	connectionUri = "http://elastic:9200"
	index         = "student"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func MakeStudent() error {
	ctx := context.Background()
	esclient, err := getESClient()
	if err != nil {
		return err
	}

	var pow = []int64{1, 2, 4, 8, 16, 32, 64, 1, 4, 3, 5, 6}

	for _, age := range pow {
		newStudent := Student{
			Name:         String(10),
			Age:          age,
			AverageScore: Float64(),
		}
		dataJSON, err := json.Marshal(newStudent)
		if err != nil {
			return err
		}

		js := string(dataJSON)
		_, err = esclient.Index().
			Index(index).
			Id("student").
			BodyJson(js).
			Do(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func FetchFirstFiveStudent() (string, error) {
	ctx := context.Background()
	esclient, err := getESClient()
	if err != nil {
		return "", err
	}

	resp, err := esclient.Get().
		Index(index).
		Id("student").
		Do(ctx)

	if err != nil {
		return "", err
	}

	result, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func getESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(connectionUri),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHttpClient(&http.Client{Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			DisableKeepAlives:  true,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		}}),
	)

	fmt.Println("ES initialized...")

	return client, err
}
