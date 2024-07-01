# NewAssetPayloadIpfs

Set to true to make default export to IPFS. To customize the
pinned files, specify an object with a spec field. False or null
means to unpin from IPFS, but it's unsupported right now.



## Supported Types

### Ipfs1

```go
newAssetPayloadIpfs := components.CreateNewAssetPayloadIpfsIpfs1(components.Ipfs1{/* values here */})
```

### 

```go
newAssetPayloadIpfs := components.CreateNewAssetPayloadIpfsBoolean(bool{/* values here */})
```

