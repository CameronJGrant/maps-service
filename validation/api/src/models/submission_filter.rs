/*
 * StrafesNET Submissions - OpenAPI 3.1
 *
 * Browse and manage map submissions in the staging pipeline.
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmissionFilter {
    #[serde(rename = "ID", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "DisplayName", skip_serializing_if = "Option::is_none")]
    pub display_name: Option<String>,
    #[serde(rename = "Creator", skip_serializing_if = "Option::is_none")]
    pub creator: Option<String>,
    #[serde(rename = "GameID", skip_serializing_if = "Option::is_none")]
    pub game_id: Option<i32>,
    #[serde(rename = "Date", skip_serializing_if = "Option::is_none")]
    pub date: Option<i64>,
}

impl SubmissionFilter {
    pub fn new() -> SubmissionFilter {
        SubmissionFilter {
            id: None,
            display_name: None,
            creator: None,
            game_id: None,
            date: None,
        }
    }
}

