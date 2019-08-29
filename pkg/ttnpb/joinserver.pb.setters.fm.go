// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"

	go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
)

func (dst *SessionKeyRequest) SetFields(src *SessionKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "session_key_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_key_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionKeyID = src.SessionKeyID
			} else {
				dst.SessionKeyID = nil
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEUI = src.DevEUI
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.EUI64
				dst.DevEUI = zero
			}
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.EUI64
				dst.JoinEUI = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *NwkSKeysResponse) SetFields(src *NwkSKeysResponse, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "f_nwk_s_int_key":
			if len(subs) > 0 {
				newDst := &dst.FNwkSIntKey
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = &src.FNwkSIntKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.FNwkSIntKey = src.FNwkSIntKey
				} else {
					var zero KeyEnvelope
					dst.FNwkSIntKey = zero
				}
			}
		case "s_nwk_s_int_key":
			if len(subs) > 0 {
				newDst := &dst.SNwkSIntKey
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = &src.SNwkSIntKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SNwkSIntKey = src.SNwkSIntKey
				} else {
					var zero KeyEnvelope
					dst.SNwkSIntKey = zero
				}
			}
		case "nwk_s_enc_key":
			if len(subs) > 0 {
				newDst := &dst.NwkSEncKey
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = &src.NwkSEncKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.NwkSEncKey = src.NwkSEncKey
				} else {
					var zero KeyEnvelope
					dst.NwkSEncKey = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AppSKeyResponse) SetFields(src *AppSKeyResponse, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "app_s_key":
			if len(subs) > 0 {
				newDst := &dst.AppSKey
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = &src.AppSKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AppSKey = src.AppSKey
				} else {
					var zero KeyEnvelope
					dst.AppSKey = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CryptoServicePayloadRequest) SetFields(src *CryptoServicePayloadRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.EndDeviceIdentifiers
				var newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "lorawan_version":
			if len(subs) > 0 {
				return fmt.Errorf("'lorawan_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LoRaWANVersion = src.LoRaWANVersion
			} else {
				var zero MACVersion
				dst.LoRaWANVersion = zero
			}
		case "payload":
			if len(subs) > 0 {
				return fmt.Errorf("'payload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Payload = src.Payload
			} else {
				dst.Payload = nil
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerID = src.ProvisionerID
			} else {
				var zero string
				dst.ProvisionerID = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CryptoServicePayloadResponse) SetFields(src *CryptoServicePayloadResponse, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "payload":
			if len(subs) > 0 {
				return fmt.Errorf("'payload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Payload = src.Payload
			} else {
				dst.Payload = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *JoinAcceptMICRequest) SetFields(src *JoinAcceptMICRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "payload_request":
			if len(subs) > 0 {
				newDst := &dst.CryptoServicePayloadRequest
				var newSrc *CryptoServicePayloadRequest
				if src != nil {
					newSrc = &src.CryptoServicePayloadRequest
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.CryptoServicePayloadRequest = src.CryptoServicePayloadRequest
				} else {
					var zero CryptoServicePayloadRequest
					dst.CryptoServicePayloadRequest = zero
				}
			}
		case "join_request_type":
			if len(subs) > 0 {
				return fmt.Errorf("'join_request_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinRequestType = src.JoinRequestType
			} else {
				var zero RejoinType
				dst.JoinRequestType = zero
			}
		case "dev_nonce":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_nonce' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevNonce = src.DevNonce
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.DevNonce
				dst.DevNonce = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeriveSessionKeysRequest) SetFields(src *DeriveSessionKeysRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.EndDeviceIdentifiers
				var newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "lorawan_version":
			if len(subs) > 0 {
				return fmt.Errorf("'lorawan_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LoRaWANVersion = src.LoRaWANVersion
			} else {
				var zero MACVersion
				dst.LoRaWANVersion = zero
			}
		case "join_nonce":
			if len(subs) > 0 {
				return fmt.Errorf("'join_nonce' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinNonce = src.JoinNonce
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.JoinNonce
				dst.JoinNonce = zero
			}
		case "dev_nonce":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_nonce' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevNonce = src.DevNonce
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.DevNonce
				dst.DevNonce = zero
			}
		case "net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NetID = src.NetID
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.NetID
				dst.NetID = zero
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerID = src.ProvisionerID
			} else {
				var zero string
				dst.ProvisionerID = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRootKeysRequest) SetFields(src *GetRootKeysRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.EndDeviceIdentifiers
				var newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerID = src.ProvisionerID
			} else {
				var zero string
				dst.ProvisionerID = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest) SetFields(src *ProvisionEndDevicesRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				newDst := &dst.ApplicationIdentifiers
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerID = src.ProvisionerID
			} else {
				var zero string
				dst.ProvisionerID = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		case "end_devices":
			if len(subs) == 0 && src == nil {
				dst.EndDevices = nil
				continue
			} else if len(subs) == 0 {
				dst.EndDevices = src.EndDevices
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "list":
					if _, ok := dst.EndDevices.(*ProvisionEndDevicesRequest_List); !ok {
						dst.EndDevices = &ProvisionEndDevicesRequest_List{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.EndDevices.(*ProvisionEndDevicesRequest_List).List
						if newDst == nil {
							newDst = &ProvisionEndDevicesRequest_IdentifiersList{}
							dst.EndDevices.(*ProvisionEndDevicesRequest_List).List = newDst
						}
						var newSrc *ProvisionEndDevicesRequest_IdentifiersList
						if src != nil {
							newSrc = src.GetList()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.EndDevices.(*ProvisionEndDevicesRequest_List).List = src.GetList()
						} else {
							dst.EndDevices.(*ProvisionEndDevicesRequest_List).List = nil
						}
					}
				case "range":
					if _, ok := dst.EndDevices.(*ProvisionEndDevicesRequest_Range); !ok {
						dst.EndDevices = &ProvisionEndDevicesRequest_Range{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.EndDevices.(*ProvisionEndDevicesRequest_Range).Range
						if newDst == nil {
							newDst = &ProvisionEndDevicesRequest_IdentifiersRange{}
							dst.EndDevices.(*ProvisionEndDevicesRequest_Range).Range = newDst
						}
						var newSrc *ProvisionEndDevicesRequest_IdentifiersRange
						if src != nil {
							newSrc = src.GetRange()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.EndDevices.(*ProvisionEndDevicesRequest_Range).Range = src.GetRange()
						} else {
							dst.EndDevices.(*ProvisionEndDevicesRequest_Range).Range = nil
						}
					}
				case "from_data":
					if _, ok := dst.EndDevices.(*ProvisionEndDevicesRequest_FromData); !ok {
						dst.EndDevices = &ProvisionEndDevicesRequest_FromData{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.EndDevices.(*ProvisionEndDevicesRequest_FromData).FromData
						if newDst == nil {
							newDst = &ProvisionEndDevicesRequest_IdentifiersFromData{}
							dst.EndDevices.(*ProvisionEndDevicesRequest_FromData).FromData = newDst
						}
						var newSrc *ProvisionEndDevicesRequest_IdentifiersFromData
						if src != nil {
							newSrc = src.GetFromData()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.EndDevices.(*ProvisionEndDevicesRequest_FromData).FromData = src.GetFromData()
						} else {
							dst.EndDevices.(*ProvisionEndDevicesRequest_FromData).FromData = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *JoinEUIPrefix) SetFields(src *JoinEUIPrefix, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.EUI64
				dst.JoinEUI = zero
			}
		case "length":
			if len(subs) > 0 {
				return fmt.Errorf("'length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Length = src.Length
			} else {
				var zero uint32
				dst.Length = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *JoinEUIPrefixes) SetFields(src *JoinEUIPrefixes, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "prefixes":
			if len(subs) > 0 {
				return fmt.Errorf("'prefixes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Prefixes = src.Prefixes
			} else {
				dst.Prefixes = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest_IdentifiersList) SetFields(src *ProvisionEndDevicesRequest_IdentifiersList, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				dst.JoinEUI = nil
			}
		case "end_device_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'end_device_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.EndDeviceIDs = src.EndDeviceIDs
			} else {
				dst.EndDeviceIDs = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest_IdentifiersRange) SetFields(src *ProvisionEndDevicesRequest_IdentifiersRange, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				dst.JoinEUI = nil
			}
		case "start_dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'start_dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.StartDevEUI = src.StartDevEUI
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.EUI64
				dst.StartDevEUI = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest_IdentifiersFromData) SetFields(src *ProvisionEndDevicesRequest_IdentifiersFromData, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				dst.JoinEUI = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
