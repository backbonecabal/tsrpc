import { EthSubmitWork } from "../generated-typings";

const eth_submitWork: EthSubmitWork = (nonce, powHash, mixHash) => {
  return Promise.resolve(false);
};

export default eth_submitWork;
