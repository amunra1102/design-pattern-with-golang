package prototype

import "fmt"

type Folder struct {
	Children []INode
	Name     string
}

func (f *Folder) Print(str string) {
	fmt.Println(str + f.Name)
	for _, i := range f.Children {
		i.Print(str + str)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildren []INode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}

