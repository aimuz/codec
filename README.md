# codec

Codec defines the interface universal uses to encode and decode messages. 

## Example

### Storm 
 
[github.com/asdine/storm](https://github.com/asdine/storm)
 
```golang
storm.Open("", storm.Codec(json.Codec))
storm.Open("", storm.Codec(gob.Codec))
storm.Open("", storm.Codec(gzip.NewGzipCodec(json.Codec)))
storm.Open("", storm.Codec(gzip.NewGzipCodec(gob.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodec(json.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodec(gob.Codec)))
```

### GRPC

codec using in grpc

```golang
encoding.RegisterCodec(json.Codec)
encoding.RegisterCodec(gob.Codec)
encoding.RegisterCodec(gzip.NewGzipCodec(json.Codec))
encoding.RegisterCodec(gzip.NewGzipCodec(gob.Codec))
encoding.RegisterCodec(snappy.NewSnappyCodec(json.Codec))
encoding.RegisterCodec(snappy.NewSnappyCodec(gob.Codec))
```

## Support

- [x] gzip
- [x] json
- [x] gob
- [x] snappy
- [x] bson
- [ ] protobuf
- [ ] ....
