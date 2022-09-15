package logger

import (
	"log"
	"os"
)

//Log структура для описания пользовательского логера
type Log struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

//NewLogger вернет новый экземпляр логера
func NewLogger(path string) (*Log, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil{
		file.Close()
		return nil, err
	}

	infoLog := log.New(file, "INFO: ", log.Ldate|log.Ltime)
	errLog := log.New(file, "ERR: ", log.Ldate|log.Ltime)

	logger := &Log{
		infoLog:  infoLog,
		errorLog: errLog,
	}

	return logger, nil
}

//Info сообщение уровня info
func (l *Log) Info(mes string){
	l.infoLog.Println(mes)
}

//Error сообщение уровня error
func (l *Log) Error(mes string) {
	l.errorLog.Println(mes)
}
