# goregextcp
This program converts Exchange Active directory  compliant values in postfix usable values.
eg: 
input smtp:someuser@domain.tld,sip:someuser@domain.tld
output someuser@domain.tld

The supported chat in postmap is 4096 chars cannot give more than that per call (so be wise on the number of aliases)

usage:
```
./concTCP 8009
postmap  -q sasluser 'pipemap:{ldap:/etc/postfix/ad-ldap.cf,tcp:127.0.0.1:8009}'
```



