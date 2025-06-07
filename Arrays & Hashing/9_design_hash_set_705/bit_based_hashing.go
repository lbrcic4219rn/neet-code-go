package __design_hash_set_705

type BitBasedSet struct {
	data []uint64
}

func Constructor2() BitBasedSet {
	size := (1000000 >> 6) + 1 // 1000000 / 64 = 15625 + 1
	return BitBasedSet{
		data: make([]uint64, size),
	}
}
