// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UnweaveSessionStatus string

const (
	UnweaveSessionStatusInitializing UnweaveSessionStatus = "initializing"
	UnweaveSessionStatusActive       UnweaveSessionStatus = "active"
	UnweaveSessionStatusTerminated   UnweaveSessionStatus = "terminated"
)

func (e *UnweaveSessionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UnweaveSessionStatus(s)
	case string:
		*e = UnweaveSessionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UnweaveSessionStatus: %T", src)
	}
	return nil
}

type NullUnweaveSessionStatus struct {
	UnweaveSessionStatus UnweaveSessionStatus
	Valid                bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUnweaveSessionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UnweaveSessionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UnweaveSessionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUnweaveSessionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UnweaveSessionStatus, nil
}

type UnweaveAccount struct {
	ID uuid.UUID `json:"id"`
}

type UnweaveProject struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uuid.UUID `json:"ownerID"`
	CreatedAt time.Time `json:"createdAt"`
}

type UnweaveSession struct {
	ID        uuid.UUID            `json:"id"`
	Name      string               `json:"name"`
	NodeID    string               `json:"nodeID"`
	CreatedBy uuid.UUID            `json:"createdBy"`
	CreatedAt time.Time            `json:"createdAt"`
	ReadyAt   sql.NullTime         `json:"readyAt"`
	ExitedAt  sql.NullTime         `json:"exitedAt"`
	Status    UnweaveSessionStatus `json:"status"`
	ProjectID uuid.UUID            `json:"projectID"`
	Provider  string               `json:"provider"`
	SshKeyID  uuid.UUID            `json:"sshKeyID"`
}

type UnweaveSshKey struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uuid.UUID `json:"ownerID"`
	CreatedAt time.Time `json:"createdAt"`
	PublicKey string    `json:"publicKey"`
}
