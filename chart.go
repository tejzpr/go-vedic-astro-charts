// Copyright (c) 2024 Tejus Pratap <tejzpr@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package parashari

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"strings"
)

// ChartType represents the type of chart to generate
type ChartType string

const (
	ChartTypeNorth ChartType = "north"
	ChartTypeSouth ChartType = "south"
	ChartTypeEast  ChartType = "east"
	ChartTypeWest  ChartType = "west"
)

// Planet represents a planet in the chart
type Planet struct {
	Rashi        string `json:"rashi"`
	IsRetrograde bool   `json:"is_retrograde"`
	IsCombust    bool   `json:"is_combust"`
	IsUpagraha   bool   `json:"upagraha,omitempty"`
	Display      string `json:"display,omitempty"` // Custom display name
}

// ChartInput contains all the data needed to generate a chart
type ChartInput struct {
	ChartType  ChartType          `json:"chart_type"`
	Planets    map[string]*Planet `json:"planets"`
	Lagna      *Planet            `json:"lagna,omitempty"`
	CenterText string             `json:"center_text,omitempty"` // Text to display in center of chart
}

// RashiToNumber converts rashi name to number (1-12)
func RashiToNumber(rashi string) int {
	rashiMap := map[string]int{
		"aries":       1,
		"taurus":      2,
		"gemini":      3,
		"cancer":      4,
		"leo":         5,
		"virgo":       6,
		"libra":       7,
		"scorpio":     8,
		"sagittarius": 9,
		"capricorn":   10,
		"aquarius":    11,
		"pisces":      12,
	}
	r := strings.ToLower(rashi)
	if num, ok := rashiMap[r]; ok {
		return num
	}
	return 0
}

// NumberToRashi converts rashi number to name
func NumberToRashi(num int) string {
	rashiMap := map[int]string{
		1:  "aries",
		2:  "taurus",
		3:  "gemini",
		4:  "cancer",
		5:  "leo",
		6:  "virgo",
		7:  "libra",
		8:  "scorpio",
		9:  "sagittarius",
		10: "capricorn",
		11: "aquarius",
		12: "pisces",
	}
	return rashiMap[num]
}

// GetPlanetAbbreviation returns the abbreviation for a planet or upagraha
func GetPlanetAbbreviation(planetName string) string {
	abbrevMap := map[string]string{
		// Planets
		"sun":     "Su",
		"moon":    "Mo",
		"mars":    "Ma",
		"mercury": "Me",
		"jupiter": "Ju",
		"venus":   "Ve",
		"saturn":  "Sa",
		"rahu":    "Ra",
		"ketu":    "Ke",
		"lagna":   "Asc",
		// Upagrahas
		"upaketu":      "Up",
		"mandi":        "Mn",
		"gulika":       "Gu",
		"yamaghantaka": "Ya",
		"ardhaprahara": "Ar",
		"kala":         "Ka",
		"dhuma":        "Dh",
		"vyatipata":    "Vy",
		"parivesha":    "Pa",
		"indrachapa":   "In",
		"upagraha":     "Up", // Generic fallback
	}
	return abbrevMap[strings.ToLower(planetName)]
}

// GetPlanetDisplayName returns the display name for a planet
// If Display field is set, it uses that, otherwise uses the abbreviation
func GetPlanetDisplayName(planetName string, planet *Planet) string {
	if planet != nil && planet.Display != "" {
		return planet.Display
	}
	return GetPlanetAbbreviation(planetName)
}

// GenerateChart generates a chart image and returns it as a base64-encoded PNG string
func GenerateChart(input ChartInput) (string, error) {
	if input.ChartType == "" {
		return "", errors.New("chart_type is required")
	}

	var img []byte
	var err error

	switch input.ChartType {
	case ChartTypeSouth:
		img, err = GenerateSouthChart(input)
	case ChartTypeNorth:
		img, err = GenerateNorthChart(input)
	default:
		return "", fmt.Errorf("unsupported chart type: %s", input.ChartType)
	}

	if err != nil {
		return "", fmt.Errorf("failed to generate chart: %w", err)
	}

	// Encode to base64
	base64Str := base64.StdEncoding.EncodeToString(img)
	return base64Str, nil
}

// Helper function to encode image to PNG bytes
func encodePNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
