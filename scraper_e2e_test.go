package scraper

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var isDebug bool

func init() {
	flag.BoolVar(&isDebug, "debug", false, "outputs some extra debugging information to STDOUT")
}

func getScraperFromFile(fileName string) (page *Scraper, err error) {
	fileHandle, err := os.Open(fmt.Sprintf("./test_assets/%v.html", fileName))
	if err != nil {
		return
	}
	defer func() { _ = fileHandle.Close() }()

	return NewFromBuffer(fileHandle)
}

func TestE2E_FindAll(t *testing.T) {
	type fields struct {
		uri     string
		filters Filter
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "example.com, p tags",
			fields: fields{
				uri:     "example.com",
				filters: Filter{Tag: "p"},
			},
			want: 2,
		},
		{
			name: "Wikipedia cats, top-level TOC items (partial class search with tags)",
			fields: fields{
				uri: "wikipedia.org_wiki_cat",
				filters: Filter{
					Tag:        "li",
					Attributes: Attributes{"class": "toclevel-1"},
				},
			},
			want: 12,
		},
		{
			name: "Wikipedia cats, number of languages (partial class search)",
			fields: fields{
				uri: "wikipedia.org_wiki_cat",
				filters: Filter{
					Attributes: Attributes{"class": "interlanguage-link"},
				},
			},
			want: 238,
		},
		{
			name: "Wikipedia cats, featured article badges (exact match search)",
			fields: fields{
				uri: "wikipedia.org_wiki_cat",
				filters: Filter{
					Attributes: Attributes{"title": "featured article badge"},
				},
			},
			want: 8,
		},
		{
			name: "Synthetic page, broken HTML, partial attribute match",
			fields: fields{
				uri: "synthetic",
				filters: Filter{
					Attributes: Attributes{"class": "beer"},
				},
			},
			want: 4,
		},
		//TODO: make this happen
		//{
		//	name: "Synthetic page, broken HTML, filter on attribute existence",
		//	fields: fields{
		//		uri: "synthetic",
		//		filters: Filter{
		//			Attributes: Attributes{"href": "*"},
		//		},
		//	},
		//	want: 2,
		//},
		{
			name: "France Passion, entry page",
			fields: fields{
				uri: "france-passion.com",
				filters: Filter{
					Tag:        "i",
					Attributes: Attributes{"class": "toolTip"},
				},
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var num int
			page, err := getScraperFromFile(tt.fields.uri)
			if err != nil {
				t.Fatal("Error while parsing page: ", err)
			}
			for element := range page.FindAll(tt.fields.filters) {
				if isDebug {
					log.Printf("%v with %v (%v)", element.Type(), element.Attributes(), element.TextOptimistic())
				}
				num++
			}
			if num != tt.want {
				t.Errorf("Matching elements: %v, want %v", num, tt.want)
			}
		})
	}
}

func TestE2E_FindAll_fluentInterface(t *testing.T) {
	type fields struct {
		uri          string
		step1Filters Filter
		step2Filters Filter
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "France Passion, entry page",
			fields: fields{
				uri: "france-passion.com",
				step1Filters: Filter{
					Tag:        "ul",
					Attributes: Attributes{"class": "carac"},
				},
				step2Filters: Filter{
					Tag:        "i",
					Attributes: Attributes{"class": "toolTip"},
				},
			},
			want: 9,
		},
		{
			name: "France Passion, entry page",
			fields: fields{
				uri: "france-passion.com",
				step1Filters: Filter{
					Tag:        "div",
					Attributes: Attributes{"class": "options"},
				},
				step2Filters: Filter{
					Tag:        "span",
					Attributes: Attributes{"class": "toolTip"},
				},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var num int
			page, err := getScraperFromFile(tt.fields.uri)
			if err != nil {
				t.Fatal("Error while parsing page: ", err)
			}
			for element := range page.FindAll(tt.fields.step1Filters) {
				for internalElement := range element.FindAll(tt.fields.step2Filters) {
					if isDebug {
						log.Printf("%v with %v (%v)", internalElement.Type(), internalElement.Attributes(), internalElement.TextOptimistic())
					}
					num++
				}
			}
			if num != tt.want {
				t.Errorf("Matching elements: %v, want %v", num, tt.want)
			}
		})
	}
}

func TestE2E_Find_FindAll_fluentInterface(t *testing.T) {
	type fields struct {
		uri          string
		step1Filters Filter
		step2Filters Filter
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "France Passion, entry page",
			fields: fields{
				uri: "france-passion.com",
				step1Filters: Filter{
					Tag:        "ul",
					Attributes: Attributes{"class": "carac"},
				},
				step2Filters: Filter{
					Tag:        "i",
					Attributes: Attributes{"class": "toolTip"},
				},
			},
			want: 9,
		},
		{
			name: "France Passion, entry page",
			fields: fields{
				uri: "france-passion.com",
				step1Filters: Filter{
					Tag:        "div",
					Attributes: Attributes{"class": "options"},
				},
				step2Filters: Filter{
					Tag:        "span",
					Attributes: Attributes{"class": "toolTip"},
				},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var num int
			page, err := getScraperFromFile(tt.fields.uri)
			if err != nil {
				t.Fatal("Error while parsing page: ", err)
			}
			element := page.Find(tt.fields.step1Filters)
			for internalElement := range element.FindAll(tt.fields.step2Filters) {
				if isDebug {
					log.Printf("%v with %v (%v)", internalElement.Type(), internalElement.Attributes(), internalElement.TextOptimistic())
				}
				num++
			}
			if num != tt.want {
				t.Errorf("Matching elements: %v, want %v", num, tt.want)
			}
		})
	}
}
