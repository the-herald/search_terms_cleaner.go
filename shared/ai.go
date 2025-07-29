package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// FlaggedTerm represents a single term flagged by the AI
type FlaggedTerm struct {
	SearchTerm string `json:"search_term"`
	FlagType   string `json:"flag_type"` // "irrelevant", "competitor", or "none"
	Reason     string `json:"reason"`
}

// AIFlagTerms sends a list of terms to the OpenAI API and returns only flagged ones
func AIFlagTerms(terms []string) ([]FlaggedTerm, error) {
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		return nil, fmt.Errorf("missing OPENAI_API_KEY")
	}

	prompt := fmt.Sprintf(
		"You are an expert Google Ads analyst. For each search term, return a JSON array of objects with 'search_term', 'flag_type' ('irrelevant', 'competitor', or 'none'), and 'reason'. Only return flagged terms.\n\nSearch terms:\n%v",
		terms,
	)

	payload := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "system", "content": "You're a senior-level Google Ads analyst. Only return flagged terms."},
			{"role": "user", "content": prompt},
		},
		"temperature": 0.2,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+openaiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var parsed struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAI response: %w", err)
	}

	var flagged []FlaggedTerm
	if err := json.Unmarshal([]byte(parsed.Choices[0].Message.Content), &flagged); err != nil {
		return nil, fmt.Errorf("failed to parse flagged terms JSON: %w\nRaw content: %s", err, parsed.Choices[0].Message.Content)
	}
	return flagged, nil
}
