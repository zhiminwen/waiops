package waiops

import (
	"testing"

	"encoding/json"

	"github.com/dsnet/try"
)

func TestEvResourceRandomField(t *testing.T) {
	res := EvResource{
		Name:      "node1",
		Hostname:  "node1",
		Interface: "GigabitEthernet0/0/0",

		Extras: make(map[string]any),
	}
	res.Extras["extraK1"] = "value1"
	res.Extras["extraK2"] = "value2"

	// t.Log(res)
	content := try.E1(json.MarshalIndent(res, "", "  "))
	t.Log(string(content))

	newres := EvResource{}
	err := json.Unmarshal(content, &newres)
	if err != nil {
		t.Error(err)
	}
	t.Log(newres)
	t.Log(newres.Extras)

}
