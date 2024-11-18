package mangadex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMangaVolumes(t *testing.T) {
	mangaVolumes := MangaAggregate{
		Volumes: []Volume{
			{
				Volume: "2",
				Chapters: []Chapter{
					{Chapter: "0.04"},
					{Chapter: "0.05"},
					{Chapter: "0.06"},
				},
			},
			{
				Volume: "1",
				Chapters: []Chapter{
					{Chapter: "0.02"},
					{Chapter: "0.01"},
					{Chapter: "0.03"},
				},
			},
			{
				Volume: "3",
				Chapters: []Chapter{
					{Chapter: "0.10"},
					{Chapter: "0.09"},
					{Chapter: "0.11"},
				},
			},
		},
	}

	expected := MangaAggregate{
		Volumes: []Volume{
			{
				Volume: "1",
				Chapters: []Chapter{
					{Chapter: "0.01"},
					{Chapter: "0.02"},
					{Chapter: "0.03"},
				},
			},
			{
				Volume: "2",
				Chapters: []Chapter{
					{Chapter: "0.04"},
					{Chapter: "0.05"},
					{Chapter: "0.06"},
				},
			},
			{
				Volume: "3",
				Chapters: []Chapter{
					{Chapter: "0.09"},
					{Chapter: "0.10"},
					{Chapter: "0.11"},
				},
			},
		},
	}

	sortMangaVolumes(mangaVolumes)

	assert.Equal(t, expected, mangaVolumes)
}