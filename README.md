# raspi-virtual-gps-receiver
Virtual GPS receiver to sync date and time in NMEA 0183


## Prerequisite

### Let UART0 available

Reference: [Raspberry Pi Documentation \- Configuration](https://www.raspberrypi.com/documentation/computers/configuration.html)

#### Disable Serial Console

1. execure the command below
```bash
sudo raspi-config
```

1. Select `3 Interface Options    Configure connections to peripherals`

1. Select `I6 Serial Port   Enable/disable shell messages on the serial connection`

1. Select `<No>` for `Would you like a login shell to be accessible over serial?`

1. Select `<Yes>` for `Would you like the serial port hardware to be enabled?`

#### Disable bluetooth

```bash
echo dtoverlay=disable-bt | tee -a /boot/config.txt
```

#### Reboot

```bash
sudo reboot
```

## Installation

```bash
curl -LO https://github.com/rakiyoshi/raspi-virtual-gps-receiver/releases/download/v0.0.8/pigps_0.0.8_arm32
sudo cp ./pigps_0.0.8_arm32 /usr/local/bin/pigps
sudo chmod +x /usr/local/bin/pigps
sudo cp -r ./service/pigps/* /etc/systemd/system/
sudo chown root /etc/systemd/system/pigps
sudo systemctl enable pigps
sudo systemctl start pigps
```


## Build guide

### Build
```bash
goreleaser build --snapshot --rm-dist
```

### Run

```bash
./dist/pigps_linux_arm64/pigps
# for ARM v6
# ./dist/pigps_linux_arm_6/pigps
```

### Reference
[NMEAフォーマット](http://www.hdlc.jp/~jh8xvh/jp/gps/nmea.html)
