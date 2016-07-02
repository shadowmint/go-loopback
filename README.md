# Loopback

A simple loopback helper.

## Usage

		conn, err := loopback.New()

		buf1 := [4]byte{0, 1, 2, 3}
		buf2 := [4]byte{0, 0, 0, 0}

		wrote, werr := conn.A.Write(buf1[:])
		read, rerr := conn.B.Read(buf2[:])
