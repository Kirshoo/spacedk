package spacedk

import (
	"testing"
	"strings"
	"encoding/json"
)

func TestDateOnlyUnmarshal(t *testing.T) {
	var date DateOnlyTime
	
	data := []byte("2025-07-27")
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

	var poi PoiSymbol

	testData := []byte(strings.Join(SymbolArray, "-"))
	if err := json.Unmarshal(data, &poi); err != nil {
		t.Errorf("Unmarshal error: %v", err)
		return
	}

	if poi.Full != string(testData) {
		t.Errorf("Error: invalid full (have '%s', need '%s')",
			poi.Full, string(testData))
	}

	if poi.Sector != testSector {
		t.Errorf("Error: unexpected sector: want '%s', got 's'", 
			testSector, poi.Sector)
	}

	if poi.System != testSystem {
		t.Errorf("Error: unexpected system: want '%s', got 's'", 
			testSystem, poi.System)
	}

	if poi.Waypoint != testSector {
		t.Errorf("Error: unexpected waypoint: want '%s', got 's'", 
			testWaypoint, poi.Waypoint)
	}
}
