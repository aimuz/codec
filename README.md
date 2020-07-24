# codec

Codec defines the interface universal uses to encode and decode messages. 

## Example

### Storm 
 
[github.com/asdine/storm](https://github.com/asdine/storm)
 
```golang
storm.Open("", storm.Codec(json.Codec))
storm.Open("", storm.Codec(gob.Codec))
storm.Open("", storm.Codec(bson.Codec))
storm.Open("", storm.Codec(xml.Codec))
storm.Open("", storm.Codec(gzip.NewGzipCodecWith(json.Codec)))
storm.Open("", storm.Codec(gzip.NewGzipCodecWith(gob.Codec)))
storm.Open("", storm.Codec(gzip.NewGzipCodecWith(bson.Codec)))
storm.Open("", storm.Codec(gzip.NewGzipCodecWith(xml.Codec)))
storm.Open("", storm.Codec(gzip.NewGzipCodecLevelWith(json.Codec, 9)))
storm.Open("", storm.Codec(gzip.NewGzipCodecLevelWith(gob.Codec, 9)))
storm.Open("", storm.Codec(gzip.NewGzipCodecLevelWith(bson.Codec, 9)))
storm.Open("", storm.Codec(gzip.NewGzipCodecLevelWith(xml.Codec, 9)))
storm.Open("", storm.Codec(snappy.NewSnappyCodecWith(json.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodecWith(gob.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodecWith(bson.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodecWith(xml.Codec)))
```

### GRPC

codec using in grpc

```golang
encoding.RegisterCodec(json.Codec)
encoding.RegisterCodec(gob.Codec)
encoding.RegisterCodec(bson.Codec)
encoding.RegisterCodec(xml.Codec)
encoding.RegisterCodec(gzip.NewGzipCodecWith(json.Codec))
encoding.RegisterCodec(gzip.NewGzipCodecWith(gob.Codec))
encoding.RegisterCodec(gzip.NewGzipCodecWith(bson.Codec))
encoding.RegisterCodec(gzip.NewGzipCodecWith(xml.Codec))
encoding.RegisterCodec(gzip.NewGzipCodecLevelWith(json.Codec, 9))
encoding.RegisterCodec(gzip.NewGzipCodecLevelWith(gob.Codec, 9))
encoding.RegisterCodec(gzip.NewGzipCodecLevelWith(bson.Codec, 9))
encoding.RegisterCodec(gzip.NewGzipCodecLevelWith(xml.Codec, 9))
encoding.RegisterCodec(snappy.NewSnappyCodecWith(json.Codec))
encoding.RegisterCodec(snappy.NewSnappyCodecWith(gob.Codec))
encoding.RegisterCodec(snappy.NewSnappyCodecWith(bson.Codec))
encoding.RegisterCodec(snappy.NewSnappyCodecWith(xml.Codec))
```

## Support

- [x] gzip
- [x] json
- [x] gob
- [x] snappy
- [x] bson
- [x] xml
- [ ] protobuf
- [ ] ....
