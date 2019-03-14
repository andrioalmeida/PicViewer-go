package main

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"os"
	"picViewer/platforms"
	"picViewer/renderers"
)

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()
	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		os.Exit(-1)
	}
	Run(platform, renderer)
}
