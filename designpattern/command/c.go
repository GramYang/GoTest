package main

import "fmt"

//命令模式就是套了多层的adapter模式
func main() {
	var tv TV
	openCommand := OpenCommand{tv}
	invoker := Invoker{openCommand}
	invoker.Do()
	invoker.SetCommand(CloseCommand{tv})
	invoker.Do()
}

type TV struct{}

func (p TV) Open() {
	fmt.Println("Play...")
}

func (p TV) Close() {
	fmt.Println("Stop...")
}

type Command interface {
	Press()
}

type OpenCommand struct {
	tv TV
}

func (p OpenCommand) Press() {
	p.tv.Open()
}

type CloseCommand struct {
	tv TV
}

func (p CloseCommand) Press() {
	p.tv.Close()
}

type Invoker struct {
	cmd Command
}

func (p *Invoker) SetCommand(cmd Command) {
	p.cmd = cmd
}

func (p *Invoker) Do() {
	p.cmd.Press()
}
