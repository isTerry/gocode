//1.Multiple packages in directory: main, checks。一个目录内不能有不同的包
//2.package中的名称最好与目录名一致，直接import目录名。如果不一致，那么import上层的目录
//3.不允许Main file has non-main package or doesn't contain main function
//4.不能从其他包内import main包，相反可以，且其他包之间可以import
package checks

import (
	"fmt"
)

type doctor struct {
	Number     int
	DoctorName string
	episodes   []string
	companions []string
}
type Doctor doctor

func main() {
	fmt.Printf("checks, world\n")
}
