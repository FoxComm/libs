package configs

import (
	"net"
	"os"
	"strconv"

	"github.com/FoxComm/libs/logger"
)

//Finds the run port by checking the following sources, in order:
//1. OS Flag: "go run main.go -p <3005>"
//2. PORT env variable: "PORT=3005 go run main.go"
//3. Configs package port.  Ie, default.
func GetRunPort(defaultPort int) int {
	//First, let's check the port that may have been passed in as part of running this app.
	if len(os.Args) > 2 && os.Args[1] == "-p" {
		if len(os.Args) == 3 {
			argPort, err := strconv.Atoi(os.Args[2])
			if err != nil {
				logger.Error("The port passed into the runner is not valid.")
				os.Exit(1)
			} else {
				logger.Debug("Mounting service with a passed-in flag port : %s", argPort)
				return argPort
			}
		}
	}

	//Then, let's check the PORT environment variable.
	if envPort := os.Getenv("PORT"); envPort != "" {
		if envPortInt, envPortErr := strconv.Atoi(envPort); envPortErr == nil {
			logger.Debug("Mounting service with an OS-environment port : %s", envPortInt)
			return envPortInt
		}
	}

	//Then, let's fallback to the config.
	return defaultPort
}

//Tells us if the port that is passed into the function is in use.
//Uses DialTCP; is probably not 100% deterministic, but should work in most cases.
func IsPortInUse(port int) bool {
	addressAndPort := net.JoinHostPort("localhost", strconv.Itoa(port))
	if addr, err := net.ResolveTCPAddr("tcp", addressAndPort); err == nil {
		if _, connErr := net.DialTCP("tcp", nil, addr); connErr == nil {
			logger.Error("Port %d is already in use!", port)
			return true
		}
	}
	return false
}

//Checks if the RunPort is in use.  If it is, it will let the system find us
// a port by bunding to port zero.
func GetSafeRunPort(defaultPort int) int {
	origRunPort := GetRunPort(defaultPort)
	if defaultPort == 8000 {
		for i := 8000; i <= 8005; i++ {
			if !IsPortInUse(i) {
				return i
			}
		}
		panic("can't find available port for router")
	} else {
		if IsPortInUse(origRunPort) {
			if addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0"); err == nil {
				if i, listenErr := net.ListenTCP("tcp", addr); listenErr == nil {
					defer i.Close()
					_, port, _ := net.SplitHostPort(i.Addr().String())
					if portInt, strErr := strconv.Atoi(port); strErr == nil {
						return portInt
					}
				}
			}
		} else {
			return origRunPort
		}

		return 0
	}
}

//Convenience function that returns the port in string form.
func GetSafeRunPortString(defaultPort int) string {
	return strconv.Itoa(GetSafeRunPort(defaultPort))
}

//Convenience function that returns the port in string form.
func GetSafeRunPortStringFromString(defaultPortStr string) string {
	if defaultPort, convErr := strconv.Atoi(defaultPortStr); convErr == nil {
		return strconv.Itoa(GetSafeRunPort(defaultPort))
	} else {
		return ""
	}
}
