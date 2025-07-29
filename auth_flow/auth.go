package auth_flow

import (
	"context"
	"fmt"
	"os"

	"github.com/googleads/google-ads-go/v16/services"
	"github.com/googleads/google-ads-go/v16"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func GetGoogleAdsClient(developerToken, loginCustomerID string) (*grpc.ClientConn, error) {
	refreshToken := os.Getenv("GOOGLE_ADS_REFRESH_TOKEN")
	clientID := os.Getenv("GOOGLE_ADS_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_ADS_CLIENT_SECRET")

	if refreshToken == "" || clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("missing Google Ads environment variables")
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"googleads.googleapis.com:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauthCredentials{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
			DeveloperToken: developerToken,
			LoginCustomerID: loginCustomerID,
		}),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Custom struct to attach OAuth2 and metadata
type oauthCredentials struct {
	ClientID        string
	ClientSecret    string
	RefreshToken    string
	DeveloperToken  string
	LoginCustomerID string
}

func (c oauthCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// You'd normally swap the refresh token for an access token here
	// For simplicity, we'll assume a valid token is injected at runtime
	return map[string]string{
		"developer-token":      c.DeveloperToken,
		"login-customer-id":    c.LoginCustomerID,
		"authorization":        fmt.Sprintf("Bearer %s", c.RefreshToken), // not real usage
	}, nil
}

func (c oauthCredentials) RequireTransportSecurity() bool {
	return true
}
