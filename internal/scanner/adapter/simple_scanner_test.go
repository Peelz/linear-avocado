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
			name: "test PublicKeyDetect",
			args: args{
				b: []byte("1st line\nnewline before public_key tailing\nfoo bar\n"),
			},
			want: []scanner.Finding{
				{
					Type:     PublicKeyDetect.Type,
					RuleId:   PublicKeyDetect.ID,
					Location: scanner.Location{},
					Metadata: scanner.Metadata{
						Description: PublicKeyDetect.Description,
						Severity:    PublicKeyDetect.Severity,
					},
				},
			},
			assertErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "test PrivateKeyDetect",
			args: args{
				b: []byte("1st line\nnewline before private_key tailing\nfoo bar\n"),
			},
			want: []scanner.Finding{
				{
					Type:     PrivateKeyDetect.Type,
					RuleId:   PrivateKeyDetect.ID,
					Location: scanner.Location{},
					Metadata: scanner.Metadata{
						Description: PrivateKeyDetect.Description,
						Severity:    PrivateKeyDetect.Severity,
					},
				},
			},
			assertErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "test PrivateKeyDetect substr",
			args: args{
				b: []byte("1st line\nnewline before headingprivate_keytailing\nfoo bar\n"),
			},
			want: []scanner.Finding{
				{
					Type:     PrivateKeyDetect.Type,
					RuleId:   PrivateKeyDetect.ID,
					Location: scanner.Location{},
					Metadata: scanner.Metadata{
						Description: PrivateKeyDetect.Description,
						Severity:    PrivateKeyDetect.Severity,
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
