package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

// LZWStrategy LZW压缩策略（模拟实现）
type LZWStrategy struct{}

func (s *LZWStrategy) Compress(data []byte) ([]byte, error) {
	// 简单模拟，实际不进行压缩
	var buf bytes.Buffer
	buf.Write(data)
	return buf.Bytes(), nil
}

func (s *LZWStrategy) Decompress(data []byte) ([]byte, error) {
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

// GUI应用程序
func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("文件压缩工具")
	myWindow.Resize(fyne.NewSize(600, 400))

	// 创建文件压缩器
	compressor := NewFileCompressor()
	
	// 状态标签
	statusLabel := widget.NewLabel("请选择操作和文件")
	
	// 用于同步UI更新的互斥锁
	var uiMutex sync.Mutex
	
	// 安全更新UI的函数
	updateStatus := func(text string) {
		uiMutex.Lock()
		statusLabel.SetText(text)
		statusLabel.Refresh()
		uiMutex.Unlock()
	}
	
	// 文件路径显示
	selectedFileLabel := widget.NewLabel("未选择文件")
	selectedFileLabel.Wrapping = fyne.TextWrapBreak
	
	// 输出路径输入框
	outputPathEntry := widget.NewEntry()
	outputPathEntry.SetPlaceHolder("输出文件路径 (留空使用默认)")
	
	// 选择文件按钮
	var selectedFilePath string
	selectFileBtn := widget.NewButton("选择文件", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}
			
			selectedFilePath = reader.URI().Path()
			selectedFileLabel.SetText(selectedFilePath)
			
			// 自动设置输出路径建议
			fileName := filepath.Base(selectedFilePath)
			dir := filepath.Dir(selectedFilePath)
			
			if strings.HasSuffix(fileName, ".gz") || strings.HasSuffix(fileName, ".zlib") || strings.HasSuffix(fileName, ".lzw") {
				// 解压缩情况，去掉扩展名
				ext := filepath.Ext(fileName)
				outputPathEntry.SetText(filepath.Join(dir, fileName[:len(fileName)-len(ext)]))
			} else {
				// 压缩情况，添加扩展名
				strategy := compressor.factory.GetStrategy(selectedFilePath)
				outputPathEntry.SetText(filepath.Join(dir, fileName + strategy.GetExtension()))
			}
		}, myWindow)
		fd.Show()
	})
	
	// 压缩按钮
	compressBtn := widget.NewButtonWithIcon("压缩文件", theme.DocumentIcon(), func() {
		if selectedFilePath == "" {
			dialog.ShowInformation("提示", "请先选择要压缩的文件", myWindow)
			return
		}
		
		outputPath := outputPathEntry.Text
		
		// 执行压缩
		updateStatus("正在压缩...")
		
		// 使用goroutine执行耗时操作
		go func() {
			err := compressor.CompressFile(selectedFilePath, outputPath)
			finalOutputPath := outputPath
			
			// 确定输出文件路径
			if finalOutputPath == "" {
				strategy := compressor.factory.GetStrategy(selectedFilePath)
				finalOutputPath = selectedFilePath + strategy.GetExtension()
			}
			
			// 使用通知而不是直接更新UI
			if err != nil {
				updateStatus("压缩失败: " + err.Error())
			} else {
				updateStatus("压缩成功: " + finalOutputPath)
			}
		}()
	})
	
	// 解压按钮
	decompressBtn := widget.NewButtonWithIcon("解压文件", theme.FolderOpenIcon(), func() {
		if selectedFilePath == "" {
			dialog.ShowInformation("提示", "请先选择要解压的文件", myWindow)
			return
		}
		
		// 检查是否是支持的压缩文件
		if !strings.HasSuffix(selectedFilePath, ".gz") && 
		   !strings.HasSuffix(selectedFilePath, ".zlib") && 
		   !strings.HasSuffix(selectedFilePath, ".lzw") {
			dialog.ShowInformation("提示", "选择的文件不是支持的压缩格式 (.gz, .zlib, .lzw)", myWindow)
			return
		}
		
		outputPath := outputPathEntry.Text
		
		// 执行解压
		updateStatus("正在解压...")
		
		// 使用goroutine执行耗时操作
		go func() {
			err := compressor.DecompressFile(selectedFilePath, outputPath)
			finalOutputPath := outputPath
			
			// 确定输出文件路径
			if finalOutputPath == "" {
				ext := filepath.Ext(selectedFilePath)
				finalOutputPath = selectedFilePath[:len(selectedFilePath)-len(ext)]
			}
			
			// 使用通知而不是直接更新UI
			if err != nil {
				updateStatus("解压失败: " + err.Error())
			} else {
				updateStatus("解压成功: " + finalOutputPath)
			}
		}()
	})
	
	// 清除按钮
	clearBtn := widget.NewButtonWithIcon("清除", theme.DeleteIcon(), func() {
		selectedFilePath = ""
		selectedFileLabel.SetText("未选择文件")
		outputPathEntry.SetText("")
		statusLabel.SetText("请选择操作和文件")
	})
	
	// 布局
	fileSelectionBox := container.NewVBox(
		widget.NewLabel("选择文件:"),
		selectFileBtn,
		container.NewHScroll(selectedFileLabel),
	)
	
	outputBox := container.NewVBox(
		widget.NewLabel("输出路径 (可选):"),
		outputPathEntry,
	)
	
	buttonsBox := container.NewHBox(
		layout.NewSpacer(),
		compressBtn,
		decompressBtn,
		clearBtn,
		layout.NewSpacer(),
	)
	
	content := container.NewVBox(
		fileSelectionBox,
		widget.NewSeparator(),
		outputBox,
		widget.NewSeparator(),
		buttonsBox,
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewLabel("状态:"),
			statusLabel,
		),
	)
	
	myWindow.SetContent(container.NewPadded(content))
	myWindow.ShowAndRun()
}