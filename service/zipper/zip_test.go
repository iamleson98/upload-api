package zipper

import (
	"testing"
)

func TestUnZip(t *testing.T) {
	err := Unzip("../../media/upload-951502072.zip")
	if err != nil {
		t.Error(err)
	}
}
