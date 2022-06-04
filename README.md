# Ransomeware Canary

Backup strategies are the main tool to mitigate and limit the harm ransomware attacks can do to companies and individuals. Often times backups run periodically and keep a limited number of recent backups (e.g. keep 4 recent backup runs). While this is a fair way to recover files lost to encryption malware, systems may overwrite the previous backups with backups of the encrypted files after an attack has happened. This can especially be the case if systems are unmonitored for a longer time.

This project's release serves as a canary to detect ransomware runs while they happen. It places a file in the home directory of the executing user and watches it for deletion or changes. Upon detection of such file event, it will log, send an email to addresses in its configuration file and exit.

## Build

Prerequisites: 
- Go 1.8.2  ([Installation Instructions](https://go.dev/doc/install))

```shell
go build -v ./...
```

## Install

By default it will build and install to `/usr/local/bin/ranscanary`. If you wish to change the prefix or binary name, adjust the `Makefile` before running: 

```shell
sudo make install
```

After installation, place the configuration at `/usr/local/etc/ranscanary/config.toml` or a location to your liking (and specify using `-config` when running).

## Configure

Make a copy of `config.example.toml` to `config.toml` and edit options.

Available configuration fields:
- `CanaryFileName`: the file name of the canary document (`string`)
- `CanaryDocument`: the content of the canary document (`string`)
- `SmtpHost`: the hostname of the SMTP server to send alerts (`string`)
- `SmtpPort`: the port of the SMTP server (`int`)
- `SmtpUser`: the username to authenticate at the SMTP server (`string`)
- `SmtpPass`: the password to authenticate at the SMTP server (`string`)
- `SmtpFrom`: the sender's email address (`string`)
- `SmtpTo`: the recipients' email addresses (`string array`)
- `SmtpSubject`: the alert email's subject (`string`)


## Run

Assuming the default installation parameters and `/usr/local/bin` is in your `PATH`:

```shell
ranscanary -config $HOME/ransomware_canary/config.toml # Specify configuration file path
# OR
ranscanary # will use the default configuration path at /usr/local/etc/ranscanary/config.toml
```
