package models

import (
	"fmt"
	"time"
)

type Salutes struct {
	Data []string `json:"salutes"`
}

type Trivia struct {
	Data []string `json:"trivia"`
}

type Stage struct {
	Count     int    `json:"count"`
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	Anomaly   string `json:"anomaly"`
	Warning   string `json:"warning"`
}

type DeepDive struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Biome  string  `json:"biome"`
	Seed   int64   `json:"seed"`
	Stages []Stage `json:"stages"`
}

func (d DeepDive) String() string {
	return fmt.Sprintf("%s %s %s", d.Type, d.Name, d.Biome)
}

type DeepDives struct {
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	Variants  []DeepDive `json:"variants"`
}

type DRGData interface {
	Salutes | Trivia | DeepDives
}
