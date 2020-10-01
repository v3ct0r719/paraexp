package main

import (
	"fmt"
	"regexp"
)

func flag_submit(flags string, io *Connection, regex string) {

	re := regexp.MustCompile(regex)
	fl := re.FindAllStringSubmatch(flags, -1)

	for _, i := range fl {

		io.sendline(i[0])
		a, _ := io.recvline()
		fmt.Println("This is what I am printing")
		fmt.Println(string(a))

	}

}
