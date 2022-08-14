package adapter

import (
	"github.com/monopeelz/linear-avocado/pkg/scanner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_codeScanner_ScanByte(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name      string
		args      args
		want      []scanner.Finding
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				b: []byte("public test"),
			},
			want: []scanner.Finding{
				{
					Type:     "swd",
					RuleId:   "G0001",
					Location: scanner.Location{},
					Metadata: scanner.Metadata{
						Description: "Sensitive keyword 'public', 'private' detected",
						Severity:    "LOW",
					},
				},
			},
			assertErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := simpleScanner{}
			got, err := c.scanByte(tt.args.b)
			tt.assertErr(t, err)
			assert.Equal(t, tt.want, got, "scanByte() got = %v, want %v", got, tt.want)
		})
	}
}

func Test_codeScanner_checkSensitiveWord(t *testing.T) {
	type args struct {
		i []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				i: []byte("private"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := simpleScanner{}
			assert.Equalf(t, tt.want, c.scanSecret(tt.args.i), "scanSecret(%v)", tt.args.i)
		})
	}
}
