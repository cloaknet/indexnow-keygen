# IndexNow KeyGen

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/cloaknet/indexnow-keygen)](https://github.com/cloaknet/indexnow-keygen/releases)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)]()

🔐 A secure command-line tool written in Go for generating IndexNow API keys compatible with major search engines.

## ✨ Features

- 🔐 **Cryptographically secure** key generation using `crypto/rand`
- 📏 **Configurable key length** (64-96 characters by default)
- 📁 **Automatic file organization** in `static/` directory
- 🔄 **Batch generation** - create multiple keys at once
- ✅ **IndexNow compliant** - follows official specifications
- 🛡️ **Robust error handling** - graceful recovery and reporting
- 🎯 **User-friendly** - interactive prompts and clear output
- 🌐 **Cross-platform** - works on Windows, macOS, and Linux

## 🚀 Quick Start

```bash
# Download and run
curl -L -O https://github.com/cloaknet/indexnow-keygen/releases/latest/download/indexnow-keygen
chmod +x indexnow-keygen
./indexnow-keygen
```

## 🔍 What is IndexNow?

IndexNow is a protocol that allows websites to instantly inform search engines about content changes. Submit your URLs to get them indexed faster!

**Supported search engines:**
- **IndexNow.org** (Official)
- **Microsoft Bing**
- **Yandex**
- **Naver**
- **Seznam.cz**
- **Yep**

## 📦 Installation

### Option 1: Download Binary

Download the latest release from the [Releases](https://github.com/cloaknet/indexnow-keygen/releases) page.

### Option 2: Build from Source

```bash
git clone https://github.com/cloaknet/indexnow-keygen.git
cd indexnow-keygen
go build -o indexnow-keygen
./indexnow-keygen
```

### Option 3: Install with Go

```bash
go install github.com/cloaknet/indexnow-keygen@latest
```

## 🎮 Usage

1. **Run the program:**
   ```bash
   ./indexnow-keygen
   ```

2. **Enter the number of keys:**
   ```
   Enter the number of keys to generate: 3
   ```

3. **Keys generated automatically:**
   ```
   1) Key: abc123def456ghi789jkl012mno345pqr678stu901vwx234yz567abc123def
      Saved to: static/abc123def456ghi789jkl012mno345pqr678stu901vwx234yz567abc123def.txt
   2) Key: def456ghi789jkl012mno345pqr678stu901vwx234yz567abc123def456ghi
      Saved to: static/def456ghi789jkl012mno345pqr678stu901vwx234yz567abc123def456ghi.txt
   3) Key: ghi789jkl012mno345pqr678stu901vwx234yz567abc123def456ghi789jkl
      Saved to: static/ghi789jkl012mno345pqr678stu901vwx234yz567abc123def456ghi789jkl.txt
   ```

## 📋 Key Specifications

Generated keys meet IndexNow requirements:

| Parameter | Value |
|-----------|-------|
| **Length** | 64-96 characters |
| **Characters** | `a-z`, `A-Z`, `0-9`, `-` |
| **Format** | Plain text files |
| **Naming** | Filename matches key content |
| **Location** | `static/` directory |

## 📂 Project Structure

```
your-website/
├── indexnow-keygen           # Executable
├── static/                   # Generated keys directory
│   ├── abc123...def.txt     # Key file 1 (key as filename)
│   ├── def456...ghi.txt     # Key file 2 (key as filename)
│   └── ghi789...jkl.txt     # Key file 3 (key as filename)
└── README.md
```

## 🌐 Using Your Generated Keys

### Step 1: Upload Key Files

Upload the generated `.txt` files to your website's root directory or `/static/` folder.

### Step 2: Choose One Key

Select any one of your generated keys for IndexNow submissions.

### Step 3: Submit URLs

#### 📝 POST Request (Recommended for Multiple URLs)

```bash
curl -X POST "https://api.indexnow.org/indexnow" \
  -H "Content-Type: application/json" \
  -d '{
    "host": "example.com",
    "key": "your-generated-key-here",
    "urlList": [
      "https://example.com/page1",
      "https://example.com/page2",
      "https://example.com/page3"
    ]
  }'
```

#### 🔗 GET Request (Single URL)

Replace `https://example.com/new-page` with your URL and `your-key` with your generated key:

| Search Engine | API Endpoint |
|---------------|--------------|
| **IndexNow (Official)** | `https://api.indexnow.org/indexnow?url=https://example.com/new-page&key=your-key` |
| **Microsoft Bing** | `https://www.bing.com/indexnow?url=https://example.com/new-page&key=your-key` |
| **Yandex** | `https://yandex.com/indexnow?url=https://example.com/new-page&key=your-key` |
| **Naver** | `https://searchadvisor.naver.com/indexnow?url=https://example.com/new-page&key=your-key` |
| **Seznam.cz** | `https://search.seznam.cz/indexnow?url=https://example.com/new-page&key=your-key` |
| **Yep** | `https://indexnow.yep.com/indexnow?url=https://example.com/new-page&key=your-key` |

#### 💡 Example Usage

```bash
# Submit to IndexNow
curl "https://api.indexnow.org/indexnow?url=https://example.com/new-article&key=abc123def456..."

# Submit to Bing
curl "https://www.bing.com/indexnow?url=https://example.com/new-article&key=abc123def456..."

# Submit to Yandex
curl "https://yandex.com/indexnow?url=https://example.com/new-article&key=abc123def456..."
```

## ⚙️ Configuration

Modify constants in source code if needed:

```go
const (
    minKeyLength = 64      // Minimum key length
    maxKeyLength = 96      // Maximum key length
    outputDir   = "static" // Output directory
)
```

## 🔒 Security Features

- **Cryptographically secure** random generation via `crypto/rand`
- **High entropy**: 382-573 bits per key (enterprise-grade security)
- **No predictable patterns** in generated keys
- **Collision detection** - prevents duplicate file creation
- **Compliance** with IndexNow security standards

## 🛠️ Error Handling

Comprehensive error management:

- ✅ Directory creation failures
- ✅ File write permission issues
- ✅ Invalid user input validation
- ✅ Cryptographic generation errors
- ✅ Duplicate file detection
- ✅ Graceful recovery and detailed reporting
- ✅ Summary of failed operations

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Useful Links

- [IndexNow Official Documentation](https://www.indexnow.org/)
- [Yandex Webmaster IndexNow Guide](https://yandex.ru/support/webmaster/indexnow/)
- [Microsoft Bing IndexNow Documentation](https://www.bing.com/indexnow)
- [IndexNow Protocol Specification](https://www.indexnow.org/documentation)

## 📝 Changelog

### v1.0.0 (Initial Release)
- ✨ Secure key generation with `crypto/rand`
- ✨ Batch processing support
- ✨ Cross-platform compatibility
- ✨ Comprehensive error handling
- ✨ User-friendly CLI interface

---

<div align="center">

**Made with ❤️ for faster web indexing**

[⭐ Star this project](https://github.com/cloaknet/indexnow-keygen) | [🐛 Report Bug](https://github.com/cloaknet/indexnow-keygen/issues) | [💡 Request Feature](https://github.com/cloaknet/indexnow-keygen/issues)

</div>