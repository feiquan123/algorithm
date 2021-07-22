/* my_buffer.h 

RUN: 
    mac:   g++ -std=c++11  my_buffer.cpp  && ./a.out && rm -f ./a.out
    linux: g++ -std=c++0x  my_buffer.cpp  && ./a.out && rm -f ./a.out
*/
#include <string>

struct MyBuffer
{
    std::string *s_;

    MyBuffer(int size)
    {
        this->s_ = new std::string(size, char('\0'));
    }

    ~MyBuffer()
    {
        delete this->s_;
    }

    int Size() const
    {
        return this->s_->size();
    }

    char *Data()
    {
        return (char *)this->s_->data();
    }
};

// int main()
// {
//     auto pBuf = new MyBuffer(1 << 10);

//     auto data = pBuf->Data();
//     auto size = pBuf->Size();
//     printf("%d\n",size);

//     delete pBuf;
// }