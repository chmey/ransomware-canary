package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"

	"github.com/chmey/ransomware_canary/cfg"
)

func SendMail(config *cfg.CanaryConfig) error {
	smtpConnection := fmt.Sprintf("%s:%s", config.SmtpHost, strconv.Itoa(config.SmtpPort))
	auth := smtp.PlainAuth("", config.SmtpUser, config.SmtpPass, config.SmtpHost)
	hostName, err := os.Hostname()

	var message []byte
	if err != nil {
		message = []byte(fmt.Sprintf("The canary file %s has been modified or deleted.", config.CanaryFileName))
	} else {
		message = []byte(fmt.Sprintf("The canary file %s on %s has been modified or deleted.", config.CanaryFileName, hostName))
	}

	return smtp.SendMail(smtpConnection, auth, config.SmtpFrom, config.SmtpTo, message)
}
