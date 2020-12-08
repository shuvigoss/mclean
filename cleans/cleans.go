package cleans

import (
	"os"
)

const (
	NOTHING         int = 0 // 什么都不是
	JAVA_MAVEN      int = 1 // maven工程
	JAVASCRIPT_NODE int = 2 // node
)

var cleaners = make([]Cleaner, 0)

type Cleaner interface {
	MarkAndRemove(p string, children *[]os.FileInfo) int
	DoClean(p string)
}

type ProjectType struct {
	Tp int
}

func GetCleaners() []Cleaner {
	return cleaners
}

func init() {
	cleaners = append(cleaners, new(JavaMavenCleaner))
	cleaners = append(cleaners, new(JavaScriptBodeCleaner))
}

func Remove(children *[]os.FileInfo, name string) {
	for i := 0; i < len(*children); i++ {
		pom := (*children)[i]
		if pom.Name() == name {
			*children = append((*children)[:i], (*children)[i+1:]...)
			i--
			break
		}
	}
}
