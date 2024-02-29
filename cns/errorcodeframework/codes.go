package errorcodeframework

type ResponseCode int

// ResponseCode definitions
// 4000 codes cannot be retried
// 5000 codes can be retried
// x100 codes are related to network, connectivity, and environment
// x300 codes are related to container, orchestrator, and virtualization
// x500 codes are related to authorization and unexpected errors
// x700 codes are related agent and internal server errors
const (
	Success                                ResponseCode = 2000
	NetworkContainerVfpProgramComplete     ResponseCode = 2352
	InvalidParameter                       ResponseCode = 4111
	MalformedSubnet                        ResponseCode = 4121
	NetworkContainerNotSpecified           ResponseCode = 4131
	ReservationNotFound                    ResponseCode = 4141
	UnknownContainerID                     ResponseCode = 4151
	UnreachableDockerDaemon                ResponseCode = 4161
	UnreachableHost                        ResponseCode = 4164
	UnspecifiedNetworkName                 ResponseCode = 4175
	UnsupportedEnvironment                 ResponseCode = 4181
	UnsupportedNetworkType                 ResponseCode = 4185
	DockerContainerNotSpecified            ResponseCode = 4311
	EmptyOrchestratorContext               ResponseCode = 4321
	InvalidPrimaryIPConfig                 ResponseCode = 4341
	InvalidRequest                         ResponseCode = 4342
	InvalidSecondaryIPConfig               ResponseCode = 4345
	NilEndpointStateStore                  ResponseCode = 4356
	PrimaryCANotSame                       ResponseCode = 4361
	UnsupportedNCVersion                   ResponseCode = 4381
	UnsupportedNetworkContainerType        ResponseCode = 4383
	UnsupportedOrchestratorContext         ResponseCode = 4385
	UnsupportedOrchestratorType            ResponseCode = 4386
	UnsupportedVerb                        ResponseCode = 4390
	StatusUnauthorized                     ResponseCode = 4561
	UnexpectedError                        ResponseCode = 4580
	UnsupportedAPI                         ResponseCode = 4591
	AddressUnavailable                     ResponseCode = 5101
	CallToHostFailed                       ResponseCode = 5105
	InconsistentIPConfigState              ResponseCode = 5110
	NetworkContainerPublishFailed          ResponseCode = 5133
	NetworkContainerUnpublishFailed        ResponseCode = 5135
	NetworkJoinFailed                      ResponseCode = 5138
	NotFound                               ResponseCode = 5140
	FailedToAllocateIPConfig               ResponseCode = 5331
	NetworkContainerVfpProgramCheckSkipped ResponseCode = 5351
	NetworkContainerVfpProgramPending      ResponseCode = 5354
	FailedToRunIPTableCmd                  ResponseCode = 5721
	NmAgentInternalServerError             ResponseCode = 5751
	NmAgentSupportedApisError              ResponseCode = 5755
)

// nolint:gocyclo
func (c ResponseCode) String() string {
	switch c {
	case AddressUnavailable:
		return "AddressUnavailable"
	case CallToHostFailed:
		return "CallToHostFailed"
	case DockerContainerNotSpecified:
		return "DockerContainerNotSpecified"
	case EmptyOrchestratorContext:
		return "EmptyOrchestratorContext"
	case FailedToAllocateIPConfig:
		return "FailedToAllocateIPConfig"
	case FailedToRunIPTableCmd:
		return "FailedToRunIPTableCmd"
	case InconsistentIPConfigState:
		return "InconsistentIPConfigState"
	case InvalidParameter:
		return "InvalidParameter"
	case InvalidPrimaryIPConfig:
		return "InvalidPrimaryIPConfig"
	case InvalidRequest:
		return "InvalidRequest"
	case InvalidSecondaryIPConfig:
		return "InvalidSecondaryIPConfig"
	case MalformedSubnet:
		return "MalformedSubnet"
	case NetworkContainerNotSpecified:
		return "NetworkContainerNotSpecified"
	case NetworkContainerPublishFailed:
		return "NetworkContainerPublishFailed"
	case NetworkContainerUnpublishFailed:
		return "NetworkContainerUnpublishFailed"
	case NetworkContainerVfpProgramCheckSkipped:
		return "NetworkContainerVfpProgramCheckSkipped"
	case NetworkContainerVfpProgramComplete:
		return "NetworkContainerVfpProgramComplete"
	case NetworkContainerVfpProgramPending:
		return "NetworkContainerVfpProgramPending"
	case NetworkJoinFailed:
		return "NetworkJoinFailed"
	case NilEndpointStateStore:
		return "NilEndpointStateStore"
	case NmAgentInternalServerError:
		return "NmAgentInternalServerError"
	case NmAgentSupportedApisError:
		return "NmAgentSupportedApisError"
	case NotFound:
		return "NotFound"
	case PrimaryCANotSame:
		return "PrimaryCANotSame"
	case ReservationNotFound:
		return "ReservationNotFound"
	case StatusUnauthorized:
		return "StatusUnauthorized"
	case Success:
		return "Success"
	case UnexpectedError:
		return "UnexpectedError"
	case UnknownContainerID:
		return "UnknownContainerID"
	case UnreachableDockerDaemon:
		return "UnreachableDockerDaemon"
	case UnreachableHost:
		return "UnreachableHost"
	case UnspecifiedNetworkName:
		return "UnspecifiedNetworkName"
	case UnsupportedAPI:
		return "UnsupportedAPI"
	case UnsupportedEnvironment:
		return "UnsupportedEnvironment"
	case UnsupportedNCVersion:
		return "UnsupportedNCVersion"
	case UnsupportedNetworkContainerType:
		return "UnsupportedNetworkContainerType"
	case UnsupportedNetworkType:
		return "UnsupportedNetworkType"
	case UnsupportedOrchestratorContext:
		return "UnsupportedOrchestratorContext"
	case UnsupportedOrchestratorType:
		return "UnsupportedOrchestratorType"
	case UnsupportedVerb:
		return "UnsupportedVerb"
	default:
		return "UnknownError"
	}
}
