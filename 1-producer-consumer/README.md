# Producer-Consumer Scenario

The producer reads in tweets from a mockstream and a consumer is processing the data to find out whether someone has tweeted about golang or not. The task is to modify the code inside `main.go` so that producer and consumer can run concurrently to increase the throughput of this program.  
生产者从一个模拟数据流中读取推文，消费者则处理这些数据以找出是否有人发了有关 Golang 的推文。任务是修改 main.go 中的代码，以便生产者和消费者可以并发运行，从而提高程序的吞吐量。
## Expected results:
Before: 
```
davecheney      tweets about golang
beertocode      does not tweet about golang
ironzeb         tweets about golang
beertocode      tweets about golang
vampirewalk666  tweets about golang
Process took 3.580866005s
```

After:
```
davecheney      tweets about golang
beertocode      does not tweet about golang
ironzeb         tweets about golang
beertocode      tweets about golang
vampirewalk666  tweets about golang
Process took 1.977756255s
```
