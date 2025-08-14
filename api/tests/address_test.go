package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jaqueline/handlers/address"
	"github.com/library/modules/address/dto"
	"github.com/library/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	router          *gin.Engine
	authHeaderUser1 string
	authHeaderUser2 string
	authHeaderAdmin string
	recordID1       string
	recordID2       string
	basePath        string
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	token1 := os.Getenv("AUTH_TOKEN_USER1")
	if token1 == "" {
		panic("AUTH_TOKEN_USER1 env var missing")
	}
	authHeaderUser1 = "Bearer " + token1

	token2 := os.Getenv("AUTH_TOKEN_USER2")
	if token2 == "" {
		panic("AUTH_TOKEN_USER2 env var missing")
	}
	authHeaderUser2 = "Bearer " + token2

	tokenAdmin := os.Getenv("AUTH_TOKEN_ADMIN")
	if tokenAdmin == "" {
		panic("AUTH_TOKEN_ADMIN env var missing")
	}
	authHeaderAdmin = "Bearer " + tokenAdmin

	basePath = address.GetBasePath()
	router = newRouter()

	cleanRecords(authHeaderUser1)
	cleanRecords(authHeaderUser2)
	cleanRecords(authHeaderAdmin)
	code := m.Run()
	cleanRecords(authHeaderUser1)
	cleanRecords(authHeaderUser2)
	cleanRecords(authHeaderAdmin)

	os.Exit(code)
}

func newRouter() *gin.Engine {
	r := gin.Default()
	group := r.Group(basePath)
	address.RegisterRoutes(group)
	return r
}

func cleanRecords(authHeader string) {
	records, err := getAll(authHeader)
	if err != nil {
		println("Cleanup fetch failed:", err.Error())
		return
	}
	for _, r := range records {
		deleteByID(r.Id, authHeader)
	}
}

func TestCreate(t *testing.T) {
	payload := dto.CreateRequestDTO{
		City:         "São Paulo",
		Country:      "Brasil",
		Neighborhood: "Jardim América",
		Number:       "123",
		State:        "SP",
		Street:       "Rua das Flores",
		Zip:          "01234-567",
	}

	body, err := json.Marshal(payload)
	require.NoError(t, err)

	resp := request(t, "POST", basePath, body, authHeaderUser1)
	require.Equal(t, http.StatusCreated, resp.Code)
	rec := parseOne(t, resp.Body.Bytes())
	assert.Equal(t, "São Paulo", rec.City)
	recordID1 = rec.Id

	payload.City = "Rio de Janeiro"
	payload.Number = "456"
	payload.Zip = "23456-789"

	body, err = json.Marshal(payload)
	require.NoError(t, err)

	resp = request(t, "POST", basePath, body, authHeaderUser1)
	require.Equal(t, http.StatusCreated, resp.Code)
	rec = parseOne(t, resp.Body.Bytes())
	assert.Equal(t, "Rio de Janeiro", rec.City)
	recordID2 = rec.Id
}

func TestGet(t *testing.T) {
	if recordID2 == "" {
		t.Skip("Skip test: no recordID2")
	}
	resp := request(t, "GET", path.Join(basePath, recordID2), nil, authHeaderUser1)
	require.Equal(t, http.StatusOK, resp.Code)

	rec := parseOne(t, resp.Body.Bytes())
	assert.Equal(t, recordID2, rec.Id)
	assert.Equal(t, "Rio de Janeiro", rec.City)
}

func TestUpdate(t *testing.T) {
	if recordID2 == "" {
		t.Skip("Skip test: no recordID2")
	}
	payload := map[string]interface{}{
		"zip":    "99999",
		"city":   "Updated City Two",
		"state":  "UT",
		"street": "Updated Street",
		"number": "999",
	}
	body, _ := json.Marshal(payload)

	resp := request(t, "PATCH", path.Join(basePath, recordID2), body, authHeaderUser1)
	require.Equal(t, http.StatusOK, resp.Code)

	rec := parseOne(t, resp.Body.Bytes())
	assert.Equal(t, recordID2, rec.Id)
	assert.Equal(t, "Updated City Two", rec.City)
}

func TestList(t *testing.T) {
	resp := request(t, "GET", basePath, nil, authHeaderUser1)
	require.Equal(t, http.StatusOK, resp.Code)

	records := parseList(t, resp.Body.Bytes())
	require.NotEmpty(t, records)
	require.Equal(t, len(records), 2)

	found1, found2 := false, false
	for _, r := range records {
		if r.Id == recordID1 {
			found1 = true
			assert.Equal(t, "São Paulo", r.City)
		}
		if r.Id == recordID2 {
			found2 = true
			assert.Equal(t, "Updated City Two", r.City)
		}
	}
	assert.True(t, found1)
	assert.True(t, found2)
}

func TestDelete(t *testing.T) {
	if recordID1 == "" {
		t.Skip("Skip test: no recordID1")
	}
	resp := request(t, "DELETE", path.Join(basePath, recordID1), nil, authHeaderUser1)
	require.Equal(t, http.StatusNoContent, resp.Code)
	recordID1 = ""
}

func TestUser2AccessDenied(t *testing.T) {
	if recordID2 == "" {
		t.Skip("Skip test: no recordID2")
	}

	respGet := request(t, "GET", path.Join(basePath, recordID2), nil, authHeaderUser2)
	assert.Equal(t, http.StatusForbidden, respGet.Code)

	payload := map[string]interface{}{
		"city": "Hacker City",
	}
	body, _ := json.Marshal(payload)
	respUpdate := request(t, "PATCH", path.Join(basePath, recordID2), body, authHeaderUser2)
	assert.Equal(t, http.StatusForbidden, respUpdate.Code)

	filter := dto.FilterRequestDTO{
		Owner:    "3aad236f-dd4a-47af-bf7e-968dc3fa4001",
		Page:     1,
		PageSize: 10,
	}
	queryParams := fmt.Sprintf("?owner=%s&page=%d&pageSize=%d", filter.Owner, filter.Page, filter.PageSize)
	listURL := basePath + queryParams

	resp := request(t, "GET", listURL, nil, authHeaderUser2)
	require.Equal(t, http.StatusForbidden, resp.Code)

	respDelete := request(t, "DELETE", path.Join(basePath, recordID2), nil, authHeaderUser2)
	assert.Equal(t, http.StatusForbidden, respDelete.Code)
}

func TestAdminGetByID(t *testing.T) {
	if recordID2 == "" {
		t.Skip("Skip test: no recordID2")
	}

	resp := request(t, "GET", path.Join(basePath, recordID2), nil, authHeaderAdmin)
	assert.Equal(t, http.StatusOK, resp.Code)

	rec := parseOne(t, resp.Body.Bytes())
	assert.Equal(t, recordID2, rec.Id)
}

func TestAdminUpdateCity(t *testing.T) {
	if recordID2 == "" {
		t.Skip("Skip test: no recordID2")
	}

	payload := map[string]interface{}{
		"city": "Admin City",
	}
	body, err := json.Marshal(payload)
	require.NoError(t, err)

	resp := request(t, "PATCH", path.Join(basePath, recordID2), body, authHeaderAdmin)
	assert.Equal(t, http.StatusOK, resp.Code)

	rec := parseOne(t, resp.Body.Bytes())
	assert.Equal(t, "Admin City", rec.City)
}

func TestAdminHasFullAccessToAnyRecord(t *testing.T) {
	filter := dto.FilterRequestDTO{
		Owner:    "3aad236f-dd4a-47af-bf7e-968dc3fa4001",
		Page:     1,
		PageSize: 10,
	}
	query := fmt.Sprintf("?owner=%s&page=%d&pageSize=%d&only_me=false", filter.Owner, filter.Page, filter.PageSize)
	url := basePath + query

	resp := request(t, "GET", url, nil, authHeaderAdmin)
	require.Equal(t, http.StatusOK, resp.Code)

	records := parseList(t, resp.Body.Bytes())
	require.NotEmpty(t, records)
	require.Equal(t, 1, len(records))
}

func TestListsAllRecords(t *testing.T) {
	filter := dto.FilterRequestDTO{
		Page:     1,
		PageSize: 10,
	}
	url := fmt.Sprintf("%s?page=%d&pageSize=%d&only_me=false", basePath, filter.Page, filter.PageSize)

	resp := request(t, "GET", url, nil, authHeaderAdmin)
	require.Equal(t, http.StatusOK, resp.Code)

	records := parseList(t, resp.Body.Bytes())
	require.NotEmpty(t, records)
	require.GreaterOrEqual(t, len(records), 1)
}

func TestListsRecordsForCurrentUser(t *testing.T) {
	filter := dto.FilterRequestDTO{
		Page:     1,
		PageSize: 10,
	}
	url := fmt.Sprintf("%s?page=%d&pageSize=%d", basePath, filter.Page, filter.PageSize)

	resp := request(t, "GET", url, nil, authHeaderAdmin)
	require.Equal(t, http.StatusOK, resp.Code)

	records := parseList(t, resp.Body.Bytes())
	require.Empty(t, records)
	require.GreaterOrEqual(t, len(records), 0)
}

func TestAdminDeleteRecord(t *testing.T) {
	if recordID2 == "" {
		t.Skip("Skip test: no recordID2")
	}

	resp := request(t, "DELETE", path.Join(basePath, recordID2), nil, authHeaderAdmin)
	assert.Equal(t, http.StatusNoContent, resp.Code)

	recordID2 = ""
}

func request(t *testing.T, method, url string, body []byte, authHeader string) *httptest.ResponseRecorder {
	t.Helper()

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Authorization", authHeader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func parseOne(t *testing.T, body []byte) *dto.ResponseDTO {
	t.Helper()

	var resp utils.APIResponse
	require.NoError(t, json.Unmarshal(body, &resp))

	msgBytes, err := json.Marshal(resp.Message)
	require.NoError(t, err)

	var rec dto.ResponseDTO
	require.NoError(t, json.Unmarshal(msgBytes, &rec))
	return &rec
}

func parseList(t *testing.T, body []byte) []dto.ResponseDTO {
	t.Helper()

	var resp utils.APIResponse
	require.NoError(t, json.Unmarshal(body, &resp))

	msgBytes, err := json.Marshal(resp.Message)
	require.NoError(t, err)

	var recs []dto.ResponseDTO
	require.NoError(t, json.Unmarshal(msgBytes, &recs))
	return recs
}

func getAll(authHeader string) ([]dto.ResponseDTO, error) {
	req, _ := http.NewRequest("GET", basePath, nil)
	req.Header.Set("Authorization", authHeader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		return nil, http.ErrHandlerTimeout
	}

	return parseListNoTest(w.Body.Bytes())
}

func parseListNoTest(body []byte) ([]dto.ResponseDTO, error) {
	var resp utils.APIResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	msgBytes, err := json.Marshal(resp.Message)
	if err != nil {
		return nil, err
	}

	var recs []dto.ResponseDTO
	if err := json.Unmarshal(msgBytes, &recs); err != nil {
		return nil, err
	}

	return recs, nil
}

func deleteByID(id string, authHeader string) {
	req, _ := http.NewRequest("DELETE", path.Join(basePath, id), nil)
	req.Header.Set("Authorization", authHeader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		println("Failed to delete record ID:", id, "Status:", w.Code)
	}
}
