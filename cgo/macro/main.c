#include <stdio.h>

#define  CGO_OS_DARWIN 1

#if defined(CGO_OS_WINDOWS)
	static const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
	static const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
	static const char* os = "linux";
#else
    static const char* os = "unkown os";
// #   error("unkown os");
#endif

int main(){
    puts(os);
    return 0;
}
