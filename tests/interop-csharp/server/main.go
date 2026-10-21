// Command server serves the CSharp service (../csharp.ridl) for the C#
// client interoperability tests.
//
// Usage: go run ./tests/interop-csharp/server -port=9990
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
)

func main() {
	port := flag.Int("port", 9990, "listen port")
	flag.Parse()

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	log.Printf("serving CSharp service at http://%s", addr)
	mux := http.NewServeMux()
	mux.Handle("/", NewCSharpServer(&server{}))
	// A proxy-like error page, not a webrpc error response.
	mux.HandleFunc("/html-error/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<html>502 Bad Gateway</html>"))
	})
	log.Fatal(http.ListenAndServe(addr, mux))
}

type server struct{}

func bigInt(s string) BigInt {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic(s)
	}
	return BigInt(*i)
}

func ptr[T any](v T) *T { return &v }

// fixtureTypes is shared with the client test expectations (Suites.cs).
func fixtureTypes() *Types {
	return &Types{
		B:        true,
		U8:       255,
		U16:      65535,
		U32:      4294967295,
		U64:      18446744073709551615,
		I8:       -128,
		I16:      -32768,
		I32:      -2147483648,
		I64:      -9223372036854775808,
		F32:      1.5,
		F64:      -0.000123456789,
		Str:      "héllo \"wörld\" <&> ☃ \U0001F600",
		Bytes:    []byte{0x00, 0x01, 0xfe, 0xff},
		Ts:       time.Date(2026, 9, 3, 12, 30, 15, 123456700, time.UTC),
		TsList:   []time.Time{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)},
		Big:      bigInt("-123456789012345678901234567890"),
		BigList:  []BigInt{bigInt("0"), bigInt("340282366920938463463374607431768211456")},
		BigMap:   map[string]BigInt{"wei": bigInt("1000000000000000000")},
		Address:  "0xabc",
		Wallet:   "0xdef",
		Labels:   Labels{"a", "b"},
		Kind:     KindAdmin,
		Kinds:    map[Kind]uint32{KindUser: 1, KindAdmin: 2},
		Country:  CountryCa,
		Reserved: ReservedClass,
		ById:     map[int64]string{-1: "minus one", 42: "answer"},
		Nested:   map[uint32][]string{7: {"x", "y"}},
		Anything: map[string]any{"n": float64(1), "s": "two", "l": []any{true, nil}},
		OptStr:   nil,
		OptBig:   ptr(bigInt("42")),
		OptTs:    ptr(time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)),
		OptKind:  ptr(KindUser),
		OptList:  nil,
		Renamed:  "renamed",
	}
}

var bigIntComparer = cmp.Comparer(func(a, b BigInt) bool { return a.AsInt().Cmp(b.AsInt()) == 0 })

func (s *server) GetTypes(ctx context.Context) (*Types, error) {
	return fixtureTypes(), nil
}

func (s *server) SendTypes(ctx context.Context, types *Types) error {
	if diff := cmp.Diff(fixtureTypes(), types, bigIntComparer); diff != "" {
		return ErrMismatch.WithCausef("types:\n%s", diff)
	}
	return nil
}

func (s *server) GetUnknownEnum(ctx context.Context) (Color, []Color, map[Color]uint32, error) {
	return Color("purple"), []Color{ColorRed, Color("blue")}, map[Color]uint32{ColorGreen: 1, Color("pink"): 2}, nil
}

func (s *server) SendUnknownEnum(ctx context.Context, color Color, colors []Color) error {
	if color != "purple" {
		return ErrMismatch.WithCausef("color: %q", color)
	}
	if !cmp.Equal(colors, []Color{ColorRed, Color("blue")}) {
		return ErrMismatch.WithCausef("colors: %q", colors)
	}
	return nil
}

func (s *server) Keywords(ctx context.Context, class string, cancellationToken string, params []string) (string, *Keywords, error) {
	event := fmt.Sprintf("class=%s cancellationToken=%s params=%s", class, cancellationToken, strings.Join(params, ","))
	return event, &Keywords{Class: "class", String: "string", Event: "event", Keywords: "keywords", ToString: "toString", Equals: "equals"}, nil
}

func (s *server) Optional(ctx context.Context, str *string, num *uint64, ts *time.Time) (*string, *uint64, *time.Time, error) {
	return str, num, ts, nil
}

func (s *server) Deprecated(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *server) GetShadowed(ctx context.Context) (*Task, *Version, *List, error) {
	return &Task{Id: 1}, &Version{Major: 2}, &List{Items: []string{"three"}}, nil
}

func (s *server) Countdown(ctx context.Context, from uint32, stream CountdownStreamWriter) error {
	at := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for n := from; n > 0; n-- {
		if err := stream.Write(n, at.Add(time.Duration(n)*time.Second)); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) Watch(ctx context.Context, req WatchRequest, stream WatchStreamWriter) error {
	for seq := uint32(1); seq <= req.Count; seq++ {
		if err := stream.Write(&WatchResponse{Key: req.Key, Seq: seq}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) CountdownFail(ctx context.Context, from uint32, stream CountdownFailStreamWriter) error {
	if err := stream.Write(from); err != nil {
		return err
	}
	return ErrCountdownFailed.WithCausef("failed after %d", from)
}

func (s *server) GetByName(ctx context.Context, name string, suffix *string) (string, *string, error) {
	return name, suffix, nil
}

func (s *server) Headers(ctx context.Context) (string, string, error) {
	r, ok := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	if !ok {
		return "", "", ErrMismatch.WithCausef("no request in context")
	}
	return r.Header.Get("Webrpc"), r.Header.Get("X-Custom"), nil
}

func (s *server) UploadDoc(ctx context.Context, doc *Doc) (string, string, uint64, string, error) {
	if doc == nil || doc.File == nil {
		return "", "", 0, "", ErrMismatch.WithCausef("doc.file is missing")
	}
	defer doc.File.Body.Close()
	content, err := io.ReadAll(doc.File.Body)
	if err != nil {
		return "", "", 0, "", err
	}
	return doc.File.Name, doc.File.ContentType, uint64(len(content)), doc.Title, nil
}

func (s *server) Attach(ctx context.Context, req AttachRequest) (*AttachResponse, error) {
	resp := &AttachResponse{Id: req.Id}
	for _, file := range req.Files {
		content, err := io.ReadAll(file.Body)
		file.Body.Close()
		if err != nil {
			return nil, err
		}
		resp.Names = append(resp.Names, file.Name)
		resp.Sizes = append(resp.Sizes, uint64(len(content)))
	}
	return resp, nil
}

func (s *server) DownloadDoc(ctx context.Context, name string) (*WebrpcFile, error) {
	content := []byte("doc:" + name)
	return &WebrpcFile{
		Name:        name,
		ContentType: "text/plain; charset=utf-8",
		Size:        int64(len(content)),
		Body:        io.NopCloser(strings.NewReader(string(content))),
	}, nil
}
