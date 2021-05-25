#include <stdio.h>

extern "C" {
	#include "sum.h"
}

// 暂时不可编译
int main(){
    printf("%d",sum(1,2));
    return 0;
}