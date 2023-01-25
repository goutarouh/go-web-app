package main

import (
	"context"
	"fmt"
	"time"
)

// キャンセル通知を知る

// 1 キャンセル済みか把握sるう
//   context.Context.Errメソッドで分かる

// 2 キャンセルされるまで待つ
//   Done()で得られるチャネルからの通知で判断する
//   業務の場合、別goroutineからの通知に使う

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() {
		fmt.Println("別goroutine")
	}()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("そして時は動き出す")
}
