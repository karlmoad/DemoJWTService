package DemoJWTService

import (
	"DemoJWTService/common"
	"testing"
)

func TestNewMicrosoftKeySource(t *testing.T) {
	_, err := common.NewMicrosoftKeySource()
	if err != nil {
		t.Error(err)
	}
}

func TestMicrosoftKeySource_GetKeys(t *testing.T) {
	msKS, err := common.NewMicrosoftKeySource()
	if err != nil {
		t.Error(err)
	}

	kids := msKS.GetKidEntries()
	if len(kids) == 0{
		t.Errorf("No KID entires found")
	}

	keys, err := msKS.GetKeys(kids[0])
	if err != nil {
		t.Error(err)
	}

	if len(keys) == 0 {
		t.Errorf("keys invalid: %s", keys)
	}

}
