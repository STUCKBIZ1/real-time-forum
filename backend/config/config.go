package config

type DBConfig struct{ 
	Driver string
	DSN string
}
type ServerCongig struct{
	Port string
	Host string
}
var DB = DBConfig{
	Driver: "sqlite3",
	DSN : "./forum.db",
}
var Server = ServerCongig{
	Port: "8080",
	Host: "localhost",
}