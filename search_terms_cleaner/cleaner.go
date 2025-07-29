package search_terms_cleaner

import (
	"context"
	"fmt"
	"log"
	"strings"

	"Search_Term_Cleaner/auth_flow"
	"Search_Term_Cleaner/shared"

	"github.com/GoogleAds/google-ads-go/v14/services"
	"github.com/GoogleAds/google-ads-go/v14/resources"
)

type FlaggedTerm struct {
	SearchTerm string
	FlagType   string
	Reason     string
}

func RunCleaner(accountID string) (map[string]interface{}, error) {
	client, err := auth_flow.LoadGoogleAdsClient(accountID)
	if err != nil {
		return nil, fmt.Errorf("auth error: %v", err)
	}
	ctx := context.Background()
	gaService := client.Services.GoogleAdsService

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

	it := gaService.Search(ctx, req)
	searchTerms := make(map[string]bool)

	for {
		resp, err := it.Next()
		if err != nil {
			break
		}
		term := resp.GetSearchTermView().GetSearchTerm()
		if term != "" {
			searchTerms[term] = true
		}
	}

	if len(searchTerms) == 0 {
		return map[string]interface{}{
			"status":     "no_data",
			"account_id": accountID,
		}, nil
	}

	var allTerms []string
	for term := range searchTerms {
		allTerms = append(allTerms, term)
	}

	autoExcluded := []string{}
	toReview := []string{}

	for _, term := range allTerms {
		if shared.IsDisqualified(term) {
			autoExcluded = append(autoExcluded, term)
		} else {
			toReview = append(toReview, term)
		}
	}

	aiFlagged := shared.AIFlagTerms(toReview)

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
		return nil, err
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
