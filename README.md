
# 🚀 Mojo GoJo: Write Go with Emojis!

Welcome to Mojo GoJo, where Go meets Emojis to make your coding experience delightful and colorful! (You will hate it)


## 📦 Installation

### 1. Pre-built Binaries:

#### Linux:

```bash
# Download the binary
wget https://github.com/EdmondTabaku/mojo-gojo/releases/download/v1.0.0/mojo-gojo-linux -O mojo-gojo

# Make the binary executable
chmod +x mojo-gojo

# Move the binary to a directory in your PATH
sudo mv mojo-gojo /usr/local/bin/
```


#### macOS:

```bash
# Download the binary
curl -Lo mojo-gojo https://github.com/EdmondTabaku/mojo-gojo/releases/download/v1.0.0/mojo-gojo-mac

# Make the binary executable
chmod +x mojo-gojo

# Move the binary to a directory in your PATH
sudo mv mojo-gojo /usr/local/bin/
```

#### Windows:
Download the .exe file from [mojo-gojo.exe](https://github.com/EdmondTabaku/mojo-gojo/releases/download/v1.0.0/mojo-gojo.exe).  

Move the .exe file to a directory of your choice.  

Add the directory to your system's PATH.  

Open Command Prompt and verify the installation with ```mojo-gojo```

### 2. Build from Source:
All Operating Systems:
```bash
# Clone the repository
git clone https://github.com/EdmondTabaku/mojo-gojo.git
cd mojo-gojo

# Build the project
go build -o mojo-gojo ./cmd/mojo-gojo

# Move the binary to a directory in your PATH (you might need sudo for Linux/macOS)
mv mojo-gojo /path/to/directory/in/your/PATH/
```


## 📖 How to Use Mojo GoJo

With Mojo GoJo, you write Go code using our emoji-keyword mappings. The mojo-gojo files have the ```.mg``` suffix. Once your code is ready, use the mojo-gojo command-line tool to convert your emoji-filled code into standard Go, which you can then run or build.

### 🗺️ Emoji-to-Keyword Mappings
Here's a quick reference of our emoji-to-keyword mappings:

| Emoji  | Go Keyword   |
|:------:|:------------:|
| 📦     | package      |
| 🧙     | func         |
| 🔄     | for          |
| 📥     | import       |
| ✋     | break        |
| 🎲     | default      |
| 🎭     | interface    |
| 🎯     | select       |
| 🧳     | case         |
| 😴     | defer        |
| 🚗     | go           |
| 🗺️    | map          |
| 🏗️    | struct       |
| 📫     | chan         |
| 🤷     | else         |
| 🚶‍♂️  | goto         |
| 🔀     | switch       |
| 🔒     | const        |
| 🤸     | fallthrough  |
| 🤔     | if           |
| 📏     | range        |
| 🔖     | type         |
| ⏭️    | continue     |
| 📤     | return       |
| 🤲     | var          |

### 🖋️ Sample Code

Here's a fun example of how you can write a simple program using Mojo GoJo:

```go
📦 main

📥 (
    "fmt"
)

🧙 main() {
    fmt.Println("Mojooo Gojooo World!")
}
```

### 🚀 Running Your Code
After writing your .mg file, simply use the mojo-gojo tool to run it:

```bash
mojo-gojo run your/project/directory
```


## 🤝 Contributing
Love Mojo GoJo? (No)   
Want to make it even better? (No)  
If yes then do your thing, and create a pr. (Deleting the project)




