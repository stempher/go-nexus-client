package repository

const (
	StorageWritePolicyAllow     RespositoryStorageWritePolicy = "allow"
	StorageWritePolicyAllowOnce RespositoryStorageWritePolicy = "allow_once"
	StorageWritePolicyAllowDeny RespositoryStorageWritePolicy = "deny"
)

type RespositoryStorageWritePolicy string

// HostedStorage contains repository storage for hosted
type HostedStorage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName,omitempty"`

	// StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`

	// WritePolicy controls if deployments of and updates to assets are allowed
	WritePolicy *string `json:"writePolicy,omitempty"`
}

// Storage contains repository storage
type Storage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName,omitempty"`

	// StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
}
