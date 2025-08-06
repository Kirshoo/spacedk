package spacedk

import (
	"testing"
	"strings"
	"encoding/json"
)

func TestDateOnlyUnmarshal(t *testing.T) {
	var date DateOnlyTime
	
	data := []byte("\"2025-07-27\"")
	if err := json.Unmarshal(data, &date); err != nil {
		t.Errorf("Unmarshal error: %v", err)
		return
	}

	t.Logf("Unmarshaled date: %v", data)
}

func TestPoiSymbolType(t *testing.T) {
	testSector := "X2"
	testSystem := "4P12"
	testWaypoint := "7H"
	SymbolArray := []string{testSector, testSystem, testWaypoint}

	var waypoint WaypointSymbol

	testData := []byte("\"" + strings.Join(SymbolArray, "-") + "\"")
	if err := waypoint.UnmarshalJSON(testData); err != nil {
		t.Errorf("Unmarshal error: %v", err)
		return
	}

	if waypoint.String() != strings.Trim(string(testData), "\"") {
		t.Errorf("Error: invalid full (have '%s', need '%s')",
			waypoint.String(), strings.Trim(string(testData), "\""))
	}
}
