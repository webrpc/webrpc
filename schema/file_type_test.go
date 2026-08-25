package schema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseFileTestSchema(t *testing.T, types string, methods string) (*WebRPCSchema, error) {
	t.Helper()
	input := fmt.Sprintf(`{
		"webrpc": "v1",
		"name": "example",
		"version": "v0.0.1",
		"types": [%s],
		"services": [
			{
				"name": "Example",
				"methods": [%s]
			}
		]
	}`, types, methods)
	return ParseSchemaJSON([]byte(input))
}

func TestFileTypeValidPlacements(t *testing.T) {
	// Single file input argument.
	_, err := parseFileTestSchema(t, ``, `
		{
			"name": "UploadAvatar",
			"inputs": [{"name": "avatar", "type": "file"}],
			"outputs": [{"name": "ok", "type": "bool"}]
		}
	`)
	require.NoError(t, err)

	// []file input argument.
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "UploadAttachments",
			"inputs": [{"name": "attachments", "type": "[]file"}],
			"outputs": []
		}
	`)
	require.NoError(t, err)

	// Struct with top-level file fields used as a method input.
	_, err = parseFileTestSchema(t, `
		{
			"kind": "struct",
			"name": "UploadReportRequest",
			"fields": [
				{"name": "title", "type": "string"},
				{"name": "document", "type": "file"},
				{"name": "attachments", "type": "[]file"}
			]
		}
	`, `
		{
			"name": "UploadReport",
			"inputs": [{"name": "req", "type": "UploadReportRequest"}],
			"outputs": [{"name": "id", "type": "uint64"}]
		}
	`)
	require.NoError(t, err)

	// Single file output (download method).
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "DownloadReport",
			"inputs": [{"name": "id", "type": "uint64"}],
			"outputs": [{"name": "report", "type": "file"}]
		}
	`)
	require.NoError(t, err)

	// File output with metadata outputs alongside it.
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "DownloadReportWithMeta",
			"inputs": [{"name": "id", "type": "uint64"}],
			"outputs": [
				{"name": "report", "type": "file"},
				{"name": "size", "type": "uint64"},
				{"name": "etag", "type": "string"}
			]
		}
	`)
	require.NoError(t, err)
}

func TestFileTypeForbiddenPlacements(t *testing.T) {
	// []file as method output.
	_, err := parseFileTestSchema(t, ``, `
		{
			"name": "DownloadAll",
			"inputs": [],
			"outputs": [{"name": "files", "type": "[]file"}]
		}
	`)
	assert.ErrorContains(t, err, "DownloadAll")
	assert.ErrorContains(t, err, "only a single 'file' output is allowed")

	// More than one file output.
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "DownloadTwo",
			"inputs": [],
			"outputs": [
				{"name": "report", "type": "file"},
				{"name": "thumbnail", "type": "file"}
			]
		}
	`)
	assert.ErrorContains(t, err, "DownloadTwo")
	assert.ErrorContains(t, err, "more than one file output")

	// file-carrying struct used as a method output.
	_, err = parseFileTestSchema(t, `
		{
			"kind": "struct",
			"name": "Payload",
			"fields": [{"name": "data", "type": "file"}]
		}
	`, `
		{
			"name": "GetPayload",
			"inputs": [],
			"outputs": [{"name": "payload", "type": "Payload"}]
		}
	`)
	assert.ErrorContains(t, err, "GetPayload")

	// file-carrying struct nested as a field of another struct.
	_, err = parseFileTestSchema(t, `
		{
			"kind": "struct",
			"name": "Payload",
			"fields": [{"name": "data", "type": "file"}]
		},
		{
			"kind": "struct",
			"name": "Wrapper",
			"fields": [{"name": "payload", "type": "Payload"}]
		}
	`, `
		{
			"name": "Upload",
			"inputs": [{"name": "req", "type": "Wrapper"}],
			"outputs": []
		}
	`)
	assert.ErrorContains(t, err, "field 'payload' of struct 'Wrapper'")

	// file inside a map.
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "UploadMap",
			"inputs": [{"name": "files", "type": "map<string,file>"}],
			"outputs": []
		}
	`)
	assert.ErrorContains(t, err, "UploadMap")
	assert.ErrorContains(t, err, "invalid file placement")

	// file nested deeper than a single list level.
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "UploadNested",
			"inputs": [{"name": "files", "type": "[][]file"}],
			"outputs": []
		}
	`)
	assert.ErrorContains(t, err, "UploadNested")
	assert.ErrorContains(t, err, "invalid file placement")

	// list of file-carrying structs as input.
	_, err = parseFileTestSchema(t, `
		{
			"kind": "struct",
			"name": "Payload",
			"fields": [{"name": "data", "type": "file"}]
		}
	`, `
		{
			"name": "UploadMany",
			"inputs": [{"name": "payloads", "type": "[]Payload"}],
			"outputs": []
		}
	`)
	assert.ErrorContains(t, err, "UploadMany")
	assert.ErrorContains(t, err, "invalid file placement")

	// alias over file.
	_, err = parseFileTestSchema(t, `
		{
			"kind": "alias",
			"name": "Upload",
			"type": "file"
		}
	`, `
		{
			"name": "Ping",
			"inputs": [],
			"outputs": []
		}
	`)
	assert.ErrorContains(t, err, "alias 'Upload'")

	// streaming methods cannot use files.
	_, err = parseFileTestSchema(t, ``, `
		{
			"name": "StreamUpload",
			"streamOutput": true,
			"inputs": [{"name": "data", "type": "file"}],
			"outputs": [{"name": "ok", "type": "bool"}]
		}
	`)
	assert.ErrorContains(t, err, "StreamUpload")
	assert.ErrorContains(t, err, "streaming")
}

func TestFileDownloadMetaHelpers(t *testing.T) {
	s, err := parseFileTestSchema(t, ``, `
		{
			"name": "DownloadReport",
			"inputs": [{"name": "id", "type": "uint64"}],
			"outputs": [
				{"name": "size", "type": "uint64"},
				{"name": "report", "type": "file"},
				{"name": "etag", "type": "string"}
			]
		},
		{
			"name": "DownloadPlain",
			"inputs": [],
			"outputs": [{"name": "report", "type": "file"}]
		},
		{
			"name": "NoFile",
			"inputs": [],
			"outputs": [{"name": "ok", "type": "bool"}]
		}
	`)
	require.NoError(t, err)

	methods := map[string]*Method{}
	for _, m := range s.Services[0].Methods {
		methods[m.Name] = m
	}

	// The file output is found wherever it is declared, and the remaining
	// outputs are returned as metadata in declaration order.
	withMeta := methods["DownloadReport"]
	assert.True(t, MethodHasFileDownload(withMeta))
	assert.True(t, MethodHasFileDownloadMeta(withMeta))
	assert.Equal(t, 1, MethodDownloadFileIndex(withMeta))
	meta := MethodDownloadMetaOutputs(withMeta)
	require.Len(t, meta, 2)
	assert.Equal(t, "size", meta[0].Name)
	assert.Equal(t, "etag", meta[1].Name)

	// A lone file output carries no metadata.
	plain := methods["DownloadPlain"]
	assert.True(t, MethodHasFileDownload(plain))
	assert.False(t, MethodHasFileDownloadMeta(plain))
	assert.Equal(t, 0, MethodDownloadFileIndex(plain))
	assert.Empty(t, MethodDownloadMetaOutputs(plain))

	// A method without a file output is not a download method at all, so its
	// outputs are not metadata.
	noFile := methods["NoFile"]
	assert.False(t, MethodHasFileDownload(noFile))
	assert.False(t, MethodHasFileDownloadMeta(noFile))
	assert.Equal(t, -1, MethodDownloadFileIndex(noFile))
	assert.Empty(t, MethodDownloadMetaOutputs(noFile))

	// Nil methods are handled without panicking.
	assert.False(t, MethodHasFileDownload(nil))
	assert.Equal(t, -1, MethodDownloadFileIndex(nil))
	assert.Empty(t, MethodDownloadMetaOutputs(nil))
}
