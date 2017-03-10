package wheel

type CloudProvider interface {
	ProvisionBuildEnvironment(name string) error
}
