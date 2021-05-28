#include <iostream>
using namespace std;

extern "C" {
	#include "hello.h"
}

void SayHelloCPP(char* s){
	cout << s << endl;
}