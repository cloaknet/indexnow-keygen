package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	// Valid characters for IndexNow keys
	// Requirements: https://yandex.ru/support/webmaster/indexnow/key.html
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

	// Key length constraints (8-128 characters allowed)
	minKeyLength = 64
	maxKeyLength = 96

	// Output directory for generated keys
	outputDir = "static"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	// Ensure output directory exists
	if err := ensureOutputDir(); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Get the number of keys to generate from user
	count, err := promptForKeyCount()
	if err != nil {
		return fmt.Errorf("failed to get key count: %w", err)
	}

	// Generate and save keys
	return generateAndSaveKeys(count)
}

func ensureOutputDir() error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return os.MkdirAll(outputDir, 0755)
	}
	return nil
}

func promptForKeyCount() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter the number of keys to generate: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(input)
		count, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if count <= 0 {
			fmt.Println("Number of keys must be positive.")
			continue
		}

		return count, nil
	}
}

func generateAndSaveKeys(count int) error {
	var errors []error

	for i := 0; i < count; i++ {
		key, err := generateSecureKey()
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to generate key %d: %w", i+1, err))
			continue
		}

		filename := key + ".txt"
		filePath := filepath.Join(outputDir, filename)

		// Check if file already exists
		if _, err := os.Stat(filePath); err == nil {
			errors = append(errors, fmt.Errorf("file already exists: %s", filePath))
			continue
		}

		if err := writeKeyToFile(filePath, key); err != nil {
			errors = append(errors, fmt.Errorf("failed to save key %d to file: %w", i+1, err))
			continue
		}

		fmt.Printf("%d) Key: %s\n   Saved to: %s\n", i+1, key, filePath)
	}

	if len(errors) > 0 {
		fmt.Printf("\nEncountered %d errors during generation:\n", len(errors))
		for _, err := range errors {
			fmt.Printf("- %v\n", err)
		}
		return fmt.Errorf("%d out of %d keys failed to generate", len(errors), count)
	}

	return nil
}

func generateSecureKey() (string, error) {
	// Generate random key length between min and max
	lengthRange := big.NewInt(int64(maxKeyLength - minKeyLength + 1))
	randomLength, err := rand.Int(rand.Reader, lengthRange)
	if err != nil {
		return "", fmt.Errorf("failed to generate random length: %w", err)
	}

	keyLength := minKeyLength + int(randomLength.Int64())

	// Generate random key
	key := make([]byte, keyLength)
	charactersLen := big.NewInt(int64(len(characters)))

	for i := 0; i < keyLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charactersLen)
		if err != nil {
			return "", fmt.Errorf("failed to generate random character: %w", err)
		}
		key[i] = characters[randomIndex.Int64()]
	}

	return string(key), nil
}

func writeKeyToFile(filename, key string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(key); err != nil {
		return fmt.Errorf("failed to write key to file: %w", err)
	}

	return nil
}