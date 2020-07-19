package scraper

import (
	"golang.org/x/net/html"
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestEmptyTarget_Content(t *testing.T) {
	type fields struct {
		name    string
		content *html.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *html.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := EmptyTarget{
				name:    tt.fields.name,
				content: tt.fields.content,
			}
			if got := em.Content(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyTarget_IsValid(t *testing.T) {
	type fields struct {
		name    string
		content *html.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := EmptyTarget{
				name:    tt.fields.name,
				content: tt.fields.content,
			}
			if got := em.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyTarget_Render(t *testing.T) {
	type fields struct {
		name    string
		content *html.Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := EmptyTarget{
				name:    tt.fields.name,
				content: tt.fields.content,
			}
			got, err := em.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Render() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromBuffer(t *testing.T) {
	type args struct {
		buffer io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    *Scraper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromBuffer(tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromBuffer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromNode(t *testing.T) {
	type args struct {
		node *html.Node
	}
	tests := []struct {
		name    string
		args    args
		want    *Scraper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromNode(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_Attributes(t *testing.T) {
	type fields struct {
		target Target
	}
	tests := []struct {
		name   string
		fields fields
		want   Attributes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.Attributes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_Content(t *testing.T) {
	type fields struct {
		target Target
	}
	tests := []struct {
		name   string
		fields fields
		want   *html.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.Content(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_Find(t *testing.T) {
	type fields struct {
		target Target
	}
	type args struct {
		filters Filters
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Scraper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.Find(tt.args.filters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_FindAll(t *testing.T) {
	type fields struct {
		target Target
	}
	type args struct {
		filters Filters
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   <-chan *Scraper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.FindAll(tt.args.filters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_Render(t *testing.T) {
	type fields struct {
		target Target
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			got, err := scraper.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Render() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_Text(t *testing.T) {
	type fields struct {
		target Target
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			got, got1 := scraper.Text()
			if got != tt.want {
				t.Errorf("Text() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Text() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScraper_TextO(t *testing.T) {
	type fields struct {
		target Target
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.TextO(); got != tt.want {
				t.Errorf("TextO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_Type(t *testing.T) {
	type fields struct {
		target Target
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScraper_getLastSubNode(t *testing.T) {
	type fields struct {
		target Target
	}
	type args struct {
		node *html.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *html.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := Scraper{
				target: tt.fields.target,
			}
			if got := scraper.getLastSubNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLastSubNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseError(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := baseError(tt.args.err, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("baseError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_htmlTarget_Content(t *testing.T) {
	type fields struct {
		content *html.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *html.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := htmlTarget{
				content: tt.fields.content,
			}
			if got := target.Content(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_htmlTarget_IsValid(t *testing.T) {
	type fields struct {
		content *html.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := htmlTarget{
				content: tt.fields.content,
			}
			if got := target.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_htmlTarget_Render(t *testing.T) {
	type fields struct {
		content *html.Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := htmlTarget{
				content: tt.fields.content,
			}
			got, err := target.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Render() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isContentValid(t *testing.T) {
	type args struct {
		content io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isContentValid(tt.args.content); got != tt.want {
				t.Errorf("isContentValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadContent(t *testing.T) {
	type args struct {
		buffer io.ReadCloser
	}
	tests := []struct {
		name     string
		args     args
		wantNode *html.Node
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, err := loadContent(tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("loadContent() gotNode = %v, want %v", gotNode, tt.wantNode)
			}
		})
	}
}

func Test_newFromTarget(t *testing.T) {
	type args struct {
		target Target
	}
	tests := []struct {
		name        string
		args        args
		wantScraper *Scraper
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScraper, err := newFromTarget(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("newFromTarget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotScraper, tt.wantScraper) {
				t.Errorf("newFromTarget() gotScraper = %v, want %v", gotScraper, tt.wantScraper)
			}
		})
	}
}

func Test_newTargetFromBuffer(t *testing.T) {
	type args struct {
		buffer io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    *htmlTarget
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTargetFromBuffer(tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTargetFromBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTargetFromBuffer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newTargetFromNode(t *testing.T) {
	type args struct {
		node *html.Node
	}
	tests := []struct {
		name string
		args args
		want *htmlTarget
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTargetFromNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTargetFromNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchTreeLayer(t *testing.T) {
	type args struct {
		operations *sync.WaitGroup
		node       *html.Node
		callable   func(node2 *html.Node)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
