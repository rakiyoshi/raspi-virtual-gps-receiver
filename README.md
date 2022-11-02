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


## Build

```bash
goreleaser build --snapshot --rm-dist
```

## Reference
[NMEAフォーマット](http://www.hdlc.jp/~jh8xvh/jp/gps/nmea.html)
