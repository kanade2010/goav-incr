// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
/*
int avcodec_open2_native(AVCodecContext* ctxt, AVCodec* c, AVDictionary** d) {
	if (d != NULL) {
		AVDictionary *tmp = NULL;
		av_dict_copy(&tmp, *d, 0);

		return avcodec_open2(ctxt, c, &tmp);
	} else {
		return avcodec_open2(ctxt, c, NULL);
	}
}
*/
import "C"
import (
	"unsafe"
	"github.com/ailumiyana/goav-incr/goav/avutil"
)

func (ctxt *Context) AvCodecGetPktTimebase() Rational {
	return Rational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctxt)))
}

// AvCodecGetPktTimebase2 returns the timebase rational number as numerator and denominator
func (ctxt *Context) AvCodecGetPktTimebase2() Rational {
	return ctxt.AvCodecGetPktTimebase()
}

func (ctxt *Context) AvCodecSetPktTimebase(r Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctxt), (C.struct_AVRational)(r))
}

func (ctxt *Context) AvCodecGetCodecDescriptor() *Descriptor {
	return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *Context) AvCodecSetCodecDescriptor(d *Descriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecDescriptor)(d))
}

func (ctxt *Context) AvCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *Context) AvCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

func (ctxt *Context) AvCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *Context) AvCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

func (ctxt *Context) AvCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *Context) AvCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt), (*C.uint16_t)(t))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctxt *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&ctxt)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctxt *Context) AvcodecGetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c)))
}

//Copy the settings of the source Context into the destination Context.
func (ctxt *Context) AvcodecCopyContext(ctxt2 *Context) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the Context to use the given Codec
func (ctxt *Context) AvcodecOpen2(c *Codec, d **Dictionary) int {
	//return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
	return int(C.avcodec_open2_native((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (ctxt *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

//The default callback for Context.get_buffer2().
func (s *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(s), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctxt *Context) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctxt *Context) AvcodecDecodeVideo2(p *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
func (ctxt *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
func (ctxt *Context) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//Encode a frame of video
func (ctxt *Context) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (ctxt *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctxt *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (ctxt *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

//Return audio frame duration.
func (ctxt *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(f)))
}

func (ctxt *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt)))
}

//Parse a packet.
//func (ctxt *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
//	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
//}

//Parse a packet.
func (ctxt *Context) AvParserParse2(ctxtp *ParserContext, pkt *Packet, b []byte, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(&pkt.data)), (*C.int)(unsafe.Pointer(&pkt.size)), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (ctxt *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *Parser) AvParserNext() *Parser {
	return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *Parser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

func (ctxt *Context) SetTimebase(num1 int, den1 int) {
	ctxt.time_base.num = C.int(num1)
	ctxt.time_base.den = C.int(den1)
}

func (ctxt *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	ctxt.width = C.int(width)
	ctxt.height = C.int(height)
	// ctxt.bit_rate = 1000000
	ctxt.gop_size = C.int(gopSize)
	// ctxt.max_b_frames = 2
	if hasBframes {
		ctxt.has_b_frames = 1
	} else {
		ctxt.has_b_frames = 0
	}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	ctxt.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

func (ctxt *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctxt.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

func (ctxt *Context) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(packet)))
}

func (ctxt *Context) AvcodecReceiveFrame(frame *Frame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(frame)))
}

func (ctxt *Context) AvcodecSendFrame(frame *Frame) int {
	return (int)(C.avcodec_send_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(frame)))
}

func (ctxt *Context) AvcodecReceivePacket(packet *Packet) int {
	return (int)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(packet)))
}

func (ctxt *Context) AvcodecParametersToContext(pctxt *AvCodecParameters) int {
	return (int)(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecParameters)(pctxt)))
}

func (ctxt *Context) SetVideoEncodeParams(bitrate int64, width, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	ctxt.width = C.int(width)
	ctxt.height = C.int(height)
	ctxt.bit_rate = C.int64_t(bitrate)
	ctxt.gop_size = C.int(gopSize)
	// ctxt.max_b_frames = 2
	if hasBframes {
		ctxt.has_b_frames = 1
	} else {
		ctxt.has_b_frames = 0
		ctxt.max_b_frames = 0
	}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	ctxt.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

func (ctxt *Context) TestVideoEncodeParams() {
	ctxt.SetTimebase(1, 30)
	ctxt.width = C.int(1920)
	ctxt.height = C.int(1080)
	//ctxt.bit_rate = C.int64_t(bitrate)
	//ctxt.gop_size = C.int(gopSize)
	// ctxt.max_b_frames = 2
	//if hasBframes {
	//	ctxt.has_b_frames = 1
	//} else {
	//	ctxt.has_b_frames = 0
	//	ctxt.max_b_frames = 0
	//}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	ctxt.pix_fmt = int32(AV_PIX_FMT_QSV)
	// C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

func (ctxt *Context) SetAudioEncodeParams(br int64, sr int, cl string, sf AvSampleFormat) {
	ctxt.bit_rate   	= C.int64_t(br)
	ctxt.sample_fmt 	= C.enum_AVSampleFormat(sf)
	ctxt.sample_rate 	= C.int(sr)

	Cstr := C.CString(cl)
	defer C.free(unsafe.Pointer(Cstr))
	ctxt.channel_layout = C.av_get_channel_layout(Cstr)

	ctxt.channels       = C.av_get_channel_layout_nb_channels(ctxt.channel_layout)
}

func (ctxt *Context) SetAudioDecodeParams(br int64, sr int, cl string, sf AvSampleFormat) {
	ctxt.bit_rate   	= C.int64_t(br)
	ctxt.sample_fmt 	= C.enum_AVSampleFormat(sf)
	ctxt.sample_rate 	= C.int(sr)

	Cstr := C.CString(cl)
	defer C.free(unsafe.Pointer(Cstr))
	ctxt.channel_layout = C.av_get_channel_layout(Cstr)

	ctxt.channels       = C.av_get_channel_layout_nb_channels(ctxt.channel_layout)
}

func (ctxt *Context) SetAudioDecodeParamsTest(channels int) {
	ctxt.channels       = C.int(channels)
}

func (ctxt *Context) SetHWFramesCtx(hw_frames_ctx *avutil.BufferRef) {
	ctxt.hw_frames_ctx = (*C.struct_AVBufferRef)(unsafe.Pointer(hw_frames_ctx))
}

func (ctxt *Context) SetHWDeviceCtx(hw_device_ctx *avutil.BufferRef) {
	ctxt.hw_device_ctx = (*C.struct_AVBufferRef)(unsafe.Pointer(hw_device_ctx))
}

func (ctxt *Context) HWFramesCtx() *avutil.BufferRef {
	return (*avutil.BufferRef)(unsafe.Pointer(ctxt.hw_frames_ctx))
}

func (ctxt *Context) HWDeviceCtx() *avutil.BufferRef {
	return (*avutil.BufferRef)(unsafe.Pointer(ctxt.hw_device_ctx))
}

/*func (ctxt *Context) SetDefaultGetFormat() {
	ctxt.get_format = (C.get_format_callback)(unsafe.Pointer(C.avcodec_default_get_format))
}

//used for qsv

func (ctxt *Context) SetDefaultQsvGetFormat() {
	ctxt.get_format = (C.get_format_callback)(unsafe.Pointer(C.qsv_get_format))
}

#include "libavutil/hwcontext.h"
#include "libavutil/hwcontext_qsv.h"

enum AVPixelFormat qsv_get_format(AVCodecContext *avctx, const enum AVPixelFormat *pix_fmts)
{
    while (*pix_fmts != AV_PIX_FMT_NONE) {
        if (*pix_fmts == AV_PIX_FMT_QSV) {
            AVHWFramesContext  *frames_ctx;
            AVQSVFramesContext *frames_hwctx;
            int ret;

            avctx->hw_frames_ctx = av_hwframe_ctx_alloc(avctx->hw_device_ctx);
            if (!avctx->hw_frames_ctx)
                return AV_PIX_FMT_NONE;
            frames_ctx   = (AVHWFramesContext*)avctx->hw_frames_ctx->data;
            frames_hwctx = frames_ctx->hwctx;

            frames_ctx->format            = AV_PIX_FMT_QSV;
            frames_ctx->sw_format         = avctx->sw_pix_fmt;
            frames_ctx->width             = FFALIGN(avctx->coded_width,  32);
            frames_ctx->height            = FFALIGN(avctx->coded_height, 32);
            frames_ctx->initial_pool_size = 64;

            frames_hwctx->frame_type = 16;//MFX_MEMTYPE_VIDEO_MEMORY_DECODER_TARGET;

            ret = av_hwframe_ctx_init(avctx->hw_frames_ctx);
            if (ret < 0)
                return AV_PIX_FMT_NONE;

            return AV_PIX_FMT_QSV;
        }

        pix_fmts++;
    }

    fprintf(stderr, "The QSV pixel format not offered in get_format()\n");

    return AV_PIX_FMT_NONE;
}*/
