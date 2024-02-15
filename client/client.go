package main

import (
	"context"
	"echelon_testing/src"
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:50051"
	asyncFlag  = flag.Bool("async", false, "Запросить несколько видео")
)

// Вытаскивает ID видео из URL
func ExtractVideoID(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)
	if err != nil {
		return "", err
	}

	if u.Host == "www.youtube.com" || u.Host == "youtube.com" {
		queryParams := u.Query()
		return queryParams.Get("v"), nil
	} else if u.Host == "youtu.be" {
		// Для коротких URL
		return strings.TrimPrefix(u.Path, "/"), nil
	}

	return "", fmt.Errorf("неизвестный формат URL видео")
}

func getThumbnail(client src.ThumbnailerServiceClient, videoID string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := client.GetThumbnail(ctx, &src.GetThumbnailRequest{VideoUrl: videoID})

	if err != nil {
		log.Printf("Невозможно получить превью для видео с ID %s: %v", videoID, err)
		return
	}

	fmt.Printf("Превью для видео с ID %s: %s\n", videoID, r.GetThumbnailPath())
}

func main() {
	flag.Parse()
	videoURLs := flag.Args()

	if len(videoURLs) == 0 {
		fmt.Println("Используйте: [--async] <video_url1> <video_url2>...")
		return
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := src.NewThumbnailerServiceClient(conn)

	var wg sync.WaitGroup

	for _, videoURL := range videoURLs {
		videoID, err := ExtractVideoID(videoURL)
		if err != nil {
			log.Printf("Ошибка извлечения ID из URL %s: %v", videoURL, err)
			continue
		}

		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			getThumbnail(client, id, &wg)
		}(videoID)

		wg.Add(1)
		if !*asyncFlag {
			wg.Wait()
		}
	}
	if *asyncFlag {
		wg.Wait()
	}
}
