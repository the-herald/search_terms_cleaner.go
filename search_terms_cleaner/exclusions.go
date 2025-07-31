package search_terms_cleaner

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/genproto/googleapis/ads/googleads/v20/common"
	"google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
	"google.golang.org/genproto/googleapis/ads/googleads/v20/resources"
	"google.golang.org/genproto/googleapis/ads/googleads/v20/services"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"search_terms_cleaner/shared"
)

const (
	lowQualitySetName = "Low-quality Searches & Words"
	competitorSetName = "Competitors"
)

func applyExclusions(client services.GoogleAdsServiceClient, accountID string, flags []FlaggedTerm) (interface{}, error) {
	ctx := context.Background()
	var operations []*services.MutateOperation

	// Step 1: Ensure shared sets exist or create them
	setIDs := map[string]int64{}
	sets := []string{lowQualitySetName, competitorSetName}
	for _, name := range sets {
		setID, err := getOrCreateSharedSet(ctx, client, accountID, name)
		if err != nil {
			return nil, fmt.Errorf("failed to get/create shared set '%s': %v", name, err)
		}
		setIDs[name] = setID
	}

	// Step 2: Prepare negative keywords
	for _, flag := range flags {
		var setName string
		if flag.FlagType == "competitor" {
			setName = competitorSetName
		} else {
			setName = lowQualitySetName
		}
		sharedSetID := setIDs[setName]

		// Add full term
		operations = append(operations, buildNegKeywordOp(accountID, sharedSetID, flag.SearchTerm))

		// Add trouble word (via AI)
		root := shared.ExtractTroubleWord(flag.SearchTerm, flag.Reason)
		if root != "" && root != flag.SearchTerm {
			operations = append(operations, buildNegKeywordOp(accountID, sharedSetID, root))
		}
	}

	if len(operations) == 0 {
		log.Println("No exclusion operations to apply.")
		return "no-op", nil
	}

	resp, err := client.Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       accountID,
		MutateOperations: operations,
	})
	if err != nil {
		return nil, fmt.Errorf("mutate failed: %v", err)
	}

	log.Printf("Applied %d negative keyword operations to account %s", len(operations), accountID)
	return resp, nil
}

func getOrCreateSharedSet(ctx context.Context, client services.GoogleAdsServiceClient, customerID, name string) (int64, error) {
	query := fmt.Sprintf(`
		SELECT shared_set.id, shared_set.name
		FROM shared_set
		WHERE shared_set.name = "%s"
		LIMIT 1`, name)

	resp, err := client.Search(ctx, &services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query:      query,
	})
	if err != nil {
		return 0, err
	}
	for _, row := range resp.Results {
		return row.GetSharedSet().GetId(), nil
	}

	op := &services.MutateOperation{
		Operation: &services.MutateOperation_SharedSetOperation{
			SharedSetOperation: &services.SharedSetOperation{
				Create: &resources.SharedSet{
					Name: wrapperspb.String(name),
					Type: enums.SharedSetTypeEnum_NEGATIVE_KEYWORDS,
				},
			},
		},
	}
	mresp, err := client.Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       customerID,
		MutateOperations: []*services.MutateOperation{op},
	})
	if err != nil || len(mresp.MutateOperationResponses) == 0 {
		return 0, fmt.Errorf("failed to create shared set '%s': %v", name, err)
	}
	return extractIDFromResource(mresp.MutateOperationResponses[0].GetSharedSetResult().GetResourceName()), nil
}

func buildNegKeywordOp(customerID string, sharedSetID int64, phrase string) *services.MutateOperation {
	neg := &common.NegativeKeyword{
		Text:      wrapperspb.String(phrase),
		MatchType: enums.KeywordMatchTypeEnum_PHRASE,
	}
	criterion := &resources.SharedCriterion{
		SharedSet: fmt.Sprintf("customers/%s/sharedSets/%d", customerID, sharedSetID),
		Criterion: &resources.SharedCriterion_Keyword{Keyword: neg},
	}
	return &services.MutateOperation{
		Operation: &services.MutateOperation_SharedCriterionOperation{
			SharedCriterionOperation: &services.SharedCriterionOperation{
				Create: criterion,
			},
		},
	}
}

func extractIDFromResource(resourceName string) int64 {
	parts := strings.Split(resourceName, "/")
	if len(parts) == 0 {
		return 0
	}
	var id int64
	fmt.Sscanf(parts[len(parts)-1], "%d", &id)
	return id
}
