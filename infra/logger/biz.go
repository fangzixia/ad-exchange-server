package logger

import (
	"bytes"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// -------------------------- 核心：完整实现 zapcore.Encoder 所有接口 --------------------------
// BizLogEncoder 包装原生 JSONEncoder，并转发所有 ObjectEncoder 方法
type BizLogEncoder struct {
	innerEncoder zapcore.Encoder // 底层复用 JSONEncoder（已实现所有接口）
}

// 编译期断言：强制验证接口实现（zap v1.27.1 无任何报错）
var _ zapcore.Encoder = (*BizLogEncoder)(nil)

// NewBizLogEncoder 创建自定义编码器
func NewBizLogEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	inner := zapcore.NewJSONEncoder(cfg)
	return &BizLogEncoder{innerEncoder: inner}
}

// -------------------------- 1. 实现 zapcore.Encoder 核心方法 --------------------------
// EncodeEntry zap v1.27.1 接口核心方法（严格匹配签名）
func (e *BizLogEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 生成 17 位时间戳：YYYYMMDDHHMMSSSSS
	now := entry.Time
	baseTime := now.Format("20060102150405")
	millisecond := now.Nanosecond() / 1e6
	timestamp := fmt.Sprintf("%s%03d", baseTime, millisecond)

	// 调用底层编码器序列化字段（正确调用 EncodeEntry）
	jsonBuf, err := e.innerEncoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	// 清理换行符 + 拼接最终格式
	jsonData := bytes.TrimSuffix(jsonBuf.Bytes(), []byte("\n"))
	jsonBuf.Free()

	resultBuf := buffer.NewPool().Get()
	_, err = resultBuf.WriteString(fmt.Sprintf("%s|%s\n", timestamp, jsonData))
	if err != nil {
		resultBuf.Free()
		return nil, fmt.Errorf("拼接格式失败: %w", err)
	}

	return resultBuf, nil
}

// Clone 实现接口：转发到底层编码器的 Clone 方法
func (e *BizLogEncoder) Clone() zapcore.Encoder {
	return &BizLogEncoder{
		innerEncoder: e.innerEncoder.Clone(),
	}
}

// -------------------------- 2. 转发所有 zapcore.ObjectEncoder 方法（核心修正） --------------------------
// 以下所有方法均转发给 innerEncoder，确保 ObjectEncoder 接口完全实现
func (e *BizLogEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	return e.innerEncoder.AddArray(key, arr)
}

func (e *BizLogEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	return e.innerEncoder.AddObject(key, obj)
}

func (e *BizLogEncoder) AddBinary(key string, val []byte) {
	e.innerEncoder.AddBinary(key, val)
}

func (e *BizLogEncoder) AddBool(key string, val bool) {
	e.innerEncoder.AddBool(key, val)
}

func (e *BizLogEncoder) AddByteString(key string, val []byte) {
	e.innerEncoder.AddByteString(key, val)
}

func (e *BizLogEncoder) AddComplex128(key string, val complex128) {
	e.innerEncoder.AddComplex128(key, val)
}

func (e *BizLogEncoder) AddComplex64(key string, val complex64) {
	e.innerEncoder.AddComplex64(key, val)
}

func (e *BizLogEncoder) AddDuration(key string, val time.Duration) {
	e.innerEncoder.AddDuration(key, val)
}

func (e *BizLogEncoder) AddFloat64(key string, val float64) {
	e.innerEncoder.AddFloat64(key, val)
}

func (e *BizLogEncoder) AddFloat32(key string, val float32) {
	e.innerEncoder.AddFloat32(key, val)
}

func (e *BizLogEncoder) AddInt(key string, val int) {
	e.innerEncoder.AddInt(key, val)
}

func (e *BizLogEncoder) AddInt64(key string, val int64) {
	e.innerEncoder.AddInt64(key, val)
}

func (e *BizLogEncoder) AddInt32(key string, val int32) {
	e.innerEncoder.AddInt32(key, val)
}

func (e *BizLogEncoder) AddInt16(key string, val int16) {
	e.innerEncoder.AddInt16(key, val)
}

func (e *BizLogEncoder) AddInt8(key string, val int8) {
	e.innerEncoder.AddInt8(key, val)
}

func (e *BizLogEncoder) AddString(key string, val string) {
	e.innerEncoder.AddString(key, val)
}

func (e *BizLogEncoder) AddTime(key string, val time.Time) {
	e.innerEncoder.AddTime(key, val)
}

func (e *BizLogEncoder) AddUint(key string, val uint) {
	e.innerEncoder.AddUint(key, val)
}

func (e *BizLogEncoder) AddUint64(key string, val uint64) {
	e.innerEncoder.AddUint64(key, val)
}

func (e *BizLogEncoder) AddUint32(key string, val uint32) {
	e.innerEncoder.AddUint32(key, val)
}

func (e *BizLogEncoder) AddUint16(key string, val uint16) {
	e.innerEncoder.AddUint16(key, val)
}

func (e *BizLogEncoder) AddUint8(key string, val uint8) {
	e.innerEncoder.AddUint8(key, val)
}

func (e *BizLogEncoder) AddUintptr(key string, val uintptr) {
	e.innerEncoder.AddUintptr(key, val)
}

func (e *BizLogEncoder) AddReflected(key string, val interface{}) error {
	return e.innerEncoder.AddReflected(key, val)
}

func (e *BizLogEncoder) OpenNamespace(key string) {
	e.innerEncoder.OpenNamespace(key)
}

// -------------------------- 3. 日志器初始化 + 测试代码 --------------------------
func NewBizLogger(logPath string, maxSizeMB int, maxAgeDay int, compress bool) (*zap.Logger, error) {
	lumberjackWriter := &lumberjack.Logger{
		Filename:  logPath,
		MaxSize:   maxSizeMB,
		MaxAge:    maxAgeDay,
		Compress:  compress,
		LocalTime: true,
	}

	writeSyncer := zapcore.AddSync(lumberjackWriter)
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:      "",
		LevelKey:     "",
		MessageKey:   "",
		CallerKey:    "",
		FunctionKey:  "",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		LineEnding:   "",
	}

	core := zapcore.NewCore(
		NewBizLogEncoder(encoderCfg),
		writeSyncer,
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.ErrorOutput(writeSyncer))
	return logger, nil
}
