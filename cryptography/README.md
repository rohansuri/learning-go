Achieving CIA with Go's `crypto` packages.

Requirements for a secure system

* Confidentiality  

  So that an adversary cannot read your message.

* Integrity  

  Ensuring data is safeguarded from modification.

* Authenticity  

  Making sure the message really came from who you think it has.

Go has the `crypto` package to offer:  
* encryption  
* hashing  
* signatures

Below we list only the recommended ones to use.  

(see GopherCon 2016 talk - Crypto for Go developers)

### Encryption

#### Symmetric key

* The same shared secret key is used to encrypt and decrypt the message.

* Diffie-Hellman could be used for key agreement over an insecure channel.

* Asymmetric key encryption could also be used to encrypt and share the secret key.

##### AES

* AES is a block cipher symmetric key algorithm.

* It only encrypts 16 bytes at a time which is called the block size.

* To encrypt data larger than the block size we use one of the block cipher's mode of operation which repeatedly applies the cipher's single block operation.

* GCM is the recommended mode to use for AES.

  This mode provides confidentialty as well as integrity and falls under the section of AEAD modes of operation.

  It requires a nonce (number used only once) per key in it's operation.  

  The suggestion is to use a large random nonce.

* Many modern CPUs have built in instructions for AES hence improving the speed and guarding against side-channel attacks.

* Use 256 bit key `crypto/aes` GCM with 96-bit random nonces.

#### Asymmetric key

* Asymmetric key encryption is not primarily used for data encryption since it is slow and works well only on small data.

* Therefore, it is rather used to encrypt and communicate the shared secret key to be used for symmetric key encryption or used in signatures.


### Hashing

* `crypto/sha256 crypto/sha512`

* Hashes used directly could be susceptible to attacks like rainbow tables, length extension, etc. and therefore it is recommended to use HMACs.

* For password hashing, use `bcrypt`.  
  
  It is an adaptive hash function meant to make the hash computation slower thereby resisting brute force attacks.

  14 is a good work factor.

  It also uses a salt to protect against rainbow tables.

### Signatures

* Digital signatures serve the purpose of integrity, authenticity and sometimes non-repudiation.

  Non-repudiation means the sender cannot deny being the one sending the message.

* Relies on public-key cryptography (asymmetric key cryptography)

* A signature is a hash digest of the message encrypted by the private key of the sender.  
  
  On the receiver side the verification involves producing a hash digest of the message and comparing it with the decrypted signature using sender's public key.

  Remember signatures don't provide you confidentiality.

* Elliptic curves based `crypto/ecdsa` P-256 with SHA 256 message digests is faster than RSA and requires smaller keys.

### Attacks

TODO 
* Side channel  
* Length extension  
* Rainbow tables
 