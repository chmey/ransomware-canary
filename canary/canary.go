package canary

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/chmey/ransomware_canary/cfg"
	"github.com/chmey/ransomware_canary/email"
	"github.com/chmey/ransomware_canary/file"
	"github.com/fsnotify/fsnotify"
)

type Canary struct {
	config       *cfg.CanaryConfig
	absolutePath string
	isSet        bool
	isWatching   bool
	originalHash []byte
}

func NewCanary(config *cfg.CanaryConfig) *Canary {
	return &Canary{
		config: config,
	}
}

func (c *Canary) Start() {
	if !c.isSet {
		err := c.setCanary()
		check(err)
	}
	if !c.isWatching {
		c.watch()
	}
}

func (c *Canary) setCanary() error {
	filePath, err := file.MakeFilePath(c.config.CanaryFileName)
	check(err)

	c.absolutePath = filePath
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		fileContent := []byte(c.config.CanaryDocument)
		err = os.WriteFile(filePath, fileContent, 0o644)
		check(err)
		log.Println("created canary at", c.absolutePath)

		c.originalHash, err = file.GetFileHash(c.absolutePath)
		check(err)

		c.isSet = true
		return nil
	} else {
		return errors.New("canary exists, won't overwrite")
	}
}

func (c *Canary) watch() {
	watcher, err := fsnotify.NewWatcher()
	check(err)
	defer watcher.Close()

	c.isWatching = true
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("carnary was deleted")
					email.SendMail(c.config)
					os.Exit(1)
				}
				newHash, err := file.GetFileHash(c.absolutePath)
				check(err)

				if !compareHashes(c.originalHash, newHash) {
					log.Println("canary was modified")
					email.SendMail(c.config)
					os.Exit(1)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(c.absolutePath)
	check(err)
	<-done
}

func compareHashes(h1 []byte, h2 []byte) bool {
	return bytes.Equal(h1, h2)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
