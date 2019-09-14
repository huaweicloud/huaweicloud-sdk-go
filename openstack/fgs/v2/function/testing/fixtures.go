package testing

import (
    "github.com/gophercloud/gophercloud/openstack/fgs/v2/function"
)

var ListBody = `
{
  "Functions": [
   {
    "FuncUrn": "urn:fss:cn-north-1:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test",
    "FuncName": "test",
    "UserDomain": "cff01_hk",
    "Namespace": "7aad83af3e8d42e99ac194e8419e2c9b",
    "ProjectName": "cn-north-1",
    "Package": "default",
    "Runtime": "Node.js6.10",
    "Timeout": 3,
    "Handler": "test.handler",
    "MemorySize": 128,
    "Cpu": 300,
    "CodeType": "inline",
    "CodeUrl": "",
    "CodeFileName": "index.js",
    "CodeSize": 272,
    "UserData": "",
    "Digest": "decbce6939297b0b5ec6d1a23bf9c725870f5e69fc338a89a6a4029264688dc26338f56d08b6535de47f15ad538e22ca66613b9a46f807d50b687bb53fded1c6",
    "Version": "latest",
    "ImageName": "latest-5qe8e",
    "Xrole": "",
    "DependencyPkg": "",
    "Description": "111",
    "VersionDescription": "",
    "LastModified": "2018-03-28T11:30:32+08:00"
   }
 ],
  "NextMarker": 45
}
`

var GetFunctionMetadata = `
{
  "CodeFileName": "index.js",
  "CodeSize": 272,
  "CodeType": "inline",
  "CodeUrl": "",
  "Cpu": 300,
  "DependencyPkg": "",
  "Description": "111",
  "Digest": "decbce6939297b0b5ec6d1a23bf9c725870f5e69fc338a89a6a4029264688dc26338f56d08b6535de47f15ad538e22ca66613b9a46f807d50b687bb53fded1c6",
  "FuncName": "test",
  "FuncUrn": "urn:fss:cn-north-1:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:latest",
  "Handler": "test.handler",
  "ImageName": "latest-5qe8e",
  "LastModified": "2018-03-28T11:30:32+08:00",
  "MemorySize": 128,
  "Namespace": "7aad83af3e8d42e99ac194e8419e2c9b",
  "Package": "default",
  "ProjectName": "cn-north-1",
  "Runtime": "Node.js6.10",
  "Timeout": 3,
  "UserData": "",
  "UserDomain": "cff01_hk",
  "Version": "latest",
  "VersionDescription": "",
  "Xrole": "cff"
}
`

var GetFunctionCode = `
{
  "CodeFileName": "index.js",
  "CodeSize": 272,
  "CodeType": "inline",
  "CodeUrl": "",
  "DependencyPkg": "",
  "Digest": "decbce6939297b0b5ec6d1a23bf9c725870f5e69fc338a89a6a4029264688dc26338f56d08b6535de47f15ad538e22ca66613b9a46f807d50b687bb53fded1c6",
  "FuncCode": {
    "File": "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGV2ZW50LCBjb250ZXh0LCBjYWxsYmFjaykgewogICAgZXZlbnQuZnVuY3Rpb25uYW1lID0gJ3Rlc3QnOwogICAgZXZlbnQucmVxdWVzdElkID0gY29udGV4dC5yZXF1ZXN0SWQ7CiAgICBldmVudC50aW1lc3RhbXAgPSAobmV3IERhdGUoKSkudG9TdHJpbmcoKTsKCiAgICBjb25zb2xlLmxvZygnZXZlbnQ6JywgSlNPTi5zdHJpbmdpZnkoZXZlbnQpKTsKICAgIGNhbGxiYWNrKG51bGwsIEpTT04uc3RyaW5naWZ5KGV2ZW50KSk7Cn0="
  },
  "FuncName": "test",
  "FuncUrn": "urn:fss:cn-hk1:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:latest",
  "LastModified": "2018-03-21T10:04:13+08:00",
  "Runtime": "Node.js6.10"
}
`

var CreateFunction = `
{
  "CodeFileName": "test.py",
  "CodeSize": 110,
  "CodeType": "inline",
  "CodeUrl": "",
  "Cpu": 300,
  "DependencyPkg": "",
  "Description": "",
  "Digest": "1c8610d1870731a818a037f1d2adf3223e8ac351aeb293fb1f8eabd2e9820069a61ed8b5d38182e760adc33a307d0e957afc357f415cd8c9c3ff6f0426fd85cd",
  "FuncName": "testfunc",
  "FuncUrn": "urn:fss:cn-hk1:7aad83af3e8d42e99ac194e8419e2c9b:function:default:testfunc:latest",
  "Handler": "test.handler",
  "ImageName": "latest-kyfoz",
  "LastModified": "2018-03-27T10:39:53+08:00",
  "MemorySize": 128,
  "Namespace": "7aad83af3e8d42e99ac194e8419e2c9b",
  "Package": "default",
  "ProjectName": "cn-hk1",
  "Runtime": "Node.js6.10",
  "Timeout": 3,
  "UserData": "",
  "UserDomain": "cff01_hk",
  "Version": "latest",
  "VersionDescription": "",
  "Xrole": ""
}
`

var UpdateFunctionCode = `
{
  "CodeFileName": "com.java",
  "CodeSize": 272,
  "CodeType": "inline",
  "CodeUrl": "",
  "Concurrency": 0,
  "DependList": null,
  "Dependencies": null,
  "DependencyPkg": "",
  "Digest": "56ea250008083af940a908e815b9e68d79a7d93f8cbb4dd881beefa667eb5647cf68881c9f8a7db4ac5856a642205bcb95db68da7bdbf66fa2953d90fb2ac5e5",
  "FuncCode": {
    "File": "",
    "Link": ""
  },
  "FuncName": "TestZip",
  "FuncUrn": "urn:fss:southchina:c3b2459a6d5e4b548e6777e57852692d:function:default:TestZip:latest",
  "LastModified": "2018-12-14T16:32:50+08:00",
  "Runtime": "Node.js6.10",
  "StrategyConfig": {
    "Concurrency": -1
  },
  "Subnet": "",
  "SubnetId": "",
  "Vpc": "",
  "VpcId": ""
}
`

var UpdateFunctionMetadata = `{
  "CodeFileName": "index.js",
  "CodeSize": 173,
  "CodeType": "inline",
  "CodeUrl": "",
  "Cpu": 300,
  "DependencyPkg": "",
  "Description": "",
  "Digest": "fe3d5f53536af46985e59332ece1a2dca785309937af90502745a565cf0c9ba8ebec45d70d2f31a83bd13cbf618a0ffb4e60470389da494a214213e15e4049fb",
  "FuncName": "HelloWorld",
  "FuncUrn": "urn:fss:cn-hk1:284f722ca42d465397ada4d9381c5eac:function:default:HelloWorld:latest",
  "Handler": "index.handler",
  "ImageName": "latest-fhzsf",
  "LastModified": "2018-08-23T14:17:17+08:00",
  "MemorySize": 128,
  "Namespace": "284f722ca42d465397ada4d9381c5eac",
  "Package": "default",
  "ProjectName": "cn-hk1",
  "Runtime": "Node.js6.10",
  "StrategyConfig": {
    "Concurrency": 0
  },
  "Subnet": "",
  "SubnetId": "",
  "Timeout": 3,
  "UserData": "",
  "UserDomain": "ecmuseryy",
  "Version": "latest",
  "VersionDescription": "",
  "Vpc": "",
  "VpcId": "",
  "Xrole": ""
}
`

var asInvoke = `{"request-id": "e834cb5b-1b2b-4c6b-b41c-8bd10fd41826"}`

var (
    FuncOne = function.Function{
        Handler:  "test.handler",
        Package:  "default",
        Runtime:  "Node.js6.10",
        CodeType: "inline",
        FuncName: "TestZip",
    }
)

var CreateVersion = `
{
  "CodeFileName": "index.js",
  "CodeSize": 272,
  "CodeType": "inline",
  "CodeUrl": "",
  "Cpu": 300,
  "DependencyPkg": "",
  "Description": "111",
  "Digest": "decbce6939297b0b5ec6d1a23bf9c725870f5e69fc338a89a6a4029264688dc26338f56d08b6535de47f15ad538e22ca66613b9a46f807d50b687bb53fded1c6",
  "FuncName": "test",
  "FuncUrn": "urn:fss:cn-north-1:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:v20180329-163907",
  "Handler": "test.handler",
  "ImageName": "v20180329-163907",
  "LastModified": "2018-03-29T16:39:07+08:00",
  "MemorySize": 128,
  "Namespace": "7aad83af3e8d42e99ac194e8419e2c9b",
  "Package": "default",
  "ProjectName": "cn-north-1",
  "Runtime": "Node.js6.10",
  "Timeout": 3,
  "UserData": "",
  "UserDomain": "cff01_hk",
  "Version": "v20180329-163907",
  "VersionDescription": "test",
  "Xrole": "cff"
}
`
var VersionList = `{
  "Versions": [{
    "CodeFileName": "index.js",
    "CodeSize": 272,
    "CodeType": "inline",
    "CodeUrl": "",
    "Cpu": 300,
    "Description": "",
    "FuncName": "test",
    "Handler": "index.handler",
    "LastModified": "2018-03-21T10:04:13+08:00",
    "MemorySize": 128,
    "Namespace": "7aad83af3e8d42e99ac194e8419e2c9b",
    "Package": "default",
    "ProjectName": "cn-hk1",
    "Runtime": "Node.js6.10",
    "Timeout": 3,
    "UserData": "",
    "UserDomain": "cff01_hk",
    "Version": "v20180329-163907",
    "VersionUrn": "urn:fss:cn-north-1:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:latest",
    "Xrole": ""
  }],
  "NextMarker": 1
}
`

var VersionAlias = `
{
  "Name": "dev",
  "Version": "latest",
  "Description": "",
  "LastModified": "2017-06-26 03:21:10",
  "AliasUrn": "urn:fss:cn-north-1: 7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:!dev"
}
`

var ListVersionAlias = `
[
    {
        "name": "dev",
        "version": "latest",
        "description": "",
        "last_modified": "2017-06-26 03:21:10",
        "alias_urn": "urn:fss:cn-north-1: 7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:!dev"
    }
]
`

var Version = function.CreateVersionOpts{
    Version: "v20180329-163907",
}

var Alias = function.AliasResult{
    Name:         "dev",
    Version:      "latest",
    Description:  "",
    LastModified: "2017-06-26 03:21:10",
    AliasUrn:     "urn:fss:cn-north-1: 7aad83af3e8d42e99ac194e8419e2c9b:function:default:test:!dev",
}
