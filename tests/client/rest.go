package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// RunRESTTests exercises the TestApiRest service, whose methods are served on
// REST routes (verb + path template) instead of the default webrpc dispatch.
func RunRESTTests(ctx context.Context, serverURL string) error {
	var errs []error

	api := NewTestApiRestClient(serverURL, &http.Client{})

	// POST /rpc/items — request fields in the JSON body.
	created, err := api.CreateItem(ctx, CreateItemRequest{Item: &Item{Id: 7, Name: "alice"}})
	if err != nil {
		errs = append(errs, fmt.Errorf("CreateItem(): %w", err))
	} else if created.Item == nil || created.Item.Name != "alice" {
		errs = append(errs, fmt.Errorf("CreateItem(): unexpected response %+v", created))
	}

	// GET /rpc/items/{id} — path param plus optional query param.
	details := true
	got, err := api.GetItem(ctx, GetItemRequest{Id: 7, Details: &details})
	if err != nil {
		errs = append(errs, fmt.Errorf("GetItem(): %w", err))
	} else if got.Item == nil || got.Item.Name != "alice (details)" {
		errs = append(errs, fmt.Errorf("GetItem(): unexpected response %+v", got.Item))
	}

	// GET /rpc/items — query params: repeated keys, required scalar, absent optional.
	list, err := api.ListItems(ctx, ListItemsRequest{Tags: []string{"a", "b"}, Limit: 5})
	if err != nil {
		errs = append(errs, fmt.Errorf("ListItems(): %w", err))
	} else {
		wantEcho := "q=false tags=[a b] limit=5"
		found := false
		for _, item := range list.Items {
			if item.Name == wantEcho {
				found = true
			}
		}
		if !found {
			errs = append(errs, fmt.Errorf("ListItems(): echo item %q not found in %+v", wantEcho, list.Items))
		}
	}

	// PUT /rpc/items/{id} — path param plus JSON body (normal form).
	updated, err := api.UpdateItem(ctx, 7, "bob")
	if err != nil {
		errs = append(errs, fmt.Errorf("UpdateItem(): %w", err))
	} else if updated == nil || updated.Name != "bob" {
		errs = append(errs, fmt.Errorf("UpdateItem(): unexpected response %+v", updated))
	}

	// PATCH /rpc/items/{id} — optional body field present.
	name := "carol"
	patched, err := api.PatchItem(ctx, 7, &name)
	if err != nil {
		errs = append(errs, fmt.Errorf("PatchItem(): %w", err))
	} else if patched == nil || patched.Name != "carol" {
		errs = append(errs, fmt.Errorf("PatchItem(): unexpected response %+v", patched))
	}

	// PATCH with absent optional body field keeps the value.
	patched, err = api.PatchItem(ctx, 7, nil)
	if err != nil {
		errs = append(errs, fmt.Errorf("PatchItem(nil): %w", err))
	} else if patched == nil || patched.Name != "carol" {
		errs = append(errs, fmt.Errorf("PatchItem(nil): unexpected response %+v", patched))
	}

	// QUERY /rpc/items/search — the experimental QUERY verb.
	found, err := api.QueryItems(ctx, "carol")
	if err != nil {
		errs = append(errs, fmt.Errorf("QueryItems(): %w", err))
	} else if len(found) != 1 || found[0].Name != "carol" {
		errs = append(errs, fmt.Errorf("QueryItems(): unexpected response %+v", found))
	}

	// GET /rpc/items/echo/params — every scalar query param type round-trips.
	ts := time.Date(2026, 9, 3, 12, 30, 15, 0, time.UTC)
	echo, err := api.EchoParams(ctx, "hello world/&?=", -42, true, 1.5, ts, Status_NOT_AVAILABLE, []string{"x y", "z"}, nil)
	if err != nil {
		errs = append(errs, fmt.Errorf("EchoParams(): %w", err))
	} else {
		want := "str=hello world/&?= num=-42 flag=true fl=1.5 ts=2026-09-03T12:30:15Z status=1 tags=[x y z] opt=<nil>"
		if echo != want {
			errs = append(errs, fmt.Errorf("EchoParams(): got %q, want %q", echo, want))
		}
	}

	// GET /rpc/items/{id}/children/{childId} — multiple path params; the
	// static route /rpc/items/echo/params must win over /rpc/items/{id}/....
	childPath, err := api.GetChild(ctx, 1, 2)
	if err != nil {
		errs = append(errs, fmt.Errorf("GetChild(): %w", err))
	} else if childPath != "/items/1/children/2" {
		errs = append(errs, fmt.Errorf("GetChild(): unexpected response %q", childPath))
	}

	// A method without a route keeps the default webrpc dispatch.
	item, err := api.RpcMethod(ctx, 7)
	if err != nil {
		errs = append(errs, fmt.Errorf("RpcMethod(): %w", err))
	} else if item == nil || item.Name != "carol" {
		errs = append(errs, fmt.Errorf("RpcMethod(): unexpected response %+v", item))
	}

	// A routed method also keeps its default webrpc path.
	resp, err := http.Post(serverURL+"/rpc/TestApiRest/GetItem", "application/json", strings.NewReader(`{"id": 7}`))
	if err != nil {
		errs = append(errs, fmt.Errorf("POST /rpc/TestApiRest/GetItem: %w", err))
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK || !strings.Contains(string(body), `"name":"carol"`) {
			errs = append(errs, fmt.Errorf("POST /rpc/TestApiRest/GetItem: expected 200 with item, got %d %s", resp.StatusCode, body))
		}
	}

	// DELETE /rpc/items/{id}.
	if err := api.DeleteItem(ctx, 7); err != nil {
		errs = append(errs, fmt.Errorf("DeleteItem(): %w", err))
	}
	if _, err := api.GetItem(ctx, GetItemRequest{Id: 7}); err == nil {
		errs = append(errs, fmt.Errorf("GetItem() after DeleteItem(): expected error, got nil"))
	}

	// Path params overwrite conflicting JSON body fields (decode order:
	// body -> path).
	if _, err := api.CreateItem(ctx, CreateItemRequest{Item: &Item{Id: 9, Name: "nine"}}); err != nil {
		errs = append(errs, fmt.Errorf("CreateItem(9): %w", err))
	}
	overwriteReq, err := http.NewRequestWithContext(ctx, "PUT", serverURL+"/rpc/items/9", strings.NewReader(`{"id": 1, "name": "dave"}`))
	if err != nil {
		errs = append(errs, fmt.Errorf("PUT /rpc/items/9: %w", err))
	} else {
		overwriteReq.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(overwriteReq)
		if err != nil {
			errs = append(errs, fmt.Errorf("PUT /rpc/items/9: %w", err))
		} else {
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				errs = append(errs, fmt.Errorf("PUT /rpc/items/9: expected 200, got %d", resp.StatusCode))
			}
			if got, err := api.GetItem(ctx, GetItemRequest{Id: 9}); err != nil || got.Item.Name != "dave" {
				errs = append(errs, fmt.Errorf("GetItem(9) after body/path conflict: %+v, %v", got, err))
			}
			if _, err := api.GetItem(ctx, GetItemRequest{Id: 1}); err == nil {
				errs = append(errs, fmt.Errorf("GetItem(1): expected error, the path param must win over the body id"))
			}
		}
	}

	// Wrong verb on an existing path: webrpc error envelope with Allow header.
	resp, err = http.Post(serverURL+"/rpc/items/7", "application/json", strings.NewReader("{}"))
	if err != nil {
		errs = append(errs, fmt.Errorf("POST /rpc/items/7: %w", err))
	} else {
		resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			errs = append(errs, fmt.Errorf("POST /rpc/items/7: expected an error status, got 200"))
		}
		if allow := resp.Header.Get("Allow"); allow == "" {
			errs = append(errs, fmt.Errorf("POST /rpc/items/7: expected Allow header"))
		}
	}

	// Path param that fails type coercion: HTTP 400 in the error envelope,
	// with a user-facing message naming the param and its expected type.
	resp, err = http.Get(serverURL + "/rpc/items/not-a-number")
	if err != nil {
		errs = append(errs, fmt.Errorf("GET /rpc/items/not-a-number: %w", err))
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			errs = append(errs, fmt.Errorf("GET /rpc/items/not-a-number: expected 400, got %d", resp.StatusCode))
		}
		wantCause := `invalid path param 'id': \"not-a-number\" is not a valid uint64`
		if !strings.Contains(string(body), wantCause) {
			errs = append(errs, fmt.Errorf("GET /rpc/items/not-a-number: error cause %q not found in %s", wantCause, body))
		}
	}

	// Body routes reject non-JSON requests, so cross-site form posts (sent
	// without a CORS preflight) can't reach handlers — with or without inputs.
	for _, path := range []string{"/rpc/items", "/rpc/items/reset"} {
		resp, err := http.Post(serverURL+path, "text/plain", strings.NewReader(`{"item":{"id":5,"name":"csrf"}}`))
		if err != nil {
			errs = append(errs, fmt.Errorf("text/plain POST %s: %w", path, err))
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			errs = append(errs, fmt.Errorf("text/plain POST %s: expected 400, got %d", path, resp.StatusCode))
		}
	}
	if _, err := api.GetItem(ctx, GetItemRequest{Id: 5}); err == nil {
		errs = append(errs, fmt.Errorf("GetItem(5): the text/plain POST must not have created an item"))
	}

	// A body route without inputs still works from the generated client.
	if err := api.ResetItems(ctx); err != nil {
		errs = append(errs, fmt.Errorf("ResetItems(): %w", err))
	}

	// A missing required query param is a 400, not a zero value.
	resp, err = http.Get(serverURL + "/rpc/items?tags=a")
	if err != nil {
		errs = append(errs, fmt.Errorf("GET /rpc/items without limit: %w", err))
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest || !strings.Contains(string(body), "missing query param 'limit'") {
			errs = append(errs, fmt.Errorf("GET /rpc/items without limit: expected 400 missing query param, got %d %s", resp.StatusCode, body))
		}
	}

	// An empty required list sends no query keys and decodes as empty.
	emptyList, err := api.ListItems(ctx, ListItemsRequest{Tags: []string{}, Limit: 1})
	if err != nil {
		errs = append(errs, fmt.Errorf("ListItems(empty tags): %w", err))
	} else if len(emptyList.Items) == 0 || emptyList.Items[0].Name != "q=false tags=[] limit=1" {
		errs = append(errs, fmt.Errorf("ListItems(empty tags): unexpected response %+v", emptyList.Items))
	}

	// Unknown path: HTTP 404 in the error envelope.
	resp, err = http.Get(serverURL + "/rpc/nope")
	if err != nil {
		errs = append(errs, fmt.Errorf("GET /rpc/nope: %w", err))
	} else {
		resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			errs = append(errs, fmt.Errorf("GET /rpc/nope: expected 404, got %d", resp.StatusCode))
		}
	}

	if len(errs) > 0 {
		errStrings := []string{}
		for _, err := range errs {
			errStrings = append(errStrings, err.Error())
		}
		return fmt.Errorf("REST tests failed:\n%v", strings.Join(errStrings, "\n"))
	}

	return nil
}
