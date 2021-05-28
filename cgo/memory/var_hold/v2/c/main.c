# include "libgostring.h"


static void printString(char* s){
	char* gs = NewGoString(s);
	PrintGoString(gs);
	FreeGoString(gs);
}

int main(){
    printString("hello world\n");
    return 0;
}