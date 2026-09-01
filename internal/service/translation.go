package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/yazu-codes/scanme.git/internal/dto"
)

type TranslationService struct {
	translationServiceURL string
	httpClient            *http.Client
	logger                *slog.Logger
}

func NewTranslationService(translationServiceURL string, logger *slog.Logger) *TranslationService {
	return &TranslationService{
		translationServiceURL: translationServiceURL,
		httpClient:            &http.Client{},
		logger:                logger,
	}
}

type TranslateMenuRequest struct {
	Menu           dto.PublicMenu `json:"menu"`
	SourceLanguage string         `json:"source_language"`
	TargetLanguage string         `json:"target_language"`
}

func (s *TranslationService) TranslateMenu(menu dto.PublicMenu, targetLanguage string) (*dto.PublicMenu, error) {
	// Build request payload
	payload := TranslateMenuRequest{
		Menu:           menu,
		SourceLanguage: "en",
		TargetLanguage: targetLanguage,
	}

	// Marshal to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		s.logger.Error("Failed to marshal request", "error", err)
		return nil, err
	}

	// Make HTTP POST request to translation service
	url := fmt.Sprintf("%s/translate", s.translationServiceURL)
	resp, err := s.httpClient.Post(
		url,
		"application/json",
		bytes.NewReader(jsonPayload),
	)
	if err != nil {
		s.logger.Error("Failed to call translation service", "error", err, "url", url)
		return nil, err
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		s.logger.Error("Translation service error", "status", resp.StatusCode, "body", string(body))
		return nil, fmt.Errorf("translation service returned status %d", resp.StatusCode)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response", "error", err)
		return nil, err
	}

	// Parse response
	var translatedMenu dto.PublicMenu
	err = json.Unmarshal(respBody, &translatedMenu)
	if err != nil {
		s.logger.Error("Failed to parse response", "error", err)
		return nil, err
	}

	translatedMenu.MenuConfiguration.Theme = menu.MenuConfiguration.Theme // Preserve the originally fetched theme

	s.logger.Info("Menu translated successfully", "language", targetLanguage)
	return &translatedMenu, nil
}
