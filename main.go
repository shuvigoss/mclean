package main

import (
	"fmt"
	"io/ioutil"
	"mclean/cleans"
	"path"
)

func main() {
	wd := "/Users/shuvigoss/Documents/bjca/front"
	list(path.Join(wd))

}

func list(parentPath string) {
	if parentPath == "" {
		fmt.Println("parentPath 为空")
		return
	}

	fmt.Println(parentPath)

	fileInfos, err := ioutil.ReadDir(parentPath)
	if err != nil {
		fmt.Printf("读取目录 %s 异常 %s", parentPath, err.Error())
		return
	}

	cleaners := cleans.GetCleaners()
	for _, c := range cleaners {
		t := c.MarkAndRemove(parentPath, &fileInfos)
		if t > cleans.NOTHING {
			c.DoClean(parentPath)
		}
	}

	for i := range fileInfos {
		var f = path.Join(parentPath, fileInfos[i].Name())
		if fileInfos[i].IsDir() {
			list(f)
		} else {
			fmt.Println(f)
		}

	}
}
