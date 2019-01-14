package main
import (

  "fmt"
  "encoding/json"
)

func main() {

  var userslist interface{}

  // define JSON data
  var s_json = []byte(`{
    "ok": true,
    "members": [
        {
            "id": "U023BECGF",
            "team_id": "T021F9ZE2",
            "name": "bobby",
            "deleted": false,
            "status": null,
            "color": "9f69e7",
            "real_name": "Bobby Tables",
            "tz": "America\/Los_Angeles",
            "tz_label": "Pacific Daylight Time",
            "tz_offset": -25200,
            "profile": {
                "avatar_hash": "ge3b51ca72de",
                "first_name": "Bobby",
                "last_name": "Tables",
                "real_name": "Bobby Tables",
                "email": "bobby@slack.com",
                "skype": "my-skype-name",
                "phone": "+1 (123) 456 7890",
                "image_24": "https://...",
                "image_32": "https://...",
                "image_48": "https://...",
                "image_72": "https://...",
                "image_192": "https://..."
            },
            "is_admin": true,
            "is_owner": true,
            "has_2fa": false,
            "has_files": true
        },
        {
            "id": "U023BECGG",
            "team_id": "T021F9ZE2",
            "name": "Henry",
            "deleted": false,
            "status": null,
            "color": "9f69e7",
            "real_name": "Henry Ford",
            "tz": "America\/Los_Angeles",
            "tz_label": "Pacific Daylight Time",
            "tz_offset": -25200,
            "profile": {
                "avatar_hash": "ge3b51ca72de",
                "first_name": "Henry",
                "last_name": "Ford",
                "real_name": "Henry Ford",
                "email": "henry@slack.com",
                "skype": "my-skype-name",
                "phone": "+1 (123) 456 7890",
                "image_24": "https://...",
                "image_32": "https://...",
                "image_48": "https://...",
                "image_72": "https://...",
                "image_192": "https://..."
            },
            "is_admin": true,
            "is_owner": true,
            "has_2fa": false,
            "has_files": true
        }

        ]
    }`)

    err := json.Unmarshal(s_json, &userslist)
    if err != nil {
      panic(err)
    }

    fmt.Println("Name is ", userslist.(map[string]interface{})["members"].([]interface{})[1].(map[string]interface{})["name"].(string))

    fmt.Println("Email is ", userslist.(map[string]interface{})["members"].([]interface{})[1].(map[string]interface{})["profile"].(map[string]interface{})["email"].(string))

}
