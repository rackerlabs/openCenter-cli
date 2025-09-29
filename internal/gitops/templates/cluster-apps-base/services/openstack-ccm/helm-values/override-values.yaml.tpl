{/*
This file was generated from overlay template comparison
Environment-specific values are templated with Go template syntax
Original source: dev environment overlay
*/}
cluster:
    name: ENC[AES256_GCM,data:kvbU,iv:HQ8kmQAd1SHildJZ4WUowcTZ+ZuvBNNFSlCTTaZ5rw0=,tag:E7LW+pQ8m8qNxzQsstSb9g==,type:str]
cloudConfig:
    global:
        auth-url: ENC[AES256_GCM,data:SmEEwcCde0bocKASf4FKI39Gti6Vc+1s4vQn6rTi90K/TML2E9UYK8RfsSiGoJw=,iv:IK+YYiGC2I7E/aoOMwWblsGOZJmLMvR4ZyuKkmJJxwE=,tag:MQFwdPrvYOgiVj1yVMi7Sw==,type:str]
        application-credential-id: ENC[AES256_GCM,data:XS3kOqpP8n0iUV0lM79fehRDO1VDZv/Quk3GNSu6yKk=,iv:wEhyHC/7TXLntVGDgkniIfH3t04WiiOs/lYm5ec7Fl8=,tag:Jp39Th+IFzZjRRxHL0V+PQ==,type:str]
        application-credential-secret: ENC[AES256_GCM,data:eG+gem3yz6405mKk0O9atYJiCQrQPgwMhWzqyII4Yo7VBHpmdc+EzdDG7kh1fJb3I4845/5xd5MbesCikHnVt7W8PldXqHgQ4UmXd+ClfwhC/pNawjw=,iv:PLx5XWx3Kd1OCvJa3gfRphwnT+wosxRqdrfggo0c4MM=,tag:je3Koz5YSy2WBMwtyq97dA==,type:str]
        domain-name: ENC[AES256_GCM,data:P4H5GZKVOPJR6djX2m8etU1BOI7cYg==,iv:wyDy/cfr2bCZgzD3kYn5R04ivDOAtDWUlMuQC1ZDxZw=,tag:+sCLzn9gk/HIXCQFzqe0Cg==,type:str]
        region: ENC[AES256_GCM,data:Ux/uoA==,iv:BLibGqZeAOtHFgKdmsf89vg94FEnSJUMng1lTXuiAeQ=,tag:5++28+lvXFmDPGJUnLTsSA==,type:str]
        tenant-name: ENC[AES256_GCM,data:n10Dj0QdIYq4fEOW6noK23zfqAZdWjfhPqAxN0hHlGqql/or,iv:XIMxM6cCf5w5NgLkCFcRiYyJ7+m7JhanuEH+uexWjtE=,tag:SQ7lVfn4wZ+4GQNo9WHQkg==,type:str]
        tls-insecure: ENC[AES256_GCM,data:V3NRZIY=,iv:WTMJohWo90c26m0pzm8nyXp1PbLAGlHrhS8surrC7nc=,tag:GS2wh5IEygzAz7mJ7NPXQg==,type:bool]
    loadBalancer:
        floating-network-id: ENC[AES256_GCM,data:hgXCCQSx9kOqLcv3Zjv+DnMAjcfv7PczDw/+O58ULAxcBtgp,iv:URo5U1TJmIpubm+yKG9b6zQQ0Zbl7wpKNy+PM2z0mQg=,tag:6WgNxtIv+Om0I5V1erbXsg==,type:str]
        subnet-id: ENC[AES256_GCM,data:lUuOoZlT66DoiFYy2JLchWobZVXCUZYly6mqt3YOUd+baXr4,iv:i0g4NbJyu3iLBDEmN03OvjF150BYcV3Z7B1YJ5SFCqo=,tag:TcWj6V292Ec2uLXwlSz9mg==,type:str]
sops:
    age:
        - recipient: age1ydr20ml7gdda4pkq5exe4amr426xtxnkwl2sdtt2zlxnu2jhg90qe2u4qk
          enc: |
              -----BEGIN AGE ENCRYPTED FILE-----
              YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IFgyNTUxOSBhNkVMM3Yvc2wxS1JqeUZh
              R3BzRERidDl1RnU0QXd3eG0walE2a0x4NEg0CmQ3MkVpTFVVa2dWdFFKb3dNc04v
              K1FIUjR2SDEveWVxRUQyOWNUM3R2MU0KLS0tIDZTSExueDZxU2FHVm4wb3hBamJa
              bTVpZjNlbytTaEpFemsvNHdMS2VkS3cKuQ7r5AwVkKtYvjOzB22uj7YV2ukZWA0R
              WTZYnu2GA2Ir/cu0esO5vbecRYIcy3kr4rLG8WynqkzqHHLePoVPcQ==
              -----END AGE ENCRYPTED FILE-----
    lastmodified: "2025-09-24T18:01:55Z"
    mac: ENC[AES256_GCM,data:f/kiqHG+xEAg+48bbcgXBMx6liNnhNhcq+ks7E4LS39BEYO4dS5Z/UhamzSIRHufzHQEjyqC/F3w9nQ7upoWu1RdBD3Du6D+G7hkEgHf9mjjOmq14ndAdq5M+JgT6S94Hfat5yliwiduQUYWuHiyR5MuAHO4X+FYbG3delTN+60=,iv:LnBhwULFJ5jI0uwmOMlwD2UZ/Gktp7y1MLfdDzjmDCY=,tag:ro24UeBqyVXTMK+77eA7WQ==,type:str]
    unencrypted_suffix: _unencrypted
    version: 3.10.2
