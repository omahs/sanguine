// Package client provides a client for the RFQ quoting API.
// nolint:,wrapcheck
package client

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"

	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"

	"github.com/dubonzi/otelresty"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
)

// AuthenticatedClient is an interface for the RFQ API.
// It provides methods for creating, retrieving and updating quotes.
type AuthenticatedClient interface {
	PutQuote(q *model.PutQuoteRequest) error
	UnauthenticatedClient
}

// UnauthenticatedClient is an interface for the RFQ API.
type UnauthenticatedClient interface {
	GetAllQuotes() ([]*model.GetQuoteResponse, error)
	GetSpecificQuote(q *model.GetQuoteSpecificRequest) ([]*model.GetQuoteResponse, error)
	GetQuoteByRelayerAddress(relayerAddr string) ([]*model.GetQuoteResponse, error)
	resty() *resty.Client
}

type unauthenticatedClient struct {
	rClient *resty.Client
}

func (c unauthenticatedClient) resty() *resty.Client {
	return c.rClient
}

type clientImpl struct {
	UnauthenticatedClient
	rClient *resty.Client
}

// NewAuthenticatedClient creates a new client for the RFQ quoting API.
// TODO: @aurelius,  you don't actually need to be authed for GET Requests.
func NewAuthenticatedClient(metrics metrics.Handler, rfqURL string, reqSigner signer.Signer) (AuthenticatedClient, error) {
	unauthedClient, err := NewUnauthenticaedClient(metrics, rfqURL)
	if err != nil {
		return nil, fmt.Errorf("could not create unauthenticated client: %w", err)
	}

	// since this is a pointer, all requests going forward will be authenticated. This is assigned
	// to a new variable for clarity.
	authedClient := unauthedClient.resty().
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			// if request.Method == "PUT" && request.URL == rfqURL+rest.QUOTE_ROUTE {
			// i.e. signature (hex encoded) = keccak(bytes.concat("\x19Ethereum Signed Message:\n", len(strconv.Itoa(time.Now().Unix()), strconv.Itoa(time.Now().Unix())))
			// so that full auth header string: auth = strconv.Itoa(time.Now().Unix()) + ":" + signature
			// Get the current Unix timestamp as a string.
			now := strconv.Itoa(int(time.Now().Unix()))

			// Prepare the data to be signed.
			data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now

			sig, err := reqSigner.SignMessage(request.Context(), []byte(data), true)

			if err != nil {
				return fmt.Errorf("failed to sign request: %w", err)
			}

			res := fmt.Sprintf("%s:%s", now, hexutil.Encode(signer.Encode(sig)))
			request.SetHeader("Authorization", res)

			return nil
		})

	return &clientImpl{
		UnauthenticatedClient: unauthedClient,
		rClient:               authedClient,
	}, nil
}

// NewUnauthenticaedClient creates a new client for the RFQ quoting API.
func NewUnauthenticaedClient(metricHandler metrics.Handler, rfqURL string) (UnauthenticatedClient, error) {
	client := resty.New().
		SetBaseURL(rfqURL).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			request.Header.Add(ginhelper.RequestIDHeader, uuid.New().String())
			return nil
		})

	otelresty.TraceClient(client, otelresty.WithTracerProvider(metricHandler.GetTracerProvider()))
	return &unauthenticatedClient{client}, nil
}

// PutQuote puts a new quote in the RFQ quoting API.
func (c *clientImpl) PutQuote(q *model.PutQuoteRequest) error {
	res, err := c.rClient.R().
		SetBody(q).
		Put(rest.QuoteRoute)

	// TODO: Figure out if there's anything to do with the response, right now it's result: Status Code 200 OK
	_ = res

	return err
}

// GetAllQuotes retrieves all quotes from the RFQ quoting API.
func (c *unauthenticatedClient) GetAllQuotes() ([]*model.GetQuoteResponse, error) {
	var quotes []*model.GetQuoteResponse
	resp, err := c.rClient.R().
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s: %w", resp.Status(), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", resp.Status())
	}

	return quotes, nil
}

// GetSpecificQuote retrieves a specific quote from the RFQ quoting API.
func (c *unauthenticatedClient) GetSpecificQuote(q *model.GetQuoteSpecificRequest) ([]*model.GetQuoteResponse, error) {
	var quotes []*model.GetQuoteResponse
	resp, err := c.rClient.R().
		SetQueryParams(map[string]string{
			"originChainId":   strconv.Itoa(q.OriginChainID),
			"originTokenAddr": q.OriginTokenAddr,
			"destChainId":     strconv.Itoa(q.DestChainID),
			"destTokenAddr":   q.DestTokenAddr,
		}).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s: %w", resp.Status(), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", resp.Status())
	}

	return quotes, nil
}

func (c *unauthenticatedClient) GetQuoteByRelayerAddress(relayerAddr string) ([]*model.GetQuoteResponse, error) {
	var quotes []*model.GetQuoteResponse
	resp, err := c.rClient.R().
		SetQueryParams(map[string]string{
			"relayerAddr": relayerAddr,
		}).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s %w", resp.Status(), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", resp.Status())
	}

	return quotes, nil
}
