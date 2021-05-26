## 1. qsort

1. 定义 `stdlib.h -> qsort` 中调用时的 cmp 函数类型 `typedef int (*qsort_cmp_func_t)(const void* a,const void*b);`
1. 申明 cmp 函数 `int _cgo_qsort_compare(void* a,void* b);` 未来会转化为 1 中 cmp
1. 使用 go 语言，全局变量的方式，实现 2 中声明函数的实现，此时, 实现是通过闭包实现，此时函数只有声明，没有执行地址
1. 通过反射限制，传入的类型必须为 slice, 并取到 slice 的首个元素地址、元素个数、元素大小，保存到全局闭包变量中，暂存
1. 调用 `stdlib.h -> qsort`  实现数组排序

## 2. 调用方

1. cmp 函数的具体实现
1. 将 cmp 的实现和 上3 中的全局闭包变量绑定
1. 调用 `qsort` 包中的排序算法，实现排序

## 3. 原因

1. go 类型函数无法直接转为 `stdlib.h -> qsort` 需要的比较函数类型, `typedef int (*qsort_cmp_func_t)(const void* a,const void*b);`
