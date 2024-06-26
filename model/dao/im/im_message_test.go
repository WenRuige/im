package im

import (
	"github.com/bytedance/sonic"
	"github.com/im/bootstrap"
	"testing"
)

func TestGetMessageList(t *testing.T) {
	t.Run("", func(t *testing.T) {
		bootstrap.InitMySQL()
		res, err := GetMessageList(nil, nil)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(sonic.MarshalString(res))
	})
}
