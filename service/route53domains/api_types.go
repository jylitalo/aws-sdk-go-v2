// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Information for one billing record.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/BillingRecord
type BillingRecord struct {
	_ struct{} `type:"structure"`

	// The date that the operation was billed, in Unix format.
	BillDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The name of the domain that the billing record applies to. If the domain
	// name contains characters other than a-z, 0-9, and - (hyphen), such as an
	// internationalized domain name, then this value is in Punycode. For more information,
	// see DNS Domain Name Format (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html)
	// in the Amazon Route 53 Developer Guidezzz.
	DomainName *string `type:"string"`

	// The ID of the invoice that is associated with the billing record.
	InvoiceId *string `type:"string"`

	// The operation that you were charged for.
	Operation OperationType `type:"string" enum:"true"`

	// The price that you were charged for the operation, in US dollars.
	//
	// Example value: 12.0
	Price *float64 `type:"double"`
}

// String returns the string representation
func (s BillingRecord) String() string {
	return awsutil.Prettify(s)
}

// ContactDetail includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ContactDetail
type ContactDetail struct {
	_ struct{} `type:"structure"`

	// First line of the contact's address.
	AddressLine1 *string `type:"string"`

	// Second line of contact's address, if any.
	AddressLine2 *string `type:"string"`

	// The city of the contact's address.
	City *string `type:"string"`

	// Indicates whether the contact is a person, company, association, or public
	// organization. If you choose an option other than PERSON, you must enter an
	// organization name, and you can't enable privacy protection for the contact.
	ContactType ContactType `type:"string" enum:"true"`

	// Code for the country of the contact's address.
	CountryCode CountryCode `type:"string" enum:"true"`

	// Email address of the contact.
	Email *string `type:"string"`

	// A list of name-value pairs for parameters required by certain top-level domains.
	ExtraParams []ExtraParam `type:"list"`

	// Fax number of the contact.
	//
	// Constraints: Phone number must be specified in the format "+[country dialing
	// code].[number including any area code]". For example, a US phone number might
	// appear as "+1.1234567890".
	Fax *string `type:"string"`

	// First name of contact.
	FirstName *string `type:"string"`

	// Last name of contact.
	LastName *string `type:"string"`

	// Name of the organization for contact types other than PERSON.
	OrganizationName *string `type:"string"`

	// The phone number of the contact.
	//
	// Constraints: Phone number must be specified in the format "+[country dialing
	// code].[number including any area code>]". For example, a US phone number
	// might appear as "+1.1234567890".
	PhoneNumber *string `type:"string"`

	// The state or province of the contact's city.
	State *string `type:"string"`

	// The zip or postal code of the contact's address.
	ZipCode *string `type:"string"`
}

// String returns the string representation
func (s ContactDetail) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ContactDetail) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ContactDetail"}
	if s.ExtraParams != nil {
		for i, v := range s.ExtraParams {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ExtraParams", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about one suggested domain name.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DomainSuggestion
type DomainSuggestion struct {
	_ struct{} `type:"structure"`

	// Whether the domain name is available for registering.
	//
	// You can register only the domains that are designated as AVAILABLE.
	//
	// Valid values:
	//
	// AVAILABLE
	//
	// The domain name is available.
	//
	// AVAILABLE_RESERVED
	//
	// The domain name is reserved under specific conditions.
	//
	// AVAILABLE_PREORDER
	//
	// The domain name is available and can be preordered.
	//
	// DONT_KNOW
	//
	// The TLD registry didn't reply with a definitive answer about whether the
	// domain name is available. Amazon Route 53 can return this response for a
	// variety of reasons, for example, the registry is performing maintenance.
	// Try again later.
	//
	// PENDING
	//
	// The TLD registry didn't return a response in the expected amount of time.
	// When the response is delayed, it usually takes just a few extra seconds.
	// You can resubmit the request immediately.
	//
	// RESERVED
	//
	// The domain name has been reserved for another person or organization.
	//
	// UNAVAILABLE
	//
	// The domain name is not available.
	//
	// UNAVAILABLE_PREMIUM
	//
	// The domain name is not available.
	//
	// UNAVAILABLE_RESTRICTED
	//
	// The domain name is forbidden.
	Availability *string `type:"string"`

	// A suggested domain name.
	DomainName *string `type:"string"`
}

// String returns the string representation
func (s DomainSuggestion) String() string {
	return awsutil.Prettify(s)
}

// Summary information about one domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DomainSummary
type DomainSummary struct {
	_ struct{} `type:"structure"`

	// Indicates whether the domain is automatically renewed upon expiration.
	AutoRenew *bool `type:"boolean"`

	// The name of the domain that the summary information applies to.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// Expiration date of the domain in Coordinated Universal Time (UTC).
	Expiry *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Indicates whether a domain is locked from unauthorized transfer to another
	// party.
	TransferLock *bool `type:"boolean"`
}

// String returns the string representation
func (s DomainSummary) String() string {
	return awsutil.Prettify(s)
}

// A complex type that contains information about whether the specified domain
// can be transferred to Amazon Route 53.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DomainTransferability
type DomainTransferability struct {
	_ struct{} `type:"structure"`

	// Whether the domain name can be transferred to Amazon Route 53.
	//
	// You can transfer only domains that have a value of TRANSFERABLE for Transferable.
	//
	// Valid values:
	//
	// TRANSFERABLE
	//
	// The domain name can be transferred to Amazon Route 53.
	//
	// UNTRANSFERRABLE
	//
	// The domain name can't be transferred to Amazon Route 53.
	//
	// DONT_KNOW
	//
	// Reserved for future use.
	Transferable Transferable `type:"string" enum:"true"`
}

// String returns the string representation
func (s DomainTransferability) String() string {
	return awsutil.Prettify(s)
}

// ExtraParam includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ExtraParam
type ExtraParam struct {
	_ struct{} `type:"structure"`

	// Name of the additional parameter required by the top-level domain. Here are
	// the top-level domains that require additional parameters and which parameters
	// they require:
	//
	//    * .com.au and .net.au: AU_ID_NUMBER and AU_ID_TYPE
	//
	//    * .ca: BRAND_NUMBER, CA_LEGAL_TYPE, and CA_BUSINESS_ENTITY_TYPE
	//
	//    * .es: ES_IDENTIFICATION, ES_IDENTIFICATION_TYPE, and ES_LEGAL_FORM
	//
	//    * .fi: BIRTH_DATE_IN_YYYY_MM_DD, FI_BUSINESS_NUMBER, FI_ID_NUMBER, FI_NATIONALITY,
	//    and FI_ORGANIZATION_TYPE
	//
	//    * .fr: BRAND_NUMBER, BIRTH_DEPARTMENT, BIRTH_DATE_IN_YYYY_MM_DD, BIRTH_COUNTRY,
	//    and BIRTH_CITY
	//
	//    * .it: BIRTH_COUNTRY, IT_PIN, and IT_REGISTRANT_ENTITY_TYPE
	//
	//    * .ru: BIRTH_DATE_IN_YYYY_MM_DD and RU_PASSPORT_DATA
	//
	//    * .se: BIRTH_COUNTRY and SE_ID_NUMBER
	//
	//    * .sg: SG_ID_NUMBER
	//
	//    * .co.uk, .me.uk, and .org.uk: UK_CONTACT_TYPE and UK_COMPANY_NUMBER
	//
	// In addition, many TLDs require VAT_NUMBER.
	//
	// Name is a required field
	Name ExtraParamName `type:"string" required:"true" enum:"true"`

	// Values corresponding to the additional parameter names required by some top-level
	// domains.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ExtraParam) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExtraParam) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExtraParam"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Nameserver includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/Nameserver
type Nameserver struct {
	_ struct{} `type:"structure"`

	// Glue IP address of a name server entry. Glue IP addresses are required only
	// when the name of the name server is a subdomain of the domain. For example,
	// if your domain is example.com and the name server for the domain is ns.example.com,
	// you need to specify the IP address for ns.example.com.
	//
	// Constraints: The list can contain only one IPv4 and one IPv6 address.
	GlueIps []string `type:"list"`

	// The fully qualified host name of the name server.
	//
	// Constraint: Maximum 255 characters
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Nameserver) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Nameserver) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Nameserver"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// OperationSummary includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/OperationSummary
type OperationSummary struct {
	_ struct{} `type:"structure"`

	// Identifier returned to track the requested action.
	//
	// OperationId is a required field
	OperationId *string `type:"string" required:"true"`

	// The current status of the requested operation in the system.
	//
	// Status is a required field
	Status OperationStatus `type:"string" required:"true" enum:"true"`

	// The date when the request was submitted.
	//
	// SubmittedDate is a required field
	SubmittedDate *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// Type of the action requested.
	//
	// Type is a required field
	Type OperationType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s OperationSummary) String() string {
	return awsutil.Prettify(s)
}

// Each tag includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The key (name) of a tag.
	//
	// Valid values: A-Z, a-z, 0-9, space, ".:/=+\-@"
	//
	// Constraints: Each key can be 1-128 characters long.
	Key *string `type:"string"`

	// The value of a tag.
	//
	// Valid values: A-Z, a-z, 0-9, space, ".:/=+\-@"
	//
	// Constraints: Each value can be 0-256 characters long.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}
