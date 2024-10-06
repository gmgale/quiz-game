/**
 * Quiz Game API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


export interface GameSession { 
    id?: string;
    code?: string;
    status?: GameSession.StatusEnum;
    currentQuestionId?: string;
    startTime?: string;
}
export namespace GameSession {
    export type StatusEnum = 'waiting' | 'active' | 'finished';
    export const StatusEnum = {
        Waiting: 'waiting' as StatusEnum,
        Active: 'active' as StatusEnum,
        Finished: 'finished' as StatusEnum
    };
}


