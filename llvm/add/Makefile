llvm-as:
	llvm-as add_fn.ll -o add_fn.bc

opt:
	opt -O2 add_fn.bc -o add_fn_O2.bc

dis:
	llvm-dis add_fn_O2.bc -o add_fn_O2.ll

asm:
	llc add_fn_O2.bc -o add_fn.s

clean:
	rm *.ll *.bc *.s