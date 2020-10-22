/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Drug
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDisease,
    EntDiseaseFromJSON,
    EntDiseaseFromJSONTyped,
    EntDiseaseToJSON,
    EntDrugType,
    EntDrugTypeFromJSON,
    EntDrugTypeFromJSONTyped,
    EntDrugTypeToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugEdges
 */
export interface EntDrugEdges {
    /**
     * 
     * @type {EntDisease}
     * @memberof EntDrugEdges
     */
    disease?: EntDisease;
    /**
     * 
     * @type {EntDrugType}
     * @memberof EntDrugEdges
     */
    drugtype?: EntDrugType;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntDrugEdges
     */
    employee?: EntEmployee;
}

export function EntDrugEdgesFromJSON(json: any): EntDrugEdges {
    return EntDrugEdgesFromJSONTyped(json, false);
}

export function EntDrugEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'disease': !exists(json, 'Disease') ? undefined : EntDiseaseFromJSON(json['Disease']),
        'drugtype': !exists(json, 'Drugtype') ? undefined : EntDrugTypeFromJSON(json['Drugtype']),
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
    };
}

export function EntDrugEdgesToJSON(value?: EntDrugEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'disease': EntDiseaseToJSON(value.disease),
        'drugtype': EntDrugTypeToJSON(value.drugtype),
        'employee': EntEmployeeToJSON(value.employee),
    };
}


