package win32

type VT uint16

const (
	VT_EMPTY           VT = 0x0
	VT_NULL            VT = 0x1
	VT_I2              VT = 0x2
	VT_I4              VT = 0x3
	VT_R4              VT = 0x4
	VT_R8              VT = 0x5
	VT_CY              VT = 0x6
	VT_DATE            VT = 0x7
	VT_BSTR            VT = 0x8
	VT_DISPATCH        VT = 0x9
	VT_ERROR           VT = 0xa
	VT_BOOL            VT = 0xb
	VT_VARIANT         VT = 0xc
	VT_UNKNOWN         VT = 0xd
	VT_DECIMAL         VT = 0xe
	VT_I1              VT = 0x10
	VT_UI1             VT = 0x11
	VT_UI2             VT = 0x12
	VT_UI4             VT = 0x13
	VT_I8              VT = 0x14
	VT_UI8             VT = 0x15
	VT_INT             VT = 0x16
	VT_UINT            VT = 0x17
	VT_VOID            VT = 0x18
	VT_HRESULT         VT = 0x19
	VT_PTR             VT = 0x1a
	VT_SAFEARRAY       VT = 0x1b
	VT_CARRAY          VT = 0x1c
	VT_USERDEFINED     VT = 0x1d
	VT_LPSTR           VT = 0x1e
	VT_LPWSTR          VT = 0x1f
	VT_RECORD          VT = 0x24
	VT_INT_PTR         VT = 0x25
	VT_UINT_PTR        VT = 0x26
	VT_FILETIME        VT = 0x40
	VT_BLOB            VT = 0x41
	VT_STREAM          VT = 0x42
	VT_STORAGE         VT = 0x43
	VT_STREAMED_OBJECT VT = 0x44
	VT_STORED_OBJECT   VT = 0x45
	VT_BLOB_OBJECT     VT = 0x46
	VT_CF              VT = 0x47
	VT_CLSID           VT = 0x48
	VT_BSTR_BLOB       VT = 0xfff
	VT_VECTOR          VT = 0x1000
	VT_ARRAY           VT = 0x2000
	VT_BYREF           VT = 0x4000
	VT_RESERVED        VT = 0x8000
	VT_ILLEGAL         VT = 0xffff
	VT_ILLEGALMASKED   VT = 0xfff
	VT_TYPEMASK        VT = 0xfff
)

// type VARIANT struct {
// 	VT         VT      //  2
// 	wReserved1 uint16  //  4
// 	wReserved2 uint16  //  6
// 	wReserved3 uint16  //  8
// 	Val        int64   // 16
// 	_          [8]byte // 24
// }

type VARIANT [24]byte
