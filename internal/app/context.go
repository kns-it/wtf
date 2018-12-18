package app

type svcOrPortAssignmentError struct {
	msg string
}

func (err svcOrPortAssignmentError) Error() string {
	return err.msg
}

type Service int16
type Port int16

type ServiceOrPort struct {
	port int16
}

func (sop *ServiceOrPort) SetService(svc Service) {
	sop.port = int16(svc)
}

func (sop *ServiceOrPort) SetPort(port Port) {
	sop.port = int16(port)
}

type CommandContext struct {
	host string
	port ServiceOrPort
}

func (ctx CommandContext) GetHost() string {
	return ctx.host
}

func (ctx CommandContext) GetPort() int16 {
	return ctx.port.port
}

func (ctx *CommandContext) SetServiceOrPort(port ServiceOrPort) error {
	if port.port != 0 {
		ctx.port = port
		return nil
	}

	return svcOrPortAssignmentError{
		msg: "Port to assign was the default port",
	}
}

func (ctx *CommandContext) SetHost(host string) {
	ctx.host = host
}