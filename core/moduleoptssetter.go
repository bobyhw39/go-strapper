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
