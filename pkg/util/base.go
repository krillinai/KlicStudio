package util

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"io"
	"krillin-ai/internal/types"
	"math"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

var strWithUpperLowerNum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func GenerateRandStringWithUpperLowerNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = strWithUpperLowerNum[rand.Intn(len(strWithUpperLowerNum))]
	}
	return string(b)
}

func GetYouTubeID(youtubeURL string) (string, error) {
	parsedURL, err := url.Parse(youtubeURL)
	if err != nil {
		return "", err
	}

	if strings.Contains(parsedURL.Path, "watch") {
		queryParams := parsedURL.Query()
		if id, exists := queryParams["v"]; exists {
			return id[0], nil
		}
	} else {
		pathSegments := strings.Split(parsedURL.Path, "/")
		return pathSegments[len(pathSegments)-1], nil
	}

	return "", fmt.Errorf("no video ID found")
}

func GetBilibiliVideoId(url string) string {
	re := regexp.MustCompile(`https://(?:www\.)?bilibili\.com/(?:video/|video/av\d+/)(BV[a-zA-Z0-9]+)`)
	matches := re.FindStringSubmatch(url)
	if len(matches) > 1 {
		// 返回匹配到的BV号
		return matches[1]
	}
	return ""
}

// 将浮点数秒数转换为HH:MM:SS,SSS格式的字符串
func FormatTime(seconds float32) string {
	totalSeconds := int(math.Floor(float64(seconds)))             // 获取总秒数
	milliseconds := int((seconds - float32(totalSeconds)) * 1000) // 获取毫秒部分

	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	secs := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, secs, milliseconds)
}

// 判断字符串是否是纯数字（字幕编号）
func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Unzip(zipFile, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("打开zip文件失败: %v", err)
	}
	defer zipReader.Close()

	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return fmt.Errorf("创建目标目录失败: %v", err)
	}

	for _, file := range zipReader.File {
		filePath := filepath.Join(destDir, file.Name)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, file.Mode())
			if err != nil {
				return fmt.Errorf("创建目录失败: %v", err)
			}
			continue
		}

		destFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("创建文件失败: %v", err)
		}
		defer destFile.Close()

		zipFileReader, err := file.Open()
		if err != nil {
			return fmt.Errorf("打开zip文件内容失败: %v", err)
		}
		defer zipFileReader.Close()

		_, err = io.Copy(destFile, zipFileReader)
		if err != nil {
			return fmt.Errorf("复制文件内容失败: %v", err)
		}
	}

	return nil
}

func GenerateID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// ChangeFileExtension 修改文件后缀
func ChangeFileExtension(path string, newExt string) string {
	ext := filepath.Ext(path)
	return path[:len(path)-len(ext)] + newExt
}

func CleanPunction(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r)
	})
}

func IsAlphabetic(r rune) bool {
	if unicode.IsLetter(r) { // 中文在IsLetter中会返回true
		switch {
		// 英语及其他拉丁字母的范围
		case r >= 'A' && r <= 'Z', r >= 'a' && r <= 'z':
			return true
		// 扩展拉丁字母（法语、西班牙语等使用的附加字符）
		case r >= '\u00C0' && r <= '\u024F':
			return true
		// 希腊字母
		case r >= '\u0370' && r <= '\u03FF':
			return true
		// 西里尔字母（俄语等）
		case r >= '\u0400' && r <= '\u04FF':
			return true
		default:
			return false
		}
	}
	return false
}

func ContainsAlphabetic(text string) bool {
	for _, r := range text {
		if IsAlphabetic(r) {
			return true
		}
	}
	return false
}

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return destinationFile.Sync()
}

// SanitizePathName 清理字符串，使其成为合法路径名
func SanitizePathName(name string) string {
	name = strings.ReplaceAll(name, ".", "_")

	var illegalChars *regexp.Regexp
	if runtime.GOOS == "windows" {
		// Windows 特殊字符
		illegalChars = regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	} else {
		// POSIX 系统：只禁用 / 和空字节
		illegalChars = regexp.MustCompile(`[/\x00]`)
	}

	sanitized := illegalChars.ReplaceAllString(name, "_")

	// 去除前后空格
	sanitized = strings.TrimSpace(sanitized)

	// 防止空字符串
	if sanitized == "" {
		sanitized = "unnamed"
	}

	// 避免 Windows 下的保留文件名
	reserved := map[string]bool{
		"CON": true, "PRN": true, "AUX": true, "NUL": true,
		"COM1": true, "COM2": true, "COM3": true, "COM4": true,
		"LPT1": true, "LPT2": true,
	}

	upper := strings.ToUpper(sanitized)
	if reserved[upper] {
		sanitized = "_" + sanitized
	}

	return sanitized
}

// FindClosestConsecutiveWords 查找 words 中 Num 连续递增的一组词，使得其拼接后的文本与 inputStr 的编辑距离最小。
func FindClosestConsecutiveWords(words []types.Word, inputStr string) []types.Word {
	if len(words) == 0 {
		return nil
	}

	// 先将输入按 Num 排序（如果你已经保证是有序的可跳过此步骤）
	// sort.Slice(words, func(i, j int) bool { return words[i].Num < words[j].Num })

	// Step 1: 获取所有 Num 连续递增的 []types.Word 组合
	var groups [][]types.Word
	var currentGroup []types.Word

	for i, word := range words {
		if i == 0 {
			currentGroup = append(currentGroup, word)
			continue
		}

		if word.Num == words[i-1].Num+1 {
			currentGroup = append(currentGroup, word)
		} else {
			if len(currentGroup) > 0 {
				groups = append(groups, currentGroup)
			}
			currentGroup = []types.Word{word}
		}
	}
	if len(currentGroup) > 0 {
		groups = append(groups, currentGroup)
	}

	// Step 2: 比较编辑距离，找最接近 inputStr 的那个组
	minDistance := -1
	var bestGroup []types.Word

	for _, group := range groups {
		var sb strings.Builder
		for _, w := range group {
			sb.WriteString(w.Text)
		}
		groupText := sb.String()

		dist := levenshtein.DistanceForStrings([]rune(groupText), []rune(inputStr), levenshtein.DefaultOptions)

		if minDistance == -1 || dist < minDistance {
			minDistance = dist
			bestGroup = group
		}
	}

	return bestGroup
}

func SaveToDisk(data any, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 美化输出
	return encoder.Encode(data)
}

func LoadFromDisk(filename string) (any, error) {
	var data any
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	return data, err
}

// 清理 Markdown 的 ```json 标记
func CleanMarkdownCodeBlock(response string) string {
	re := regexp.MustCompile("(?m)^```(json|[a-zA-Z]*)?\n?|```$")
	cleaned := re.ReplaceAllString(response, "")
	return strings.TrimSpace(cleaned)
}
