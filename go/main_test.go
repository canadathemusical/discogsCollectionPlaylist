package main

import (
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		str      string
		expected bool
	}{
		{"found", []string{"a", "b", "c"}, "b", true},
		{"not found", []string{"a", "b", "c"}, "d", false},
		{"empty slice", []string{}, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := contains(tt.slice, tt.str)
			if result != tt.expected {
				t.Errorf("contains(%v, %q) = %v; want %v", tt.slice, tt.str, result, tt.expected)
			}
		})
	}
}

func TestHandleGenreFilter(t *testing.T) {
	releases := []iRelease{
		{
			BasicInformation: iBasicInformation{
				Title:  "Album 1",
				Genres: []string{"Rock"},
				Styles: []string{"Doom Metal"},
			},
		},
		{
			BasicInformation: iBasicInformation{
				Title:  "Album 2",
				Genres: []string{"Electronic"},
				Styles: []string{"Techno"},
			},
		},
		{
			BasicInformation: iBasicInformation{
				Title:  "Album 3",
				Genres: []string{"Rock"},
				Styles: []string{"Punk"},
			},
		},
	}

	tests := []struct {
		name          string
		genre         string
		expectedCount int
	}{
		{"filter by rock", "rock", 2},
		{"filter by metal", "metal", 1},
		{"filter by electronic", "electronic", 1},
		{"no filter", "", 3},
		{"not found", "jazz", 3}, // returns all when not found
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handleGenreFilter(releases, tt.genre)
			if len(result) != tt.expectedCount {
				t.Errorf("handleGenreFilter with genre %q returned %d releases; want %d",
					tt.genre, len(result), tt.expectedCount)
			}
		})
	}
}

func TestHandleArtistsFilter(t *testing.T) {
	releases := []iRelease{
		{
			BasicInformation: iBasicInformation{
				Title: "Album 1",
				Artists: []iArtist{
					{Name: "Metallica"},
				},
			},
		},
		{
			BasicInformation: iBasicInformation{
				Title: "Album 2",
				Artists: []iArtist{
					{Name: "Wu-Tang Clan"},
				},
			},
		},
		{
			BasicInformation: iBasicInformation{
				Title: "Album 3",
				Artists: []iArtist{
					{Name: "Wu-Tang Clan"},
					{Name: "Method Man"},
				},
			},
		},
	}

	tests := []struct {
		name          string
		artist        string
		expectedCount int
	}{
		{"filter by metallica", "metallica", 1},
		{"filter by wu-tang", "wu-tang", 2},
		{"filter by method", "method", 1},
		{"no filter", "", 3},
		{"not found", "beatles", 3}, // returns all when not found
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handleArtistsFilter(releases, tt.artist)
			if len(result) != tt.expectedCount {
				t.Errorf("handleArtistsFilter with artist %q returned %d releases; want %d",
					tt.artist, len(result), tt.expectedCount)
			}
		})
	}
}

func TestHandleRecentFilter(t *testing.T) {
	now := time.Now()
	
	releases := []iRelease{
		{
			DateAdded: now.AddDate(0, 0, -5).Format(time.RFC3339), // 5 days ago
			BasicInformation: iBasicInformation{
				Title: "Recent Album 1",
			},
		},
		{
			DateAdded: now.AddDate(0, 0, -20).Format(time.RFC3339), // 20 days ago
			BasicInformation: iBasicInformation{
				Title: "Recent Album 2",
			},
		},
		{
			DateAdded: now.AddDate(0, 0, -50).Format(time.RFC3339), // 50 days ago
			BasicInformation: iBasicInformation{
				Title: "Old Album",
			},
		},
	}

	tests := []struct {
		name          string
		days          int
		expectedCount int
	}{
		{"filter last 7 days", 7, 1},
		{"filter last 30 days", 30, 2},
		{"filter last 60 days", 60, 3},
		{"no filter (0 days)", 0, 3},
		{"no filter (negative)", -1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handleRecentFilter(releases, tt.days)
			if len(result) != tt.expectedCount {
				t.Errorf("handleRecentFilter with %d days returned %d releases; want %d",
					tt.days, len(result), tt.expectedCount)
			}
		})
	}
}
