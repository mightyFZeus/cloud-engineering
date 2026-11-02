# Day 27 — Private S3 Transfers with Presigned URLs

**Verified:** November 2, 2026. User-ready drafts; external publication is not verified.

## LinkedIn

Day 27 of my AWS cloud engineering journey: I transferred a file through presigned requests while keeping its S3 bucket private.

In this guided Go lab, I kept all four S3 public-access blocks enabled and ACLs disabled. The SDK generated presigned PUT and GET requests using temporary profile credentials.

Both signed requests returned HTTP 200, and the downloaded bytes matched the fictional payload exactly. A fresh unsigned GET to the same existing object returned 403.

The useful lesson: a private bucket can still accept an authorized signed request. The URL is a bearer credential, so anyone holding it may use its permitted operation while it remains valid.

I configured a five-minute lifetime; I did not test expiration by waiting. The program kept URLs in memory and printed only safe check results.

Cleanup was verified too: object and bucket deleted, SDK BucketMissing confirmed, environment variable cleared, and CLI logged out.

Next: recover an overwritten object using S3 versioning.

https://github.com/mightyFZeus/cloud-engineering

#AWS #Golang #S3

## X / Threads

Day 27: private S3 transfers with Go presigned PUT/GET. Signed requests returned 200 with exact bytes; unsigned GET returned 403. URLs stayed private. Deleted the object/bucket and verified cleanup.
https://github.com/mightyFZeus/cloud-engineering
#AWS #Golang
