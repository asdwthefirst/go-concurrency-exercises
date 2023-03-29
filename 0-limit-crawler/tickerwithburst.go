package main

import (
	"context"
	"time"
)

const rateLimit = time.Second / 10 // 10 calls per second

// Client is an interface that calls something with a payload.
//type Client interface {
//	Call(*Payload)
//}

//// Payload is some payload a Client would send in a call.
//type Payload struct{}
//
//// BurstRateLimitCall allows burst rate limiting client calls with the
//// payloads.
//func BurstRateLimitCall(ctx context.Context, client Client, payloads []*Payload, burstLimit int) {
//	throttle := make(chan time.Time, burstLimit)
//
//	ctx, cancel := context.WithCancel(ctx)//context的用法
//	defer cancel()
//
//	go func() {
//		ticker := time.NewTicker(rateLimit)
//		defer ticker.Stop()
//		for t := range ticker.C {
//			select {
//			case throttle <- t:
//			case <-ctx.Done():
//				return // exit goroutine when surrounding function returns
//			}
//		}
//	}()
//
//	for _, payload := range payloads {
//		<-throttle // rate limit our client calls
//		go client.Call(payload)
//	}
//}

type Client interface {
	Call(*Payload)
}

type Payload struct{}

func BurstRateLimitCall(ctx context.Context, client Client, payloads []*Payload, burstLimit int) {
	throttle := make(chan time.Time, burstLimit)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() //管理子任务

	go func() {
		ticker := time.NewTicker(rateLimit)
		defer ticker.Stop()
		for { //这样写不对
			select {
			case throttle <- <-ticker.C:
				//这样写是把ticker.c的值取出来放到一个临时变量再放到throttle，
			// 这样操作如果throttle的写入比较慢的话，会导致数据被丢弃，因为下一次循环已经开始了
			case <-ctx.Done():
				return
			}
		}
		for t := range ticker.C {
			select {
			case throttle <- t: //说是这样会等待throttle被处理完毕才执行下一次循环
			case <-ctx.Done(): //对于这里能不能被执行我还是存疑？
				return
			}
		}
		var t time.Time
		for { //这是我修改过的。
			select {
			case t = <-ticker.C:
				throttle <- t
			case <-ctx.Done():
				return
			}
		}
	}()

	for _, payload := range payloads {
		<-throttle // rate limit our client calls
		go client.Call(payload)
	}

}
