package toy

func Tiling(n int) int {
	/*
		세로 길이 2, 가로 길이 n인 2 x n 보드가 있습니다. 2 x 1 크기의 타일을 가지고 이 보드를 채우는 모든 경우의 수를 리턴해야 합니다.
	*/
	m := []int{1, 2} // 1칸이면1 2칸이면11, 2 => 2, 3 => tilint(1) + tiling(2) => fibo
	if n == 1 || n == 2 {
		return m[n-1]
	}
	return Tiling(n-1) + Tiling(n-2)
}

func Fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	fibo := Fibo(n-1) + Fibo(n-2)
	return fibo
}
