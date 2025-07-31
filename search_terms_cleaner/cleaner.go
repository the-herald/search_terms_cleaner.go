package search_terms_cleaner

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"google.golang.org/genproto/googleapis/ads/googleads/v20/resources"
	"google.golang.org/genproto/googleapis/ads/googleads/v20/services"
)

type FlaggedTerm struct {
	SearchTerm string
	FlagType   string
	Reason     string
}

func RunCleaner(accountID string) (map[string]interface{}, error) {
	conn, err := auth_flow.GetGoogleAdsClient()
	if err != nil {
		return nil, fmt.Errorf("auth error: %v", err)
	}
	defer conn.Close()

	// Set a timeout for the API request
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := services.NewGoogleAdsServiceClient(conn)

	query := `
		SELECT
			search_term_view.search_term,
			metrics.impressions,
			metrics.clicks,
			metrics.conversions
		FROM search_term_view
		WHERE segments.date DURING LAST_30_DAYS
		  AND campaign.advertising_channel_type = 'SEARCH'`

	req := &services.SearchGoogleAdsRequest{
		CustomerId: accountID,
		Query:      query,
	}

	resp, err := client.Search(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Google Ads API search failed for account %s: %v", accountID, err)
	}

	searchTerms := make(map[string]bool)
	for _, row := range resp.Results {
		stv := row.GetSearchTermView()
		if stv == nil {
			continue
		}
		term := stv.GetSearchTerm()
		if term != "" {
			searchTerms[term] = true
			log.Printf("Term: %-40s | Impressions: %d | Clicks: %d | Conversions: %d",
				term,
				row.GetMetrics().GetImpressions(),
				row.GetMetrics().GetClicks(),
				row.GetMetrics().GetConversions())
		}
	}

	if len(searchTerms) == 0 {
		log.Printf("No search terms found for account %s", accountID)
		return map[string]interface{}{
			"status":     "no_data",
			"account_id": accountID,
		}, nil
	}

	// Extract terms, sort for consistency
	var allTerms []string
	for term := range searchTerms {
		allTerms = append(allTerms, term)
	}
	sort.Strings(allTerms)

	log.Printf("Account %s: Pulled %d search terms", accountID, len(allTerms))

	// Split into disqualified and to-review
	autoExcluded, toReview := splitTerms(allTerms)
	log.Printf("Account %s: Auto-excluded %d terms, flagged %d for AI review",
		accountID, len(autoExcluded), len(toReview))

	aiFlagged := shared.AIFlagTerms(toReview)
	logAIFlags(accountID, aiFlagged)

	// Build flag list
	var allFlags []FlaggedTerm
	for _, term := range autoExcluded {
		allFlags = append(allFlags, FlaggedTerm{
			SearchTerm: term,
			FlagType:   "irrelevant",
			Reason:     "Matched disqualifier list",
		})
	}
	for _, item := range aiFlagged {
		allFlags = append(allFlags, FlaggedTerm{
			SearchTerm: item.SearchTerm,
			FlagType:   item.FlagType,
			Reason:     item.Reason,
		})
	}

	result, err := applyExclusions(client, accountID, allFlags)
	if err != nil {
		return nil, fmt.Errorf("exclusion error for account %s: %v", accountID, err)
	}

	return map[string]interface{}{
		"status":           "success",
		"account_id":       accountID,
		"auto_excluded":    autoExcluded,
		"ai_flagged":       aiFlagged,
		"flagged_count":    len(aiFlagged),
		"exclusion_result": result,
	}, nil
}

// splitTerms filters search terms into auto-excluded and AI-review groups
func splitTerms(terms []string) (autoExcluded []string, toReview []string) {
	for _, term := range terms {
		if shared.IsDisqualified(term) {
			autoExcluded = append(autoExcluded, term)
		} else {
			toReview = append(toReview, term)
		}
	}
	return
}

// logAIFlags summarizes AI-flagged results in the logs
func logAIFlags(accountID string, flags []FlaggedTerm) {
	if len(flags) == 0 {
		return
	}
	log.Printf("Account %s: AI flagged %d terms for manual review", accountID, len(flags))
	for _, f := range flags {
		log.Printf("  → [%s] %s | Reason: %s", f.FlagType, f.SearchTerm, f.Reason)
	}
}
