gen:
	gcc -S -O2 function_call.c -o function_call.s

output:
	as -o function_call.o function_call.s

run:
	gcc -o function_call function_call.o && ./function_call

clean:
	rm -f function_call.s function_call.o function_call