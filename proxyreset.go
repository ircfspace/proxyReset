package main

import (
	"fmt"
	"os/exec"
)

func main() {
    cmd := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "0", "/f")
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error disabling proxy:", err)
        return
    }
    fmt.Println("Proxy successfully disabled.")
}