/*
GO->C:
	go build -buildmode=c-shared -o person.so object.go person.go person_capi.go
RUN: 
    mac:   g++ -std=c++11  main.cpp ./person.so -o person && ./person && rm -f ./person
    linux: g++ -std=c++0x  main.cpp ./person.so -o person && ./person && rm -f ./person
*/
#include "person_class.h"

#include <stdio.h>

// // v1
// int main(){
// 	auto p = new Person("gopher",10);

// 	char buf[64];
// 	char* name = p->GetName(buf,sizeof(buf)-1);
// 	int age = p->GetAge();

// 	printf("%s, %d years old.\n",name,age);
// 	delete p;

// 	return 0;
// }

// v2
int main(){
	Person person;
	auto p = person.New("gopher",10);

	char buf[64];
	char* name = p->GetName(buf,sizeof(buf)-1);
	int age = p->GetAge();

	printf("%s, %d years old.\n",name,age);
	p->Delete();

	return 0;
}