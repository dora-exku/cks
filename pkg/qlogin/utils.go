package qlogin

func Hash33(t string) int {
	var e int
	for _, item := range t {
		e += (e << 5) + int(item)
	}
	return 2147483647 & e
}

func Gtk(pskey string) int {
	hash := 5381
	for _, item := range pskey {
		hash += (hash << 5) + int(item)
	}
	return hash & 0x7fffffff
}
