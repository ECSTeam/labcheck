# labcheck

Simple app to track information on lab environment.  Integrated with Slack's slash command.

Usage:
`/labs`
Returns the all lab information


`/labs checkout <labname>
/labs checkin <labname>`
Checkout adds your slack/name and marks the lab 'Available=false' for when you need to work on a lab.
Checkin when you are done with the lab, clears your username and mark lab as Available=true.

`/labs status <labname>`
Returns information on a single lab

`/labs update <labname> {"version"="x.x", "desc"="..."}`
Updates lab version and description. Wrap version/desc in JSON format.
Use the description to add tiles, apps, services or whatever is relevant to the particular lab.

TODO:
Add and delete labs.  Search to be used for answering the question 'which lab has -blah-'
/labs add labxx {version='x.x', desc='....'}
/labs delete labxx
/labs search <searchterm>
