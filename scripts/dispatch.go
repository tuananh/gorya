//go:build ignore
// +build ignore

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/nduyphuong/gorya/internal/queue"
	"github.com/nduyphuong/gorya/internal/worker"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	dispatch := flag.Bool("dispatch", true, "run dispatcher")
	process := flag.Bool("process", true, "run processor")
	queueName := flag.String("queueName", "", "queue name")
	flag.Parse()
	if *queueName == "" {
		fmt.Println("nothing to do")
		return
	}
	ctx := context.TODO()
	w := worker.NewClient(worker.Options{
		Log: logrus.New(),
		QueueOpts: queue.Options{
			Addr:          "localhost:6379",
			Name:          *queueName,
			FetchInterval: 2 * time.Second,
		},
	})
	numWorkers := 10
	var input []worker.QueueElem
	for i := 0; i < 100; i++ {
		input = append(input, worker.QueueElem{
			RequestURI: "/tasks/change_state",
			Project:    "test-aws-account",
			TagKey:     "phuong",
			TagValue:   "test",
			Action:     0,
		})
	}
	if *dispatch {
		for _, v := range input {
			if err := w.Dispatch(ctx, &v); err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
		return
	}
	if *process {
		resultChan := make(chan string)
		errCh := make(chan error)
		for i := 0; i < numWorkers; i++ {
			go w.Process(ctx, resultChan, errCh)
		}
		for res := range resultChan {
			fmt.Printf("receive result %v from worker \n", res)
		}
		var resErr error
		for err := range errCh {
			resErr = errors.Join(resErr, err)
		}
		fmt.Printf("errs: %v", resErr)
		return
	}

}
