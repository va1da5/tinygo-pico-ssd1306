# Sreaming Images via UART

## Serial Monitor

```bash
# Pico UART adapter
minicom -b 115200 -8 -D /dev/ttyACM0

# USB to UART adapter
minicom -b 115200 -8 -D /dev/ttyUSB0
```

## Getting Started

```bash
# build binary and flash pico
make

# start streaming frames
make stream
```

## References

- [Image2CCP](https://javl.github.io/image2cpp/)
- [Interface OLED Graphic Display Module with Arduino](https://lastminuteengineers.com/oled-display-arduino-tutorial/)
  https://lastminuteengineers.com/oled-display-arduino-tutorial/
