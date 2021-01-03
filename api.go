package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func getIPs() (ip string) {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	str := string(getIPs())
	fmt.Fprintf(w,"go crazy 2 " + str)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
