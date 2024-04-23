package main

import (
    "fmt"
    "net"
    "gopkg.in/ini.v1"
)

type Config struct {
    Host string
    ApiKey string
    HostName string
    Comment string
    TTL int
    Interface string
}

var config Config
const IS_DEBUG = true

func init() {
    // Load config
    cfg, err := ini.Load("conf/config.ini")
    if err != nil {
        panic(err)
    }
    
    config = Config{
        Host: cfg.Section("dns_server").Key("host").String(),
        ApiKey: cfg.Section("dns_server").Key("key").String(),
        HostName: cfg.Section("me").Key("hostname").String(),
        Comment: cfg.Section("me").Key("comment").String(),
        TTL: cfg.Section("me").Key("ttl").MustInt(),
        Interface: cfg.Section("me").Key("interface").String(),
    }
}

func logger(msg string) {
    if IS_DEBUG {
        fmt.Println(msg)
    }
}

func main() {
    logger(fmt.Sprintf("Host: %s", config.Host))
    logger(fmt.Sprintf("ApiKey: %s", config.ApiKey))
    logger(fmt.Sprintf("HostName: %s", config.HostName))
    logger(fmt.Sprintf("Comment: %s", config.Comment))
    logger(fmt.Sprintf("TTL: %d", config.TTL))
    logger(fmt.Sprintf("Interface: %s", config.Interface))

    logger(getMyIp())
}

func getMyIp() string {
    iface, err := net.InterfaceByName(config.Interface)
    if err != nil {
        panic(err)
    }

    if iface.Flags&net.FlagUp == 0 {
        panic("Interface is down")
    }

    // インタフェースに設定されているアドレスを取得
    addrs, err := iface.Addrs()
    if err != nil {
        panic(err)
    }
    for _, addr := range addrs {
        switch v := addr.(type) {
        case *net.IPNet:
            if v.IP.To4() != nil {
            return v.IP.String()
            }
        }
    }
    
    return ""
}
