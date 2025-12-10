package main

const fenwickInitVal = 0 // -1e18

type fenwick []int

func newFenwickTree(n int) fenwick {
	t := make(fenwick, n+1)
	for i := range t {
		t[i] = fenwickInitVal
	}
	return t
}

func (fenwick) op(a, b int) int {
	return a + b // max(a, b)
}

// a[i] 增加 val
// 1<=i<=n
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = f.op(f[i], val)
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1<=i<=n
func (f fenwick) pre(i int) int {
	res := fenwickInitVal
	i = min(i, len(f)-1)
	for ; i > 0; i &= i - 1 {
		res = f.op(res, f[i])
	}
	return res
}

// 求区间和 a[l] + ... + a[r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func main() {

}
