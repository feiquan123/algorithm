#include "my_main.h"

#include <stdio.h>

// gcc -lpthread -o a.out _test_main.c main.a && ./a.out && rm -f a.out
int main(){
	int a=10;
	int b=5;
	int c = 12;
	int x = number_add_mod(a,b,c);
	printf("(%d+%d)%%%d = %d\n",a,b,c,x);

	goPrintln("done");
	return 0;
}