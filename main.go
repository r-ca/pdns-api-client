package main

import (
    "fmt"
    "gopkg.in/ini.v1"
)

type Config struct {
    Host string
    ApiKey string
    HostName string
    Comment string
    TTL int
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
}


