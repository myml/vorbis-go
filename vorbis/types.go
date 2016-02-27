// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 28 Feb 2016 02:05:39 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package vorbis

/*
#cgo pkg-config: ogg vorbis
#include "ogg/ogg.h"
#include "vorbis/codec.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// OggIovec as declared in https://xiph.org/ogg/doc/libogg/ogg_iovec_t.html
type OggIovec struct {
	IovBase        unsafe.Pointer
	IovLen         uint
	ref92a8e63c    *C.ogg_iovec_t
	allocs92a8e63c interface{}
}

// OggPage as declared in https://xiph.org/ogg/doc/libogg/ogg_page.html
type OggPage struct {
	Header         []byte
	HeaderLen      int
	Body           []byte
	BodyLen        int
	refb80411d1    *C.ogg_page
	allocsb80411d1 interface{}
}

// OggStreamState as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_state.html
type OggStreamState struct {
	BodyData       []byte
	BodyStorage    int
	BodyFill       int
	BodyReturned   int
	LacingVals     []int32
	GranuleVals    []OggInt64
	LacingStorage  int
	LacingFill     int
	LacingPacket   int
	LacingReturned int
	Header         [282]byte
	HeaderFill     int32
	EOS            int32
	BOS            int32
	Serialno       int
	Pageno         int
	Packetno       OggInt64
	Granulepos     OggInt64
	refd2ebdd43    *C.ogg_stream_state
	allocsd2ebdd43 interface{}
}

// OggPacket as declared in https://xiph.org/ogg/doc/libogg/ogg_packet.html
type OggPacket struct {
	Packet         []byte
	Bytes          int
	BOS            int
	EOS            int
	Granulepos     OggInt64
	Packetno       OggInt64
	ref7954eecd    *C.ogg_packet
	allocs7954eecd interface{}
}

// OggSyncState as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_state.html
type OggSyncState struct {
	Data           []byte
	Storage        int32
	Fill           int32
	Returned       int32
	Unsynced       int32
	Headerbytes    int32
	Bodybytes      int32
	ref69d6025b    *C.ogg_sync_state
	allocs69d6025b interface{}
}

// OggInt16 type as declared in ogg/config_types.h:19
type OggInt16 int16

// OggUint16 type as declared in ogg/config_types.h:20
type OggUint16 uint16

// OggInt32 type as declared in ogg/config_types.h:21
type OggInt32 int32

// OggUint32 type as declared in ogg/config_types.h:22
type OggUint32 uint32

// OggInt64 type as declared in ogg/config_types.h:23
type OggInt64 int64

// Info as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info.html
type Info struct {
	Version        int32
	Channels       int32
	Rate           int
	BitrateUpper   int
	BitrateNominal int
	BitrateLower   int
	BitrateWindow  int
	CodecSetup     unsafe.Pointer
	refd271fa3c    *C.vorbis_info
	allocsd271fa3c interface{}
}

// DspState as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_dsp_state.html
type DspState struct {
	Analysisp      int32
	Vi             []Info
	Pcm            [][]float32
	Pcmret         [][]float32
	PcmStorage     int32
	PcmCurrent     int32
	PcmReturned    int32
	Preextrapolate int32
	Eofflag        int32
	Lw             int
	W              int
	Nw             int
	Centerw        int
	Granulepos     OggInt64
	Sequence       OggInt64
	GlueBits       OggInt64
	TimeBits       OggInt64
	FloorBits      OggInt64
	ResBits        OggInt64
	BackendState   unsafe.Pointer
	refde21b9b1    *C.vorbis_dsp_state
	allocsde21b9b1 interface{}
}

// Block as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_block.html
type Block struct {
	Pcm            [][]float32
	Lw             int
	W              int
	Nw             int
	Pcmend         int32
	Mode           int32
	Eofflag        int32
	Granulepos     OggInt64
	Sequence       OggInt64
	Vd             []DspState
	Localstore     unsafe.Pointer
	Localtop       int
	Localalloc     int
	Totaluse       int
	GlueBits       int
	TimeBits       int
	FloorBits      int
	ResBits        int
	Internal       unsafe.Pointer
	ref5962d739    *C.vorbis_block
	allocs5962d739 interface{}
}

// Comment as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment.html
type Comment struct {
	UserComments   [][]byte
	CommentLengths []int32
	Comments       int32
	Vendor         []byte
	ref4b9e021f    *C.vorbis_comment
	allocs4b9e021f interface{}
}
