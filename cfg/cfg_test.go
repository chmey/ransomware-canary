package cfg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	var tests = map[string]struct {
		configPath  string
		wantErr     bool
		expectedCfg *CanaryConfig
		wantWrite   bool
		writeCfg    string
	}{
		"valid config": {
			configPath: "config.test.toml",
			wantErr:    false,
			expectedCfg: &CanaryConfig{
				ForceOverwrite: true,
				CanaryFileName: "test_canary.txt",
				CanaryDocument: `First Line Of Document
Last Line Of Document
`,
				SendMail:    true,
				SmtpHost:    "example.com",
				SmtpPort:    465,
				SmtpProto:   "SSL",
				SmtpUser:    "example@example.com",
				SmtpPass:    "INSECURE_PASSWORD",
				SmtpFrom:    "example@example.com",
				SmtpSubject: "Email Subject",
				SmtpTo:      []string{"admin@example.com"},
			},
			wantWrite: true,
			writeCfg: `
ForceOverwrite = true
SendMail = true
canaryFileName = "test_canary.txt"
canaryDocument = """
First Line Of Document
Last Line Of Document
"""
smtpHost = "example.com"
smtpPort = 465
smtpProto = "SSL"
smtpUser = "example@example.com"
smtpPass = "INSECURE_PASSWORD"
smtpFrom = "example@example.com"
smtpSubject = "Email Subject"
smtpTo = [ "admin@example.com" ]
`,
		},
		"config does not exist": {
			configPath:  "DOESNOTEXIST",
			wantErr:     true,
			expectedCfg: nil,
			wantWrite:   false,
			writeCfg:    "",
		},
		"invalid config": {
			configPath:  "config.test.toml",
			wantErr:     true,
			expectedCfg: nil,
			wantWrite:   true,
			writeCfg: `
canaryFileName ={= "test_canary.txt"
`,
		},
		"array type is string": {
			configPath:  "config.test.toml",
			wantErr:     true,
			expectedCfg: nil,
			wantWrite:   true,
			writeCfg: `
canaryFileName = "test_canary.txt"
canaryDocument = """
First Line Of Document
Last Line Of Document
"""
smtpHost = "example.com"
smtpPort = 465
smtpProto = "SSL"
smtpUser = "example@example.com"
smtpPass = "INSECURE_PASSWORD"
smtpFrom = "example@example.com"
smtpSubject = "Email Subject"
smtpTo = "admin@example.com"
`,
		},
		"int type is string": {
			configPath:  "config.test.toml",
			wantErr:     true,
			expectedCfg: nil,
			wantWrite:   true,
			writeCfg: `
canaryFileName = "test_canary.txt"
canaryDocument = """
First Line Of Document
Last Line Of Document
"""
smtpHost = "example.com"
smtpPort = "465"
smtpProto = "SSL"
smtpUser = "example@example.com"
smtpPass = "INSECURE_PASSWORD"
smtpFrom = "example@example.com"
smtpSubject = "Email Subject"
smtpTo = [ "admin@example.com" ]
`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.wantWrite {
				err := os.WriteFile(tc.configPath, []byte(tc.writeCfg), 0644)
				require.NoError(err)
			}

			cfg, err := NewConfig(tc.configPath)
			if !tc.wantErr {
				require.NoError(err)
				assert.Equal(*cfg, *tc.expectedCfg)
			} else {
				require.Error(err)
			}

			if tc.wantWrite {
				err := os.Remove(tc.configPath)
				require.NoError(err)
			}
		})
	}
}
