package schema

//
// This file is automatically generated by schema/generator
//
// **** DO NOT EDIT ****
//
const DiscoveryJSON = `{
  "kind": "discovery#restDescription",
  "discoveryVersion": "v1",
  "id": "bridge:v1",
  "name": "schema",
  "version": "v1",
  "title": "Bridge API",
  "description": "",
  "documentationLink": "http://github.com/coreos-inc/bridge",
  "protocol": "rest",
  "icons": {
    "x16": "",
    "x32": ""
  },
  "labels": [],
  "rootUrl": "http://localhost:9090",
  "servicePath": "/api/bridge/v1/",
  "batchPath": "batch",
  "parameters": {},
  "auth": {},
  "schemas": {
    "EtcdState": {
      "id": "EtcdState",
      "type": "object",
      "properties": {
        "checkSuccess": {
          "type": "boolean"
        },
        "members": {
          "type": "array",
          "items": {
            "$ref": "EtcdMember"
          }
        }
      }
    },
    "EtcdMember": {
      "id": "EtcdMember",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "ControlService": {
      "id": "ControlService",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "unitStates": {
          "type": "array",
          "items": {
            "$ref": "UnitState"
          }
        }
      }
    },
    "UnitState": {
      "id": "UnitState",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "hash": {
          "type": "string"
        },
        "machineID": {
          "type": "string"
        },
        "systemdLoadState": {
          "type": "string"
        },
        "systemdActiveState": {
          "type": "string"
        },
        "systemdSubState": {
          "type": "string"
        }
      }
    },
    "UnitStatePage": {
      "id": "UnitStatePage",
      "type": "object",
      "properties": {
        "states": {
          "type": "array",
          "items": {
            "$ref": "UnitState"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    }
  },
  "resources": {
    "cluster": {
      "methods": {
        "controlServices": {
          "id": "bridge.cluster.status.controlServices",
          "description": "Retrieve status of all control cluster services.",
          "httpMethod": "GET",
          "path": "cluster/status/control-services",
          "response": {
            "type": "array",
            "items": {
              "$ref": "ControlService"
            }
          }
        },
        "etcd": {
          "id": "bridge.cluster.status.etcd",
          "description": "Retrieve etcd machine statuses.",
          "httpMethod": "GET",
          "path": "cluster/status/etcd",
          "response": {
            "type": "array",
            "items": {
              "$ref": "EtcdMember"
            }
          }
        }
      }
    }
  }
}
`
