#include "./number.h"

/*
生成静态库:
	gcc -c -o number.o number.c
	ar rsc libnumber.a number.o
*/
int number_add_mod(int a, int b, int mod){
	return (a+b)%mod;
}