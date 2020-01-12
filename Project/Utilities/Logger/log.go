package logger

import (
	"fmt"
	"log"
)

const (
	normalText  = "\033[0m"
	boldText    = "\033[1m"
	blackText   = "\033[30;1m"
	redText     = "\033[31;1m"
	greenText   = "\033[32;1m"
	yellowText  = "\033[33;1m"
	blueText    = "\033[34;1m"
	magentaText = "\033[35;1m"
	cyanText    = "\033[36;1m"
	whiteText   = "\033[37;1m"
)

var (
	infoTag    = greenText + "Info: " + normalText
	warningTag = yellowText + "Warning: " + normalText
	errorTag   = redText + "Error: " + normalText
	debugTag   = blueText + "Debug: " + normalText
	fatalTag   = redText + "Fatal: " + normalText
)

// Infoln ... Logs an information message
func Infoln(message string) {
	fmt.Print(infoTag + message + "\n")
}

// Warnln ... Logs an warning message
func Warnln(message string) {
	fmt.Print(warningTag + message + "\n")
}

// Errorln ... Logs an error message
func Errorln(message string) {
	fmt.Print(errorTag + message + "\n")
}

// Debugln ... Logs an debug message
func Debugln(message string) {
	fmt.Print(debugTag + message + "\n")
}

// Fatalln ... Logs an fatal message and calls panic()
func Fatalln(message string) {
	log.Print(fatalTag + message + "\n")
	panic("")
}

// Infof ... Logs an information message formatted
func Infof(message string, values ...interface{}) {
	fmt.Printf(infoTag+message, values...)
}

// Warnf ... Logs an warning message formatted
func Warnf(message string, values ...interface{}) {
	fmt.Printf(warningTag+message, values...)
}

// Errorf ... Logs an error message formatted
func Errorf(message string, values ...interface{}) {
	fmt.Printf(errorTag+message, values...)
}

// Debugf ... Logs an debug message formatted
func Debugf(message string, values ...interface{}) {
	fmt.Printf(debugTag+message, values...)
}

// Fatalf ... Logs an fatal message formatted then calls panic()
func Fatalf(message string, values ...interface{}) {
	log.Printf(fatalTag+message, values...)
	panic("")
}
