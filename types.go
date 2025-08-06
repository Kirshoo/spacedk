package spacedk

import (
	"fmt"
	"time"
	"strings"
	"encoding/json"
)

// Currency is a shorthand for int64
// Similar to https://github.com/HOWZ1T/space_trader
type Currency int64

const dtLayout string = "2006-01-02"
type DateOnlyTime struct {
	time.Time
}

func (dt *DateOnlyTime) UnmarshalJSON(data []byte) error {
	timeString := strings.Trim(string(data), "\"")
	if timeString == "null" {
		dt.Time = time.Time{}
		return nil
	}

	tmp, err := time.Parse(dtLayout, timeString)
	if err != nil {
		return err
	}

	dt.Time = tmp
	return nil
}

func (dt DateOnlyTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", dt.Format(dtLayout))), nil
}

type SectorSymbol struct {
	Sector string
}

func (s *SectorSymbol) String() string {
	return s.Sector
}

// Represents a Symbol of a System
type SystemSymbol struct {
	Sector string
	System string
}

func (s *SystemSymbol) String() string {
	return strings.Join([]string{s.Sector, s.System}, "-")
}

func (s SystemSymbol) ToSectorSymbol() *SectorSymbol {
	return &SectorSymbol{Sector: s.Sector}
}

type WaypointSymbol struct {
	Sector string
	System string
	Waypoint string
}

func (s *WaypointSymbol) String() string {
	return strings.Join([]string{s.Sector, s.System, s.Waypoint}, "-")
}

func (s WaypointSymbol) ToSystemSymbol() *SystemSymbol {
	return &SystemSymbol{Sector: s.Sector, System: s.System}
}

func marshalDeliminerStrings(parts ...string) ([]byte, error) {
	return json.Marshal(strings.Join(parts, "-"))
}

func unmarshalDeliminerStrings(data []byte, count int) ([]string, error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	tokens := strings.Split(s, "-")
	if len(tokens) != count {
		return nil, fmt.Errorf("expected %d components, got %d", count, len(tokens))
	}

	result := make([]string, count)
	for i, t := range tokens {
		result[i] = t
	}

	return result, nil
}

func (s *SectorSymbol) MarshalJSON() ([]byte, error) {
	return marshalDeliminerStrings(s.Sector)
}

func (s *SectorSymbol) UnmarshalJSON(data []byte) error {
	parts, err := unmarshalDeliminerStrings(data, 1)
	if err != nil {
		return err
	}

	s.Sector = parts[0]
	return nil
}

func (s *SystemSymbol) MarshalJSON() ([]byte, error) {
	return marshalDeliminerStrings(s.Sector, s.System)
}

func (s *SystemSymbol) UnmarshalJSON(data []byte) error {
	parts, err := unmarshalDeliminerStrings(data, 2)
	if err != nil {
		return err
	}

	s.Sector = parts[0]
	s.System = parts[1]
	return nil
}

func (s *WaypointSymbol) MarshalJSON() ([]byte, error) {
	return marshalDeliminerStrings(s.Sector, s.System, s.Waypoint)
}

func (s *WaypointSymbol) UnmarshalJSON(data []byte) error {
	parts, err := unmarshalDeliminerStrings(data, 3)
	if err != nil {
		return err
	}

	s.Sector = parts[0]
	s.System = parts[1]
	s.Waypoint = parts[2]
	return nil
}
