package waiops

import (
	"encoding/json"
	"testing"

	"github.com/dsnet/try"
)

func TestEvResourceMarshalJSONOmitsEmptyStrings(t *testing.T) {
	res := EvResource{
		Name:      "node1",
		Hostname:  "",
		Interface: "GigabitEthernet0/0/0",
		Port:      0,

		Extras: map[string]any{
			"extraK1":     "value1",
			"extraEmpty":  "",
			"extraNumber": 7,
		},
	}

	content := try.E1(json.Marshal(res))

	var payload map[string]any
	if err := json.Unmarshal(content, &payload); err != nil {
		t.Fatal(err)
	}

	if payload["name"] != "node1" {
		t.Fatalf("expected name to be present, got %v", payload["name"])
	}
	if _, ok := payload["hostname"]; ok {
		t.Fatalf("expected hostname to be omitted, got %v", payload["hostname"])
	}
	if payload["interface"] != "GigabitEthernet0/0/0" {
		t.Fatalf("expected interface to be present, got %v", payload["interface"])
	}
	if _, ok := payload["extraEmpty"]; ok {
		t.Fatalf("expected empty extra field to be omitted, got %v", payload["extraEmpty"])
	}
	if payload["extraK1"] != "value1" {
		t.Fatalf("expected non-empty extra string to be present, got %v", payload["extraK1"])
	}
	if payload["extraNumber"] != float64(7) {
		t.Fatalf("expected non-string extra to be present, got %v", payload["extraNumber"])
	}
	if payload["port"] != float64(0) {
		t.Fatalf("expected port zero value to remain present, got %v", payload["port"])
	}
}

func TestEvResourceUnmarshalJSONMissingFieldsUseZeroValues(t *testing.T) {
	content := []byte(`{"name":"node1","extraK1":"value1"}`)

	var newres EvResource
	if err := json.Unmarshal(content, &newres); err != nil {
		t.Fatal(err)
	}

	if newres.Name != "node1" {
		t.Fatalf("expected name to be unmarshaled, got %q", newres.Name)
	}
	if newres.Hostname != "" {
		t.Fatalf("expected missing hostname to stay empty, got %q", newres.Hostname)
	}
	if newres.Interface != "" {
		t.Fatalf("expected missing interface to stay empty, got %q", newres.Interface)
	}
	if newres.Extras["extraK1"] != "value1" {
		t.Fatalf("expected extra field to be preserved, got %v", newres.Extras["extraK1"])
	}
}
