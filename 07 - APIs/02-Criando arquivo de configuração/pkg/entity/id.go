package entity

import (
	"database/sql/driver" // Import necessário
	"fmt"
	"github.com/google/uuid"
)

// ... (Outras funções NewID, ParseID, String...)

type ID uuid.UUID

// Implementação de driver.Valuer (Como salvar no DB)
// Retorna o valor do ID como []byte (que o SQL driver reconhece como UUID)
func (id ID) Value() (driver.Value, error) {
	return uuid.UUID(id).MarshalBinary()
}

// Implementação de sql.Scanner (Como ler do DB)
// Converte o valor lido do banco de dados de volta para o tipo ID.
func (id *ID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// O valor do DB pode vir como []byte (para UUID) ou string
	switch v := value.(type) {
	case []byte:
		u, err := uuid.FromBytes(v)
		if err != nil {
			return err
		}
		*id = ID(u)
		return nil
	case string:
		u, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		*id = ID(u)
		return nil
	default:
		return fmt.Errorf("tipo não suportado para Scan de ID: %T", v)
	}
}

// Outras funções (NewID, ParseID, String) permanecem as mesmas
func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}