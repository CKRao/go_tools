package rx_go_lib

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"go.uber.org/zap"
)

// rxgo demo

func Run() {
	run4()
}

func run1() {
	observable := rxgo.Just(1, 2, 3, 4, 5, 6, 7)()
	c := observable.Observe()
	for item := range c {
		zap.S().Info(item.V)
	}
}

func run2() {
	observable := rxgo.Just(1, 2, 3, 4, errors.New("--"), 6, 7)()
	<-observable.ForEach(
		func(v interface{}) {
			fmt.Println("received: ", v)
		}, func(e error) {
			fmt.Println("err: ", e.Error())
		}, func() {
			fmt.Println("complete")
		})
}

func run3() {
	// 使用create创建observable
	observable := rxgo.Create([]rxgo.Producer{
		func(ctx context.Context, next chan<- rxgo.Item) {
			next <- rxgo.Of(1)
			next <- rxgo.Of(2)
			next <- rxgo.Of(3)
			next <- rxgo.Error(errors.New("unknown"))
			next <- rxgo.Of(5)
			next <- rxgo.Of(6)
			next <- rxgo.Of(7)
		},
	})

	// 分成两个 rxgo.Producer slice 也是一样的
	//observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
	//	next <- rxgo.Of(1)
	//	next <- rxgo.Of(2)
	//	next <- rxgo.Of(3)
	//	next <- rxgo.Error(errors.New("unknown"))
	//}, func(ctx context.Context, next chan<- rxgo.Item) {
	//	next <- rxgo.Of(4)
	//	next <- rxgo.Of(5)
	//}})

	<-observable.Map(func(ctx context.Context, v interface{}) (v2 interface{}, e error) {
		fmt.Printf("v: %+v\n", v)
		return v.(int) + 1, nil
	}, rxgo.WithErrorStrategy(rxgo.ContinueOnError)).ForEach(func(v interface{}) {
		fmt.Println("received: ", v)
	}, func(e error) {
		fmt.Println("received err: ", e.Error())
	}, func() {
		fmt.Println("complete")
	})
}

func run4() {
	// 使用create创建observable
	// 分成两个 rxgo.Producer slice 也是一样的
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}, func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(4)
		next <- rxgo.Of(5)
	}})

	<-observable.Filter(func(v interface{}) bool {
		return v.(int) != 4
	}, rxgo.WithErrorStrategy(rxgo.ContinueOnError)).ForEach(func(v interface{}) {
		fmt.Println("received: ", v)
	}, func(e error) {
		fmt.Println("received err: ", e.Error())
	}, func() {
		fmt.Println("complete")
	})
}
