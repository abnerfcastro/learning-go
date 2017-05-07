package utils

// MyName will be exported because it starts with a capital letter
var MyName = "Abner Castro"
var whoami = "You're " + MyName

// WhoAmI returns unexported whoami var string from package utils
func WhoAmI() string {
	return whoami
}
