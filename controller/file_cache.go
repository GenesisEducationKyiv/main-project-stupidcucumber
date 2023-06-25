package controller

import (
	"api/bitcoin-api/models"
	"api/bitcoin-api/tools/creators"
	"api/bitcoin-api/tools/loaders"
	"encoding/json"
	"fmt"
	"os"
)

type FileCache struct {
	FileName string `json:"fileName"`
}

func NewFileCache() (*FileCache, error) {
	fileName, err := loaders.GetEnvVariable("CACHE_PATH")
	if err != nil {
		return nil, fmt.Errorf("instantiating FileCache %s: %w", fileName, err)
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if err = creators.CreateFile(fileName); err != nil {
			return nil, fmt.Errorf("instantiating FileCache: %w", err)
		}

		fileCache := FileCache{FileName: fileName}
		err := fileCache.Write(*models.NewCachedPrice())
		if err != nil {
			return nil, fmt.Errorf("instantiating FileCache: %w", err)
		}
	}

	return &FileCache{FileName: fileName}, nil
}

func (fileCache *FileCache) Write(cache models.CachedPrice) error {
	path, err := loaders.GetEnvVariable("CACHE_PATH")
	if err != nil {
		return fmt.Errorf("getting .env vavriable in writeCahce: %w", err)
	}

	file, err := os.OpenFile(fileCache.FileName, os.O_WRONLY|os.O_CREATE, 0o766)
	if err != nil {
		return fmt.Errorf("writing cache to %s: %w", fileCache.FileName, err)
	}

	cachedJSON, err := json.Marshal(cache)
	if err != nil {
		return fmt.Errorf("jsonifying CachedPrice: %w", err)
	}

	if _, err := file.Write(cachedJSON); err != nil {
		return fmt.Errorf("writing to the file %s: %w", path, err)
	}

	return nil
}

func (fileCache *FileCache) Read() (*models.CachedPrice, error) {
	var cache models.CachedPrice

	fileContent, err := os.ReadFile(fileCache.FileName)
	if err != nil {
		return nil, fmt.Errorf("reading cache from %s: %w", fileCache.FileName, err)
	}

	if err = json.Unmarshal(fileContent, &cache); err != nil {
		return nil, fmt.Errorf("unmarshalling JSON object into CachedPrice: %w", err)
	}

	return &cache, nil
}
