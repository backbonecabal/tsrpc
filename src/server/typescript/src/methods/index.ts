import web3_clientVersion from "./web3_clientVersion";
import web3_sha3 from "./web3_sha3";
import net_listening from "./net_listening";
import net_peerCount from "./net_peerCount";
import net_version from "./net_version";
import eth_blockNumber from "./eth_blockNumber";
import eth_call from "./eth_call";
import eth_chainId from "./eth_chainId";
import eth_coinbase from "./eth_coinbase";
import eth_estimateGas from "./eth_estimateGas";
import eth_gasPrice from "./eth_gasPrice";
import eth_getBalance from "./eth_getBalance";
import eth_getBlockByHash from "./eth_getBlockByHash";
import eth_getBlockByNumber from "./eth_getBlockByNumber";
import eth_getBlockTransactionCountByHash from "./eth_getBlockTransactionCountByHash";
import eth_getBlockTransactionCountByNumber from "./eth_getBlockTransactionCountByNumber";
import eth_getCode from "./eth_getCode";
import eth_getFilterChanges from "./eth_getFilterChanges";
import eth_getFilterLogs from "./eth_getFilterLogs";
import eth_getRawTransactionByHash from "./eth_getRawTransactionByHash";
import eth_getRawTransactionByBlockHashAndIndex from "./eth_getRawTransactionByBlockHashAndIndex";
import eth_getRawTransactionByBlockNumberAndIndex from "./eth_getRawTransactionByBlockNumberAndIndex";
import eth_getLogs from "./eth_getLogs";
import eth_getStorageAt from "./eth_getStorageAt";
import eth_getTransactionByBlockHashAndIndex from "./eth_getTransactionByBlockHashAndIndex";
import eth_getTransactionByBlockNumberAndIndex from "./eth_getTransactionByBlockNumberAndIndex";
import eth_getTransactionByHash from "./eth_getTransactionByHash";
import eth_getTransactionCount from "./eth_getTransactionCount";
import eth_getTransactionReceipt from "./eth_getTransactionReceipt";
import eth_getUncleByBlockHashAndIndex from "./eth_getUncleByBlockHashAndIndex";
import eth_getUncleByBlockNumberAndIndex from "./eth_getUncleByBlockNumberAndIndex";
import eth_getUncleCountByBlockHash from "./eth_getUncleCountByBlockHash";
import eth_getUncleCountByBlockNumber from "./eth_getUncleCountByBlockNumber";
import eth_getProof from "./eth_getProof";
import eth_getWork from "./eth_getWork";
import eth_hashrate from "./eth_hashrate";
import eth_mining from "./eth_mining";
import eth_newBlockFilter from "./eth_newBlockFilter";
import eth_newFilter from "./eth_newFilter";
import eth_newPendingTransactionFilter from "./eth_newPendingTransactionFilter";
import eth_pendingTransactions from "./eth_pendingTransactions";
import eth_protocolVersion from "./eth_protocolVersion";
import eth_sendRawTransaction from "./eth_sendRawTransaction";
import eth_submitHashrate from "./eth_submitHashrate";
import eth_submitWork from "./eth_submitWork";
import eth_syncing from "./eth_syncing";
import eth_uninstallFilter from "./eth_uninstallFilter";

const methods = {
  web3_clientVersion,
  web3_sha3,
  net_listening,
  net_peerCount,
  net_version,
  eth_blockNumber,
  eth_call,
  eth_chainId,
  eth_coinbase,
  eth_estimateGas,
  eth_gasPrice,
  eth_getBalance,
  eth_getBlockByHash,
  eth_getBlockByNumber,
  eth_getBlockTransactionCountByHash,
  eth_getBlockTransactionCountByNumber,
  eth_getCode,
  eth_getFilterChanges,
  eth_getFilterLogs,
  eth_getRawTransactionByHash,
  eth_getRawTransactionByBlockHashAndIndex,
  eth_getRawTransactionByBlockNumberAndIndex,
  eth_getLogs,
  eth_getStorageAt,
  eth_getTransactionByBlockHashAndIndex,
  eth_getTransactionByBlockNumberAndIndex,
  eth_getTransactionByHash,
  eth_getTransactionCount,
  eth_getTransactionReceipt,
  eth_getUncleByBlockHashAndIndex,
  eth_getUncleByBlockNumberAndIndex,
  eth_getUncleCountByBlockHash,
  eth_getUncleCountByBlockNumber,
  eth_getProof,
  eth_getWork,
  eth_hashrate,
  eth_mining,
  eth_newBlockFilter,
  eth_newFilter,
  eth_newPendingTransactionFilter,
  eth_pendingTransactions,
  eth_protocolVersion,
  eth_sendRawTransaction,
  eth_submitHashrate,
  eth_submitWork,
  eth_syncing,
  eth_uninstallFilter,
};

export default methods;
