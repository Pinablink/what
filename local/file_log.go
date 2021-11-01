package local

import (
	"errors"
	"fmt"
	"strings"
)

const ERRO_STR_NAME_FILE string = "Não utilize caracter especial no nome do arquivo."

//
type Filelog struct {
	logNameFile string
	dateFile    string
	nameFileLog string
}

// New Retorna uma nova instancia da Struct Filelog
func NewFilelog(strFileLogName string, strDateFile string) *Filelog {
	return &Filelog{logNameFile: strFileLogName,
		dateFile: strDateFile}
}

// InitFilelog Inicializa o Filelog especifico
func (ref *Filelog) InitFilelog(pathLog string) error {

	error := ref.validName()

	if error != nil {
		return error
	}

	strLog := fmt.Sprintf("%s/%s_%s.log", pathLog, ref.logNameFile, ref.dateFile)
	ref.nameFileLog = strLog

	return nil
}

// Func validName Verifica se existe conteúdo na string informada
func (ref *Filelog) validName() error {

	f := func(r rune) bool {
		return (r < 'A' || r > 'z')
	}

	if strings.IndexFunc(ref.logNameFile, f) != -1 {
		error := errors.New(ERRO_STR_NAME_FILE)
		return error
	}

	return nil
}

// Func GetNameLog Retorna o caminho e o nome do arquivo configurado
func (ref *Filelog) GetNameLog() string {
	return ref.nameFileLog
}
