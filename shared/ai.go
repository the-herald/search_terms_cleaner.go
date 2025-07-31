package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type FlaggedTerm struct {
	SearchTerm string `json:"search_term"`
	FlagType   string `json:"flag_type"` // "irrelevant", "competitor", or "none"
	Reason     string `json:"reason"`
}

const batchSize = 30
const maxRetries = 3

func AIFlagTerms(terms []string) ([]FlaggedTerm, error) {
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-4"
	}
	temp := 0.2
	if t := os.Getenv("OPENAI_TEMPERATURE"); t != "" {
		fmt.Sscanf(t, "%f", &temp)
	}

	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return nil, fmt.Errorf("missing OPENAI_API_KEY")
	}

	var allFlags []FlaggedTerm
	for i := 0; i < len(terms); i += batchSize {
		end := i + batchSize
		if end > len(terms) {
			end = len(terms)
		}
		batch := terms[i:end]

		flags, err := callOpenAIFlagTerms(batch, key, model, temp)
		if err != nil {
			return nil, err
		}
		allFlags = append(allFlags, flags...)
	}

	log.Printf("AI flagged %d of %d terms", len(allFlags), len(terms))
	return allFlags, nil
}

func callOpenAIFlagTerms(batch []string, key, model string, temp float64) ([]FlaggedTerm, error) {
	examples := `[{"search_term":"how to become a dentist","flag_type":"irrelevant","reason":"educational intent"},{"search_term":"smile direct club reviews","flag_type":"competitor","reason":"competing dental brand"},{"search_term":"affordable dental schools","flag_type":"irrelevant","reason":"looking for education, not services"}]`

	prompt := fmt.Sprintf(
		"You are an expert Google Ads analyst. For each search term, return a JSON array of objects with 'search_term', 'flag_type' ('irrelevant', 'competitor', or 'none'), and 'reason'. Only return flagged terms.\n\nExamples:\n%s\n\nSearch terms:\n%v",
		examples, batch,
	)

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": "You're a senior Google Ads analyst. Only return flagged terms as a valid JSON array."},
			{"role": "user", "content": prompt},
		},
		"temperature": temp,
	}

	body, _ := json.Marshal(payload)

	var resp *http.Response
	var err error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+key)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{Timeout: 20 * time.Second}
		resp, err = client.Do(req)

		if err == nil && resp.StatusCode == 200 {
			break
		}
		if resp != nil {
			if resp.StatusCode == 429 || resp.StatusCode == 502 {
				log.Printf("Retrying due to API error (%d)...", resp.StatusCode)
				time.Sleep(2 * time.Second)
				continue
			}
		}
		break
	}
	if err != nil || resp == nil {
		return nil, fmt.Errorf("OpenAI request failed: %v", err)
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
	if err := json.Unmarshal(raw, &parsed); err != nil || len(parsed.Choices) == 0 {
		return nil, fmt.Errorf("failed to parse OpenAI response (empty or malformed): %v", err)
	}

	var flagged []FlaggedTerm
	if err := json.Unmarshal([]byte(parsed.Choices[0].Message.Content), &flagged); err != nil {
		return nil, fmt.Errorf("failed to unmarshal flagged terms JSON: %w\nContent: %s", err, parsed.Choices[0].Message.Content)
	}
	return flagged, nil
}

func ExtractTroubleWord(term, reason string) string {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return ""
	}
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-4"
	}

	prompt := fmt.Sprintf(`You're a Google Ads intent expert. For the search term and reason below, extract the single most problematic word.

Search Term: %q
Reason: %q

Only return that one word (no punctuation).`, term, reason)

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are an AI that identifies the most problematic word in a flagged search term. Respond with just one word."},
			{"role": "user", "content": prompt},
		},
		"temperature": 0.0,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ""
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
	if err := json.Unmarshal(raw, &parsed); err != nil || len(parsed.Choices) == 0 {
		return ""
	}

	word := strings.TrimSpace(parsed.Choices[0].Message.Content)
	word = strings.Trim(word, `"'.,`)
	return word
}
