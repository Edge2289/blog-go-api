package elastic

import (
	//"blog-go-api/app/config"
	"gopkg.in/olivere/elastic/v7"
	"context"
	"fmt"
)

type Tweet struct {
	User    string
	Message string
}

func TestAdd()  {

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://139.9.38.4:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("conn es succ")

	ctx := context.Background()
	for i := 0; i < 20; i++ {
		tweet := Tweet{User: "olivere", Message: "Take Five"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
			return
		}
	}

	fmt.Println("insert succ")
}
