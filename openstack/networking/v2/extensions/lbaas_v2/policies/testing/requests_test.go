package testing

import (
	"testing"

	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/policies"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePolicySuccessfully(t)

	pages := 0
	err := policies.List(fake.ServiceClient(), policies.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := policies.ExtractPolcies(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 policies, got %d", len(actual))
		}
		th.CheckDeepEquals(t, PoliciesOne, actual[0])
		th.CheckDeepEquals(t, PoliciesTwo, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllPolicies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePoliciesALLSuccessfully(t)

	allPages, err := policies.List(fake.ServiceClient(), policies.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := policies.ExtractPolcies(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, PoliciesOne, actual[0])
	th.CheckDeepEquals(t, PoliciesTwo, actual[1])
}

func TestCreatePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePolicyCreationSuccessfully(t)

	actual, err := policies.Create(fake.ServiceClient(), policies.CreateOpts{
		RedirectPoolID: "431a03eb-81bb-408e-ae37-7ce19023692b",
		ListenerID:     "26058b64-6185-4e06-874e-4bd68b7633d0",
		Action:         "REDIRECT_TO_POOL",
		Name:           "test1",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, PoliciesOne, *actual)
}

//
func TestGetPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePolicyGetSuccessfully(t)

	client := fake.ServiceClient()
	actual, err := policies.Get(client, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, PoliciesOne, *actual)
}

func TestDeletePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePolicyDeletionSuccessfully(t)

	res := policies.Delete(fake.ServiceClient(), "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586")
	th.AssertNoErr(t, res.Err)
}

func TestUpdatePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePolicyUpdateSuccessfully(t)
	AdminStateUp := false
	client := fake.ServiceClient()
	actual, err := policies.Update(client, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", policies.UpdateOpts{
		AdminStateUp: &AdminStateUp,
	}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, PoliciesOne, *actual)
}

func TestListRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRulesListALLSuccessfully(t)

	pages := 0
	err := policies.ListRules(fake.ServiceClient(), policies.RulesListOpts{}, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586").EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := policies.ExtractPolicyRules(page)
		if err != nil {
			return false, err
		}

		if len(actual.Rules) != 2 {
			t.Fatalf("Expected 2 policies, got %d", len(actual.Rules))
		}
		th.CheckDeepEquals(t, RuleOne, actual.Rules[0])
		th.CheckDeepEquals(t, RuleTwo, actual.Rules[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRulesListALLSuccessfully(t)

	allPages, err := policies.ListRules(fake.ServiceClient(), policies.RulesListOpts{}, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586").AllPages()
	th.AssertNoErr(t, err)
	actual, err := policies.ExtractPolicyRules(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, RuleOne, actual.Rules[0])
	th.CheckDeepEquals(t, RuleTwo, actual.Rules[1])
}

func TestCreateRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRuleCreationSuccessfully(t)

	actual, err := policies.CreateRule(fake.ServiceClient(), policies.CreateRuleOpts{
		Type:        "HOST_NAME",
		CompareType: "EQUAL_TO",
		Value:       "www.test.com",
	}, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586").Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, RuleOne, actual.Rule)
}

//
func TestGetRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRulesGetSuccessfully(t)

	client := fake.ServiceClient()
	actual, err := policies.GetRule(client, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, RuleOne, actual.Rule)
}

func TestDeleteRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRuleDeletionSuccessfully(t)

	res := policies.DeleteRule(fake.ServiceClient(), "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRuleUpdateSuccessfully(t)
	AdminStateUp := false
	client := fake.ServiceClient()
	actual, err := policies.UpdateRule(client, policies.RuleUpdateOpts{
		AdminStateUp: &AdminStateUp,
	}, "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3").Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, RuleOne, actual.Rule)
}
