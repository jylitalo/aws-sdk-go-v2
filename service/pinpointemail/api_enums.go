// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

// The action that you want Amazon Pinpoint to take if it can't read the required
// MX record for a custom MAIL FROM domain. When you set this value to UseDefaultValue,
// Amazon Pinpoint uses amazonses.com as the MAIL FROM domain. When you set
// this value to RejectMessage, Amazon Pinpoint returns a MailFromDomainNotVerified
// error, and doesn't attempt to deliver the email.
//
// These behaviors are taken when the custom MAIL FROM domain configuration
// is in the Pending, Failed, and TemporaryFailure states.
type BehaviorOnMxFailure string

// Enum values for BehaviorOnMxFailure
const (
	BehaviorOnMxFailureUseDefaultValue BehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	BehaviorOnMxFailureRejectMessage   BehaviorOnMxFailure = "REJECT_MESSAGE"
)

func (enum BehaviorOnMxFailure) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BehaviorOnMxFailure) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The status of a predictive inbox placement test. If the status is IN_PROGRESS,
// then the predictive inbox placement test is currently running. Predictive
// inbox placement tests are usually complete within 24 hours of creating the
// test. If the status is COMPLETE, then the test is finished, and you can use
// the GetDeliverabilityTestReport operation to view the results of the test.
type DeliverabilityTestStatus string

// Enum values for DeliverabilityTestStatus
const (
	DeliverabilityTestStatusInProgress DeliverabilityTestStatus = "IN_PROGRESS"
	DeliverabilityTestStatusCompleted  DeliverabilityTestStatus = "COMPLETED"
)

func (enum DeliverabilityTestStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliverabilityTestStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The location where Amazon Pinpoint finds the value of a dimension to publish
// to Amazon CloudWatch. If you want Amazon Pinpoint to use the message tags
// that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the
// SendEmail/SendRawEmail API, choose messageTag. If you want Amazon Pinpoint
// to use your own email headers, choose emailHeader. If you want Amazon Pinpoint
// to use link tags, choose linkTags.
type DimensionValueSource string

// Enum values for DimensionValueSource
const (
	DimensionValueSourceMessageTag  DimensionValueSource = "MESSAGE_TAG"
	DimensionValueSourceEmailHeader DimensionValueSource = "EMAIL_HEADER"
	DimensionValueSourceLinkTag     DimensionValueSource = "LINK_TAG"
)

func (enum DimensionValueSource) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DimensionValueSource) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The DKIM authentication status of the identity. The status can be one of
// the following:
//
//    * PENDING – The DKIM verification process was initiated, and Amazon
//    Pinpoint is still waiting for the required CNAME records to appear in
//    the DNS configuration for the domain.
//
//    * SUCCESS – The DKIM authentication process completed successfully.
//
//    * FAILED – The DKIM authentication process failed. This can happen when
//    Amazon Pinpoint fails to find the required CNAME records in the DNS configuration
//    of the domain.
//
//    * TEMPORARY_FAILURE – A temporary issue is preventing Amazon Pinpoint
//    from determining the DKIM authentication status of the domain.
//
//    * NOT_STARTED – The DKIM verification process hasn't been initiated
//    for the domain.
type DkimStatus string

// Enum values for DkimStatus
const (
	DkimStatusPending          DkimStatus = "PENDING"
	DkimStatusSuccess          DkimStatus = "SUCCESS"
	DkimStatusFailed           DkimStatus = "FAILED"
	DkimStatusTemporaryFailure DkimStatus = "TEMPORARY_FAILURE"
	DkimStatusNotStarted       DkimStatus = "NOT_STARTED"
)

func (enum DkimStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DkimStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// An email sending event type. For example, email sends, opens, and bounces
// are all email events.
type EventType string

// Enum values for EventType
const (
	EventTypeSend             EventType = "SEND"
	EventTypeReject           EventType = "REJECT"
	EventTypeBounce           EventType = "BOUNCE"
	EventTypeComplaint        EventType = "COMPLAINT"
	EventTypeDelivery         EventType = "DELIVERY"
	EventTypeOpen             EventType = "OPEN"
	EventTypeClick            EventType = "CLICK"
	EventTypeRenderingFailure EventType = "RENDERING_FAILURE"
)

func (enum EventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The email identity type. The identity type can be one of the following:
//
//    * EMAIL_ADDRESS – The identity is an email address.
//
//    * DOMAIN – The identity is a domain.
type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeEmailAddress  IdentityType = "EMAIL_ADDRESS"
	IdentityTypeDomain        IdentityType = "DOMAIN"
	IdentityTypeManagedDomain IdentityType = "MANAGED_DOMAIN"
)

func (enum IdentityType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IdentityType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The status of the MAIL FROM domain. This status can have the following values:
//
//    * PENDING – Amazon Pinpoint hasn't started searching for the MX record
//    yet.
//
//    * SUCCESS – Amazon Pinpoint detected the required MX record for the
//    MAIL FROM domain.
//
//    * FAILED – Amazon Pinpoint can't find the required MX record, or the
//    record no longer exists.
//
//    * TEMPORARY_FAILURE – A temporary issue occurred, which prevented Amazon
//    Pinpoint from determining the status of the MAIL FROM domain.
type MailFromDomainStatus string

// Enum values for MailFromDomainStatus
const (
	MailFromDomainStatusPending          MailFromDomainStatus = "PENDING"
	MailFromDomainStatusSuccess          MailFromDomainStatus = "SUCCESS"
	MailFromDomainStatusFailed           MailFromDomainStatus = "FAILED"
	MailFromDomainStatusTemporaryFailure MailFromDomainStatus = "TEMPORARY_FAILURE"
)

func (enum MailFromDomainStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MailFromDomainStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The warmup status of a dedicated IP.
type WarmupStatus string

// Enum values for WarmupStatus
const (
	WarmupStatusInProgress WarmupStatus = "IN_PROGRESS"
	WarmupStatusDone       WarmupStatus = "DONE"
)

func (enum WarmupStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WarmupStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
