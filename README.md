# labcheck

Labcheck is a checkin/checkout system for keeping track of ECS lab environment.  Use it to quickly reference PCF versions, tiles, apps, etc.  Responses are *ephemeral* _only you can see the response_, and *in_channel* _response appears to everyone in the channel_

Usage:


`/labs`   <_*ephemeral* returns info for all labs_>

`/labs checkout _labname_ {"_optional comment_"}`  <_"in_channel" returns info on the lab you checked out with a comment.  The comment is not stored and only used for the Slack response_>

`/labs checkin <labname>`   <_"in_channel" returns a lab to available state_>

`/labs status <labname>` <_"ephemeral" Interested in a particular lab?  Use this._>

`/labs update <labname> {"version":"x.x", "desc":"_tiles, apps, etc._"}` <_"ephemeral" Update lab version and
    descripion ie. tiles, apps.  Use JSON notation for version and desc._>

 `/labs help`  <_"ephemeral" to see this description. Feel free to make comments, improvements, bugs on github._>`
TODO:
Add and delete labs.  Search to be used for answering the question 'which lab has -blah-'
/labs add labxx {"version":"x.x", "desc":"...."}
/labs delete labxx
/labs search <searchterm>
