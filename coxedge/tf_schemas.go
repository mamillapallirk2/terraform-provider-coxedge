/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func getOrganizationSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"organizations": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOrganizationSchema(),
			},
			Description: "Organization details",
		},
	}
}

func getOrganizationBillingInfoSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the billing information",
		},
		"organizations_billing_info": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOrganizationBillingInfoSchema(),
			},
			Description: "The organization billing information",
		},
	}
}

func getOrganizationBillingInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The id of the billing information",
		},
		"organization_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_provider_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_masked_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_exp": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_line_one": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_line_two": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_city": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_province": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_postal_code": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_postal_country": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func getOrganizationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The id of the organization",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name of the organization",
		},
		"entry_point": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The entry point of the organization is the subdomain of the organization in the Cox Edge URL : [entryPoint].Cox Edge",
		},
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Tags associated to the organization",
		},
		"service_connections": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The services for which the organization is allowed to provision resources. Includes: id,serviceCode",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"service_code": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}

func getEnvironmentSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"environments": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Environment descriptions",
			Elem: &schema.Resource{
				Schema: getEnvironmentSchema(),
			},
		},
	}
}

func getRolesSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"roles": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getRolesSchema(),
			},
		},
	}
}

func getRolesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"is_system": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"default_scope": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func getEnvironmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "The id of the environment.",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the new environment. Should be unique in the environment and only contain lower case characters, numbers, dashes and underscores.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The description of the environment.",
		},
		"membership": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of membership of the environment. ALL_ORG_USERS will add every user in the organization to this environment with the default role. MANY_USERS will allow you to choose the users you want in the environment and assigned them specific roles. Defaults to MANY_USERS.",
		},
		"organization_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The organization that the environment should be created in. Defaults to your organization.Required: id",
		},
		"service_connection_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The service connection that the environment should be created in. Required: id",
		},
		"creation_date": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date in ISO 8601 that the environment was created.",
		},
		"roles": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The roles of the environment and the users assigned to them. Also, defines the default role of the environment.\nrequired: name, users.id\noptional: isDefault",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Name of the role",
					},
					"is_default": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Set to default role",
					},
					"users": {
						Type:        schema.TypeList,
						Required:    true,
						Description: "Array of users Id",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func getUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Username of the new user. Should be unique across the organization.",
		},
		"first_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "First name of the user.",
		},
		"last_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Last name of the user.",
		},
		"email": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Email of the user. Should be unique across the organization.",
		},
		"organization_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Organization in which the user will be created. Defaults to your organization. Required: id",
		},
		"roles": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Role Id",
					},
				},
			},
			Optional:    true,
			Description: "The system and environment roles to give to the user.",
		},
		"last_updated": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func getWorkloadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the workload.",
		},
		"image": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Either the location of a Docker image to run as a container or the image to use for the virtual machine. If for a virtual machine, this is in the format of /[:]. If the image tag portion is omitted, 'default' is assumed which is the most recently created, ready, and non-deprecated image of that slug. A set of common images is present on the 'cox-edge' stack.",
		},
		"specs": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Specification type for resources which are allocated to each instance in a workload. Supported specifications are SP-1 (1 vCPU, 2 GB RAM),SP-2 (2 vCPU, 4 GB RAM),SP-3 (2 vCPU, 8GB RAM),SP-4 (4 vCPU, 16 GB RAM),SP-5 (8 vCPU, 32 GB RAM).",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Specify whether a workload is a VM-based workload or container-based. Can be either VM or CONTAINER.",
		},
		"deployment": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The name of the deployment.",
					},
					"pops": {
						Type: schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Required:    true,
						Description: "The points of presence of a deployment. In the regex format [A-Z]{3, 3}.",
					},
					"enable_autoscaling": {
						Type:        schema.TypeBool,
						Default:     false,
						Optional:    true,
						Description: "Specifies if autoscaling is enabled. If enabled, then cpuUtilization , minInstancesPerPop and maxInstancesPerPop are required.",
					},
					"instances_per_pop": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     -1,
						Description: "The number of instances per point of presence. Only applicable if autoscaling is not enabled.",
					},
					"max_instances_per_pop": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "The maximum number of instances per PoP. Only applicable if autoscaling is enabled. Should be greater than zero and less than 50.",
					},
					"min_instances_per_pop": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "The minimum number of instances per PoP. Only applicable if autoscaling is enabled. Should be greater than zero and less than 50.",
					},
					"cpu_utilization": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "The percentage of CPU utilization. Only applicable if autoscaling is enabled.",
					},
				},
			},
		},
		"add_anycast_ip_address": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Option to AnyCast IP Address.",
		},
		"anycast_ip_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "The Anycast IP address assigned to a workload. If there is no IP assigned to the workload then the value of this attribute will be None.",
		},
		"commands": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "The commands that start a container. Only applicable to workloads of type 'CONTAINER'. Commands cannot be updated or removed after workload creation.",
		},
		"container_email": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The email address to use for the docker registry account",
		},
		"container_username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username used to authenticate the image pull.",
		},
		"container_password": {
			Type:        schema.TypeString,
			Sensitive:   true,
			Optional:    true,
			Description: "The password used to authenticate the image pull.",
		},
		"container_server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The server that the credentials should be used with. This value will default to the docker hub registry when not set.",
		},
		"environment_variables": {
			Type:        schema.TypeMap,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "A list of environment variables. Only applicable to workloads of type 'CONTAINER'.",
		},
		"first_boot_ssh_key": {
			Type:        schema.TypeString,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "If creating a VM-based workload, SSH keys are required. Multiple SSH keys can be separated by newlines \\n.",
		},
		"ports": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of network interfaces that will be created for each workload instance.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"protocol": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Protocol for the network policy rule. Supported protocols are: TCP, UDP and TCP_UDP.",
					},
					"public_port": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "A single port, such as 80 or a port range, such as 1024-65535 for which a network policy rule will be created for the workload.",
					},
					"public_port_desc": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "A summary of what the network policy rule does or a name for it. It is highly recommended to give a unique description to easily identify a network policy rule. Defaults to an empty string if not provided.",
					},
					"public_port_src": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "A subnet that will define all the IPs allowed by the network policy rule. Defaults to 0.0.0.0/0 if not specified.",
					},
				},
			},
		},
		"persistent_storages": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Persistent storage volumes used by the workload.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"path": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The path in an instance to mount a volume.",
					},
					"size": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The size of the mounted volume (in GB).",
					},
				},
			},
		},
		"secret_environment_variables": {
			Type:        schema.TypeMap,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "A list of sensitive environment variables. Only applicable to workloads of type 'CONTAINER'.",
		},
		"slug": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A workload's programmatic name. Workload slugs are used to build its instances names. If not provided, defaults to workload's name. It must not exceed 18 characters.",
		},
	}
}

func getImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Required: true,
		},
		"family": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tag": {
			Type:     schema.TypeString,
			Required: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"reference": {
			Type:     schema.TypeString,
			Required: true,
		},
		"status": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func getImageSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"environment": {
			Type:     schema.TypeString,
			Required: true,
		},
		"images": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getImageSchema(),
			},
		},
	}
}

func getNetworkPolicyRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the network policy rule, in the form networkProfileId/type/hashCode/occurrence.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the stack to which the network policy belongs.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"workload_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The UUID of the workload to which the network policy rule is applied. Corresponds to the first workload ID in the network policy's list of instance selectors.",
		},
		"network_policy_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the network policy to which the network policy rule belongs.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A summary of what this rule does or a name of this rule. It is highly recommended to give a unique description to easily identify a rule.",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The type of network policy rule, either INBOUND or OUTBOUND.",
		},
		"source": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A subnet that will define all the IPs allowed or denied by this rule.",
		},
		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The network policy rule action: ALLOW (allow traffic) or BLOCK (deny traffic).",
		},
		"protocol": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Supported protocols are: TCP, UDP, TCP_UDP, ESP, AH, ICMP or GRE.",
		},
		"port_range": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "This specifies on which ports traffic will be allowed or denied by this rule. It can be a range of ports separated by a hyphen.",
		},
	}
}

func getSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"services": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required:    true,
			Description: "Services list that will be used on the site. Possibles values are CDN,SERVERLESS_EDGE_ENGINE or WAF.",
		},
		"protocol": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Protocol that will be used to communicate with the hostname. Possibles values are HTTP or HTTPS.",
		},
		"domain": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The domain name that will be used for the site.",
		},
		"hostname": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The hostname to that will be used to get the information from. The hostname can be an IP or a name. It may include a specific port and a precise path as well (e.g. 199.250.204.212:80/test).",
		},
		"auth_method": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The authentication method to communicate with the hostname. Possibles values are NONE or BASIC. If not provided, it will default to NONE unless the username or password is provided. It would then default to BASIC.",
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username for the basic authentication. Required if authMethod is BASIC or if the password id provided.",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The password for the basic authentication. Required if authMethod is BASIC or if the password id provided.",
		},
		"operation": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "To enable or disable CDN, WAF and Serverless Scripts. Required values: CDN -> enable_cdn/disable_cdn, WAF -> enable_waf/disable_waf, Serverless Scripts -> enable_scripts/disable_scripts",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				val := false
				switch value {
				case "enable_cdn":
					val = true
					break
				case "disable_cdn":
					val = true
					break
				case "enable_waf":
					val = true
					break
				case "disable_waf":
					val = true
					break
				case "enable_scripts":
					val = true
					break
				case "disable_scripts":
					val = true
					break
				}
				if !val {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("opertaion field: %q should be either one of following - enable_cdn, disable_cdn, enable_waf, disable_waf, enable_scripts, disable_scripts", value),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		//Computed properties
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A sites's unique identifier.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the stack that a site belongs to.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The status of the site. It can either be ACTIVE, PENDING, or PROVISIONING.",
		},
		"edge_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The edge address of the site.",
		},
		"anycast_ip": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Anycast IP address that domains should be pointed to.",
		},
		"delivery_domains": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "List of delivery domains of the site.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"domain": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "A delivery domain of the site.",
						ForceNew:    true,
					},
					"validated_at": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The date the domain was validated to be pointing to Cox.",
						ForceNew:    true,
					},
				},
			},
		},
	}
}

func getOriginSettingSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"origin_settings": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOriginSettingsSchema(),
			},
		},
	}
}

func getOriginSettingOriginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "An origin's unique identifier.",
		},
		"address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The address of the primary origin that the CDN uses to pull content from. Can be a valid IPv4 address or a valid domain name. It may include a specific port and a precise path as well (e.g. 199.250.204.212:80/test). Port must be one of [80, 8080, 443, 1935, 9091].",
		},
		"common_certificate_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Common name to validate SSL origin requests against.",
		},
		"auth_method": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the authentication method that the origin uses. Must be one of [\"NONE\", \"BASIC\"].",
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Username to use when authenticating with the origin.",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Password to use when authenticating with the origin.",
		},
	}
}

func getOriginSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"site_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A site's unique identifier.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the stack that a site belongs to.",
		},
		"scope_configuration_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the scope of the site that the origins are connected to.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The domain of the site.",
		},
		"websockets_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies if web socket connections to the origin server are enabled.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"ssl_validation_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies if SSL validation for the origins is enabled.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"pull_protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The type of protocol used to pull content from the origin. Must be one of [\"HTTP\", \"HTTPS\", \"MATCH\"]. \"MATCH\" is equivalent to \"HTTP or HTTPS\".",
		},
		"host_header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The host header to be used to pull content from the origin. \"Dynamic\" refers to using the requested domain name (Host: %client.request.host%) as the host header.",
		},
		"origin": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: getOriginSettingOriginSchema(),
			},
			Optional:    true,
			Description: "The primary origin that the CDN uses to pull content from.",
		},
		"backup_origin_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies if a backup origin for the site is configured.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"backup_origin": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: getOriginSettingOriginSchema(),
			},
			Optional:    true,
			Description: "The secondary origin that the CDN uses to pull content from when the primary origin is not available.",
		},
		"backup_origin_exclude_codes": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "Requests are made to the backup origin on any 4xx or 5xx response codes returned from the primary origin. This property specifies the response status codes for which calls to the backup origin must not be made. Multiple response codes can be excluded. e.g: [\"410\", \"411\", \"412\"]. Asterisks can be used to cover a range of codes. e.g. All the 4xx codes can be covered using \"4*\".",
		},
	}
}

func getDeliveryDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The delivery domain unique identifier.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the stack that the site belongs to.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"domain": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The site's domain name.",
		},
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the site for which to list delivery domains. This parameter is required.",
		},
	}
}

func getCDNSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A site's unique identifier.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"cache_expire_policy": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A site's cache expiry policy. Can be ORIGIN_CONTROLLED, SPECIFY_CDN_TTL, NEVER_EXPIRE, or DO_NOT_CACHE.",
		},
		"cache_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The time to live for the cache, in seconds. Depends on the cache expiry policy.",
		},
		"query_string_control": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "The strategy for caching query strings. Can be IGNORE, CACHE_ALL or CUSTOM.",
		},
		"custom_cached_query_strings": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "List of custom cached query strings. Only visible if the queryStringControl attribute is CUSTOM.",
		},
		"dynamic_caching_by_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable dynamic caching by headers.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"custom_cached_headers": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "A list of custom cached headers. Only visible if dynamicCachingByHeaderEnabled is true.",
		},
		"gzip_compression_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable gzip compression.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"gzip_compression_level": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The level for the gzip compression. Values are between 1 to 6. Only visible is gzipCompressionEnabled is true.",
		},
		"content_persistence_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not make cached content available after its expiration time.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"maximum_stale_file_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The maximum time to live for stale files, in seconds. Only visible if contentPersistenceEnabled is true.",
		},
		"vary_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable honoring the vary header in a request.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"browser_cache_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Sets the default browser expiration time for cached assets, in seconds.",
		},
		"cors_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets the Access-Control-Allow-Origin header to allow browsers to access this domain from other origins.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"allowed_cors_origins": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The strategy for allowing cors origins. Can be SPECIFY_ORIGINS or ALL_ORIGINS.",
		},
		"origins_to_allow_cors": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "A list of origins to allow cors requests from. Only visible if allowedCorsOrigins is set to SPECIFY_ORIGINS.",
		},
		"http2_support_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable supporting applications using HTTP/2 protocol.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"http2_server_push_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to push assets to the client or browser (user) in advance (before the user requests these assets) which enables faster load times.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"link_header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The link header for http2ServerPush, only visible if http2ServerPushEnabled is true.",
		},
		"canonical_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable setting Link: http://{hostname}/URI; rel=\"canonical\" header on each response.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"canonical_header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The hostname for the canonicalHeader, only visible if canonicalHeaderEnabled is true.",
		},
		"url_caching_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable caching of URLs without file extensions.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"url_caching_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The time to live for the url cache. Only visible if urlCachingEnabled is true.",
		},
	}
}

func getCDNPurgeResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "A site's unique identifier.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"purge_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "URL",
			ForceNew:    true,
			Description: "The type of cache purge. Can be URL or PATH. Default value is URL.",
		},
		"items": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "The items to purge from the CDN.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"url": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The URL or path at which to delete content.",
					},
					"recursive": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether or not to recursively delete content from the CDN.",
					},
					"invalidate_only": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether or not to mark the asset as expired and re-validate instead of deleting.",
					},
					"purge_all_dynamic": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether or not to purge dynamic versions of assets.",
					},
					"headers": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "A list of HTTP request headers used to construct a cache key to purge content by. These headers must be configured in the site configuration's DynamicContent.headerFields property.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"purge_selector": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "A key/value pair definition of content to purge from the CDN.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"selector_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The name of the type of content to purge. For example, the name of the HTTP response header. Names are case sensitive.",
								},
								"selector_type": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The kinds of content that can be purged from the CDN. One of: HEADER (Purge content based on an HTTP response header), TAG (Purge content based on an X-TAG HTTP header value. Purging by tag can be useful when content on the origin is tagged).",
								},
								"selector_value": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The value of the content to purge. For example, the value of the HTTP response header. Values are case sensitive and may be wild-carded, but cannot match a \"/\".",
								},
								"selector_value_delimiter": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The delimiter to separate multiple values with. Defaults to \",\".",
								},
							},
						},
					},
				},
			},
		},
	}
}

func getWAFSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the site for which the WAF is applied to.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the stack that a site belongs to.",
		},
		"domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The domain of the site.",
		},
		"api_urls": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "List of configured API urls.",
		},
		"ddos_settings": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The DDoS Setting containing the different threshold values.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"global_threshold": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The number of overall requests per ten seconds that can trigger DDoS protection.",
					},
					"burst_threshold": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The number of requests per two seconds that can trigger DDoS protection.",
					},
					"subsecond_burst_threshold": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The number of requests per 0.1 seconds that can trigger DDoS protection.",
					},
				},
			},
		},
		"monitoring_mode_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If the monitoring mode is enabled.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"owasp_threats": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Cox’s core rule set & OWASP’s most critical Web application security risks.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"sql_injection": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of being a SQL injection attack attempt. SQL injection attacks attempt to exploit vulnerabilities in a Web application's code and seek to gain access and control over the database. A successful attack would typically result in stolen data or the site being defaced or taken down.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"xss_attack": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of being a Cross-Site-Scripting attack attempt. Cross Site Scripting attacks attempt to exploit vulnerabilities in a Web application and seek to inject a client side script either across an entire site or to a specific user's session. A successful attack would typically allow forbidden access to a user's actions and data.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"shell_shock_attack": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of being a Shellshock attack attempt. A Shellshock attack is an attempt to exploit a server's vulnerabilities to gain full access and control over them. A successful attack would typically either abuse a server's resources or hack the website.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"remote_file_inclusion": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of being a Remote File Inclusion attempt. Remote File Inclusion attempts to exploit vulnerabilities in a Web application (typically in PHP) to execute a script from a 3rd party server. RFI attacks provide a backdoor for the hacker to change the behaviour of a server and Web application.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"apache_struts_exploit": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Patch known vulnerabilities in the Apache Struts framework by blocking requests suspected of exploiting these vulnerabilities.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"local_file_inclusion": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of a Local File Inclusion attempt. Local File Inclusion attempts seek to exploit vulnerabilities in a Web application to execute potentially harmful scripts on your servers.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"common_web_application_vulnerabilities": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block attempts to access and potentially harm your servers through common backdoors, such as common control panels, configuration scripts etc. which may be accessible to unwanted users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"webshell_execution_attempt": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of Web shell attempts. A Web shell is a script that can be uploaded to a Web server to enable remote administration of the machine. Infected Web servers can either be internet-facing or internal to the network, where the Web shell is used to further pivot to internal hosts.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"protocol_attack": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"csrf": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Cox WAF will generate a CSRF token that is added to forms. Requests without a valid CSRF token will be blocked.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"open_redirect": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of being an Open Redirect attempt. Open Redirect attempts to exploit vulnerabilities in a Web application to redirect a user to a new website without any validation of the target of redirect.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"shell_injection": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block requests suspected of being a shell injection attack attempt. Shell Injection is an attack in which the goal is execution of arbitrary commands on the host operating system via a vulnerable application. Command injection attacks are possible when an application passes unsafe user supplied data (forms, cookies, HTTP headers etc.) to a system shell.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"code_injection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sensitive_data_exposure": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"xml_external_entity": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"personal_identifiable_info": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"serverside_template_injection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"general_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"block_invalid_user_agents": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"block_unknown_user_agents": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"http_method_validation": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"traffic_sources": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Real-time threat intelligence for IP addresses, source location, and information on malicious IPs.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"via_tor_nodes": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge traffic from The Onion Ring exit nodes to block bots and known bad devices. While TOR is used sometimes purely for Web anonymity, it is commonly used by hackers, scrapers, and spammers to crawl or hack Web applications.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"via_proxy_networks": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge traffic from any known proxy network to block bots and known bad devices. While proxy services are used sometimes purely for Web anonymity, they are also commonly used by hackers, scrapers, and spammers to crawl or hack Web applications.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"via_hosting_services": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge traffic from IP addresses known to be of hosting service companies. This rule is unlikely to see legitimate human traffic on these IP spaces since they are typically used for server hosting. In most cases, traffic from these IP spaces originate from infected servers that are controlled by hackers.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"via_vpn": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge traffic from any known VPN to block bots and known bad devices. While VPNs are sometimes used purely for Web anonymity, they are also commonly used by hackers, scrapers, and spammers to crawl or hack Web applications.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"convicted_bot_traffic": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge traffic from IP addresses that have been convicted of automated activities (bots) on this site or on others. These IP addresses are used by malicious automated agents while no legitimate traffic has been observed on them.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"traffic_from_suspicious_nat_ranges": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge traffic from suspicious NAT ranges.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"external_reputation_block_list": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"traffic_via_cdn": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"anti_automation_bot_protection": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Block automated traffic from scanning and browsing your online application.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"force_browser_validation_on_traffic_anomalies": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge and block requests if the user or device behind them does not keep session cookies and does not execute JavaScripts correctly. Most malicious automated activities (bots) do not meet these conditions and will, therefore, effectively be blocked by the JavaScript challenge triggered in any suspected situation. Clients can also be blocked depending on whether they act in an abnormal to the specific domain—by scraping content in a way that most sessions on this domain don't—or clients that try to, for example, avoid detection by switching IPs.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"challenge_automated_clients": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Captcha-challenge and block sessions conducted by standard Web browsers if there is evidence that these sessions are being automated and not driven by a human user. Such automation is used primarily for screen scraping and other very targeted, site-specific malicious automation.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"challenge_headless_browsers": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge requests if the user or device behind them uses an automation tool that initiates browsers but is actually an automation tool without real display—such as phantomJS, Selenium, or other. While such tools are favored by programmers, they are also extremely popular with scrapers, hackers and even in sophisticated DDoS attacks to circumvent standard anti-bot measures.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"anti_scraping": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A more hardened anti-automation policy that is meant to stop scrapers by using faster and harsher convictions.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"behavioral_waf": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Cox's sophisticated user behaviour and reputation analysis rules.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"spam_protection": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge and block user sessions and activities that seem to be aggressively using forms on your website to post spam content, generate new accounts, and more. Also, require a handshake (if not already provided) to clients making POST requests.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"block_probing_and_forced_browsing": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge or block sessions and users that seem to make brute-forced requests on random URLs seeking to discover a Web application's structure and hidden directories.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"obfuscated_attacks_and_zeroday_mitigation": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Block clients performing multiple injection attacks.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"repeated_violations": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge or block clients that failed to answer previous challenges.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"bruteforce_protection": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Challenge and block attempts seeking to guess user names and passwords on Web login forms.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"cms_protection": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Whitelist admin users.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"wordpress_waf_ruleset": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_wordpress": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable whitelist WordPress admin logged-in users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_modx": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable whitelist MODX admin logged-in users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_drupal": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable whitelist Drupal admin logged-in users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_joomla": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable whitelist Joomla admin logged-in users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_magento": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable whitelist Magento admin logged-in users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_origin_ip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable this policy to whitelist requests coming from the origin for plugin updates and general CMS updates",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_umbraco": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable whitelist Umbraco admin logged-in users.",
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"allow_known_bots": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "An object containing known bots.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"acquia_uptime": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"add_search_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"adestra_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"adjust_servers": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ahrefs_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"alerta_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"alexa_ia_archiver": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"alexa_technologies": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"amazon_route_53_health_check_service": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"applebot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"apple_news_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ask_jeeves_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"audisto_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"baidu_spider_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"baidu_spider_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"binary_canary": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"bitbucket_webhook": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"blekko_scout_jet_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"chrome_compression_proxy": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"coccocbot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"cookie_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"cybersource": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"daumoa_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"detectify_scanner": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"digi_cert_dcv_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"dotmic_dot_bot_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"duck_duck_go_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"facebook_external_hit_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"feeder_co": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"feed_press": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"feed_wind": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"freshping_monitoring": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"geckoboard": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ghost_inspector": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"gomez": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"goo_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_ads_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_cloud_monitoring_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_feed_fetcher_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_image_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_image_proxy": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_mediapartners_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_mobile_ads_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_news_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_page_speed_insights": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_structured_data_testing_tool": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_verification_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_video_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_web_light": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"grapeshot_bot_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"gree_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"hetrix_tools": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"hi_pay": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"hyperspin_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ias_crawler_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"internet_archive_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"jetpack_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"jike_spider_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"j_word_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"kakao_user_agent": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"kyoto_tohoku_crawler": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"landau_media_spider": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"lets_encrypt": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"line_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"linked_in_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"livedoor_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"mail_ru_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"manage_wp": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_bing_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_bing_preview_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_msn_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_skype_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"mixi_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"mobage_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"naver_yeti_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"new_relic_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ocn_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"panopta_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"parse_ly_scraper": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"pay_pal_ipn": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"petal_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"pingdom": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"pinterest_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"qwantify_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"roger_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sage_pay": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sectigo_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"semrush_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"server_density_service_monitoring_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"seznam_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"shareaholic_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"site_24_x_7_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"siteimprove_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"site_lock_spider": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"slack_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sogou_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"soso_spider_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"spatineo": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"spring_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"stackify": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"status_cake_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"stripe": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sucuri_uptime_monitor_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"telegram_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"testomato_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"the_find_crawler": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"twitter_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"uptime_robot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"vkontakte_external_hit_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"w_3_c": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"wordfence_central": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"workato": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"xml_sitemaps": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_inktomi_slurp_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_link_preview": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_seeker_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_slurp_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yandex_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yisou_spider_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yodao_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"zendesk_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"zoho_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"zum_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
	}
}

func getFirewallRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier for the rule.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the site for which the firewall rule is applied to.",
		},
		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Either ALLOW or BLOCK.",
		},
		"ip_start": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The start ip address for the rule. When no ipEnd attribute is provided, the rule only applies for the ip provided in ipStart.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the rule.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether or not the rule is enabled. The default value is false.",
		},
		"ip_end": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The end ip address for the rule.",
		},
	}
}

func getScriptSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier for the script.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the stack that the script belongs to.",
		},
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the site that the script belongs to.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the script.",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Creation timestamp of the script.",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date on which the script was last updated.",
		},
		"version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The version number of the script.",
		},
		"code": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The JavaScript code used for the script.",
		},
		"routes": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required:    true,
			Description: "The routes that incoming requests should respond with a script.",
		},
	}
}
