package model

type DeployResponse struct {
	TxID string `json:"txId"`
}

type TransactionResponse struct {
	BlockHash         string `json:"blockHash"`
	Status            bool   `json:"status"`
	BlockNumber       int    `json:"blockNumber"`
	From              string `json:"from"`
	Gas               int    `json:"gas"`
	GasPrice          int    `json:"gasPrice"`
	TransactionHash   string `json:"transactionHash"`
	Input             string `json:"input"`
	Nonce             int    `json:"nonce"`
	To                string `json:"to"`
	TransactionIndex  int    `json:"transactionIndex"`
	Value             string `json:"value"`
	GasUsed           int    `json:"gasUsed"`
	CumulativeGasUsed int    `json:"cumulativeGasUsed"`
	ContractAddress   string `json:"contractAddress"`
	Logs              []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		LogIndex         int      `json:"logIndex"`
		BlockNumber      int      `json:"blockNumber"`
		BlockHash        string   `json:"blockHash"`
		TransactionIndex int      `json:"transactionIndex"`
		TransactionHash  string   `json:"transactionHash"`
	} `json:"logs"`
}
