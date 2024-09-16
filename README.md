# Deployment Manager CLI

Deployment Manager is a command-line tool to manage deployments for cloud-native applications. It provides commands to install, upgrade, rollback, and uninstall deployments.

## Features

- Install new deployments
- Upgrade existing deployments
- Rollback to previous versions of deployments
- Uninstall deployments

## Requirements

- Go 1.16+
- Cobra CLI library
- Kubernetes cluster (for deployment management)

## Initialize Go Modules
Initialize a new Go module:

```bash
go mod init deployment-manager
```

## Install Dependencies
To install the required dependencies (like the Cobra library), use:

```bash
go get github.com/spf13/cobra
```

## Build the tool

Clone the repository and build the tool:

```bash
git clone https://github.com/yourusername/deployment-manager.git
cd deployment-manager
go build -o deployment-manager main.go
```

## Usage
Once installed, the CLI provides several commands:

## Install a deployment
```bash
deployment-manager install --namespace <namespace> --timeout <seconds> --ignore=<true|false>
--namespace (required): The namespace where the deployment will be installed.
--timeout (required): Timeout in seconds for the deployment.
--ignore (optional): Whether to ignore pre-install checks. Default is false.
```

## Upgrade a deployment
```bash
deployment-manager upgrade --version <version>
--version (required): The version to upgrade the deployment to.
```
## Rollback a deployment
```bash
deployment-manager rollback --version <version>
--version (required): The version to rollback the deployment to.
```

## Uninstall a deployment
```bash
deployment-manager uninstall --namespace <namespace>
--namespace (required): The namespace where the deployment will be uninstalled.
```