#include "./my_buffer.h"

extern "C" {
	#include "./my_buffer_capi.h"
}


// MyBuffer_T public 继承  MyBuffer
struct MyBuffer_T: MyBuffer{
	MyBuffer_T(int size): MyBuffer(size){} //有参构造函数
	~MyBuffer_T(){}
};

MyBuffer_T* NewMyBuffer(int size){
	auto p = new MyBuffer_T(size);
	return p;
}

void DeleteMyBufffer(MyBuffer_T* p){
	delete p;
}

char* MyBuffer_Data(MyBuffer_T* p){
	return p->Data();
}

int MyBuffer_Size(MyBuffer_T* p){
	return p->Size();
}
