/**
 *
 * Hex representation of a Keccak 256 hash
 *
 */
export type Keccak = string;
/**
 *
 * Null
 *
 */
export type Null = null;
/**
 *
 * The block hash or null when its the pending block
 *
 */
export type BlockHashOrNull = Keccak | Null;
/**
 *
 * The hex representation of the block's height
 *
 */
export type BlockNumber = string;
/**
 *
 * The block number or null when its the pending block
 *
 */
export type BlockNumberOrNull = BlockNumber | Null;
/**
 *
 * The sender of the transaction
 *
 */
export type From = string;
/**
 *
 * The gas limit provided by the sender in Wei
 *
 */
export type TransactionGas = string;
/**
 *
 * The gas price willing to be paid by the sender in Wei
 *
 */
export type TransactionGasPrice = string;
/**
 *
 * Keccak 256 Hash of the RLP encoding of a transaction
 *
 */
export type TransactionHash = string;
/**
 *
 * The data field sent with the transaction
 *
 */
export type TransactionInput = string;
/**
 *
 * The total number of prior transactions made by the sender
 *
 */
export type TransactionNonce = string;
export type Address = string;
/**
 *
 * Destination address of the transaction. Null if it was a contract create.
 *
 */
export type To = Address | Null;
/**
 *
 * Hex representation of the integer
 *
 */
export type Integer = string;
/**
 *
 * The index of the transaction. null when its pending
 *
 */
export type TransactionIndex = Integer | Null;
/**
 *
 * Value of Ether being transferred in Wei
 *
 */
export type TransactionValue = string;
/**
 *
 * ECDSA recovery id
 *
 */
export type TransactionSigV = string;
/**
 *
 * ECDSA signature r
 *
 */
export type TransactionSigR = string;
/**
 *
 * ECDSA signature s
 *
 */
export type TransactionSigS = string;
/**
 *
 * The optional block height description
 *
 */
export type BlockNumberTag = "earliest" | "latest" | "pending";
/**
 *
 * List of contract addresses from which to monitor events
 *
 */
export type Addresses = Address[];
export type OneOrArrayOfAddresses = Address | Addresses;
/**
 *
 * 32 Bytes DATA of indexed log arguments. (In solidity: The first topic is the hash of the signature of the event (e.g. Deposit(address,bytes32,uint256))
 *
 */
export type Topic = string;
/**
 *
 * Topics are order-dependent. Each topic can also be an array of DATA with 'or' options.
 *
 */
export type LogTopics = Topic[];
/**
 *
 * The key used to get the storage slot in its account tree.
 *
 */
export type StorageProofKey = string;
/**
 *
 * The hex representation of the Keccak 256 of the RLP encoded block
 *
 */
export type BlockHash = string;
/**
 *
 * A number only to be used once
 *
 */
export type Nonce = string;
export type NonceOrNull = Nonce | Null;
/**
 *
 * Keccak hash of the uncles data in the block
 *
 */
export type BlockShaUncles = string;
/**
 *
 * The bloom filter for the logs of the block or null when its the pending block
 *
 */
export type BlockLogsBloom = string;
/**
 *
 * The root of the transactions trie of the block.
 *
 */
export type BlockTransactionsRoot = string;
/**
 *
 * The root of the final state trie of the block
 *
 */
export type BlockStateRoot = string;
/**
 *
 * The root of the receipts trie of the block
 *
 */
export type BlockReceiptsRoot = string;
export type AddressOrNull = Address | Null;
/**
 *
 * Integer of the difficulty for this block
 *
 */
export type BlockDifficulty = string;
/**
 *
 * Integer of the total difficulty of the chain until this block
 *
 */
export type BlockTotalDifficulty = Integer | Null;
/**
 *
 * The 'extra data' field of this block
 *
 */
export type BlockExtraData = string;
/**
 *
 * Integer the size of this block in bytes
 *
 */
export type BlockSize = string;
/**
 *
 * The maximum gas allowed in this block
 *
 */
export type BlockGasLimit = string;
/**
 *
 * The total used gas by all transactions in this block
 *
 */
export type BlockGasUsed = string;
/**
 *
 * The unix timestamp for when the block was collated
 *
 */
export type BlockTimeStamp = string;
export interface Transaction {
  blockHash?: BlockHashOrNull;
  blockNumber?: BlockNumberOrNull;
  from?: From;
  gas: TransactionGas;
  gasPrice: TransactionGasPrice;
  hash?: TransactionHash;
  input?: TransactionInput;
  nonce: TransactionNonce;
  to?: To;
  transactionIndex?: TransactionIndex;
  value?: TransactionValue;
  v?: TransactionSigV;
  r?: TransactionSigR;
  s?: TransactionSigS;
  [k: string]: any;
}
export type TransactionOrTransactionHash = Transaction | TransactionHash;
/**
 *
 * Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
 *
 */
export type TransactionsOrHashes = TransactionOrTransactionHash[];
/**
 *
 * Block hash of the RLP encoding of an uncle block
 *
 */
export type UncleHash = string;
/**
 *
 * Array of uncle hashes
 *
 */
export type UncleHashes = UncleHash[];
/**
 *
 * The Block is the collection of relevant pieces of information (known as the block header), together with information corresponding to the comprised transactions, and a set of other block headers that are known to have a parent equal to the present block’s parent’s parent.
 *
 */
export interface Block {
  number?: BlockNumberOrNull;
  hash?: BlockHashOrNull;
  parentHash?: BlockHash;
  nonce?: NonceOrNull;
  sha3Uncles?: BlockShaUncles;
  logsBloom?: BlockLogsBloom;
  transactionsRoot?: BlockTransactionsRoot;
  stateRoot?: BlockStateRoot;
  receiptsRoot?: BlockReceiptsRoot;
  miner?: AddressOrNull;
  difficulty?: BlockDifficulty;
  totalDifficulty?: BlockTotalDifficulty;
  extraData?: BlockExtraData;
  size?: BlockSize;
  gasLimit?: BlockGasLimit;
  gasUsed?: BlockGasUsed;
  timestamp?: BlockTimeStamp;
  transactions?: TransactionsOrHashes;
  uncles?: UncleHashes;
  [k: string]: any;
}
/**
 *
 * Sender of the transaction
 *
 */
export type LogAddress = string;
/**
 *
 * The data/input string sent along with the transaction
 *
 */
export type LogData = string;
/**
 *
 * The index of the event within its transaction, null when its pending
 *
 */
export type LogIndex = string;
/**
 *
 * Whether or not the log was orphaned off the main chain
 *
 */
export type LogIsRemoved = boolean;
/**
 *
 * An indexed event generated during a transaction
 *
 */
export interface Log {
  address?: LogAddress;
  blockHash?: BlockHash;
  blockNumber?: BlockNumber;
  data?: LogData;
  logIndex?: LogIndex;
  removed?: LogIsRemoved;
  topics?: LogTopics;
  transactionHash?: TransactionHash;
  transactionIndex?: TransactionIndex;
  [k: string]: any;
}
/**
 *
 * The contract address created, if the transaction was a contract creation, otherwise null
 *
 */
export type ReceiptContractAddress = Address | Null;
/**
 *
 * The gas units used by the transaction
 *
 */
export type ReceiptCumulativeGasUsed = string;
/**
 *
 * The total gas used by the transaction
 *
 */
export type ReceiptGasUsed = string;
/**
 *
 * An array of all the logs triggered during the transaction
 *
 */
export type Logs = Log[];
/**
 *
 * A 2048 bit bloom filter from the logs of the transaction. Each log sets 3 bits though taking the low-order 11 bits of each of the first three pairs of bytes in a Keccak 256 hash of the log's byte series
 *
 */
export type BloomFilter = string;
/**
 *
 * The intermediate stateRoot directly after transaction execution.
 *
 */
export type ReceiptPostTransactionState = string;
/**
 *
 * Whether or not the transaction threw an error.
 *
 */
export type ReceiptStatus = boolean;
/**
 *
 * The receipt of a transaction
 *
 */
export interface Receipt {
  blockHash: BlockHash;
  blockNumber: BlockNumber;
  contractAddress: ReceiptContractAddress;
  cumulativeGasUsed: ReceiptCumulativeGasUsed;
  from: From;
  gasUsed: ReceiptGasUsed;
  logs: Logs;
  logsBloom: BloomFilter;
  to: To;
  transactionHash: TransactionHash;
  transactionIndex: TransactionIndex;
  postTransactionState?: ReceiptPostTransactionState;
  status?: ReceiptStatus;
  [k: string]: any;
}
/**
 *
 * The address of the account or contract of the request
 *
 */
export type ProofAccountAddress = string;
/**
 *
 * An individual node used to prove a path down a merkle-patricia-tree
 *
 */
export type ProofNode = string;
/**
 *
 * The set of node values needed to traverse a patricia merkle tree (from root to leaf) to retrieve a value
 *
 */
export type ProofNodes = ProofNode[];
/**
 *
 * The Ether balance of the account or contract of the request
 *
 */
export type ProofAccountBalance = string;
/**
 *
 * The code hash of the contract of the request (keccak(NULL) if external account)
 *
 */
export type ProofAccountCodeHash = string;
/**
 *
 * The transaction count of the account or contract of the request
 *
 */
export type ProofAccountNonce = string;
/**
 *
 * The storage hash of the contract of the request (keccak(rlp(NULL)) if external account)
 *
 */
export type ProofAccountStorageHash = string;
/**
 *
 * The value of the storage slot in its account tree
 *
 */
export type StorageProofValue = string;
/**
 *
 * Object proving a relationship of a storage value to an account's storageHash.
 *
 */
export interface StorageProof {
  key?: StorageProofKey;
  value?: StorageProofValue;
  proof?: ProofNodes;
  [k: string]: any;
}
/**
 *
 * Current block header PoW hash.
 *
 */
export type StorageProofSet = StorageProof[];
/**
 *
 * The merkle proofs of the specified account connecting them to the blockhash of the block specified
 *
 */
export interface ProofAccount {
  address?: ProofAccountAddress;
  accountProof?: ProofNodes;
  balance?: ProofAccountBalance;
  codeHash?: ProofAccountCodeHash;
  nonce?: ProofAccountNonce;
  storageHash?: ProofAccountStorageHash;
  storageProof?: StorageProofSet;
  [k: string]: any;
}
/**
 *
 * Current block header PoW hash.
 *
 */
export type PowHash = string;
/**
 *
 * The seed hash used for the DAG.
 *
 */
export type SeedHash = string;
/**
 *
 * The boundary condition ('target'), 2^256 / difficulty.
 *
 */
export type Difficulty = string;
/**
 *
 * Block at which the import started (will only be reset, after the sync reached his head)
 *
 */
export type SyncingDataStartingBlock = string;
/**
 *
 * The current block, same as eth_blockNumber
 *
 */
export type SyncingDataCurrentBlock = string;
/**
 *
 * The estimated highest block
 *
 */
export type SyncingDataHighestBlock = string;
/**
 *
 * The known states
 *
 */
export type SyncingDataKnownStates = string;
/**
 *
 * The pulled states
 *
 */
export type SyncingDataPulledStates = string;
/**
 *
 * An object with sync status data
 *
 */
export interface SyncingData {
  startingBlock?: SyncingDataStartingBlock;
  currentBlock?: SyncingDataCurrentBlock;
  highestBlock?: SyncingDataHighestBlock;
  knownStates?: SyncingDataKnownStates;
  pulledStates?: SyncingDataPulledStates;
  [k: string]: any;
}
export type BooleanVyG3AETh = boolean;
export type Data = string;
export type BlockNumberOrTag = BlockNumber | BlockNumberTag;
export type IsTransactionsIncluded = boolean;
/**
 *
 * An identifier used to reference the filter.
 *
 */
export type FilterId = string;
/**
 *
 * A filter used to monitor the blockchain for log/events
 *
 */
export interface Filter {
  fromBlock?: BlockNumber;
  toBlock?: BlockNumber;
  address?: OneOrArrayOfAddresses;
  topics?: LogTopics;
  [k: string]: any;
}
/**
 *
 * Hex representation of the storage slot where the variable exists
 *
 */
export type Position = string;
/**
 *
 * A storage key is indexed from the solidity compiler by the order it is declared. For mappings it uses the keccak of the mapping key with its position (and recursively for X-dimensional mappings)
 *
 */
export type StorageKeys = any;
/**
 *
 * Hex representation of a variable length byte array
 *
 */
export type Bytes = string;
/**
 *
 * Hex representation of a 256 bit unit of data
 *
 */
export type DataWord = string;
/**
 *
 * The mix digest.
 *
 */
export type MixHash = string;
export type ClientVersion = string;
export type IsNetListening = boolean;
/**
 *
 * Hex representation of number of connected peers
 *
 */
export type NumConnectedPeers = string;
export type NetworkId = string;
export type ChainId = string;
/**
 *
 * Integer of the current gas price
 *
 */
export type GasPriceResult = string;
export type IntegerOrNull = Integer | Null;
export type BlockOrNull = Block | Null;
export type LogResult = Log[];
export type SetOfLogs = Log[];
export type TransactionOrNull = Transaction | Null;
export type TransactionReceiptOrNull = Receipt | Null;
export type ProofAccountOrNull = ProofAccount | Null;
export type GetWorkResults = [PowHash, SeedHash, Difficulty];
/**
 *
 * An array of transactions
 *
 */
export type Transactions = Transaction[];
export type IsSyncingResult = SyncingData | BooleanVyG3AETh;
/**
 *
 * Generated! Represents an alias to any of the provided schemas
 *
 */
export type AnyOfDataTransactionBlockNumberOrTagTransactionAddressBlockNumberBlockHashIsTransactionsIncludedBlockNumberOrTagIsTransactionsIncludedBlockHashBlockNumberOrTagAddressBlockNumberFilterIdFilterIdTransactionHashBlockHashIntegerBlockNumberOrTagIntegerFilterAddressPositionBlockNumberOrTagBlockHashIntegerBlockNumberOrTagIntegerTransactionHashAddressBlockNumberOrTagTransactionHashBlockHashIntegerBlockNumberIntegerBlockHashBlockNumberOrTagAddressStorageKeysBlockNumberOrTagFilterBytesDataWordDataWordNoncePowHashMixHashFilterIdClientVersionKeccakIsNetListeningNumConnectedPeersNetworkIdBlockNumberOrTagBytesChainIdAddressIntegerGasPriceResultIntegerOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullBytesLogResultSetOfLogsBytesBytesBytesSetOfLogsDataWordTransactionOrNullTransactionOrNullTransactionOrNullNonceOrNullTransactionReceiptOrNullBlockOrNullBlockOrNullIntegerOrNullIntegerOrNullProofAccountOrNullGetWorkResultsIntegerBooleanVyG3AEThFilterIdIntegerFilterIdTransactionsIntegerKeccakBooleanVyG3AEThBooleanVyG3AEThIsSyncingResultBooleanVyG3AETh = Data | Transaction | BlockNumberOrTag | Address | BlockNumber | BlockHash | IsTransactionsIncluded | FilterId | TransactionHash | Integer | Filter | Position | StorageKeys | Bytes | DataWord | Nonce | PowHash | MixHash | ClientVersion | Keccak | IsNetListening | NumConnectedPeers | NetworkId | ChainId | GasPriceResult | IntegerOrNull | BlockOrNull | LogResult | SetOfLogs | TransactionOrNull | NonceOrNull | TransactionReceiptOrNull | ProofAccountOrNull | GetWorkResults | BooleanVyG3AETh | Transactions | IsSyncingResult;
export type Web3ClientVersion = () => Promise<ClientVersion>;
export type Web3Sha3 = (data: Data) => Promise<Keccak>;
export type NetListening = () => Promise<IsNetListening>;
export type NetPeerCount = () => Promise<NumConnectedPeers>;
export type NetVersion = () => Promise<NetworkId>;
export type EthBlockNumber = () => Promise<BlockNumberOrTag>;
export type EthCall = (transaction: Transaction, blockNumber: BlockNumberOrTag) => Promise<Bytes>;
export type EthChainId = () => Promise<ChainId>;
export type EthCoinbase = () => Promise<Address>;
export type EthEstimateGas = (transaction: Transaction) => Promise<Integer>;
export type EthGasPrice = () => Promise<GasPriceResult>;
export type EthGetBalance = (address: Address, blockNumber: BlockNumber) => Promise<IntegerOrNull>;
export type EthGetBlockByHash = (blockHash: BlockHash, includeTransactions: IsTransactionsIncluded) => Promise<BlockOrNull>;
export type EthGetBlockByNumber = (blockNumber: BlockNumberOrTag, includeTransactions: IsTransactionsIncluded) => Promise<BlockOrNull>;
export type EthGetBlockTransactionCountByHash = (blockHash: BlockHash) => Promise<IntegerOrNull>;
export type EthGetBlockTransactionCountByNumber = (blockNumber: BlockNumberOrTag) => Promise<IntegerOrNull>;
export type EthGetCode = (address: Address, blockNumber: BlockNumber) => Promise<Bytes>;
export type EthGetFilterChanges = (filterId: FilterId) => Promise<LogResult>;
export type EthGetFilterLogs = (filterId: FilterId) => Promise<SetOfLogs>;
export type EthGetRawTransactionByHash = (transactionHash: TransactionHash) => Promise<Bytes>;
export type EthGetRawTransactionByBlockHashAndIndex = (blockHash: BlockHash, index: Integer) => Promise<Bytes>;
export type EthGetRawTransactionByBlockNumberAndIndex = (blockNumber: BlockNumberOrTag, index: Integer) => Promise<Bytes>;
export type EthGetLogs = (filter: Filter) => Promise<SetOfLogs>;
export type EthGetStorageAt = (address: Address, key: Position, blockNumber: BlockNumberOrTag) => Promise<DataWord>;
export type EthGetTransactionByBlockHashAndIndex = (blockHash: BlockHash, index: Integer) => Promise<TransactionOrNull>;
export type EthGetTransactionByBlockNumberAndIndex = (blockNumber: BlockNumberOrTag, index: Integer) => Promise<TransactionOrNull>;
export type EthGetTransactionByHash = (transactionHash: TransactionHash) => Promise<TransactionOrNull>;
export type EthGetTransactionCount = (address: Address, blockNumber: BlockNumberOrTag) => Promise<NonceOrNull>;
export type EthGetTransactionReceipt = (transactionHash: TransactionHash) => Promise<TransactionReceiptOrNull>;
export type EthGetUncleByBlockHashAndIndex = (blockHash: BlockHash, index: Integer) => Promise<BlockOrNull>;
export type EthGetUncleByBlockNumberAndIndex = (uncleBlockNumber: BlockNumber, index: Integer) => Promise<BlockOrNull>;
export type EthGetUncleCountByBlockHash = (blockHash: BlockHash) => Promise<IntegerOrNull>;
export type EthGetUncleCountByBlockNumber = (blockNumber: BlockNumberOrTag) => Promise<IntegerOrNull>;
export type EthGetProof = (address: Address, storageKeys: StorageKeys, blockNumber: BlockNumberOrTag) => Promise<ProofAccountOrNull>;
export type EthGetWork = () => Promise<GetWorkResults>;
export type EthHashrate = () => Promise<Integer>;
export type EthMining = () => Promise<BooleanVyG3AETh>;
export type EthNewBlockFilter = () => Promise<FilterId>;
export type EthNewFilter = (filter: Filter) => Promise<Integer>;
export type EthNewPendingTransactionFilter = () => Promise<FilterId>;
export type EthPendingTransactions = () => Promise<Transactions>;
export type EthProtocolVersion = () => Promise<Integer>;
export type EthSendRawTransaction = (signedTransactionData: Bytes) => Promise<Keccak>;
export type EthSubmitHashrate = (hashRate: DataWord, id: DataWord) => Promise<BooleanVyG3AETh>;
export type EthSubmitWork = (nonce: Nonce, powHash: PowHash, mixHash: MixHash) => Promise<BooleanVyG3AETh>;
export type EthSyncing = () => Promise<IsSyncingResult>;
export type EthUninstallFilter = (filterId: FilterId) => Promise<BooleanVyG3AETh>;