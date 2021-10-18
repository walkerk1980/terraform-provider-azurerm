package servicebus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// AccessRightsListen ...
	AccessRightsListen AccessRights = "Listen"
	// AccessRightsManage ...
	AccessRightsManage AccessRights = "Manage"
	// AccessRightsSend ...
	AccessRightsSend AccessRights = "Send"
)

// PossibleAccessRightsValues returns an array of possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{AccessRightsListen, AccessRightsManage, AccessRightsSend}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// DefaultActionAllow ...
	DefaultActionAllow DefaultAction = "Allow"
	// DefaultActionDeny ...
	DefaultActionDeny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{DefaultActionAllow, DefaultActionDeny}
}

// EndPointProvisioningState enumerates the values for end point provisioning state.
type EndPointProvisioningState string

const (
	// EndPointProvisioningStateCanceled ...
	EndPointProvisioningStateCanceled EndPointProvisioningState = "Canceled"
	// EndPointProvisioningStateCreating ...
	EndPointProvisioningStateCreating EndPointProvisioningState = "Creating"
	// EndPointProvisioningStateDeleting ...
	EndPointProvisioningStateDeleting EndPointProvisioningState = "Deleting"
	// EndPointProvisioningStateFailed ...
	EndPointProvisioningStateFailed EndPointProvisioningState = "Failed"
	// EndPointProvisioningStateSucceeded ...
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	// EndPointProvisioningStateUpdating ...
	EndPointProvisioningStateUpdating EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns an array of possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{EndPointProvisioningStateCanceled, EndPointProvisioningStateCreating, EndPointProvisioningStateDeleting, EndPointProvisioningStateFailed, EndPointProvisioningStateSucceeded, EndPointProvisioningStateUpdating}
}

// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// EntityStatusActive ...
	EntityStatusActive EntityStatus = "Active"
	// EntityStatusCreating ...
	EntityStatusCreating EntityStatus = "Creating"
	// EntityStatusDeleting ...
	EntityStatusDeleting EntityStatus = "Deleting"
	// EntityStatusDisabled ...
	EntityStatusDisabled EntityStatus = "Disabled"
	// EntityStatusReceiveDisabled ...
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	// EntityStatusRenaming ...
	EntityStatusRenaming EntityStatus = "Renaming"
	// EntityStatusRestoring ...
	EntityStatusRestoring EntityStatus = "Restoring"
	// EntityStatusSendDisabled ...
	EntityStatusSendDisabled EntityStatus = "SendDisabled"
	// EntityStatusUnknown ...
	EntityStatusUnknown EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns an array of possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{EntityStatusActive, EntityStatusCreating, EntityStatusDeleting, EntityStatusDisabled, EntityStatusReceiveDisabled, EntityStatusRenaming, EntityStatusRestoring, EntityStatusSendDisabled, EntityStatusUnknown}
}

// FilterType enumerates the values for filter type.
type FilterType string

const (
	// FilterTypeCorrelationFilter ...
	FilterTypeCorrelationFilter FilterType = "CorrelationFilter"
	// FilterTypeSQLFilter ...
	FilterTypeSQLFilter FilterType = "SqlFilter"
)

// PossibleFilterTypeValues returns an array of possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{FilterTypeCorrelationFilter, FilterTypeSQLFilter}
}

// KeySource enumerates the values for key source.
type KeySource string

const (
	// KeySourceMicrosoftKeyVault ...
	KeySourceMicrosoftKeyVault KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns an array of possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{KeySourceMicrosoftKeyVault}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// KeyTypePrimaryKey ...
	KeyTypePrimaryKey KeyType = "PrimaryKey"
	// KeyTypeSecondaryKey ...
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{KeyTypePrimaryKey, KeyTypeSecondaryKey}
}

// ManagedServiceIdentityType enumerates the values for managed service identity type.
type ManagedServiceIdentityType string

const (
	// ManagedServiceIdentityTypeNone ...
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	// ManagedServiceIdentityTypeSystemAssigned ...
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	// ManagedServiceIdentityTypeSystemAssignedUserAssigned ...
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	// ManagedServiceIdentityTypeUserAssigned ...
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns an array of possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{ManagedServiceIdentityTypeNone, ManagedServiceIdentityTypeSystemAssigned, ManagedServiceIdentityTypeSystemAssignedUserAssigned, ManagedServiceIdentityTypeUserAssigned}
}

// NetworkRuleIPAction enumerates the values for network rule ip action.
type NetworkRuleIPAction string

const (
	// NetworkRuleIPActionAllow ...
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns an array of possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{NetworkRuleIPActionAllow}
}

// PrivateLinkConnectionStatus enumerates the values for private link connection status.
type PrivateLinkConnectionStatus string

const (
	// PrivateLinkConnectionStatusApproved ...
	PrivateLinkConnectionStatusApproved PrivateLinkConnectionStatus = "Approved"
	// PrivateLinkConnectionStatusDisconnected ...
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	// PrivateLinkConnectionStatusPending ...
	PrivateLinkConnectionStatusPending PrivateLinkConnectionStatus = "Pending"
	// PrivateLinkConnectionStatusRejected ...
	PrivateLinkConnectionStatusRejected PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns an array of possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{PrivateLinkConnectionStatusApproved, PrivateLinkConnectionStatusDisconnected, PrivateLinkConnectionStatusPending, PrivateLinkConnectionStatusRejected}
}

// ProvisioningStateDR enumerates the values for provisioning state dr.
type ProvisioningStateDR string

const (
	// ProvisioningStateDRAccepted ...
	ProvisioningStateDRAccepted ProvisioningStateDR = "Accepted"
	// ProvisioningStateDRFailed ...
	ProvisioningStateDRFailed ProvisioningStateDR = "Failed"
	// ProvisioningStateDRSucceeded ...
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
)

// PossibleProvisioningStateDRValues returns an array of possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{ProvisioningStateDRAccepted, ProvisioningStateDRFailed, ProvisioningStateDRSucceeded}
}

// PublicNetworkAccessFlag enumerates the values for public network access flag.
type PublicNetworkAccessFlag string

const (
	// PublicNetworkAccessFlagDisabled ...
	PublicNetworkAccessFlagDisabled PublicNetworkAccessFlag = "Disabled"
	// PublicNetworkAccessFlagEnabled ...
	PublicNetworkAccessFlagEnabled PublicNetworkAccessFlag = "Enabled"
)

// PossiblePublicNetworkAccessFlagValues returns an array of possible values for the PublicNetworkAccessFlag const type.
func PossiblePublicNetworkAccessFlagValues() []PublicNetworkAccessFlag {
	return []PublicNetworkAccessFlag{PublicNetworkAccessFlagDisabled, PublicNetworkAccessFlagEnabled}
}

// RoleDisasterRecovery enumerates the values for role disaster recovery.
type RoleDisasterRecovery string

const (
	// RoleDisasterRecoveryPrimary ...
	RoleDisasterRecoveryPrimary RoleDisasterRecovery = "Primary"
	// RoleDisasterRecoveryPrimaryNotReplicating ...
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	// RoleDisasterRecoverySecondary ...
	RoleDisasterRecoverySecondary RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns an array of possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{RoleDisasterRecoveryPrimary, RoleDisasterRecoveryPrimaryNotReplicating, RoleDisasterRecoverySecondary}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// SkuNameBasic ...
	SkuNameBasic SkuName = "Basic"
	// SkuNamePremium ...
	SkuNamePremium SkuName = "Premium"
	// SkuNameStandard ...
	SkuNameStandard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{SkuNameBasic, SkuNamePremium, SkuNameStandard}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierPremium ...
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierPremium, SkuTierStandard}
}

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// UnavailableReasonInvalidName ...
	UnavailableReasonInvalidName UnavailableReason = "InvalidName"
	// UnavailableReasonNameInLockdown ...
	UnavailableReasonNameInLockdown UnavailableReason = "NameInLockdown"
	// UnavailableReasonNameInUse ...
	UnavailableReasonNameInUse UnavailableReason = "NameInUse"
	// UnavailableReasonNone ...
	UnavailableReasonNone UnavailableReason = "None"
	// UnavailableReasonSubscriptionIsDisabled ...
	UnavailableReasonSubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// UnavailableReasonTooManyNamespaceInCurrentSubscription ...
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns an array of possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{UnavailableReasonInvalidName, UnavailableReasonNameInLockdown, UnavailableReasonNameInUse, UnavailableReasonNone, UnavailableReasonSubscriptionIsDisabled, UnavailableReasonTooManyNamespaceInCurrentSubscription}
}