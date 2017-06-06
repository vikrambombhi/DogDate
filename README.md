# DogDate
DogDate is a dating app for dogs. The main use case for this app is to find/organize play dates for your dog. 

### Routes
| Route | Description | Response type |
| ------ | ------ | ------ |
| /login             | get JWT token for authentication         | String           |
| /matches/available | get dogs which avaiable for dates        | list of type Dog |
| /matches/matched   | get dogs that you have been paired with  | list of type Dog |
| /matches/history   | get dogs you have liked in the past      | list of type Dog |
| /matches/purposals | get dogs who have shown interest in you  | list of type Dog |

### Response types
##### Dog
```json
{
  "id": 2,
  "owner": 2,
  "name": "dasiy",
  "breed": "mix",
  "size": "medium"
}
 ```
 `note above response type is populated with sample data`


