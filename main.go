package built

// these vars are built at compile time, DO NOT ALTER
var (
	// Version adds build information
	version string
	// BuildTimestamp adds build information
	timestamp string
	// CompiledBy adds the make/model that was used to compile
	compiledBy string
)

// Version returns the build version set when the binary was built
func Version() string {
	return version
}

// Timestamp returns the build timestamp set when the binary was built
func Timestamp() string {
	return timestamp
}

// CompiledBy returns the build timestamp set when the binary was built
func CompiledBy() string {
	return compiledBy
}
