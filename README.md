# labcheck

Labcheck is a checkin/checkout system for keeping track of ECS lab environment.  Use it to quickly reference PCF versions, tiles, apps, etc.  <br>
Slack response modes are:
   * **ephemeral** - only you can see the response
   * **in_channel** - response appears to everyone in the channel

| Command | Description | Slack Mode |
|---|---|---|
|`/labs`| Returns info for all labs | ephemeral |
|`/labs checkout labxx {"_optional comment_"}`|Returns info on the lab you checked out with a comment.  The comment is not stored and only used for the Slack response | in_channel |
|`/labs checkin labxx`| Returns a lab to available state | in_channel |
|`/labs status labxx`|  Interested in a particular lab?  Use this.|ephemeral|
|`/labs update labxx {"version":"x.x", "desc":"..."}`|  Update lab version and description ie. tiles, apps.  Use JSON notation for version and desc.|ephemeral|
|`/labs help`| Link to this page on github.| ephemeral |

#TODOs: <br>
 Search to be used for answering questions like 'which labs have the mysql tile installed?' <br>
`/labs search <searchterm>`

Add and delete labs.
