package cmd

type configFlags struct {
	configFile string
}

type domainFlags struct {
	domainName string
}

type listLocalInfo struct {
	configFlags
	domainFlags
}
