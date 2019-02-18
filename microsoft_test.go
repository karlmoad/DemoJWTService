package DemoJWTService

import (
	"DemoJWTService/common"
	"testing"
)

func TestNewMicrosoftKeySource(t *testing.T) {
	msKS, err := common.NewMicrosoftKeySource()
	if err != nil {
		t.Error(err)
	}
}

func TestMicrosoftKeySource_GetKeys(t *testing.T) {
	msKS, err := common.NewMicrosoftKeySource()
	if err != nil {
		t.Error(err)
	}

	keys, err := msKS.GetKeys("-sxMJMLCIDWMTPvZyJ6tx-CDxw0")
	if err != nil {
		t.Error(err)
	}

	if len(keys) == 0 {
		t.Errorf("keys invalid: %s", keys)
	}

}
