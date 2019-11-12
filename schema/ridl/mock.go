package ridl

var mockImport bool

func enableMockImport() {
	mockImport = true
}

func disableMockImport() {
	mockImport = false
}
