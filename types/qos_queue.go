//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Elaugaste/contrail-go-api"
)

const (
	qos_queue_min_bandwidth uint64 = 1 << iota
	qos_queue_max_bandwidth
	qos_queue_id_perms
	qos_queue_display_name
	qos_queue_qos_forwarding_class_back_refs
)

type QosQueue struct {
        contrail.ObjectBase
	min_bandwidth int
	max_bandwidth int
	id_perms IdPermsType
	display_name string
	qos_forwarding_class_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *QosQueue) GetType() string {
        return "qos-queue"
}

func (obj *QosQueue) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *QosQueue) GetDefaultParentType() string {
        return "project"
}

func (obj *QosQueue) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *QosQueue) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *QosQueue) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *QosQueue) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *QosQueue) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *QosQueue) GetMinBandwidth() int {
        return obj.min_bandwidth
}

func (obj *QosQueue) SetMinBandwidth(value int) {
        obj.min_bandwidth = value
        obj.modified |= qos_queue_min_bandwidth
}

func (obj *QosQueue) GetMaxBandwidth() int {
        return obj.max_bandwidth
}

func (obj *QosQueue) SetMaxBandwidth(value int) {
        obj.max_bandwidth = value
        obj.modified |= qos_queue_max_bandwidth
}

func (obj *QosQueue) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *QosQueue) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= qos_queue_id_perms
}

func (obj *QosQueue) GetDisplayName() string {
        return obj.display_name
}

func (obj *QosQueue) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= qos_queue_display_name
}

func (obj *QosQueue) readQosForwardingClassBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & qos_queue_qos_forwarding_class_back_refs == 0) {
                err := obj.GetField(obj, "qos_forwarding_class_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosQueue) GetQosForwardingClassBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readQosForwardingClassBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.qos_forwarding_class_back_refs, nil
}

func (obj *QosQueue) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & qos_queue_min_bandwidth != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.min_bandwidth)
                if err != nil {
                        return nil, err
                }
                msg["min_bandwidth"] = &value
        }

        if obj.modified & qos_queue_max_bandwidth != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.max_bandwidth)
                if err != nil {
                        return nil, err
                }
                msg["max_bandwidth"] = &value
        }

        if obj.modified & qos_queue_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & qos_queue_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *QosQueue) UnmarshalJSON(body []byte) error {
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
                case "min_bandwidth":
                        err = json.Unmarshal(value, &obj.min_bandwidth)
                        if err == nil {
                                obj.valid |= qos_queue_min_bandwidth
                        }
                        break
                case "max_bandwidth":
                        err = json.Unmarshal(value, &obj.max_bandwidth)
                        if err == nil {
                                obj.valid |= qos_queue_max_bandwidth
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= qos_queue_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= qos_queue_display_name
                        }
                        break
                case "qos_forwarding_class_back_refs":
                        err = json.Unmarshal(value, &obj.qos_forwarding_class_back_refs)
                        if err == nil {
                                obj.valid |= qos_queue_qos_forwarding_class_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosQueue) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & qos_queue_min_bandwidth != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.min_bandwidth)
                if err != nil {
                        return nil, err
                }
                msg["min_bandwidth"] = &value
        }

        if obj.modified & qos_queue_max_bandwidth != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.max_bandwidth)
                if err != nil {
                        return nil, err
                }
                msg["max_bandwidth"] = &value
        }

        if obj.modified & qos_queue_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & qos_queue_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *QosQueue) UpdateReferences() error {

        return nil
}

func QosQueueByName(c contrail.ApiClient, fqn string) (*QosQueue, error) {
    obj, err := c.FindByName("qos-queue", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*QosQueue), nil
}

func QosQueueByUuid(c contrail.ApiClient, uuid string) (*QosQueue, error) {
    obj, err := c.FindByUuid("qos-queue", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*QosQueue), nil
}
