package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallRuleCreate,
		ReadContext:   resourceFirewallRuleRead,
		UpdateContext: resourceFirewallRuleUpdate,
		DeleteContext: resourceFirewallRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getFirewallRuleSchema(),
	}
}

func resourceFirewallRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newFirewallRule := convertResourceDataToFirewallRuleCreateAPIObject(d)

	//Call the API
	createdFirewallRule, err := coxEdgeClient.CreateFirewallRule(newFirewallRule)
	if err != nil {
		return diag.FromErr(err)
	}

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, createdFirewallRule.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceFirewallRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	firewallRule, err := coxEdgeClient.GetFirewallRule(d.Get("site_id").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertFirewallRuleAPIObjectToResourceData(d, firewallRule)

	//Update state
	resourceFirewallRuleRead(ctx, d, m)

	return diags
}

func resourceFirewallRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedFirewallRule := convertResourceDataToFirewallRuleCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateFirewallRule(resourceId, updatedFirewallRule)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceFirewallRuleRead(ctx, d, m)
}

func resourceFirewallRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Delete the FirewallRule
	err := coxEdgeClient.DeleteFirewallRule(d.Get("site_id").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToFirewallRuleCreateAPIObject(d *schema.ResourceData) apiclient.FirewallRule {
	//Create update firewallRule struct
	updatedFirewallRule := apiclient.FirewallRule{
		Action:  d.Get("action").(string),
		Enabled: d.Get("enabled").(bool),
		Id:      d.Get("id").(string),
		IpEnd:   d.Get("ip_end").(string),
		IpStart: d.Get("ip_start").(string),
		Name:    d.Get("name").(string),
		SiteId:  d.Get("site_id").(string),
	}

	return updatedFirewallRule
}

func convertFirewallRuleAPIObjectToResourceData(d *schema.ResourceData, firewallRule *apiclient.FirewallRule) {
	//Store the data
	d.Set("id", firewallRule.Id)
	d.Set("site_id", firewallRule.SiteId)
	d.Set("action", firewallRule.Action)
	d.Set("ip_start", firewallRule.IpStart)
	d.Set("name", firewallRule.Name)
	d.Set("enabled", firewallRule.Enabled)
	d.Set("ip_end", firewallRule.IpEnd)
}
