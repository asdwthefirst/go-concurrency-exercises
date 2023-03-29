//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweets chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweets)
			return
		}

		//tweets = append(tweets, tweet)
		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	return
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	//tweets := producer(stream)
	tweets := make(chan *Tweet) //有无缓冲在这里使用上区别不大
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		producer(stream, tweets)
		wg.Done()
	}()
	// Consumer
	go func() {
		consumer(tweets)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
