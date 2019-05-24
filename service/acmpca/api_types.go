// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains information about the certificate subject. The certificate can be
// one issued by your private certificate authority (CA) or it can be your private
// CA certificate. The Subject field in the certificate identifies the entity
// that owns or controls the public key in the certificate. The entity can be
// a user, computer, device, or service. The Subject must contain an X.500 distinguished
// name (DN). A DN is a sequence of relative distinguished names (RDNs). The
// RDNs are separated by commas in the certificate. The DN must be unique for
// each entity, but your private CA can issue more than one certificate with
// the same DN to the same entity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ASN1Subject
type ASN1Subject struct {
	_ struct{} `type:"structure"`

	// Fully qualified domain name (FQDN) associated with the certificate subject.
	CommonName *string `type:"string"`

	// Two-digit code that specifies the country in which the certificate subject
	// located.
	Country *string `type:"string"`

	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string `type:"string"`

	// Typically a qualifier appended to the name of an individual. Examples include
	// Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string `type:"string"`

	// First name.
	GivenName *string `type:"string"`

	// Concatenation that typically contains the first letter of the GivenName,
	// the first letter of the middle name if one exists, and the first letter of
	// the SurName.
	Initials *string `type:"string"`

	// The locality (such as a city or town) in which the certificate subject is
	// located.
	Locality *string `type:"string"`

	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string `type:"string"`

	// A subdivision or unit of the organization (such as sales or finance) with
	// which the certificate subject is affiliated.
	OrganizationalUnit *string `type:"string"`

	// Typically a shortened version of a longer GivenName. For example, Jonathan
	// is often shortened to John. Elizabeth is often shortened to Beth, Liz, or
	// Eliza.
	Pseudonym *string `type:"string"`

	// The certificate serial number.
	SerialNumber *string `type:"string"`

	// State in which the subject of the certificate is located.
	State *string `type:"string"`

	// Family name. In the US and the UK, for example, the surname of an individual
	// is ordered last. In Asian cultures the surname is typically ordered first.
	Surname *string `type:"string"`

	// A title such as Mr. or Ms., which is pre-pended to the name to refer formally
	// to the certificate subject.
	Title *string `type:"string"`
}

// String returns the string representation
func (s ASN1Subject) String() string {
	return awsutil.Prettify(s)
}

// Contains information about your private certificate authority (CA). Your
// private CA can issue and revoke X.509 digital certificates. Digital certificates
// verify that the entity named in the certificate Subject field owns or controls
// the public key contained in the Subject Public Key Info field. Call the CreateCertificateAuthority
// operation to create your private CA. You must then call the GetCertificateAuthorityCertificate
// operation to retrieve a private CA certificate signing request (CSR). Take
// the CSR to your on-premises CA and sign it with the root CA certificate or
// a subordinate certificate. Call the ImportCertificateAuthorityCertificate
// operation to import the signed certificate into AWS Certificate Manager (ACM).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CertificateAuthority
type CertificateAuthority struct {
	_ struct{} `type:"structure"`

	// Amazon Resource Name (ARN) for your private certificate authority (CA). The
	// format is 12345678-1234-1234-1234-123456789012 .
	Arn *string `min:"5" type:"string"`

	// Your private CA configuration.
	CertificateAuthorityConfiguration *CertificateAuthorityConfiguration `type:"structure"`

	// Date and time at which your private CA was created.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Reason the request to create your private CA failed.
	FailureReason FailureReason `type:"string" enum:"true"`

	// Date and time at which your private CA was last updated.
	LastStateChangeAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Date and time after which your private CA certificate is not valid.
	NotAfter *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Date and time before which your private CA certificate is not valid.
	NotBefore *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The period during which a deleted CA can be restored. For more information,
	// see the PermanentDeletionTimeInDays parameter of the DeleteCertificateAuthorityRequest
	// operation.
	RestorableUntil *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Information about the certificate revocation list (CRL) created and maintained
	// by your private CA.
	RevocationConfiguration *RevocationConfiguration `type:"structure"`

	// Serial number of your private CA.
	Serial *string `type:"string"`

	// Status of your private CA.
	Status CertificateAuthorityStatus `type:"string" enum:"true"`

	// Type of your private CA.
	Type CertificateAuthorityType `type:"string" enum:"true"`
}

// String returns the string representation
func (s CertificateAuthority) String() string {
	return awsutil.Prettify(s)
}

// Contains configuration information for your private certificate authority
// (CA). This includes information about the class of public key algorithm and
// the key pair that your private CA creates when it issues a certificate. It
// also includes the signature algorithm that it uses when issuing certificates,
// and its X.500 distinguished name. You must specify this information when
// you call the CreateCertificateAuthority operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CertificateAuthorityConfiguration
type CertificateAuthorityConfiguration struct {
	_ struct{} `type:"structure"`

	// Type of the public key algorithm and size, in bits, of the key pair that
	// your key pair creates when it issues a certificate.
	//
	// KeyAlgorithm is a required field
	KeyAlgorithm KeyAlgorithm `type:"string" required:"true" enum:"true"`

	// Name of the algorithm your private CA uses to sign certificate requests.
	//
	// SigningAlgorithm is a required field
	SigningAlgorithm SigningAlgorithm `type:"string" required:"true" enum:"true"`

	// Structure that contains X.500 distinguished name information for your private
	// CA.
	//
	// Subject is a required field
	Subject *ASN1Subject `type:"structure" required:"true"`
}

// String returns the string representation
func (s CertificateAuthorityConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CertificateAuthorityConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CertificateAuthorityConfiguration"}
	if len(s.KeyAlgorithm) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("KeyAlgorithm"))
	}
	if len(s.SigningAlgorithm) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SigningAlgorithm"))
	}

	if s.Subject == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subject"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains configuration information for a certificate revocation list (CRL).
// Your private certificate authority (CA) creates base CRLs. Delta CRLs are
// not supported. You can enable CRLs for your new or an existing private CA
// by setting the Enabled parameter to true. Your private CA writes CRLs to
// an S3 bucket that you specify in the S3BucketName parameter. You can hide
// the name of your bucket by specifying a value for the CustomCname parameter.
// Your private CA copies the CNAME or the S3 bucket name to the CRL Distribution
// Points extension of each certificate it issues. Your S3 bucket policy must
// give write permission to ACM PCA.
//
// Your private CA uses the value in the ExpirationInDays parameter to calculate
// the nextUpdate field in the CRL. The CRL is refreshed at 1/2 the age of next
// update or when a certificate is revoked. When a certificate is revoked, it
// is recorded in the next CRL that is generated and in the next audit report.
// Only time valid certificates are listed in the CRL. Expired certificates
// are not included.
//
// CRLs contain the following fields:
//
//    * Version: The current version number defined in RFC 5280 is V2. The integer
//    value is 0x1.
//
//    * Signature Algorithm: The name of the algorithm used to sign the CRL.
//
//    * Issuer: The X.500 distinguished name of your private CA that issued
//    the CRL.
//
//    * Last Update: The issue date and time of this CRL.
//
//    * Next Update: The day and time by which the next CRL will be issued.
//
//    * Revoked Certificates: List of revoked certificates. Each list item contains
//    the following information. Serial Number: The serial number, in hexadecimal
//    format, of the revoked certificate. Revocation Date: Date and time the
//    certificate was revoked. CRL Entry Extensions: Optional extensions for
//    the CRL entry. X509v3 CRL Reason Code: Reason the certificate was revoked.
//
//    * CRL Extensions: Optional extensions for the CRL. X509v3 Authority Key
//    Identifier: Identifies the public key associated with the private key
//    used to sign the certificate. X509v3 CRL Number:: Decimal sequence number
//    for the CRL.
//
//    * Signature Algorithm: Algorithm used by your private CA to sign the CRL.
//
//    * Signature Value: Signature computed over the CRL.
//
// Certificate revocation lists created by ACM PCA are DER-encoded. You can
// use the following OpenSSL command to list a CRL.
//
// openssl crl -inform DER -text -in crl_path -noout
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CrlConfiguration
type CrlConfiguration struct {
	_ struct{} `type:"structure"`

	// Name inserted into the certificate CRL Distribution Points extension that
	// enables the use of an alias for the CRL distribution point. Use this value
	// if you don't want the name of your S3 bucket to be public.
	CustomCname *string `type:"string"`

	// Boolean value that specifies whether certificate revocation lists (CRLs)
	// are enabled. You can use this value to enable certificate revocation for
	// a new CA when you call the CreateCertificateAuthority operation or for an
	// existing CA when you call the UpdateCertificateAuthority operation.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`

	// Number of days until a certificate expires.
	ExpirationInDays *int64 `min:"1" type:"integer"`

	// Name of the S3 bucket that contains the CRL. If you do not provide a value
	// for the CustomCname argument, the name of your S3 bucket is placed into the
	// CRL Distribution Points extension of the issued certificate. You can change
	// the name of your bucket by calling the UpdateCertificateAuthority operation.
	// You must specify a bucket policy that allows ACM PCA to write the CRL to
	// your bucket.
	S3BucketName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s CrlConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CrlConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CrlConfiguration"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}
	if s.ExpirationInDays != nil && *s.ExpirationInDays < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ExpirationInDays", 1))
	}
	if s.S3BucketName != nil && len(*s.S3BucketName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("S3BucketName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Permissions designate which private CA operations can be performed by an
// AWS service or entity. In order for ACM to automatically renew private certificates,
// you must give the ACM service principal all available permissions (IssueCertificate,
// GetCertificate, and ListPermissions). Permissions can be assigned with the
// CreatePermission operation, removed with the DeletePermission operation,
// and listed with the ListPermissions operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/Permission
type Permission struct {
	_ struct{} `type:"structure"`

	// The private CA operations that can be performed by the designated AWS service.
	Actions []ActionType `min:"1" type:"list"`

	// The Amazon Resource Number (ARN) of the private CA from which the permission
	// was issued.
	CertificateAuthorityArn *string `min:"5" type:"string"`

	// The time at which the permission was created.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The name of the policy that is associated with the permission.
	Policy *string `type:"string"`

	// The AWS service or entity that holds the permission. At this time, the only
	// valid principal is acm.amazonaws.com.
	Principal *string `type:"string"`

	// The ID of the account that assigned the permission.
	SourceAccount *string `type:"string"`
}

// String returns the string representation
func (s Permission) String() string {
	return awsutil.Prettify(s)
}

// Certificate revocation information used by the CreateCertificateAuthority
// and UpdateCertificateAuthority operations. Your private certificate authority
// (CA) can create and maintain a certificate revocation list (CRL). A CRL contains
// information about certificates revoked by your CA. For more information,
// see RevokeCertificate.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/RevocationConfiguration
type RevocationConfiguration struct {
	_ struct{} `type:"structure"`

	// Configuration of the certificate revocation list (CRL), if any, maintained
	// by your private CA.
	CrlConfiguration *CrlConfiguration `type:"structure"`
}

// String returns the string representation
func (s RevocationConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevocationConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevocationConfiguration"}
	if s.CrlConfiguration != nil {
		if err := s.CrlConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CrlConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Tags are labels that you can use to identify and organize your private CAs.
// Each tag consists of a key and an optional value. You can associate up to
// 50 tags with a private CA. To add one or more tags to a private CA, call
// the TagCertificateAuthority operation. To remove a tag, call the UntagCertificateAuthority
// operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// Key (name) of the tag.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// Value of the tag.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Length of time for which the certificate issued by your private certificate
// authority (CA), or by the private CA itself, is valid in days, months, or
// years. You can issue a certificate by calling the IssueCertificate operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/Validity
type Validity struct {
	_ struct{} `type:"structure"`

	// Specifies whether the Value parameter represents days, months, or years.
	//
	// Type is a required field
	Type ValidityPeriodType `type:"string" required:"true" enum:"true"`

	// Time period.
	//
	// Value is a required field
	Value *int64 `min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s Validity) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Validity) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Validity"}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}
	if s.Value != nil && *s.Value < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
