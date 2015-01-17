# zendesk

[![Build Status](https://snap-ci.com/savaki/zendesk/branch/master/build_image)](https://snap-ci.com/savaki/zendesk/branch/master)

Work in progress implementation of the zendesk api in Go.

Starting with the user api and working from there.

# User Api

Request | Path | Status
-------- | ------ | :---------:
List Users | GET /api/v2/users.json | done
 | GET /api/v2/groups/{id}/users.json | -
 | GET /api/v2/organizations/{id}/users.json | -
Show User | GET /api/v2/users/{id}.json | done
Show Many Users |  GET /api/v2/users/show_many.json?ids={ids} | -
User Related Information | GET /api/v2/users/{id}/related.json | done
Create User | POST /api/v2/users.json | done
Merge Self With Another User | PUT /api/v2/users/me/merge.json | -
Merge Users | PUT /api/v2/users/{user_id}/merge.json | -
Create Many Users | POST /api/v2/users/create_many.json | -
Update User | PUT /api/v2/users/{id}.json | -
Suspend User | PUT /api/v2/users/{id}.json | -
Delete User | DELETE /api/v2/users/{id}.json | done
Search Users | GET /api/v2/users/search.json?query={query} | done
 | GET /api/v2/users/search.json?external_id={external_id} | -
Autocomplete Users | GET /api/v2/users/autocomplete.json?name={name} | done
Update Profile Image | PUT /api/v2/users/{id}.json | -
Show Current User | GET /api/v2/users/me.json | done
Set User's Password | POST /api/v2/users/{user_id}/password.json | done 
Change Your Password | PUT /api/v2/users/{user_id}/password.json | done

 
