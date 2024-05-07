package config

type ServiceConfigFile struct {
	ModuleName     string         `toml:"module_name"`
	Service        ServiceConfig  `toml:"service"`
	Tables         Table `toml:"table"`
	Consumers      []string       `toml:"consumers"`
	GopVersion     string         `toml:"gop_version"`
	GenerationDate string
	GopVersionGen  string
}

type ServiceConfig struct {
	Name     string    `toml:"name"`
	Handlers []Handler `toml:"handler"`
	Port 	 int       `toml:"port"`
    Namespace string   `toml:"namespace"`
    Image string       `toml:"image"`
}

type Handler struct {
	Method      string  `toml:"method"`
	Path        string  `toml:"path"`
	Tables      []string `toml:"tables"`
	Topics      []string `toml:"topics"`
}

type DatabaseConfig struct {
	Tables []Table `toml:"table"`
}

type Table struct {
	Name    string   `toml:"name"`
	Columns []Column `toml:"column"`
}

type Column struct {
	Name       string   `toml:"name"`
	Type       string   `toml:"type"`
	Attributes []string `toml:"attributes"`
}
