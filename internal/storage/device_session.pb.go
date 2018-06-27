// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device_session.proto

package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceSessionPBChannel struct {
	// Frequency (Hz).
	Frequency uint32 `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Min. data-rate.
	MinDr uint32 `protobuf:"varint,2,opt,name=min_dr,json=minDr,proto3" json:"min_dr,omitempty"`
	// Max. data-rate.
	MaxDr                uint32   `protobuf:"varint,3,opt,name=max_dr,json=maxDr,proto3" json:"max_dr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSessionPBChannel) Reset()         { *m = DeviceSessionPBChannel{} }
func (m *DeviceSessionPBChannel) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPBChannel) ProtoMessage()    {}
func (*DeviceSessionPBChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_session_82b396c549ab2196, []int{0}
}
func (m *DeviceSessionPBChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPBChannel.Unmarshal(m, b)
}
func (m *DeviceSessionPBChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPBChannel.Marshal(b, m, deterministic)
}
func (dst *DeviceSessionPBChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPBChannel.Merge(dst, src)
}
func (m *DeviceSessionPBChannel) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPBChannel.Size(m)
}
func (m *DeviceSessionPBChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPBChannel.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPBChannel proto.InternalMessageInfo

func (m *DeviceSessionPBChannel) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DeviceSessionPBChannel) GetMinDr() uint32 {
	if m != nil {
		return m.MinDr
	}
	return 0
}

func (m *DeviceSessionPBChannel) GetMaxDr() uint32 {
	if m != nil {
		return m.MaxDr
	}
	return 0
}

type DeviceSessionPBUplinkADRHistory struct {
	// Uplink frame-counter.
	FCnt uint32 `protobuf:"varint,1,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Max SNR (of deduplicated frames received by one or multiple gateways).
	MaxSnr float32 `protobuf:"fixed32,2,opt,name=max_snr,json=maxSnr,proto3" json:"max_snr,omitempty"`
	// TX Power (as known by the network-server).
	TxPowerIndex uint32 `protobuf:"varint,3,opt,name=tx_power_index,json=txPowerIndex,proto3" json:"tx_power_index,omitempty"`
	// Number of receiving gateways.
	GatewayCount         uint32   `protobuf:"varint,4,opt,name=gateway_count,json=gatewayCount,proto3" json:"gateway_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSessionPBUplinkADRHistory) Reset()         { *m = DeviceSessionPBUplinkADRHistory{} }
func (m *DeviceSessionPBUplinkADRHistory) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPBUplinkADRHistory) ProtoMessage()    {}
func (*DeviceSessionPBUplinkADRHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_session_82b396c549ab2196, []int{1}
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Unmarshal(m, b)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Marshal(b, m, deterministic)
}
func (dst *DeviceSessionPBUplinkADRHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Merge(dst, src)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Size(m)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPBUplinkADRHistory.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPBUplinkADRHistory proto.InternalMessageInfo

func (m *DeviceSessionPBUplinkADRHistory) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *DeviceSessionPBUplinkADRHistory) GetMaxSnr() float32 {
	if m != nil {
		return m.MaxSnr
	}
	return 0
}

func (m *DeviceSessionPBUplinkADRHistory) GetTxPowerIndex() uint32 {
	if m != nil {
		return m.TxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPBUplinkADRHistory) GetGatewayCount() uint32 {
	if m != nil {
		return m.GatewayCount
	}
	return 0
}

type DeviceSessionPBUplinkGatewayHistory struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSessionPBUplinkGatewayHistory) Reset()         { *m = DeviceSessionPBUplinkGatewayHistory{} }
func (m *DeviceSessionPBUplinkGatewayHistory) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPBUplinkGatewayHistory) ProtoMessage()    {}
func (*DeviceSessionPBUplinkGatewayHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_session_82b396c549ab2196, []int{2}
}
func (m *DeviceSessionPBUplinkGatewayHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPBUplinkGatewayHistory.Unmarshal(m, b)
}
func (m *DeviceSessionPBUplinkGatewayHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPBUplinkGatewayHistory.Marshal(b, m, deterministic)
}
func (dst *DeviceSessionPBUplinkGatewayHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPBUplinkGatewayHistory.Merge(dst, src)
}
func (m *DeviceSessionPBUplinkGatewayHistory) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPBUplinkGatewayHistory.Size(m)
}
func (m *DeviceSessionPBUplinkGatewayHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPBUplinkGatewayHistory.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPBUplinkGatewayHistory proto.InternalMessageInfo

type DeviceSessionPB struct {
	// ID of the device-profile.
	DeviceProfileId string `protobuf:"bytes,1,opt,name=device_profile_id,json=deviceProfileId,proto3" json:"device_profile_id,omitempty"`
	// ID of the service-profile.
	ServiceProfileId string `protobuf:"bytes,2,opt,name=service_profile_id,json=serviceProfileId,proto3" json:"service_profile_id,omitempty"`
	// ID of the routing-profile.
	RoutingProfileId string `protobuf:"bytes,3,opt,name=routing_profile_id,json=routingProfileId,proto3" json:"routing_profile_id,omitempty"`
	// Device address.
	DevAddr []byte `protobuf:"bytes,4,opt,name=dev_addr,json=devAddr,proto3" json:"dev_addr,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,5,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Join EUI.
	JoinEui []byte `protobuf:"bytes,6,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	// FNwkSIntKey.
	FNwkSIntKey []byte `protobuf:"bytes,7,opt,name=f_nwk_s_int_key,json=fNwkSIntKey,proto3" json:"f_nwk_s_int_key,omitempty"`
	// SNwkSIntKey.
	SNwkSIntKey []byte `protobuf:"bytes,8,opt,name=s_nwk_s_int_key,json=sNwkSIntKey,proto3" json:"s_nwk_s_int_key,omitempty"`
	// NwkSEncKey.
	NwkSEncKey []byte `protobuf:"bytes,9,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	// Uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,10,opt,name=f_cnt_up,json=fCntUp,proto3" json:"f_cnt_up,omitempty"`
	// Downlink frame-counter (network-server).
	NFCntDown uint32 `protobuf:"varint,11,opt,name=n_f_cnt_down,json=nFCntDown,proto3" json:"n_f_cnt_down,omitempty"`
	// Uplink frame-counter (application-server).
	// Note: this frame-counter is managed by the application-server.
	AFCntDown uint32 `protobuf:"varint,12,opt,name=a_f_cnt_down,json=aFCntDown,proto3" json:"a_f_cnt_down,omitempty"`
	// Frame-counter holding the last confirmed downlink frame (n_f_cnt_down or a_f_cnt_down).
	ConfFCnt uint32 `protobuf:"varint,39,opt,name=conf_f_cnt,json=confFCnt,proto3" json:"conf_f_cnt,omitempty"`
	// Skip uplink frame-counter validation.
	SkipFCntCheck bool `protobuf:"varint,13,opt,name=skip_f_cnt_check,json=skipFCntCheck,proto3" json:"skip_f_cnt_check,omitempty"`
	// RX Delay.
	RxDelay uint32 `protobuf:"varint,14,opt,name=rx_delay,json=rxDelay,proto3" json:"rx_delay,omitempty"`
	// RX1 data-rate offset.
	Rx1DrOffset uint32 `protobuf:"varint,15,opt,name=rx1_dr_offset,json=rx1DrOffset,proto3" json:"rx1_dr_offset,omitempty"`
	// RX2 data-rate.
	Rx2Dr uint32 `protobuf:"varint,16,opt,name=rx2_dr,json=rx2Dr,proto3" json:"rx2_dr,omitempty"`
	// RX2 frequency.
	Rx2Frequency uint32 `protobuf:"varint,17,opt,name=rx2_frequency,json=rx2Frequency,proto3" json:"rx2_frequency,omitempty"`
	// TXPowerIndex which the node is using. The possible values are defined
	// by the lorawan/band package and are region specific. By default it is
	// assumed that the node is using TXPower 0. This value is controlled by
	// the ADR engine.
	TxPowerIndex uint32 `protobuf:"varint,18,opt,name=tx_power_index,json=txPowerIndex,proto3" json:"tx_power_index,omitempty"`
	// DR defines the (last known) data-rate at which the node is operating.
	// This value is controlled by the ADR engine.
	Dr uint32 `protobuf:"varint,19,opt,name=dr,proto3" json:"dr,omitempty"`
	// ADR defines if the device has ADR enabled.
	Adr bool `protobuf:"varint,20,opt,name=adr,proto3" json:"adr,omitempty"`
	// MaxSupportedTXPowerIndex defines the maximum supported tx-power index
	// by the node, or 0 when not set.
	MaxSupportedTxPowerIndex uint32 `protobuf:"varint,21,opt,name=max_supported_tx_power_index,json=maxSupportedTxPowerIndex,proto3" json:"max_supported_tx_power_index,omitempty"`
	// MaxSupportedDR defines the maximum supported DR index by the node,
	// or 0 when not set.
	MaxSupportedDr uint32 `protobuf:"varint,22,opt,name=max_supported_dr,json=maxSupportedDr,proto3" json:"max_supported_dr,omitempty"`
	// NbTrans defines the number of transmissions for each unconfirmed uplink
	// frame. In case of 0, the default value is used.
	// This value is controlled by the ADR engine.
	NbTrans uint32 `protobuf:"varint,23,opt,name=nb_trans,json=nbTrans,proto3" json:"nb_trans,omitempty"`
	// Channels that are activated on the device.
	EnabledUplinkChannels []uint32 `protobuf:"varint,24,rep,packed,name=enabled_uplink_channels,json=enabledUplinkChannels,proto3" json:"enabled_uplink_channels,omitempty"`
	// Extra uplink channels, configured by the user.
	ExtraUplinkChannels map[uint32]*DeviceSessionPBChannel `protobuf:"bytes,25,rep,name=extra_uplink_channels,json=extraUplinkChannels,proto3" json:"extra_uplink_channels,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Frequency of each channel.
	ChannelFrequencies []uint32 `protobuf:"varint,26,rep,packed,name=channel_frequencies,json=channelFrequencies,proto3" json:"channel_frequencies,omitempty"`
	// Uplink history for ADR (last 20 uplink transmissions).
	UplinkAdrHistory []*DeviceSessionPBUplinkADRHistory `protobuf:"bytes,27,rep,name=uplink_adr_history,json=uplinkAdrHistory,proto3" json:"uplink_adr_history,omitempty"`
	// Uplink gateway history (for Class B / C).
	// Note that as bytes are not supported in a Protobuf map, the gateway ID is HEX encoded.
	UplinkGatewayHistory map[string]*DeviceSessionPBUplinkGatewayHistory `protobuf:"bytes,28,rep,name=uplink_gateway_history,json=uplinkGatewayHistory,proto3" json:"uplink_gateway_history,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Last device-status requested timestamp (Unix ns)
	LastDeviceStatusRequestTimeUnixNs int64 `protobuf:"varint,29,opt,name=last_device_status_request_time_unix_ns,json=lastDeviceStatusRequestTimeUnixNs,proto3" json:"last_device_status_request_time_unix_ns,omitempty"`
	// Last downlink timestamp (Unix ns).
	LastDownlinkTxTimestampUnixNs int64 `protobuf:"varint,32,opt,name=last_downlink_tx_timestamp_unix_ns,json=lastDownlinkTxTimestampUnixNs,proto3" json:"last_downlink_tx_timestamp_unix_ns,omitempty"`
	// Class-B beacon is locked.
	BeaconLocked bool `protobuf:"varint,33,opt,name=beacon_locked,json=beaconLocked,proto3" json:"beacon_locked,omitempty"`
	// Class-B ping-slot nb.
	PingSlotNb uint32 `protobuf:"varint,34,opt,name=ping_slot_nb,json=pingSlotNb,proto3" json:"ping_slot_nb,omitempty"`
	// Class-B ping-slot data-rate.
	PingSlotDr uint32 `protobuf:"varint,35,opt,name=ping_slot_dr,json=pingSlotDr,proto3" json:"ping_slot_dr,omitempty"`
	// Class-B ping-slot tx frequency.
	PingSlotFrequency uint32 `protobuf:"varint,36,opt,name=ping_slot_frequency,json=pingSlotFrequency,proto3" json:"ping_slot_frequency,omitempty"`
	// LoRaWAN mac-version.
	MacVersion string `protobuf:"bytes,37,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	// MinSupportedTXPowerIndex defines the minimum supported tx-power index
	// by the node (default 0).
	MinSupportedTxPowerIndex uint32 `protobuf:"varint,38,opt,name=min_supported_tx_power_index,json=minSupportedTxPowerIndex,proto3" json:"min_supported_tx_power_index,omitempty"`
	// RejoinRequestEnabled defines if the rejoin-request is enabled on the
	// device.
	RejoinRequestEnabled bool `protobuf:"varint,44,opt,name=rejoin_request_enabled,json=rejoinRequestEnabled,proto3" json:"rejoin_request_enabled,omitempty"`
	// RejoinRequestMaxCountN defines the 2^(C+4) uplink message interval for
	// the rejoin-request.
	RejoinRequestMaxCountN uint32 `protobuf:"varint,40,opt,name=rejoin_request_max_count_n,json=rejoinRequestMaxCountN,proto3" json:"rejoin_request_max_count_n,omitempty"`
	// RejoinRequestMaxTimeN defines the 2^(T+10) time interval (seconds)
	// for the rejoin-request.
	RejoinRequestMaxTimeN uint32 `protobuf:"varint,41,opt,name=rejoin_request_max_time_n,json=rejoinRequestMaxTimeN,proto3" json:"rejoin_request_max_time_n,omitempty"`
	// Rejoin counter (RJCount0).
	// This counter is reset to 0 after each successful join-accept.
	RejoinCount_0 uint32 `protobuf:"varint,42,opt,name=rejoin_count_0,json=rejoinCount0,proto3" json:"rejoin_count_0,omitempty"`
	// Pending rejoin device-session contains a device-session which has not
	// yet been activated by the device (by sending a first uplink).
	PendingRejoinDeviceSession []byte   `protobuf:"bytes,43,opt,name=pending_rejoin_device_session,json=pendingRejoinDeviceSession,proto3" json:"pending_rejoin_device_session,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *DeviceSessionPB) Reset()         { *m = DeviceSessionPB{} }
func (m *DeviceSessionPB) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPB) ProtoMessage()    {}
func (*DeviceSessionPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_session_82b396c549ab2196, []int{3}
}
func (m *DeviceSessionPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPB.Unmarshal(m, b)
}
func (m *DeviceSessionPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPB.Marshal(b, m, deterministic)
}
func (dst *DeviceSessionPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPB.Merge(dst, src)
}
func (m *DeviceSessionPB) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPB.Size(m)
}
func (m *DeviceSessionPB) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPB.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPB proto.InternalMessageInfo

func (m *DeviceSessionPB) GetDeviceProfileId() string {
	if m != nil {
		return m.DeviceProfileId
	}
	return ""
}

func (m *DeviceSessionPB) GetServiceProfileId() string {
	if m != nil {
		return m.ServiceProfileId
	}
	return ""
}

func (m *DeviceSessionPB) GetRoutingProfileId() string {
	if m != nil {
		return m.RoutingProfileId
	}
	return ""
}

func (m *DeviceSessionPB) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *DeviceSessionPB) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *DeviceSessionPB) GetJoinEui() []byte {
	if m != nil {
		return m.JoinEui
	}
	return nil
}

func (m *DeviceSessionPB) GetFNwkSIntKey() []byte {
	if m != nil {
		return m.FNwkSIntKey
	}
	return nil
}

func (m *DeviceSessionPB) GetSNwkSIntKey() []byte {
	if m != nil {
		return m.SNwkSIntKey
	}
	return nil
}

func (m *DeviceSessionPB) GetNwkSEncKey() []byte {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func (m *DeviceSessionPB) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *DeviceSessionPB) GetNFCntDown() uint32 {
	if m != nil {
		return m.NFCntDown
	}
	return 0
}

func (m *DeviceSessionPB) GetAFCntDown() uint32 {
	if m != nil {
		return m.AFCntDown
	}
	return 0
}

func (m *DeviceSessionPB) GetConfFCnt() uint32 {
	if m != nil {
		return m.ConfFCnt
	}
	return 0
}

func (m *DeviceSessionPB) GetSkipFCntCheck() bool {
	if m != nil {
		return m.SkipFCntCheck
	}
	return false
}

func (m *DeviceSessionPB) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *DeviceSessionPB) GetRx1DrOffset() uint32 {
	if m != nil {
		return m.Rx1DrOffset
	}
	return 0
}

func (m *DeviceSessionPB) GetRx2Dr() uint32 {
	if m != nil {
		return m.Rx2Dr
	}
	return 0
}

func (m *DeviceSessionPB) GetRx2Frequency() uint32 {
	if m != nil {
		return m.Rx2Frequency
	}
	return 0
}

func (m *DeviceSessionPB) GetTxPowerIndex() uint32 {
	if m != nil {
		return m.TxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPB) GetDr() uint32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func (m *DeviceSessionPB) GetAdr() bool {
	if m != nil {
		return m.Adr
	}
	return false
}

func (m *DeviceSessionPB) GetMaxSupportedTxPowerIndex() uint32 {
	if m != nil {
		return m.MaxSupportedTxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPB) GetMaxSupportedDr() uint32 {
	if m != nil {
		return m.MaxSupportedDr
	}
	return 0
}

func (m *DeviceSessionPB) GetNbTrans() uint32 {
	if m != nil {
		return m.NbTrans
	}
	return 0
}

func (m *DeviceSessionPB) GetEnabledUplinkChannels() []uint32 {
	if m != nil {
		return m.EnabledUplinkChannels
	}
	return nil
}

func (m *DeviceSessionPB) GetExtraUplinkChannels() map[uint32]*DeviceSessionPBChannel {
	if m != nil {
		return m.ExtraUplinkChannels
	}
	return nil
}

func (m *DeviceSessionPB) GetChannelFrequencies() []uint32 {
	if m != nil {
		return m.ChannelFrequencies
	}
	return nil
}

func (m *DeviceSessionPB) GetUplinkAdrHistory() []*DeviceSessionPBUplinkADRHistory {
	if m != nil {
		return m.UplinkAdrHistory
	}
	return nil
}

func (m *DeviceSessionPB) GetUplinkGatewayHistory() map[string]*DeviceSessionPBUplinkGatewayHistory {
	if m != nil {
		return m.UplinkGatewayHistory
	}
	return nil
}

func (m *DeviceSessionPB) GetLastDeviceStatusRequestTimeUnixNs() int64 {
	if m != nil {
		return m.LastDeviceStatusRequestTimeUnixNs
	}
	return 0
}

func (m *DeviceSessionPB) GetLastDownlinkTxTimestampUnixNs() int64 {
	if m != nil {
		return m.LastDownlinkTxTimestampUnixNs
	}
	return 0
}

func (m *DeviceSessionPB) GetBeaconLocked() bool {
	if m != nil {
		return m.BeaconLocked
	}
	return false
}

func (m *DeviceSessionPB) GetPingSlotNb() uint32 {
	if m != nil {
		return m.PingSlotNb
	}
	return 0
}

func (m *DeviceSessionPB) GetPingSlotDr() uint32 {
	if m != nil {
		return m.PingSlotDr
	}
	return 0
}

func (m *DeviceSessionPB) GetPingSlotFrequency() uint32 {
	if m != nil {
		return m.PingSlotFrequency
	}
	return 0
}

func (m *DeviceSessionPB) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *DeviceSessionPB) GetMinSupportedTxPowerIndex() uint32 {
	if m != nil {
		return m.MinSupportedTxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPB) GetRejoinRequestEnabled() bool {
	if m != nil {
		return m.RejoinRequestEnabled
	}
	return false
}

func (m *DeviceSessionPB) GetRejoinRequestMaxCountN() uint32 {
	if m != nil {
		return m.RejoinRequestMaxCountN
	}
	return 0
}

func (m *DeviceSessionPB) GetRejoinRequestMaxTimeN() uint32 {
	if m != nil {
		return m.RejoinRequestMaxTimeN
	}
	return 0
}

func (m *DeviceSessionPB) GetRejoinCount_0() uint32 {
	if m != nil {
		return m.RejoinCount_0
	}
	return 0
}

func (m *DeviceSessionPB) GetPendingRejoinDeviceSession() []byte {
	if m != nil {
		return m.PendingRejoinDeviceSession
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceSessionPBChannel)(nil), "storage.DeviceSessionPBChannel")
	proto.RegisterType((*DeviceSessionPBUplinkADRHistory)(nil), "storage.DeviceSessionPBUplinkADRHistory")
	proto.RegisterType((*DeviceSessionPBUplinkGatewayHistory)(nil), "storage.DeviceSessionPBUplinkGatewayHistory")
	proto.RegisterType((*DeviceSessionPB)(nil), "storage.DeviceSessionPB")
	proto.RegisterMapType((map[uint32]*DeviceSessionPBChannel)(nil), "storage.DeviceSessionPB.ExtraUplinkChannelsEntry")
	proto.RegisterMapType((map[string]*DeviceSessionPBUplinkGatewayHistory)(nil), "storage.DeviceSessionPB.UplinkGatewayHistoryEntry")
}

func init() {
	proto.RegisterFile("device_session.proto", fileDescriptor_device_session_82b396c549ab2196)
}

var fileDescriptor_device_session_82b396c549ab2196 = []byte{
	// 1120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0x6d, 0x53, 0x22, 0x47,
	0x10, 0xc7, 0x0b, 0x38, 0x9f, 0x1a, 0x50, 0x6e, 0x7c, 0xb8, 0xd1, 0x78, 0x25, 0x87, 0x77, 0x91,
	0x5c, 0x2c, 0x73, 0x92, 0x87, 0xba, 0xba, 0x17, 0xa9, 0xf2, 0x44, 0x13, 0x2b, 0x89, 0xb1, 0x56,
	0xbd, 0xb7, 0x53, 0xcb, 0xce, 0xa0, 0x1b, 0x60, 0x76, 0x33, 0x3b, 0x0b, 0xcb, 0x57, 0xc9, 0xc7,
	0xcb, 0x27, 0x49, 0x4d, 0xcf, 0x20, 0x0f, 0x42, 0xde, 0x41, 0xf7, 0xaf, 0xbb, 0xe7, 0xa1, 0xff,
	0x3d, 0x0b, 0x5b, 0x5c, 0xf4, 0xc3, 0x40, 0xb0, 0x44, 0x24, 0x49, 0x18, 0xc9, 0x93, 0x58, 0x45,
	0x3a, 0x22, 0x2b, 0x89, 0x8e, 0x94, 0xff, 0x20, 0x6a, 0x1c, 0x76, 0x9a, 0x08, 0xdc, 0x5a, 0xff,
	0xcd, 0xe7, 0xf3, 0x47, 0x5f, 0x4a, 0xd1, 0x25, 0xfb, 0xb0, 0xd6, 0x56, 0xe2, 0xef, 0x54, 0xc8,
	0x60, 0x48, 0x73, 0xd5, 0x5c, 0xbd, 0xec, 0x8d, 0x0d, 0x64, 0x1b, 0x96, 0x7b, 0xa1, 0x64, 0x5c,
	0xd1, 0x3c, 0xba, 0x96, 0x7a, 0xa1, 0x6c, 0x2a, 0x34, 0xfb, 0x99, 0x31, 0x17, 0x9c, 0xd9, 0xcf,
	0x9a, 0xaa, 0xf6, 0x4f, 0x0e, 0x0e, 0x66, 0xca, 0xdc, 0xc7, 0xdd, 0x50, 0x76, 0xce, 0x9a, 0xde,
	0xaf, 0xa1, 0x59, 0xcb, 0x90, 0x6c, 0xc2, 0x52, 0x9b, 0x05, 0x52, 0xbb, 0x5a, 0x2f, 0xda, 0xe7,
	0x52, 0x93, 0x57, 0xb0, 0x62, 0xf2, 0x25, 0xd2, 0xd6, 0xc9, 0x7b, 0x26, 0xfd, 0xad, 0x54, 0xe4,
	0x2d, 0xac, 0xeb, 0x8c, 0xc5, 0xd1, 0x40, 0x28, 0x16, 0x4a, 0x2e, 0x32, 0x57, 0xb0, 0xa4, 0xb3,
	0x1b, 0x63, 0xbc, 0x32, 0x36, 0x72, 0x08, 0xe5, 0x07, 0x5f, 0x8b, 0x81, 0x3f, 0x64, 0x41, 0x94,
	0x4a, 0x4d, 0x5f, 0x58, 0xc8, 0x19, 0xcf, 0x8d, 0xad, 0xf6, 0x0e, 0x0e, 0xe7, 0xae, 0xed, 0x17,
	0x0b, 0xb9, 0xf5, 0xd5, 0xfe, 0xad, 0xc0, 0xc6, 0x0c, 0x47, 0xde, 0xc3, 0x4b, 0x77, 0xbc, 0xb1,
	0x8a, 0xda, 0x61, 0x57, 0xb0, 0x90, 0xe3, 0xfa, 0xd7, 0xbc, 0x0d, 0xeb, 0xb8, 0xb1, 0xf6, 0x2b,
	0x4e, 0x8e, 0x81, 0x24, 0x42, 0xcd, 0xc2, 0x79, 0x84, 0x2b, 0xce, 0x33, 0x45, 0xab, 0x28, 0xd5,
	0xa1, 0x7c, 0x98, 0xa4, 0x0b, 0x96, 0x76, 0x9e, 0x31, 0xbd, 0x0b, 0xab, 0x5c, 0xf4, 0x99, 0xcf,
	0xb9, 0xc2, 0x2d, 0x96, 0xbc, 0x15, 0x2e, 0xfa, 0x67, 0x9c, 0x2b, 0x73, 0x82, 0xc6, 0x25, 0xd2,
	0x90, 0x2e, 0xa1, 0x67, 0x99, 0x8b, 0xfe, 0x45, 0x1a, 0x9a, 0x98, 0xbf, 0xa2, 0x50, 0xa2, 0x67,
	0xd9, 0xc6, 0x98, 0xff, 0xc6, 0xf5, 0x16, 0x36, 0xda, 0x4c, 0x0e, 0x3a, 0x2c, 0x61, 0xa1, 0xd4,
	0xac, 0x23, 0x86, 0x74, 0x05, 0x89, 0x62, 0xfb, 0x7a, 0xd0, 0xb9, 0xbd, 0x92, 0xfa, 0x37, 0x31,
	0x34, 0x54, 0x32, 0x43, 0xad, 0x5a, 0x2a, 0x99, 0xa0, 0xde, 0x40, 0xd9, 0x32, 0x42, 0x06, 0xc8,
	0xac, 0x21, 0x03, 0x72, 0xd0, 0xb9, 0xbd, 0x90, 0x81, 0x41, 0x28, 0xac, 0xe2, 0xcd, 0xb3, 0x34,
	0xa6, 0x80, 0x17, 0xb4, 0x6c, 0x2e, 0xff, 0x3e, 0x26, 0x07, 0x50, 0x92, 0xcc, 0xfa, 0x78, 0x34,
	0x90, 0xb4, 0x68, 0xdb, 0x50, 0x5e, 0x9e, 0x4b, 0xdd, 0x8c, 0x06, 0xd2, 0x00, 0xfe, 0x24, 0x50,
	0xb2, 0x80, 0xff, 0x04, 0xec, 0x03, 0x04, 0x91, 0x6c, 0x5b, 0x86, 0x1e, 0xa1, 0x7b, 0xd5, 0x58,
	0x0c, 0x41, 0x8e, 0xa0, 0x92, 0x74, 0xc2, 0xd8, 0x65, 0x08, 0x1e, 0x45, 0xd0, 0xa1, 0xe5, 0x6a,
	0xae, 0xbe, 0xea, 0x95, 0x8d, 0xdd, 0x30, 0xe7, 0xc6, 0x68, 0x0e, 0x4b, 0x65, 0x8c, 0x8b, 0xae,
	0x3f, 0xa4, 0xeb, 0x98, 0x64, 0x45, 0x65, 0x4d, 0xf3, 0x97, 0xd4, 0xa0, 0xac, 0xb2, 0x53, 0xc6,
	0x15, 0x8b, 0xda, 0xed, 0x44, 0x68, 0xba, 0x81, 0xfe, 0xa2, 0xca, 0x4e, 0x9b, 0xea, 0x4f, 0x34,
	0x19, 0x59, 0xa8, 0xac, 0x61, 0x64, 0x51, 0xb1, 0xb2, 0x50, 0x59, 0xa3, 0xa9, 0x4c, 0x7b, 0x1a,
	0xf3, 0x58, 0x66, 0x2f, 0x6d, 0x7b, 0xaa, 0xac, 0x71, 0xf9, 0xa4, 0xb4, 0xe7, 0x9d, 0x4e, 0xe6,
	0x74, 0xfa, 0x3a, 0xe4, 0xb9, 0xa2, 0x9b, 0xe8, 0xc9, 0x73, 0x45, 0x2a, 0x50, 0xf0, 0xb9, 0xa2,
	0x5b, 0xb8, 0x19, 0xf3, 0x93, 0xfc, 0x0c, 0xfb, 0x28, 0xa5, 0x34, 0x8e, 0x23, 0xa5, 0x05, 0x67,
	0x33, 0x59, 0xb7, 0x31, 0x96, 0x1a, 0x7d, 0x8d, 0x90, 0xbb, 0xc9, 0x0a, 0x75, 0xa8, 0x4c, 0xc7,
	0x73, 0x45, 0x77, 0x30, 0x66, 0x7d, 0x32, 0xa6, 0xa9, 0xcc, 0x61, 0xc9, 0x16, 0xd3, 0xca, 0x97,
	0x09, 0x7d, 0x65, 0x0f, 0x4b, 0xb6, 0xee, 0xcc, 0x5f, 0xf2, 0x13, 0xbc, 0x12, 0xd2, 0x6f, 0x75,
	0x05, 0x67, 0x29, 0x8a, 0x8c, 0x05, 0x76, 0xdc, 0x24, 0x94, 0x56, 0x0b, 0xf5, 0xb2, 0xb7, 0xed,
	0xdc, 0x56, 0x82, 0x6e, 0x16, 0x25, 0x44, 0xc0, 0xb6, 0xc8, 0xb4, 0xf2, 0x9f, 0x45, 0xed, 0x56,
	0x0b, 0xf5, 0x62, 0xe3, 0xf4, 0xc4, 0xcd, 0xb3, 0x93, 0x19, 0x85, 0x9e, 0x5c, 0x98, 0xa8, 0xe9,
	0x64, 0x17, 0x52, 0xab, 0xa1, 0xb7, 0x29, 0x9e, 0x7b, 0xc8, 0x77, 0xb0, 0xe9, 0x32, 0x3f, 0x5d,
	0x4a, 0x28, 0x12, 0xba, 0x87, 0x4b, 0x23, 0xce, 0x75, 0x39, 0xf6, 0x90, 0x2f, 0x40, 0xdc, 0x8a,
	0x7c, 0xae, 0xd8, 0xa3, 0x1d, 0x15, 0xf4, 0x2b, 0x5c, 0x54, 0x7d, 0xd1, 0xa2, 0x66, 0x47, 0x9f,
	0x57, 0xb1, 0x39, 0xce, 0xb8, 0x1a, 0x0d, 0xc3, 0x47, 0xd8, 0x71, 0x79, 0x47, 0xf3, 0x6b, 0x94,
	0x7b, 0x1f, 0x73, 0x37, 0x16, 0x6e, 0x78, 0xde, 0xec, 0xb2, 0x3b, 0xde, 0x4a, 0xe7, 0xb8, 0x88,
	0x07, 0x47, 0x5d, 0x3f, 0xd1, 0x6c, 0xf4, 0x4c, 0x68, 0x5f, 0xa7, 0x09, 0xc3, 0x2d, 0x26, 0x9a,
	0xe9, 0xb0, 0x27, 0x58, 0x2a, 0xc3, 0x8c, 0xc9, 0x84, 0xbe, 0xae, 0xe6, 0xea, 0x05, 0xef, 0x8d,
	0xc1, 0x5d, 0x55, 0x84, 0x3d, 0xcb, 0xde, 0x85, 0x3d, 0x71, 0x2f, 0xc3, 0xec, 0x3a, 0x21, 0x57,
	0x50, 0xb3, 0x39, 0xa3, 0x81, 0xc4, 0x4d, 0xe8, 0x0c, 0x33, 0x25, 0xda, 0xef, 0xc5, 0x4f, 0xe9,
	0xaa, 0x98, 0xee, 0x35, 0xa6, 0x73, 0xe0, 0x5d, 0x76, 0x37, 0xc2, 0x5c, 0xaa, 0x43, 0x28, 0xb7,
	0x84, 0x1f, 0x44, 0x92, 0x75, 0xa3, 0xa0, 0x23, 0x38, 0x7d, 0x83, 0x1d, 0x5d, 0xb2, 0xc6, 0xdf,
	0xd1, 0x46, 0xaa, 0x50, 0x8a, 0xcd, 0xa4, 0x4c, 0xba, 0x91, 0x66, 0xb2, 0x45, 0x6b, 0xd8, 0x74,
	0x60, 0x6c, 0xb7, 0xdd, 0x48, 0x5f, 0xb7, 0xa6, 0x09, 0xae, 0xe8, 0xe1, 0x34, 0xd1, 0x54, 0xe4,
	0x04, 0x36, 0xc7, 0xc4, 0x58, 0x91, 0x6f, 0x11, 0x7c, 0x39, 0x02, 0xc7, 0xb2, 0x3c, 0x80, 0x62,
	0xcf, 0x0f, 0x58, 0x5f, 0x28, 0x73, 0xf0, 0xf4, 0x1d, 0x4e, 0x66, 0xe8, 0xf9, 0xc1, 0x17, 0x6b,
	0x41, 0xbd, 0x85, 0x72, 0xb1, 0xde, 0xbe, 0x76, 0x7a, 0x0b, 0xe5, 0x7c, 0xbd, 0xfd, 0x00, 0x3b,
	0x4a, 0xe0, 0x84, 0x1e, 0x5d, 0x86, 0x93, 0x06, 0x3d, 0xc6, 0x23, 0xd8, 0xb2, 0x5e, 0x77, 0xfa,
	0x17, 0xd6, 0x47, 0x3e, 0xc1, 0xde, 0x4c, 0x94, 0x11, 0x2d, 0x3e, 0x7e, 0x4c, 0xd2, 0x3a, 0xd6,
	0xdc, 0x99, 0x8a, 0xfc, 0xc3, 0xcf, 0xf0, 0x1d, 0xbc, 0x26, 0x1f, 0x61, 0x77, 0x4e, 0x2c, 0xb6,
	0x80, 0xa4, 0xdf, 0x60, 0xe8, 0xf6, 0x6c, 0xa8, 0xb9, 0xaf, 0x6b, 0x33, 0xa3, 0x5c, 0xa4, 0xad,
	0xf4, 0x81, 0xbe, 0x77, 0x93, 0x0c, 0xad, 0x98, 0xff, 0x03, 0x39, 0x83, 0xd7, 0xb1, 0x90, 0xdc,
	0x9c, 0xb2, 0xa3, 0xa7, 0xbf, 0x4d, 0xe8, 0xb7, 0xf8, 0x34, 0xec, 0x39, 0xc8, 0x43, 0x66, 0xaa,
	0xbf, 0xf7, 0x1e, 0x80, 0x2e, 0x52, 0xb4, 0x19, 0x79, 0xe6, 0x7d, 0xb1, 0x9f, 0x0f, 0xe6, 0x27,
	0xf9, 0x11, 0x96, 0xfa, 0x7e, 0x37, 0x15, 0xf8, 0xca, 0x16, 0x1b, 0x07, 0x8b, 0x44, 0xe3, 0xf2,
	0x78, 0x96, 0xfe, 0x94, 0xff, 0x98, 0xdb, 0x4b, 0x61, 0x77, 0xa1, 0x92, 0x26, 0x2b, 0xad, 0xd9,
	0x4a, 0x9f, 0xa7, 0x2b, 0x1d, 0xff, 0xbf, 0xf4, 0xa7, 0x73, 0x4e, 0x94, 0x6d, 0x2d, 0xe3, 0xe7,
	0xd9, 0xf7, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x66, 0xad, 0xc0, 0xb6, 0x09, 0x00, 0x00,
}