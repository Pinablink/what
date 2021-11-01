package what

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Pinablink/what/local"
)

const (
	FORMAT_STR_DATE        string = "20060102"
	ERROR_CREATE_DIR_LOG   string = "Ocorreu erro na criação do diretório de log."
	ERROR_CREATE_FILE_LOOG string = "Ocorreu erro na criação de file log"
)

// Struct What
type What struct {
	strDateCurrent string
	strNameLog     string
	refFile        *os.File
	warninglogger  *log.Logger
	errorlogger    *log.Logger
	infologger     *log.Logger
	localLog       *local.Local
	fileLog        *local.Filelog
}

// Function New
func NewWhat(path_log string, name_log string) *What {
	var dateCurrent string = dateCurrent()
	return &What{strDateCurrent: dateCurrent,
		strNameLog: name_log,
		localLog:   local.New(path_log),
		fileLog:    local.NewFilelog(name_log, dateCurrent),
	}
}

// Function dateCurrent
func dateCurrent() string {
	const formatDate = FORMAT_STR_DATE
	var time = time.Now()
	return time.Format(formatDate)
}

// InitWhat
func (ref_what *What) InitWhat() error {

	error := ref_what.localLog.InitLocal()

	if error != nil {
		errorFmt := fmt.Sprintf("%s \n Detalhe: \n %s", ERROR_CREATE_DIR_LOG, error.Error())
		return errors.New(errorFmt)
	}

	error = ref_what.fileLog.InitFilelog(ref_what.localLog.GetOsLocal())

	if error != nil {
		errorFmt := fmt.Sprintf("%s \n Detalhe: \n %s", ERROR_CREATE_FILE_LOOG, error.Error())
		return errors.New(errorFmt)
	}

	error = ref_what.initLog()

	if error != nil {
		return error
	}

	return nil
}

//
func (ref_what *What) initLog() error {
	var mError error
	ref_what.refFile, mError = os.OpenFile(ref_what.fileLog.GetNameLog(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if mError != nil {
		errorFmt := fmt.Sprintf("%s \n Detalhe: \n %s", ERROR_CREATE_FILE_LOOG, mError.Error())
		return errors.New(errorFmt)
	}

	ref_what.startLog()

	return nil
}

func (ref_what *What) startLog() {
	ref_what.warninglogger = log.New(ref_what.refFile, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ref_what.infologger = log.New(ref_what.refFile, "INFO: ", log.LstdFlags|log.Lshortfile)
	ref_what.errorlogger = log.New(ref_what.refFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

//
func (ref_what *What) validDate() {
	if ref_what.strDateCurrent != dateCurrent() {
		ref_what.refFile.Close()
		ref_what.fileLog = local.NewFilelog(ref_what.strNameLog, dateCurrent())
		ref_what.initLog()
	}
}

//
func (ref_what *What) Info() *log.Logger {
	ref_what.validDate()
	return ref_what.infologger
}

//
func (ref_what *What) Warning() *log.Logger {
	ref_what.validDate()
	return ref_what.warninglogger
}

//
func (ref_what *What) Error() *log.Logger {
	ref_what.validDate()
	return ref_what.errorlogger
}
