# codec

Codec defines the interface universal uses to encode and decode messages. 

## examples

### with storm 

```golang
storm.Open("", storm.Codec(json.Codec))
storm.Open("", storm.Codec(gob.Codec))
storm.Open("", storm.Codec(gzip.NewGzipCodec(json.Codec)))
storm.Open("", storm.Codec(gzip.NewGzipCodec(gob.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodec(json.Codec)))
storm.Open("", storm.Codec(snappy.NewSnappyCodec(gob.Codec)))
```

## support

- [x] gzip
- [x] json
- [x] gob
- [x] snappy
- [ ] bson
- [ ] protobuf
- [ ] ....
