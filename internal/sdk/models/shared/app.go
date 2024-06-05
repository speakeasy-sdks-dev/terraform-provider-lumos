// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type App struct {
	// The ID of this app.
	ID string `json:"id"`
	// The ID of the service associated with this app.
	AppClassID string `json:"app_class_id"`
	// The ID of the instance associated with this app.
	InstanceID string `json:"instance_id"`
	// The user-friendly label of this app.
	UserFriendlyLabel string `json:"user_friendly_label"`
	// The status of this app.
	Status DomainAppStatus `json:"status"`
	// The sources of this app.
	Sources []DiscoverySource `json:"sources"`
	// Whether the app is configured to allow multiple permissions to be requested at a time. This field will be removed in subsequent API versions.
	AllowMultiplePermissionSelection bool `json:"allow_multiple_permission_selection"`
}

func (o *App) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *App) GetAppClassID() string {
	if o == nil {
		return ""
	}
	return o.AppClassID
}

func (o *App) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

func (o *App) GetUserFriendlyLabel() string {
	if o == nil {
		return ""
	}
	return o.UserFriendlyLabel
}

func (o *App) GetStatus() DomainAppStatus {
	if o == nil {
		return DomainAppStatus("")
	}
	return o.Status
}

func (o *App) GetSources() []DiscoverySource {
	if o == nil {
		return []DiscoverySource{}
	}
	return o.Sources
}

func (o *App) GetAllowMultiplePermissionSelection() bool {
	if o == nil {
		return false
	}
	return o.AllowMultiplePermissionSelection
}
