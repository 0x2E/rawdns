# rawdns

DNS messages (un)marshaller

## Overview

This project is forked from [cirocosta/rawdns](https://github.com/cirocosta/rawdns) and has been modified to make it more suitable for constructing and parsing the UDP packet of DNS.

For complete DNS functionality, see [miekg/dns](https://github.com/miekg/dns).

## Example

```go
package main

import (
	"github.com/0x2E/rawdns"
	"net"
)

func main() {
	// create socket
	conn, _ := net.Dial("udp", "8.8.8.8:53")
	defer conn.Close()

	// construct DNS packet content
	payload, _ := rawdns.Marshal(33, 1, "github.com", rawdns.QTypeA)

	// send UDP packet
	_, _ = conn.Write(payload)

	// receive UDP packet
	buf := make([]byte, 0, 1024)
	n, _ := conn.Read(buf)

	// parse
	resp, _ := rawdns.Unmarshal(buf[:n])
}
```
