package configs // Pacote configs do arquivo config.go


var cfg *config // Variável para armazenar as configurações do aplicativo.

type config struct { // Estrutura struct config APIConfig DBConfig serve para armazenar as configurações de API e DB.
	API APIConfig
	DB DBConfig 
}

type APIConfig struct { // A estrutura APIConfig representa as configurações da API.
	Port string
}

type DBConfig struct { // A estrutura DBConfig representa as configurações do banco de dados.
	Host string
	Port string
	User string
	Pass string
	Database string
}

