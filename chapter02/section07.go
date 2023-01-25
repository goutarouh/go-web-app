package main

import (
	"context"
	"fmt"
)

// context.Contextに値を含める方法
// Valueを使用する
// 値のアクセサ?のヘルパーメソッドをパッケージ内で定義する。

type TraceID string

// デフォルトの値
const ZeroTraceID = ""

// キーは独自型の空構造体をしようすることで
// 他パッケージとキーの衝突を避ける
type traceIDKey struct{}

func SetTraceId(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
}

func GetTraceId(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

// trace id = ""
// trace id = "test-id"

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceId(ctx))
	ctx = SetTraceId(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceId(ctx))
}
