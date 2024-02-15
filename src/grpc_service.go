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

func (s *Server) CheckCache(videoID string) (string, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	url, found := thumbnailCache[videoID]
	return url, found
}

func (s *Server) GetThumbnail(ctx context.Context, req *GetThumbnailRequest) (*GetThumbnailResponse, error) {
	videoID := req.GetVideoUrl()

	// Попытка извлечения URL из кэша
	cacheMutex.RLock()
	url, found := thumbnailCache[videoID]
	cacheMutex.RUnlock()
	if found {
		return &GetThumbnailResponse{ThumbnailPath: url}, nil
	}

	apiKey := "AIzaSyCuRb8FAAbzDWZ1nl0tE5iSeONivMuREYM"
	apiUrl := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=snippet", videoID, apiKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to YouTube API: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error decode response from YouTube API: %v", err)
	}

	// Десериализация JSON ответа
	var apiResponse YouTubeApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error decode response from YouTube API (2): %v", err)
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
