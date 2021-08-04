#include "./number.h"

/*
生成动态库:
	gcc -shared -o libnumber.so number.c
*/
int number_add_mod(int a, int b, int mod){
	return (a+b)%mod;
}