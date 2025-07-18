package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestParseConfig(t *testing.T) {
	testCases := []struct {
		name string
		tmpl string
		exp  string
		err  bool
	}{
		{
			name: "simple",
			tmpl: `<!--- {{ "hello!" | toUpper | repeat 5 }} --->`,
			exp: `<!--- {{ "hello!" | toUpper | repeat 5 }} --->
HELLO!HELLO!HELLO!HELLO!HELLO!
`,
		},
		{
			name: "exec",
			tmpl: `<!--- {{ exec "echo hallo" | repeat 3 | truncate }} --->`,
			exp: `<!--- {{ exec "echo hallo" | repeat 3 | truncate }} --->
hallo
hallo
hallo
`,
		},
		{
			name: "fle",
			tmpl: `<!--- {{ file "testdata/cfg.yml" | code "yml" }} --->`,
			exp: `<!--- {{ file "testdata/cfg.yml" | code "yml" }} --->` + "\n```yml" + `
settings:
  cfg: true

` + "```\n",
		},
		{
			name: "tmpl",
			tmpl: `<!--- {{ tmpl "testdata/tmpl.tmpl" | truncate }} --->`,
			exp: `<!--- {{ tmpl "testdata/tmpl.tmpl" | truncate }} --->
This is a test template
`,
		},
		{
			name: "tmplWithVars",
			tmpl: `<!--- {{ tmplWithVars "testdata/template.tmpl" (file "./testdata/values.yml" | fromYAML) }} --->`,
			exp: `<!--- {{ tmplWithVars "testdata/template.tmpl" (file "./testdata/values.yml" | fromYAML) }} --->
username=admin
password=password

`,
		},
		{
			name: "regularComment",
			tmpl: `<!--- regular comment --->`,
			exp: `<!--- regular comment --->
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := strings.NewReader(tc.tmpl)
			res, err := parse(s)

			if tc.err {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.exp, string(res))
		})
	}
}
