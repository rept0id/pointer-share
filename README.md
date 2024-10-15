# pointer-share

![](doc/assets/img/header1.png)

An UNSAFE pointer shared between 2 programs. For educational purposes.

```
sudo make build
```

```
./src/writer/writer 57
```
EXAMPE OUTPUT : `Memory address: 824634289880 PID: 51428`

```
sudo ./src/reader/reader 51428 824634289880
```
EXAMPLE OUTPUT : `Value at address 824634289880: 57`
