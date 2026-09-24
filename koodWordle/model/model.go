package model

type User struct {
	Username string
}

type GameResult struct {
	Username    string
	SecretWord  string
	Attempts    int
	Result      string
}
