package lang

import (
	"github.com/Aptomi/aptomi/pkg/runtime"
)

// ACLRule defines which users have which roles in Aptomi. They should be configured by Aptomi domain admins in the
// policy. ACLRules allow to pick groups of users and assign ACL roles to them (e.g. give access to a particular
// namespace)
type ACLRule = Rule

// ACLRuleObject is an informational data structure with Kind and Constructor for ACLRule
var ACLRuleObject = &runtime.Info{
	Kind:        "aclrule",
	Storable:    true,
	Versioned:   true,
	Deletable:   true,
	Constructor: func() runtime.Object { return &ACLRule{} },
}

// Allows to define a role which spans across all namespaces (e.g. "domain admin")
const namespaceAll = "*"

// ACLRole is a struct for defining user roles and their privileges.
// Aptomi has 4 built-in user roles: domain admin, namespace admin, service consumer, and nobody.
// Domain admin has full access rights to all namespaces. It can manage global objects in 'system' namespace (clusters,
// rules, and ACL rules).
// Namespace admin has full access right to a given set of namespaces, but it cannot global objects in 'system' namespace (clusters,
// rules, and ACL rules).
// Service consumer can only consume services within a given set of namespaces. Service consumption is treated as capability
// to instantiate services in a given namespace.
// Nobody cannot do anything except viewing the policy.
type ACLRole struct {
	ID         string
	Name       string
	Privileges *Privileges
}

// Privileges defines a set of privileges for a particular role in Aptomi
type Privileges struct {
	// AllNamespaces, when set to true, indicated that user privileges apply to all namespaces. Otherwise it applies
	// to a set of given namespaces
	AllNamespaces bool

	// NamespaceObjects specifies whether or not this role can view/manage a certain object kind within a non-system namespace
	NamespaceObjects map[string]*Privilege

	// GlobalObjects specifies whether or not this role can view/manage a certain object kind within a system namespace
	GlobalObjects map[string]*Privilege
}

// Returns privileges for a given object
func (privileges *Privileges) getObjectPrivileges(obj Base) *Privilege {
	var result *Privilege
	if obj.GetNamespace() == runtime.SystemNS {
		result = privileges.GlobalObjects[obj.GetKind()]
	} else {
		result = privileges.NamespaceObjects[obj.GetKind()]
	}
	if result == nil {
		return noAccess
	}
	return result
}

// Privilege is a unit of privilege for any single given object
type Privilege struct {
	// View indicates whether or not a user can view an object (R)
	View bool

	// Manage indicates whether or not a user can manage an object, i.e. perform operations (CUD)
	Manage bool
}

// Full access privilege
var fullAccess = &Privilege{
	View:   true,
	Manage: true,
}

// View access privilege
var viewAccess = &Privilege{
	View: true,
}

// No access privilege
var noAccess = &Privilege{}

// DomainAdmin is a built-in domain admin role
var DomainAdmin = &ACLRole{
	ID:   "domain-admin",
	Name: "Domain Admin",
	Privileges: &Privileges{
		AllNamespaces: true,
		NamespaceObjects: map[string]*Privilege{
			ServiceObject.Kind:    fullAccess,
			ContractObject.Kind:   fullAccess,
			DependencyObject.Kind: fullAccess,
			RuleObject.Kind:       fullAccess,
		},
		GlobalObjects: map[string]*Privilege{
			ClusterObject.Kind: fullAccess,
			RuleObject.Kind:    fullAccess,
			ACLRuleObject.Kind: fullAccess,
		},
	},
}

// NamespaceAdmin is a built-in admin role
var NamespaceAdmin = &ACLRole{
	ID:   "namespace-admin",
	Name: "Namespace Admin",
	Privileges: &Privileges{
		NamespaceObjects: map[string]*Privilege{
			ServiceObject.Kind:    fullAccess,
			ContractObject.Kind:   fullAccess,
			DependencyObject.Kind: fullAccess,
			RuleObject.Kind:       fullAccess,
		},
		GlobalObjects: map[string]*Privilege{
			ClusterObject.Kind: viewAccess,
			RuleObject.Kind:    viewAccess,
			ACLRuleObject.Kind: viewAccess,
		},
	},
}

// ServiceConsumer is a built-in service consumer role
var ServiceConsumer = &ACLRole{
	ID:   "service-consumer",
	Name: "Service Consumer",
	Privileges: &Privileges{
		NamespaceObjects: map[string]*Privilege{
			ServiceObject.Kind:    viewAccess,
			ContractObject.Kind:   viewAccess,
			DependencyObject.Kind: fullAccess,
			RuleObject.Kind:       viewAccess,
		},
		GlobalObjects: map[string]*Privilege{
			ClusterObject.Kind: viewAccess,
			RuleObject.Kind:    viewAccess,
			ACLRuleObject.Kind: viewAccess,
		},
	},
}

// Nobody role
var nobody = &ACLRole{
	ID:   "nobody",
	Name: "Nobody",
	Privileges: &Privileges{
		NamespaceObjects: map[string]*Privilege{
			ServiceObject.Kind:    viewAccess,
			ContractObject.Kind:   viewAccess,
			DependencyObject.Kind: viewAccess,
			RuleObject.Kind:       viewAccess,
		},
		GlobalObjects: map[string]*Privilege{
			ClusterObject.Kind: viewAccess,
			RuleObject.Kind:    viewAccess,
			ACLRuleObject.Kind: viewAccess,
		},
	},
}

// ACLRolesOrderedList represents the ordered list of ACL roles (from most "powerful" to least "powerful")
var ACLRolesOrderedList = []*ACLRole{
	DomainAdmin,
	NamespaceAdmin,
	ServiceConsumer,
	nobody,
}

// ACLRolesMap represents the map of ACL roles (Role ID -> Role)
var ACLRolesMap = map[string]*ACLRole{
	DomainAdmin.ID:     DomainAdmin,
	NamespaceAdmin.ID:  NamespaceAdmin,
	ServiceConsumer.ID: ServiceConsumer,
	nobody.ID:          nobody,
}
