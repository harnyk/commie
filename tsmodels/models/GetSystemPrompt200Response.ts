/* tslint:disable */
/* eslint-disable */
/**
 * Commie Plugin API
 * API for communication between Commie and plugins via CPCP (Commie Plugin Connectivity Protocol).
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface GetSystemPrompt200Response
 */
export interface GetSystemPrompt200Response {
    /**
     * Optional system prompt provided by the plugin
     * @type {string}
     * @memberof GetSystemPrompt200Response
     */
    systemPrompt?: string | null;
}

/**
 * Check if a given object implements the GetSystemPrompt200Response interface.
 */
export function instanceOfGetSystemPrompt200Response(value: object): value is GetSystemPrompt200Response {
    return true;
}

export function GetSystemPrompt200ResponseFromJSON(json: any): GetSystemPrompt200Response {
    return GetSystemPrompt200ResponseFromJSONTyped(json, false);
}

export function GetSystemPrompt200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetSystemPrompt200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'systemPrompt': json['system_prompt'] == null ? undefined : json['system_prompt'],
    };
}

export function GetSystemPrompt200ResponseToJSON(json: any): GetSystemPrompt200Response {
    return GetSystemPrompt200ResponseToJSONTyped(json, false);
}

export function GetSystemPrompt200ResponseToJSONTyped(value?: GetSystemPrompt200Response | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'system_prompt': value['systemPrompt'],
    };
}

