# Etcd Operations CLI

**Author:** Morteza Khazamipour

## Overview

`etcd-operations-cli` is a command-line tool designed for backing up and defragmenting etcd clusters. It provides functionality to perform local backups as well as backups to S3-compatible storage.

## Installation

To install the CLI, run the following command:


```bash
go install github.com/mkhazamipour/etcd-operations-cli@latest
```

## Usage

### Global Flags

--config: Specify the path to the configuration file (default is $HOME/etcd-configs.yaml).

-e, --etcd-endpoints: Etcd endpoints separated by a comma (e.g., 127.0.0.1:2379,127.0.0.2:2379).

-a, --ca-cert: Etcd CA certificate file.

-k, --key: Etcd key file.

-c, --cert: Etcd cert file.


### Commands

### 1. Defrag

```bash
etcd-operations-cli defrag
```
The defrag command performs defragmentation on the specified etcd endpoints.

### 2. Backup

#### Subcommands

#### Local
```bash
etcd-operations-cli backup local
```
The local subcommand performs a backup of the etcd cluster to the local disk.

#### Flags

-l, --backup-location: Location to save the etcd backup on disk (e.g., /tmp/backup1.db).


#### S3
```bash
etcd-operations-cli backup s3
```

The s3 subcommand performs a backup of the etcd cluster to S3-compatible storage.
#### Flags

-b, --bucket-name: Bucket name to save etcd's snapshot.

-p, --s3-endpoint: S3 Region, can be MinIO endpoint.

-r, --region: S3 Region.

-n, --s3-access-key: S3 Access Key.

-s, --s3-secret-key: S3 Secret Key.

### 3. Size

```bash
etcd-operations-cli size
```

Get the DB size for each endpoint


## Configuration

The CLI supports configuration through a YAML file. An example configuration:

```yaml
etcd:
  endpoints:
    - "127.0.0.1:2379"
    - "127.0.0.2:2379"
caPath:
  cert: "/etcd/kubernetes/pki/etcd/certs/cert.crt"
  key: "/etcd/kubernetes/pki/etcd/certs/cert.key"
  cacert: "/etcd/kubernetes/pki/etcd/certs/cacert.crt"
backupLocation: "/tmp/backup"
s3BucketName: "etcd-backup"
s3Endpoint: "minio-cluster.cluster.morteza.dev"
s3Region: "eu-central"
s3AccessKey: "accesskey"
s3SecretKey: "secretkey"
```
