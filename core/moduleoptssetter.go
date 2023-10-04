package core

func ServiceName(name string) ModuleOptsSetter {
	return func(options *ModuleOptions) {
		options.ServiceName = name
	}
}

func ServiceVersion(version string) ModuleOptsSetter {
	return func(options *ModuleOptions) {
		options.ServiceVersion = version
	}
}

func ServiceAddress(address string) ModuleOptsSetter {
	return func(options *ModuleOptions) {
		options.HttpAddress = address
	}
}

func ServiceConfigurationPath(path string) ModuleOptsSetter {
	return func(options *ModuleOptions) {
		options.Configuration.Path = path
	}
}

func ServiceConfigurationEnvPrefix(prefix string) ModuleOptsSetter {
	return func(options *ModuleOptions) {
		options.Configuration.EnvironmentPrefix = prefix
	}
}

func ServiceConfigurationRootKey(rootKey string) ModuleOptsSetter {
	return func(options *ModuleOptions) {
		options.Configuration.RootKey = rootKey
	}
}
