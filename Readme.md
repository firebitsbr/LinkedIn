| URL       | Method           | Description  |
| ------------- |-------------| -----|
| /v1/login/ | GET      |    show the login page |
| /v1/login/      | POST      |   do the login |
| /v1/register/ | POST      |    do the registration |
| /v1/register/ | GET      |    show the register page |
| /v1/logout/ | POST      |    log the user out |
| /v1/me      | GET | show the profile page of the current user|
| /v1/users/{user_id}      | GET | show the profile page of a user|
| /v1/users/{user_id}/skills/ | POST      |    add a skill |
| /v1/users/{user_id}/skills/{skill_id} | DELETE      |    delete a skill |
| /v1/users/{user_id}/skills/{skill_id}/endorser/ | PUT      |    endorse a skill |