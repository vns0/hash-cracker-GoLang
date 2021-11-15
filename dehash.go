package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ripemd160"
	"math"

	"github.com/pedroalbanese/edgetk/groestl"
	"github.com/pedroalbanese/edgetk/jh"
	"github.com/pedroalbanese/gmsm/sm3"
	"github.com/pedroalbanese/gogost/gost28147"
	"github.com/pedroalbanese/gogost/gost34112012256"
	"github.com/pedroalbanese/gogost/gost34112012512"
	"github.com/pedroalbanese/gogost/gost341194"
	"github.com/pedroalbanese/tiger"
	"github.com/pedroalbanese/whirlpool"
	"github.com/pedroalbanese/skein"
)

type HASHCracker struct {
	hash     []uint8
	chars    []uint8
	minLen   uint8
	maxLen   uint8
	solution chan string
	stopchan chan struct{}
	endchan  chan struct{}
	hashFunc func([]byte) []byte
}

func NewHASHCracker(algorithm string, hashString string, chars []uint8, minLen uint8, maxLen uint8) *HASHCracker {
	// decode hash
	hash, err := hex.DecodeString(hashString)
	if err != nil {
		panic("Hash is not valid!")
	}

	// getting hash function
	var hashFunc func([]byte) []byte
	switch algorithm {
	case "sha1":
		hashFunc = func(msg []byte) []byte {
			res := sha1.Sum(msg)
			return res[:]
		}
	case "sha256":
		hashFunc = func(msg []byte) []byte {
			res := sha256.Sum256(msg)
			return res[:]
		}
	case "sha512":
		hashFunc = func(msg []byte) []byte {
			res := sha512.Sum512(msg)
			return res[:]
		}
	case "md5":
		hashFunc = func(msg []byte) []byte {
			res := md5.Sum(msg)
			return res[:]
		}
	case "blake2b256":
		hashFunc = func(msg []byte) []byte {
			res := blake2b.Sum256(msg)
			return res[:]
		}
	case "blake2b512":
		hashFunc = func(msg []byte) []byte {
			res := blake2b.Sum512(msg)
			return res[:]
		}
	case "blake2s256":
		hashFunc = func(msg []byte) []byte {
			res := blake2s.Sum256(msg)
			return res[:]
		}
	case "sha3-256":
		hashFunc = func(msg []byte) []byte {
			res := sha3.Sum256(msg)
			return res[:]
		}
	case "sha3-512":
		hashFunc = func(msg []byte) []byte {
			res := sha3.Sum512(msg)
			return res[:]
		}
	case "keccak256":
		hashFunc = func(msg []byte) []byte {
			res := sha3.NewLegacyKeccak256()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "keccak512":
		hashFunc = func(msg []byte) []byte {
			res := sha3.NewLegacyKeccak512()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "groestl":
		hashFunc = func(msg []byte) []byte {
			res := groestl.Sum256(msg)
			return res[:]
		}
	case "jh":
		hashFunc = func(msg []byte) []byte {
			res := jh.Sum256(msg)
			return res[:]
		}
	case "tiger":
		hashFunc = func(msg []byte) []byte {
			res := tiger.New()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "rmd160":
		hashFunc = func(msg []byte) []byte {
			res := ripemd160.New()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "sm3":
		hashFunc = func(msg []byte) []byte {
			res := sm3.New()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "gost94":
		hashFunc = func(msg []byte) []byte {
			res := gost341194.New(&gost28147.SboxIdGostR341194CryptoProParamSet)
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "streebog256":
		hashFunc = func(msg []byte) []byte {
			res := gost34112012256.New()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "streebog512":
		hashFunc = func(msg []byte) []byte {
			res := gost34112012512.New()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "whirlpool":
		hashFunc = func(msg []byte) []byte {
			res := whirlpool.New()
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "skein256":
		hashFunc = func(msg []byte) []byte {
			res := skein.New256(nil)
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	case "skein512":
		hashFunc = func(msg []byte) []byte {
			res := skein.New512(nil)
			res.Write(msg)
			hash := res.Sum(nil)
			return []byte(hash)
		}
	default:
		panic("Wrong hashing algorithm name was passed!")
	}

	return &HASHCracker{
		hash, chars, minLen, maxLen,
		make(chan string), make(chan struct{}), make(chan struct{}),
		hashFunc}
}

func (s *HASHCracker) checkIsSuit(bNum *ByteNumber) bool {
	// translate premutation to current charset
	msg := bNum.Translate(s.chars)
	hashSum := []uint8(s.hashFunc(msg))
	return bytes.Equal(hashSum, s.hash)
}

func (s *HASHCracker) bruteForce(bNum *ByteNumber, iterations uint64) {
	for iterations > 0 {
		select {
		default:
			if s.checkIsSuit(bNum) {
				s.solution <- string(bNum.Translate(s.chars))
				return
			}
			bNum.Inc()
			iterations--
		case <-s.stopchan:
			return
		}
	}
	s.endchan <- struct{}{}
}

func (s *HASHCracker) Crack(goroutines uint32) string {
	fmt.Printf("Start cracking hash %x\n", s.hash)
	for msgLen := s.minLen; msgLen < s.maxLen+1; msgLen++ {
		// number of all possible variants
		var variants uint64 = uint64(math.Pow(float64(len(s.chars)), float64(msgLen)))
		var jobs uint32 = goroutines
		if variants <= uint64(goroutines) {
			jobs = 1
		}
		// number of hashes to check in one goroutine
		var itersInGorutine uint64 = variants / uint64(jobs)

		fmt.Println("Check mesages with length:", msgLen, "| Possible variants:", variants)

		// start goroutines
		for i := uint32(0); i < jobs-1; i++ {
			variants -= itersInGorutine
			// make object that represents premutation on (i * premutation) step
			bNum := NewByteNumber(uint64(i)*itersInGorutine, uint8(len(s.chars)), msgLen)
			// brute force all premutations from (i * itersInGorutine) to (i * itersInGorutine + itersInGorutine)
			go s.bruteForce(bNum, itersInGorutine)
		}
		bNum := NewByteNumber(uint64(jobs-1)*itersInGorutine, uint8(len(s.chars)), msgLen)
		go s.bruteForce(bNum, variants-1)

		// wait for all goroutines to finish work or get solution
		var group uint32 = jobs
		for group > 0 {
			select {
			case <-s.endchan:
				group--
				continue
			case solution := <-s.solution:
				// if the solution came then tell other goroutines to stop
				close(s.stopchan)
				return solution
			}
		}
	}
	return ""
}
