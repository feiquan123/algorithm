#include "number.h"

#include <stdio.h>

// gcc -lpthread -o a.out _test_main.c number.a && ./a.out && rm -f a.out
int main(){
	int a=10;
	int b=5;
	int c = 12;
	int x = number_add_mod(a,b,c);
	printf("(%d+%d)%%%d = %d\n",a,b,c,x);
	return 0;
}