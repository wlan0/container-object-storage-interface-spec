# Object Storage Interface (OSI)

Object Storage Interface (OSI) is an industry standard that defines a single API intended to work across a number of Object Storage (OS) systems for provisioning and accessing buckets.

Authors:

- [Sidhartha Mani](https://github.com/wlan0), MinIO [<img alt="email" src="https://raw.githubusercontent.com/google/material-design-icons/master/src/content/mail/materialiconsround/24px.svg" height=16 width=16/>](mailto:sid@min.io)
- [Jeff Vance](https://github.com/jeffvance), RedHat [<img alt="email" src="https://raw.githubusercontent.com/google/material-design-icons/master/src/content/mail/materialiconsround/24px.svg" height=16 width=16/>](mailto:jvance@redhat.com)

## Notational Conventions

The keywords "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" are to be interpreted as described in [RFC 2119](http://tools.ietf.org/html/rfc2119) (Bradner, S., "Key words for use in RFCs to Indicate Requirement Levels", BCP 14, RFC 2119, March 1997).

The key words "unspecified", "undefined", and "implementation-defined" are to be interpreted as described in the [rationale for the C99 standard](http://www.open-std.org/jtc1/sc22/wg14/www/C99RationaleV5.10.pdf#page=18).

An implementation is not compliant if it fails to satisfy one or more of the MUST, REQUIRED, or SHALL requirements for the protocols it implements.
An implementation is compliant if it satisfies all the MUST, REQUIRED, and SHALL requirements for the protocols it implements.

## Terminology

| Term                    | Definition                                       |
|-------------------------|--------------------------------------------------|
| Bucket                  | Unit of provisioned Object Storage with its own endpoint |
| Object Storage Protocol | API specification for the data path of a particular OS system. eg. s3v4, gs |
