package types

import (
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog"
)

type Status string

const (
	RuntimeProviderKey        = "Provider"
	StatusInitializing Status = "initializing"
	StatusRunning      Status = "running"
	StatusTerminated   Status = "terminated"
	StatusError        Status = "error"
	StatusFailed       Status = "failed"
	StatusSuccess      Status = "success"
	StatusSnapshotting Status = "snapshotting"
)

type NoOpLogHook struct{}

func (d NoOpLogHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {}

var NewErrLogHook = func() zerolog.Hook { return NoOpLogHook{} }

// Provider is the platform that the node is spawned on. This is where the user
// code runs
type Provider string

func (r Provider) String() string {
	return string(r)
}

const (
	LambdaLabsProvider Provider = "lambdalabs"
	UnweaveProvider    Provider = "unweave"
)

func (r Provider) DisplayName() string {
	switch r {
	case LambdaLabsProvider:
		return "LambdaLabs"
	case UnweaveProvider:
		return "Unweave"
	default:
		return "Unknown"
	}
}

type Build struct {
	BuildID     string     `json:"buildID"`
	Name        string     `json:"name"`
	ProjectID   string     `json:"projectID"`
	Status      string     `json:"status"`
	BuilderType string     `json:"builderType"`
	CreatedAt   time.Time  `json:"createdAt"`
	StartedAt   *time.Time `json:"startedAt,omitempty"`
	FinishedAt  *time.Time `json:"finishedAt,omitempty"`
}

type LogEntry struct {
	TimeStamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Level     string    `json:"level"`
}

func (l LogEntry) String() string {
	return fmt.Sprintf("%s %s %s", l.TimeStamp.Format(time.RFC3339), l.Level, l.Message)
}

type NodeSpecs struct {
	VCPUs int `json:"vCPUs"`
	// Memory is the RAM in GB
	Memory int `json:"memory"`
	// Storage is the storage in GB
	Storage int `json:"storage"`
	// GPUType is the type of GPU
	GPUType string `json:"gpuType"`
	// GPUCount is the number of GPUs
	GPUCount int `json:"gpuCount"`
	// GPUMemory is the GPU RAM in GB
	GPUMemory *int `json:"gpuMemory"`
}

type NodeType struct {
	ID       string    `json:"id"`
	Name     *string   `json:"name"`
	Price    *int      `json:"price"`
	Regions  []string  `json:"regions"`
	Provider Provider  `json:"provider"`
	Specs    NodeSpecs `json:"specs"`
}

type Node struct {
	ID       string    `json:"id"`
	TypeID   string    `json:"typeID"`
	OwnerID  string    `json:"ownerID"`
	Price    int       `json:"price"`
	Region   string    `json:"region"`
	KeyPair  SSHKey    `json:"sshKeyPair"`
	Status   Status    `json:"status"`
	Provider Provider  `json:"provider"`
	Specs    NodeSpecs `json:"specs"`
	Host     string    `json:"host"`
	User     string    `json:"user"`
	Port     int       `json:"port"`
}

type Project struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type SSHKey struct {
	Name      string     `json:"name"`
	PublicKey *string    `json:"publicKey,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type ConnectionInfo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
}

type Exec struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	SSHKey       SSHKey          `json:"sshKey"`
	Connection   *ConnectionInfo `json:"connection,omitempty"`
	Status       Status          `json:"status"`
	CreatedAt    *time.Time      `json:"createdAt,omitempty"`
	NodeTypeID   string          `json:"nodeTypeID"`
	Region       string          `json:"region"`
	Provider     Provider        `json:"provider"`
	PersistentFS bool            `json:"persistentFS"`
}

type ExecCtx struct {
	Command  []string      `json:"command"`
	CommitID *string       `json:"commitID,omitempty"`
	GitURL   *string       `json:"gitURL,omitempty"`
	BuildID  *string       `json:"buildID,omitempty"`
	Context  io.ReadCloser `json:"-"`
}
