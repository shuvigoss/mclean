package cleans

import (
	"fmt"
	"os"
	"path"
)

type JavaScriptBodeCleaner struct {
	ProjectType
}

func (jn JavaScriptBodeCleaner) MarkAndRemove(p string, children *[]os.FileInfo) int {
	for _, packageJson := range *children {
		if packageJson.Name() == "package.json" {
			fmt.Printf("%s 是一个javascript node 工程\n", p)
			Remove(children, "node_modules")
			return JAVASCRIPT_NODE
		}
	}
	return NOTHING
}

func (jn JavaScriptBodeCleaner) DoClean(p string) {
	err := os.RemoveAll(path.Join(p, "node_modules"))
	if err != nil {
		fmt.Println("JavaScriptBodeCleaner remove error " + err.Error())
	}
}
