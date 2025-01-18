gen:
	gcc -S -O2 hello.c -o hello.s

output:
	as -o hello.o hello.s

run:
	gcc -o hello hello.o && ./hello

clean:
	rm -f hello.s hello.o hello