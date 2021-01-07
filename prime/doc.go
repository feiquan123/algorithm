// Package prime query prime from range
package prime

/*
1. 素数的定义
	- 素数又称为质数，指在大于1的自然数中，除了1和它本身以外不再有其他因数的自然数
	- 2是最小的质数

2. 算法
	1. 暴力枚举	IsPrimeByRangeEnum
	2. 开方	因数都是成对出现,
		比如，100的因数有：1和100、2和50、4和25、5和20、10和10
		即成对的因数，其中一个必然小于等于100的开平方，另一个大于等于100的开平方
		因此只要判断2~sqrt(n)的因数即可
	3. 埃拉托斯特尼(Eratosthenes)
		1. 创建长度为 (n -2 + 1) 的数组, 初始值都为 true
		2. 遍历[2-n] , 每次将 i 的倍数的下标设置为 false
			- 倍数筛选 [j=2, i*j<=n, j++]
			- 二次筛选 [j=i*i, j<=n, j+=i]
			- 线性筛选
		3. 最后遍历数组中为 true 的元素就是素数

3. 素数定理
	素数的分布是越往后越稀疏。或者说，素数的密度是越来越低。而素数定理，
	说白了就是数学家找到了一些公式，用来估计某个范围内的素数，大概有几个。
	在这些公式中，最简洁的就是 x/ln(x).假设要估计1,000,000以内有多少质数，
	用该公式算出是72,382个，而实际有78,498个，误差约8个百分点。
	该公式的特点是：估算的范围越大，偏差率越小。
	一般用x/ln x来估计某个范围内素数的个数(误差小于15%)
*/
