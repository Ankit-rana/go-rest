package main

import (
    "os/exec"
    "strings"
)

type Interface struct {
        Name string
}

func (i Interface) show() ([]string, error){
        /*
         *run ifconfig and find all interface Name
        */
        cmd := exec.Command("../scripts/showInterface.sh")
        stdout,err := cmd.Output()
        if err != nil {
                return nil,err
        }
        s1 := strings.Trim(string(stdout),"\n")
        s := strings.Split(string(s1), "\n")
        return s, nil
}

