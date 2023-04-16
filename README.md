# Terrap by Sirrend
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) </br>
Simplify your Provider version upgrades with Terrap - a powerful CLI tool that scans your system and identifies any required changes. </br>
The tool offers clear and actionable notifications, helping you streamline the upgrade process and avoid any potential errors or complications. 

## How to Download
#### Clone the Terrap-CLI Repository
```shell
git clone https://github.com/sirrend/terrap-cli
cd terrap-cli

go build -o terrap .

chmod +x terrap
mv terrap /usr/local/bin/
```

#### Brew
```shell
brew install terrap-cli
```

## Quick Start
1. Go to your local IaC repository folder.
2. Initialize a new terrap workspace where you would run `terraform apply` with `terrap init -c`.
3. Scan your workspace with: `terrap scan`


https://user-images.githubusercontent.com/47568615/232330641-f7409ee9-def1-42a0-aace-cb4139b9d72d.mov
