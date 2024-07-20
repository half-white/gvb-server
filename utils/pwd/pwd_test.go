package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$p2fkjuHmNGlfa.suyzGe8.HtnsPphpyXbQKgXlCXslEfjC9Sn8c8y", "1234"))
}
