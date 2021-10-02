# lookres

This utility can be used to **look**up and **res**olve information about IP address or domain name.

## Usage

```
./lookres <IPv4 or IPv6 or domain name>
```

## Examples

```
./lookres 127.0.0.1

IP: 127.0.0.1
Domains: localhost
Loopback: yes
IPv4: 127.0.0.1
IP addresses: 127.0.0.1
CNAME: localhost.
```

```
./lookres localhost

IP: 127.0.0.1
Domains: localhost
Loopback: yes
IPv4: 127.0.0.1
IP addresses: 127.0.0.1
CNAME: localhost.
```

```
./lookres google.com

IP: 74.125.131.138
Domains: lu-in-f138.1e100.net.
Global unicast: yes
IPv4: 74.125.131.138
IPv6: 2a00:1450:4010:c1e::65
IP addresses: 74.125.131.113, 74.125.131.102, 74.125.131.101, 74.125.131.139, 74.125.131.100, 74.125.131.138, 2a00:1450:4010:c1e::8a, 2a00:1450:4010:c1e::64, 2a00:1450:4010:c1e::71, 2a00:1450:4010:c1e::65
CNAME: google.com.
MX: aspmx.l.google.com., alt1.aspmx.l.google.com., alt2.aspmx.l.google.com., alt3.aspmx.l.google.com., alt4.aspmx.l.google.com.
NS: ns1.google.com., ns3.google.com., ns2.google.com., ns4.google.com.
TXT: google-site-verification=wD8N7i1JTNTkezJ49swvWW48f8_9xveREV4oB-0Hf5o, MS=E4A68B9AB2BB9670BCE15412F62916164C0B20BB, globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=, google-site-verification=TV9-DBe4R80X4v0M4U_bd_J9cpOJM0nikft0jAgjmsQ, docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e, docusign=1b0a6754-49b1-4db5-8540-d2c12664b289, facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95, v=spf1 include:_spf.google.com ~all, apple-domain-verification=30afIBcvSuDV2PLX
```

```
./lookres 5.255.255.55

IP: 5.255.255.55
Domains: yandex.ru.
Global unicast: yes
IPv4: 77.88.55.55
IPv6: 2a02:6b8:a::a
IP addresses: 5.255.255.50, 77.88.55.50, 5.255.255.5, 77.88.55.55, 2a02:6b8:a::a
CNAME: yandex.ru.
MX: mx.yandex.ru.
NS: ns1.yandex.ru., ns2.yandex.ru., ns9.z5h64q92x9.net.
TXT: MS=ms75457885, mailru-verification: 530c425b1458283e, facebook-domain-verification=e750ewnqm68u4f83wvp6qp7iiphkj0, google-site-verification=XyQDB5000-0rTv33yw7AX-EiuH1v5yW5PjkYeYxxPEg, 2e35680fa5ac784cf58deca180385b5eff74dfeb831c2d73830425e8a8deb7d5, _globalsign-domain-verification=lD5-OgV_QE93G8rzNaeJKvtqe9tlP5AZtyDodrldYh, have-i-been-pwned-verification=13c7b50cd0b12f85dabe796e6178fb74, 6cf645ec7ee0f9fdbd1bbfa00671b6641a91ff8b87fbdd82f086cb38cc946aa0, v=spf1 redirect=_spf.yandex.ru
```

## Build

```
go build .
```

## Name resolution

[GO net](https://pkg.go.dev/net@go1.17.1#hdr-Name_Resolution).
