# labcheck

Labcheck is a checkin/checkout system integrated with Slack for keeping track of the ECS lab environment.  Use it to quickly reference PCF versions, tiles, apps, etc. and to see if anyone is working in a particular lab.  It uses custom slash commands in Slack to call a service which stores data related to the labs.  The responses can either be seen by everyone in a channel or just by you, see response modes.     <br>
Slack response modes are:
   * **ephemeral** - only you can see the response
   * **in_channel** - response appears to everyone in the channel

# NOTE: _Turn off Smart Quotes on MAC OS, or the JSON won't parse from the Slack Desktop app_

| Command | Description | Slack Mode |
|---|---|---|
|`/labs`| Returns info for all labs | ephemeral |
|`/labs checkout labxx {"_optional comment_"}`|Returns info on the lab you checked out with a comment.  The comment is not stored and only used for the Slack response | in_channel |
|`/labs checkin labxx`| Returns a lab to available state | in_channel |
|`/labs status labxx`|  Interested in a particular lab?  Use this.|ephemeral|
|`/labs update labxx {"version":"x.x", "desc":"..."}`|  Update lab version and description ie. tiles, apps.  Use JSON notation for version and desc. Note: updates will overwrite existing content. |ephemeral|
|`/labs help`| Link to this page on github.| ephemeral |

#TODOs: <br>
 Search to be used for answering questions like 'which labs have the mysql tile installed?' <br>
`/labs search <searchterm>`

Add and delete labs.

# Deployment
Labcheck is currently deployed to `lab02 ecsteam | development` and uses google's Datastore cloud service to store it's data.
