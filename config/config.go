package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
}

type System struct {
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}


