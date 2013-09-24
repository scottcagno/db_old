// * 
// * Copyright 2013, Scott Cagno. All rights Reserved
// * License: sites.google.com/site/bsdc3license
// * 
// * -------
// * serv.go ::: database server
// * -------
// * 

package db

// server
type Server struct {
	m 	*Manager
}

// initialize server
func InitServer() *Server {
	return &Server{
		m: InitManager(),
	}
}

// listen and serve clients
func (self *Server) Serve() {
	addr, err := net.ResolveTCPAddr("tcp", LISTENING_PORT)
	if err != nil {
		log.Panicln(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Panicln(err)
	}
	for {
		c, err := l.AcceptTCP()
		if err != nil {
			log.Panicln(err)
		}
		go self.handleClient(c)
	}
}

// handle client connection
func (self *Server) handleClient(c *net.TCPConn) {
	// open new reader, and set client timeout
	r := bufio.NewReader(c)
	self.extend(c, CLIENT_TIMEOUT)
	for {
		// read bytes int b
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			self.close(c)
			return
		} else {
			self.extend(c, CLIENT_TIMEOUT)
		}
		// parse read bytes
		cmd, store, args := parse(b)
		switch cmd {
		case "":
			// send error
			send(c, []byte("-> empty"))
		case "help":
			// help menu
			send(c, []byte("-> help"))
		case "exit":
			// close connection
			send(c, []byte("-> exit"))
			c.SetDeadline(time.Now())
		case "rand":
			// send random hash
			send(c, random())
		case "save":
			// save archive
			send(c, []byte("-> save"))
		case "load":
			// load archive
			send(c, []byte("-> load"))
		case "incr":
			// incrment store pk
			if args == nil && self.m.Has(store) == []byte(1) {
				b := self.m.Stores[store].NextPK()
				send(c, b)
				break
			}
			// reset store pk
			if args != nil && self.m.Has(store) == []byte(1) {
				if args[0] == []byte("reset") {
					b := self.m.Stores[store].ResetPK()
					send(c, b)
					break
				}
				send(c, []byte("-> pk err"))
			}
			send(c, []byte("-> save"))
		case "has":
			// has comment
			send(c, []byte("-> has"))
		case "add":
			// add comment
			send(c, []byte("-> add"))
		case "set":
			// set comment
			send(c, []byte("-> set"))
		case "get":
			// get comment
			send(c, []byte("-> get"))
		case "del":
			// del comment
			send(c, []byte("-> del"))
		case "exp":
			// exp comment
			send(c, []byte("-> exp"))
		case "ttl":
			// ttl comment
			send(c, []byte("-> ttl"))
		case "loc":
			// loc comment
			send(c, []byte("-> loc"))
		default:
			// send error
			send(c, []byte("-> err"))
		}
	}
}

// parse bytes into easy to use chunks
func parse(b []byte) (string, string, [][]byte) {
	b = bytes.TrimRight(b, "\r\n")
	if b != nil {
		arg := bytes.Split(b, []byte(" "))
		if len(arg) == 1 {
			return string(arg[0]), "", nil
		}
		if len(arg) == 2 {
			return string(arg[0]), string(arg[1]), nil	
		}
		return string(arg[0]), string(arg[1]), arg[2:]
	}
	return "", "", nil
}

// send byte data out to client
func send(c *net.TCPConn, b []byte) {
	c.Write(b=append(b, '\n'))
}


// extend connection ttl
func (self *Server) extend(c *net.TCPConn, ttl int64) {
	if ttl > 0 {
		c.SetDeadline(time.Now().Add(time.Duration(ttl) * time.Second))
	}
}

// close connection
func (self *Server) close(c *net.TCPConn) {
	c.Write([]byte("Goodbye\r\n"))
	log.Printf("CLOSED CONNECTION TO CLIENT [%s]\n", c.RemoteAddr().String())
	c.Close()
	c = nil
}

