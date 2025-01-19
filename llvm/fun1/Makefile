#ir:
#	clang -emit-llvm -S fun1.c -o fun1.ll

irO2:
	clang -emit-llvm -S -O2 fun1.c -o fun1.ll

llvm-as:
	llvm-as fun1.ll -o fun1.bc

asm:
	clang -S fun1.bc -o fun1.s

clean:
	rm *.ll *.bc *.s