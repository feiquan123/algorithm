// Package rsa models
package rsa

/*
gcd 是求最大公约数

rsa 算法:
	1. 选择两个 p,q 素数，n = p * q
	2. 计算 Q = (p-1) * (q-1)
	3. 找出满足条件的e:
		1 < e < Q, gcd(e,Q) = 1
		方法:
			在 (1,Q) 中选择一个素数, 使 gcd(e,Q) = 1, 既使得 gcd(e,p-1) =1 && gcd(e,q-1) = 1
		推论:
			gcd(e,Q) = 1 	=> 	gcd(e,p-1) = 1 && gcd(e,q-1) = 1
			如:	gcd(4,7) = 1 && gcd(4,9) = 1 	=> 	gcd(4,36)
	4. 找出满足条件的d
		1< d < Q, ed % Q = 1 	=> 	1< Q/e < d < Q, ed % Q = 1
		方法:
			 => (ed - 1)  / Q = 0
	5. 公钥: (n,e) , 私钥: (n,d) . p,q,Q 的值不公开

加密:
	c = exp(m,e) % n

解密:
	m = exp(c,d) % n
*/
