package assets

import "embed"

// FS contains the optimized card and mission images used by the app.
//
//go:embed cards/*/*.jpg cards/*/*.jpeg missions/*/*.jpg
var FS embed.FS
