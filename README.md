# Terrap, by Sirrend
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/sirrend/terrap-cli?filename=go.mod)</br>
<img src="./docs/terrap-cover.png"/>
</br></br>
Simplify your Provider version upgrades with **Terrap** - a powerful CLI tool that scans your system and identifies any required changes. </br>
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
</br>

https://user-images.githubusercontent.com/47568615/232331582-998cb9dc-4ad4-465e-af31-4fab0c77877b.mov


## Features
### Scan for changes with `scan`
Scan your infrastructure for changes in the following provider version for a safe and easy upgrade!</br>
Looking for a specific resource type changes? Use the `--data-sources` `--resources` and `--provider` flags.

### Stay up-to-date with `whats-new`
Ready to explore what's new in the following version of your provider? Simply execute `terrap whats-new`.</br>
Looking to delve into a specific version of your provider? Specify the desired version with `--fixed-providers <provider>:<version>` to explore what's new and improved.

### Which Providers are Supported?
Run `terrap providers get-supported` to get a list of all supported providers and version ranges.</br>
You can use the `--filter` flag if you're looking for something specific.


## What's the Future Hold?
1. Bulk updates - straight to a version of your choosing.
2. Automatic Upgrades - you write, **Terrap** upgrades.
3. Expanding Terrap's providers support.

## Something's Wrong? Tell Us!
You can open an issue either directly from the CLI using `terrap open-issue` or through the GitHub UI.

## Want to contribute? 🍀 Lucky us!
1. Checkout from the `main` branch.
2. Add your code with the proper documentation.
3. Open a PR with a detailed explanation of the functionality you want to add.


