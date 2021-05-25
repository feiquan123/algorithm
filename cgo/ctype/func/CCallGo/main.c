#include <stdio.h>
#include "sum.h"


// go build -buildmode=c-shared -o sum.so main.go
// gcc main.c ./sum.so -o main && ./main && rm main
int main(){
    printf("%d\n",sum(1,2));
    return 0;
}