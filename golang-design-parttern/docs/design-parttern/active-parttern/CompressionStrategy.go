package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CompressionStrategy 压缩策略接口
type CompressionStrategy interface {
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
	GetExtension() string
	GetName() string
}

// GzipStrategy Gzip压缩策略
type GzipStrategy struct{}

func (s *GzipStrategy) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	
	_, err := writer.Write(data)
	if err != nil {
		return nil, fmt.Errorf("gzip压缩失败: %w", err)
	}
	
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("gzip关闭失败: %w", err)
	}
	
	return buf.Bytes(), nil
}

func (s *GzipStrategy) Decompress(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("gzip解压失败: %w", err)
	}
	defer reader.Close()
	
	return ioutil.ReadAll(reader)
}

func (s *GzipStrategy) GetExtension() string {
	return ".gz"
}

func (s *GzipStrategy) GetName() string {
	return "Gzip"
}

// ZlibStrategy Zlib压缩策略
type ZlibStrategy struct{}

func (s *ZlibStrategy) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	
	_, err := writer.Write(data)
	if err != nil {
		return nil, fmt.Errorf("zlib压缩失败: %w", err)
	}
	
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("zlib关闭失败: %w", err)
	}
	
	return buf.Bytes(), nil
}

func (s *ZlibStrategy) Decompress(data []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("zlib解压失败: %w", err)
	}
	defer reader.Close()
	
	return ioutil.ReadAll(reader)
}

func (s *ZlibStrategy) GetExtension() string {
	return ".zlib"
}

func (s *ZlibStrategy) GetName() string {
	return "Zlib"
}

// NoCompressionStrategy 无压缩策略（用于不需要压缩的文件类型）
type NoCompressionStrategy struct{}

func (s *NoCompressionStrategy) Compress(data []byte) ([]byte, error) {
	return data, nil // 直接返回原始数据
}

func (s *NoCompressionStrategy) Decompress(data []byte) ([]byte, error) {
	return data, nil // 直接返回原始数据
}

func (s *NoCompressionStrategy) GetExtension() string {
	return "" // 不添加扩展名
}

func (s *NoCompressionStrategy) GetName() string {
	return "无压缩"
}

// LZWStrategy LZW压缩策略（模拟实现，实际应使用真实的LZW库）
type LZWStrategy struct{}

func (s *LZWStrategy) Compress(data []byte) ([]byte, error) {
	// 这里是模拟实现，实际应使用真实的LZW库
	// 例如可以使用 github.com/golang/compress/lzw
	
	// 简单模拟，实际不进行压缩
	var buf bytes.Buffer
	buf.Write(data)
	return buf.Bytes(), nil
}

func (s *LZWStrategy) Decompress(data []byte) ([]byte, error) {
	// 同上，这是模拟实现
	return data, nil
}

func (s *LZWStrategy) GetExtension() string {
	return ".lzw"
}

func (s *LZWStrategy) GetName() string {
	return "LZW"
}

// CompressionContext 压缩上下文
type CompressionContext struct {
	strategy CompressionStrategy
}

// NewCompressionContext 创建压缩上下文
func NewCompressionContext(strategy CompressionStrategy) *CompressionContext {
	return &CompressionContext{
		strategy: strategy,
	}
}

// SetStrategy 设置压缩策略
func (c *CompressionContext) SetStrategy(strategy CompressionStrategy) {
	c.strategy = strategy
}

// CompressFile 压缩文件
func (c *CompressionContext) CompressFile(inputPath, outputPath string) error {
	// 读取输入文件
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}
	
	// 压缩数据
	compressed, err := c.strategy.Compress(data)
	if err != nil {
		return err
	}
	
	// 如果没有指定输出路径，则使用输入路径加上压缩扩展名
	if outputPath == "" {
		outputPath = inputPath + c.strategy.GetExtension()
	}
	
	// 写入压缩后的数据
	if err := ioutil.WriteFile(outputPath, compressed, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	
	fmt.Printf("文件已使用%s压缩: %s -> %s\n", c.strategy.GetName(), inputPath, outputPath)
	fmt.Printf("原始大小: %d 字节, 压缩后大小: %d 字节\n", len(data), len(compressed))
	
	return nil
}

// DecompressFile 解压文件
func (c *CompressionContext) DecompressFile(inputPath, outputPath string) error {
	// 读取压缩文件
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}
	
	// 解压数据
	decompressed, err := c.strategy.Decompress(data)
	if err != nil {
		return err
	}
	
	// 如果没有指定输出路径，则使用输入路径去除压缩扩展名
	if outputPath == "" {
		ext := c.strategy.GetExtension()
		if ext != "" && strings.HasSuffix(inputPath, ext) {
			outputPath = inputPath[:len(inputPath)-len(ext)]
		} else {
			outputPath = inputPath + ".decompressed"
		}
	}
	
	// 写入解压后的数据
	if err := ioutil.WriteFile(outputPath, decompressed, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	
	fmt.Printf("文件已使用%s解压: %s -> %s\n", c.strategy.GetName(), inputPath, outputPath)
	
	return nil
}

// StrategyFactory 策略工厂，根据文件类型选择合适的压缩策略
type StrategyFactory struct {
	strategies map[string]CompressionStrategy
}

// NewStrategyFactory 创建策略工厂
func NewStrategyFactory() *StrategyFactory {
	factory := &StrategyFactory{
		strategies: make(map[string]CompressionStrategy),
	}
	
	// 注册默认策略
	factory.RegisterStrategy(".txt", &GzipStrategy{})
	factory.RegisterStrategy(".log", &GzipStrategy{})
	factory.RegisterStrategy(".csv", &GzipStrategy{})
	factory.RegisterStrategy(".xml", &GzipStrategy{})
	factory.RegisterStrategy(".json", &GzipStrategy{})
	
	factory.RegisterStrategy(".jpg", &ZlibStrategy{})
	factory.RegisterStrategy(".jpeg", &ZlibStrategy{})
	factory.RegisterStrategy(".png", &ZlibStrategy{})
	
	factory.RegisterStrategy(".html", &LZWStrategy{})
	factory.RegisterStrategy(".htm", &LZWStrategy{})
	factory.RegisterStrategy(".css", &LZWStrategy{})
	factory.RegisterStrategy(".js", &LZWStrategy{})
	
	// 对于已经压缩的文件类型，不再压缩
	factory.RegisterStrategy(".zip", &NoCompressionStrategy{})
	factory.RegisterStrategy(".gz", &NoCompressionStrategy{})
	factory.RegisterStrategy(".rar", &NoCompressionStrategy{})
	factory.RegisterStrategy(".7z", &NoCompressionStrategy{})
	factory.RegisterStrategy(".mp3", &NoCompressionStrategy{})
	factory.RegisterStrategy(".mp4", &NoCompressionStrategy{})
	factory.RegisterStrategy(".avi", &NoCompressionStrategy{})
	
	return factory
}

// RegisterStrategy 注册文件扩展名对应的压缩策略
func (f *StrategyFactory) RegisterStrategy(extension string, strategy CompressionStrategy) {
	f.strategies[strings.ToLower(extension)] = strategy
}

// GetStrategy 根据文件名获取合适的压缩策略
func (f *StrategyFactory) GetStrategy(filename string) CompressionStrategy {
	ext := strings.ToLower(filepath.Ext(filename))
	
	if strategy, ok := f.strategies[ext]; ok {
		return strategy
	}
	
	// 默认使用Gzip
	return &GzipStrategy{}
}

// FileCompressor 文件压缩器
type FileCompressor struct {
	context *CompressionContext
	factory *StrategyFactory
}

// NewFileCompressor 创建文件压缩器
func NewFileCompressor() *FileCompressor {
	factory := NewStrategyFactory()
	return &FileCompressor{
		context: NewCompressionContext(&GzipStrategy{}), // 默认策略
		factory: factory,
	}
}

// CompressFile 智能压缩文件
func (fc *FileCompressor) CompressFile(inputPath, outputPath string) error {
	// 根据文件类型选择合适的压缩策略
	strategy := fc.factory.GetStrategy(inputPath)
	fc.context.SetStrategy(strategy)
	
	return fc.context.CompressFile(inputPath, outputPath)
}

// DecompressFile 解压文件
func (fc *FileCompressor) DecompressFile(inputPath, outputPath string) error {
	// 根据文件扩展名选择解压策略
	var strategy CompressionStrategy
	
	if strings.HasSuffix(inputPath, ".gz") {
		strategy = &GzipStrategy{}
	} else if strings.HasSuffix(inputPath, ".zlib") {
		strategy = &ZlibStrategy{}
	} else if strings.HasSuffix(inputPath, ".lzw") {
		strategy = &LZWStrategy{}
	} else {
		return fmt.Errorf("无法识别的压缩文件格式: %s", inputPath)
	}
	
	fc.context.SetStrategy(strategy)
	return fc.context.DecompressFile(inputPath, outputPath)
}

// CompressDirectory 压缩目录中的所有文件
func (fc *FileCompressor) CompressDirectory(dirPath string, recursive bool) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// 跳过目录和已压缩的文件
		if info.IsDir() {
			if path != dirPath && !recursive {
				return filepath.SkipDir
			}
			return nil
		}
		
		// 跳过已经是压缩文件的文件
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".gz" || ext == ".zlib" || ext == ".lzw" || ext == ".zip" || ext == ".rar" || ext == ".7z" {
			fmt.Printf("跳过已压缩文件: %s\n", path)
			return nil
		}
		
		// 压缩文件
		return fc.CompressFile(path, "")
	})
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法: program [compress|decompress|batch] [文件路径|目录路径] [输出路径(可选)]")
		return
	}
	
	command := os.Args[1]
	path := os.Args[2]
	outputPath := ""
	
	if len(os.Args) > 3 {
		outputPath = os.Args[3]
	}
	
	compressor := NewFileCompressor()
	
	switch command {
	case "compress":
		if err := compressor.CompressFile(path, outputPath); err != nil {
			fmt.Printf("压缩失败: %v\n", err)
		}
	case "decompress":
		if err := compressor.DecompressFile(path, outputPath); err != nil {
			fmt.Printf("解压失败: %v\n", err)
		}
	case "batch":
		fmt.Printf("批量压缩目录: %s\n", path)
		if err := compressor.CompressDirectory(path, true); err != nil {
			fmt.Printf("批量压缩失败: %v\n", err)
		}
	default:
		fmt.Println("未知命令，请使用 compress、decompress 或 batch")
	}
}