package main

import (
	"github.com/flint92/play-with-compiler/craft/script"
	"os"
)

func main() {

	///
	//>>2;
	//2
	//>>2+
	//	3*5
	//;
	//17
	//>>age;
	//unknown variable: age
	//>>int age = 45;
	//45
	//>>int newAge = age + 10 * 2;
	//65
	//>>newAge;
	//65
	//>>exit();
	//Bye!
	///
	script.NewScript().Run(os.Stdin, os.Stdout)
}
