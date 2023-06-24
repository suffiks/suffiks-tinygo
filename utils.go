package suffiks

import (
	"unsafe"
)

func byteToPtr(b []byte) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.SliceData(b))
	return uint32(uintptr(ptr)), uint32(len(b))
}

func byteToPtrAndSize(b []byte) uint64 {
	ptr, size := byteToPtr(b)
	return uint64(ptr)<<32 | uint64(size)
}

type vtMarshaler interface {
	MarshalVT() ([]byte, error)
}

func marshal(m vtMarshaler) []byte {
	b, err := m.MarshalVT()
	if err != nil {
		panic(err)
	}
	return b
}

/*

  function getArray(ptrAndSize: u32): Uint8Array {
    var size: u32 = ptrAndSize & 0xffff;
    // Shift the combined value to the right by 16 bits to get the pointer address
    var ptr: u32 = ptrAndSize >>> 16;

    const b = new Uint8Array(size);
    for (let i = u32(0); i < size; i++) {
      b[i] = load<u8>(ptr + i);
    }

    heap.free(ptr);
    return b;
  }

	In go:
*/

func getByteSlice(ptrAndSize uint64) []byte {
	ptr := uint32(ptrAndSize >> 32)
	size := uint32(ptrAndSize)

	uptr := unsafe.Pointer(uintptr(ptr))
	b := unsafe.Slice((*byte)(uptr), size)

	return b
}

func ptrString(s string) *string {
	return &s
}

func stringToWasiPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}
