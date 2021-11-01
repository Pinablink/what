package local

import (
	"os"
)

// Local
type Local struct {
	os_local string
}

// New Retorna uma instância da estrutura
func New(val_os_local string) *Local {
	return &Local{
		os_local: val_os_local,
	}
}

// InitLog Inicializa o log no sistema
func (refLocal *Local) InitLocal() error {

	_, ok := os.Stat(refLocal.os_local)

	if os.IsNotExist(ok) {
		return os.Mkdir(refLocal.os_local, 0755)
	}

	return nil
}

// Func GetOsLocal
func (refLocal *Local) GetOsLocal() string {
	return refLocal.os_local
}
