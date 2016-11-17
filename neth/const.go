package neth

// UDPGuaranteedPayloadSize is maximal guaranteed size of UDP payload which does not fragment and must be delivered.
// = 576 (RFC792) - 60 (Max IP Header Size) — 8 (UDP Header Size)
const UDPGuaranteedPayloadSize = 508

// UDPUsualMaxPayloadSize is usual size of UDP payload which does not fragment and must be delivered via internet (if no encapsulation is used).
// Primary used for servers where encapsulation more or less can be checked.
// 1500 (MTU) — 60(Max IP Header Size) — 8(UDP Header Size)
const UDPUsualMaxPayloadSize = 1432
