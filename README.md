# ✦•┈๑⋅⋯ 𝐄𝐌𝐋𝐂𝐇𝐄𝐂𝐊 ⋯⋅๑┈•✦

<img src="https://i.pinimg.com/1200x/4d/6b/d3/4d6bd38e850925dbc2ff1ca106377edd.jpg?width=770&height=578&fit=crop&format=pjpg&auto=webp" alt="img" align="right" width="400px"> <br><br>

This project is a simple command-line tool built in Go for analyzing `.eml` files and extracting relevant email headers for quick inspection. It is designed for fast triage, allowing users to identify spoofing indicators, authentication results, and suspicious patterns directly from raw email data.

The tool focuses on simplicity and speed, making it useful for beginners, analysts, or anyone needing a quick look into email metadata without using heavy forensic tools.

<br><br><br>

## 【𝑭𝒆𝒂𝒕𝒖𝒓𝒆𝒔】

* **Header Extraction**
  Parses key headers such as `From`, `Reply-To`, `Return-Path`, and `Received`.

* **Authentication Insights**
  Displays SPF, DKIM, and DMARC results when available.

* **Exchange Awareness**
  Detects and shows `X-MS-Exchange-*` headers for Microsoft environments.

* **Fast Execution**
  Built in Go for instant performance.

* **Minimal & Clean Output**
  Focused results without unnecessary noise.

---

## 【𝑰𝒏𝒔𝒕𝒂𝒍𝒍𝒂𝒕𝒊𝒐𝒏】

### Prerequisites

* Go (1.18+ recommended)

---

### 【𝑨𝒖𝒕𝒐 𝑰𝒏𝒔𝒕𝒂𝒍𝒍】

```bash
chmod +x install.sh
./install.sh
```

The script will:

* Compile the binary
* Install `emlcheck` into your system path

---

### 【𝑴𝒂𝒏𝒖𝒂𝒍 𝑰𝒏𝒔𝒕𝒂𝒍𝒍】

```bash
git clone https://github.com/yourusername/emlcheck.git
cd emlcheck
go build -o emlcheck main.go
```

Move binary:

```bash
sudo mv emlcheck /usr/local/bin/
```

Or without sudo:

```bash
mkdir -p ~/.local/bin
mv emlcheck ~/.local/bin/
```

---

## 【𝑼𝒔𝒂𝒈𝒆】

Analyze an `.eml` file:

```bash
emlcheck file.eml
```

---

### Example Output

```
From: attacker@evil.com
Reply-To: fake@gmail.com
Return-Path: spoofed@evil.com

SPF: fail
DKIM: fail
DMARC: fail

X-MS-Exchange-Organization-AuthAs: Anonymous
```

---

### Help

```bash
emlcheck --help
```

---

## 【𝑾𝒉𝒂𝒕 𝒊𝒕 𝒅𝒆𝒕𝒆𝒄𝒕𝒔】

* Mismatched headers (`From` vs `Reply-To`)
* Authentication failures (SPF / DKIM / DMARC)
* Suspicious routing chains (`Received`)
* Exchange-specific indicators
* Basic spoofing patterns

---

## 【𝑪𝒐𝒎𝒎𝒂𝒏𝒅𝒔】

| Command           | Description         |
| ----------------- | ------------------- |
| `emlcheck -r <file>` | Analyze an EML file |
| `--help`          | Show help           |

---

## 【𝑵𝒐𝒕𝒆𝒔】

* This tool is focused on **header-level analysis**
* Some headers can be spoofed; results should be correlated
* Intended for **quick inspection**, not full forensic workflows

---

## 【𝑫𝒊𝒔𝒄𝒍𝒂𝒊𝒎𝒆𝒓】

This project isn't a way to know if the eml is dangerous or not, it only shows you the initial info without noise. 

