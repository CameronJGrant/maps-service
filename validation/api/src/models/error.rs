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

/// Error : Represents error object
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Error {
    #[serde(rename = "code")]
    pub code: i64,
    #[serde(rename = "message")]
    pub message: String,
}

impl Error {
    /// Represents error object
    pub fn new(code: i64, message: String) -> Error {
        Error {
            code,
            message,
        }
    }
}

