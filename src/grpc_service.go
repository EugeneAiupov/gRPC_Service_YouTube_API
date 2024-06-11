package src

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	// Кэш для хранения URL thumbnails, где ключом является ID видео
	thumbnailCache = make(map[string]string)
	// Mutex для синхронизации доступа к кэшу
	cacheMutex = sync.RWMutex{}
)

type Server struct {
	UnimplementedThumbnailerServiceServer
}

// Десериализация ответа YouTube API
type YouTubeApiResponse struct {
	Items []struct {
		Snippet struct {
			Thumbnails struct {
				Default struct {
					Url string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

// Отчистка кэша
func (s *Server) ClearCache() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	for k := range thumbnailCache {
		delete(thumbnailCache, k)
	}
}

// CheckCache проверяет наличие превью по videoID
func (s *Server) CheckCache(videoID string) (string, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	url, found := thumbnailCache[videoID]
	return url, found
}

// Возвращает videoID через YouTube API или ошибку в случае неудачи
func (s *Server) GetThumbnail(ctx context.Context, req *GetThumbnailRequest) (*GetThumbnailResponse, error) {
	videoID := req.GetVideoUrl()

	// Попытка извлечения URL из кэша
	cacheMutex.RLock()
	url, found := thumbnailCache[videoID]
	cacheMutex.RUnlock()
	if found {
		return &GetThumbnailResponse{ThumbnailPath: url}, nil
	}

	// Плохая реализация API key
	youtubeAPIKEY := "API key here"
	youtubeAPIURL := fmt.Sprintf("API URL", videoID, youtubeAPIKEY)

	resp, err := http.Get(youtubeAPIURL)
	if err != nil {
		return nil, fmt.Errorf("Ошибка запроса к YouTube API: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения ответа от YouTube API: %v", err)
	}

	// Десериализация JSON ответа
	var apiResponse YouTubeApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("Ошибка декодирования ответа от YouTube API (2): %v", err)
	}

	if len(apiResponse.Items) == 0 {
		return nil, fmt.Errorf("Таких видео нет по данному ID")
	}

	thumbnailUrl := apiResponse.Items[0].Snippet.Thumbnails.Default.Url

	// Добавление URL в кэш
	cacheMutex.Lock()
	thumbnailCache[videoID] = thumbnailUrl
	cacheMutex.Unlock()

	return &GetThumbnailResponse{ThumbnailPath: thumbnailUrl}, nil
}
