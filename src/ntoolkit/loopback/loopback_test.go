package loopback_test

import (
	"ntoolkit/assert"
	"ntoolkit/loopback"
	"testing"
	"time"
)

func TestNew(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		conn, err := loopback.New()
		T.Assert(err == nil)
		defer conn.Close()

		T.Assert(conn != nil)
		T.Assert(conn.A != nil)
		T.Assert(conn.B != nil)
	})
}

func TestReadWrite(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		conn, err := loopback.New()
		T.Assert(err == nil)
		defer conn.Close()

		buf1 := [4]byte{0, 1, 2, 3}
		buf2 := [4]byte{0, 0, 0, 0}

		wrote, werr := conn.A.Write(buf1[:])
		T.Assert(werr == nil)
		T.Assert(wrote == 4)

		read, rerr := conn.B.Read(buf2[:])
		T.Assert(rerr == nil)
		T.Assert(read == 4)
		T.Assert(buf2 == [4]byte{0, 1, 2, 3})
	})
}

func TestReadWriteRepeat(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		conn, err := loopback.New()
		T.Assert(err == nil)
		defer conn.Close()

		buf1 := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		buf2 := [4]byte{0, 0, 0, 0}

		conn.A.SetDeadline(time.Now().Add(time.Second))
		conn.B.SetDeadline(time.Now().Add(time.Second))

		for i := 0; i < 10; i++ {
			wrote, werr := conn.A.Write(buf1[:8])
			T.Assert(werr == nil)
			T.Assert(wrote == 8)

			wrote, werr = conn.A.Write(buf1[8:])
			T.Assert(werr == nil)
			T.Assert(wrote == 8)
		}

		for i := 0; i < 10; i++ {
			read, rerr := conn.B.Read(buf2[:])
			T.Assert(rerr == nil)
			T.Assert(read == 4)

			read, rerr = conn.B.Read(buf2[:])
			T.Assert(rerr == nil)
			T.Assert(read == 4)

			read, rerr = conn.B.Read(buf2[:])
			T.Assert(rerr == nil)
			T.Assert(read == 4)

			read, rerr = conn.B.Read(buf2[:])
			T.Assert(rerr == nil)
			T.Assert(read == 4)
		}
	})
}

func TestReadWriteTimeout(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		conn, err := loopback.New()
		T.Assert(err == nil)
		defer conn.Close()

		conn.A.Close()

		one := []byte{0}
		conn.B.SetReadDeadline(time.Now())
		_, err = conn.B.Read(one)
		T.Assert(err != nil)
	})
}
