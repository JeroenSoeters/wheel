package wheel

type CloudProvider interface {
	ProvisionBuildEnvironment() error
}
