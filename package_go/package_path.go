package packagego

import (
	"fmt"
	"path"
	"path/filepath"
)

func MainPackagePath() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "southeast"))
	fmt.Println(path.Join("hello", "world", "southeast", "base.go"))

	// Filepath is path support differen OS, it is commonly used to file system paths
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "southeast.go"))
}
