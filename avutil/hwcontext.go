package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/hwcontext.h>
//#include <libavutil/buffer.h>
import "C"
import (
	"unsafe"
)

type (
	AVHWDeviceType     C.enum_AVHWDeviceType
	HWDeviceContext	   C.struct_AVHWDeviceContext
	HWFramesContext	   C.struct_AVHWFramesContext
	BufferRef		   C.struct_AVBufferRef
	//AVQSVFramesContext C.struct_AVQSVFramesContext
)

const (
	AV_HWDEVICE_TYPE_NONE  = C.AV_HWDEVICE_TYPE_NONE
	AV_HWDEVICE_TYPE_VAAPI = C.AV_HWDEVICE_TYPE_VAAPI
	AV_HWDEVICE_TYPE_QSV   = C.AV_HWDEVICE_TYPE_QSV
	AV_HWDEVICE_TYPE_OPENCL	= C.AV_HWDEVICE_TYPE_OPENCL
)

func AVHwdeviceCtxAlloc(typ AVHWDeviceType) *BufferRef {
	return (*BufferRef)(C.av_hwdevice_ctx_alloc(C.enum_AVHWDeviceType(typ)))
}

func AVHwdeviceCtxCreate(dev_ctx **BufferRef, typ AVHWDeviceType, device string, opts *Dictionary, flags int) int{

	if device == "" {
		return int(C.av_hwdevice_ctx_create((**C.struct_AVBufferRef)(unsafe.Pointer(dev_ctx)), (C.enum_AVHWDeviceType)(typ),
		nil, (*C.struct_AVDictionary)(opts), (C.int)(flags)))
	}

	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))

	return int(C.av_hwdevice_ctx_create((**C.struct_AVBufferRef)(unsafe.Pointer(dev_ctx)), (C.enum_AVHWDeviceType)(typ),
			   dev, (*C.struct_AVDictionary)(opts), (C.int)(flags)))
}

func AVHwframeCtxAlloc(device_ctx *BufferRef) *BufferRef {
	return (*BufferRef)(C.av_hwframe_ctx_alloc((*C.struct_AVBufferRef)(device_ctx)))
}

func AVHwFrameCtxInit(ref *BufferRef) int {
	return (int)(C.av_hwframe_ctx_init((*C.struct_AVBufferRef)(ref)))
}

func AVBufferRef(buf *BufferRef) *BufferRef {
	return (*BufferRef)(C.av_buffer_ref((*C.struct_AVBufferRef)(buf)))
}

func AVBufferUnref(buf **BufferRef) {
	C.av_buffer_unref((**C.struct_AVBufferRef)(unsafe.Pointer(buf)))
}

/*
func AVHwdeviceCtxAlloc(typ AVHWDeviceType) *BufferRef {
	return (*BufferRef)(C.av_hwdevice_ctx_alloc(C.enum_AVHWDeviceType(typ)))
}

func (b *BufferRef)AVHwdeviceCtxCreate(typ AVHWDeviceType, device string, opts *Dictionary, flags int) int{

	if device == "" {
		return int(C.av_hwdevice_ctx_create((**C.struct_AVBufferRef)(unsafe.Pointer(&b)), (C.enum_AVHWDeviceType)(typ),
		nil, (*C.struct_AVDictionary)(opts), (C.int)(flags)))
	}

	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))

	return int(C.av_hwdevice_ctx_create((**C.struct_AVBufferRef)(unsafe.Pointer(&b)), (C.enum_AVHWDeviceType)(typ),
			   dev, (*C.struct_AVDictionary)(opts), (C.int)(flags)))
}

func (b *BufferRef)AVHwframeCtxAlloc() *BufferRef {
	return (*BufferRef)(C.av_hwframe_ctx_alloc((*C.struct_AVBufferRef)(b)))
}

//func (b *BufferRef)AVHwFrameCtxInit() int {
//	return (int)(C.av_hwframe_ctx_init((*C.struct_AVBufferRef)(b)))
//}

func AVHwFrameCtxInit(ref *BufferRef) int {
	return (int)(C.av_hwframe_ctx_init((*C.struct_AVBufferRef)(ref)))
}

func AVBufferRef(buf *BufferRef) *BufferRef {
	return (*BufferRef)(C.av_buffer_ref((*C.struct_AVBufferRef)(buf)))
}

func AVBufferUnref(buf **BufferRef) {
	C.av_buffer_unref((**C.struct_AVBufferRef)(unsafe.Pointer(buf)))
}*/

func (b *BufferRef)HWFramesContext() *HWFramesContext {
	return (*HWFramesContext)((*C.struct_AVHWFramesContext)(unsafe.Pointer(b.data)))
}

func (c *HWFramesContext) Width() int {
	return (int)(c.width)
}

func (c *HWFramesContext) Height() int {
	return (int)(c.height)
}

func (c *HWFramesContext)SetHWFramesContextVideoSize(w, h int) { //initial_pool_size ) {
	c.width     = (C.int)(w) 
	c.height	= (C.int)(h)
	//c.initial_pool_size = 32
}

func (c *HWFramesContext)SetHWFramesContextPrarms(format, sw_format PixelFormat, w, h int) { //initial_pool_size ) {
	c.format    = (C.enum_AVPixelFormat)(format)
	c.sw_format = (C.enum_AVPixelFormat)(sw_format)
	c.width     = (C.int)(w) 
	c.height	= (C.int)(h)
	//c.initial_pool_size = 32
}

func (c *HWFramesContext)SetHWFramesContextPrarms2(format, sw_format PixelFormat, w, h, s int) { //initial_pool_size ) {
	c.format    = (C.enum_AVPixelFormat)(format)
	c.sw_format = (C.enum_AVPixelFormat)(sw_format)
	c.width     = (C.int)(w) 
	c.height	= (C.int)(h)
	c.initial_pool_size = (C.int)(s)
}

/*func (c *HWFramesContext)SetQsvHWFramesContextPrarms(format, sw_format PixelFormat, w, h, s int) { //initial_pool_size ) {
	c.format    = (C.enum_AVPixelFormat)(format)
	c.sw_format = (C.enum_AVPixelFormat)(sw_format)
	c.width     = (C.int)(w) 
	c.height	= (C.int)(h)
	c.initial_pool_size = (C.int)(s)
	((*AVQSVFramesContext)((*C.struct_AVQSVFramesContext)(unsafe.Pointer(c.hwctx)))).frame_type = (C.int)(16)
}*/