# natat

natat is an abbreviation for nat assert. It's a NAT sniffing tool that you can use to determine your NAT type.

### Installation

Go get with:

```
go get github.com/songjiayang/natat
```

### Usage

- Testing with Symmetric NAT

```
$ natat
2019/04/12 16:39:33 start stun server ping...
2019/04/12 16:39:34 stun.l.google.com:19302 mapped: 0.0.0.0:3489 -> 103.90.76.164:35598
2019/04/12 16:39:34 stun1.l.google.com:19302 mapped: 0.0.0.0:3489 -> 45.64.52.117:6078
2019/04/12 16:39:34 start NAT type assert...
2019/04/12 16:39:34 It's Symmetric NAT
```

- Testing with Cone NAT

```
$ natat
2019/04/12 20:42:47 start stun server ping...
2019/04/12 20:42:48 stun.l.google.com:19302 mapped: 0.0.0.0:3489 -> 101.84.52.37:26262
2019/04/12 20:42:48 stun1.l.google.com:19302 mapped: 0.0.0.0:3489 -> 101.84.52.37:26262
2019/04/12 20:42:48 start NAT type assert...
2019/04/12 20:42:48 It's Cone NAT
```

- More configuration:

```
$ natat -h
Usage of natat:
  -bind string
    	ping with local address bind. (default "0.0.0.0:3489")
  -studs string
    	stun servers for ping. (default "stun.l.google.com:19302,stun1.l.google.com:19302")
```

###  Supported NAT Type

- [ ] Full Cone
- [ ] Restricted Cone
- [ ] Port Restricted Cone
- [x] Symmetric

### Thanks

- RFC [3489](https://www.ietf.org/rfc/rfc3489.txt)

-  stun servers: 
    * stun.l.google.com:19302
    * stun1.l.google.com:19302
    * stun2.l.google.com:19302
    * stun3.l.google.com:19302
    * stun4.l.google.com:19302
    * [more](https://gist.github.com/mondain/b0ec1cf5f60ae726202e)