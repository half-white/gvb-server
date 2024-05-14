package config

type Logger struct {
	Level        string `yame:"level"`
	Prefix       string `yame:"prefix"`
	Director     string `yame:"director"`
	ShowLine     bool   `yame:"show_line"`      //显示行号
	LogInConsole bool   `yame:"log_in_console"` //显示打印路径
}
