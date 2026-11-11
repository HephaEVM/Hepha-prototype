package main

type target struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	CompilerVersion  string `json:"compilerVersion"`
	OptimizerEnabled *bool  `json:"optimizerEnabled"`
	Runs             int    `json:"runs"`
	SourceURL        string `json:"etherscanURL"`
}
