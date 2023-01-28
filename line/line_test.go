package line

import (
	"testing"
)

func TestIsLimitApiReturnError(t *testing.T) {
	if ERROR_INT64 > ERROR_SETTING_INT64 {
		t.Log("checking limit")
	} else {
		t.Error("no checking limit")
	}

	t.Log("TestIsLimitApiReturnError終了")
}
