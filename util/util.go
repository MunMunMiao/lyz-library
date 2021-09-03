package util

import (
    "errors"
    "net"
)

func GetAddr() (*net.UDPAddr, error) {
    conn, err := net.Dial("udp", "1.1.1.1:80")
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
    if !ok{
        return nil, errors.New("can not find net addr")
    }

    return localAddr, nil
}

