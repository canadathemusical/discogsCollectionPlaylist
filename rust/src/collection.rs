use serde_derive::Deserialize;
use serde_derive::Serialize;

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct CollectionReleasesResponse {
    pub pagination: Pagination,
    pub collection: Collection,
}

pub type Collection = Vec<Album>;

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Album {
    pub id: i64,
    #[serde(rename = "instance_id")]
    pub instance_id: i64,
    #[serde(rename = "date_added")]
    pub date_added: String,
    pub rating: i64,
    #[serde(rename = "basic_information")]
    pub basic_information: BasicInformation,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct BasicInformation {
    pub id: i64,
    #[serde(rename = "master_id")]
    pub master_id: i64,
    #[serde(rename = "master_url")]
    pub master_url: Option<String>,
    #[serde(rename = "resource_url")]
    pub resource_url: String,
    pub thumb: String,
    #[serde(rename = "cover_image")]
    pub cover_image: String,
    pub title: String,
    pub year: i64,
    pub formats: Vec<Format>,
    pub artists: Vec<Artist>,
    pub labels: Vec<Label>,
    pub genres: Vec<String>,
    pub styles: Vec<String>,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Format {
    pub name: String,
    pub qty: String,
    #[serde(default)]
    pub descriptions: Vec<String>,
    pub text: Option<String>,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Artist {
    pub name: String,
    pub anv: String,
    pub join: String,
    pub role: String,
    pub tracks: String,
    pub id: i64,
    #[serde(rename = "resource_url")]
    pub resource_url: String,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Label {
    pub name: String,
    pub catno: String,
    #[serde(rename = "entity_type")]
    pub entity_type: String,
    #[serde(rename = "entity_type_name")]
    pub entity_type_name: String,
    pub id: i64,
    #[serde(rename = "resource_url")]
    pub resource_url: String,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Pagination {
    pub page: i64,
    pub pages: i64,
    #[serde(rename = "per_page")]
    pub per_page: i64,
    pub items: i64,
    pub urls: Urls,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Urls {
    pub last: String,
    pub next: String,
}
