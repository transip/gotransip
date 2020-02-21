package vps

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VpsRepository VpsRepository service
type VpsRepository gotransip.Service

// AddIPv6AddressToAVPSOpts Optional parameters for the method 'AddIPv6AddressToAVPS'
type AddIPv6AddressToAVPSOpts struct {
	InlineObject42 optional.Interface
}

/*
AddIPv6AddressToAVPS Add IPv6 address to a VPS
TransIP VPSes are deployed with an &#x60;/64&#x60; IPv6 range. In order to set ReverseDNS for specific ipv6 addresses, you will have to add the IPv6 address via this command.  After adding an IPv6 address, you can set the reverse DNS for this address using the [Update Reverse DNS](#vps-ip-addresses-put) API call.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *AddIPv6AddressToAVPSOpts - Optional Parameters:
 * @param "InlineObject42" (optional.Interface of InlineObject42) -
*/
func (a *VpsRepository) AddIPv6AddressToAVPS(vpsName string, localVarOptionals *AddIPv6AddressToAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/ip-addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject42.IsSet() {
		localVarOptionalInlineObject42, localVarOptionalInlineObject42ok := localVarOptionals.InlineObject42.Value().(inline_objects.InlineObject42)
		if !localVarOptionalInlineObject42ok {
			return nil, reportError("inlineObject42 should be InlineObject42")
		}
		localVarPostBody = &localVarOptionalInlineObject42
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CancelAPrivateNetworkOpts Optional parameters for the method 'CancelAPrivateNetwork'
type CancelAPrivateNetworkOpts struct {
	InlineObject33 optional.Interface
}

/*
CancelAPrivateNetwork Cancel a private network
Cancel a private network.  You can set the &#x60;endTime&#x60; attribute to &#x60;end&#x60; or &#x60;immediately&#x60;, this has the following implications:  * **end**: The private network will be terminated from the end date of the agreement as can be found in the applicable quote;  * **immediately**: The private network will be terminated immediately.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param privateNetworkName Name of the private network
 * @param optional nil or *CancelAPrivateNetworkOpts - Optional Parameters:
 * @param "InlineObject33" (optional.Interface of InlineObject33) -
*/
func (a *VpsRepository) CancelAPrivateNetwork(privateNetworkName string, localVarOptionals *CancelAPrivateNetworkOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/private-networks/{privateNetworkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkName"+"}", _neturl.QueryEscape(parameterToString(privateNetworkName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject33.IsSet() {
		localVarOptionalInlineObject33, localVarOptionalInlineObject33ok := localVarOptionals.InlineObject33.Value().(inline_objects.InlineObject33)
		if !localVarOptionalInlineObject33ok {
			return nil, reportError("inlineObject33 should be InlineObject33")
		}
		localVarPostBody = &localVarOptionalInlineObject33
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CancelAVPSOpts Optional parameters for the method 'CancelVps'
type CancelAVPSOpts struct {
	InlineObject37 optional.Interface
}

/*
CancelVps Cancel a VPS
Using the DELETE method on a VPS will cancel the VPS, thus deleting it.  Upon cancellation This will wipe all data on the VPS and permanently destroy it.  You can set the &#x60;endTime&#x60; attribute to &#39;end&#39; or &#39;immediately&#39;, this has the following implications:  * **end**: The VPS will be terminated from the end date of the agreement as can be found in the applicable quote;  * **immediately**: The VPS will be terminated immediately.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *CancelAVPSOpts - Optional Parameters:
 * @param "InlineObject37" (optional.Interface of InlineObject37) -
*/
func (a *VpsRepository) CancelVps(vpsName string, localVarOptionals *CancelAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject37.IsSet() {
		localVarOptionalInlineObject37, localVarOptionalInlineObject37ok := localVarOptionals.InlineObject37.Value().(inline_objects.InlineObject37)
		if !localVarOptionalInlineObject37ok {
			return nil, reportError("inlineObject37 should be InlineObject37")
		}
		localVarPostBody = &localVarOptionalInlineObject37
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
CancelAnAddonForAVPS Cancel an addon for a VPS
By using this API call, you can cancel an add-on by name, specifying the VPS name as well. Due to technical restrictions (possible dataloss) storage add-ons cannot be cancelled.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param addonName Addon name
*/
func (a *VpsRepository) CancelAnAddonForAVPS(vpsName string, addonName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/addons/{addonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"addonName"+"}", _neturl.QueryEscape(parameterToString(addonName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CancelBigStorageOpts Optional parameters for the method 'CancelBigStorage'
type CancelBigStorageOpts struct {
	InlineObject2 optional.Interface
}

/*
CancelBigStorage Cancel big storage
Cancels a big storage for the specified ‘endTime’. You can set the &#x60;endTime&#x60; attribute to &#x60;end&#x60; or &#x60;immediately&#x60;, this has the following implications:  * **end**: The Big Storage will be terminated from the end date of the agreement as can be found in the applicable quote;  * **immediately**: The Big Storage will be terminated immediately.   Note that canceling a Big Storage will wipe all data stored on it as well as off-site back-ups as well if these are activated.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bigStorageName The name of the big storage
 * @param optional nil or *CancelBigStorageOpts - Optional Parameters:
 * @param "InlineObject2" (optional.Interface of InlineObject2) -
*/
func (a *VpsRepository) CancelBigStorage(bigStorageName string, localVarOptionals *CancelBigStorageOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages/{bigStorageName}"
	localVarPath = strings.Replace(localVarPath, "{"+"bigStorageName"+"}", _neturl.QueryEscape(parameterToString(bigStorageName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject2.IsSet() {
		localVarOptionalInlineObject2, localVarOptionalInlineObject2ok := localVarOptionals.InlineObject2.Value().(inline_objects.InlineObject2)
		if !localVarOptionalInlineObject2ok {
			return nil, reportError("inlineObject2 should be InlineObject2")
		}
		localVarPostBody = &localVarOptionalInlineObject2
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CloneAVPSOpts Optional parameters for the method 'CloneAVPS'
type CloneAVPSOpts struct {
	InlineObject35 optional.Interface
}

/*
CloneAVPS Clone a VPS
Use this API call in order to clone an existing VPS. There are a few things to take into account when you want to clone an existing VPS to a new VPS:  * If the original VPS (which you’re going to clone) is currently locked, the clone will fail;  * Cloned control panels can be used on the VPS, but as the IP address changes, this does require you to synchronise the new license on the new VPS (licenses are often IP-based);  * Possibly, your VPS has its network interface(s) configured using (a) static IP(‘s) rather than a dynamic allocation using DHCP. If this is the case, you have to configure the new IP(‘s) on the new VPS. Do note that this is not the case with our pre-installed control panel images;  * VPS add-ons such as Big Storage aren’t affected by cloning - these will stay attached to the original VPS and can’t be swapped automatically   ::: warning   &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: As cloning is a paid service, an invoice will be generated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CloneAVPSOpts - Optional Parameters:
 * @param "InlineObject35" (optional.Interface of InlineObject35) -
*/
func (a *VpsRepository) CloneAVPS(localVarOptionals *CloneAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject35.IsSet() {
		localVarOptionalInlineObject35, localVarOptionalInlineObject35ok := localVarOptionals.InlineObject35.Value().(inline_objects.InlineObject35)
		if !localVarOptionalInlineObject35ok {
			return nil, reportError("inlineObject35 should be InlineObject35")
		}
		localVarPostBody = &localVarOptionalInlineObject35
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// ConvertBackupToSnapshotOpts Optional parameters for the method 'ConvertBackupToSnapshot'
type ConvertBackupToSnapshotOpts struct {
	InlineObject40 optional.Interface
}

/*
ConvertBackupToSnapshot Convert backup to snapshot
With this API call you can convert a backup to a snapshot for the VPS.  In case the creation of this snapshot leads to exceeding the maximum allowed snapshots, the API call will return an error and the snapshot will not be created please order extra snapshots before proceeding.  To convert a backup to a VPS snapshot, send a PATCH request with the &#x60;action&#x60; attribute set to &#x60;convert&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param backupId Id of the backup
 * @param optional nil or *ConvertBackupToSnapshotOpts - Optional Parameters:
 * @param "InlineObject40" (optional.Interface of InlineObject40) -
*/
func (a *VpsRepository) ConvertBackupToSnapshot(vpsName string, backupId float32, localVarOptionals *ConvertBackupToSnapshotOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/backups/{backupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"backupId"+"}", _neturl.QueryEscape(parameterToString(backupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject40.IsSet() {
		localVarOptionalInlineObject40, localVarOptionalInlineObject40ok := localVarOptionals.InlineObject40.Value().(inline_objects.InlineObject40)
		if !localVarOptionalInlineObject40ok {
			return nil, reportError("inlineObject40 should be InlineObject40")
		}
		localVarPostBody = &localVarOptionalInlineObject40
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CreateAContactOpts Optional parameters for the method 'CreateAContact'
type CreateAContactOpts struct {
	InlineObject29 optional.Interface
}

/*
CreateAContact Create a contact
Create a monitoring contact in your TransIP account.  You can later use this contact by adding them to your [TCP Monitor](#vps-tcp-monitoring).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateAContactOpts - Optional Parameters:
 * @param "InlineObject29" (optional.Interface of InlineObject29) -
*/
func (a *VpsRepository) CreateAContact(localVarOptionals *CreateAContactOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/monitoring-contacts"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject29.IsSet() {
		localVarOptionalInlineObject29, localVarOptionalInlineObject29ok := localVarOptionals.InlineObject29.Value().(inline_objects.InlineObject29)
		if !localVarOptionalInlineObject29ok {
			return nil, reportError("inlineObject29 should be InlineObject29")
		}
		localVarPostBody = &localVarOptionalInlineObject29
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CreateATCPMonitorForAVPSOpts Optional parameters for the method 'CreateATCPMonitorForAVPS'
type CreateATCPMonitorForAVPSOpts struct {
	InlineObject47 optional.Interface
}

/*
CreateATCPMonitorForAVPS Create a TCP monitor for a VPS
Create a TCP monitor and specify which ports you would like to monitor.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *CreateATCPMonitorForAVPSOpts - Optional Parameters:
 * @param "InlineObject47" (optional.Interface of InlineObject47) -
*/
func (a *VpsRepository) CreateATCPMonitorForAVPS(vpsName string, localVarOptionals *CreateATCPMonitorForAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/tcp-monitors"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject47.IsSet() {
		localVarOptionalInlineObject47, localVarOptionalInlineObject47ok := localVarOptionals.InlineObject47.Value().(inline_objects.InlineObject47)
		if !localVarOptionalInlineObject47ok {
			return nil, reportError("inlineObject47 should be InlineObject47")
		}
		localVarPostBody = &localVarOptionalInlineObject47
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// CreateSnapshotOfAVPSOpts Optional parameters for the method 'CreateSnapshotOfAVPS'
type CreateSnapshotOfAVPSOpts struct {
	InlineObject45 optional.Interface
}

/*
CreateSnapshotOfAVPS Create snapshot of a VPS
With this API call you can create a snapshot of a VPS.  In case the creation of this snapshot leads to exceeding the maximum allowed snapshots, the API call will return an error and the snapshot will not be created - please order extra snapshots before proceeding.  Creating a snapshot allows for restoring it on another VPS using the [Revert snapshot to a VPS](#vps-snapshots-patch) given that its specifications equals or exceeds those of the snapshot&#39;s source VPS.  ::: warning  We strongly recommend shutting the VPS down before taking a snapshot in order to prevent data loss, etc.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *CreateSnapshotOfAVPSOpts - Optional Parameters:
 * @param "InlineObject45" (optional.Interface of InlineObject45) -
*/
func (a *VpsRepository) CreateSnapshotOfAVPS(vpsName string, localVarOptionals *CreateSnapshotOfAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject45.IsSet() {
		localVarOptionalInlineObject45, localVarOptionalInlineObject45ok := localVarOptionals.InlineObject45.Value().(inline_objects.InlineObject45)
		if !localVarOptionalInlineObject45ok {
			return nil, reportError("inlineObject45 should be InlineObject45")
		}
		localVarPostBody = &localVarOptionalInlineObject45
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
DeleteAContact Delete a contact
Permanently deletes a monitoring contact from your TransIP account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param contactId Id number of the contact
*/
func (a *VpsRepository) DeleteAContact(contactId float32) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/monitoring-contacts/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", _neturl.QueryEscape(parameterToString(contactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
DeleteASnapshot Delete a snapshot
Delete a VPS snapshot using this API call.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param snapshotName Name of the snapshot
*/
func (a *VpsRepository) DeleteASnapshot(vpsName string, snapshotName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/snapshots/{snapshotName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"snapshotName"+"}", _neturl.QueryEscape(parameterToString(snapshotName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
DeleteATCPMonitorForAVPS Delete a TCP monitor for a VPS
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param ipAddress IP Address that is monitored
*/
func (a *VpsRepository) DeleteATCPMonitorForAVPS(vpsName string, ipAddress string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/tcp-monitors/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.QueryEscape(parameterToString(ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// DetachVpsFromPrivateNetworkOpts Optional parameters for the method 'DetachVpsFromPrivateNetwork'
type DetachVpsFromPrivateNetworkOpts struct {
	InlineObject34 optional.Interface
}

/*
DetachVpsFromPrivateNetwork Detach vps from privateNetwork
Detach VPSes from the private network one at a time. Send a PATCH request with the &#x60;action&#x60; attribute set to &#x60;removevps&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param privateNetworkName Name of the private network
 * @param optional nil or *DetachVpsFromPrivateNetworkOpts - Optional Parameters:
 * @param "InlineObject34" (optional.Interface of InlineObject34) -
*/
func (a *VpsRepository) DetachVpsFromPrivateNetwork(privateNetworkName string, localVarOptionals *DetachVpsFromPrivateNetworkOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/private-networks/{privateNetworkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkName"+"}", _neturl.QueryEscape(parameterToString(privateNetworkName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject34.IsSet() {
		localVarOptionalInlineObject34, localVarOptionalInlineObject34ok := localVarOptionals.InlineObject34.Value().(inline_objects.InlineObject34)
		if !localVarOptionalInlineObject34ok {
			return nil, reportError("inlineObject34 should be InlineObject34")
		}
		localVarPostBody = &localVarOptionalInlineObject34
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
GetBigStorageByName Get big storage by name
Get information about a specific Big Storage and its current status. If the Big Storage is attached to a VPS, the output will contain the VPS name it’s attached to.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bigStorageName The name of the big storage
@return InlineResponse2003
*/
func (a *VpsRepository) GetBigStorageByName(bigStorageName string) (inline_objects.InlineResponse2003, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2003
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages/{bigStorageName}"
	localVarPath = strings.Replace(localVarPath, "{"+"bigStorageName"+"}", _neturl.QueryEscape(parameterToString(bigStorageName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2003
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetBigStorageUsageStatisticsOpts Optional parameters for the method 'GetBigStorageUsageStatistics'
type GetBigStorageUsageStatisticsOpts struct {
	InlineObject4 optional.Interface
}

/*
GetBigStorageUsageStatistics Get big storage usage statistics
Get the usage statistics for a big storage. You can specify a &#x60;dateTimeStart&#x60; and &#x60;dateTimeEnd&#x60; parameter in the UNIX timestamp format. When none given, traffic for the past 24 hours are returned. The maximum period is one month.  When the big storage is not attached to a vps, there are no usage statistics available. Therefore, the response returned will be a 406 exception. If the big storage is re-attached to another vps then the old statistics are no longer available.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bigStorageName The name of the big storage
 * @param optional nil or *GetBigStorageUsageStatisticsOpts - Optional Parameters:
 * @param "InlineObject4" (optional.Interface of InlineObject4) -
@return InlineResponse2005
*/
func (a *VpsRepository) GetBigStorageUsageStatistics(bigStorageName string, localVarOptionals *GetBigStorageUsageStatisticsOpts) (inline_objects.InlineResponse2005, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2005
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages/{bigStorageName}/usage"
	localVarPath = strings.Replace(localVarPath, "{"+"bigStorageName"+"}", _neturl.QueryEscape(parameterToString(bigStorageName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject4.IsSet() {
		localVarOptionalInlineObject4, localVarOptionalInlineObject4ok := localVarOptionals.InlineObject4.Value().(inline_objects.InlineObject4)
		if !localVarOptionalInlineObject4ok {
			return localVarReturnValue, nil, reportError("inlineObject4 should be InlineObject4")
		}
		localVarPostBody = &localVarOptionalInlineObject4
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2005
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetIPAddressInfoByAddress Get IP address info by address
Only return network information for the specified IP address.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param ipAddress The IP address of the VPS
@return InlineResponse2009
*/
func (a *VpsRepository) GetIPAddressInfoByAddress(vpsName string, ipAddress string) (inline_objects.InlineResponse2009, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2009
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/ip-addresses/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.QueryEscape(parameterToString(ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2009
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetPrivateNetworkByName Get private network by name
Gather detailed information about a private network. As one of the returned attributes includes an array of the VPSes it’s attached to, you can determine if the private network is already attached to a specific VPS and if not, you can attach it.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param privateNetworkName Name of the private network
@return InlineResponse20036
*/
func (a *VpsRepository) GetPrivateNetworkByName(privateNetworkName string) (inline_objects.InlineResponse20036, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20036
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/private-networks/{privateNetworkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkName"+"}", _neturl.QueryEscape(parameterToString(privateNetworkName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20036
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetSnapshotByName Get snapshot by name
Specifying the snapshot ID and the VPS name it’s associated with, allows for insight in snapshot details.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param snapshotName Name of the snapshot
@return InlineResponse20048
*/
func (a *VpsRepository) GetSnapshotByName(vpsName string, snapshotName string) (inline_objects.InlineResponse20048, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20048
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/snapshots/{snapshotName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"snapshotName"+"}", _neturl.QueryEscape(parameterToString(snapshotName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20048
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetTrafficInformationForAVPS Get traffic information for a VPS
Traffic information for a specific VPS can be retrieved using this API call. Statistics such as consumed bandwidth and network usage statistics are classified as traffic information.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20041
*/
func (a *VpsRepository) GetTrafficInformationForAVPS(vpsName string) (inline_objects.InlineResponse20041, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20041
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/traffic/{vpsName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20041
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetTrafficPoolInformation Get traffic pool information
All the traffic of your VPSes combined, overusage will also be billed based on this information.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return InlineResponse20041
*/
func (a *VpsRepository) GetTrafficPoolInformation(ctx _context.Context) (inline_objects.InlineResponse20041, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20041
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/traffic"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20041
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetUsageDataForAVPSOpts Optional parameters for the method 'GetUsageDataForAVPS'
type GetUsageDataForAVPSOpts struct {
	InlineObject50 optional.Interface
}

/*
GetUsageDataForAVPS Get usage data for a VPS
Use this API call to retrieve usage data for a specific VPS. Make sure to specify the &#x60;dateTimeStart&#x60; and &#x60;dateTimeEnd&#x60; parameters in UNIX timestamp format.  Please take the following into account:  * The &#x60;dateTimeStart&#x60; and &#x60;dateTimeEnd&#x60; parameters allow for gathering information about a specific time period, when not specified the output will contain data for the past 24 hours;  * The difference between &#x60;dateTimeStart&#x60; and &#x60;dateTimeEnd&#x60; parameters may not exceed one month.   For traffic-related information and statistics, use the [Get traffic information for a VPS](#vps-traffic-get) API call.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *GetUsageDataForAVPSOpts - Optional Parameters:
 * @param "InlineObject50" (optional.Interface of InlineObject50) -
@return InlineResponse20051
*/
func (a *VpsRepository) GetUsageDataForAVPS(vpsName string, localVarOptionals *GetUsageDataForAVPSOpts) (inline_objects.InlineResponse20051, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20051
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/usage"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject50.IsSet() {
		localVarOptionalInlineObject50, localVarOptionalInlineObject50ok := localVarOptionals.InlineObject50.Value().(inline_objects.InlineObject50)
		if !localVarOptionalInlineObject50ok {
			return localVarReturnValue, nil, reportError("inlineObject50 should be InlineObject50")
		}
		localVarPostBody = &localVarOptionalInlineObject50
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20051
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetVNCDataForAVPS Get VNC data for a VPS
Get the location, token and password in order to connect directly to the VNC console of your VPS.  Please note, you cannot directly connect to the proxy using a VNC client, use a client that supports websockets like https://github.com/novnc/noVNC  Use the information as URL query parameters, or use the URL in the url parameter directly. Then enter the password. See the link below for an example of how to provide URL query parameters to a hosted novnc instance.  By default novnc understands the following url parameters:  - **host** the host hosting the vnc proxy, this should be &#x60;vncproxy.transip.nl&#x60;  - **path** the path to request on the host, &#x60;websockify?token&#x3D;YOURTOKEN&#x60;  - **password** the vnc password  - **autoconnect** whether or not to start connecting once you loaded the page   An example of all parameters together in one url would be https://novnc.com/noVNC/vnc.html?host&#x3D;vncproxy.transip.nl&amp;path&#x3D;websockify?token&#x3D;esco024gzqwyeeb5nexayi2gve09paw9dytumyxqzurxj5t642o5p6myzisn5gch&amp;password&#x3D;fVpTyDrhMiuYBXxn&amp;autoconnect&#x3D;true   ::: warning   &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: We do recommend running novnc locally or hosting your own novnc page, this way you can make sure your token and password remain private.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20052
*/
func (a *VpsRepository) GetVNCDataForAVPS(vpsName string) (inline_objects.InlineResponse20052, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20052
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/vnc-data"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20052
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetVPSByName Get VPS by name
Get information on specific VPS by name.  **Note**: for &#x60;vpsName&#x60;, use the TransIP provided name (format: username-vpsXX).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20043
*/
func (a *VpsRepository) GetVPSByName(vpsName string) (inline_objects.InlineResponse20043, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20043
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20043
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// HandoverAVPSOpts Optional parameters for the method 'HandoverAVPS'
type HandoverAVPSOpts struct {
	InlineObject38 optional.Interface
}

/*
HandoverAVPS Handover a VPS
Handover a VPS to another TransIP Account. This call will initiate the handover process. the actual handover will be done when the target customer accepts the handover.  Note: the VPS will be shut down in order to handover.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *HandoverAVPSOpts - Optional Parameters:
 * @param "InlineObject38" (optional.Interface of InlineObject38) -
*/
func (a *VpsRepository) HandoverAVPS(vpsName string, localVarOptionals *HandoverAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject38.IsSet() {
		localVarOptionalInlineObject38, localVarOptionalInlineObject38ok := localVarOptionals.InlineObject38.Value().(inline_objects.InlineObject38)
		if !localVarOptionalInlineObject38ok {
			return nil, reportError("inlineObject38 should be InlineObject38")
		}
		localVarPostBody = &localVarOptionalInlineObject38
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// InstallAnOperatingSystemOnAVPSOpts Optional parameters for the method 'InstallAnOperatingSystemOnAVPS'
type InstallAnOperatingSystemOnAVPSOpts struct {
	InlineObject44 optional.Interface
}

/*
InstallAnOperatingSystemOnAVPS Install an operating system on a VPS
With this method you can install operating systems and preinstalled images on a VPS.  A relatively important aspect regarding this feature is the ability to specify if the installation should be unattended using the &#x60;base64InstallText&#x60; parameter, allowing for automatic deployment of operating systems.  ::: warning  &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: This could potentially create an invoice when a commercial operating system or a commercial control panel is chosen.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *InstallAnOperatingSystemOnAVPSOpts - Optional Parameters:
 * @param "InlineObject44" (optional.Interface of InlineObject44) -
*/
func (a *VpsRepository) InstallAnOperatingSystemOnAVPS(vpsName string, localVarOptionals *InstallAnOperatingSystemOnAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/operating-systems"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject44.IsSet() {
		localVarOptionalInlineObject44, localVarOptionalInlineObject44ok := localVarOptionals.InlineObject44.Value().(inline_objects.InlineObject44)
		if !localVarOptionalInlineObject44ok {
			return nil, reportError("inlineObject44 should be InlineObject44")
		}
		localVarPostBody = &localVarOptionalInlineObject44
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
ListAddonsForAVPS List addons for a VPS
This method will return all active, cancelable and available add-ons for a VPS.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20044
*/
func (a *VpsRepository) ListAddonsForAVPS(vpsName string) (inline_objects.InlineResponse20044, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20044
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/addons"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20044
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListAllBigStoragesOpts Optional parameters for the method 'ListAllBigStorages'
type ListAllBigStoragesOpts struct {
	Body optional.String
}

/*
ListAllBigStorages List all big storages
Returns an array of all Big Storages in the given account.  After all Big Storages have been returned as an array, you can extract it and use a specific Big Storage for the other API calls documented below.  Should you only want to get the big storages attached to a specific VPS, set the &#x60;vpsName&#x60; parameter and only big storages that are attached to the given vps will be shown like https://api.transip.nl/v6/big-storages?vpsName&#x3D;example-vps  ::: note  This method supports pagination, using this methods you can limit the amount of big storages returned by the api, which might be usefull if you expect a lot of response objects and you want to spread that over multiple requests. See the [documentation on pages](#header-pages) for more information on how to use this functionality.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ListAllBigStoragesOpts - Optional Parameters:
 * @param "Body" (optional.String) -  Filters on a given vps name.
@return InlineResponse2002
*/
func (a *VpsRepository) ListAllBigStorages(localVarOptionals *ListAllBigStoragesOpts) (inline_objects.InlineResponse2002, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2002
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarPostBody = localVarOptionals.Body.Value()
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2002
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListAllContacts List all contacts
Get a list of all monitoring contacts attached to your TransIP account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return InlineResponse20034
*/
func (a *VpsRepository) ListAllContacts(ctx _context.Context) (inline_objects.InlineResponse20034, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20034
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/monitoring-contacts"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20034
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListAllPrivateNetworksOpts Optional parameters for the method 'ListAllPrivateNetworks'
type ListAllPrivateNetworksOpts struct {
	Body optional.String
}

/*
ListAllPrivateNetworks List all private networks
List all private networks in your account.  Should you only want to get the private networks attached to a specific VPS, set the &#x60;vpsName&#x60; parameter and only attached private networks will be shown like https://api.transip.nl/v6/private-networks?vpsName&#x3D;example-vps  If this parameter is not set, all private networks will be listed along with the VPSes it’s attached to.  ::: note  This method supports pagination, using this methods you can limit the amount of private networks returned by the api, which might be usefull if you expect a lot of response objects and you want to spread that over multiple requests. See the [documentation on pages](#header-pages) for more information on how to use this functionality.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ListAllPrivateNetworksOpts - Optional Parameters:
 * @param "Body" (optional.String) -  Filter private networks by a given VPS
@return InlineResponse20035
*/
func (a *VpsRepository) ListAllPrivateNetworks(localVarOptionals *ListAllPrivateNetworksOpts) (inline_objects.InlineResponse20035, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20035
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/private-networks"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarPostBody = localVarOptionals.Body.Value()
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20035
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListAllTCPMonitorsForAVPS List all TCP monitors for a VPS
Get an overview of all existing monitors attached to your VPS.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20049
*/
func (a *VpsRepository) ListAllTCPMonitorsForAVPS(vpsName string) (inline_objects.InlineResponse20049, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20049
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/tcp-monitors"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20049
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListAllVPSsOpts Optional parameters for the method 'ListAllVPSs'
type ListAllVPSsOpts struct {
	Body optional.String
}

/*
ListAllVPSs List all VPSs
Returns a list of all VPSs in the TransIP account.  ::: note  This method supports pagination, using this methods you can limit the amount of VPSs returned by the api, which might be usefull if you expect a lot of response objects and you want to spread that over multiple requests. See the [documentation on pages](#header-pages) for more information on how to use this functionality.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ListAllVPSsOpts - Optional Parameters:
 * @param "Body" (optional.String) -  Tags to filter by, separated by a comma.
@return InlineResponse20042
*/
func (a *VpsRepository) ListAllVPSs(localVarOptionals *ListAllVPSsOpts) (inline_objects.InlineResponse20042, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20042
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarPostBody = localVarOptionals.Body.Value()
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20042
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListAvailableUpgradesForAVPS List available upgrades for a VPS
List all available product upgrades for a VPS.  Upgrades differentiate from add-ons in the sense that upgrades are VPS products like the &#x60;vps-bladevps-pro-x16&#x60; VPS product.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20050
*/
func (a *VpsRepository) ListAvailableUpgradesForAVPS(vpsName string) (inline_objects.InlineResponse20050, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20050
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/upgrades"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20050
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListBackupsForABigStorage List backups for a big storage
Using this API call, you are able to list all backups belonging to a specific big storage.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bigStorageName The name of the big storage
@return InlineResponse2004
*/
func (a *VpsRepository) ListBackupsForABigStorage(bigStorageName string) (inline_objects.InlineResponse2004, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2004
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages/{bigStorageName}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"bigStorageName"+"}", _neturl.QueryEscape(parameterToString(bigStorageName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2004
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListBackupsForAVPS List backups for a VPS
TransIP offers multiple backup types, every VPS has 4 hourly backups by default, weekly backups are available for a small fee. This API call returns backups for both types.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse2004
*/
func (a *VpsRepository) ListBackupsForAVPS(vpsName string) (inline_objects.InlineResponse2004, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2004
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2004
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListFirewallForAVPS List firewall for a VPS
The VPS firewall works as a whitelist stateful firewall for **incoming** traffic. Enable the Firewall to block everything, add rules to exclude certain traffic from being blocked.  To further filter traffic, IP&#39;s can be whitelisted per rule. when no whitelist has been given for a specific rule, all traffic is allowed to this port.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20045
*/
func (a *VpsRepository) ListFirewallForAVPS(vpsName string) (inline_objects.InlineResponse20045, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20045
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/firewall"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20045
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListIPAddressesForAVPS List IP addresses for a VPS
This API call will return all IPv4 and IPv6 addresses attached to the VPS including Relevant network information like the gateway and subnet mask.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse2008
*/
func (a *VpsRepository) ListIPAddressesForAVPS(vpsName string) (inline_objects.InlineResponse2008, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse2008
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/ip-addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse2008
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListInstallableOperatingSystemsForAVPS List installable operating systems for a VPS
TransIP offers a number of operating systems and preinstalled images ready to be installed on any VPS. Using this API call, you can get a list of operating systems and preinstalled images available.  Commercial operating systems (such as Windows Server editions) and images shipping a commercial control panel contain the ‘price’ attribute, showing the price per month charged on top of the VPS itself.  A list with operating systems can also be found on the TransIP website: [https://www.transip.nl/vps/](https://www.transip.nl/vps/)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20046
*/
func (a *VpsRepository) ListInstallableOperatingSystemsForAVPS(vpsName string) (inline_objects.InlineResponse20046, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20046
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/operating-systems"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20046
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListSnapshotsForAVPS List snapshots for a VPS
This method allows you to list all snapshots that are taken of your VPS main disk.  A snapshot status can have the following values: ‘active‘, ‘creating‘, ‘reverting‘, ‘deleting‘, ‘pendingDeletion‘, ‘syncing‘, ‘moving‘ when status is ‘active‘ you can perform actions on it.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
@return InlineResponse20047
*/
func (a *VpsRepository) ListSnapshotsForAVPS(vpsName string) (inline_objects.InlineResponse20047, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  inline_objects.InlineResponse20047
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v inline_objects.InlineResponse20047
			err = a.client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// OrderANewPrivateNetworkOpts Optional parameters for the method 'OrderANewPrivateNetwork'
type OrderANewPrivateNetworkOpts struct {
	InlineObject31 optional.Interface
}

/*
OrderANewPrivateNetwork Order a new private network
Order a new private network. After ordering a private network you’re able to attach it to a VPS t o make use of the private network.  ::: warning  &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: This API call will create an invoice!
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OrderANewPrivateNetworkOpts - Optional Parameters:
 * @param "InlineObject31" (optional.Interface of InlineObject31) -
*/
func (a *VpsRepository) OrderANewPrivateNetwork(localVarOptionals *OrderANewPrivateNetworkOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/private-networks"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject31.IsSet() {
		localVarOptionalInlineObject31, localVarOptionalInlineObject31ok := localVarOptionals.InlineObject31.Value().(inline_objects.InlineObject31)
		if !localVarOptionalInlineObject31ok {
			return nil, reportError("inlineObject31 should be InlineObject31")
		}
		localVarPostBody = &localVarOptionalInlineObject31
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// OrderAddonsForAVPSOpts Optional parameters for the method 'OrderAddonsForAVPS'
type OrderAddonsForAVPSOpts struct {
	InlineObject39 optional.Interface
}

/*
OrderAddonsForAVPS Order addons for a VPS
In order to extend a specific VPS with add-ons, use this API call.  The type of add-ons that can be ordered range from extra IP addresses to hardware add-ons such as an extra core or additional SSD disk space.  ::: warning  &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: This API call will create a new invoice for the specified add-on(s)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *OrderAddonsForAVPSOpts - Optional Parameters:
 * @param "InlineObject39" (optional.Interface of InlineObject39) -
*/
func (a *VpsRepository) OrderAddonsForAVPS(vpsName string, localVarOptionals *OrderAddonsForAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/addons"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject39.IsSet() {
		localVarOptionalInlineObject39, localVarOptionalInlineObject39ok := localVarOptionals.InlineObject39.Value().(inline_objects.InlineObject39)
		if !localVarOptionalInlineObject39ok {
			return nil, reportError("inlineObject39 should be InlineObject39")
		}
		localVarPostBody = &localVarOptionalInlineObject39
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
RegenerateVNCTokenForAVps Regenerate VNC token for a vps
call this method to regenerate the VNC credentials for a VPS.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
*/
func (a *VpsRepository) RegenerateVNCTokenForAVps(vpsName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/vnc-data"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
RemoveAnIPv6AddressFromAVPS Remove an IPv6 address from a VPS
This method allows you to remove specific IPv6 addresses from the registered list of IPv6 addresses within the VPS&#39;s &#x60;/64&#x60; IPv6 range.  Note that deleting an IP address will also wipe its reverse DNS information.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param ipAddress The IP address of the VPS
*/
func (a *VpsRepository) RemoveAnIPv6AddressFromAVPS(vpsName string, ipAddress string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/ip-addresses/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.QueryEscape(parameterToString(ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// RevertABigStorageBackupOpts Optional parameters for the method 'RevertABigStorageBackup'
type RevertABigStorageBackupOpts struct {
	InlineObject3 optional.Interface
}

/*
RevertABigStorageBackup Revert a big storage backup
To revert a backup from a big storage, retrieve the &#x60;backupId&#x60; from the [backups](#vps-backups-get-1) resource. Please note this is only possible when any backups are created with the off-site backups feature, otherwise no backups will be made nor listed.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bigStorageName The name of the big storage
 * @param backupId Id of the backup
 * @param optional nil or *RevertABigStorageBackupOpts - Optional Parameters:
 * @param "InlineObject3" (optional.Interface of InlineObject3) -
*/
func (a *VpsRepository) RevertABigStorageBackup(bigStorageName string, backupId float32, localVarOptionals *RevertABigStorageBackupOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages/{bigStorageName}/backups/{backupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bigStorageName"+"}", _neturl.QueryEscape(parameterToString(bigStorageName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"backupId"+"}", _neturl.QueryEscape(parameterToString(backupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject3.IsSet() {
		localVarOptionalInlineObject3, localVarOptionalInlineObject3ok := localVarOptionals.InlineObject3.Value().(inline_objects.InlineObject3)
		if !localVarOptionalInlineObject3ok {
			return nil, reportError("inlineObject3 should be InlineObject3")
		}
		localVarPostBody = &localVarOptionalInlineObject3
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// RevertSnapshotToAVPSOpts Optional parameters for the method 'RevertSnapshotToAVPS'
type RevertSnapshotToAVPSOpts struct {
	InlineObject46 optional.Interface
}

/*
RevertSnapshotToAVPS Revert snapshot to a VPS
This method can be used to revert a snapshot to a VPS. Specifying the &#x60;destinationVpsName&#x60; attribute makes sure the snapshot is restored onto another VPS.  Networking may be configured statically on the source VPS, therefore breaking connectivity when restored onto another VPS with a new assigned IP. You should be able to alter this through the KVM console in the TransIP control panel in case SSH is unavailable.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param snapshotName Name of the snapshot
 * @param optional nil or *RevertSnapshotToAVPSOpts - Optional Parameters:
 * @param "InlineObject46" (optional.Interface of InlineObject46) -
*/
func (a *VpsRepository) RevertSnapshotToAVPS(vpsName string, snapshotName string, localVarOptionals *RevertSnapshotToAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/snapshots/{snapshotName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"snapshotName"+"}", _neturl.QueryEscape(parameterToString(snapshotName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject46.IsSet() {
		localVarOptionalInlineObject46, localVarOptionalInlineObject46ok := localVarOptionals.InlineObject46.Value().(inline_objects.InlineObject46)
		if !localVarOptionalInlineObject46ok {
			return nil, reportError("inlineObject46 should be InlineObject46")
		}
		localVarPostBody = &localVarOptionalInlineObject46
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdateAContactOpts Optional parameters for the method 'UpdateAContact'
type UpdateAContactOpts struct {
	InlineObject30 optional.Interface
}

/*
UpdateAContact Update a contact
Updates a specified contact. This call will override existing fields.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param contactId Id number of the contact
 * @param optional nil or *UpdateAContactOpts - Optional Parameters:
 * @param "InlineObject30" (optional.Interface of InlineObject30) -
*/
func (a *VpsRepository) UpdateAContact(contactId float32, localVarOptionals *UpdateAContactOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/monitoring-contacts/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", _neturl.QueryEscape(parameterToString(contactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject30.IsSet() {
		localVarOptionalInlineObject30, localVarOptionalInlineObject30ok := localVarOptionals.InlineObject30.Value().(inline_objects.InlineObject30)
		if !localVarOptionalInlineObject30ok {
			return nil, reportError("inlineObject30 should be InlineObject30")
		}
		localVarPostBody = &localVarOptionalInlineObject30
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdateATCPMonitorForAVPSOpts Optional parameters for the method 'UpdateATCPMonitorForAVPS'
type UpdateATCPMonitorForAVPSOpts struct {
	InlineObject48 optional.Interface
}

/*
UpdateATCPMonitorForAVPS Update a TCP monitor for a VPS
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param ipAddress IP Address that is monitored
 * @param optional nil or *UpdateATCPMonitorForAVPSOpts - Optional Parameters:
 * @param "InlineObject48" (optional.Interface of InlineObject48) -
*/
func (a *VpsRepository) UpdateATCPMonitorForAVPS(vpsName string, ipAddress string, localVarOptionals *UpdateATCPMonitorForAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/tcp-monitors/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.QueryEscape(parameterToString(ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject48.IsSet() {
		localVarOptionalInlineObject48, localVarOptionalInlineObject48ok := localVarOptionals.InlineObject48.Value().(inline_objects.InlineObject48)
		if !localVarOptionalInlineObject48ok {
			return nil, reportError("inlineObject48 should be InlineObject48")
		}
		localVarPostBody = &localVarOptionalInlineObject48
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdateAVPSOpts Optional parameters for the method 'UpdateAVPS'
type UpdateAVPSOpts struct {
	InlineObject36 optional.Interface
}

/*
UpdateAVPS Update a VPS
In this API call you can lock/unlock a VPS, update VPS description, and add/remove tags.  ### Locking a VPS  Locking a VPS prevents accidental execution of API calls and manual actions through the control panel.  For locking the VPS, set &#x60;isCustomerLocked&#x60; to &#x60;true&#x60;. Set the value to &#x60;false&#x60; for unlocking the VPS.  ### Change a VPS description  You can change your VPS description by simply changing the &#x60;description&#x60; attribute. Note that the identifier key &#x60;name&#x60; will not be changed. The description can be maximum 32 character longs  ### VPS Tags  To add/remove tags, you must update the &#x60;tags&#x60; attribute. Every time you make a call with a changed tags attribute, the existing tags are overridden.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *UpdateAVPSOpts - Optional Parameters:
 * @param "InlineObject36" (optional.Interface of InlineObject36) -
*/
func (a *VpsRepository) UpdateAVPS(vpsName string, localVarOptionals *UpdateAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject36.IsSet() {
		localVarOptionalInlineObject36, localVarOptionalInlineObject36ok := localVarOptionals.InlineObject36.Value().(inline_objects.InlineObject36)
		if !localVarOptionalInlineObject36ok {
			return nil, reportError("inlineObject36 should be InlineObject36")
		}
		localVarPostBody = &localVarOptionalInlineObject36
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdateBigStorageOpts Optional parameters for the method 'UpdateBigStorage'
type UpdateBigStorageOpts struct {
	InlineObject1 optional.Interface
}

/*
UpdateBigStorage Update big storage
This API calls allows for altering a big storage in several ways outlined below:  * Changing the description of a Big Storage;  * One Big Storages can only be attached to one VPS at a time;  * One VPS can have a maximum of 10 bigstorages attached;  * Set the &#x60;vpsName&#x60; property to the VPS name to attach to for attaching Big Storage;  * Set the &#x60;vpsName&#x60; property to null to detach the Big Storage from the currently attached VPS.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bigStorageName The name of the big storage
 * @param optional nil or *UpdateBigStorageOpts - Optional Parameters:
 * @param "InlineObject1" (optional.Interface of InlineObject1) -
*/
func (a *VpsRepository) UpdateBigStorage(bigStorageName string, localVarOptionals *UpdateBigStorageOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages/{bigStorageName}"
	localVarPath = strings.Replace(localVarPath, "{"+"bigStorageName"+"}", _neturl.QueryEscape(parameterToString(bigStorageName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject1.IsSet() {
		localVarOptionalInlineObject1, localVarOptionalInlineObject1ok := localVarOptionals.InlineObject1.Value().(inline_objects.InlineObject1)
		if !localVarOptionalInlineObject1ok {
			return nil, reportError("inlineObject1 should be InlineObject1")
		}
		localVarPostBody = &localVarOptionalInlineObject1
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdateFirewallForAVPSOpts Optional parameters for the method 'UpdateFirewallForAVPS'
type UpdateFirewallForAVPSOpts struct {
	InlineObject41 optional.Interface
}

/*
UpdateFirewallForAVPS Update firewall for a VPS
Update the ruleset for a VPS. This will override all current rules set for this VPS.  The VPS Firewall works as a whitelist. when no entries are given, but the firewall is enabled. All **incoming** traffic will be blocked.  IP&#39;s or IP Ranges (v4/v6) can be whitelisted per rule. When no whitelist has been given for a specific rule, all incoming traffic is allowed to this port.  Protocol parameter can either be &#x60;tcp&#x60;, &#x60;udp&#x60; or &#x60;tcp_udp&#x60;.  There is a maximum of 50 rules with each a maximum of 20 whitelist entries.  Any change to the firewall will temporary lock the VPS while the new rules are being applied.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *UpdateFirewallForAVPSOpts - Optional Parameters:
 * @param "InlineObject41" (optional.Interface of InlineObject41) -
*/
func (a *VpsRepository) UpdateFirewallForAVPS(vpsName string, localVarOptionals *UpdateFirewallForAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/firewall"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject41.IsSet() {
		localVarOptionalInlineObject41, localVarOptionalInlineObject41ok := localVarOptionals.InlineObject41.Value().(inline_objects.InlineObject41)
		if !localVarOptionalInlineObject41ok {
			return nil, reportError("inlineObject41 should be InlineObject41")
		}
		localVarPostBody = &localVarOptionalInlineObject41
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdatePrivateNetworkOpts Optional parameters for the method 'UpdatePrivateNetwork'
type UpdatePrivateNetworkOpts struct {
	InlineObject32 optional.Interface
}

/*
UpdatePrivateNetwork Update private network
This method can also be used to change the &#x60;description&#x60; attribute. The description can be maximum 32 character longs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param privateNetworkName Name of the private network
 * @param optional nil or *UpdatePrivateNetworkOpts - Optional Parameters:
 * @param "InlineObject32" (optional.Interface of InlineObject32) -
*/
func (a *VpsRepository) UpdatePrivateNetwork(privateNetworkName string, localVarOptionals *privatenetwork.PrivateNetwork) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/private-networks/{privateNetworkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkName"+"}", _neturl.QueryEscape(parameterToString(privateNetworkName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject32.IsSet() {
		localVarOptionalInlineObject32, localVarOptionalInlineObject32ok := localVarOptionals.InlineObject32.Value().(inline_objects.InlineObject32)
		if !localVarOptionalInlineObject32ok {
			return nil, reportError("inlineObject32 should be InlineObject32")
		}
		localVarPostBody = &localVarOptionalInlineObject32
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpdateReverseDNSForAVPSOpts Optional parameters for the method 'UpdateReverseDNSForAVPS'
type UpdateReverseDNSForAVPSOpts struct {
	InlineObject43 optional.Interface
}

/*
UpdateReverseDNSForAVPS Update reverse DNS for a VPS
Reverse DNS for IPv4 addresses as well as IPv6 addresses can be updated using this API call.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param ipAddress The IP address of the VPS
 * @param optional nil or *UpdateReverseDNSForAVPSOpts - Optional Parameters:
 * @param "InlineObject43" (optional.Interface of InlineObject43) -
*/
func (a *VpsRepository) UpdateReverseDNSForAVPS(vpsName string, ipAddress string, localVarOptionals *UpdateReverseDNSForAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/ip-addresses/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.QueryEscape(parameterToString(ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject43.IsSet() {
		localVarOptionalInlineObject43, localVarOptionalInlineObject43ok := localVarOptionals.InlineObject43.Value().(inline_objects.InlineObject43)
		if !localVarOptionalInlineObject43ok {
			return nil, reportError("inlineObject43 should be InlineObject43")
		}
		localVarPostBody = &localVarOptionalInlineObject43
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpgradeAVPSOpts Optional parameters for the method 'UpgradeAVPS'
type UpgradeAVPSOpts struct {
	InlineObject49 optional.Interface
}

/*
UpgradeAVPS Upgrade a VPS
This API call allows you too upgrade a VPS by name and productName.  It’s not possible to downgrade a VPS, as most upgrades cannot be deallocated due to technical reasons (data loss when shrinking the disk space).  ::: warning  &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: This API call will create an invoice for the upgrade.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vpsName VPS name
 * @param optional nil or *UpgradeAVPSOpts - Optional Parameters:
 * @param "InlineObject49" (optional.Interface of InlineObject49) -
*/
func (a *VpsRepository) UpgradeAVPS(vpsName string, localVarOptionals *UpgradeAVPSOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/vps/{vpsName}/upgrades"
	localVarPath = strings.Replace(localVarPath, "{"+"vpsName"+"}", _neturl.QueryEscape(parameterToString(vpsName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject49.IsSet() {
		localVarOptionalInlineObject49, localVarOptionalInlineObject49ok := localVarOptionals.InlineObject49.Value().(inline_objects.InlineObject49)
		if !localVarOptionalInlineObject49ok {
			return nil, reportError("inlineObject49 should be InlineObject49")
		}
		localVarPostBody = &localVarOptionalInlineObject49
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UpgradeBigStorageOpts Optional parameters for the method 'UpgradeBigStorage'
type UpgradeBigStorageOpts struct {
	InlineObject optional.Interface
}

/*
UpgradeBigStorage Upgrade big storage
With this method you are able to upgrade a bigstorage diskSize or enable backups.  The minimum size is 2 TB and storage can be extended with up to maximum of 40 TB. Make sure to use a multitude of 2 TB.  Optionally, to create back-ups of your Big Storage, enable off-site back-ups. We highly recommend activating back-ups.  ::: warning  &lt;i class&#x3D;\&quot;fa fa-warning\&quot;&gt;&lt;/i&gt; **Warning**: This API call will create an invoice!
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *UpgradeBigStorageOpts - Optional Parameters:
 * @param "InlineObject" (optional.Interface of InlineObject) -
*/
func (a *VpsRepository) UpgradeBigStorage(localVarOptionals *UpgradeBigStorageOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/big-storages"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.InlineObject.IsSet() {
		localVarOptionalInlineObject, localVarOptionalInlineObjectok := localVarOptionals.InlineObject.Value().(inline_objects.InlineObject)
		if !localVarOptionalInlineObjectok {
			return nil, reportError("inlineObject should be InlineObject")
		}
		localVarPostBody = &localVarOptionalInlineObject
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := gotransip.GenericError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
