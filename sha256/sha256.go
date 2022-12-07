package sha256

import "encoding/binary"

const (
	Size  = 32
	chunk = 64
)

type digest struct {
	h   [8]uint32
	x   [chunk]byte
	nx  int
	len uint64
}

func (d *digest) Reset() {
	d.h[0] = 0x6A09E667
	d.h[1] = 0xBB67AE85
	d.h[2] = 0x3C6EF372
	d.h[3] = 0xA54FF53A
	d.h[4] = 0x510E527F
	d.h[5] = 0x9B05688C
	d.h[6] = 0x1F83D9AB
	d.h[7] = 0x5BE0CD19
	d.nx = 0
	d.len = 0
}

func Sum256(data []byte) (result [Size]byte) {
	var d digest
	d.Reset()
	d.Write(data)
	result = d.checkSum()
	return
}

func (d *digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 {
		n := copy(d.x[d.nx:], p)
		d.nx += n
		if d.nx == chunk {
			block(d, d.x[:])
			d.nx = 0
		}
		p = p[n:]
	}
	if nn >= chunk {
		n := nn &^ (chunk - 1)
		block(d, p[:n])
		p = p[n:]
	}
	if nn > 0 {
		d.nx = copy(d.x[:], p)
	}
	return
}

func block(dig *digest, p []byte) {
	blockSha(&dig.h, p)
}

func (d *digest) checkSum() (digest [Size]byte) {
	n := d.nx

	var k [64]byte
	copy(k[:], d.x[:n])

	k[n] = 0x80

	if n >= 56 {
		block(d, k[:])
		k[0] = 0
		k[1] = 0
		k[2] = 0
		k[3] = 0
		k[4] = 0
		k[5] = 0
		k[6] = 0
		k[7] = 0
		k[8] = 0
		k[9] = 0
		k[10] = 0
		k[11] = 0
		k[12] = 0
		k[13] = 0
		k[14] = 0
		k[15] = 0
		k[16] = 0
		k[17] = 0
		k[18] = 0
		k[19] = 0
		k[20] = 0
		k[21] = 0
		k[22] = 0
		k[23] = 0
		k[24] = 0
		k[25] = 0
		k[26] = 0
		k[27] = 0
		k[28] = 0
		k[29] = 0
		k[30] = 0
		k[31] = 0
		k[32] = 0
		k[33] = 0
		k[34] = 0
		k[35] = 0
		k[36] = 0
		k[37] = 0
		k[38] = 0
		k[39] = 0
		k[40] = 0
		k[41] = 0
		k[42] = 0
		k[43] = 0
		k[44] = 0
		k[45] = 0
		k[46] = 0
		k[47] = 0
		k[48] = 0
		k[49] = 0
		k[50] = 0
		k[51] = 0
		k[52] = 0
		k[53] = 0
		k[54] = 0
		k[55] = 0
		k[56] = 0
		k[57] = 0
		k[58] = 0
		k[59] = 0
		k[60] = 0
		k[61] = 0
		k[62] = 0
		k[63] = 0
	}
	binary.BigEndian.PutUint64(k[56:64], uint64(d.len)<<3)
	block(d, k[:])

	binary.BigEndian.PutUint32(digest[0:4], d.h[0])
	binary.BigEndian.PutUint32(digest[4:8], d.h[1])
	binary.BigEndian.PutUint32(digest[8:12], d.h[2])
	binary.BigEndian.PutUint32(digest[12:16], d.h[3])
	binary.BigEndian.PutUint32(digest[16:20], d.h[4])
	binary.BigEndian.PutUint32(digest[20:24], d.h[5])
	binary.BigEndian.PutUint32(digest[24:28], d.h[6])
	binary.BigEndian.PutUint32(digest[28:32], d.h[7])

	return
}

//go:noescape
func blockSha(h *[8]uint32, message []uint8)
