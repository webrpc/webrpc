package client

import (
	"context"
	"fmt"
	"net/http"
)

func RunTests(serverURL string) error {
	if err := TestComplexData(serverURL); err != nil {
		return err
	}

	return nil
}

func TestComplexData(serverURL string) error {
	httpClient := NewComplexApiClient(serverURL, &http.Client{})

	// get complex data from server
	resp, err := httpClient.GetComplex(context.Background())
	if err != nil {
		return fmt.Errorf("error getting complex data from server: %w", err)
	}

	// send the same data to server & validate
	err = httpClient.SendComplex(context.Background(), resp)
	if err != nil {
		return fmt.Errorf("error sending complex data back to server: %w", err)
	}

	return nil
}
