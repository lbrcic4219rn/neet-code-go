package __design_hash_set_705

type MyHashSet struct {
	buckets    [][]bool
	bucketSize int
}

func Constructor() *MyHashSet {
	size := 1000
	return &MyHashSet{
		buckets:    make([][]bool, size),
		bucketSize: size,
	}
}

func (s *MyHashSet) getHash(key int) (int, int) {
	return key % s.bucketSize, key / s.bucketSize
}

func (s *MyHashSet) Add(key int) {
	bucketIndex, indexInBucket := s.getHash(key)
	if s.buckets[bucketIndex] == nil {
		if bucketIndex == 0 {
			s.buckets[bucketIndex] = make([]bool, s.bucketSize+1)
		} else {
			s.buckets[bucketIndex] = make([]bool, s.bucketSize)
		}
	}
	s.buckets[bucketIndex][indexInBucket] = true
}

func (s *MyHashSet) Remove(key int) {
	bucketIndex, indexInBucket := s.getHash(key)
	if s.buckets[bucketIndex] != nil {
		s.buckets[bucketIndex][indexInBucket] = false
	}
}

func (s *MyHashSet) Contains(key int) bool {
	bucketIndex, indexInBucket := s.getHash(key)
	return s.buckets[bucketIndex] != nil && s.buckets[bucketIndex][indexInBucket]
}
