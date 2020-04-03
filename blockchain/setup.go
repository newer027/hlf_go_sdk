package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	// mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	// "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	// packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	// "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
	"sync"
	"io"
	"time"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile      string
	OrgID           string
	OrdererID		string
	ChannelID       string
	ChainCodeID     string
	initialized     bool
	ChannelConfig   string
	ChaincodeGoPath string
	ChaincodePath   string
	OrgAdmin        string
	OrgName         string
	UserName        string
	client          *channel.Client
	admin           *resmgmt.Client
	sdk             *fabsdk.FabricSDK
	event           *event.Client
}

// Initialize reads the configuration file and sets up the client, chain and event hub
func (setup *FabricSetup) Initialize() error {

	// Add parameters for the initialization
	if setup.initialized {
		return errors.New("sdk already initialized")
	}

	// Initialize the SDK with the configuration file
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}
	setup.sdk = sdk
	fmt.Println("SDK created")

	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity")
	}
	setup.admin = resMgmtClient
	fmt.Println("Resource management client created")
 	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
 	setup.client, err = channel.New(clientContext)
 	setup.event, err = event.New(clientContext)
 	if err != nil {
 		return errors.WithMessage(err, "failed to create new event client")
 	}
	 // The MSP client allow us to retrieve user information from their identity, like its signing identity which we will need to save the channel
// 	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(setup.OrgName))
 	if err != nil {
 		return errors.WithMessage(err, "failed to create MSP client")
 	}
// 	adminIdentity, err := mspClient.GetSigningIdentity(setup.OrgAdmin)
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to get admin signing identity")
// 	}
// 	req := resmgmt.SaveChannelRequest{ChannelID: setup.ChannelID, ChannelConfigPath: setup.ChannelConfig, SigningIdentities: []msp.SigningIdentity{adminIdentity}}
// 	txID, err := setup.admin.SaveChannel(req, resmgmt.WithOrdererEndpoint(setup.OrdererID))
// 	if err != nil || txID.TransactionID == "" {
// 		return errors.WithMessage(err, "failed to save channel")
// 	}
// 	fmt.Println("Channel created")

	// Make admin user join the previously created channel
// 	if err = setup.admin.JoinChannel(setup.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(setup.OrdererID)); err != nil {
// 		return errors.WithMessage(err, "failed to make admin join channel")
// 	}
//	fmt.Println("Channel joined")

	fmt.Println("Initialization Successful")
	setup.initialized = true
	return nil
}

// func (setup *FabricSetup) InstallAndInstantiateCC() error {
// 	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath, setup.ChaincodeGoPath)
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to create chaincode package")
// 	}
// 	fmt.Println("ccpkg created")
// 	installCCReq := resmgmt.InstallCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodePath, Version: "0", Package: ccPkg}
// 	_, err = setup.admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to install chaincode")
// 	}
// 	fmt.Println("Chaincode installed")
// 	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.hf.chainhero.io"})
// 	resp, err := setup.admin.InstantiateCC(setup.ChannelID, resmgmt.InstantiateCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodeGoPath, Version: "0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy})
// 	if err != nil || resp.TransactionID =="" {
// 		return errors.WithMessage(err, "failed to instantiate the chaincode")
// 	}
// 	fmt.Println("Chaincode instantiated")
// 	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
// 	setup.client, err = channel.New(clientContext)
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to")
// 	}
// 	fmt.Println("Channel client created")
// 	setup.event, err = event.New(clientContext)
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to create new event client")
// 	}
// 	fmt.Println("Event client created")
// 	fmt.Println("Chaincode Install & Instantiate")
// 	return nil
// }

func (setup *FabricSetup) CloseSDK() {
 	setup.sdk.Close()
}

var DefaultChannelOpts = Opts{
	Attempts:       DefaultAttempts,
	InitialBackoff: DefaultInitialBackoff,
	MaxBackoff:     DefaultMaxBackoff,
	BackoffFactor:  DefaultBackoffFactor,
	RetryableCodes: ChannelClientRetryableCodes,
}

func queryCC(client *channel.Client, t *T, targetEndpoints ...string) []byte {
	response, err := client.Query(channel.Request{ChaincodeID: "marbles", Fcn: "invoke", Args: ExampleCCQueryArgs()},
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(targetEndpoints...),
	)
	if err != nil {
		fmt.Println("Failed to query funds: %s", err)
	}
	return response.Payload
}

type Opts struct {
	// Attempts the number retry attempts
	Attempts int
	// InitialBackoff the backoff interval for the first retry attempt
	InitialBackoff time.Duration
	// MaxBackoff the maximum backoff interval for any retry attempt
	MaxBackoff time.Duration
	// BackoffFactor the factor by which the InitialBackoff is exponentially
	// incremented for consecutive retry attempts.
	// For example, a backoff factor of 2.5 will result in a backoff of
	// InitialBackoff * 2.5 * 2.5 on the second attempt.
	BackoffFactor float64
	// RetryableCodes defines the status codes, mapped by group, returned by fabric-sdk-go
	// that warrant a retry. This will default to retry.DefaultRetryableCodes.
	RetryableCodes map[Group][]Code
}

type Group int32
type Code uint32
// DefaultBackoffFactor default backoff factor
type Status int32

const (
	// DefaultAttempts number of retry attempts made by default
	DefaultAttempts = 3
	// DefaultInitialBackoff default initial backoff
	DefaultInitialBackoff = 500 * time.Millisecond
	// DefaultMaxBackoff default maximum backoff
	DefaultMaxBackoff = 60 * time.Second
	// DefaultBackoffFactor default backoff factor
	DefaultBackoffFactor = 2.0
)

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_NOT_IMPLEMENTED          Status = 501
	Status_SERVICE_UNAVAILABLE      Status = 503
)

const (
	// UnknownStatus unknown status group
	UnknownStatus Group = iota

	// TransportStatus defines the status returned by the transport layer of
	// the connections made by fabric-sdk-go

	// GRPCTransportStatus is the status associated with requests made over
	// gRPC connections
	GRPCTransportStatus
	// HTTPTransportStatus is the status associated with requests made over HTTP
	// connections
	HTTPTransportStatus

	// ServerStatus defines the status returned by various servers that fabric-sdk-go
	// is a client to

	// EndorserServerStatus status returned by the endorser server
	EndorserServerStatus
	// EventServerStatus status returned by the eventhub
	EventServerStatus
	// OrdererServerStatus status returned by the ordering service
	OrdererServerStatus
	// FabricCAServerStatus status returned by the Fabric CA server
	FabricCAServerStatus

	// ClientStatus defines the status from responses inferred by fabric-sdk-go.
	// This could be a result of response validation performed by the SDK - for example,
	// a client status could be produced by validating endorsements

	// EndorserClientStatus status returned from the endorser client
	EndorserClientStatus
	// OrdererClientStatus status returned from the orderer client
	OrdererClientStatus
	// ClientStatus is a generic client status
	ClientStatus

	// ChaincodeStatus defines the status codes returned by chaincode
	ChaincodeStatus

	// DiscoveryServerStatus status returned by the Discovery Server
	DiscoveryServerStatus
)

const (
	// OK is returned on success.
	OK Code = 0

	// Unknown represents status codes that are uncategorized or unknown to the SDK
	Unknown Code = 1

	// ConnectionFailed is returned when a network connection attempt from the SDK fails
	ConnectionFailed Code = 2

	// EndorsementMismatch is returned when there is a mismatch in endorsements received by the SDK
	EndorsementMismatch Code = 3

	// EmptyCert is return when an empty cert is returned
	EmptyCert Code = 4

	// Timeout operation timed out
	Timeout Code = 5

	// NoPeersFound No peers were discovered/configured
	NoPeersFound Code = 6

	// MultipleErrors multiple errors occurred
	MultipleErrors Code = 7

	// SignatureVerificationFailed is when signature fails verification
	SignatureVerificationFailed Code = 8

	// MissingEndorsement is if an endoresement is missing
	MissingEndorsement Code = 9

	// QueryEndorsers error indicates that no endorser group was found that would
	// satisfy the chaincode policy
	QueryEndorsers Code = 11

	// PrematureChaincodeExecution indicates that an attempt was made to invoke a chaincode that's
	// in the process of being launched.
	PrematureChaincodeExecution Code = 21

	// ChaincodeAlreadyLaunching indicates that an attempt for multiple simultaneous invokes was made to launch chaincode
	ChaincodeAlreadyLaunching Code = 22
)

var ChannelClientRetryableCodes = map[Group][]Code{
	EndorserClientStatus: {
		ConnectionFailed, EndorsementMismatch,
		PrematureChaincodeExecution,
		ChaincodeAlreadyLaunching,
	},
	EndorserServerStatus: {
		Code(Status_SERVICE_UNAVAILABLE),
		Code(Status_INTERNAL_SERVER_ERROR),
	},
	OrdererClientStatus: {
		ConnectionFailed,
	},
	OrdererServerStatus: {
		Code(Status_SERVICE_UNAVAILABLE),
		Code(Status_INTERNAL_SERVER_ERROR),
	},
	EventServerStatus: {
		Code(TxValidationCode_DUPLICATE_TXID),
		Code(TxValidationCode_ENDORSEMENT_POLICY_FAILURE),
		Code(TxValidationCode_MVCC_READ_CONFLICT),
		Code(TxValidationCode_PHANTOM_READ_CONFLICT),
	},
	// TODO: gRPC introduced retries in v1.8.0. This can be replaced with the
	// gRPC fail fast option, once available
	GRPCTransportStatus: {
		Code(Unavailable),
	},
}

type TxValidationCode int32

const (
	TxValidationCode_VALID                        TxValidationCode = 0
	TxValidationCode_NIL_ENVELOPE                 TxValidationCode = 1
	TxValidationCode_BAD_PAYLOAD                  TxValidationCode = 2
	TxValidationCode_BAD_COMMON_HEADER            TxValidationCode = 3
	TxValidationCode_BAD_CREATOR_SIGNATURE        TxValidationCode = 4
	TxValidationCode_INVALID_ENDORSER_TRANSACTION TxValidationCode = 5
	TxValidationCode_INVALID_CONFIG_TRANSACTION   TxValidationCode = 6
	TxValidationCode_UNSUPPORTED_TX_PAYLOAD       TxValidationCode = 7
	TxValidationCode_BAD_PROPOSAL_TXID            TxValidationCode = 8
	TxValidationCode_DUPLICATE_TXID               TxValidationCode = 9
	TxValidationCode_ENDORSEMENT_POLICY_FAILURE   TxValidationCode = 10
	TxValidationCode_MVCC_READ_CONFLICT           TxValidationCode = 11
	TxValidationCode_PHANTOM_READ_CONFLICT        TxValidationCode = 12
	TxValidationCode_UNKNOWN_TX_TYPE              TxValidationCode = 13
	TxValidationCode_TARGET_CHAIN_NOT_FOUND       TxValidationCode = 14
	TxValidationCode_MARSHAL_TX_ERROR             TxValidationCode = 15
	TxValidationCode_NIL_TXACTION                 TxValidationCode = 16
	TxValidationCode_EXPIRED_CHAINCODE            TxValidationCode = 17
	TxValidationCode_CHAINCODE_VERSION_CONFLICT   TxValidationCode = 18
	TxValidationCode_BAD_HEADER_EXTENSION         TxValidationCode = 19
	TxValidationCode_BAD_CHANNEL_HEADER           TxValidationCode = 20
	TxValidationCode_BAD_RESPONSE_PAYLOAD         TxValidationCode = 21
	TxValidationCode_BAD_RWSET                    TxValidationCode = 22
	TxValidationCode_ILLEGAL_WRITESET             TxValidationCode = 23
	TxValidationCode_NOT_VALIDATED                TxValidationCode = 254
	TxValidationCode_INVALID_OTHER_REASON         TxValidationCode = 255
)

const (

	// Canceled indicates the operation was canceled (typically by the caller).
	Canceled Code = 1

	

	// InvalidArgument indicates client specified an invalid argument.
	// Note that this differs from FailedPrecondition. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	InvalidArgument Code = 3

	// DeadlineExceeded means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	DeadlineExceeded Code = 4

	// NotFound means some requested entity (e.g., file or directory) was
	// not found.
	NotFound Code = 5

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	AlreadyExists Code = 6

	// PermissionDenied indicates the caller does not have permission to
	// execute the specified operation. It must not be used for rejections
	// caused by exhausting some resource (use ResourceExhausted
	// instead for those errors). It must not be
	// used if the caller cannot be identified (use Unauthenticated
	// instead for those errors).
	PermissionDenied Code = 7

	// ResourceExhausted indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	ResourceExhausted Code = 8

	// FailedPrecondition indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FailedPrecondition, Aborted, and Unavailable:
	//  (a) Use Unavailable if the client can retry just the failing call.
	//  (b) Use Aborted if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FailedPrecondition if the client should not retry until
	//      the system state has been explicitly fixed. E.g., if an "rmdir"
	//      fails because the directory is non-empty, FailedPrecondition
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FailedPrecondition if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	FailedPrecondition Code = 9

	// Aborted indicates the operation was aborted, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FailedPrecondition,
	// Aborted, and Unavailable.
	Aborted Code = 10

	// OutOfRange means operation was attempted past the valid range.
	// E.g., seeking or reading past end of file.
	//
	// Unlike InvalidArgument, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate InvalidArgument if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OutOfRange if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FailedPrecondition and
	// OutOfRange. We recommend using OutOfRange (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OutOfRange error to detect when
	// they are done.
	OutOfRange Code = 11

	// Unimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	Unimplemented Code = 12

	// Internal errors. Means some invariants expected by underlying
	// system has been broken. If you see one of these errors,
	// something is very broken.
	Internal Code = 13

	// Unavailable indicates the service is currently unavailable.
	// This is a most likely a transient condition and may be corrected
	// by retrying with a backoff.
	//
	// See litmus test above for deciding between FailedPrecondition,
	// Aborted, and Unavailable.
	Unavailable Code = 14

	// DataLoss indicates unrecoverable data loss or corruption.
	DataLoss Code = 15

	// Unauthenticated indicates the request does not have valid
	// authentication credentials for the operation.
	Unauthenticated Code = 16

	_maxCode = 17
)
var queryArgs = [][]byte{[]byte("invoke"), []byte("initOrder")}

func ExampleCCQueryArgs() [][]byte {
	return queryArgs
}

type T struct {
	common
	isParallel bool
	context    *testContext // For running tests and subtests.
}

type testContext struct {
	match *matcher

	mu sync.Mutex

	// Channel used to signal tests that are ready to be run in parallel.
	startParallel chan bool

	// running is the number of tests currently running in parallel.
	// This does not include tests that are waiting for subtests to complete.
	running int

	// numWaiting is the number tests waiting to be run in parallel.
	numWaiting int

	// maxParallel is a copy of the parallel flag.
	maxParallel int
}

type matcher struct {
	filter    []string
	matchFunc func(pat, str string) (bool, error)

	mu       sync.Mutex
	subNames map[string]int64
}

type common struct {
	mu      sync.RWMutex        // guards this group of fields
	output  []byte              // Output generated by test or benchmark.
	w       io.Writer           // For flushToParent.
	ran     bool                // Test or benchmark (or one of its subtests) was executed.
	failed  bool                // Test or benchmark has failed.
	skipped bool                // Test of benchmark has been skipped.
	done    bool                // Test is finished and all subtests have completed.
	helpers map[string]struct{} // functions to be skipped when writing file/line info

	chatty     bool   // A copy of the chatty flag.
	finished   bool   // Test function has completed.
	hasSub     int32  // written atomically
	raceErrors int    // number of races detected during test
	runner     string // function name of tRunner running the test

	parent   *common
	level    int       // Nesting depth of test or benchmark.
	creator  []uintptr // If level > 0, the stack trace at the point where the parent called t.Run.
	name     string    // Name of test or benchmark.
	start    time.Time // Time test or benchmark started
	duration time.Duration
	barrier  chan bool // To signal parallel subtests they may start.
	signal   chan bool // To signal a test is done.
	sub      []*T      // Queue of subtests to be run in parallel.
}
