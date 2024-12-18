//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Elaugaste/contrail-go-api"
)

const (
	route_table_routes uint64 = 1 << iota
	route_table_id_perms
	route_table_display_name
	route_table_virtual_network_back_refs
)

type RouteTable struct {
        contrail.ObjectBase
	routes RouteTableType
	id_perms IdPermsType
	display_name string
	virtual_network_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *RouteTable) GetType() string {
        return "route-table"
}

func (obj *RouteTable) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *RouteTable) GetDefaultParentType() string {
        return "project"
}

func (obj *RouteTable) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *RouteTable) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *RouteTable) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *RouteTable) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *RouteTable) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *RouteTable) GetRoutes() RouteTableType {
        return obj.routes
}

func (obj *RouteTable) SetRoutes(value *RouteTableType) {
        obj.routes = *value
        obj.modified |= route_table_routes
}

func (obj *RouteTable) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *RouteTable) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= route_table_id_perms
}

func (obj *RouteTable) GetDisplayName() string {
        return obj.display_name
}

func (obj *RouteTable) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= route_table_display_name
}

func (obj *RouteTable) readVirtualNetworkBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & route_table_virtual_network_back_refs == 0) {
                err := obj.GetField(obj, "virtual_network_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *RouteTable) GetVirtualNetworkBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_back_refs, nil
}

func (obj *RouteTable) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & route_table_routes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.routes)
                if err != nil {
                        return nil, err
                }
                msg["routes"] = &value
        }

        if obj.modified & route_table_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & route_table_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *RouteTable) UnmarshalJSON(body []byte) error {
        var m map[string]json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
                return err
        }
        err = obj.UnmarshalCommon(m)
        if err != nil {
                return err
        }
        for key, value := range m {
                switch key {
                case "routes":
                        err = json.Unmarshal(value, &obj.routes)
                        if err == nil {
                                obj.valid |= route_table_routes
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= route_table_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= route_table_display_name
                        }
                        break
                case "virtual_network_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_network_back_refs)
                        if err == nil {
                                obj.valid |= route_table_virtual_network_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *RouteTable) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & route_table_routes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.routes)
                if err != nil {
                        return nil, err
                }
                msg["routes"] = &value
        }

        if obj.modified & route_table_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & route_table_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *RouteTable) UpdateReferences() error {

        return nil
}

func RouteTableByName(c contrail.ApiClient, fqn string) (*RouteTable, error) {
    obj, err := c.FindByName("route-table", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*RouteTable), nil
}

func RouteTableByUuid(c contrail.ApiClient, uuid string) (*RouteTable, error) {
    obj, err := c.FindByUuid("route-table", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*RouteTable), nil
}
