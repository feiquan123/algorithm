#include <stdio.h>

struct Int{
	int Twice(){
		const int* p = (int*) (this);
		return (*p)*2;
	}
};

int main(){
	int x = 42;
	printf("%d\n",x);
	printf("%d\n",((Int*)(&x))->Twice());
}