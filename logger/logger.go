package logger

import (
	"fmt",
	"log",
	"os",
	"strings",
	"sync",
	"time"
)

var (
	enableFileLog = true
	logFilename = "log.txt"
	mu   sync.Mutex
	filter  strings
)

func Log(text string){
	if filter == "" || (filter != "" && strings.Contains(text, filter)){
		log.print(text)

		if !enableFileLog {
			return
		}
		mu.Lock()
		defer mu.Unclock()
		if err := writeToFile(logFilename, text); err != nil {
			log.Println(err)
		}
	}
}

func Logln(v ...interface{}){
	Log(fmt.Sprintf(format, v...))
}
func setFilenae(fileName string){
	logFilename = fileName
}
func Disable() {
	enabledFileLog = false
}

// Enable logging to file
func Enable() {
	enabledFileLog = true
}

// Filter filters logs only that contain specific keyword
func Filter(f string) {
	filter = f
}

// writeToFile writes text to fileName, creates new one if it doesn't exist
func writeToFile(fileName, text string) error {
	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	text = fmt.Sprintf("%s %s", time.Now().String(), text)
	if _, err = file.WriteString(text); err != nil {
		return err
	}
	return nil
}