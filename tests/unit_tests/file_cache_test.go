package unitTest

import (
	"api/bitcoin-api/models"
	"encoding/json"
	"os"
	"testing"
	"time"
)

const (
	testFileName = "test.cache"
)

func cleanupTests() {
	if _, err := os.Stat(testFileName); err == nil {
		os.Remove(testFileName)
	}
}

func TestFileCacheWrite(t *testing.T) {
	database := models.FileCache{
		FileName: testFileName,
	}
	expectedTestData := models.CachedPrice{
		Price:     1000000,
		TimeStamp: time.Date(2000, time.January, 1, 1, 0, 0, 0, time.UTC),
	}
	var actualTestData models.CachedPrice

	err := database.Write(expectedTestData)
	if err != nil {
		t.Errorf("unexpected error encountered during writing to the file: %v", err)
	}

	body, err := os.ReadFile(testFileName)
	if err != nil {
		t.Errorf("unexpected error encountered during reading file: %v", err)
	}

	err = json.Unmarshal(body, &actualTestData)
	if err != nil {
		t.Errorf("data written to the file is corrupted: %v", err)
	}

	if actualTestData != expectedTestData {
		t.Errorf("actual data (%v) written to the file is different from what is expected (%v)", actualTestData, expectedTestData)
	}

	t.Cleanup(cleanupTests)
}

func TestFileCacheRead(t *testing.T) {
	database := models.FileCache{
		FileName: testFileName,
	}
	expectedTestData := models.CachedPrice{
		Price:     1000000,
		TimeStamp: time.Date(2000, time.January, 1, 1, 0, 0, 0, time.UTC),
	}

	if err := database.Write(expectedTestData); err != nil {
		t.Errorf("unexpected error during writing data to file: %v", err)
	}

	actualTestData, err := database.Read()
	if err != nil {
		t.Errorf("error during reading file: %v", err)
	}

	if *actualTestData != expectedTestData {
		t.Errorf("actual data (%#v) is different from expected (%#v)", *actualTestData, expectedTestData)
	}

	t.Cleanup(cleanupTests)
}
