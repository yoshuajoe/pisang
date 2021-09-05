# pisang
Pisang is fun. This is a simple interpreter, just for fun, hobby and upskilling.

## How to build
Please use this command to build Pisang to your machine
```
  make buildexec
```

## Current working syntax
```
alas := 9;
tinggi := 5;
luas := alas*tinggi/2;
assert luas;

siswaDalamKelas := ["Aldo", "Aldi", "Andi", "Ali", "Ando"];
if siswaDalamKelas[1] == "Aldi"{
    assert "Halo " + siswaDalamKelas[1] + ", salam kenal!";
}

assert siswaDalamKelas[0];

```
Yeah I know this is still needing semicolon and flying old-style Pascal-like. Anyway, if you're interested to contribute please create an issues then fork it.
Please see `bin/sample.pi`

## To do 
1. Unary operator
2. Tidying the codes
3. Argparse
4. Procedure/function
5. Adding abstract syntax tree (AST)
6. Variable type declaration
7. Other primitive types (currently I'm only working on integer data type)
... (will be updated soon)
