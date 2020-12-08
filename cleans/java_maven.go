package cleans

import (
	"fmt"
	"os"
	"path"
)

type JavaMavenCleaner struct {
	ProjectType
}

func (jm JavaMavenCleaner) MarkAndRemove(p string, children *[]os.FileInfo) int {
	markResult := false
	for i := 0; i < len(*children); i++ {
		pom := (*children)[i]
		if pom.Name() == "pom.xml" {
			Remove(children, "target")
			fmt.Printf("%s 是一个java maven 工程 移除target目录 \n", p)
			markResult = true
			break
		}
	}
	if markResult {
		return JAVA_MAVEN
	}
	return NOTHING
}

func (jm JavaMavenCleaner) DoClean(p string) {
	err := os.RemoveAll(path.Join(p, "target"))
	if err != nil {
		fmt.Println("JavaMavenCleaner remove error " + err.Error())
	}
}
