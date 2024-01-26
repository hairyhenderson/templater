package gomplate

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/gomplate/v4/internal/config"
	"github.com/hairyhenderson/gomplate/v4/internal/datafs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderTemplate(t *testing.T) {
	wd, _ := os.Getwd()
	t.Cleanup(func() {
		_ = os.Chdir(wd)
	})
	_ = os.Chdir("/")

	fsys := fstest.MapFS{}
	fsp := fsimpl.NewMux()
	fsp.Add(datafs.EnvFS)
	fsp.Add(datafs.StdinFS)
	fsp.Add(datafs.WrappedFSProvider(fsys, "mem", ""))
	ctx := datafs.ContextWithFSProvider(context.Background(), fsp)

	// no options - built-in function
	tr := NewRenderer(Options{})
	out := &bytes.Buffer{}
	err := tr.Render(ctx, "test", "{{ `hello world` | toUpper }}", out)
	require.NoError(t, err)
	assert.Equal(t, "HELLO WORLD", out.String())

	// with datasource and context
	hu, _ := url.Parse("stdin:")
	wu, _ := url.Parse("env:WORLD")

	t.Setenv("WORLD", "world")

	tr = NewRenderer(Options{
		Context: map[string]config.DataSource{
			"hi": {URL: hu},
		},
		Datasources: map[string]config.DataSource{
			"world": {URL: wu},
		},
	})
	ctx = datafs.ContextWithStdin(ctx, strings.NewReader("hello"))
	out = &bytes.Buffer{}
	err = tr.Render(ctx, "test", `{{ .hi | toUpper }} {{ (ds "world") | toUpper }}`, out)
	require.NoError(t, err)
	assert.Equal(t, "HELLO WORLD", out.String())

	// with a nested template
	nu, _ := url.Parse("nested.tmpl")
	fsys["nested.tmpl"] = &fstest.MapFile{Data: []byte(
		`<< . | toUpper >>`)}

	tr = NewRenderer(Options{
		Templates: map[string]config.DataSource{
			"nested": {URL: nu},
		},
		LDelim: "<<",
		RDelim: ">>",
	})
	out = &bytes.Buffer{}
	err = tr.Render(ctx, "test", `<< template "nested" "hello" >>`, out)
	require.NoError(t, err)
	assert.Equal(t, "HELLO", out.String())

	// errors contain the template name
	tr = NewRenderer(Options{})
	err = tr.Render(ctx, "foo", `{{ bogus }}`, &bytes.Buffer{})
	assert.ErrorContains(t, err, "template: foo:")
}

//// examples

func ExampleRenderer() {
	ctx := context.Background()

	// create a new template renderer
	tr := NewRenderer(Options{})

	// render a template to stdout
	err := tr.Render(ctx, "mytemplate",
		`{{ "hello, world!" | toUpper }}`,
		os.Stdout)
	if err != nil {
		fmt.Println("gomplate error:", err)
	}

	// Output:
	// HELLO, WORLD!
}

func ExampleRenderer_manyTemplates() {
	ctx := context.Background()

	// create a new template renderer
	tr := NewRenderer(Options{})

	templates := []Template{
		{
			Name:   "one.tmpl",
			Text:   `contents of {{ tmpl.Path }}`,
			Writer: &bytes.Buffer{},
		},
		{
			Name:   "two.tmpl",
			Text:   `{{ "hello world" | toUpper }}`,
			Writer: &bytes.Buffer{},
		},
		{
			Name:   "three.tmpl",
			Text:   `1 + 1 = {{ math.Add 1 1 }}`,
			Writer: &bytes.Buffer{},
		},
	}

	// render the templates
	err := tr.RenderTemplates(ctx, templates)
	if err != nil {
		panic(err)
	}

	for _, t := range templates {
		fmt.Printf("%s: %s\n", t.Name, t.Writer.(*bytes.Buffer).String())
	}

	// Output:
	// one.tmpl: contents of one.tmpl
	// two.tmpl: HELLO WORLD
	// three.tmpl: 1 + 1 = 2
}

func ExampleRenderer_datasources() {
	ctx := context.Background()

	// a datasource that retrieves JSON from a public API
	u, _ := url.Parse("https://ipinfo.io/1.1.1.1")
	tr := NewRenderer(Options{
		Context: map[string]config.DataSource{
			"info": {URL: u},
		},
	})

	err := tr.Render(ctx, "jsontest",
		`{{"\U0001F30E"}} {{ .info.hostname }} is served from {{ .info.city }}, {{ .info.region }}`,
		os.Stdout)
	if err != nil {
		panic(err)
	}

	// Output:
	// 🌎 one.one.one.one is served from Los Angeles, California
}
