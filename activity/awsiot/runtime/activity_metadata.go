package awsiot

var jsonMetadata = `{
  "name": "tibco-awsiot",
  "version": "0.0.1",
  "title": "Update AWS Device Shadow",
  "description": "Simple AWS IoT",
  "inputs":[
    {
      "name": "thingName",
      "type": "string",
      "required": true
    },
    {
      "name": "awsEndpoint",
      "type": "string",
      "required": true
    },
    {
      "name": "desired",
      "type": "params"
    },
    {
      "name": "reported",
      "type": "params"
    }
  ]
}
`
