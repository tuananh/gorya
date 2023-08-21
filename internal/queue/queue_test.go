package queue

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nduyphuong/gorya/internal/queue/options"
	"github.com/stretchr/testify/assert"
)

func TestClient_Enqueue(t *testing.T) {
	ctx := context.TODO()
	queue := newQueue(
		options.WithQueueAddr("localhost:6379"),
		options.WithFetchInterval(2*time.Second),
		options.WithQueueName("my-queue"))
	items := make([]any, 0)
	for i := 0; i < 10; i++ {
		items = append(items,
			fmt.Sprintf("val-%d", i),
		)
	}
	for _, v := range items {
		err := queue.Enqueue(ctx, v)
		assert.NoError(t, err)
	}
	resultChan := make(chan string)
	errChan := make(chan error)
	go queue.Dequeue(ctx, resultChan, errChan)
	for {
		select {
		case resp := <-resultChan:
			fmt.Printf("receive response from channel %v\n", resp)
		case err := <-errChan:
			fmt.Printf("receive error from channel %v\n", err)
			return
		}
	}
}
