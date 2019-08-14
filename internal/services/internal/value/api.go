package value

// Transformer allows a value to be transformed before being read from or
// written to the underlying store. The methods must be able to undo the
// transformation caused by the other.
type Transformer interface {
	// TransformFromStorage may transform the provided data from its underlying
	// storage representation or return an error.
	// Stale is true if the object on disk is stale and a write to etcd should
	// be issued, even if the contents of the object have not changed.
	TransformFromStorage(data []byte) (out []byte, stale bool, err error)
	// TransformToStorage may transform the provided data into the appropriate
	// form in storage or return an error.
	TransformToStorage(data []byte) (out []byte, err error)
}
