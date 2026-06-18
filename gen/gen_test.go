package gen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/schema"
)

func TestWebRPCGenHeaderSchemaVersion(t *testing.T) {
	tt := []struct {
		name          string
		schemaName    string
		schemaVersion string
		schemaHash    string
		wantSuffix    string
	}{
		{
			name:          "versioned schema ignores hash",
			schemaName:    "svc",
			schemaVersion: "v2.0.0",
			schemaHash:    "abc123",
			wantSuffix:    "svc@v2.0.0",
		},
		{
			name:       "unversioned schema falls back to hash",
			schemaName: "svc",
			schemaHash: "abc123",
			wantSuffix: "svc@v0.0.0-abc123",
		},
		{
			name:       "unversioned schema with omitted hash has no dangling dash",
			schemaName: "svc",
			schemaHash: "",
			wantSuffix: "svc@v0.0.0",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			header := webRPCGenHeader(TemplateVars{
				WebRPCSchema: &schema.WebRPCSchema{
					SchemaName:    tc.schemaName,
					SchemaVersion: tc.schemaVersion,
				},
				SchemaHash: tc.schemaHash,
			})

			// The schema-version segment is the third ";"-delimited field.
			fields := strings.Split(header, ";")
			require.Len(t, fields, 3, "header %q should have 3 fields", header)
			require.Equal(t, tc.wantSuffix, fields[2])
		})
	}
}
