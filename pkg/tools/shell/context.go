package shell

import (
	"os"
	"os/user"
	"runtime"
)

type EnvironmentContext struct {
	OS                  string
	Shell               string
	User                string
	IsWindowsStyleFlags bool
}

func NewEnvironmentContext() (EnvironmentContext, error) {
	user, err := getUser()
	if err != nil {
		return EnvironmentContext{}, err
	}

	sh, isWindowsStyleFlags := getShell()

	envContext := EnvironmentContext{
		OS:                  getOS(),
		Shell:               sh,
		IsWindowsStyleFlags: isWindowsStyleFlags,
		User:                user,
	}

	return envContext, nil
}

func getOS() string {
	return runtime.GOOS
}

func getShell() (shell string, isWindowsStyleFlags bool) {
	if os.Getenv("PSModulePath") != "" {
		return "powershell", true
	}
	csp := os.Getenv("ComSpec")
	if csp != "" {
		return csp, true
	}
	return os.Getenv("SHELL"), false
}

func getUser() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}
