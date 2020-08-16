package config

import (
	"strings"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gofiber/fiber"
)

// Config holds the configuration for the middleware
type PermissionMiddleware struct {
	// ModelFilePath is path to model file for Casbin.
	// Optional. Default: "./model.conf".
	Enforcer *casbin.Enforcer

	// PolicyAdapter is an interface for different persistent providers.
	// Optional. Default: fileadapter.NewAdapter("./policy.csv").
	PolicyAdapter *gormadapter.Adapter

	// Lookup is a function that is used to look up current subject.
	// An empty string is considered as unauthenticated user.
	// Optional. Default: func(c *fiber.Ctx) string { return "" }
	Lookup func(*fiber.Ctx) string

	// Unauthorized defines the response body for unauthorized responses.
	// Optional. Default: func(c *fiber.Ctx) string { c.SendStatus(401) }
	Unauthorized func(*fiber.Ctx)

	// Forbidden defines the response body for forbidden responses.
	// Optional. Default: func(c *fiber.Ctx) string { c.SendStatus(403) }
	Forbidden func(*fiber.Ctx)
}

type validationRule int

const (
	matchAll validationRule = iota
	atLeastOne
)

// MatchAll is an option that defines all permissions
// or roles should match the user.
var MatchAll = func(o *Options) { //nolint:gochecknoglobals
	o.ValidationRule = matchAll
}

// AtLeastOne is an option that defines at least on of
// permissions or roles should match to pass.
var AtLeastOne = func(o *Options) { //nolint:gochecknoglobals
	o.ValidationRule = atLeastOne
}

// PermissionParserFunc is used for parsing the permission
// to extract object and action usually
type PermissionParserFunc func(str string) []string

func permissionParserWithSeperator(sep string) PermissionParserFunc {
	return func(str string) []string {
		return strings.Split(str, sep)
	}
}

// PermissionParserWithSeperator is an option that parses permission
// with seperators
func PermissionParserWithSeperator(sep string) func(o *Options) {
	return func(o *Options) {
		o.PermissionParser = permissionParserWithSeperator(sep)
	}
}

// Options holds options of middleware
type Options struct {
	ValidationRule   validationRule
	PermissionParser PermissionParserFunc
}

// RequiresPermissions tries to find the current subject and determine if the
// subject has the required permissions according to predefined Casbin policies.
func (cm *PermissionMiddleware) RequiresPermissions(permissions []string, opts ...func(o *Options)) func(*fiber.Ctx) {

	options := &Options{
		ValidationRule:   matchAll,
		PermissionParser: permissionParserWithSeperator(":"),
	}

	for _, o := range opts {
		o(options)
	}

	return func(c *fiber.Ctx) {
		if len(permissions) == 0 {
			c.Next()
			return
		}

		sub := cm.Lookup(c)
		if len(sub) == 0 {
			cm.Unauthorized(c)
			return
		}

		if options.ValidationRule == matchAll {
			for _, permission := range permissions {
				vals := append([]string{sub}, options.PermissionParser(permission)...)
				if ok, err := cm.Enforcer.Enforce(convertToInterface(vals)...); err != nil {
					c.SendStatus(fiber.StatusInternalServerError)
					return
				} else if !ok {
					cm.Forbidden(c)
					return
				}
			}
			c.Next()
			return
		} else if options.ValidationRule == atLeastOne {
			for _, permission := range permissions {
				vals := append([]string{sub}, options.PermissionParser(permission)...)
				if ok, err := cm.Enforcer.Enforce(convertToInterface(vals)...); err != nil {
					c.SendStatus(fiber.StatusInternalServerError)
					return
				} else if ok {
					c.Next()
					return
				}
			}
			cm.Forbidden(c)
			return
		}

		c.Next()
	}
}

// RoutePermission tries to find the current subject and determine if the
// subject has the required permissions according to predefined Casbin policies.
// This method uses http Path and Method as object and action.
func (cm *PermissionMiddleware) RoutePermission() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		sub := cm.Lookup(c)
		if len(sub) == 0 {
			cm.Unauthorized(c)
			return
		}
		if ok, err := cm.Enforcer.Enforce(sub, c.Path(), c.Method()); err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return
		} else if !ok {
			cm.Forbidden(c)
			return
		}

		c.Next()
		return
	}
}

// RequiresRoles tries to find the current subject and determine if the
// subject has the required roles according to predefined Casbin policies.
func (cm *PermissionMiddleware) RequiresRoles(roles []string, opts ...func(o *Options)) func(*fiber.Ctx) {
	options := &Options{
		ValidationRule:   matchAll,
		PermissionParser: permissionParserWithSeperator(":"),
	}

	for _, o := range opts {
		o(options)
	}
	return func(c *fiber.Ctx) { //nolint:wsl
		if len(roles) == 0 {
			c.Next()
			return
		}

		sub := cm.Lookup(c)
		if len(sub) == 0 {
			cm.Unauthorized(c)
			return
		}

		userRoles, err := cm.Enforcer.GetRolesForUser(sub)
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return
		}

		if options.ValidationRule == matchAll {
			for _, role := range roles {
				if !contains(userRoles, role) {
					cm.Forbidden(c)
					return
				}
			}
			c.Next() //nolint:wsl
			return
		} else if options.ValidationRule == atLeastOne {
			for _, role := range roles {
				if contains(userRoles, role) {
					c.Next()
					return
				}
			}
			cm.Forbidden(c)
			return
		}

		c.Next()
	}
}

func contains(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false //nolint:wsl
}

func convertToInterface(arr []string) []interface{} {
	in := make([]interface{}, 0)
	for _, a := range arr {
		in = append(in, a)
	}
	return in //nolint:wsl
}
