


### Auto_Cert
* fixed bug that was referencing the wrong s3 bucket for
sgt config
  * This means that prior version will possibly have problems?  How does this upgrade work?
  
* adds LE autocert capabilities

* adds support for credentials from env vars and sts credentials 