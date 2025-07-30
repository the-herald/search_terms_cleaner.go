package auth_flow

import (
    "context"
    "fmt"
    "os"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

func GetGoogleAdsClient() (*grpc.ClientConn, error) {
    developerToken := os.Getenv("GOOGLE_ADS_DEVELOPER_TOKEN")
    loginCustomerID := os.Getenv("GOOGLE_ADS_LOGIN_CUSTOMER_ID")
    refreshToken := os.Getenv("GOOGLE_ADS_REFRESH_TOKEN")
    clientID := os.Getenv("GOOGLE_ADS_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_ADS_CLIENT_SECRET")

    if developerToken == "" || loginCustomerID == "" || refreshToken == "" || clientID == "" || clientSecret == "" {
        return nil, fmt.Errorf("missing one or more required Google Ads environment variables")
    }

    ctx := context.Background()
    conn, err := grpc.DialContext(
        ctx,
        "googleads.googleapis.com:443",
        grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
        grpc.WithPerRPCCredentials(oauthCredentials{
            ClientID:        clientID,
            ClientSecret:    clientSecret,
            RefreshToken:    refreshToken,
            DeveloperToken:  developerToken,
            LoginCustomerID: loginCustomerID,
        }),
    )
    if err != nil {
        return nil, err
    }

    return conn, nil
}

type oauthCredentials struct {
    ClientID        string
    ClientSecret    string
    RefreshToken    string
    DeveloperToken  string
    LoginCustomerID string
}

func (c oauthCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
    return map[string]string{
        "developer-token":   c.DeveloperToken,
        "login-customer-id": c.LoginCustomerID,
        "authorization":     fmt.Sprintf("Bearer %s", c.RefreshToken),
    }, nil
}

func (c oauthCredentials) RequireTransportSecurity() bool {
    return true
}