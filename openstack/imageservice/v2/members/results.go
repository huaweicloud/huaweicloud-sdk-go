package members

import (
	"time"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Member represents a member of an Image.
type Member struct {
	CreatedAt time.Time `json:"created_at"`
	ImageID   string    `json:"image_id"`
	MemberID  string    `json:"member_id"`
	Schema    string    `json:"schema"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Extract Member model from a request.
func (r commonResult) Extract() (*Member, error) {
	var s *Member
	err := r.ExtractInto(&s)
	return s, err
}

// MemberPage is a single page of Members results.
type MemberPage struct {
	pagination.SinglePageBase
}

// ExtractMembers returns a slice of Members contained in a single page
// of results.
func ExtractMembers(r pagination.Page) ([]Member, error) {
	var s struct {
		Members []Member `json:"members"`
	}
	err := r.(MemberPage).ExtractInto(&s)
	return s.Members, err
}

// IsEmpty determines whether or not a MemberPage contains any results.
func (r MemberPage) IsEmpty() (bool, error) {
	members, err := ExtractMembers(r)
	return len(members) == 0, err
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a Create operation. Call its Extract
// method to interpret it as a Member.
type CreateResult struct {
	commonResult
}

// DetailsResult represents the result of a Get operation. Call its Extract
// method to interpret it as a Member.
type DetailsResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation. Call its Extract
// method to interpret it as a Member.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// MemberSchemas presents the result of getting member schemas request
type MemberSchemas struct {
	// Name is the name of schemas
	Name string `json:"name"`
	// Properties is the explaination of schemas properties
	Properties *json.RawMessage `json:"properties"`
}

// MemberSchemasResult represents the result of member schemas request
type MemberSchemasResult struct {
	commonResult
}

// Extract interprets the result as an MemberSchemas
func (r MemberSchemasResult) Extract() (*MemberSchemas, error) {
	var s *MemberSchemas
	err := r.ExtractInto(&s)
	return s, err
}

// MembersSchemas presents the result of getting members schemas request
type MembersSchemas struct {
	// Name is the name of schemas
	Name string `json:"name"`
	// Links is the links of schemas
	Links []map[string]string `json:"links"`
	// Properties is the explaination of schemas properties
	Properties *json.RawMessage `json:"properties"`
}

// MembersSchemasResult represents the result of members schemas request
type MembersSchemasResult struct {
	commonResult
}

// Extract interprets the result as an MemberSchemas
func (r MembersSchemasResult) Extract() (*MembersSchemas, error) {
	var s *MembersSchemas
	err := r.ExtractInto(&s)
	return s, err
}
