package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/fsnotify/fsnotify"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <directory-to-watch> <command-to-run>", os.Args[0])
	}

	dirToWatch := os.Args[1]
	commandToRun := os.Args[2]

	log.Printf("Watching directory: %s", dirToWatch)
	log.Printf("Running command: %s", commandToRun)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	// Handle OS signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("Received shutdown signal, exiting...")
		done <- true
	}()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("Event detected:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Modified file:", event.Name)
					runCommand(commandToRun)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	err = filepath.Walk(dirToWatch, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func runCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = getWorkingDirectory()

	// Run command in a separate goroutine to avoid blocking
	go func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Command execution failed: %v", err)
		}
		fmt.Printf("Command output: %s\n", output)
	}()
}

func getWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	return dir
}
