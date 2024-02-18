package tree

import (
	_ "io/ioutil"
	_ "os"
	_ "path/filepath"
)

type FloderNode struct {
	Name     string
	IsDir    bool
	Children []FloderNode
}
