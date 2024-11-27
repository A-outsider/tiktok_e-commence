package bloom_filter

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"github.com/redis/go-redis/v9"
	"hash"
	"hash/crc32"
	"hash/fnv"
	"hash/maphash"

	"github.com/spaolacci/murmur3"
	"golang.org/x/net/context"
)

// BloomFilter 使用 Redis 的布隆过滤器
type BloomFilter struct {
	client    *redis.Client         // Redis 客户端
	key       string                // Redis 键名
	bitSize   uint64                // 位数组大小
	hashFuncs []func([]byte) uint64 // 6 种哈希函数
}

// NewBloomFilter 创建布隆过滤器
func NewBloomFilter(client *redis.Client, key string, bitSize uint64) *BloomFilter {
	// 定义 6 种不同的哈希函数
	hashFuncs := []func([]byte) uint64{
		fnvHash,     // FNV-1a 哈希
		murmurHash,  // MurmurHash3
		maphashHash, // Go maphash
		crc32Hash,   // CRC32 哈希
		sha1Hash,    // SHA1 前 64 位
		md5Hash,     // MD5 前 64 位
	}

	return &BloomFilter{
		client:    client,
		key:       key,
		bitSize:   bitSize,
		hashFuncs: hashFuncs,
	}
}

// fnvHash 使用 FNV-1a
func fnvHash(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}

// murmurHash 使用 MurmurHash3
func murmurHash(data []byte) uint64 {
	return murmur3.Sum64(data)
}

// maphashHash 使用 maphash
func maphashHash(data []byte) uint64 {
	var h hash.Hash64 = new(maphash.Hash)
	h.Write(data)
	return h.Sum64()
}

// crc32Hash 使用 CRC32
func crc32Hash(data []byte) uint64 {
	return uint64(crc32.ChecksumIEEE(data))
}

// sha1Hash 使用 SHA1，截取前 64 位
func sha1Hash(data []byte) uint64 {
	h := sha1.New()
	h.Write(data)
	return binary.BigEndian.Uint64(h.Sum(nil)[:8])
}

// md5Hash 使用 MD5，截取前 64 位
func md5Hash(data []byte) uint64 {
	h := md5.New()
	h.Write(data)
	return binary.BigEndian.Uint64(h.Sum(nil)[:8])
}

func (bf *BloomFilter) Add(ctx context.Context, data []byte) error {
	for _, hashFunc := range bf.hashFuncs {
		hashValue := hashFunc(data)
		index := hashValue % bf.bitSize // 计算位数组中的位置
		_, err := bf.client.SetBit(ctx, bf.key, int64(index), 1).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

// Exist 检查元素是否存在于布隆过滤器中
func (bf *BloomFilter) Exist(ctx context.Context, data []byte) (bool, error) {
	for _, hashFunc := range bf.hashFuncs {
		hashValue := hashFunc(data)
		index := hashValue % bf.bitSize
		result, err := bf.client.GetBit(ctx, bf.key, int64(index)).Result()
		if err != nil {
			return false, err
		}
		if result == 0 {
			return false, nil // 如果某个位为 0，元素一定不存在
		}
	}
	return true, nil // 所有位都为 1，元素可能存在
}
