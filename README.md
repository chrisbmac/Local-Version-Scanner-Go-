# Local-Version-Scanner-Go-

# 🔍 versions — A Fast, Friendly Developer Environment Scanner

`versions` is a lightweight, cross‑platform CLI tool that scans your machine and shows all your installed developer tool versions in a clean, readable table.  
It’s perfect for onboarding, debugging “works on my machine” issues, and keeping your environment consistent across teams.

---

## ✨ Features

- 🚀 **Instant environment scan**  
  Detects versions of common developer tools (Git, Go, Node, Java, Docker, Python, Rust, Terraform, and more)

- 🎨 **Beautiful terminal output**  
  Color‑coded table with emojis for each tool

- ⚡ **Fast & concurrent**  
  Runs all version checks in parallel

- 🧹 **Clean version formatting**  
  Multi‑line version outputs are automatically trimmed

- 🧩 **Modular detector system**  
  Easy to add new tools or customize your own

- 🛠️ **Built in Go**  
  Small, fast, portable static binaries

---

## 📦 Installation (macOS)

Download the latest binary from the Releases page:
For MacOS
```
sudo sh -c "$(curl -fsSL https://raw.githubusercontent.com/chrisbmac/Local-Version-Scanner-Go-/main/install.sh)"
```
Then run `versions` to see the output of your dev versions!
