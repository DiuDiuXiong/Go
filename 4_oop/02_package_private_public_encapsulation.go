package tree

import "fmt"

/**
- first character upper case: public
- first character lower case: private
*/

/** About package
1. Every directory can have only one package
2. main package contains main function which is starting point of execution
3. functions for a struct must within same package (see canDefineHere), but can define functions for struct on another file
4. import other file from other directory, then <package_name>.<struct_name> to access

*/

func (node TreeNode) canDefineHere() {
	fmt.Println("Do nothing here.")
}
