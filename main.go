package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func writeHosts(hostsPath string, hostsFile string) {
	hostsFiles, err := ioutil.ReadDir(hostsPath)
	if err != nil {
		log.Fatal(err)
	}
	tmpFile, err := ioutil.TempFile("", "hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpFile.Close()

	for _, file := range hostsFiles {
		if !file.IsDir() {
			fullPath := filepath.Join(hostsPath, file.Name())
			contents, err := ioutil.ReadFile(fullPath)
			if err != nil {
				log.Print(err)
			}
			tmpFile.WriteString("# " + fullPath + "\n")
			tmpFile.Write(contents)
			tmpFile.Write([]byte("\n"))
		}
	}

	err = tmpFile.Chmod(0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(tmpFile.Name(), hostsFile)
	if err == nil {
		log.Println("Successfuly wrote hosts file")
	} else {
		log.Fatal("Could not write hosts file:", err)
	}
}

func main() {
	hostsPath := flag.String("hostsfiles", "/etc/hosts.d", "directory for hosts files")
	hostsFile := flag.String("hostsfile", "/etc/hosts", "target hosts file")

	writeHosts(*hostsPath, *hostsFile)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Name[len(event.Name)-4:] != ".swp" || event.Name[len(event.Name)-1:] != "~" {
					log.Println("event:", event)
					writeHosts(*hostsPath, *hostsFile)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/etc/hosts.d/")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
