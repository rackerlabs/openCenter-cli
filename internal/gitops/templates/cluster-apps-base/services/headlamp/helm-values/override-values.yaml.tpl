{/*
This file was generated from overlay template comparison
Environment-specific values are templated with Go template syntax
Original source: dev environment overlay
*/}
config:
  oidc:
    externalSecret:
      enabled: ENC[AES256_GCM,data:3+ceAlg=,iv:l5SpJ/iDynCgR5joykOF6ezrHSFArLhrJyWD8xhxttI=,tag:oTTwBCCvlKdJz6XayMuL+A==,type:bool]
    secret:
      create: ENC[AES256_GCM,data:Fx5+7Q==,iv:vNSwmlLTEz9BZ6n3+/HU/Vmk2SxRAABv7c3eo0OS29s=,tag:RWf7joDB/68UO3oVPhyN0g==,type:bool]
    clientID: ENC[AES256_GCM,data:CJb5Lke09tjNQw==,iv:jShX1UWyVLzfg2kQdO9OvZrcg+2DKJtErMe0jhYn04c=,tag:eCN170Usj+r/5LKPcmMHqw==,type:str]
    clientSecret: ENC[AES256_GCM,data:xScFJ0cNEGLeC0PpzWvyF2woFaOl6XuxkUEK1l2R1/A=,iv:Sgnc4L4KSLTk+zpxVB68bTg8KLSHkzz80+SDL31FQVw=,tag:aPJaRvAz2ucY1p3/0tjTwg==,type:str]
    issuerURL: ENC[AES256_GCM,data:ddDBtDifWxyiXc1ehZbKcqzTUz8BQYV9QBue143Qhso2C9c1N5MFW/0Hk3jOA6E2F7baNwgP1Igz0GWTN0XvHAoTwg==,iv:LpR61+X6E2A6VB5X9Qpt4pb+FvaH1J/1VyhUVrVF3AQ=,tag:vLYWof7Voe4efq0aLiFXuA==,type:str]
    scopes: ENC[AES256_GCM,data:Zy42nOtWRglvGiqdXPRGv376gFFodCggXfYr,iv:xKz3ok+wdrjQHfZGmOLmpyqVUQxRjePLOt/35lWBIB4=,tag:xn74XCY3YHxTOCXCMSxOwA==,type:str]
    callbackURL: ENC[AES256_GCM,data:uH+F3KrIpEVJGHul4OHp5QW6D5iUhXg4Ic7ZCwzXn7uRNeoW9CkE+dYgY+Mo+EFqyYUeZnMQE138jD7yqXL6+cS7mg==,iv:xi1lDCI17DOrsu+W16qdtdHXkOhZ2i9LMd05CXXjknI=,tag:0+PRTb65c1Sp1d9/mm40OA==,type:str]
sops:
  age:
    - recipient: age1ydr20ml7gdda4pkq5exe4amr426xtxnkwl2sdtt2zlxnu2jhg90qe2u4qk
      enc: |
        -----BEGIN AGE ENCRYPTED FILE-----
        YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IFgyNTUxOSBPV0QrWjhPSzV4OEJHc2po
        VTN1ZEJaNnF2K2xGOHRKcENIZ3VSR3ZDU2hBClNPTVMxcFpXRmMrKzFwcmNybFc0
        VGIxQllzYjZ1S2JXU1Z6TXdQcW13M0EKLS0tIGN4TU5JdURQNlUrYmNFRCt1SnRl
        dGJJcUg2L05LYy9maVdwT1NZVFg5aWMKEVk1HvQ4aatdebo4fIj7146v3bdON3Kd
        BIRkv4E2B4id+DupkIFF15SnBvPOKLsdaDf539Qnp9JSFVraFDnsEg==
        -----END AGE ENCRYPTED FILE-----
  lastmodified: "2025-09-24T19:46:50Z"
  mac: ENC[AES256_GCM,data:4w1tlnt/gH9CTeCYlVdD0X7BEfmMf2IFEK3yuR3USejAiW4P1S0XOgs7Jv0epe5YruFQgDUAse/JK7mgaITGpyZzCpUVwBVQiTm2nmo8+Zsno0gXR1AHQQgdPCII1nO+cwOyghHAMc4RE7Jh66u+qgPtNOLFDTLF8CRwHKWNTRc=,iv:Iu5npj7x517BHGRPdGCmLXhjAwrZ88AZTWTR3Fdx6ho=,tag:I9TxVWVVFO1lXentGzL1eA==,type:str]
  unencrypted_suffix: _unencrypted
  version: 3.10.2
