package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    go parentProcess(ctx)
    time.Sleep(1000 * time.Second)
}

func parentProcess(ctx context.Context) {
    childCtx, cancel := context.WithCancel(ctx)
    childCtx2, _ := context.WithCancel(childCtx)
    go childProcess(childCtx, "child1")
    go childProcess(childCtx2, "child2")
    //10秒まってキャンセル。
    time.Sleep(9 * time.Second)
    // childCtxから派生したchildCtx2もキャンセルが伝播される。実際にはchildCtx2のcancelFuncもどこかで呼ばないとgo vet error
    cancel()
    //１０秒後にキャンセル サンプル用にcancelFuncを捨ててるけど実際にはcancelFuncを明示的にどこかで呼ばないとgo vet error
    ctxWithTimeout10, _ := context.WithTimeout(ctx, time.Second*10)
    go childProcess(ctxWithTimeout10, "with timeout")
}

func childProcess(ctx context.Context, prefix string) {
    for i := 1; i <= 1000; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("%s: canceld \n", prefix)
            return
        case <-time.After(1 * time.Second):
            fmt.Printf("%s:%d sec..\n", prefix, i)
        }
      fmt.Println("This time is ", i, " time(s).")
    }
}
