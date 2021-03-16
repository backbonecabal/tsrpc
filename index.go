import "encoding/json"
import "errors"
// Hex representation of a Keccak 256 hash
type Keccak string
// Null
type Null interface{}
// The block hash or null when its the pending block
type BlockHashOrNull struct {
	Keccak *Keccak
	Null   *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *BlockHashOrNull) UnmarshalJSON(bytes []byte) error {
	var myKeccak Keccak
	if err := json.Unmarshal(bytes, &myKeccak); err == nil {
		o.Keccak = &myKeccak
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o BlockHashOrNull) MarshalJSON() ([]byte, error) {
	if o.Keccak != nil {
		return json.Marshal(o.Keccak)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// The hex representation of the block's height
type BlockNumber string
// The block number or null when its the pending block
type BlockNumberOrNull struct {
	BlockNumber *BlockNumber
	Null        *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *BlockNumberOrNull) UnmarshalJSON(bytes []byte) error {
	var myBlockNumber BlockNumber
	if err := json.Unmarshal(bytes, &myBlockNumber); err == nil {
		o.BlockNumber = &myBlockNumber
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o BlockNumberOrNull) MarshalJSON() ([]byte, error) {
	if o.BlockNumber != nil {
		return json.Marshal(o.BlockNumber)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// The sender of the transaction
type From string
// The gas limit provided by the sender in Wei
type TransactionGas string
// The gas price willing to be paid by the sender in Wei
type TransactionGasPrice string
// Keccak 256 Hash of the RLP encoding of a transaction
type TransactionHash string
// The data field sent with the transaction
type TransactionInput string
// The total number of prior transactions made by the sender
type TransactionNonce string
type Address string
// Destination address of the transaction. Null if it was a contract create.
type To struct {
	Address *Address
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *To) UnmarshalJSON(bytes []byte) error {
	var myAddress Address
	if err := json.Unmarshal(bytes, &myAddress); err == nil {
		o.Address = &myAddress
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o To) MarshalJSON() ([]byte, error) {
	if o.Address != nil {
		return json.Marshal(o.Address)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Hex representation of the integer
type Integer string
// The index of the transaction. null when its pending
type TransactionIndex struct {
	Integer *Integer
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *TransactionIndex) UnmarshalJSON(bytes []byte) error {
	var myInteger Integer
	if err := json.Unmarshal(bytes, &myInteger); err == nil {
		o.Integer = &myInteger
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o TransactionIndex) MarshalJSON() ([]byte, error) {
	if o.Integer != nil {
		return json.Marshal(o.Integer)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Value of Ether being transferred in Wei
type TransactionValue string
// ECDSA recovery id
type TransactionSigV string
// ECDSA signature r
type TransactionSigR string
// ECDSA signature s
type TransactionSigS string
// The optional block height description
type BlockNumberTag string
const (
	BlockNumberTagEnum0 BlockNumberTag = "earliest"
	BlockNumberTagEnum1 BlockNumberTag = "latest"
	BlockNumberTagEnum2 BlockNumberTag = "pending"
)
// List of contract addresses from which to monitor events
type Addresses []Address
type OneOrArrayOfAddresses struct {
	Address   *Address
	Addresses *Addresses
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *OneOrArrayOfAddresses) UnmarshalJSON(bytes []byte) error {
	var myAddress Address
	if err := json.Unmarshal(bytes, &myAddress); err == nil {
		o.Address = &myAddress
		return nil
	}
	var myAddresses Addresses
	if err := json.Unmarshal(bytes, &myAddresses); err == nil {
		o.Addresses = &myAddresses
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o OneOrArrayOfAddresses) MarshalJSON() ([]byte, error) {
	if o.Address != nil {
		return json.Marshal(o.Address)
	}
	if o.Addresses != nil {
		return json.Marshal(o.Addresses)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// 32 Bytes DATA of indexed log arguments. (In solidity: The first topic is the hash of the signature of the event (e.g. Deposit(address,bytes32,uint256))
type Topic string
// Topics are order-dependent. Each topic can also be an array of DATA with 'or' options.
type LogTopics []Topic
// The key used to get the storage slot in its account tree.
type StorageProofKey string
// The hex representation of the Keccak 256 of the RLP encoded block
type BlockHash string
// A number only to be used once
type Nonce string
type NonceOrNull struct {
	Nonce *Nonce
	Null  *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *NonceOrNull) UnmarshalJSON(bytes []byte) error {
	var myNonce Nonce
	if err := json.Unmarshal(bytes, &myNonce); err == nil {
		o.Nonce = &myNonce
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o NonceOrNull) MarshalJSON() ([]byte, error) {
	if o.Nonce != nil {
		return json.Marshal(o.Nonce)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Keccak hash of the uncles data in the block
type BlockShaUncles string
// The bloom filter for the logs of the block or null when its the pending block
type BlockLogsBloom string
// The root of the transactions trie of the block.
type BlockTransactionsRoot string
// The root of the final state trie of the block
type BlockStateRoot string
// The root of the receipts trie of the block
type BlockReceiptsRoot string
type AddressOrNull struct {
	Address *Address
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *AddressOrNull) UnmarshalJSON(bytes []byte) error {
	var myAddress Address
	if err := json.Unmarshal(bytes, &myAddress); err == nil {
		o.Address = &myAddress
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o AddressOrNull) MarshalJSON() ([]byte, error) {
	if o.Address != nil {
		return json.Marshal(o.Address)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Integer of the difficulty for this block
type BlockDifficulty string
// Integer of the total difficulty of the chain until this block
type BlockTotalDifficulty struct {
	Integer *Integer
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *BlockTotalDifficulty) UnmarshalJSON(bytes []byte) error {
	var myInteger Integer
	if err := json.Unmarshal(bytes, &myInteger); err == nil {
		o.Integer = &myInteger
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o BlockTotalDifficulty) MarshalJSON() ([]byte, error) {
	if o.Integer != nil {
		return json.Marshal(o.Integer)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// The 'extra data' field of this block
type BlockExtraData string
// Integer the size of this block in bytes
type BlockSize string
// The maximum gas allowed in this block
type BlockGasLimit string
// The total used gas by all transactions in this block
type BlockGasUsed string
// The unix timestamp for when the block was collated
type BlockTimeStamp string
type Transaction struct {
	BlockHash        *BlockHashOrNull     `json:"blockHash,omitempty"`
	BlockNumber      *BlockNumberOrNull   `json:"blockNumber,omitempty"`
	From             *From                `json:"from,omitempty"`
	Gas              *TransactionGas      `json:"gas"`
	GasPrice         *TransactionGasPrice `json:"gasPrice"`
	Hash             *TransactionHash     `json:"hash,omitempty"`
	Input            *TransactionInput    `json:"input,omitempty"`
	Nonce            *TransactionNonce    `json:"nonce"`
	To               *To                  `json:"to,omitempty"`
	TransactionIndex *TransactionIndex    `json:"transactionIndex,omitempty"`
	Value            *TransactionValue    `json:"value,omitempty"`
	V                *TransactionSigV     `json:"v,omitempty"`
	R                *TransactionSigR     `json:"r,omitempty"`
	S                *TransactionSigS     `json:"s,omitempty"`
}
type TransactionOrTransactionHash struct {
	Transaction     *Transaction
	TransactionHash *TransactionHash
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *TransactionOrTransactionHash) UnmarshalJSON(bytes []byte) error {
	var myTransaction Transaction
	if err := json.Unmarshal(bytes, &myTransaction); err == nil {
		o.Transaction = &myTransaction
		return nil
	}
	var myTransactionHash TransactionHash
	if err := json.Unmarshal(bytes, &myTransactionHash); err == nil {
		o.TransactionHash = &myTransactionHash
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o TransactionOrTransactionHash) MarshalJSON() ([]byte, error) {
	if o.Transaction != nil {
		return json.Marshal(o.Transaction)
	}
	if o.TransactionHash != nil {
		return json.Marshal(o.TransactionHash)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
type TransactionsOrHashes []TransactionOrTransactionHash
// Block hash of the RLP encoding of an uncle block
type UncleHash string
// Array of uncle hashes
type UncleHashes []UncleHash
// The Block is the collection of relevant pieces of information (known as the block header), together with information corresponding to the comprised transactions, and a set of other block headers that are known to have a parent equal to the present block’s parent’s parent.
type Block struct {
	Number           *BlockNumberOrNull     `json:"number,omitempty"`
	Hash             *BlockHashOrNull       `json:"hash,omitempty"`
	ParentHash       *BlockHash             `json:"parentHash,omitempty"`
	Nonce            *NonceOrNull           `json:"nonce,omitempty"`
	Sha3Uncles       *BlockShaUncles        `json:"sha3Uncles,omitempty"`
	LogsBloom        *BlockLogsBloom        `json:"logsBloom,omitempty"`
	TransactionsRoot *BlockTransactionsRoot `json:"transactionsRoot,omitempty"`
	StateRoot        *BlockStateRoot        `json:"stateRoot,omitempty"`
	ReceiptsRoot     *BlockReceiptsRoot     `json:"receiptsRoot,omitempty"`
	Miner            *AddressOrNull         `json:"miner,omitempty"`
	Difficulty       *BlockDifficulty       `json:"difficulty,omitempty"`
	TotalDifficulty  *BlockTotalDifficulty  `json:"totalDifficulty,omitempty"`
	ExtraData        *BlockExtraData        `json:"extraData,omitempty"`
	Size             *BlockSize             `json:"size,omitempty"`
	GasLimit         *BlockGasLimit         `json:"gasLimit,omitempty"`
	GasUsed          *BlockGasUsed          `json:"gasUsed,omitempty"`
	Timestamp        *BlockTimeStamp        `json:"timestamp,omitempty"`
	Transactions     *TransactionsOrHashes  `json:"transactions,omitempty"`
	Uncles           *UncleHashes           `json:"uncles,omitempty"`
}
// Sender of the transaction
type LogAddress string
// The data/input string sent along with the transaction
type LogData string
// The index of the event within its transaction, null when its pending
type LogIndex string
// Whether or not the log was orphaned off the main chain
type LogIsRemoved bool
// An indexed event generated during a transaction
type Log struct {
	Address          *LogAddress       `json:"address,omitempty"`
	BlockHash        *BlockHash        `json:"blockHash,omitempty"`
	BlockNumber      *BlockNumber      `json:"blockNumber,omitempty"`
	Data             *LogData          `json:"data,omitempty"`
	LogIndex         *LogIndex         `json:"logIndex,omitempty"`
	Removed          *LogIsRemoved     `json:"removed,omitempty"`
	Topics           *LogTopics        `json:"topics,omitempty"`
	TransactionHash  *TransactionHash  `json:"transactionHash,omitempty"`
	TransactionIndex *TransactionIndex `json:"transactionIndex,omitempty"`
}
// The contract address created, if the transaction was a contract creation, otherwise null
type ReceiptContractAddress struct {
	Address *Address
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ReceiptContractAddress) UnmarshalJSON(bytes []byte) error {
	var myAddress Address
	if err := json.Unmarshal(bytes, &myAddress); err == nil {
		o.Address = &myAddress
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ReceiptContractAddress) MarshalJSON() ([]byte, error) {
	if o.Address != nil {
		return json.Marshal(o.Address)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// The gas units used by the transaction
type ReceiptCumulativeGasUsed string
// The total gas used by the transaction
type ReceiptGasUsed string
// An array of all the logs triggered during the transaction
type Logs []Log
// A 2048 bit bloom filter from the logs of the transaction. Each log sets 3 bits though taking the low-order 11 bits of each of the first three pairs of bytes in a Keccak 256 hash of the log's byte series
type BloomFilter string
// The intermediate stateRoot directly after transaction execution.
type ReceiptPostTransactionState string
// Whether or not the transaction threw an error.
type ReceiptStatus bool
// The receipt of a transaction
type Receipt struct {
	BlockHash            *BlockHash                   `json:"blockHash"`
	BlockNumber          *BlockNumber                 `json:"blockNumber"`
	ContractAddress      *ReceiptContractAddress      `json:"contractAddress"`
	CumulativeGasUsed    *ReceiptCumulativeGasUsed    `json:"cumulativeGasUsed"`
	From                 *From                        `json:"from"`
	GasUsed              *ReceiptGasUsed              `json:"gasUsed"`
	Logs                 *Logs                        `json:"logs"`
	LogsBloom            *BloomFilter                 `json:"logsBloom"`
	To                   *To                          `json:"to"`
	TransactionHash      *TransactionHash             `json:"transactionHash"`
	TransactionIndex     *TransactionIndex            `json:"transactionIndex"`
	PostTransactionState *ReceiptPostTransactionState `json:"postTransactionState,omitempty"`
	Status               *ReceiptStatus               `json:"status,omitempty"`
}
// The address of the account or contract of the request
type ProofAccountAddress string
// An individual node used to prove a path down a merkle-patricia-tree
type ProofNode string
// The set of node values needed to traverse a patricia merkle tree (from root to leaf) to retrieve a value
type ProofNodes []ProofNode
// The Ether balance of the account or contract of the request
type ProofAccountBalance string
// The code hash of the contract of the request (keccak(NULL) if external account)
type ProofAccountCodeHash string
// The transaction count of the account or contract of the request
type ProofAccountNonce string
// The storage hash of the contract of the request (keccak(rlp(NULL)) if external account)
type ProofAccountStorageHash string
// The value of the storage slot in its account tree
type StorageProofValue string
// Object proving a relationship of a storage value to an account's storageHash.
type StorageProof struct {
	Key   *StorageProofKey   `json:"key,omitempty"`
	Value *StorageProofValue `json:"value,omitempty"`
	Proof *ProofNodes        `json:"proof,omitempty"`
}
// Current block header PoW hash.
type StorageProofSet []StorageProof
// The merkle proofs of the specified account connecting them to the blockhash of the block specified
type ProofAccount struct {
	Address      *ProofAccountAddress     `json:"address,omitempty"`
	AccountProof *ProofNodes              `json:"accountProof,omitempty"`
	Balance      *ProofAccountBalance     `json:"balance,omitempty"`
	CodeHash     *ProofAccountCodeHash    `json:"codeHash,omitempty"`
	Nonce        *ProofAccountNonce       `json:"nonce,omitempty"`
	StorageHash  *ProofAccountStorageHash `json:"storageHash,omitempty"`
	StorageProof *StorageProofSet         `json:"storageProof,omitempty"`
}
// Current block header PoW hash.
type PowHash string
// The seed hash used for the DAG.
type SeedHash string
// The boundary condition ('target'), 2^256 / difficulty.
type Difficulty string
// Block at which the import started (will only be reset, after the sync reached his head)
type SyncingDataStartingBlock string
// The current block, same as eth_blockNumber
type SyncingDataCurrentBlock string
// The estimated highest block
type SyncingDataHighestBlock string
// The known states
type SyncingDataKnownStates string
// The pulled states
type SyncingDataPulledStates string
// An object with sync status data
type SyncingData struct {
	StartingBlock *SyncingDataStartingBlock `json:"startingBlock,omitempty"`
	CurrentBlock  *SyncingDataCurrentBlock  `json:"currentBlock,omitempty"`
	HighestBlock  *SyncingDataHighestBlock  `json:"highestBlock,omitempty"`
	KnownStates   *SyncingDataKnownStates   `json:"knownStates,omitempty"`
	PulledStates  *SyncingDataPulledStates  `json:"pulledStates,omitempty"`
}
type BooleanVyG3AETh bool
type Data string
type BlockNumberOrTag struct {
	BlockNumber    *BlockNumber
	BlockNumberTag *BlockNumberTag
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *BlockNumberOrTag) UnmarshalJSON(bytes []byte) error {
	var myBlockNumber BlockNumber
	if err := json.Unmarshal(bytes, &myBlockNumber); err == nil {
		o.BlockNumber = &myBlockNumber
		return nil
	}
	var myBlockNumberTag BlockNumberTag
	if err := json.Unmarshal(bytes, &myBlockNumberTag); err == nil {
		o.BlockNumberTag = &myBlockNumberTag
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o BlockNumberOrTag) MarshalJSON() ([]byte, error) {
	if o.BlockNumber != nil {
		return json.Marshal(o.BlockNumber)
	}
	if o.BlockNumberTag != nil {
		return json.Marshal(o.BlockNumberTag)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type IsTransactionsIncluded bool
// An identifier used to reference the filter.
type FilterId string
// A filter used to monitor the blockchain for log/events
type Filter struct {
	FromBlock *BlockNumber           `json:"fromBlock,omitempty"`
	ToBlock   *BlockNumber           `json:"toBlock,omitempty"`
	Address   *OneOrArrayOfAddresses `json:"address,omitempty"`
	Topics    *LogTopics             `json:"topics,omitempty"`
}
// Hex representation of the storage slot where the variable exists
type Position string
// A storage key is indexed from the solidity compiler by the order it is declared. For mappings it uses the keccak of the mapping key with its position (and recursively for X-dimensional mappings)
type StorageKeys interface{}
// Hex representation of a variable length byte array
type Bytes string
// Hex representation of a 256 bit unit of data
type DataWord string
// The mix digest.
type MixHash string
type ClientVersion string
type IsNetListening bool
// Hex representation of number of connected peers
type NumConnectedPeers string
type NetworkId string
type ChainId string
// Integer of the current gas price
type GasPriceResult string
type IntegerOrNull struct {
	Integer *Integer
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *IntegerOrNull) UnmarshalJSON(bytes []byte) error {
	var myInteger Integer
	if err := json.Unmarshal(bytes, &myInteger); err == nil {
		o.Integer = &myInteger
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o IntegerOrNull) MarshalJSON() ([]byte, error) {
	if o.Integer != nil {
		return json.Marshal(o.Integer)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type BlockOrNull struct {
	Block *Block
	Null  *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *BlockOrNull) UnmarshalJSON(bytes []byte) error {
	var myBlock Block
	if err := json.Unmarshal(bytes, &myBlock); err == nil {
		o.Block = &myBlock
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o BlockOrNull) MarshalJSON() ([]byte, error) {
	if o.Block != nil {
		return json.Marshal(o.Block)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type LogResult []Log
type SetOfLogs []Log
type TransactionOrNull struct {
	Transaction *Transaction
	Null        *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *TransactionOrNull) UnmarshalJSON(bytes []byte) error {
	var myTransaction Transaction
	if err := json.Unmarshal(bytes, &myTransaction); err == nil {
		o.Transaction = &myTransaction
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o TransactionOrNull) MarshalJSON() ([]byte, error) {
	if o.Transaction != nil {
		return json.Marshal(o.Transaction)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type TransactionReceiptOrNull struct {
	Receipt *Receipt
	Null    *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *TransactionReceiptOrNull) UnmarshalJSON(bytes []byte) error {
	var myReceipt Receipt
	if err := json.Unmarshal(bytes, &myReceipt); err == nil {
		o.Receipt = &myReceipt
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o TransactionReceiptOrNull) MarshalJSON() ([]byte, error) {
	if o.Receipt != nil {
		return json.Marshal(o.Receipt)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type ProofAccountOrNull struct {
	ProofAccount *ProofAccount
	Null         *Null
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *ProofAccountOrNull) UnmarshalJSON(bytes []byte) error {
	var myProofAccount ProofAccount
	if err := json.Unmarshal(bytes, &myProofAccount); err == nil {
		o.ProofAccount = &myProofAccount
		return nil
	}
	var myNull Null
	if err := json.Unmarshal(bytes, &myNull); err == nil {
		o.Null = &myNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o ProofAccountOrNull) MarshalJSON() ([]byte, error) {
	if o.ProofAccount != nil {
		return json.Marshal(o.ProofAccount)
	}
	if o.Null != nil {
		return json.Marshal(o.Null)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
type GetWorkResults (PowHash, SeedHash, Difficulty)
// An array of transactions
type Transactions []Transaction
type IsSyncingResult struct {
	SyncingData     *SyncingData
	BooleanVyG3AETh *BooleanVyG3AETh
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *IsSyncingResult) UnmarshalJSON(bytes []byte) error {
	var mySyncingData SyncingData
	if err := json.Unmarshal(bytes, &mySyncingData); err == nil {
		o.SyncingData = &mySyncingData
		return nil
	}
	var myBooleanVyG3AETh BooleanVyG3AETh
	if err := json.Unmarshal(bytes, &myBooleanVyG3AETh); err == nil {
		o.BooleanVyG3AETh = &myBooleanVyG3AETh
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o IsSyncingResult) MarshalJSON() ([]byte, error) {
	if o.SyncingData != nil {
		return json.Marshal(o.SyncingData)
	}
	if o.BooleanVyG3AETh != nil {
		return json.Marshal(o.BooleanVyG3AETh)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// Generated! Represents an alias to any of the provided schemas
type AnyOfDataTransactionBlockNumberOrTagTransactionAddressBlockNumberBlockHashIsTransactionsIncludedBlockNumberOrTagIsTransactionsIncludedBlockHashBlockNumberOrTagAddressBlockNumberFilterIdFilterIdTransactionHashBlockHashIntegerBlockNumberOrTagIntegerFilterAddressPositionBlockNumberOrTagBlockHashIntegerBlockNumberOrTagIntegerTransactionHashAddressBlockNumberOrTagTransactionHashBlockHashIntegerBlockNumberIntegerBlockHashBlockNumberOrTagAddressStorageKeysBlockNumberOrTagFilterBytesDataWordDataWordNoncePowHashMixHashFilterIdClientVersionKeccakIsNetListeningNumConnectedPeersNetworkIdBlockNumberOrTagBytesChainIdAddressIntegerGasPriceResultIntegerOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullBytesLogResultSetOfLogsBytesBytesBytesSetOfLogsDataWordTransactionOrNullTransactionOrNullTransactionOrNullNonceOrNullTransactionReceiptOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullProofAccountOrNullGetWorkResultsIntegerBooleanVyG3AEThFilterIdIntegerFilterIdTransactionsIntegerKeccakBooleanVyG3AEThBooleanVyG3AEThIsSyncingResultBooleanVyG3AETh struct {
	Data                     *Data
	Transaction              *Transaction
	BlockNumberOrTag         *BlockNumberOrTag
	Address                  *Address
	BlockNumber              *BlockNumber
	BlockHash                *BlockHash
	IsTransactionsIncluded   *IsTransactionsIncluded
	FilterId                 *FilterId
	TransactionHash          *TransactionHash
	Integer                  *Integer
	Filter                   *Filter
	Position                 *Position
	StorageKeys              *StorageKeys
	Bytes                    *Bytes
	DataWord                 *DataWord
	Nonce                    *Nonce
	PowHash                  *PowHash
	MixHash                  *MixHash
	ClientVersion            *ClientVersion
	Keccak                   *Keccak
	IsNetListening           *IsNetListening
	NumConnectedPeers        *NumConnectedPeers
	NetworkId                *NetworkId
	ChainId                  *ChainId
	GasPriceResult           *GasPriceResult
	IntegerOrNull            *IntegerOrNull
	BlockOrNull              *BlockOrNull
	LogResult                *LogResult
	SetOfLogs                *SetOfLogs
	TransactionOrNull        *TransactionOrNull
	NonceOrNull              *NonceOrNull
	TransactionReceiptOrNull *TransactionReceiptOrNull
	ProofAccountOrNull       *ProofAccountOrNull
	GetWorkResults           *GetWorkResults
	BooleanVyG3AETh          *BooleanVyG3AETh
	Transactions             *Transactions
	IsSyncingResult          *IsSyncingResult
}
func (a *AnyOfDataTransactionBlockNumberOrTagTransactionAddressBlockNumberBlockHashIsTransactionsIncludedBlockNumberOrTagIsTransactionsIncludedBlockHashBlockNumberOrTagAddressBlockNumberFilterIdFilterIdTransactionHashBlockHashIntegerBlockNumberOrTagIntegerFilterAddressPositionBlockNumberOrTagBlockHashIntegerBlockNumberOrTagIntegerTransactionHashAddressBlockNumberOrTagTransactionHashBlockHashIntegerBlockNumberIntegerBlockHashBlockNumberOrTagAddressStorageKeysBlockNumberOrTagFilterBytesDataWordDataWordNoncePowHashMixHashFilterIdClientVersionKeccakIsNetListeningNumConnectedPeersNetworkIdBlockNumberOrTagBytesChainIdAddressIntegerGasPriceResultIntegerOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullBytesLogResultSetOfLogsBytesBytesBytesSetOfLogsDataWordTransactionOrNullTransactionOrNullTransactionOrNullNonceOrNullTransactionReceiptOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullProofAccountOrNullGetWorkResultsIntegerBooleanVyG3AEThFilterIdIntegerFilterIdTransactionsIntegerKeccakBooleanVyG3AEThBooleanVyG3AEThIsSyncingResultBooleanVyG3AETh) UnmarshalJSON(bytes []byte) error {
	var ok bool
	var myData Data
	if err := json.Unmarshal(bytes, &myData); err == nil {
		ok = true
		a.Data = &myData
	}
	var myTransaction Transaction
	if err := json.Unmarshal(bytes, &myTransaction); err == nil {
		ok = true
		a.Transaction = &myTransaction
	}
	var myBlockNumberOrTag BlockNumberOrTag
	if err := json.Unmarshal(bytes, &myBlockNumberOrTag); err == nil {
		ok = true
		a.BlockNumberOrTag = &myBlockNumberOrTag
	}
	var myAddress Address
	if err := json.Unmarshal(bytes, &myAddress); err == nil {
		ok = true
		a.Address = &myAddress
	}
	var myBlockNumber BlockNumber
	if err := json.Unmarshal(bytes, &myBlockNumber); err == nil {
		ok = true
		a.BlockNumber = &myBlockNumber
	}
	var myBlockHash BlockHash
	if err := json.Unmarshal(bytes, &myBlockHash); err == nil {
		ok = true
		a.BlockHash = &myBlockHash
	}
	var myIsTransactionsIncluded IsTransactionsIncluded
	if err := json.Unmarshal(bytes, &myIsTransactionsIncluded); err == nil {
		ok = true
		a.IsTransactionsIncluded = &myIsTransactionsIncluded
	}
	var myFilterId FilterId
	if err := json.Unmarshal(bytes, &myFilterId); err == nil {
		ok = true
		a.FilterId = &myFilterId
	}
	var myTransactionHash TransactionHash
	if err := json.Unmarshal(bytes, &myTransactionHash); err == nil {
		ok = true
		a.TransactionHash = &myTransactionHash
	}
	var myInteger Integer
	if err := json.Unmarshal(bytes, &myInteger); err == nil {
		ok = true
		a.Integer = &myInteger
	}
	var myFilter Filter
	if err := json.Unmarshal(bytes, &myFilter); err == nil {
		ok = true
		a.Filter = &myFilter
	}
	var myPosition Position
	if err := json.Unmarshal(bytes, &myPosition); err == nil {
		ok = true
		a.Position = &myPosition
	}
	var myStorageKeys StorageKeys
	if err := json.Unmarshal(bytes, &myStorageKeys); err == nil {
		ok = true
		a.StorageKeys = &myStorageKeys
	}
	var myBytes Bytes
	if err := json.Unmarshal(bytes, &myBytes); err == nil {
		ok = true
		a.Bytes = &myBytes
	}
	var myDataWord DataWord
	if err := json.Unmarshal(bytes, &myDataWord); err == nil {
		ok = true
		a.DataWord = &myDataWord
	}
	var myNonce Nonce
	if err := json.Unmarshal(bytes, &myNonce); err == nil {
		ok = true
		a.Nonce = &myNonce
	}
	var myPowHash PowHash
	if err := json.Unmarshal(bytes, &myPowHash); err == nil {
		ok = true
		a.PowHash = &myPowHash
	}
	var myMixHash MixHash
	if err := json.Unmarshal(bytes, &myMixHash); err == nil {
		ok = true
		a.MixHash = &myMixHash
	}
	var myClientVersion ClientVersion
	if err := json.Unmarshal(bytes, &myClientVersion); err == nil {
		ok = true
		a.ClientVersion = &myClientVersion
	}
	var myKeccak Keccak
	if err := json.Unmarshal(bytes, &myKeccak); err == nil {
		ok = true
		a.Keccak = &myKeccak
	}
	var myIsNetListening IsNetListening
	if err := json.Unmarshal(bytes, &myIsNetListening); err == nil {
		ok = true
		a.IsNetListening = &myIsNetListening
	}
	var myNumConnectedPeers NumConnectedPeers
	if err := json.Unmarshal(bytes, &myNumConnectedPeers); err == nil {
		ok = true
		a.NumConnectedPeers = &myNumConnectedPeers
	}
	var myNetworkId NetworkId
	if err := json.Unmarshal(bytes, &myNetworkId); err == nil {
		ok = true
		a.NetworkId = &myNetworkId
	}
	var myChainId ChainId
	if err := json.Unmarshal(bytes, &myChainId); err == nil {
		ok = true
		a.ChainId = &myChainId
	}
	var myGasPriceResult GasPriceResult
	if err := json.Unmarshal(bytes, &myGasPriceResult); err == nil {
		ok = true
		a.GasPriceResult = &myGasPriceResult
	}
	var myIntegerOrNull IntegerOrNull
	if err := json.Unmarshal(bytes, &myIntegerOrNull); err == nil {
		ok = true
		a.IntegerOrNull = &myIntegerOrNull
	}
	var myBlockOrNull BlockOrNull
	if err := json.Unmarshal(bytes, &myBlockOrNull); err == nil {
		ok = true
		a.BlockOrNull = &myBlockOrNull
	}
	var myLogResult LogResult
	if err := json.Unmarshal(bytes, &myLogResult); err == nil {
		ok = true
		a.LogResult = &myLogResult
	}
	var mySetOfLogs SetOfLogs
	if err := json.Unmarshal(bytes, &mySetOfLogs); err == nil {
		ok = true
		a.SetOfLogs = &mySetOfLogs
	}
	var myTransactionOrNull TransactionOrNull
	if err := json.Unmarshal(bytes, &myTransactionOrNull); err == nil {
		ok = true
		a.TransactionOrNull = &myTransactionOrNull
	}
	var myNonceOrNull NonceOrNull
	if err := json.Unmarshal(bytes, &myNonceOrNull); err == nil {
		ok = true
		a.NonceOrNull = &myNonceOrNull
	}
	var myTransactionReceiptOrNull TransactionReceiptOrNull
	if err := json.Unmarshal(bytes, &myTransactionReceiptOrNull); err == nil {
		ok = true
		a.TransactionReceiptOrNull = &myTransactionReceiptOrNull
	}
	var myProofAccountOrNull ProofAccountOrNull
	if err := json.Unmarshal(bytes, &myProofAccountOrNull); err == nil {
		ok = true
		a.ProofAccountOrNull = &myProofAccountOrNull
	}
	var myGetWorkResults GetWorkResults
	if err := json.Unmarshal(bytes, &myGetWorkResults); err == nil {
		ok = true
		a.GetWorkResults = &myGetWorkResults
	}
	var myBooleanVyG3AETh BooleanVyG3AETh
	if err := json.Unmarshal(bytes, &myBooleanVyG3AETh); err == nil {
		ok = true
		a.BooleanVyG3AETh = &myBooleanVyG3AETh
	}
	var myTransactions Transactions
	if err := json.Unmarshal(bytes, &myTransactions); err == nil {
		ok = true
		a.Transactions = &myTransactions
	}
	var myIsSyncingResult IsSyncingResult
	if err := json.Unmarshal(bytes, &myIsSyncingResult); err == nil {
		ok = true
		a.IsSyncingResult = &myIsSyncingResult
	}
	if ok {
		return nil
	}
	return errors.New("failed to unmarshal any of the object properties")
}
func (o AnyOfDataTransactionBlockNumberOrTagTransactionAddressBlockNumberBlockHashIsTransactionsIncludedBlockNumberOrTagIsTransactionsIncludedBlockHashBlockNumberOrTagAddressBlockNumberFilterIdFilterIdTransactionHashBlockHashIntegerBlockNumberOrTagIntegerFilterAddressPositionBlockNumberOrTagBlockHashIntegerBlockNumberOrTagIntegerTransactionHashAddressBlockNumberOrTagTransactionHashBlockHashIntegerBlockNumberIntegerBlockHashBlockNumberOrTagAddressStorageKeysBlockNumberOrTagFilterBytesDataWordDataWordNoncePowHashMixHashFilterIdClientVersionKeccakIsNetListeningNumConnectedPeersNetworkIdBlockNumberOrTagBytesChainIdAddressIntegerGasPriceResultIntegerOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullBytesLogResultSetOfLogsBytesBytesBytesSetOfLogsDataWordTransactionOrNullTransactionOrNullTransactionOrNullNonceOrNullTransactionReceiptOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullProofAccountOrNullGetWorkResultsIntegerBooleanVyG3AEThFilterIdIntegerFilterIdTransactionsIntegerKeccakBooleanVyG3AEThBooleanVyG3AEThIsSyncingResultBooleanVyG3AETh) MarshalJSON() ([]byte, error) {
	out := []interface{}{}
	if o.Data != nil {
		out = append(out, o.Data)
	}
	if o.Transaction != nil {
		out = append(out, o.Transaction)
	}
	if o.BlockNumberOrTag != nil {
		out = append(out, o.BlockNumberOrTag)
	}
	if o.Address != nil {
		out = append(out, o.Address)
	}
	if o.BlockNumber != nil {
		out = append(out, o.BlockNumber)
	}
	if o.BlockHash != nil {
		out = append(out, o.BlockHash)
	}
	if o.IsTransactionsIncluded != nil {
		out = append(out, o.IsTransactionsIncluded)
	}
	if o.FilterId != nil {
		out = append(out, o.FilterId)
	}
	if o.TransactionHash != nil {
		out = append(out, o.TransactionHash)
	}
	if o.Integer != nil {
		out = append(out, o.Integer)
	}
	if o.Filter != nil {
		out = append(out, o.Filter)
	}
	if o.Position != nil {
		out = append(out, o.Position)
	}
	if o.StorageKeys != nil {
		out = append(out, o.StorageKeys)
	}
	if o.Bytes != nil {
		out = append(out, o.Bytes)
	}
	if o.DataWord != nil {
		out = append(out, o.DataWord)
	}
	if o.Nonce != nil {
		out = append(out, o.Nonce)
	}
	if o.PowHash != nil {
		out = append(out, o.PowHash)
	}
	if o.MixHash != nil {
		out = append(out, o.MixHash)
	}
	if o.ClientVersion != nil {
		out = append(out, o.ClientVersion)
	}
	if o.Keccak != nil {
		out = append(out, o.Keccak)
	}
	if o.IsNetListening != nil {
		out = append(out, o.IsNetListening)
	}
	if o.NumConnectedPeers != nil {
		out = append(out, o.NumConnectedPeers)
	}
	if o.NetworkId != nil {
		out = append(out, o.NetworkId)
	}
	if o.ChainId != nil {
		out = append(out, o.ChainId)
	}
	if o.GasPriceResult != nil {
		out = append(out, o.GasPriceResult)
	}
	if o.IntegerOrNull != nil {
		out = append(out, o.IntegerOrNull)
	}
	if o.BlockOrNull != nil {
		out = append(out, o.BlockOrNull)
	}
	if o.LogResult != nil {
		out = append(out, o.LogResult)
	}
	if o.SetOfLogs != nil {
		out = append(out, o.SetOfLogs)
	}
	if o.TransactionOrNull != nil {
		out = append(out, o.TransactionOrNull)
	}
	if o.NonceOrNull != nil {
		out = append(out, o.NonceOrNull)
	}
	if o.TransactionReceiptOrNull != nil {
		out = append(out, o.TransactionReceiptOrNull)
	}
	if o.ProofAccountOrNull != nil {
		out = append(out, o.ProofAccountOrNull)
	}
	if o.GetWorkResults != nil {
		out = append(out, o.GetWorkResults)
	}
	if o.BooleanVyG3AETh != nil {
		out = append(out, o.BooleanVyG3AETh)
	}
	if o.Transactions != nil {
		out = append(out, o.Transactions)
	}
	if o.IsSyncingResult != nil {
		out = append(out, o.IsSyncingResult)
	}
	return json.Marshal(out)
}
type EthereumJSONRPC interface {
	Web3ClientVersion() (ClientVersion, error)
	Web3Sha3(data Data) (Keccak, error)
	NetListening() (IsNetListening, error)
	NetPeerCount() (NumConnectedPeers, error)
	NetVersion() (NetworkId, error)
	EthBlockNumber() (BlockNumberOrTag, error)
	EthCall(transaction Transaction, blockNumber BlockNumberOrTag) (Bytes, error)
	EthChainId() (ChainId, error)
	EthCoinbase() (Address, error)
	EthEstimateGas(transaction Transaction) (Integer, error)
	EthGasPrice() (GasPriceResult, error)
	EthGetBalance(address Address, blockNumber BlockNumber) (IntegerOrNull, error)
	EthGetBlockByHash(blockHash BlockHash, includeTransactions IsTransactionsIncluded) (BlockOrNull, error)
	EthGetBlockByNumber(blockNumber BlockNumberOrTag, includeTransactions IsTransactionsIncluded) (BlockOrNull, error)
	EthGetBlockTransactionCountByHash(blockHash BlockHash) (IntegerOrNull, error)
	EthGetBlockTransactionCountByNumber(blockNumber BlockNumberOrTag) (IntegerOrNull, error)
	EthGetCode(address Address, blockNumber BlockNumber) (Bytes, error)
	EthGetFilterChanges(filterId FilterId) (LogResult, error)
	EthGetFilterLogs(filterId FilterId) (SetOfLogs, error)
	EthGetRawTransactionByHash(transactionHash TransactionHash) (Bytes, error)
	EthGetRawTransactionByBlockHashAndIndex(blockHash BlockHash, index Integer) (Bytes, error)
	EthGetRawTransactionByBlockNumberAndIndex(blockNumber BlockNumberOrTag, index Integer) (Bytes, error)
	EthGetLogs(filter Filter) (SetOfLogs, error)
	EthGetStorageAt(address Address, key Position, blockNumber BlockNumberOrTag) (DataWord, error)
	EthGetTransactionByBlockHashAndIndex(blockHash BlockHash, index Integer) (TransactionOrNull, error)
	EthGetTransactionByBlockNumberAndIndex(blockNumber BlockNumberOrTag, index Integer) (TransactionOrNull, error)
	EthGetTransactionByHash(transactionHash TransactionHash) (TransactionOrNull, error)
	EthGetTransactionCount(address Address, blockNumber BlockNumberOrTag) (NonceOrNull, error)
	EthGetTransactionReceipt(transactionHash TransactionHash) (TransactionReceiptOrNull, error)
	EthGetUncleByBlockHashAndIndex(blockHash BlockHash, index Integer) (BlockOrNull, error)
	EthGetUncleByBlockNumberAndIndex(uncleBlockNumber BlockNumber, index Integer) (BlockOrNull, error)
	EthGetUncleCountByBlockHash(blockHash BlockHash) (IntegerOrNull, error)
	EthGetUncleCountByBlockNumber(blockNumber BlockNumberOrTag) (IntegerOrNull, error)
	EthGetProof(address Address, storageKeys StorageKeys, blockNumber BlockNumberOrTag) (ProofAccountOrNull, error)
	EthGetWork() (GetWorkResults, error)
	EthHashrate() (Integer, error)
	EthMining() (BooleanVyG3AETh, error)
	EthNewBlockFilter() (FilterId, error)
	EthNewFilter(filter Filter) (Integer, error)
	EthNewPendingTransactionFilter() (FilterId, error)
	EthPendingTransactions() (Transactions, error)
	EthProtocolVersion() (Integer, error)
	EthSendRawTransaction(signedTransactionData Bytes) (Keccak, error)
	EthSubmitHashrate(hashRate DataWord, id DataWord) (BooleanVyG3AETh, error)
	EthSubmitWork(nonce Nonce, powHash PowHash, mixHash MixHash) (BooleanVyG3AETh, error)
	EthSyncing() (IsSyncingResult, error)
	EthUninstallFilter(filterId FilterId) (BooleanVyG3AETh, error)
}