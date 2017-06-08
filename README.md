# DogDate
DogDate is a dating app for dogs. The main use case for this app is to find/organize play dates for your dog. 

### Routes
| Route | Method | Description | Response type |
| ------ | ------ | ------ | ------ |
| /login             | GET | get JWT token for authentication         | String           |
| /matches           | GET | get dogs which avaiable for dates        | list of type Dog |
| /matches           | POST | like a dog(swipe right)                 | status String    |
| /matches/matched   | GET | get dogs that you have been paired with  | list of type Dog |
| /matches/history   | GET | get dogs you have liked in the past      | list of type Dog |
| /matches/purposals | GET | get dogs who have shown interest in you  | list of type Dog |
| /users/{userID}    | GET | get user info using the users ID         | JSON with User and list of Dogs |

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


