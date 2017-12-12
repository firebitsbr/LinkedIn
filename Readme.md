| Endpoint       | Method           | Description  |    Response  |     Return  |
| ------------- |-------------| -----| -----| -----|
| /v1/login      | POST      |   do the login | 200 OK | JWT |
| /v1/register | POST      |    do the registration | 201 Created | - |
| /v1/logout | GET      |    log the user out | 200 OK | - |
| /v1/me      | GET | show the profile page of the current user| 200 OK | user object |
| /v1/users/{uid}      | GET | show the profile page of a user| 200 OK | user object |
| /v1/me/skills | POST      |    add a skill | 201 Created | - |
| /v1/me/skills/{sid} | DELETE      |    remove a skill | 204 No Content | - |
| /v1/users/{uid}/skills/{sid}/endorse | PUT      |    endorse a skill | 204 No Content | - |
| /v1/users/{uid}/skills/{sid}/undo | PUT     |    undo endorsing a skill | 204 No Content | - |
 /v1/skills      | GET | show all skill tags in descending order | 200 OK | skill object |
| /v1/skills/{sid}      | GET | show the profile page of a skill tag | 200 OK | skill object |