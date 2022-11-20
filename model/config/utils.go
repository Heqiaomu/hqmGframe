package config

func GetConfig() *App {
	return &cfg
}

func GetHttpServer() *HTTPServer {
	return &cfg.HTTPServer
}

func GetAppDebug() bool {
	return cfg.AppDebug
}

func GetZap() *Logs {
	return &cfg.Logs
}

func GetJWT() *Token {
	return &cfg.Token
}
