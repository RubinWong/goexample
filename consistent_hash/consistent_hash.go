package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"os"
	"sort"
	"strconv"
	"unsafe"
)

const (
	defaultReplica = 100
	defaultWeight = 1
)

type HashRing []uint32

func (ring HashRing) Len() int {
	return len(ring)
}

func (ring HashRing) Less(i, j int) bool {
	return ring[i] < ring[j]
}

func (ring HashRing) Swap(i, j int) {
	ring[i], ring[j] = ring[j], ring[i]
}

type ConsistentHash struct{
	replica int;
	buckets map[uint32]string
	ring  HashRing
	hashFn func([]byte) uint32
}

func (h* ConsistentHash) AddNode(str string, weight int) {
	if weight <= 0 {
		weight = 1
	} else if weight >= 5 {
		weight = 5
	}
	for i := 0; i < h.replica * weight; i++ {
		vnode := str + strconv.Itoa(i)
		hash := h.hashFn([]byte(vnode))
		h.buckets[hash] = str
		h.ring = append(h.ring, hash)
	}

	sort.Sort(h.ring)
}

func (h* ConsistentHash) GetConsistentNode(str string) string {
	hash := h.hashFn([]byte(str))
	i := sort.Search(len(h.ring), func(i int) bool { return h.ring[i] >= hash })

	pos := i - 1
	if pos < 0 {
		pos = len(h.ring) - 1
	}

	// fmt.Println(h.ring)
	// fmt.Println("value of pos , i ", "  ", pos, ", ", i, " hash: ", hash, " len of ring: ", len(h.ring))
	return h.buckets[h.ring[pos]]
}


func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage count mur|ieee|sha1")
		return
	}

	count,_ := strconv.ParseInt(os.Args[1], 10, 32)
	hashtab  := &ConsistentHash {
		replica: int(count),
		buckets: make(map[uint32]string),
	}

	if os.Args[2] == "mur" {
		hashtab.hashFn = hash_mur32
	} else if os.Args[2] == "ieee" {
		hashtab.hashFn = hash_ieee
	} else {
		hashtab.hashFn = hash_sha1
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			hashtab.AddNode("stream_tab" + strconv.Itoa(i), 2)
		} else {
			hashtab.AddNode("stream_tab" + strconv.Itoa(i), 0)
		}
		// fmt.Println("add bucket ", i)
	}

	fmt.Println("size of the hash ring ", unsafe.Sizeof(hashtab))
	
	loopNum, _ := strconv.ParseInt(os.Args[3], 10, 32)
	result := make(map[string]int)
	for i := 0; i < int(loopNum); i++ {
		target := "0" + "lubing.wan" + strconv.Itoa(i)
		if _, ok := result[hashtab.GetConsistentNode(target)]; ok {
			result[hashtab.GetConsistentNode(target)] += 1
		}else {
			result[hashtab.GetConsistentNode(target)] = 1
		}

	}
	fmt.Println(result)



	// var ring HashRing

	// for i := 1; i < 45; i++ {
	// 	if i % 3 == 0 {
	// 		ring = append(ring, uint32(i))
	// 	}
	// }
	// fmt.Println(ring)
	// sort.Sort(ring)
	// fmt.Println(ring)
	// var j uint32 = 11
	// i := sort.Search(len(ring), func(i int) bool { return ring[i] >= j })
	// fmt.Println(i, " ", ring[i]);

	// j = 108
	// i = sort.Search(len(ring), func(i int) bool { return ring[i] >= j })
	// fmt.Println(i);

	// j = 1
	// i = sort.Search(len(ring), func(i int) bool { return ring[i] >= j })
	// fmt.Println(i);

	// j = 4
	// i = sort.Search(len(ring), func(i int) bool { return ring[i] >= j })
	// fmt.Println(i);
}

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
 )
 
 // GetHash returns a murmur32 hash for the data slice.
 func hash_mur32(data []byte) uint32 {
	// Seed is set to 37, same as C# version of emitter
	var h1 uint32 = 37
 
	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
	   p = uintptr(unsafe.Pointer(&data[0]))
	}
 
	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
	   k1 := *(*uint32)(unsafe.Pointer(p))
 
	   k1 *= c1_32
	   k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
	   k1 *= c2_32
 
	   h1 ^= k1
	   h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
	   h1 = h1*5 + 0xe6546b64
	}
 
	tail := data[nblocks*4:]
 
	var k1 uint32
	switch len(tail) & 3 {
	case 3:
	   k1 ^= uint32(tail[2]) << 16
	   fallthrough
	case 2:
	   k1 ^= uint32(tail[1]) << 8
	   fallthrough
	case 1:
	   k1 ^= uint32(tail[0])
	   k1 *= c1_32
	   k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
	   k1 *= c2_32
	   h1 ^= k1
	}
 
	h1 ^= uint32(len(data))
 
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16
 
	return (h1 << 24) | (((h1 >> 8) << 16) & 0xFF0000) | (((h1 >> 16) << 8) & 0xFF00) | (h1 >> 24)
 }


 func hash_sha1(data []byte) uint32 {
	hash := sha1.New()
	hash.Write(data)
	hb := hash.Sum(nil)

	if len(hb) < 4 {
		return 0
	}
	v := (uint32(hb[3]) << 24) | (uint32(hb[2]) << 16) | (uint32(hb[1]) << 8) | (uint32(hb[0]))
	return v
}

func hash_ieee(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}