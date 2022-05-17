# Scale Codec in Golang

Scale Codec written by golang, reference from [polkascan/py-scale-codec](https://github.com/polkascan/py-scale-codec).

## Features

- [x] low level codec decode
- [x] metadata decode
- [x] extrinsic decode
- [x] LogDigest decode
- [x] event decode
- [x] custom type reg
- [x] reg special version type
- [ ] low level codec encode

## Installation 

```bash
go get -u github.com/yspk/scale.go
```

## Test

```bash
go test ./... --cover -v
```

## Resources

- [Subscan](https://www.subscan.io)
- [Subscan Open-source](https://github.com/itering/subscan-essentials)
- [Polkadot.js](http://polkadot.js.org/)

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/yspk/scale.go

## License

The package is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
