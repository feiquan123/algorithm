#include "number.h"

#include <stdio.h>

// gcc -o a.out _test_main.c number.so && LD_LIBRARY_PATH=. ./a.out && rm -rf a.out
int main(){
	int a=10;
	int b=5;
	int c = 12;
	int x = number_add_mod(a,b,c);
	printf("(%d+%d)%%%d = %d\n",a,b,c,x);
	return 0;
}