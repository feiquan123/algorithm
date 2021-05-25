#include <stdio.h>
#include <stdlib.h>

#define DIM(x)  (sizeof(x)/sizeof((x)[0]))  // 取数组 x 的长度

static int cmp(const void* a, const void* b){
    const int* pa = (int*) a;
    const int* pb = (int*) b;
    return *pa - *pb;
}

// gcc main.c -o main && ./main && rm main
int main(){
    int values[] =  {42, 8, 109, 97, 23, 25};
    int i;
    qsort(values,DIM(values),sizeof(values[0]),cmp);

    for(i = 0;i < DIM(values); i++){
        printf("%d ",values[i]);
    }
    return 0;
}