package initializers

import (
	"bufio"
	"crypto/sha1"
	"os"
	"strconv"
	"strings"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/willf/bloom"
)

var BloomFilter *bloom.BloomFilter

func InitializeBloomFilter() {
	// ADJUST FILTER EXPECTED COUNT AND FALSE POSITIVE RATE
	filterCount := os.Getenv("BLOOM_FILTER_COUNT")
	filterCountInt, err := strconv.Atoi(filterCount)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Failed to initialize bloom filter")
	}
	filterFalsePositiveRate := os.Getenv("BLOOM_FILTER_FALSE_POSITIVE_RATE")
	filterFalsePositiveRateFloat, err := strconv.ParseFloat(filterFalsePositiveRate, 64)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Failed to initialize bloom filter")
	}
	BloomFilter = bloom.NewWithEstimates(uint(filterCountInt), filterFalsePositiveRateFloat)

	// LOAD HASHES
	filePath := os.Getenv("COMMON_PASS_LIB_PATH")
	file, err := os.Open(filePath)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Failed to initialize bloom filter")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		hash := sha1.Sum([]byte(password))
		BloomFilter.Add(hash[:])
	}

	if err := scanner.Err(); err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Failed to initialize bloom filter")
	}

	helpers.LogWithSeverity(constants.LOGGER_SEVERITY_INFO, "[INFO] Bloom filter initialized")
}
