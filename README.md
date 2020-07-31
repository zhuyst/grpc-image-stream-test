# GRPC 一个图片的base64的两种发送方式比较测试

```proto
service StreamService {
    rpc GetImageStream(GetImageRequest) returns (stream GetImageResponse);
    rpc GetImage(GetImageRequest) returns (GetImageResponse);
}

message GetImageRequest{}

message GetImageResponse {
    bytes images = 1;
}
```

验证、比较一张1920*1080图片的stream发送与普通发送的性能差距。