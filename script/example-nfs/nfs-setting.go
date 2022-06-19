package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"os"
)

func Exec(command string) error {
	// http://docscn.studygolang.com/pkg/bytes/#NewBuffer
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("bash")
	cmd.Stdin = in
	in.WriteString(command)
	//in.WriteString("exit\n")

	if err := cmd.Run();
	err != nil{
		return err
	}

	return nil
}

func  Cmdinit(parentCommand string) string{
	var path = flag.String("path","-h","please input the new nfs path")
	flag.Parse()
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	flag.PrintDefaults()

	var cmd = parentCommand + " "+ *path
	return cmd
}

func Inputfile(filename string, input string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0655)
	if err != nil{
		fmt.Printf("%v文件打开失败，错误为：%v\n", filename, err)
		return nil
	}

	defer file.Close()

	file.Write([]byte(input))
	file.WriteString("dddddd")

	return nil
}

func main() {

	Inputfile("/tmp/test.txt", "hellp")
	/*
	var cmd = Cmdinit("echo ")
	fmt.Println(cmd)

	if err := Exec(cmd)
	err != nil{
		fmt.Println("error")
	}

	 */
}
