package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"testing"
)

type ResponseUnauthorized struct {
	BaseType string `json:"baseType"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Type     string `json:"type"`
}

type Response struct {
	Uuid           string `json:"uuid"`
	UserPrincipal  string `json:"userPrincipal"`
	UserId         string `json:"userId"`
	Type           string `json:"type"`
	Ttl            int    `json:"ttl"`
	Token          string `json:"token"`
	Name           string `json:"name"`
	Links          Link   `json:"links"`
	LastUpdateTime string `json:"lastUpdateTime"`
	Labels         Label  `json:"labels"`
	IsDerived      bool   `json:"isDerived"`
	Id             string `json:"id"`
	ExpiresAt      string `json:"expiresAt"`
	Expired        bool   `json:"expired"`
	Enabled        bool   `json:"enabled"`
	Current        bool   `json:"current"`
	Description    string `json:"description"`
	CreatorId      string `json:"creatorId"`
	Created        string `json:"created"`
	BaseType       string `json:"baseType"`
	AuthProvider   string `json:"authProvider"`
	ClusterId      string `json:"clusterId"`
	CreatedTS      int    `json:"createdTS"`
}

type Label struct {
	Kind        string `json:"authn.management.cattle.io/kind"`
	TokenUserId string `json:"authn.management.cattle.io/token-userId"`
	Creator     string `json:"cattle.io/creator"`
}

type Link struct {
	Self string `json:"self"`
}

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestLoginAPI(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		body       interface{}
		statusCode string
	}{
		{
			name:   "POST valid user for action=login",
			method: http.MethodPost,
			url:    "https://localhost/v3-public/localProviders/local?action=login",
			body: Body{
				Username: "admin",
				Password: "1234567abcdef",
			},
			statusCode: "201 Created",
		},
		{
			name:   "POST invalid user for action=login",
			method: http.MethodPost,
			url:    "https://localhost/v3-public/localProviders/local?action=login",
			body: Body{
				Username: "admin",
				Password: "errorerror",
			},
			statusCode: "401 Unauthorized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jsonData []byte
			if tt.body != nil {
				var err error
				jsonData, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("could not marshal body: %v", err)
				}
			}

			req, err := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			customClient := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}

			resp, err := customClient.Do(req)
			if err != nil {
				t.Fatalf("could not send request: %v", err)
			}
			defer resp.Body.Close()

			if status := resp.Status; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.statusCode)
			}
		})
	}
}
