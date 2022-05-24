package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

var fileHash []byte

func main() {
	canaryPath, err := createCanary()
	check(err)

	log.Println("created canary at", canaryPath)
	fileHash, err = getFileHash(canaryPath)
	check(err)

	watcher, err := fsnotify.NewWatcher()
	check(err)
	defer watcher.Close()

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
					os.Exit(1)
				}
				newHash, err := getFileHash(canaryPath)
				check(err)

				if !compareHashes(fileHash, newHash) {
					log.Println("canary was modified")
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

	err = watcher.Add(canaryPath)
	check(err)
	<-done
}

func createCanary() (string, error) {
	fileName, err := makeFileName()
	check(err)

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		fileContent := []byte(canaryDocument)
		err = os.WriteFile(fileName, fileContent, 0o644)
		check(err)
		return fileName, nil
	} else {
		return "", errors.New("canary exists, won't overwrite")
	}
}

func makeFileName() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	} else {
		return home + string(os.PathSeparator) + canaryFileName, nil
	}
}

func getFileHash(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func compareHashes(h1 []byte, h2 []byte) bool {
	return bytes.Equal(h1, h2)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
