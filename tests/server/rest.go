package server

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// TestRESTServer implements the TestApiRest service, whose methods are served
// on REST routes (verb + path template) instead of the default webrpc
// dispatch.
type TestRESTServer struct {
	mu    sync.Mutex
	items map[uint64]*Item
}

func NewTestRESTServer() *TestRESTServer {
	return &TestRESTServer{items: map[uint64]*Item{}}
}

func (s *TestRESTServer) GetItem(ctx context.Context, req GetItemRequest) (*GetItemResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[req.Id]
	if !ok {
		return nil, ErrUnexpectedValue.WithCausef("no item %d", req.Id)
	}
	out := *item
	if req.Details != nil && *req.Details {
		out.Name += " (details)"
	}
	return &GetItemResponse{Item: &out}, nil
}

func (s *TestRESTServer) ListItems(ctx context.Context, req ListItemsRequest) (*ListItemsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Echo the decoded query params back as a synthetic item, so the client
	// can assert the query string decoding.
	echo := fmt.Sprintf("q=%v tags=%v limit=%d", req.Q != nil, req.Tags, req.Limit)
	items := []*Item{{Id: 0, Name: echo}}
	for _, item := range s.items {
		items = append(items, item)
	}
	return &ListItemsResponse{Items: items}, nil
}

func (s *TestRESTServer) CreateItem(ctx context.Context, req CreateItemRequest) (*CreateItemResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Item == nil {
		return nil, ErrMissingArgument.WithCausef("item is required")
	}
	item := *req.Item
	s.items[item.Id] = &item
	return &CreateItemResponse{Item: &item}, nil
}

func (s *TestRESTServer) UpdateItem(ctx context.Context, id uint64, name string) (*Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[id]
	if !ok {
		return nil, ErrUnexpectedValue.WithCausef("no item %d", id)
	}
	item.Name = name
	return item, nil
}

func (s *TestRESTServer) PatchItem(ctx context.Context, id uint64, name *string) (*Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[id]
	if !ok {
		return nil, ErrUnexpectedValue.WithCausef("no item %d", id)
	}
	if name != nil {
		item.Name = *name
	}
	return item, nil
}

func (s *TestRESTServer) DeleteItem(ctx context.Context, id uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[id]; !ok {
		return ErrUnexpectedValue.WithCausef("no item %d", id)
	}
	delete(s.items, id)
	return nil
}

func (s *TestRESTServer) QueryItems(ctx context.Context, q string) ([]*Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := []*Item{}
	for _, item := range s.items {
		if strings.Contains(item.Name, q) {
			items = append(items, item)
		}
	}
	return items, nil
}

func (s *TestRESTServer) EchoParams(ctx context.Context, str string, num int64, flag bool, fl float64, ts time.Time, status Status, tags []string, opt *string) (string, error) {
	optStr := "<nil>"
	if opt != nil {
		optStr = *opt
	}
	return fmt.Sprintf("str=%s num=%d flag=%v fl=%v ts=%s status=%d tags=%v opt=%s",
		str, num, flag, fl, ts.UTC().Format(time.RFC3339), status, tags, optStr), nil
}

func (s *TestRESTServer) GetChild(ctx context.Context, id uint64, childId uint64) (string, error) {
	return fmt.Sprintf("/items/%d/children/%d", id, childId), nil
}

func (s *TestRESTServer) ResetItems(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = map[uint64]*Item{}
	return nil
}

func (s *TestRESTServer) RpcMethod(ctx context.Context, id uint64) (*Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[id]
	if !ok {
		return nil, ErrUnexpectedValue.WithCausef("no item %d", id)
	}
	return item, nil
}
