package main

import (
	"testing"
)

func TestGetAlbum(t *testing.T) {
	tests := []struct {
		name     string
		release  iRelease
		expected string
	}{
		{
			name: "simple title",
			release: iRelease{
				BasicInformation: iBasicInformation{
					Title: "Sorrow And Extinction",
				},
			},
			expected: "Sorrow And Extinction",
		},
		{
			name: "title with special characters",
			release: iRelease{
				BasicInformation: iBasicInformation{
					Title: "Enter The Wu-Tang (36 Chambers)",
				},
			},
			expected: "Enter The Wu-Tang (36 Chambers)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getAlbum(tt.release)
			if result != tt.expected {
				t.Errorf("getAlbum() = %q; want %q", result, tt.expected)
			}
		})
	}
}

func TestGetArtist(t *testing.T) {
	tests := []struct {
		name     string
		release  iRelease
		expected string
	}{
		{
			name: "single artist",
			release: iRelease{
				BasicInformation: iBasicInformation{
					Artists: []iArtist{
						{Name: "Pallbearer"},
					},
				},
			},
			expected: "Pallbearer",
		},
		{
			name: "artist with number suffix",
			release: iRelease{
				BasicInformation: iBasicInformation{
					Artists: []iArtist{
						{Name: "Metallica (2)"},
					},
				},
			},
			expected: "Metallica",
		},
		{
			name: "multiple artists",
			release: iRelease{
				BasicInformation: iBasicInformation{
					Artists: []iArtist{
						{Name: "Artist One"},
						{Name: "Artist Two"},
					},
				},
			},
			expected: "Artist One & Artist Two",
		},
		{
			name: "multiple artists with numbers",
			release: iRelease{
				BasicInformation: iBasicInformation{
					Artists: []iArtist{
						{Name: "Wu-Tang Clan (3)"},
						{Name: "Method Man (2)"},
					},
				},
			},
			expected: "Wu-Tang Clan & Method Man",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getArtist(tt.release)
			if result != tt.expected {
				t.Errorf("getArtist() = %q; want %q", result, tt.expected)
			}
		})
	}
}
