package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func callBoth(ctx context.Context, errVal, slowURL, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()
	wg.Wait()
	fmt.Println("done with both")
}

var client = http.Client{Timeout: 10 * time.Second}

func callServer(ctx context.Context, label, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err: ", err)
		return err
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "response err: ", err)
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(label, "read err: ", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result: ", result)
	}
	if result == "error" {
		fmt.Println("cancelling from ", label)
		return errors.New("error happened")
	}
	return nil
}
