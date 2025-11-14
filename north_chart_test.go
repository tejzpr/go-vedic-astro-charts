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
	"encoding/base64"
	"os"
	"testing"
)

func TestNorthChart_AllPlanets(t *testing.T) {
	// Test 1: All planets in different rashis
	input := ChartInput{
		ChartType: ChartTypeNorth,
		Lagna: &Planet{
			Rashi:        "aries",
			IsRetrograde: false,
			IsCombust:    false,
		},
		Planets: map[string]*Planet{
			"sun": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"moon": {
				Rashi:        "taurus",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mars": {
				Rashi:        "gemini",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mercury": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    true,
			},
			"jupiter": {
				Rashi:        "leo",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"venus": {
				Rashi:        "cancer",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"saturn": {
				Rashi:        "libra",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"rahu": {
				Rashi:        "scorpio",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"ketu": {
				Rashi:        "sagittarius",
				IsRetrograde: false,
				IsCombust:    false,
			},
		},
	}

	base64Image, err := GenerateChart(input)
	if err != nil {
		t.Fatalf("Error generating chart: %v", err)
	}

	if base64Image == "" {
		t.Fatal("Generated base64 image is empty")
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		t.Fatalf("Error decoding base64: %v", err)
	}

	err = os.WriteFile("test_north_all_planets.png", imageData, 0644)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	t.Logf("Test 1 passed: All planets chart generated successfully (%d bytes)", len(imageData))
}

func TestNorthChart_AllPlanetsWithLagna(t *testing.T) {
	// Test 2: All planets with Lagna in different rashi
	input := ChartInput{
		ChartType: ChartTypeNorth,
		Lagna: &Planet{
			Rashi:        "libra",
			IsRetrograde: false,
			IsCombust:    false,
		},
		Planets: map[string]*Planet{
			"sun": {
				Rashi:        "scorpio",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"moon": {
				Rashi:        "sagittarius",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mars": {
				Rashi:        "capricorn",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mercury": {
				Rashi:        "scorpio",
				IsRetrograde: false,
				IsCombust:    true,
			},
			"jupiter": {
				Rashi:        "pisces",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"venus": {
				Rashi:        "aquarius",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"saturn": {
				Rashi:        "taurus",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"rahu": {
				Rashi:        "gemini",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"ketu": {
				Rashi:        "cancer",
				IsRetrograde: false,
				IsCombust:    false,
			},
		},
	}

	base64Image, err := GenerateChart(input)
	if err != nil {
		t.Fatalf("Error generating chart: %v", err)
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		t.Fatalf("Error decoding base64: %v", err)
	}

	err = os.WriteFile("test_north_all_planets_with_lagna.png", imageData, 0644)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	t.Logf("Test 2 passed: All planets with Lagna chart generated successfully (%d bytes)", len(imageData))
}

func TestNorthChart_AllPlanetsWithUpagrahas(t *testing.T) {
	// Test 3: All planets with upagrahas
	input := ChartInput{
		ChartType: ChartTypeNorth,
		Lagna: &Planet{
			Rashi:        "aries",
			IsRetrograde: false,
			IsCombust:    false,
		},
		Planets: map[string]*Planet{
			// Regular planets
			"sun": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"moon": {
				Rashi:        "taurus",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mars": {
				Rashi:        "gemini",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mercury": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    true,
			},
			"jupiter": {
				Rashi:        "leo",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"venus": {
				Rashi:        "cancer",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"saturn": {
				Rashi:        "libra",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"rahu": {
				Rashi:        "scorpio",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"ketu": {
				Rashi:        "sagittarius",
				IsRetrograde: false,
				IsCombust:    false,
			},
			// Upagrahas
			"upaketu": {
				Rashi:        "capricorn",
				IsRetrograde: false,
				IsCombust:    false,
				IsUpagraha:   true,
			},
			"mandi": {
				Rashi:        "aquarius",
				IsRetrograde: false,
				IsCombust:    false,
				IsUpagraha:   true,
			},
			"gulika": {
				Rashi:        "pisces",
				IsRetrograde: false,
				IsCombust:    false,
				IsUpagraha:   true,
			},
		},
	}

	base64Image, err := GenerateChart(input)
	if err != nil {
		t.Fatalf("Error generating chart: %v", err)
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		t.Fatalf("Error decoding base64: %v", err)
	}

	err = os.WriteFile("test_north_all_planets_with_upagrahas.png", imageData, 0644)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	t.Logf("Test 3 passed: All planets with upagrahas chart generated successfully (%d bytes)", len(imageData))
}

func TestNorthChart_AllPlanetsUpagrahasLagnaSameRashi(t *testing.T) {
	// Test 4: All planets, upagrahas, and lagna in the same rashi (Aries)
	input := ChartInput{
		ChartType: ChartTypeNorth,
		Lagna: &Planet{
			Rashi:        "aries",
			IsRetrograde: false,
			IsCombust:    false,
		},
		Planets: map[string]*Planet{
			"sun": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"moon": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mars": {
				Rashi:        "aries",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"mercury": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    true,
			},
			"jupiter": {
				Rashi:        "aries",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"venus": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"saturn": {
				Rashi:        "aries",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"rahu": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"ketu": {
				Rashi:        "aries",
				IsRetrograde: false,
				IsCombust:    false,
			},
		},
	}

	base64Image, err := GenerateChart(input)
	if err != nil {
		t.Fatalf("Error generating chart: %v", err)
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		t.Fatalf("Error decoding base64: %v", err)
	}

	err = os.WriteFile("test_north_all_same_rashi.png", imageData, 0644)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	t.Logf("Test 4 passed: All planets, upagrahas, and lagna in same rashi chart generated successfully (%d bytes)", len(imageData))
}

func TestNorthChart_WithLagnaInLeo(t *testing.T) {
	// Test 5: Lagna in Leo (rashi 5) to test rashi number rotation
	input := ChartInput{
		ChartType: ChartTypeNorth,
		Lagna: &Planet{
			Rashi:        "leo",
			IsRetrograde: false,
			IsCombust:    false,
		},
		Planets: map[string]*Planet{
			"sun": {
				Rashi:        "taurus",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"jupiter": {
				Rashi:        "taurus",
				IsRetrograde: true,
				IsCombust:    false,
			},
			"moon": {
				Rashi:        "gemini",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mars": {
				Rashi:        "cancer",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"mercury": {
				Rashi:        "virgo",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"venus": {
				Rashi:        "libra",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"saturn": {
				Rashi:        "scorpio",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"rahu": {
				Rashi:        "sagittarius",
				IsRetrograde: false,
				IsCombust:    false,
			},
			"ketu": {
				Rashi:        "pisces",
				IsRetrograde: false,
				IsCombust:    false,
			},
		},
	}

	base64Image, err := GenerateChart(input)
	if err != nil {
		t.Fatalf("Error generating chart: %v", err)
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		t.Fatalf("Error decoding base64: %v", err)
	}

	err = os.WriteFile("test_north_lagna_leo.png", imageData, 0644)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	t.Logf("Test 5 passed: Lagna in Leo chart generated successfully (%d bytes)", len(imageData))
}
