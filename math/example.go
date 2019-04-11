package main

import (
	"math"
)

// math包提供了基本的数学常数和数学函数
func main() {

	// 返回一个NaN, IEEE 754“这不是一个数字”值
	math.NaN()

	// 判断f是否表示一个NaN（Not A Number）值
	math.IsNaN(12.34)

	// 如果sign>=0函数返回正无穷大，否则返回负无穷大
	f := math.Inf(0)

	// 如果sign > 0，f是正无穷大时返回真；如果sign<0，f是负无穷大时返回真；sign==0则f是两种无穷大时都返回真
	math.IsInf(12.34, 1)
	math.IsInf(f, 0)

	// 返回浮点数f的IEEE 754格式二进制表示对应的4字节无符号整数
	// float32转uint32
	u32 := math.Float32bits(12.34)

	// 返回无符号整数b对应的IEEE 754格式二进制表示的4字节浮点数
	// uint32转float32
	math.Float32frombits(u32)

	// 返回浮点数f的IEEE 754格式二进制表示对应的8字节无符号整数
	// float64转uint64
	u64 := math.Float64bits(12.34)

	// 返回无符号整数b对应的IEEE 754格式二进制表示的8字节浮点数
	// uint64转float64
	math.Float64frombits(u64)

	// 如果x是一个负数或者负零，返回真
	// 判断是否小于0
	math.Signbit(-1)

	// 返回x的绝对值与y的正负号的浮点数
	math.Copysign(-12.34, -22.54) // -12.34

	// 返回不小于x的最小整数（的浮点值）
	// 计算数字只入不舍
	math.Ceil(12.34)

	// 返回不大于x的最小整数（的浮点值）
	// 计算数字只舍不入
	math.Floor(12.34)

	// 四舍五入
	math.Round(12.34)

	// 返回x的整数部分（的浮点值）
	math.Trunc(12.34)

	// 返回f的整数部分和小数部分，结果的正负号都相同
	math.Modf(12.34)

	// 参数x到参数y的方向上，下一个可表示的数值；如果x==y将返回x
	math.Nextafter(12.34, 15.67)

	// 同Nextafter类似
	math.Nextafter32(12.34, 15.67)

	// 返回x的绝对值
	math.Abs(-12.34) // 12.34

	// 返回x和y中最大值
	math.Max(12.3, 22)

	// 返回x和y中最小值
	math.Min(12.3, 22)

	// 返回x-y和0中的最大值
	math.Dim(12.3, 12)

	// 取余运算，可以理解为 x-Trunc(x/y)*y，结果的正负号和x相同
	math.Mod(22, 3)

	// IEEE 754差数求值，即x减去最接近x/y的整数值（如果有两个整数与x/y距离相同，则取其中的偶数）与y的乘积
	math.Remainder(12, 34)

	// 返回x的二次方根
	math.Sqrt(144) // 12

	// 返回x的三次方根
	math.Cbrt(27)

	// 返回Sqrt(p*p + q*q)，注意要避免不必要的溢出或下溢
	math.Hypot(2, 3)

	// 求正弦
	math.Sin(3)

	// 求余弦
	math.Cos(4)

	// 求正切
	math.Tan(5)

	// 函数返回Sin(x), Cos(x)
	math.Sincos(3)

	// 求反正弦（x是弧度）
	// Asin(±0) = ±0
	// Asin(x) = NaN if x < -1 or x > 1
	math.Asin(0.5)

	// 求反余弦（x是弧度）
	// Acos(x) = NaN if x < -1 or x > 1
	math.Acos(0.5)

	// 求反正切（x是弧度）
	// Atan(±0) = ±0
	// Atan(±Inf) = ±Pi/2
	math.Atan(0.5)

	// 类似Atan(y/x)，但会根据x，y的正负号确定象限
	math.Atan2(0.5, 0.2)

	// 求双曲正弦
	math.Sinh(2)

	// 求双曲余弦
	math.Cosh(2)

	// 求双曲正切
	math.Tanh(2)

	// 求反双曲正弦
	math.Asinh(2)

	// 求反双曲余弦
	math.Acosh(2)

	// 求反双曲正切
	math.Atanh(2)

	// 求自然对数
	math.Log(2)

	// 等价于Log(1+x)。但是在x接近0时，本函数更加精确
	math.Log(0.0001)

	// 求2为底的对数；特例和Log相同
	math.Log2(2)

	// 求10为底的对数；特例和Log相同
	math.Log10(10)

	// 返回x的二进制指数值，可以理解为Trunc(Log2(x))
	math.Logb(10)

	// 类似Logb，但返回值是整型
	math.Ilogb(12.34)

	// 返回一个标准化小数frac和2的整型指数exp，满足f == frac * 2**exp，且0.5 <= Abs(frac) < 1
	frac, exp := math.Frexp(12.34)

	// Frexp的反函数，返回 frac * 2**exp
	math.Ldexp(frac, exp)

	// 返回E**x；x绝对值很大时可能会溢出为0或者+Inf，x绝对值很小时可能会下溢为1
	math.Exp(12)

	// 等价于Exp(x)-1，但是在x接近零时更精确；x绝对值很大时可能会溢出为-1或+Inf
	math.Expm1(12)

	// 返回2**x
	math.Exp2(12)

	// 返回x**y
	math.Pow(3, 4)

	// 返回10**e
	math.Pow10(2)

	// 伽玛函数（当x为正整数时，值为(x-1)!）
	math.Gamma(10)

	// 返回Gamma(x)的自然对数和正负号
	math.Lgamma(10)

	// 误差函数
	math.Erf(10.01)

	// 余补误差函数
	math.Erfc(10.01)

	// 第一类贝塞尔函数，0阶
	math.J0(2)

	// 第一类贝塞尔函数，1阶
	math.J1(2)

	// 第一类贝塞尔函数，n阶
	math.Jn(2, 2)

	// 第二类贝塞尔函数，0阶
	math.Y0(2)

	// 第二类贝塞尔函数，1阶
	math.Y1(2)

	// 第二类贝塞尔函数，n阶
	math.Yn(2, 2)
}
