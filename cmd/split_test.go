package cmd

import (
	"testing"
	"testing/fstest"
)

func TestSplitFile(t *testing.T) {
	const (
		file_content = `line 1
line 2
line 3`
	)

	file_dir := fstest.MapFS{"myfile.csv": {Data: []byte(file_content)}}

	err := SplitFile(file_dir, "myfile.csv", 2)

	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if len(file_dir) == 1 {
		t.Error("file not split")
	}

	if len(file_dir) > 2 {
		t.Error("file split too many times")
	}

}
