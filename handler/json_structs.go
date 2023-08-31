package handler

type singUpJSON struct {
	Username  string `json:"username"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Is_author bool   `json:"is_author"`
}

type logInJSON struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
