package tests

import (
	"context"
	"echelon_testing/src"
	"testing"
)

// Тест на проверку получения thumbnail без кэша
func TestGetThumbnail(t *testing.T) {
	s := src.Server{}

	videoID := "SItyvIMLac8"
	req := &src.GetThumbnailRequest{VideoUrl: videoID}

	resp, err := s.GetThumbnail(context.Background(), req)
	if err != nil {
		t.Fatalf("GetThumbNail() error = %v", err)
	}

	if resp.GetThumbnailPath() == "" {
		t.Error("Ожидался URL превью, получена пустая строка")
	}
}

// Тест на првоерку работы кэша
func TestGetThumbnailCache(t *testing.T) {
	s := src.Server{}

	videoID := "QS01z1mkWYM"
	req := &src.GetThumbnailRequest{VideoUrl: videoID}

	s.ClearCache()

	if _, found := s.CheckCache(videoID); found {
		t.Fatal("Кэш должен быть пустым")
	}
	_, err := s.GetThumbnail(context.Background(), req)
	if err != nil {
		t.Fatalf("First GetThumbnail() call failed: %v", err)
	}

	if _, found := s.CheckCache(videoID); !found {
		t.Fatal("Ожидалось найти ответ в кэше после первого запроса")
	}
	_, err = s.GetThumbnail(context.Background(), req)
	if err != nil {
		t.Fatalf("Second GetThumbnail() call failed: %v", err)
	}
}
