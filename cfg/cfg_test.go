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

	var tests = []struct {
		configPath  string
		wantErr     bool
		expectedCfg *CanaryConfig
		wantWrite   bool
		writeCfg    string
	}{
		{
			configPath: "config.test.toml",
			wantErr:    false,
			expectedCfg: &CanaryConfig{
				CanaryFileName: "test_canary.txt",
				CanaryDocument: `First Line Of Document
Last Line Of Document
`,
				SmtpHost:    "example.com",
				SmtpPort:    465,
				SmtpProto:   "SSL",
				SmtpUser:    "example@example.com",
				SmtpPass:    "INSECURE_PASSWORD",
				SmtpFrom:    "example@example.com",
				SmtpSubject: "Email Subject",
			},
			wantWrite: true,
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
`,
		},
		{
			configPath:  "DOESNOTEXIST",
			wantErr:     true,
			expectedCfg: nil,
			wantWrite:   false,
			writeCfg:    "",
		},
		{
			configPath:  "config.test.toml",
			wantErr:     true,
			expectedCfg: nil,
			wantWrite:   true,
			writeCfg: `
canaryFileName ={= "test_canary.txt"
`,
		},
	}
	for _, tc := range tests {
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
	}

}
