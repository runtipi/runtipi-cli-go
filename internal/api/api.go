package api

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/steveiliop56/runtipi-cli-go/internal/env"
)

func GenerateJWT() (string, error) {
	secret, envErr := env.GetEnvValue("JWT_SECRET")
	
	if envErr != nil {
		return "", envErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"skill": "issue",
	})

	tokenString, tokenErr := token.SignedString([]byte(secret))

	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenString, nil
}

func ApiRequest(path string, method string) (error) {
	token, tokenErr := GenerateJWT()

	if tokenErr != nil {
		return tokenErr
	}

	port, portErr := env.GetEnvValue("NGINX_PORT")

	if portErr != nil {
		return portErr
	}

	apiUrl := fmt.Sprintf("http://localhost:%s/worker-api/%s", port, path)

	request, requestErr := http.NewRequest(method, apiUrl, bytes.NewBuffer([]byte("")))

	if requestErr != nil {
		return requestErr
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	response, clientErr := client.Do(request)

	if clientErr != nil {
		return clientErr
	}

	defer response.Body.Close()
	
	return nil
} 