package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	stream "grpc-image-stream-test/proto"
	"io"
	"log"
	"sync"
	"time"
)

func testStream(client stream.StreamServiceClient) (int64, error) {
	startTime := time.Now().UnixNano()
	resp, err := client.GetImageStream(context.Background(), &stream.GetImageRequest{})
	if err != nil {
		return 0, err
	}

	var images []byte
	for {
		image, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			println(err.Error())
			return 0, err
		}
		images = append(images, image.Images...)
	}
	return time.Now().UnixNano() - startTime, nil
}

func testRequest(client stream.StreamServiceClient) (int64, error) {
	startTime := time.Now().UnixNano()
	_, err := client.GetImage(context.Background(), &stream.GetImageRequest{})
	if err != nil {
		return 0, err
	}
	return  time.Now().UnixNano() - startTime, nil
}

func Request(port int) {
	addr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		println(err.Error())
		return
	}

	client := stream.NewStreamServiceClient(conn)
	var (
		requestTotal int64
		streamTotal int64
	)

	var (
		num int64 = 100
		i int64 = 0
	)
	for i = 0;i < num;i++ {
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			requestTime, err := testRequest(client)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			requestTotal += requestTime
		}()
		go func() {
			defer wg.Done()
			streamTime, err := testStream(client)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			streamTotal += streamTime
		}()
		wg.Wait()
	}
	fmt.Printf("requestTotal: %d, requestAvg: %d\n", requestTotal, requestTotal / num)
	fmt.Printf("streamTotal: %d, streamAvg: %d", streamTotal, streamTotal / num)
}
