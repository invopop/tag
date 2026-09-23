package tag

import "testing"

func TestVersionForBranch(t *testing.T) {
	const ts = "v20260619T1530"

	tests := []struct {
		name   string
		branch string
		want   string
	}{
		{name: "main is bare timestamp", branch: "main", want: ts},
		{name: "empty branch is bare timestamp", branch: "", want: ts},
		{name: "simple branch", branch: "fix-thing", want: "fix-thing-" + ts},
		{name: "slash collapsed to hyphen", branch: "feat/foo", want: "feat-foo-" + ts},
		{name: "multiple slashes", branch: "feat/foo/bar", want: "feat-foo-bar-" + ts},
		{name: "preserves dots and hyphens", branch: "chore/vuln-bumps-2026-06", want: "chore-vuln-bumps-2026-06-" + ts},
		{name: "collapses runs and trims", branch: "weird@@name//x", want: "weird-name-x-" + ts},
		{name: "all-unsafe branch falls back to timestamp", branch: "///", want: ts},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := versionForBranch(tt.branch, ts); got != tt.want {
				t.Errorf("versionForBranch(%q) = %q, want %q", tt.branch, got, tt.want)
			}
		})
	}
}
