package main

import (
    "fmt"
    "net/http"
    "net"
    "strings"
)

func ReverseIPAddress(ip net.IP) string {
    addressSlice := strings.Split(ip.String(), ".")
    reverseSlice := []string{}
    for i := range addressSlice {
        octet := addressSlice[len(addressSlice)-1-i]
            reverseSlice = append(reverseSlice, octet)
    }
    return strings.Join(reverseSlice, ".")
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        ipAddress := net.ParseIP(r.FormValue("ip"))
        if ipAddress.To4() != nil {
            reverseIpAddress := ReverseIPAddress(ipAddress)
            ip, err := net.LookupIP(reverseIpAddress + ".score.senderscore.com")
            if err != nil {
                http.Error(w, "Your IP hasn't sent out enough email for senderscore to do a measure.", http.StatusNotFound)
                return
            } else {
                fmt.Fprint(w, ip[0][3], "\n")
            }
        } else {
            http.Error(w, "Invalid IP-address supplied.", http.StatusInternalServerError)
        }
    })
    http.ListenAndServe(":80", nil)
}
