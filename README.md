# Terrap, by Sirrend
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/sirrend/terrap-cli?filename=go.mod)</br>
<img src="./docs/terrap-cover.png"/>
</br></br>
Simplify your Provider version upgrades with **Terrap** - a powerful CLI tool that scans your system and identifies any required changes. </br>
The tool offers clear and actionable notifications, helping you streamline the upgrade process and avoid any potential errors or complications.</br></br>
🔍 **Terrap is an alpha version project, therefore some data might be partial.**

## Resources
* Documentation - <a href="https://www.sirrend.com/terrap-docs">sirrend.com/terrap-docs</a>

## Constraints 🧱
1. Supported Terraform Core versions: `>=0.13`.
2. Every provider which uses `Terraform Core 0.13` or higher.

## Good To Know 💡
Terrap decides which Terraform version to use in the following order:
1. The latest installed Terraform version found locally.
2. If the `TERRAP_TERRAFORM_VERSION` environment variable is set, Terrap will use the version specified as long as it matches the `>=0.13` constraint.</br>
    Set environment variable on mac/linux:</br>
    ```shell
   export <var>=<value>
   ```
   Set environment variable on windows:</br>
    ```shell
   $Env:TERRAP_TERRAFORM_VERSION = "0.13"
   ```
   
4. If none of the above is applicable, Terrap will download the latest available version.

## How to Download ⬇️
### Clone sirrend/terrap-cli
```shell
git clone https://github.com/sirrend/terrap-cli
cd terrap-cli

go build -o terrap .

chmod +x terrap
mv terrap /usr/local/bin/
```

### Brew
```shell
brew tap sirrend/sirrend
brew install terrap
```

Validate terrap is working by executing `terrap`.

## Quick Start ⏩

### Initialize my First Workspace
1. `CD` to the local Terraform folder you want to work with.</br>
   `cd < /terraform/folder/path >`</br></br>

2. Initialize a new Terrap workspace where you would run terraform apply with `terrap init`.</br></br>
    <strong>Important!</strong> </br>
    As Terrap runs <code>terraform init</code> under the hood, it would need every configuration component you normally use when executing the command.</br>
    It can be environment variables, the <code>.aws/credentials</code> file, etc.


3. Scan your workspace with: `terrap scan`

https://user-images.githubusercontent.com/47568615/232331582-998cb9dc-4ad4-465e-af31-4fab0c77877b.mov

## Features 🚀
### Scan for changes with `scan`
Scan your infrastructure for changes in the following provider version for a safe and easy upgrade!</br>
Looking for a specific resource type changes? Use the `--data-sources` `--resources` and `--provider` flags.

### Stay up-to-date with `whats-new`
Ready to explore what's new in the following version of your provider? Simply execute `terrap whats-new`.</br>
Looking to delve into a specific version of your provider? Specify the desired version with `--fixed-providers <provider>:<version>` to explore what's new and improved.

### Which Providers are Supported?
Run `terrap providers get-supported` to get a list of all supported providers and version ranges.</br>
You can use the `--filter` flag if you're looking for something specific.


## What's the Future Hold 🔮
1. Bulk updates - straight to a version of your choosing.
2. Automatic Upgrades - you write, **Terrap** upgrades.
3. Expanding Terrap's providers support.

## Something's Wrong? Tell Us! 🚨
You can open an issue either directly from the CLI using `terrap open-issue` or through the GitHub UI.

## Want to contribute? 🍀 Lucky us!
1. Checkout from the `main` branch.
2. Add your code with the proper documentation.
3. Open a PR with a detailed explanation of the functionality you want to add.


