// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/awsutil"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
)

const opDetectEntities = "DetectEntities"

// DetectEntitiesRequest generates a "aws/request.Request" representing the
// client's request for the DetectEntities operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DetectEntities for more information on using the DetectEntities
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DetectEntitiesRequest method.
//    req, resp := client.DetectEntitiesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectEntities
func (c *ComprehendMedical) DetectEntitiesRequest(input *DetectEntitiesInput) (req *request.Request, output *DetectEntitiesOutput) {
	op := &request.Operation{
		Name:       opDetectEntities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectEntitiesInput{}
	}

	output = &DetectEntitiesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DetectEntities API operation for AWS Comprehend Medical.
//
// Inspects the clinical text for a variety of medical entities and returns
// specific information about them such as entity category, location, and confidence
// score on that information .
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Comprehend Medical's
// API operation DetectEntities for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalServerException "InternalServerException"
//   An internal server error occurred. Retry your request.
//
//   * ErrCodeServiceUnavailableException "ServiceUnavailableException"
//   The Comprehend Medical service is temporarily unavailable. Please wait and
//   then retry your request.
//
//   * ErrCodeTooManyRequestsException "TooManyRequestsException"
//   You have made too many requests within a short period of time. Wait for a
//   short time and then try your request again. Contact customer support for
//   more information about a service limit increase.
//
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request that you made is invalid. Check your request to determine why
//   it's invalid and then retry the request.
//
//   * ErrCodeInvalidEncodingException "InvalidEncodingException"
//   The input text was not in valid UTF-8 character encoding. Check your text
//   then retry your request.
//
//   * ErrCodeTextSizeLimitExceededException "TextSizeLimitExceededException"
//   The size of the text you submitted exceeds the size limit. Reduce the size
//   of the text or use a smaller document and then retry your request.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectEntities
func (c *ComprehendMedical) DetectEntities(input *DetectEntitiesInput) (*DetectEntitiesOutput, error) {
	req, out := c.DetectEntitiesRequest(input)
	return out, req.Send()
}

// DetectEntitiesWithContext is the same as DetectEntities with the addition of
// the ability to pass a context and additional request options.
//
// See DetectEntities for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ComprehendMedical) DetectEntitiesWithContext(ctx aws.Context, input *DetectEntitiesInput, opts ...request.Option) (*DetectEntitiesOutput, error) {
	req, out := c.DetectEntitiesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetectPHI = "DetectPHI"

// DetectPHIRequest generates a "aws/request.Request" representing the
// client's request for the DetectPHI operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DetectPHI for more information on using the DetectPHI
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DetectPHIRequest method.
//    req, resp := client.DetectPHIRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectPHI
func (c *ComprehendMedical) DetectPHIRequest(input *DetectPHIInput) (req *request.Request, output *DetectPHIOutput) {
	op := &request.Operation{
		Name:       opDetectPHI,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectPHIInput{}
	}

	output = &DetectPHIOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DetectPHI API operation for AWS Comprehend Medical.
//
// Inspects the clinical text for personal health information (PHI) entities
// and entity category, location, and confidence score on that information.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Comprehend Medical's
// API operation DetectPHI for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalServerException "InternalServerException"
//   An internal server error occurred. Retry your request.
//
//   * ErrCodeServiceUnavailableException "ServiceUnavailableException"
//   The Comprehend Medical service is temporarily unavailable. Please wait and
//   then retry your request.
//
//   * ErrCodeTooManyRequestsException "TooManyRequestsException"
//   You have made too many requests within a short period of time. Wait for a
//   short time and then try your request again. Contact customer support for
//   more information about a service limit increase.
//
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request that you made is invalid. Check your request to determine why
//   it's invalid and then retry the request.
//
//   * ErrCodeInvalidEncodingException "InvalidEncodingException"
//   The input text was not in valid UTF-8 character encoding. Check your text
//   then retry your request.
//
//   * ErrCodeTextSizeLimitExceededException "TextSizeLimitExceededException"
//   The size of the text you submitted exceeds the size limit. Reduce the size
//   of the text or use a smaller document and then retry your request.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectPHI
func (c *ComprehendMedical) DetectPHI(input *DetectPHIInput) (*DetectPHIOutput, error) {
	req, out := c.DetectPHIRequest(input)
	return out, req.Send()
}

// DetectPHIWithContext is the same as DetectPHI with the addition of
// the ability to pass a context and additional request options.
//
// See DetectPHI for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ComprehendMedical) DetectPHIWithContext(ctx aws.Context, input *DetectPHIInput, opts ...request.Option) (*DetectPHIOutput, error) {
	req, out := c.DetectPHIRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// An extracted segment of the text that is an attribute of an entity, or otherwise
// related to an entity, such as the dosage of a medication taken. It contains
// information about the attribute such as id, begin and end offset within the
// input text, and the segment of the input text.
type Attribute struct {
	_ struct{} `type:"structure"`

	// The 0-based character offset in the input text that shows where the attribute
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int64 `type:"integer"`

	// The 0-based character offset in the input text that shows where the attribute
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int64 `type:"integer"`

	// The numeric identifier for this attribute. This is a monotonically increasing
	// id unique within this response rather than a global unique identifier.
	Id *int64 `type:"integer"`

	// The level of confidence that Comprehend Medical has that this attribute is
	// correctly related to this entity.
	RelationshipScore *float64 `type:"float"`

	// The level of confidence that Comprehend Medical has that the segment of text
	// is correctly recognized as an attribute.
	Score *float64 `type:"float"`

	// The segment of input text extracted as this attribute.
	Text *string `min:"1" type:"string"`

	// Contextual information for this attribute.
	Traits []*Trait `type:"list"`

	// The type of attribute.
	Type *string `type:"string" enum:"EntitySubType"`
}

// String returns the string representation
func (s Attribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Attribute) GoString() string {
	return s.String()
}

// SetBeginOffset sets the BeginOffset field's value.
func (s *Attribute) SetBeginOffset(v int64) *Attribute {
	s.BeginOffset = &v
	return s
}

// SetEndOffset sets the EndOffset field's value.
func (s *Attribute) SetEndOffset(v int64) *Attribute {
	s.EndOffset = &v
	return s
}

// SetId sets the Id field's value.
func (s *Attribute) SetId(v int64) *Attribute {
	s.Id = &v
	return s
}

// SetRelationshipScore sets the RelationshipScore field's value.
func (s *Attribute) SetRelationshipScore(v float64) *Attribute {
	s.RelationshipScore = &v
	return s
}

// SetScore sets the Score field's value.
func (s *Attribute) SetScore(v float64) *Attribute {
	s.Score = &v
	return s
}

// SetText sets the Text field's value.
func (s *Attribute) SetText(v string) *Attribute {
	s.Text = &v
	return s
}

// SetTraits sets the Traits field's value.
func (s *Attribute) SetTraits(v []*Trait) *Attribute {
	s.Traits = v
	return s
}

// SetType sets the Type field's value.
func (s *Attribute) SetType(v string) *Attribute {
	s.Type = &v
	return s
}

type DetectEntitiesInput struct {
	_ struct{} `type:"structure"`

	// A UTF-8 text string containing the clinical content being examined for entities.
	// Each string must contain fewer than 20,000 bytes of characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DetectEntitiesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectEntitiesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetectEntitiesInput"}
	if s.Text == nil {
		invalidParams.Add(request.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetText sets the Text field's value.
func (s *DetectEntitiesInput) SetText(v string) *DetectEntitiesInput {
	s.Text = &v
	return s
}

type DetectEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// The collection of medical entities extracted from the input text and their
	// associated information. For each entity, the response provides the entity
	// text, the entity category, where the entity text begins and ends, and the
	// level of confidence that Comprehend Medical has in the detection and analysis.
	// Attributes and traits of the entity are also returned.
	//
	// Entities is a required field
	Entities []*Entity `type:"list" required:"true"`

	// If the result of the previous request to DetectEntities was truncated, include
	// the Paginationtoken to fetch the next page of entities.
	PaginationToken *string `min:"1" type:"string"`

	// Attributes extracted from the input text that we were unable to relate to
	// an entity.
	UnmappedAttributes []*UnmappedAttribute `type:"list"`
}

// String returns the string representation
func (s DetectEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DetectEntitiesOutput) GoString() string {
	return s.String()
}

// SetEntities sets the Entities field's value.
func (s *DetectEntitiesOutput) SetEntities(v []*Entity) *DetectEntitiesOutput {
	s.Entities = v
	return s
}

// SetPaginationToken sets the PaginationToken field's value.
func (s *DetectEntitiesOutput) SetPaginationToken(v string) *DetectEntitiesOutput {
	s.PaginationToken = &v
	return s
}

// SetUnmappedAttributes sets the UnmappedAttributes field's value.
func (s *DetectEntitiesOutput) SetUnmappedAttributes(v []*UnmappedAttribute) *DetectEntitiesOutput {
	s.UnmappedAttributes = v
	return s
}

type DetectPHIInput struct {
	_ struct{} `type:"structure"`

	// A UTF-8 text string containing the clinical content being examined for PHI
	// entities. Each string must contain fewer than 20,000 bytes of characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectPHIInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DetectPHIInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectPHIInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetectPHIInput"}
	if s.Text == nil {
		invalidParams.Add(request.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetText sets the Text field's value.
func (s *DetectPHIInput) SetText(v string) *DetectPHIInput {
	s.Text = &v
	return s
}

type DetectPHIOutput struct {
	_ struct{} `type:"structure"`

	// The collection of PHI entities extracted from the input text and their associated
	// information. For each entity, the response provides the entity text, the
	// entity category, where the entity text begins and ends, and the level of
	// confidence that Comprehend Medical has in its detection.
	//
	// Entities is a required field
	Entities []*Entity `type:"list" required:"true"`

	// If the result of the previous request to DetectPHI was truncated, include
	// the Paginationtoken to fetch the next page of PHI entities.
	PaginationToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DetectPHIOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DetectPHIOutput) GoString() string {
	return s.String()
}

// SetEntities sets the Entities field's value.
func (s *DetectPHIOutput) SetEntities(v []*Entity) *DetectPHIOutput {
	s.Entities = v
	return s
}

// SetPaginationToken sets the PaginationToken field's value.
func (s *DetectPHIOutput) SetPaginationToken(v string) *DetectPHIOutput {
	s.PaginationToken = &v
	return s
}

// Provides information about an extracted medical entity.
type Entity struct {
	_ struct{} `type:"structure"`

	// The extracted attributes that relate to this entity.
	Attributes []*Attribute `type:"list"`

	// The 0-based character offset in the input text that shows where the entity
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int64 `type:"integer"`

	// The category of the entity.
	Category *string `type:"string" enum:"EntityType"`

	// The 0-based character offset in the input text that shows where the entity
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int64 `type:"integer"`

	// The numeric identifier for the entity. This is a monotonically increasing
	// id unique within this response rather than a global unique identifier.
	Id *int64 `type:"integer"`

	// The level of confidence that Comprehend Medical has in the accuracy of the
	// detection.
	Score *float64 `type:"float"`

	// The segment of input text extracted as this entity.
	Text *string `min:"1" type:"string"`

	// Contextual information for the entity
	Traits []*Trait `type:"list"`

	// Describes the specific type of entity with category of entities.
	Type *string `type:"string" enum:"EntitySubType"`
}

// String returns the string representation
func (s Entity) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Entity) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *Entity) SetAttributes(v []*Attribute) *Entity {
	s.Attributes = v
	return s
}

// SetBeginOffset sets the BeginOffset field's value.
func (s *Entity) SetBeginOffset(v int64) *Entity {
	s.BeginOffset = &v
	return s
}

// SetCategory sets the Category field's value.
func (s *Entity) SetCategory(v string) *Entity {
	s.Category = &v
	return s
}

// SetEndOffset sets the EndOffset field's value.
func (s *Entity) SetEndOffset(v int64) *Entity {
	s.EndOffset = &v
	return s
}

// SetId sets the Id field's value.
func (s *Entity) SetId(v int64) *Entity {
	s.Id = &v
	return s
}

// SetScore sets the Score field's value.
func (s *Entity) SetScore(v float64) *Entity {
	s.Score = &v
	return s
}

// SetText sets the Text field's value.
func (s *Entity) SetText(v string) *Entity {
	s.Text = &v
	return s
}

// SetTraits sets the Traits field's value.
func (s *Entity) SetTraits(v []*Trait) *Entity {
	s.Traits = v
	return s
}

// SetType sets the Type field's value.
func (s *Entity) SetType(v string) *Entity {
	s.Type = &v
	return s
}

// Provides contextual information about the extracted entity.
type Trait struct {
	_ struct{} `type:"structure"`

	// Provides a name or contextual description about the trait.
	Name *string `type:"string" enum:"AttributeName"`

	// The level of confidence that Comprehend Medical has in the accuracy of this
	// trait.
	Score *float64 `type:"float"`
}

// String returns the string representation
func (s Trait) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Trait) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *Trait) SetName(v string) *Trait {
	s.Name = &v
	return s
}

// SetScore sets the Score field's value.
func (s *Trait) SetScore(v float64) *Trait {
	s.Score = &v
	return s
}

// An attribute that we extracted, but were unable to relate to an entity.
type UnmappedAttribute struct {
	_ struct{} `type:"structure"`

	// The specific attribute that has been extracted but not mapped to an entity.
	Attribute *Attribute `type:"structure"`

	// The type of the attribute, could be one of the following values: "MEDICATION",
	// "MEDICAL_CONDITION", "ANATOMY", "TEST_AND_TREATMENT_PROCEDURE" or "PERSONAL_HEALTH_INFORMATION".
	Type *string `type:"string" enum:"EntityType"`
}

// String returns the string representation
func (s UnmappedAttribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UnmappedAttribute) GoString() string {
	return s.String()
}

// SetAttribute sets the Attribute field's value.
func (s *UnmappedAttribute) SetAttribute(v *Attribute) *UnmappedAttribute {
	s.Attribute = v
	return s
}

// SetType sets the Type field's value.
func (s *UnmappedAttribute) SetType(v string) *UnmappedAttribute {
	s.Type = &v
	return s
}

const (
	// AttributeNameSign is a AttributeName enum value
	AttributeNameSign = "SIGN"

	// AttributeNameSymptom is a AttributeName enum value
	AttributeNameSymptom = "SYMPTOM"

	// AttributeNameDiagnosis is a AttributeName enum value
	AttributeNameDiagnosis = "DIAGNOSIS"

	// AttributeNameNegation is a AttributeName enum value
	AttributeNameNegation = "NEGATION"
)

const (
	// EntitySubTypeName is a EntitySubType enum value
	EntitySubTypeName = "NAME"

	// EntitySubTypeDosage is a EntitySubType enum value
	EntitySubTypeDosage = "DOSAGE"

	// EntitySubTypeRouteOrMode is a EntitySubType enum value
	EntitySubTypeRouteOrMode = "ROUTE_OR_MODE"

	// EntitySubTypeForm is a EntitySubType enum value
	EntitySubTypeForm = "FORM"

	// EntitySubTypeFrequency is a EntitySubType enum value
	EntitySubTypeFrequency = "FREQUENCY"

	// EntitySubTypeDuration is a EntitySubType enum value
	EntitySubTypeDuration = "DURATION"

	// EntitySubTypeGenericName is a EntitySubType enum value
	EntitySubTypeGenericName = "GENERIC_NAME"

	// EntitySubTypeBrandName is a EntitySubType enum value
	EntitySubTypeBrandName = "BRAND_NAME"

	// EntitySubTypeStrength is a EntitySubType enum value
	EntitySubTypeStrength = "STRENGTH"

	// EntitySubTypeRate is a EntitySubType enum value
	EntitySubTypeRate = "RATE"

	// EntitySubTypeAcuity is a EntitySubType enum value
	EntitySubTypeAcuity = "ACUITY"

	// EntitySubTypeTestName is a EntitySubType enum value
	EntitySubTypeTestName = "TEST_NAME"

	// EntitySubTypeTestValue is a EntitySubType enum value
	EntitySubTypeTestValue = "TEST_VALUE"

	// EntitySubTypeTestUnits is a EntitySubType enum value
	EntitySubTypeTestUnits = "TEST_UNITS"

	// EntitySubTypeProcedureName is a EntitySubType enum value
	EntitySubTypeProcedureName = "PROCEDURE_NAME"

	// EntitySubTypeTreatmentName is a EntitySubType enum value
	EntitySubTypeTreatmentName = "TREATMENT_NAME"

	// EntitySubTypeDate is a EntitySubType enum value
	EntitySubTypeDate = "DATE"

	// EntitySubTypeAge is a EntitySubType enum value
	EntitySubTypeAge = "AGE"

	// EntitySubTypeContactPoint is a EntitySubType enum value
	EntitySubTypeContactPoint = "CONTACT_POINT"

	// EntitySubTypeEmail is a EntitySubType enum value
	EntitySubTypeEmail = "EMAIL"

	// EntitySubTypeIdentifier is a EntitySubType enum value
	EntitySubTypeIdentifier = "IDENTIFIER"

	// EntitySubTypeUrl is a EntitySubType enum value
	EntitySubTypeUrl = "URL"

	// EntitySubTypeAddress is a EntitySubType enum value
	EntitySubTypeAddress = "ADDRESS"

	// EntitySubTypeProfession is a EntitySubType enum value
	EntitySubTypeProfession = "PROFESSION"

	// EntitySubTypeSystemOrganSite is a EntitySubType enum value
	EntitySubTypeSystemOrganSite = "SYSTEM_ORGAN_SITE"

	// EntitySubTypeDirection is a EntitySubType enum value
	EntitySubTypeDirection = "DIRECTION"

	// EntitySubTypeQuality is a EntitySubType enum value
	EntitySubTypeQuality = "QUALITY"

	// EntitySubTypeQuantity is a EntitySubType enum value
	EntitySubTypeQuantity = "QUANTITY"
)

const (
	// EntityTypeMedication is a EntityType enum value
	EntityTypeMedication = "MEDICATION"

	// EntityTypeMedicalCondition is a EntityType enum value
	EntityTypeMedicalCondition = "MEDICAL_CONDITION"

	// EntityTypeProtectedHealthInformation is a EntityType enum value
	EntityTypeProtectedHealthInformation = "PROTECTED_HEALTH_INFORMATION"

	// EntityTypeTestTreatmentProcedure is a EntityType enum value
	EntityTypeTestTreatmentProcedure = "TEST_TREATMENT_PROCEDURE"

	// EntityTypeAnatomy is a EntityType enum value
	EntityTypeAnatomy = "ANATOMY"
)
