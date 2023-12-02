package providusbank_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/CreateDynamicAccount-auth-failed.json
var createDynamicAccountAuthFailed []byte

//go:embed testdata/CreateDynamicAccount-success.json
var createDynamicAccountSuccess []byte

//go:embed testdata/CreateReservedAccount-success.json
var createReservedAccountSuccess []byte

//go:embed testdata/UpdateAccountName-fail.json
var updateAccountNameFail []byte

//go:embed testdata/UpdateAccountName-success.json
var updateAccountNameSuccess []byte

//go:embed testdata/BlacklistAccount-fail.json
var blacklistAccountFail []byte

//go:embed testdata/BlacklistAccount-success.json
var blacklistAccountSuccess []byte

//go:embed testdata/VerifyTransaction-fail.json
var verifyTransactionFail []byte

//go:embed testdata/VerifyTransaction-success.json
var verifyTransactionSuccess []byte

//go:embed testdata/RepushTransaction-fail.json
var repushTransactionFail []byte

func TestCreateDynamicAccount_AuthFailed(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createDynamicAccountAuthFailed))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateDynamicAccount(context.TODO(), providusbank.DynamicAccountPayload{})
	require.NoError(t, err)

	assert.False(t, got.Success)
}

func TestCreateDynamicAccount_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createDynamicAccountSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateDynamicAccount(context.TODO(), providusbank.DynamicAccountPayload{})
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestCreateReservedAccount_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createReservedAccountSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateReservedAccount(context.TODO(), providusbank.ReservedAccountPayload{})
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestUpdateAccountName_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(updateAccountNameFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.UpdateAccountName(context.TODO(), providusbank.UpdateAccountNamePayload{})
	require.NoError(t, err)

	assert.False(t, got.Success)
}

func TestUpdateAccountName_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(updateAccountNameSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.UpdateAccountName(context.TODO(), providusbank.UpdateAccountNamePayload{})
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestBlacklistAccount_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(blacklistAccountFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.BlacklistAccount(context.TODO(), providusbank.BlacklistAccountPayload{})
	require.NoError(t, err)

	assert.False(t, got.Success)
}

func TestBlacklistAccount_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(blacklistAccountSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.BlacklistAccount(context.TODO(), providusbank.BlacklistAccountPayload{})
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestVerifyTransaction_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(verifyTransactionFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.VerifyTransaction(context.TODO(), "")
	require.NoError(t, err)

	assert.Equal(t, "", got.SessionID)
}

func TestVerifyTransaction_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(verifyTransactionSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.VerifyTransaction(context.TODO(), "123456789")
	require.NoError(t, err)

	assert.Equal(t, "123456789", got.SessionID)
}

func TestVerifyTransactionWithSettlementID_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(verifyTransactionSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.VerifyTransactionWithSettlementID(context.TODO(), "204210202000000700001")
	require.NoError(t, err)

	assert.Equal(t, "204210202000000700001", got.SettlementID)
}

func TestRepushTransaction_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(repushTransactionFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.RepushTransaction(context.TODO(), providusbank.RepushTransactionPayload{})
	require.NoError(t, err)

	assert.False(t, got.Success)
}