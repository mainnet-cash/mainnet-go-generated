/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchTransactionReceipt struct for SmartBchTransactionReceipt
type SmartBchTransactionReceipt struct {
	To string `json:"to"`
	From string `json:"from"`
	ContractAddress string `json:"contractAddress"`
	TransactionIndex float32 `json:"transactionIndex"`
	Root string `json:"root,omitempty"`
	GasUsed string `json:"gasUsed"`
	LogsBloom string `json:"logsBloom"`
	BlockHash string `json:"blockHash"`
	TransactionHash string `json:"transactionHash"`
	Logs []interface{} `json:"logs"`
	BlockNumber float32 `json:"blockNumber"`
	Confirmations float32 `json:"confirmations"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	Byzantium bool `json:"byzantium"`
	Type float32 `json:"type"`
	Status float32 `json:"status,omitempty"`
}
