package couchdbgo

// Connection information for couch db
type connection struct {
	host     string
	port     string
	password string
	username string
	protocol string
}

var connectionData connection

// Config function to activate connection data for couch database
// [host: <>,  port: <>, protocol: [], password: [], username: []]
func Config(conf map[string]string) {
	for key, value := range conf {
		switch key {
		case "host":
			connectionData.host = value
		case "port":
			connectionData.port = value
		case "password":
			connectionData.password = value
		case "username":
			connectionData.username = value
		case "protocol":
			connectionData.protocol = value
		}
	}
}
