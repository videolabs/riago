package riago

// Performs a Riak Ping request.
func (c *Conn) Ping() error {
	return c.do(MsgRpbPingReq, nil, nil)
}

// Performs a Riak Server info request.
func (c *Conn) ServerInfo() (resp *RpbGetServerInfoResp, err error) {
	resp = new(RpbGetServerInfoResp)
	err = c.do(MsgRpbGetServerInfoReq, nil, resp)
	return
}