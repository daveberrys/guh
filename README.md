<div style="display: flex; gap: 15px;">
    <div>
        <img src=".github/readme/Guh.svg" width="128" height="128" alt="Guh">
    </div>
    <div>
        <h1 style="color: #F03C2E;"> GUH (Git Utility(For) Helper(Noobs)) </h1>
        <span>Want to use the Git CLI and don't want to use the GitHub Desktop because you're "too cool"? But you don't know how to even use the git CLI? We've got you covered.</span>
    </div>
</div>

---

This tool is a simple CLI tool that wraps the git CLI (from `os/exec`) and provides a simple command-line interface for new Git CLI users.

## How to use it?
Check out the [documentation](DOCS.md) for more information. That has all the CLI commands and usage instructions.

## Installation
### Nightly Downloads
You can download the latest nightly build from the [GitHub Actions](https://github.com/daveberrys/guh/actions) page. Or if you're lazy, here's the nightly.link
- **Windows**: [nightly.link](https://nightly.link/daveberrys/guh/workflows/compile.yaml/main/GUH-Windows.exe.zip)
- **Linux**: [nightly.link](https://nightly.link/daveberrys/guh/workflows/compile.yaml/main/GUH-Linux.zip)
- **macOS**: [nightly.link](https://nightly.link/daveberrys/guh/workflows/compile.yaml/main/GUH-macOS.zip)

### Build from source
#### Requirements
- [Go 1.2x or higher](https://go.dev/dl/)

#### Unix:
```bash
git clone https://github.com/daveberrys/guh.git
cd guh
go mod download
bash install_guh.sh
```

#### Windows:
```powershell
git clone https://github.com/daveberrys/guh.git
cd guh
go mod download
go mod tidy
go build
```

# LICENSE
[MIT License](LICENSE)