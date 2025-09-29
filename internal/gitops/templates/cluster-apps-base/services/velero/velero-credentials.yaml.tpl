{/*
This file was generated from overlay template comparison
Environment-specific values are templated with Go template syntax
Original source: dev environment overlay
*/}
apiVersion: v1
data:
    cloud: ENC[AES256_GCM,data:YX+HpULnmB9OfI/obNneqFcdFHk2vX7Byi7uqUaBI5QC4L99BBv0PlRGhT+/djFcIJAVivz7GyF/K8XQYUh0mtOZSrQIxbpVp3+34UdGFQ313DmWmCcKuuAJd7OvH4RLtUnxIDXu8QNChdPzhWEhKFTKewzn9tYLxTTW2Rg+N21m+5k8zuQzhGtO6vly/TPFpUh4PkXd7fk7BuNq,iv:9wkj2jIiQeXwBdDQFR+VCOpgG4nTwSu16M4b9salNZU=,tag:ScxdZuHjtp8fpnAXalwlqA==,type:str]
kind: Secret
metadata:
    name: velero-s3-credentials
    namespace: velero
type: Opaque
sops:
    age:
        - recipient: age1ydr20ml7gdda4pkq5exe4amr426xtxnkwl2sdtt2zlxnu2jhg90qe2u4qk
          enc: |
              -----BEGIN AGE ENCRYPTED FILE-----
              YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IFgyNTUxOSBrN0dSL2FCclBnMjhDRDNN
              VUtEc09vSElSci9oUUIwaGllSFkvWHRkeWhRCnVJSjBnMEtaaEprWFQ4UVNoVW9t
              bU00WlhXemZKY29TTjVzQU1vZEU1aDQKLS0tIGZpSXpDajI1SHRLUTFXOU1QTDRm
              ME1pb25FcklKY2NjWEd5UlppSUlYTGsKGevNJMaTYIU/YaVlK362Ud4LcQ3Wn54P
              ky3RCr3txOy7jpBOLigHYej+Gz7dHivs1UoY7soHZ0QeHpd+9xzVag==
              -----END AGE ENCRYPTED FILE-----
    lastmodified: "2025-09-24T18:36:40Z"
    mac: ENC[AES256_GCM,data:GXkMiy+1rSZZZUqeaQeizfkRpN488yaBTJC0YeFJHUFYqYyiWq9I+zjh28zymcwtwHd24Bw6p4v2B4Hix5D+gIhlGmOYX09lNMtkgvOQ8+eOKVHlRgTuxkbmywMh/mEmeIM2KIUfoxuT41fg/LpQ6LDoLu9upnKOI3j2FdCXze8=,iv:NEbUDmpyLYIlltioZZ05jU6i4Qzee8CU+/meotd9u60=,tag:w9osDGDvCmxnqfhsg65GPw==,type:str]
    encrypted_regex: ^(data|stringData)$
    version: 3.10.2
