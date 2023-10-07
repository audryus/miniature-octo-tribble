package types

type contextKey string

const (
	Mongodb        = contextKey("mongodb")
	Config         = contextKey("config")
	EmpresaRepo    = contextKey("empresaRepo")
	VencimentoRepo = contextKey("vencimentoRepo")
	SerieRepo      = contextKey("serieRepo")
	EmpresaUC      = contextKey("empresaUC")
	OpcoesUC       = contextKey("opcoesUC")
	SerieUC        = contextKey("serieUC")
)
